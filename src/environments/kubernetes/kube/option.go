package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	if len(commandArgs) == 0 {
		return []prompt.Suggest{}
	}
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
	case "x-open":
		suggests = []prompt.Suggest{
			{Text: "--add_dir_header", Description: "If true, adds the file directory to the header of the log messages"},
			{Text: "--alsologtostderr", Description: "log to standard error as well as files"},
			{Text: "--as", Description: "Username to impersonate for the operation"},
			{Text: "--as-group", Description: "Group to impersonate for the operation, this flag can be repeated to specify multiple groups."},
			{Text: "--cache-dir", Description: "Default cache directory (default \"/root/.kube/cache\")"},
			{Text: "--certificate-authority", Description: "Path to a cert file for the certificate authority"},
			{Text: "--client-certificate", Description: "Path to a client certificate file for TLS"},
			{Text: "--client-key", Description: "Path to a client key file for TLS"},
			{Text: "--cluster", Description: "The name of the kubeconfig cluster to use"},
			{Text: "--insecure-skip-tls-verify", Description: "If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure"},
			{Text: "--log_backtrace_at", Description: "When logging hits line file:N, emit a stack trace (default :0)"},
			{Text: "--log_dir", Description: "If non-empty, write log files in this directory"},
			{Text: "--log_file", Description: "If non-empty, use this log file"},
			{Text: "--log_file_max_size", Description: "Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited. (default 1800)"},
			{Text: "--logtostderr", Description: "Log to standard error instead of files (default true)"},
			{Text: "--one_output", Description: "If true, only write logs to their native severity level (vs also writing to each lower severity level)"},
			{Text: "--request-timeout", Description: "The length of time to wait before giving up on a single server request"},
			{Text: "--scheme", Description: "The scheme for connections between the apiserver and the service. It must be \"http\" or \"https\" if specfied."},
			{Text: "--server", Description: "The address and port of the Kubernetes API server"},
			{Text: "--skip_headers", Description: "If true, avoid header prefixes in the log messages"},
			{Text: "--skip_log_headers", Description: "If true, avoid headers when opening log files"},
			{Text: "--stderrthreshold", Description: "Logs at or above this threshold go to stderr (default 2)"},
			{Text: "--tls-server-name", Description: "Server name to use for server certificate validation"},
			{Text: "--token", Description: "Bearer token for authentication to the API server"},
			{Text: "--user", Description: "The name of the kubeconfig user to use"},
			{Text: "--v", Alias: "-v", Description: "Number for the log level verbosity"},
			{Text: "--vmodule", Description: "Comma-separated list of pattern=N settings for file-filtered logging"},
			{Text: "--port", Alias: "-p", Description: "The port on which to run the proxy. Set to 0 to pick a random port. (default 8001)"},
			{Text: "--address", Description: "The IP address on which to serve on. (default \"127.0.0.1\")"},
			{Text: "--keepalive", Description: "keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keepalive."},
		}
	default:
		suggests = globalOptions
	}

	suggests = append(suggests, globalOptions...)
	return prompt.FilterHasPrefix(suggests, args[len(args)-1], true)
}

var globalOptions = []prompt.Suggest{
	{Text: "--namespace", Alias: "-n", Description: "Temporarily set the namespace for a request"},
	{Text: "--server", Alias: "-s", Description: "Specify the address and port of the Kubernetes API server"},
	{Text: "--user", Description: "Take the user if this flag exists."},
	{Text: "--cluster", Description: "Take the cluster if this flag exists."},
	{Text: "--help", Description: "Show helo information."},
}
