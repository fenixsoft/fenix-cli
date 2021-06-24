package kubernetes

import (
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
)

const (
	Pod                   = "pod"
	Namespace             = "namespace"
	Secret                = "secret"
	ServiceAccount        = "serviceaccount"
	Context               = "context"
	ConfigMap             = "configmap"
	ComponentStatus       = "componentstatus"
	DaemonSet             = "daemonset"
	Deployment            = "deployment"
	Endpoint              = "endpoint"
	Ingress               = "ingress"
	Job                   = "job"
	LimitRange            = "limitrange"
	Node                  = "node"
	PersistentVolumeClaim = "persistentvolumeclaim"
	PersistentVolume      = "persistentvolume"
	PodSecurityPolicy     = "podsecuritypolicy"
	PodTemplate           = "podtemplate"
	ReplicaSet            = "replicaset"
	ReplicationController = "replicationcontroller"
	ResourceQuota         = "resourcequota"
	Service               = "service"
	Resource              = "Resource"
)

func providePodSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetPodSuggestions(Completer.Client, Completer.Namespace)
}

func provideNamespaceSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetNameSpaceSuggestions(Completer)
}

func provideSecretSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetSecretSuggestions(Completer.Client, Completer.Namespace)
}

func provideServiceAccountSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetServiceAccountSuggestions(Completer.Client, Completer.Namespace)
}

func provideContextSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetContextSuggestions()
}

func provideConfigMapSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetConfigMapSuggestions(Completer.Client, Completer.Namespace)
}

func provideComponentStatusSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetComponentStatusCompletions(Completer.Client)
}

func provideDaemonSetSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetDaemonSetSuggestions(Completer.Client, Completer.Namespace)
}

func provideDeploymentSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetDeploymentSuggestions(Completer.Client, Completer.Namespace)
}

func provideEndpointSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetEndpointsSuggestions(Completer.Client, Completer.Namespace)
}

func provideIngressSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetIngressSuggestions(Completer.Client, Completer.Namespace)
}

func provideJobSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetJobSuggestions(Completer.Client, Completer.Namespace)
}

func provideLimitRangeSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetLimitRangeSuggestions(Completer.Client, Completer.Namespace)
}

func provideNodeSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetNodeSuggestions(Completer.Client)
}

func providePersistentVolumeClaimSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetPersistentVolumeClaimSuggestions(Completer.Client, Completer.Namespace)
}

func providePersistentVolumeSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetPersistentVolumeSuggestions(Completer.Client)
}

func providePodSecurityPolicySuggestion(arg ...string) []prompt.Suggest {
	return kube.GetPodSecurityPolicySuggestions(Completer.Client)
}

func providePodTemplateSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetPodTemplateSuggestions(Completer.Client, Completer.Namespace)
}

func provideReplicaSetSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetReplicaSetSuggestions(Completer.Client, Completer.Namespace)
}

func provideReplicationControllerSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetReplicationControllerSuggestions(Completer.Client, Completer.Namespace)
}

func provideResourceQuotaSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetResourceQuotasSuggestions(Completer.Client, Completer.Namespace)
}

func provideServiceSuggestion(arg ...string) []prompt.Suggest {
	return kube.GetServiceSuggestions(Completer.Client, Completer.Namespace)
}

func provideResourceSuggestion(arg ...string) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "pod", Alias: "po"},
		{Text: "namespaces", Alias: "ns"},
		{Text: "secrets"},
		{Text: "serviceaccounts", Alias: "sa"},
		{Text: "configmaps", Alias: "cm"},
		{Text: "componentstatuses", Alias: "cc"},
		{Text: "daemonsets", Alias: "ds"},
		{Text: "deployments", Alias: "deploy"},
		{Text: "endpoints", Alias: "ep"},
		{Text: "ingresses", Alias: "ing"},
		{Text: "jobs", Alias: "job"},
		{Text: "limitranges", Alias: "limit"},
		{Text: "nodes", Alias: "no"},
		{Text: "persistentvolumeclaims", Alias: "pvc"},
		{Text: "persistentvolumes", Alias: "pv"},
		{Text: "podsecuritypolicies", Alias: "psp"},
		{Text: "podtemplates"},
		{Text: "replicasets", Alias: "rs"},
		{Text: "replicationcontrollers", Alias: "rc"},
		{Text: "resourcequotas", Alias: "quota"},
		{Text: "services", Alias: "svc"},
	}
}
