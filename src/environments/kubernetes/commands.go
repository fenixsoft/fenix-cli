package kubernetes

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"io"
	k8s "k8s.io/client-go/kubernetes"
	"strings"
)

var executor = func(args string) {
	environments.GetKubernetes().Executor(args)
}

var ExtraCommands = []environments.Command{
	{
		Text:         "x-context",
		Description:  "Select the cluster for kubernetes management",
		Environments: []environments.Environment{environments.Kubernetes, environments.Istio},
		MatchFn:      environments.IgnoreCaseMatch,
		Fn: func(args []string, writer io.Writer) {
			// refresh for new namespace
			GetNameSpaceSuggestions(Client)

			var ctx []string
			sug := GetContextSuggestions()
			for _, s := range sug {
				ctx = append(ctx, s.Text)
			}

			var ret string
			pt := &survey.Select{
				Message: "Select active cluster context: ",
				Options: ctx,
			}
			if err := survey.AskOne(pt, &ret); err != nil {
				_, _ = writer.Write([]byte(err.Error() + "\n"))
			} else {
				executor("config use-context " + ret)
			}
		},
	},
	{
		Text:         "x-namespace",
		Alias:        "x-ns",
		Provider:     Namespace,
		Description:  "Select the namespace for kubernetes management",
		Environments: []environments.Environment{environments.Kubernetes, environments.Istio},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			var ret string
			if len(args) < 2 {
				// refresh for new namespace
				GetNameSpaceSuggestions(Client)

				ns := make([]string, len(Client.NamespaceList.Items))
				for i := range Client.NamespaceList.Items {
					ns[i] = Client.NamespaceList.Items[i].Name
				}

				pt := &survey.Select{
					Message: "Select active namespace: ",
					Options: ns,
				}
				if err := survey.AskOne(pt, &ret); err != nil {
					_, _ = writer.Write([]byte(err.Error() + "\n"))
					return
				}
			} else {
				ret = args[1]
			}

			// execute : kubectl config set-context --current --namespace=<NS>
			executor("config set-context --current --namespace=" + ret)
			Client.Namespace = ret

		},
	},
	{
		Text:         "x-batch",
		Provider:     Resource,
		Description:  "Batch management of kubernetes resources",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			defaultOp := []string{"delete", "describe", "edit"}
			resources := map[string]struct {
				fn  func(*k8s.Clientset, string) []prompt.Suggest
				ops []string
			}{
				"componentstatuses": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"cc": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"namespaces": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetNameSpaceSuggestions(Client)
					},
					ops: defaultOp,
				},
				"ns": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetNameSpaceSuggestions(Client)
					},
					ops: defaultOp,
				},
				"nodes": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"no": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"persistentvolumes": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"pv": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"podsecuritypolicies": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetPodSecurityPolicySuggestions(c)
					},
					ops: defaultOp,
				},
				"psp": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return GetPodSecurityPolicySuggestions(c)
					},
					ops: defaultOp,
				},
				"configmaps":             {fn: GetConfigMapSuggestions, ops: defaultOp},
				"cm":                     {fn: GetConfigMapSuggestions, ops: defaultOp},
				"daemonsets":             {fn: GetDaemonSetSuggestions, ops: defaultOp},
				"ds":                     {fn: GetDaemonSetSuggestions, ops: defaultOp},
				"deployments":            {fn: GetDeploymentSuggestions, ops: defaultOp},
				"deploy":                 {fn: GetDeploymentSuggestions, ops: defaultOp},
				"endpoints":              {fn: GetEndpointsSuggestions, ops: defaultOp},
				"ep":                     {fn: GetEndpointsSuggestions, ops: defaultOp},
				"ingresses":              {fn: GetIngressSuggestions, ops: defaultOp},
				"ing":                    {fn: GetIngressSuggestions, ops: defaultOp},
				"limitranges":            {fn: GetLimitRangeSuggestions, ops: defaultOp},
				"limits":                 {fn: GetLimitRangeSuggestions, ops: defaultOp},
				"pods":                   {fn: GetPodSuggestions, ops: defaultOp},
				"po":                     {fn: GetPodSuggestions, ops: defaultOp},
				"pod":                    {fn: GetPodSuggestions, ops: defaultOp},
				"persistentvolumeclaims": {fn: GetPersistentVolumeClaimSuggestions, ops: defaultOp},
				"pvc":                    {fn: GetPersistentVolumeClaimSuggestions, ops: defaultOp},
				"podtemplates":           {fn: GetPodTemplateSuggestions, ops: defaultOp},
				"replicasets":            {fn: GetReplicaSetSuggestions, ops: defaultOp},
				"rs":                     {fn: GetReplicaSetSuggestions, ops: defaultOp},
				"replicationcontrollers": {fn: GetReplicationControllerSuggestions, ops: defaultOp},
				"rc":                     {fn: GetReplicationControllerSuggestions, ops: defaultOp},
				"resourcequotas":         {fn: GetResourceQuotasSuggestions, ops: defaultOp},
				"quota":                  {fn: GetResourceQuotasSuggestions, ops: defaultOp},
				"secrets":                {fn: GetSecretSuggestions, ops: defaultOp},
				"serviceaccounts":        {fn: GetServiceAccountSuggestions, ops: defaultOp},
				"sa":                     {fn: GetServiceAccountSuggestions, ops: defaultOp},
				"services":               {fn: GetServiceSuggestions, ops: defaultOp},
				"svc":                    {fn: GetServiceSuggestions, ops: defaultOp},
				"jobs":                   {fn: GetJobSuggestions, ops: defaultOp},
				"job":                    {fn: GetJobSuggestions, ops: defaultOp},
			}

			if len(args) < 2 {
				_, _ = writer.Write([]byte("The resource type should be provided in the parameter\n" +
					"Try \"x-batch pod\" or \"x-batch deployment\" again\n"))
				return
			}
			resType := strings.ToLower(args[1])
			if _, ok := resources[resType]; !ok {
				_, _ = writer.Write([]byte("The resource type should be provided in the parameter\n" +
					"Try \"x-batch pod\" or \"x-batch deployment\" again\n"))
				return
			}

			resource := resources[resType]
			var opts []string
			col := int(util.GetWindowWidth() - 15)
			for _, v := range resource.fn(Client.Set, Client.Namespace) {
				opts = append(opts, v.Text+" | "+util.SubString(v.Description, 0, col-len(v.Text)))
			}
			var qs = []*survey.Question{
				{
					Name: "objective",
					Prompt: &survey.MultiSelect{
						Message:  "Which resource you want to operate ?",
						Options:  opts,
						PageSize: 25,
					},
					Validate: func(val interface{}) error {
						if ans, ok := val.([]survey.OptionAnswer); !ok || len(ans) == 0 {
							return errors.New("please select a least one resource")
						}
						return nil
					},
				},
				{
					Name: "operation",
					Prompt: &survey.Select{
						Message: "What you want to do with the these resources ?",
						Options: resource.ops,
					},
				},
			}
			answers := struct {
				Objective []string
				Operation string
				Confirm   bool
			}{}

			err := survey.Ask(qs, &answers, survey.WithKeepFilter(true))
			if err == terminal.InterruptErr {
				_, _ = writer.Write([]byte("operation interrupted\n"))
			} else if err != nil {
				if err.Error() == "please provide options to select from" {
					_, _ = writer.Write([]byte("there is no " + resType + " resource here, do you forgot setting namespace? \n"))
				} else {
					_, _ = writer.Write([]byte(err.Error() + "\n"))
				}
			}

			if answers.Operation == "delete" {
				pt := &survey.Confirm{
					Message: "Are you sure you want to continue ?",
				}
				_ = survey.AskOne(pt, &answers.Confirm)
			} else {
				answers.Confirm = true
			}
			if answers.Confirm {
				for _, o := range answers.Objective {
					executor(answers.Operation + " " + resType + " " + strings.TrimSpace(strings.Split(o, "|")[0]))
				}
			}
		},
	},
	{
		Text:         "x-sniff",
		Provider:     Pod,
		Description:  "Get a capture of the network activity between services.",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			if len(args) <= 1 || args[1] == "" {
				_, _ = writer.Write([]byte("Provide an available pod name for x-sniff command\n"))
				return
			}
			if !krew.IsTSharkAvailable() {
				_, _ = writer.Write([]byte("" +
					"TShark is not available, install it first\n" +
					"eg: " +
					"sudo add-apt-repository ppa:wireshark-dev/stable\n" +
					"sudo apt install tshark\n"))
				return
			}
			pt := &survey.Select{
				Message: "Which type you want to start capture it's traffic ?",
				Options: []string{"summary", "detail"},
			}
			var op string
			_ = survey.AskOne(pt, &op)
			if op == "summary" {
				op = " -p -f \"port 80\" -o - | tshark -Y http --export-objects \"http,data\" -r -"
			} else {
				op = " -p -f \"port 80\" -o - | tshark -Y http -V -r -"
			}
			krew.CheckAndInstall("sniff")
			executor("sniff " + args[1] + op)
		},
	},
	{
		Text:         "x-lens",
		Provider:     Pod,
		Description:  "Show pod-related resource information.",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			if len(args) <= 1 || args[1] == "" {
				_, _ = writer.Write([]byte("Provide an available pod name for x-lens command\n"))
				return
			}
			krew.CheckAndInstall("pod-lens")
			executor("pod-lens " + args[1])
		},
	},
	{
		Text:         "x-status",
		Provider:     Resource,
		Description:  "Print a human-friendly output that focuses on the status fields of the resources in kubernetes.",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			if len(args) <= 1 || args[1] == "" {
				_, _ = writer.Write([]byte("Provide an available resource type for x-status command\n"))
				return
			}

			krew.CheckAndInstall("status")
			executor("status " + args[1])
		},
	},
	{
		Text:         "x-open",
		Provider:     Service,
		Description:  "Open the kubernetes url for the specified service in your browser.",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			if len(args) <= 1 || args[1] == "" {
				_, _ = writer.Write([]byte("Provide an available service name for x-open command\n"))
				return
			}
			if !krew.IsXDGAvailable() {
				_, _ = writer.Write([]byte("" +
					"xdg-open is not available, install it first\n" +
					"eg: " +
					"sudo apt-get install -y xdg-utils\n"))
				return
			}

			krew.CheckAndInstall("open-svc")
			executor("open-svc " + strings.Join(args[1:], " "))
		},
	},
}
