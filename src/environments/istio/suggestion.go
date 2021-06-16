package istio

import (
	"github.com/c-bata/go-prompt"
)

var MajorCommands = []prompt.Suggest{
	{Text: "admin", Description: "A group of commands used to manage istiod configuration"},
	{Text: "analyze", Description: "Analyze Istio configuration and print validation messages"},
	{Text: "authz", Description: "Check prints the AuthorizationPolicy applied to a pod by directly checking the Envoy configuration of the pod"},
	{Text: "bug-report", Description: "bug-report selectively captures cluster information and logs into an archive to help diagnose problems"},
	{Text: "dashboard", Alias: "d", Description: "Access to Istio web UIs"},
	{Text: "install", Description: "The install command generates an Istio install manifest and applies it to a cluster"},
	{Text: "experimental", Description: "Experimental commands that may be modified or deprecated"},
	{Text: "kube-inject", Description: "kube-inject manually injects the Envoy sidecar into Kubernetes workloads"},
	{Text: "manifest", Description: "The manifest command generates and diffs Istio manifests."},
	{Text: "operator", Description: "The operator command installs, dumps, removes and shows the status of the operator controller"},
	{Text: "options", Description: "Displays istioctl global options"},
	{Text: "profile", Description: "The profile command lists, dumps or diffs Istio configuration profiles"},
	{Text: "proxy-config", Description: "A group of commands used to retrieve information about proxy configuration from the Envoy config dump"},
	{Text: "proxy-status", Description: "Retrieves last sent and last acknowledged xDS sync from Istiod to each Envoy in the mesh"},
	{Text: "upgrade", Description: "The upgrade command checks for upgrade version eligibility and, if eligible, upgrades the Istio control plane components in-place"},
	{Text: "validate", Description: "Validate Istio policy and rules files"},
	{Text: "verify-install", Description: "verify-install verifies Istio installation status against the installation file you specified when you installed Istio"},
	{Text: "version", Description: "Prints out build version information"},
}

var optionGlobal = []prompt.Suggest{
	{Text: "--context", Description: "The name of the kubeconfig context to use (default ``)"},
	{Text: "--istioNamespace", Alias: "-i", Description: "Istio system namespace (default `istio-system`)"},
	{Text: "--kubeconfig", Alias: "-c", Description: "Kubernetes configuration file (default ``)"},
	{Text: "--namespace", Alias: "-n", Description: "Config namespace (default ``)"},
}

var optionAnalyze = append([]prompt.Suggest{
	{Text: "--all-namespaces", Alias: "-A", Description: "Analyze all namespaces"},
	{Text: "--color", Description: "Default true. Disable with '=false' or set $TERM to dumb"},
	{Text: "--failure-threshold", Description: "The severity level of analysis at which to set a non-zero exit code. Valid values: [Info Warning Error] (default `Error`)"},
	{Text: "--list-analyzers", Alias: "-L", Description: "List the analyzers available to run. Suppresses normal execution."},
	{Text: "--meshConfigFile", Description: "Overrides the mesh config values to use for analysis. (default ``)"},
	{Text: "--output", Alias: "-o", Description: "Output format: one of [log json yaml] (default `log`)"},
	{Text: "--output-threshold", Description: "The severity level of analysis at which to display messages. Valid values: [Info Warning Error] (default `Info`)"},
	{Text: "--recursive", Alias: "-R", Description: "Process directory arguments recursively. Useful when you want to analyze related manifests organized within the same directory."},
	{Text: "--suppress", Alias: "-S", Description: "Suppress reporting a message code on a specific resource. Values are supplied in the form <code>=<resource>. (default `[]`)"},
	{Text: "--timeout", Description: "The duration to wait before failing (default `30s`)"},
	{Text: "--use-kube", Alias: "-k", Description: "Use live Kubernetes cluster for analysis. Set --use-kube=false to analyze files only."},
	{Text: "--verbose", Alias: "-v", Description: "Enable verbose output"},
}, optionGlobal...)

var optionAdminLog = append([]prompt.Suggest{
	{Text: "--ctrlz_port", Description: "ControlZ port (default `9876`)"},
	{Text: "--level", Description: "Comma-separated list of output logging level for scopes in format <scope>:<level>[,<scope>:<level>,...]Possible values for <level>: none, error, warn, info, debug (default ``)"},
	{Text: "--output", Description: "Output format: one of json|short (default `short`)"},
	{Text: "--reset", Description: "Reset levels to default value. (info)"},
	{Text: "--selector", Description: "label selector (default `app=istiod`)"},
	{Text: "--stack-trace-level", Description: "Comma-separated list of stack trace level for scopes in format <scope>:<stack-trace-level>[,<scope>:<stack-trace-level>,...] Possible values for <stack-trace-level>: none, error, warn, info, debug (default ``)"},
}, optionGlobal...)

var optionAdmin = append([]prompt.Suggest{
	{Text: "--selector", Description: "label selector (default `app=istiod`)"},
}, optionGlobal...)

var optionBugReport = append([]prompt.Suggest{
	{Text: "--critical-errs", Description: "List of comma separated glob patterns to match against log error strings"},
	{Text: "--dir", Description: "Set a specific directory for temporary artifact storage. (default ``)"},
	{Text: "--dry-run", Description: "Only log commands that would be run, don't fetch or write."},
	{Text: "--duration", Description: "How far to go back in time from end-time for log entries to include in the archive"},
	{Text: "--end-time", Description: "End time for the range of log entries to include in the archive"},
	{Text: "--exclude", Description: "Spec for which pod's proxy logs to exclude from the archive, after the include spec is processed"},
	{Text: "--filename", Alias: "-f", Description: "Path to a file containing configuration in YAML format"},
	{Text: "--full-secrets", Description: "If set, secret contents are included in output"},
	{Text: "--ignore-errs", Description: "List of comma separated glob patterns to match against log error strings"},
	{Text: "--include", Description: "Spec for which pod's proxy logs to include in the archive"},
	{Text: "--start-time", Description: "Start time for the range of log entries to include in the archive"},
	{Text: "--timeout", Description: "Maximum amount of time to spend fetching logs"},
}, optionGlobal...)

var optionDashboard = append([]prompt.Suggest{
	{Text: "--address", Description: "Address to listen on. Only accepts IP address or localhost as a value"},
	{Text: "--browser", Description: "When --browser is supplied as false, istioctl dashboard will not open the browser"},
	{Text: "--port", Alias: "-p", Description: "Local port to listen to (default `0`)"},
}, optionGlobal...)

var optionDashboardControlz = append([]prompt.Suggest{
	{Text: "--ctrlz_port", Description: "ControlZ port (default `9876`)"},
	{Text: "--selector", Alias: "-l", Description: "Label selector (default ``)"},
}, optionDashboard...)

var optionDashboardEnvoy = append([]prompt.Suggest{
	{Text: "--selector", Alias: "-l", Description: "Label selector (default ``)"},
}, optionDashboard...)

var optionInstall = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--filename", Alias: "-f", Description: "Path to file containing IstioOperator custom resource This flag can be specified multiple times to overlay multiple files"},
	{Text: "--force", Description: "Proceed even with validation errors"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--readiness-timeout", Description: "Maximum time to wait for Istio resources in each component to be ready. (default `5m0s`)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
	{Text: "--set", Alias: "-s", Description: "Override an IstioOperator value"},
	{Text: "--skip-confirmation", Alias: "-y", Description: "The skipConfirmation determines whether the user is prompted for confirmation"},
	{Text: "--verify", Description: "Verify the Istio control plane after installation/in-place upgrade"},
}, optionGlobal...)

var optionKubeInject = append([]prompt.Suggest{
	{Text: "--filename", Alias: "-f", Description: "Input Kubernetes resource filename (default ``)"},
	{Text: "--injectConfigFile", Description: "Injection configuration filename. Cannot be used with --injectConfigMapName (default ``)"},
	{Text: "--meshConfigFile", Description: "Mesh configuration filename. Takes precedence over --meshConfigMapName if set (default ``)"},
	{Text: "--meshConfigMapName", Description: "ConfigMap name for Istio mesh configuration, key should be `mesh` (default `istio`)"},
	{Text: "--output", Alias: "-o", Description: "Modified output Kubernetes resource filename (default ``)"},
	{Text: "--revision", Alias: "-r", Description: "Control plane revision (default ``)"},
	{Text: "--valuesFile", Description: "injection values configuration filename. (default ``)"},
	{Text: "--webhookConfig", Description: "MutatingWebhookConfiguration name for Istio (default `istio-sidecar-injector`)"},
}, optionGlobal...)

var optionManifest = append([]prompt.Suggest{
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
}, optionGlobal...)

var optionManifestDiff = append([]prompt.Suggest{
	{Text: "--directory", Alias: "-r", Description: "Compare directory."},
	{Text: "--ignore", Description: "Ignore all listed items during comparison, using the same list format as selectResources. (default ``)"},
	{Text: "--rename", Description: "Rename resources before comparison"},
	{Text: "--select", Description: "Constrain the list of resources to compare to only the ones in this list, ignoring all others"},
	{Text: "--verbose", Description: "Verbose output"},
}, optionManifest...)

var optionManifestGenerate = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--component", Description: "Specify which component to generate manifests for. (default `[]`)"},
	{Text: "--filename", Alias: "-f", Description: "Path to file containing IstioOperator custom resource This flag can be specified multiple times to overlay multiple files"},
	{Text: "--force", Description: "Proceed even with validation errors."},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--output", Alias: "-o", Description: "Manifest output directory path. (default ``)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
	{Text: "--set", Alias: "-s", Description: "Override an IstioOperator value"},
}, optionManifest...)

var optionManifestInstall = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--filename", Alias: "-f", Description: "Path to file containing IstioOperator custom resource This flag can be specified multiple times to overlay multiple files"},
	{Text: "--force", Description: "Proceed even with validation errors."},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--readiness-timeout", Description: "Maximum time to wait for Istio resources in each component to be ready. (default `5m0s`)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
	{Text: "--set", Alias: "-s", Description: "Override an IstioOperator value"},
	{Text: "--skip-confirmation", Alias: "-y", Description: "The skipConfirmation determines whether the user is prompted for confirmation"},
	{Text: "--verify", Description: "Verify the Istio control plane after installation/in-place upgrade"},
}, optionManifest...)

var optionOperator = optionGlobal

var optionOperatorDump = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--hub", Description: "The hub for the operator controller image. (default `unknown`)"},
	{Text: "--imagePullSecrets", Description: "The imagePullSecrets are used to pull the operator image from the private registry"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--operatorNamespace", Description: "The namespace the operator controller is installed into. (default `istio-operator`)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
	{Text: "--tag", Description: "The tag for the operator controller image. (default `unknown`)"},
}, optionOperator...)

var optionOperatorInit = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--filename", Alias: "-f", Description: "Path to file containing IstioOperator custom resource"},
	{Text: "--hub", Description: "The hub for the operator controller image. (default `unknown`)"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--imagePullSecrets", Description: "The imagePullSecrets are used to pull the operator image from the private registry"},
	{Text: "--operatorNamespace", Description: "The namespace the operator controller is installed into. (default `istio-operator`)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
	{Text: "--tag", Description: "The tag for the operator controller image. (default `unknown`)"},
	{Text: "--watchedNamespaces", Description: "The namespaces the operator controller watches"},
}, optionOperator...)

var optionOperatorRemove = append([]prompt.Suggest{
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--force", Description: "Proceed even with validation errors"},
	{Text: "--manifests", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--operatorNamespace", Description: "The namespace the operator controller is installed into. (default `istio-operator`)"},
	{Text: "--revision", Description: "Target control plane revision for the command. (default ``)"},
}, optionOperator...)

var optionOptions = optionGlobal

var optionProfile = append([]prompt.Suggest{
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
}, optionGlobal...)

var optionProfileDiff = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
}, optionProfile...)

var optionProfileDump = append([]prompt.Suggest{
	{Text: "--config-path", Alias: "-p", Description: "The path the root of the configuration subtree to dump"},
	{Text: "--output", Alias: "-o", Description: "Output format: one of json|yaml|flags (default `yaml`)"},
}, optionProfileDiff...)

var optionProfileList = optionProfileDiff

var optionProxyConfig = append([]prompt.Suggest{
	{Text: "--output", Alias: "-o", Description: "Output format: one of json|yaml|short (default `short`)"},
}, optionGlobal...)

var optionProxyConfigAll = append([]prompt.Suggest{
	{Text: "--address", Description: "Filter listeners by address field (default ``)"},
	{Text: "--direction", Description: "Filter clusters by Direction field (default ``)"},
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--fqdn", Description: "Filter clusters by substring of Service FQDN field (default ``)"},
	{Text: "--name", Description: "Filter listeners by route name field (default ``)"},
	{Text: "--port", Description: "Filter clusters and listeners by Port field (default `0`)"},
	{Text: "--subset", Description: "Filter clusters by substring of Subset field (default ``)"},
	{Text: "--type", Description: "Filter listeners by type field (default ``)"},
	{Text: "--verbose", Description: "Output more information"},
}, optionProxyConfig...)

var optionProxyConfigBootstrap = append([]prompt.Suggest{
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
}, optionProxyConfig...)

var optionProxyConfigCluster = append([]prompt.Suggest{
	{Text: "--direction", Description: "Filter clusters by Direction field (default ``)"},
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--fqdn", Description: "Filter clusters by substring of Service FQDN field (default ``)"},
	{Text: "--port", Description: "Filter clusters and listeners by Port field (default `0`)"},
	{Text: "--subset", Description: "Filter clusters by substring of Subset field (default ``)"},
}, optionProxyConfig...)

var optionProxyConfigEndpoint = append([]prompt.Suggest{
	{Text: "--address", Description: "Filter listeners by address field (default ``)"},
	{Text: "--cluster", Description: "Filter endpoints by cluster name field (default ``)"},
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--port", Description: "Filter clusters and listeners by Port field (default `0`)"},
	{Text: "--status", Description: "Filter endpoints by status field (default ``)"},
}, optionProxyConfig...)

var optionProxyConfigListener = append([]prompt.Suggest{
	{Text: "--address", Description: "Filter listeners by address field (default ``)"},
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--port", Description: "Filter clusters and listeners by Port field (default `0`)"},
	{Text: "--type", Description: "Filter listeners by type field (default ``)"},
	{Text: "--verbose", Description: "Output more information"},
}, optionProxyConfig...)

var optionProxyConfigLog = append([]prompt.Suggest{
	{Text: "-level", Description: "Comma-separated minimum per-logger level of messages to output"},
	{Text: "--reset", Alias: "-r", Description: "Reset levels to default value (warning)."},
	{Text: "--selector", Alias: "-l", Description: "Label selector (default ``)"},
}, optionProxyConfig...)

var optionProxyConfigRoute = append([]prompt.Suggest{
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--name", Description: "Filter listeners by route name field (default ``)"},
	{Text: "--verbose", Description: "Output more information"},
}, optionProxyConfig...)

var optionProxyConfigSecret = append([]prompt.Suggest{
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
}, optionProxyConfig...)

var optionProxyStatus = append([]prompt.Suggest{
	{Text: "--file", Alias: "-f", Description: "Envoy config dump file (default ``)"},
	{Text: "--revision", Alias: "-r", Description: "Target control plane revision for the command. (default ``)"},
}, optionGlobal...)

var optionUpgrade = append([]prompt.Suggest{
	{Text: "--charts", Description: "Deprecated, use --manifests instead. (default ``)"},
	{Text: "--dry-run", Description: "Console/log output only, make no changes."},
	{Text: "--filename", Alias: "-f", Description: "Path to file containing IstioOperator custom resource This flag can be specified multiple times to overlay multiple files"},
	{Text: "--force", Description: "Apply the upgrade without eligibility checks"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--readiness-timeout", Description: "Maximum time to wait for Istio resources in each component to be ready. (default `5m0s`)"},
	{Text: "--set", Alias: "-s", Description: "Override an IstioOperator value, e.g. to choose a profile"},
	{Text: "--skip-confirmation", Alias: "-y", Description: "If skip-confirmation is set, skips the prompting confirmation for value changes in this upgrade"},
	{Text: "--verify", Description: "Verify the Istio control plane after installation/in-place upgrade"},
}, optionGlobal...)

var optionValidate = append([]prompt.Suggest{
	{Text: "--filename", Alias: "-f", Description: "Names of files to validate (default `[]`)"},
	{Text: "--referential", Alias: "-x", Description: "Enable structural validation for policy and telemetry"},
}, optionGlobal...)

var optionVerifyInstall = append([]prompt.Suggest{
	{Text: "--filename", Alias: "-f", Description: "Istio YAML installation file. (default `[]`)"},
	{Text: "--manifests", Alias: "-d", Description: "Specify a path to a directory of charts and profiles"},
	{Text: "--revision", Alias: "-r", Description: "Control plane revision (default ``)"},
}, optionGlobal...)

var optionVersion = append([]prompt.Suggest{
	{Text: "--output", Alias: "-o", Description: "One of 'yaml' or 'json'. (default ``)"},
	{Text: "--remote", Description: "Use --remote=false to suppress control plane check"},
	{Text: "--revision", Alias: "-r", Description: "Control plane revision (default ``)"},
	{Text: "--short", Alias: "-s", Description: "Use --short=false to generate full version information"},
}, optionGlobal...)
