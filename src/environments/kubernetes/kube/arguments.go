package kube

import (
	prompt "github.com/c-bata/go-prompt"
	"strings"
)

var Commands = []prompt.Suggest{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "replace", Description: "Replace a resource by filename or stdin."},
	{Text: "patch", Description: "Update field(s) of a resource using strategic merge patch."},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "logs", Description: "Print the logs for a container in a pod."},
	{Text: "rolling-update", Description: "Perform a rolling update of the given ReplicationController."},
	{Text: "scale", Description: "Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job."},
	{Text: "cordon", Description: "Mark node as unschedulable"},
	{Text: "drain", Description: "Drain node in preparation for maintenance"},
	{Text: "uncordon", Description: "Mark node as schedulable"},
	{Text: "attach", Description: "Attach to a running container."},
	{Text: "exec", Description: "Execute a command in a container."},
	{Text: "port-forward", Description: "Forward one or more local ports to a pod."},
	{Text: "proxy", Description: "Run a proxy to the Kubernetes API server"},
	{Text: "run", Description: "Run a particular image on the cluster."},
	{Text: "expose", Description: "Take a replication controller, service, or pod and expose it as a new Kubernetes Service"},
	{Text: "autoscale", Description: "Auto-scale a Deployment, ReplicaSet, or ReplicationController"},
	{Text: "rollout", Description: "rollout manages a deployment"},
	{Text: "label", Description: "Update the labels on a resource"},
	{Text: "annotate", Description: "Update the annotations on a resource"},
	{Text: "config", Description: "config modifies kubeconfig files"},
	{Text: "cluster-info", Description: "Display cluster info"},
	{Text: "api-versions", Description: "Print the supported API versions on the server, in the form of 'group/version'."},
	{Text: "api-resources", Description: "Print the supported API resources on the server'."},
	{Text: "version", Description: "Print the client and server version information."},
	{Text: "explain", Description: "Documentation of resources."},
	{Text: "convert", Description: "Convert config files between different API versions"},
	{Text: "top", Description: "Display Resource (CPU/Memory/Storage) usage"},
	{Text: "options", Description: "Print the list of flags inherited by all commands"},
	{Text: "plugin", Description: "Provides utilities for interacting with plugins."},
	{Text: "taint", Description: "Update the taints on one or more nodes."},
	{Text: "debug", Description: "Debug cluster resources using interactive debugging containers."},
	{Text: "kustomize", Description: "Build a set of KRM resources using a 'kustomization.yaml' file."},
}

var xBatchTypes = []prompt.Suggest{
	{Text: "ComponentStatuses", Alias: "cs", Description: "ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+"},
	{Text: "ConfigMaps", Alias: "cm", Description: "ConfigMap holds configuration data for pods to consume."},
	{Text: "DaemonSets", Alias: "ds", Description: "DaemonSet represents the configuration of a daemon set."},
	{Text: "Deployments", Alias: "deploy", Description: "Deployment enables declarative updates for Pods and ReplicaSets."},
	{Text: "Endpoints", Alias: "ep", Description: "Endpoints is a collection of endpoints that implement the actual service."},
	{Text: "Ingresses", Alias: "ing", Description: "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend."},
	{Text: "Jobs", Description: "Job represents the configuration of a single job."},
	{Text: "Limitranges", Alias: "limits", Description: "LimitRange sets resource usage limits for each kind of resource in a Namespace."},
	{Text: "Namespaces", Alias: "ns", Description: "Namespace provides a scope for Names. Use of multiple namespaces is optional."},
	{Text: "Nodes", Alias: "no", Description: "Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd)."},
	{Text: "PersistentVolumeClaims", Alias: "pvc", Description: "PersistentVolumeClaim is a user's request for and claim to a persistent volume"},
	{Text: "PersistentVolumes", Alias: "pv", Description: "PersistentVolume (PV) is a storage resource provisioned by an administrator."},
	{Text: "Pod", Alias: "po", Description: "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts."},
	{Text: "PodSecurityPolicies", Alias: "psp", Description: "PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container. Deprecated in 1.21."},
	{Text: "PodTemplates", Description: "PodTemplate describes a template for creating copies of a predefined pod."},
	{Text: "Replicasets", Alias: "rs", Description: "ReplicaSet ensures that a specified number of pod replicas are running at any given time."},
	{Text: "ReplicationControllers", Alias: "rc", Description: "ReplicationController represents the configuration of a replication controller."},
	{Text: "ResourceQuotas", Alias: "quota", Description: "ResourceQuota sets aggregate quota restrictions enforced per namespace"},
	{Text: "Secrets", Description: "Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes."},
	{Text: "ServiceAccounts", Alias: "sa", Description: "ServiceAccount binds together: a name, understood by users, and perhaps by peripheral systems, for an identity a principal that can be authenticated and authorized a set of secrets"},
	{Text: "Services", Alias: "svc", Description: "Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy."},
}

var resourceTypes = []prompt.Suggest{
	{Text: "ComponentStatuses", Alias: "cs", Description: "ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+"},
	{Text: "ConfigMaps", Alias: "cm", Description: "ConfigMap holds configuration data for pods to consume."},
	{Text: "DaemonSets", Alias: "ds", Description: "DaemonSet represents the configuration of a daemon set."},
	{Text: "Deployments", Alias: "deploy", Description: "Deployment enables declarative updates for Pods and ReplicaSets."},
	{Text: "Endpoints", Alias: "ep", Description: "Endpoints is a collection of endpoints that implement the actual service."},
	{Text: "Events", Alias: "ev", Description: "Event is a report of an event somewhere in the cluster."},
	{Text: "HorizontalPodAutoscaler", Alias: "hpa", Description: "HorizontalPodAutoscaler configuration of a horizontal pod autoscaler."},
	{Text: "Ingresses", Alias: "ing", Description: "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend."},
	{Text: "Jobs", Description: "Job represents the configuration of a single job."},
	{Text: "CronJobs", Alias: "cj", Description: "CronJob represents the configuration of a single cron job."},
	{Text: "limitranges", Alias: "limits", Description: "LimitRange sets resource usage limits for each kind of resource in a Namespace."},
	{Text: "Namespaces", Alias: "ns", Description: "Namespace provides a scope for Names. Use of multiple namespaces is optional."},
	{Text: "NetworkPolicies", Alias: "netpol", Description: "NetworkPolicy describes what network traffic is allowed for a set of Pods"},
	{Text: "Nodes", Alias: "no", Description: "Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd)."},
	{Text: "PersistentVolumeClaims", Alias: "pvc", Description: "PersistentVolumeClaim is a user's request for and claim to a persistent volume"},
	{Text: "PersistentVolumes", Alias: "pv", Description: "PersistentVolume (PV) is a storage resource provisioned by an administrator."},
	{Text: "Pod", Alias: "po", Description: "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts."},
	{Text: "PodSecurityPolicies", Alias: "psp", Description: "PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container. Deprecated in 1.21."},
	{Text: "PodTemplates", Description: "PodTemplate describes a template for creating copies of a predefined pod."},
	{Text: "Replicasets", Alias: "rs", Description: "ReplicaSet ensures that a specified number of pod replicas are running at any given time."},
	{Text: "ReplicationControllers", Alias: "rc", Description: "ReplicationController represents the configuration of a replication controller."},
	{Text: "ResourceQuotas", Alias: "quota", Description: "ResourceQuota sets aggregate quota restrictions enforced per namespace"},
	{Text: "Secrets", Description: "Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes."},
	{Text: "ServiceAccounts", Alias: "sa", Description: "ServiceAccount binds together: a name, understood by users, and perhaps by peripheral systems, for an identity a principal that can be authenticated and authorized a set of secrets"},
	{Text: "Services", Alias: "svc", Description: "Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy."},
	{Text: "StatefulSets", Alias: "sts", Description: "StatefulSet represents a set of pods with consistent identities."},
	{Text: "StorageClasses", Alias: "sc", Description: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned."},
	{Text: "Bindings", Description: "Binding ties one object to another"},
	{Text: "MutatingWebhookConfigurations", Description: "MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object."},
	{Text: "ValidatingWebhookConfigurations", Description: "ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it."},
	{Text: "CustomResourceDefinitions", Alias: "crd", Description: "CustomResourceDefinition represents a resource that should be exposed on the API server."},
	{Text: "APIServices", Description: "APIService represents a server for a particular GroupVersion. Name must be \"version.group\"."},
	{Text: "ControllerRevisions", Description: "ControllerRevision implements an immutable snapshot of state data."},
	{Text: "TokenReviews", Description: "TokenReview attempts to authenticate a token to a known user."},
	{Text: "LocalSubjectAccessReviews", Description: "LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace."},
	{Text: "SelfSubjectAccessReviews", Description: "SelfSubjectAccessReview checks whether or the current user can perform an action."},
	{Text: "selfsubjectrulesreviews", Description: "SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace."},
	{Text: "SubjectAccessReviews", Description: "SubjectAccessReview checks whether or not a user or group can perform an action."},
	{Text: "CertificateSigningRequests", Alias: "csr", Description: "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued."},
	{Text: "Leases", Description: "Lease defines a lease concept."},
	{Text: "EndpointSlices", Description: "EndpointSlice represents a subset of the endpoints that implement a service."},
	{Text: "FlowSchemas", Description: "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a \"flow distinguisher\"."},
	{Text: "PriorityLevelConfigurations", Description: "PriorityLevelConfiguration represents the configuration of a priority level."},
	{Text: "IngressClasses", Description: "IngressClass represents the class of the Ingress, referenced by the Ingress Spec."},
	{Text: "RuntimeClasses", Description: "RuntimeClass defines a class of container runtime supported in the cluster."},
	{Text: "PodDisruptionBudgets", Alias: "pdb", Description: "PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods"},
	{Text: "ClusterRoleBindings", Description: "ClusterRoleBinding references a ClusterRole, but not contain it. It can reference a ClusterRole in the global namespace, and adds who information  via Subject."},
	{Text: "ClusterRoles", Description: "ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding."},
	{Text: "RoleBindings", Description: "RoleBinding references a role, but does not contain it. It can reference a Role in the same namespace or a ClusterRole in the global namespace."},
	{Text: "Roles", Description: "Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding."},
	{Text: "PriorityClasses", Alias: "pc", Description: "PriorityClass defines mapping from a priority class name to the priority integer value."},
	{Text: "CSIDrivers", Description: "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster."},
	{Text: "CSINodes", Description: "CSINode holds information about all CSI drivers installed on a node. CSI  drivers do not need to create the CSINode object directly."},
	{Text: "CSIStorageCapacities", Description: "CSIStorageCapacity stores the result of one CSI GetCapacity call. For a given StorageClass, this describes the available capacity in a particular topology segment."},
	{Text: "VolumeAttachments", Description: "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node."},
}

func (c *Completer) argumentsCompleter(namespace string, args []string) []prompt.Suggest {
	if len(args) == 0 {
		return c.KubernetesRuntime.MainSuggestion
	} else if len(args) == 1 {
		return prompt.FilterHasPrefix(c.KubernetesRuntime.MainSuggestion, args[0], true)
	}

	majorCmd := strings.ToLower(args[0])
	subCmd := ""
	if len(args) > 1 {
		subCmd = strings.ToLower(args[1])
	}
	argument := ""
	if len(args) > 2 {
		argument = strings.ToLower(args[2])
	}
	switch majorCmd {
	case "get":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, subCmd, true)
		}

		if len(args) == 3 {
			switch subCmd {
			case "componentstatuses", "cs":
				return prompt.FilterContains(GetComponentStatusCompletions(c.Client), argument, true)
			case "configmaps", "cm":
				return prompt.FilterContains(GetConfigMapSuggestions(c.Client, namespace), argument, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(GetDaemonSetSuggestions(c.Client, namespace), argument, true)
			case "deploy", "deployments":
				return prompt.FilterContains(GetDeploymentSuggestions(c.Client, namespace), argument, true)
			case "endpoints", "ep":
				return prompt.FilterContains(GetEndpointsSuggestions(c.Client, namespace), argument, true)
			case "ingresses", "ing":
				return prompt.FilterContains(GetIngressSuggestions(c.Client, namespace), argument, true)
			case "limitranges", "limits":
				return prompt.FilterContains(GetLimitRangeSuggestions(c.Client, namespace), argument, true)
			case "namespaces", "ns":
				return prompt.FilterContains(GetNameSpaceSuggestions(c.NamespaceList), argument, true)
			case "no", "nodes":
				return prompt.FilterContains(GetNodeSuggestions(c.Client), argument, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(GetPersistentVolumeClaimSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(GetPersistentVolumeSuggestions(c.Client), argument, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(GetPodSecurityPolicySuggestions(c.Client), argument, true)
			case "podtemplates":
				return prompt.FilterContains(GetPodTemplateSuggestions(c.Client, namespace), argument, true)
			case "replicasets", "rs":
				return prompt.FilterContains(GetReplicaSetSuggestions(c.Client, namespace), argument, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), argument, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(GetResourceQuotasSuggestions(c.Client, namespace), argument, true)
			case "secrets":
				return prompt.FilterContains(GetSecretSuggestions(c.Client, namespace), argument, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(GetServiceAccountSuggestions(c.Client, namespace), argument, true)
			case "svc", "services":
				return prompt.FilterContains(GetServiceSuggestions(c.Client, namespace), argument, true)
			case "job", "jobs":
				return prompt.FilterContains(GetJobSuggestions(c.Client, namespace), argument, true)
			}
		}
	case "describe":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, subCmd, true)
		}
		if len(args) == 3 {
			switch subCmd {
			case "componentstatuses", "cs":
				return prompt.FilterContains(GetComponentStatusCompletions(c.Client), argument, true)
			case "configmaps", "cm":
				return prompt.FilterContains(GetConfigMapSuggestions(c.Client, namespace), argument, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(GetDaemonSetSuggestions(c.Client, namespace), argument, true)
			case "deploy", "deployments":
				return prompt.FilterContains(GetDeploymentSuggestions(c.Client, namespace), argument, true)
			case "endpoints", "ep":
				return prompt.FilterContains(GetEndpointsSuggestions(c.Client, namespace), argument, true)
			case "ingresses", "ing":
				return prompt.FilterContains(GetIngressSuggestions(c.Client, namespace), argument, true)
			case "limitranges", "limits":
				return prompt.FilterContains(GetLimitRangeSuggestions(c.Client, namespace), argument, true)
			case "namespaces", "ns":
				return prompt.FilterContains(GetNameSpaceSuggestions(c.NamespaceList), argument, true)
			case "no", "nodes":
				return prompt.FilterContains(GetNodeSuggestions(c.Client), argument, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(GetPersistentVolumeClaimSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(GetPersistentVolumeSuggestions(c.Client), argument, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(GetPodSecurityPolicySuggestions(c.Client), argument, true)
			case "podtemplates":
				return prompt.FilterContains(GetPodTemplateSuggestions(c.Client, namespace), argument, true)
			case "replicasets", "rs":
				return prompt.FilterContains(GetReplicaSetSuggestions(c.Client, namespace), argument, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), argument, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(GetResourceQuotasSuggestions(c.Client, namespace), argument, true)
			case "secrets":
				return prompt.FilterContains(GetSecretSuggestions(c.Client, namespace), argument, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(GetServiceAccountSuggestions(c.Client, namespace), argument, true)
			case "svc", "services":
				return prompt.FilterContains(GetServiceSuggestions(c.Client, namespace), argument, true)
			case "job", "jobs":
				return prompt.FilterContains(GetJobSuggestions(c.Client, namespace), argument, true)
			}
		}
	case "create":
		subcommands := []prompt.Suggest{
			{Text: "ConfigMap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "Deployment", Description: "Create a deployment with the specified name."},
			{Text: "Namespace", Description: "Create a namespace with the specified name"},
			{Text: "Quota", Description: "Create a quota with the specified name."},
			{Text: "Secret", Description: "Create a secret using specified subcommand"},
			{Text: "Service", Description: "Create a service using specified subcommand."},
			{Text: "ServiceAccount", Description: "Create a service account with the specified name"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "delete":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, subCmd, true)
		}
		if len(args) == 3 {
			switch subCmd {
			case "componentstatuses", "cs":
				return prompt.FilterContains(GetComponentStatusCompletions(c.Client), argument, true)
			case "configmaps", "cm":
				return prompt.FilterContains(GetConfigMapSuggestions(c.Client, namespace), argument, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(GetDaemonSetSuggestions(c.Client, namespace), argument, true)
			case "deploy", "deployments":
				return prompt.FilterContains(GetDeploymentSuggestions(c.Client, namespace), argument, true)
			case "endpoints", "ep":
				return prompt.FilterContains(GetEndpointsSuggestions(c.Client, namespace), argument, true)
			case "ingresses", "ing":
				return prompt.FilterContains(GetIngressSuggestions(c.Client, namespace), argument, true)
			case "limitranges", "limits":
				return prompt.FilterContains(GetLimitRangeSuggestions(c.Client, namespace), argument, true)
			case "namespaces", "ns":
				return prompt.FilterContains(GetNameSpaceSuggestions(c.NamespaceList), argument, true)
			case "no", "nodes":
				return prompt.FilterContains(GetNodeSuggestions(c.Client), argument, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(GetPersistentVolumeClaimSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(GetPersistentVolumeSuggestions(c.Client), argument, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(GetPodSecurityPolicySuggestions(c.Client), argument, true)
			case "podtemplates":
				return prompt.FilterContains(GetPodTemplateSuggestions(c.Client, namespace), argument, true)
			case "replicasets", "rs":
				return prompt.FilterContains(GetReplicaSetSuggestions(c.Client, namespace), argument, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), argument, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(GetResourceQuotasSuggestions(c.Client, namespace), argument, true)
			case "secrets":
				return prompt.FilterContains(GetSecretSuggestions(c.Client, namespace), argument, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(GetServiceAccountSuggestions(c.Client, namespace), argument, true)
			case "svc", "services":
				return prompt.FilterContains(GetServiceSuggestions(c.Client, namespace), argument, true)
			case "job", "jobs":
				return prompt.FilterContains(GetJobSuggestions(c.Client, namespace), argument, true)
			}
		}
	case "edit":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, args[1], true)
		}
		if len(args) == 3 {
			switch subCmd {
			case "componentstatuses", "cs":
				return prompt.FilterContains(GetComponentStatusCompletions(c.Client), argument, true)
			case "configmaps", "cm":
				return prompt.FilterContains(GetConfigMapSuggestions(c.Client, namespace), argument, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(GetDaemonSetSuggestions(c.Client, namespace), argument, true)
			case "deploy", "deployments":
				return prompt.FilterContains(GetDeploymentSuggestions(c.Client, namespace), argument, true)
			case "endpoints", "ep":
				return prompt.FilterContains(GetEndpointsSuggestions(c.Client, namespace), argument, true)
			case "ingresses", "ing":
				return prompt.FilterContains(GetIngressSuggestions(c.Client, namespace), argument, true)
			case "limitranges", "limits":
				return prompt.FilterContains(GetLimitRangeSuggestions(c.Client, namespace), argument, true)
			case "namespaces", "ns":
				return prompt.FilterContains(GetNameSpaceSuggestions(c.NamespaceList), argument, true)
			case "no", "nodes":
				return prompt.FilterContains(GetNodeSuggestions(c.Client), argument, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(GetPersistentVolumeClaimSuggestions(c.Client, namespace), argument, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(GetPersistentVolumeSuggestions(c.Client), argument, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(GetPodSecurityPolicySuggestions(c.Client), argument, true)
			case "podtemplates":
				return prompt.FilterContains(GetPodTemplateSuggestions(c.Client, namespace), argument, true)
			case "replicasets", "rs":
				return prompt.FilterContains(GetReplicaSetSuggestions(c.Client, namespace), argument, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), argument, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(GetResourceQuotasSuggestions(c.Client, namespace), argument, true)
			case "secrets":
				return prompt.FilterContains(GetSecretSuggestions(c.Client, namespace), argument, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(GetServiceAccountSuggestions(c.Client, namespace), argument, true)
			case "svc", "services":
				return prompt.FilterContains(GetServiceSuggestions(c.Client, namespace), argument, true)
			case "job", "jobs":
				return prompt.FilterContains(GetJobSuggestions(c.Client, namespace), argument, true)
			}
		}
	case "logs", "x-sniff", "x-lens":
		if len(args) == 2 {
			return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), args[1], true)
		}
	case "rolling-update", "rollingupdate":
		if len(args) == 2 {
			return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), args[1], true)
		} else if len(args) == 3 {
			return prompt.FilterContains(GetReplicationControllerSuggestions(c.Client, namespace), args[2], true)
		}
	case "scale", "resize":
		if len(args) == 2 {
			// Deployment, ReplicaSet, Replication Controller, or Job.
			r := GetDeploymentSuggestions(c.Client, namespace)
			r = append(r, GetReplicaSetSuggestions(c.Client, namespace)...)
			r = append(r, GetReplicationControllerSuggestions(c.Client, namespace)...)
			return prompt.FilterContains(r, args[1], true)
		}
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	case "uncordon":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(GetNodeSuggestions(c.Client), subCmd, true)
		}
	case "attach":
		if len(args) == 2 {
			return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), subCmd, true)
		}
	case "exec":
		if len(args) == 2 {
			return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), subCmd, true)
		}
	case "port-forward":
		if len(args) == 2 {
			return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), subCmd, true)
		}
		if len(args) == 3 {
			return prompt.FilterHasPrefix(getPortsFromPodName(namespace, subCmd), argument, true)
		}
	case "rollout":
		subCommands := []prompt.Suggest{
			{Text: "history", Description: "view rollout history"},
			{Text: "pause", Description: "Mark the provided resource as paused"},
			{Text: "resume", Description: "Resume a paused resource"},
			{Text: "undo", Description: "undoes a previous rollout"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, subCmd, true)
		}
	case "annotate":
	case "config":
		subCommands := []prompt.Suggest{
			{Text: "current-context", Description: "Displays the current-context"},
			{Text: "delete-cluster", Description: "Delete the specified cluster from the kubeconfig"},
			{Text: "delete-context", Description: "Delete the specified context from the kubeconfig"},
			{Text: "get-clusters", Description: "Display clusters defined in the kubeconfig"},
			{Text: "get-contexts", Description: "Describe one or many contexts"},
			{Text: "set", Description: "Sets an individual value in a kubeconfig file"},
			{Text: "set-cluster", Description: "Sets a cluster entry in kubeconfig"},
			{Text: "set-context", Description: "Sets a context entry in kubeconfig"},
			{Text: "set-credentials", Description: "Sets a user entry in kubeconfig"},
			{Text: "unset", Description: "Unsets an individual value in a kubeconfig file"},
			{Text: "use-context", Description: "Sets the current-context in a kubeconfig file"},
			{Text: "view", Description: "Display merged kubeconfig settings or a specified kubeconfig file"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, subCmd, true)
		}
		if len(args) == 3 {
			switch subCmd {
			case "use-context":
				return prompt.FilterContains(GetContextSuggestions(), argument, true)
			}
		}
	case "cluster-info":
		subCommands := []prompt.Suggest{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, subCmd, true)
		}
	case "explain":
		return prompt.FilterHasPrefix(resourceTypes, subCmd, true)
	case "top":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "nodes", Alias: "no", Description: "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts."},
				{Text: "pod", Alias: "po", Description: "Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd)."},
			}
			return prompt.FilterHasPrefix(subcommands, subCmd, true)
		}
		if len(args) == 3 {
			switch subCmd {
			case "no", "node", "nodes":
				return prompt.FilterContains(GetNodeSuggestions(c.Client), argument, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(GetPodSuggestions(c.Client, namespace), argument, true)
			}
		}
	case "x-batch":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(xBatchTypes, subCmd, true)
		}
	case "x-open":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(GetServiceSuggestions(c.Client, namespace), subCmd, true)
		}
	case "x-status":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, subCmd, true)
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
