package kubernetes

import (
	"github.com/c-bata/go-prompt"
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

	Resource = "Resource"
	Verb     = "verb"
	Cascade  = "cascade"
)

func providePodSuggestion(arg ...string) []prompt.Suggest {
	return GetPodSuggestions(Client.Set, Client.Namespace)
}

func provideNamespaceSuggestion(arg ...string) []prompt.Suggest {
	return GetNameSpaceSuggestions(Client)
}

func provideSecretSuggestion(arg ...string) []prompt.Suggest {
	return GetSecretSuggestions(Client.Set, Client.Namespace)
}

func provideServiceAccountSuggestion(arg ...string) []prompt.Suggest {
	return GetServiceAccountSuggestions(Client.Set, Client.Namespace)
}

func provideContextSuggestion(arg ...string) []prompt.Suggest {
	return GetContextSuggestions()
}

func provideConfigMapSuggestion(arg ...string) []prompt.Suggest {
	return GetConfigMapSuggestions(Client.Set, Client.Namespace)
}

func provideComponentStatusSuggestion(arg ...string) []prompt.Suggest {
	return GetComponentStatusCompletions(Client.Set)
}

func provideDaemonSetSuggestion(arg ...string) []prompt.Suggest {
	return GetDaemonSetSuggestions(Client.Set, Client.Namespace)
}

func provideDeploymentSuggestion(arg ...string) []prompt.Suggest {
	return GetDeploymentSuggestions(Client.Set, Client.Namespace)
}

func provideEndpointSuggestion(arg ...string) []prompt.Suggest {
	return GetEndpointsSuggestions(Client.Set, Client.Namespace)
}

func provideIngressSuggestion(arg ...string) []prompt.Suggest {
	return GetIngressSuggestions(Client.Set, Client.Namespace)
}

func provideJobSuggestion(arg ...string) []prompt.Suggest {
	return GetJobSuggestions(Client.Set, Client.Namespace)
}

func provideLimitRangeSuggestion(arg ...string) []prompt.Suggest {
	return GetLimitRangeSuggestions(Client.Set, Client.Namespace)
}

func provideNodeSuggestion(arg ...string) []prompt.Suggest {
	return GetNodeSuggestions(Client.Set)
}

func providePersistentVolumeClaimSuggestion(arg ...string) []prompt.Suggest {
	return GetPersistentVolumeClaimSuggestions(Client.Set, Client.Namespace)
}

func providePersistentVolumeSuggestion(arg ...string) []prompt.Suggest {
	return GetPersistentVolumeSuggestions(Client.Set)
}

func providePodSecurityPolicySuggestion(arg ...string) []prompt.Suggest {
	return GetPodSecurityPolicySuggestions(Client.Set)
}

func providePodTemplateSuggestion(arg ...string) []prompt.Suggest {
	return GetPodTemplateSuggestions(Client.Set, Client.Namespace)
}

func provideReplicaSetSuggestion(arg ...string) []prompt.Suggest {
	return GetReplicaSetSuggestions(Client.Set, Client.Namespace)
}

func provideReplicationControllerSuggestion(arg ...string) []prompt.Suggest {
	return GetReplicationControllerSuggestions(Client.Set, Client.Namespace)
}

func provideResourceQuotaSuggestion(arg ...string) []prompt.Suggest {
	return GetResourceQuotasSuggestions(Client.Set, Client.Namespace)
}

func provideServiceSuggestion(arg ...string) []prompt.Suggest {
	return GetServiceSuggestions(Client.Set, Client.Namespace)
}

func provideResourceSuggestion(arg ...string) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "pod", Alias: "po", Description: "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts."},
		{Text: "namespaces", Alias: "ns", Description: "Namespace provides a scope for Names. Use of multiple namespaces is optional."},
		{Text: "secrets", Description: "Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes."},
		{Text: "serviceaccounts", Alias: "sa", Description: "ServiceAccount binds together: a name"},
		{Text: "configmaps", Alias: "cm", Description: "ConfigMap holds configuration data for pods to consume."},
		{Text: "componentstatuses", Alias: "cc", Description: "ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+"},
		{Text: "daemonsets", Alias: "ds", Description: "DaemonSet represents the configuration of a daemon set."},
		{Text: "deployments", Alias: "deploy", Description: "Deployment enables declarative updates for Pods and ReplicaSets."},
		{Text: "endpoints", Alias: "ep", Description: "Endpoints is a collection of endpoints that implement the actual service."},
		{Text: "ingresses", Alias: "ing", Description: "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend."},
		{Text: "jobs", Alias: "job", Description: "Job represents the configuration of a single job."},
		{Text: "limitranges", Alias: "limit", Description: "LimitRange sets resource usage limits for each kind of resource in a Namespace."},
		{Text: "nodes", Alias: "no", Description: "Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd)."},
		{Text: "persistentvolumeclaims", Alias: "pvc", Description: "PersistentVolumeClaim is a user's request for and claim to a persistent volume"},
		{Text: "persistentvolumes", Alias: "pv", Description: "PersistentVolume (PV) is a storage resource provisioned by an administrator."},
		{Text: "podsecuritypolicies", Alias: "psp", Description: "PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container. Deprecated in 1.21."},
		{Text: "podtemplates", Description: "PodTemplate describes a template for creating copies of a predefined pod."},
		{Text: "replicasets", Alias: "rs", Description: "ReplicaSet ensures that a specified number of pod replicas are running at any given time."},
		{Text: "replicationcontrollers", Alias: "rc", Description: "ReplicationController represents the configuration of a replication controller."},
		{Text: "resourcequotas", Alias: "quota", Description: "ResourceQuota sets aggregate quota restrictions enforced per namespace"},
		{Text: "services", Alias: "svc", Description: "Service is a named abstraction of software service"},
	}
}
