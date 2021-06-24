package kubernetes

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"io"
	k8s "k8s.io/client-go/kubernetes"
	"strings"
)

var executor = func(args string) {
	Completer.KubernetesRuntime.Executor(args)
}

var ExtraCommands = []environments.Command{
	{
		Text:         "x-context",
		Description:  "Select the cluster for kubernetes management",
		Environments: []environments.Environment{environments.Kubernetes, environments.Istio},
		MatchFn:      environments.IgnoreCaseMatch,
		Fn: func(args []string, writer io.Writer) {
			// refresh for new namespace
			if c, err := kube.NewCompleter(); err != nil {
				Completer = c
			}

			var ctx []string
			sug := kube.GetContextSuggestions()
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
				kube.GetNameSpaceSuggestions(Completer)

				ns := make([]string, len(Completer.NamespaceList.Items))
				for i := range Completer.NamespaceList.Items {
					ns[i] = Completer.NamespaceList.Items[i].Name
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
			Completer.Namespace = ret

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
						return kube.GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"cc": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"namespaces": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetNameSpaceSuggestions(Completer)
					},
					ops: defaultOp,
				},
				"ns": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetNameSpaceSuggestions(Completer)
					},
					ops: defaultOp,
				},
				"nodes": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"no": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"persistentvolumes": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"pv": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"podsecuritypolicies": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetPodSecurityPolicySuggestions(c)
					},
					ops: defaultOp,
				},
				"psp": {
					fn: func(c *k8s.Clientset, s string) []prompt.Suggest {
						return kube.GetPodSecurityPolicySuggestions(c)
					},
					ops: defaultOp,
				},
				"configmaps":             {fn: kube.GetConfigMapSuggestions, ops: defaultOp},
				"cm":                     {fn: kube.GetConfigMapSuggestions, ops: defaultOp},
				"daemonsets":             {fn: kube.GetDaemonSetSuggestions, ops: defaultOp},
				"ds":                     {fn: kube.GetDaemonSetSuggestions, ops: defaultOp},
				"deployments":            {fn: kube.GetDeploymentSuggestions, ops: defaultOp},
				"deploy":                 {fn: kube.GetDeploymentSuggestions, ops: defaultOp},
				"endpoints":              {fn: kube.GetEndpointsSuggestions, ops: defaultOp},
				"ep":                     {fn: kube.GetEndpointsSuggestions, ops: defaultOp},
				"ingresses":              {fn: kube.GetIngressSuggestions, ops: defaultOp},
				"ing":                    {fn: kube.GetIngressSuggestions, ops: defaultOp},
				"limitranges":            {fn: kube.GetLimitRangeSuggestions, ops: defaultOp},
				"limits":                 {fn: kube.GetLimitRangeSuggestions, ops: defaultOp},
				"pods":                   {fn: kube.GetPodSuggestions, ops: defaultOp},
				"po":                     {fn: kube.GetPodSuggestions, ops: defaultOp},
				"pod":                    {fn: kube.GetPodSuggestions, ops: defaultOp},
				"persistentvolumeclaims": {fn: kube.GetPersistentVolumeClaimSuggestions, ops: defaultOp},
				"pvc":                    {fn: kube.GetPersistentVolumeClaimSuggestions, ops: defaultOp},
				"podtemplates":           {fn: kube.GetPodTemplateSuggestions, ops: defaultOp},
				"replicasets":            {fn: kube.GetReplicaSetSuggestions, ops: defaultOp},
				"rs":                     {fn: kube.GetReplicaSetSuggestions, ops: defaultOp},
				"replicationcontrollers": {fn: kube.GetReplicationControllerSuggestions, ops: defaultOp},
				"rc":                     {fn: kube.GetReplicationControllerSuggestions, ops: defaultOp},
				"resourcequotas":         {fn: kube.GetResourceQuotasSuggestions, ops: defaultOp},
				"quota":                  {fn: kube.GetResourceQuotasSuggestions, ops: defaultOp},
				"secrets":                {fn: kube.GetSecretSuggestions, ops: defaultOp},
				"serviceaccounts":        {fn: kube.GetServiceAccountSuggestions, ops: defaultOp},
				"sa":                     {fn: kube.GetServiceAccountSuggestions, ops: defaultOp},
				"services":               {fn: kube.GetServiceSuggestions, ops: defaultOp},
				"svc":                    {fn: kube.GetServiceSuggestions, ops: defaultOp},
				"jobs":                   {fn: kube.GetJobSuggestions, ops: defaultOp},
				"job":                    {fn: kube.GetJobSuggestions, ops: defaultOp},
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
			for _, v := range resource.fn(Completer.Client, Completer.Namespace) {
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
