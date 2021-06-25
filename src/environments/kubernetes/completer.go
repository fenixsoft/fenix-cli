package kubernetes

import (
	"context"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	Namespace     string
	NamespaceList *corev1.NamespaceList
	Set           *kubernetes.Clientset
}

var Client *KubeClient

func New() (*suggestions.GenericCompleter, error) {
	// Determine if docker is running
	if c, err := NewClient(); err != nil {
		return nil, err
	} else {
		Client = c
	}
	return suggestions.NewGenericCompleter(arguments, options, func(c *suggestions.GenericCompleter) {
		c.SuggestionProviders.Add(suggestions.Argument, suggestions.BuildStaticCompletionProvider(c.Arguments, suggestions.LengthFilter))
		c.SuggestionProviders.Add(suggestions.Option, suggestions.BuildStaticCompletionProvider(c.Options, suggestions.BestEffortFilter))

		suggestions.RegisterSharedProvider(Pod, providePodSuggestion)
		suggestions.RegisterSharedProvider(Namespace, provideNamespaceSuggestion)
		suggestions.RegisterSharedProvider(Secret, provideSecretSuggestion)
		suggestions.RegisterSharedProvider(ServiceAccount, provideServiceAccountSuggestion)
		suggestions.RegisterSharedProvider(Context, provideContextSuggestion)
		suggestions.RegisterSharedProvider(ConfigMap, provideConfigMapSuggestion)
		suggestions.RegisterSharedProvider(ComponentStatus, provideComponentStatusSuggestion)
		suggestions.RegisterSharedProvider(DaemonSet, provideDaemonSetSuggestion)
		suggestions.RegisterSharedProvider(Deployment, provideDeploymentSuggestion)
		suggestions.RegisterSharedProvider(Endpoint, provideEndpointSuggestion)
		suggestions.RegisterSharedProvider(Ingress, provideIngressSuggestion)
		suggestions.RegisterSharedProvider(Job, provideJobSuggestion)
		suggestions.RegisterSharedProvider(LimitRange, provideLimitRangeSuggestion)
		suggestions.RegisterSharedProvider(Node, provideNodeSuggestion)
		suggestions.RegisterSharedProvider(PersistentVolume, providePersistentVolumeSuggestion)
		suggestions.RegisterSharedProvider(PersistentVolumeClaim, providePersistentVolumeClaimSuggestion)
		suggestions.RegisterSharedProvider(PodSecurityPolicy, providePodSecurityPolicySuggestion)
		suggestions.RegisterSharedProvider(PodTemplate, providePodTemplateSuggestion)
		suggestions.RegisterSharedProvider(ReplicaSet, provideReplicaSetSuggestion)
		suggestions.RegisterSharedProvider(ReplicationController, provideReplicationControllerSuggestion)
		suggestions.RegisterSharedProvider(ResourceQuota, provideResourceQuotaSuggestion)
		suggestions.RegisterSharedProvider(Service, provideServiceSuggestion)
		suggestions.RegisterSharedProvider(Resource, provideResourceSuggestion)

		suggestions.RegisterSharedProvider(Verb, suggestions.BuildFixedSelectionProvider("get", "list", "watch", "create", "update", "patch", "delete"))
		suggestions.RegisterSharedProvider(Cascade, suggestions.BuildFixedSelectionProvider("background", "orphan", "foreground"))
	}), nil
}

func NewClient() (*KubeClient, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.WarnIfAllMissing = false
	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

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

	return &KubeClient{
		Namespace:     namespace,
		NamespaceList: namespaces,
		Set:           client,
	}, nil
}
