package istio

import (
	"errors"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
	"strings"
)

type IstioCompleter struct {
	IstioRuntime *environments.Runtime
}

func NewCompleter() (*IstioCompleter, error) {
	// only support Istio over Kubernetes
	// so that must already installed Kubernetes Env
	if _, err := kube.NewCompleter(); err != nil {
		return nil, err
	}
	completer := &IstioCompleter{}
	// check if istioctl is valid in PATH
	code, msg := environments.ExecuteAndGetResult("echo", "version")
	if code == 0 {
		return completer, nil
	} else {
		return nil, errors.New(msg)
	}
}

func (c *IstioCompleter) Complete(d prompt.Document) []prompt.Suggest {
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
		return optionCompleter(args)
	}

	l := len(args)
	if l > 2 {
		switch args[l-2] {
		case "--meshConfigFile", "--injectConfigFile", "--filename", "--file", "--valuesFile", "--config-path", "-f":
			return environments.GetPathSuggestion(args[l-1])
		case "--istioNamespace", "--namespace", "--operatorNamespace", "--watchedNamespaces", "-n":
			return kube.GetNameSpaceSuggestions(kubernetes.Completer.NamespaceList)
		}
	}

	return argumentsCompleter(c.IstioRuntime, kubernetes.Completer.Namespace, args, w)
}

func optionCompleter(args []string) []prompt.Suggest {
	var suggests []prompt.Suggest
	l := len(args)
	if l <= 1 {
		suggests = optionGlobal
	}

	switch args[0] {
	case "analyze":
		suggests = optionAnalyze
	case "admin":
		if len(args) >= 3 {
			switch args[1] {
			case "log":
				suggests = optionAdminLog
			}
		} else {
			suggests = optionAdmin
		}
	case "bug-report":
		suggests = optionBugReport
	case "dashboard":
		if len(args) >= 3 {
			switch args[1] {
			case "controlz":
				suggests = optionDashboardControlz
			case "envoy":
				suggests = optionDashboardEnvoy
			case "grafana", "jaeger", "kiali", "prometheus", "zipkin":
				suggests = optionDashboard
			default:
				suggests = optionDashboard
			}
		} else {
			suggests = optionDashboard
		}
	case "install":
		suggests = optionInstall
	case "kube-inject":
		suggests = optionKubeInject
	case "manifest":
		if len(args) >= 3 {
			switch args[1] {
			case "diff":
				suggests = optionManifestDiff
			case "generate":
				suggests = optionManifestGenerate
			case "install":
				suggests = optionManifestInstall
			}
		} else {
			suggests = optionManifest
		}
	case "operator":
		if len(args) >= 3 {
			switch args[1] {
			case "dump":
				suggests = optionOperatorDump
			case "init":
				suggests = optionOperatorInit
			case "remove":
				suggests = optionOperatorRemove
			}
		} else {
			suggests = optionOperator
		}
	case "options":
		suggests = optionOptions
	case "profile":
		if len(args) >= 3 {
			switch args[1] {
			case "diff":
				suggests = optionProfileDiff
			case "dump":
				suggests = optionProfileDump
			case "list":
				suggests = optionProfileList
			}
		} else {
			suggests = optionProfile
		}
	case "proxy-config":
		if len(args) >= 3 {
			switch args[1] {
			case "all":
				suggests = optionProxyConfigAll
			case "bootstrap":
				suggests = optionProxyConfigBootstrap
			case "cluster":
				suggests = optionProxyConfigCluster
			case "endpoint":
				suggests = optionProxyConfigEndpoint
			case "listener":
				suggests = optionProxyConfigListener
			case "log":
				suggests = optionProxyConfigLog
			case "route":
				suggests = optionProxyConfigRoute
			case "secret":
				suggests = optionProxyConfigSecret
			}
		} else {
			suggests = optionProxyConfig
		}
	case "proxy-status":
		suggests = optionProxyStatus
	case "upgrade":
		suggests = optionUpgrade
	case "validate":
		suggests = optionValidate
	case "verify-install":
		suggests = optionVerifyInstall
	case "version":
		suggests = optionVersion
	default:
		suggests = optionGlobal
	}

	return prompt.FilterContains(suggests, args[l-1], false)
}

func getPodSuggestion(namespace string) []prompt.Suggest {
	return kube.GetPodSuggestions(kubernetes.Completer.Client, namespace)
}

func argumentsCompleter(istio *environments.Runtime, namespace string, args []string, currentArg string) []prompt.Suggest {
	if len(args) == 0 {
		return istio.MainSuggestion
	} else if len(args) == 1 {
		return prompt.FilterHasPrefix(istio.MainSuggestion, args[0], true)
	}

	first := args[0]
	switch first {
	case "admin":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "log", Alias: "l", Description: "Retrieve or update logging levels of istiod components."},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		} else if len(args) == 3 {
			return prompt.FilterFuzzy(getPodSuggestion(namespace), currentArg, true)
		}
	case "analyze":
		return environments.GetPathSuggestion(currentArg)
	case "dashboard":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "controlz", Description: "Open the ControlZ web UI for a pod in the Istio control plane"},
				{Text: "envoy", Description: "Open the Envoy admin dashboard for a sidecar"},
				{Text: "grafana", Description: "Open Istio's Grafana dashboard"},
				{Text: "jaeger", Description: "Open Istio's Jaeger dashboard"},
				{Text: "kiali", Description: "Open Istio's Kiali dashboard"},
				{Text: "prometheus", Description: "Open Istio's Prometheus dashboard"},
				{Text: "zipkin", Description: "Open Istio's Zipkin dashboard"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "experimental":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "add-to-mesh", Description: "restarts pods with an Istio sidecar or configures meshed pod access to external services"},
				{Text: "authz", Description: "THIS COMMAND IS UNDER ACTIVE DEVELOPMENT AND NOT READY FOR PRODUCTION USE"},
				{Text: "config", Description: "Configure istioctl defaults"},
				{Text: "create-remote-secret", Description: "Create a secret with credentials to allow Istio to access remote Kubernetes apiservers"},
				{Text: "describe", Description: "Describe resource and related Istio configuration"},
				{Text: "injector", Description: "List sidecar injector and sidecar versions"},
				{Text: "kube-uninject", Description: "kube-uninject is used to prevent Istio from adding a sidecar and also provides the inverse of `istioctl kube-inject -f`"},
				{Text: "metrics", Description: "Prints the metrics for the specified service(s) when running in Kubernetes"},
				{Text: "precheck", Description: "precheck inspects a Kubernetes cluster for Istio install and upgrade requirements."},
				{Text: "proxy-status", Description: "Retrieves last sent and last acknowledged xDS sync from Istiod to each Envoy in the mesh"},
				{Text: "remove-from-mesh", Description: "istioctl experimental remove-from-mesh' restarts pods without an Istio sidecar or removes external service access configuration"},
				{Text: "revision", Description: "The revision command provides a revision centric view of istio deployments"},
				{Text: "uninstall", Description: "The uninstall command uninstalls Istio from a cluster"},
				{Text: "wait", Description: "Waits for the specified condition to be true of an Istio resource"},
				{Text: "workload", Description: "Commands to assist in configuring and deploying workloads running on VMs and other non-Kubernetes environments"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "manifest":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "diff", Description: "The diff subcommand compares manifests from two files or directories"},
				{Text: "generate", Description: "The generate subcommand generates an Istio install manifest and outputs to the console by default"},
				{Text: "install", Description: "The install command generates an Istio install manifest and applies it to a cluster"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		} else if len(args) == 3 && args[1] == "diff" {
			return environments.GetPathSuggestion(currentArg)
		}
	case "operator":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "dump", Description: "The dump subcommand dumps the Istio operator controller manifest"},
				{Text: "init", Description: "The init subcommand installs the Istio operator controller in the cluster"},
				{Text: "remove", Description: "The remove subcommand removes the Istio operator controller from the cluster"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "profile":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "diff", Description: "The diff subcommand displays the differences between two Istio configuration profiles"},
				{Text: "dump", Description: "The dump subcommand dumps the values in an Istio configuration profile"},
				{Text: "list", Description: "The list subcommand lists the available Istio configuration profiles"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		} else if len(args) == 3 && args[1] == "diff" {
			return environments.GetPathSuggestion(currentArg)
		}
	case "proxy-config":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "all", Description: "Retrieve information about all configuration for the Envoy instance in the specified pod"},
				{Text: "bootstrap", Description: "Retrieve information about bootstrap configuration for the Envoy instance in the specified pod"},
				{Text: "cluster", Description: "Retrieve information about cluster configuration for the Envoy instance in the specified"},
				{Text: "endpoint", Description: "Retrieve information about endpoint configuration for the Envoy instance in the specified pod"},
				{Text: "listener", Description: "Retrieve information about listener configuration for the Envoy instance in the specified pod"},
				{Text: "log", Description: "Retrieve information about logging levels of the Envoy instance in the specified pod, and update optionally"},
				{Text: "route", Description: "Retrieve information about route configuration for the Envoy instance in the specified pod"},
				{Text: "secret", Description: "Retrieve information about secret configuration for the Envoy instance in the specified pod"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		} else {
			return getPodSuggestion(namespace)
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
