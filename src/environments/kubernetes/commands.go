package kubernetes

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"io"
	"k8s.io/client-go/kubernetes"
	"strings"
)

var ExtraCommands = []environments.Command{
	{
		Text:         "x-namespace",
		Description:  "Select a namespace for current kubernetes management",
		Environments: []environments.Environment{environments.Kubernetes, environments.Istio},
		MatchFn:      environments.IgnoreCaseMatch,
		Fn: func(args []string, writer io.Writer) {
			// refresh for new namespace
			if c, err := kube.NewCompleter(); err != nil {
				Completer = c
			}

			ns := make([]string, len(Completer.NamespaceList.Items))
			for i := range Completer.NamespaceList.Items {
				ns[i] = Completer.NamespaceList.Items[i].Name
			}

			var ret string
			pt := &survey.Select{
				Message: "Select active namespace: ",
				Options: ns,
			}
			util.AssertNoError(survey.AskOne(pt, &ret))

			// execute : kubectl config set-context --current --namespace=<NS>
			environments.Executor("kubectl", "config set-context --current --namespace="+ret)
			Completer.Namespace = ret
		},
	},
	{
		Text:         "x-batch",
		Description:  "Batch management of kubernetes resources",
		Environments: []environments.Environment{environments.Kubernetes},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			defaultOp := []string{"delete", "describe", "edit"}
			resources := map[string]struct {
				fn  func(*kubernetes.Clientset, string) []prompt.Suggest
				ops []string
			}{
				"componentstatuses": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"cc": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetComponentStatusCompletions(c)
					},
					ops: defaultOp,
				},
				"namespaces": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetNameSpaceSuggestions(Completer.NamespaceList)
					},
					ops: defaultOp,
				},
				"ns": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetNameSpaceSuggestions(Completer.NamespaceList)
					},
					ops: defaultOp,
				},
				"nodes": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"no": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetNodeSuggestions(c)
					},
					ops: defaultOp,
				},
				"persistentvolumes": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"pv": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetPersistentVolumeSuggestions(c)
					},
					ops: defaultOp,
				},
				"podsecuritypolicies": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
						return kube.GetPodSecurityPolicySuggestions(c)
					},
					ops: defaultOp,
				},
				"psp": {
					fn: func(c *kubernetes.Clientset, s string) []prompt.Suggest {
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
					environments.Executor("kubectl", answers.Operation+" "+resType+" "+strings.TrimSpace(strings.Split(o, "|")[0]))
				}
			}
		},
	},
}
