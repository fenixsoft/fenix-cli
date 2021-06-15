package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "get":
		suggests = getOptions
	case "describe":
		suggests = describeOptions
	case "create":
		suggests = createOptions
	case "replace":
		suggests = replaceOptions
	case "patch":
		suggests = patchOptions
	case "delete":
		suggests = deleteOptions
	case "edit":
		suggests = editOptions
	case "apply":
		suggests = applyOptions
	case "logs":
		suggests = logsOptions
	case "rolling-update":
		suggests = rollingUpdateOptions
	case "scale", "resize":
		suggests = scaleOptions
	case "attach":
		suggests = attachOptions
	case "exec":
		suggests = execOptions
	case "port-forward":
		suggests = portForwardOptions
	case "proxy":
		suggests = proxyOptions
	case "run", "run-container":
		suggests = runOptions
	case "expose":
		suggests = exposeOptions
	case "autoscale":
		suggests = autoscaleOptions
	case "rollout":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "history":
				suggests = rolloutHistoryOptions
			case "pause":
				suggests = rolloutPauseOptions
			case "resume":
				suggests = rolloutResumeOptions
			case "status":
				suggests = rolloutStatusOptions
			case "undo":
				suggests = rolloutUndoOptions
			}
		}
	case "label":
		suggests = labelOptions
	case "cluster-info":
		suggests = clusterInfoOptions
	case "explain":
		suggests = explainOptions
	case "cordon":
		suggests = cordonOptions
	case "drain":
		suggests = drainOptions
	case "uncordon":
		suggests = uncordonOptions
	case "annotate":
		suggests = annotateOptions
	case "convert":
		suggests = convertOptions
	case "top":
		if len(commandArgs) >= 2 {
			switch commandArgs[1] {
			case "no", "node", "nodes":
				suggests = topNodeOptions
			case "po", "pod", "pods":
				suggests = topPodOptions
			}
		}
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "get-contexts":
				suggests = configGetContextsOptions
			case "view":
				suggests = configViewOptions
			case "set-cluster":
				suggests = configSetClusterOptions
			case "set-credentials":
				suggests = configSetCredentialsOptions
			case "set":
				suggests = configSetOptions
			}
		}
	default:
		suggests = globalOptions
	}

	suggests = append(suggests, globalOptions...)
	return prompt.FilterContains(suggests, args[len(args)-1], true)
}

var globalOptions = []prompt.Suggest{
	{Text: "--namespace", Alias: "-n", Description: "Temporarily set the namespace for a request"},
	{Text: "--server", Alias: "-s", Description: "Specify the address and port of the Kubernetes API server"},
	{Text: "--user", Description: "Take the user if this flag exists."},
	{Text: "--cluster", Description: "Take the cluster if this flag exists."},
	{Text: "--help", Description: "Show helo information."},
}
