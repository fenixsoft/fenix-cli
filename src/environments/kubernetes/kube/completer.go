package kube

import (
	"context"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewCompleter() (*Completer, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.WarnIfAllMissing = false
	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	config, err := loader.ClientConfig()
	if err != nil {
		return nil, err
	}

	namespace, _, err := loader.Namespace()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		if statusError, ok := err.(*errors.StatusError); ok && statusError.Status().Code == 403 {
			namespaces = nil
		} else {
			return nil, err
		}
	}

	return &Completer{
		Namespace:     namespace,
		NamespaceList: namespaces,
		Client:        client,
	}, nil
}

type Completer struct {
	Namespace         string
	NamespaceList     *corev1.NamespaceList
	Client            *kubernetes.Clientset
	KubernetesRuntime *environments.Runtime
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}

	// Return suggestions for option
	if suggests, found := c.completeOptionArguments(d); found {
		return suggests
	}

	namespace := checkNamespaceArg(d)
	if namespace == "" {
		namespace = c.Namespace
	}
	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		return optionValueCompleter(args, w)
	}
	return c.argumentsCompleter(namespace, commandArgs)
}

func optionValueCompleter(args []string, currentArg string) []prompt.Suggest {
	l := len(args)
	var suggest []prompt.Suggest
	if l < 2 {
		return suggest
	}
	option := args[l-2]

	switch option {
	case "--context":
		suggest = GetContextSuggestions()
	case "--output", "-o":
		suggest = []prompt.Suggest{{Text: "json"}, {Text: "yaml"}, {Text: "table"}, {Text: "short"}}
	case "--cascade":
		suggest = []prompt.Suggest{{Text: "background"}, {Text: "orphan"}, {Text: "foreground"}}
	case "--restart":
		suggest = []prompt.Suggest{{Text: "Always"}, {Text: "OnFailure"}, {Text: "Never"}}
	case "--session-affinity":
		suggest = []prompt.Suggest{{Text: "None"}, {Text: "ClientIP"}}
	}

	return prompt.FilterContains(suggest, currentArg, true)
}

func checkNamespaceArg(d prompt.Document) string {
	args := strings.Split(d.Text, " ")
	var found bool
	for i := 0; i < len(args); i++ {
		if found {
			return args[i]
		}
		if args[i] == "--namespace" || args[i] == "-n" {
			found = true
			continue
		}
	}
	return ""
}

/* Option arguments */

var yamlFileCompleter = completer.FilePathCompleter{
	IgnoreCase: true,
	Filter: func(fi os.FileInfo) bool {
		if fi.IsDir() {
			return true
		}
		if strings.HasSuffix(fi.Name(), ".yaml") || strings.HasSuffix(fi.Name(), ".yml") {
			return true
		}
		return false
	},
}

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func GetPathSuggestion(path string) []prompt.Suggest {
	if !strings.HasSuffix(path, "*") {
		path = path + "*"
	}
	files, _ := filepath.Glob(path)
	var ret []prompt.Suggest
	for i, file := range files {
		if i > 16 {
			return ret
		} else {
			ret = append(ret, prompt.Suggest{Text: file})
		}
	}
	return ret
}

func (c *Completer) completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}

	// namespace
	if option == "-n" || option == "--namespace" {
		return prompt.FilterHasPrefix(
			GetNameSpaceSuggestions(c),
			d.GetWordBeforeCursor(),
			true,
		), true
	}

	// filename
	switch cmd {
	case "get", "describe", "create", "delete", "replace", "patch",
		"edit", "apply", "expose", "rolling-update", "rollout",
		"label", "annotate", "scale", "convert", "autoscale", "top":
		if option == "-f" || option == "--filename" {
			// return yamlFileCompleter.Complete(d), true
			return GetPathSuggestion(d.GetWordBeforeCursor()), true
		}
	}

	// container
	switch cmd {
	case "exec", "logs", "run", "attach", "port-forward", "cp":
		if option == "-c" || option == "--container" {
			cmdArgs := getCommandArgs(d)
			var suggestions []prompt.Suggest
			if cmdArgs == nil || len(cmdArgs) < 2 {
				suggestions = getContainerNamesFromCachedPods(c.Client, c.Namespace)
			} else {
				suggestions = getContainerName(c.Client, c.Namespace, cmdArgs[1])
			}
			return prompt.FilterHasPrefix(
				suggestions,
				d.GetWordBeforeCursor(),
				true,
			), true
		}
	}
	return []prompt.Suggest{}, false
}

func getCommandArgs(d prompt.Document) []string {
	args := strings.Split(d.TextBeforeCursor(), " ")

	// If PIPE is in text before the cursor, returns empty.
	for i := range args {
		if args[i] == "|" {
			return nil
		}
	}

	commandArgs, _ := excludeOptions(args)
	return commandArgs
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	cmd := args[0]
	filtered := make([]string, 0, l)

	var skipNextArg bool
	for i := 0; i < len(args)-1; i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		// ignore first or continuous blank spaces
		if (i == 0 && args[0] == "") || (args[i] == "" && args[i-1] == "") {
			continue
		}

		if cmd == "logs" && args[i] == "-f" {
			continue
		}

		for _, s := range []string{
			"-f", "--filename", "-n", "--namespace", "-s", "--server", "--kubeconfig", "--cluster", "--user",
			"-o", "--output", "-c", "--container", "--clusterrole", "--role", "--clusterip", "--external-name",
			"--tcp", "--template", "--field-manager", "--chunk-size", "--field-selector", "--selector", "--sort-by",
			"--serviceaccount", "--aggregation-rule", "--rule", "--verb", "--from-file", "--from-literal", "--from",
			"--image", "--port", "--replicas", "--annotation", "--max-unavailable", "--min-available", "--description",
			"--preemption-policy", "--resource-name", "--resource", "--docker-server", "--docker-password",
			"--external-name", "--allow-missing-template-keys", "--chunk-size", "--label-columns", "--kustomize",
			"--raw", "--sort-by", "--cascade", "--envc", "--grace-period", "--hostport", "--image-pull-policy",
			"--limits", "--overrides", "--pod-running-timeout", "--requests", "--restart", "--timeout",
			"--container-port", "--target-port", "--external-ip", "--generator", "--protocol", "--session-affinity",
			"--type", "--resource-version", "--cpu-percent", "--max", "--min", "--copy-to",
		} {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		filtered = append(filtered, args[i])
	}
	filtered = append(filtered, args[len(args)-1])

	return filtered, skipNextArg
}
