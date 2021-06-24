package kubernetes

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
)

var Completer *kube.Completer

//wrapper for kubernetes completer changed (while namespace change)
func Complete(d prompt.Document) []prompt.Suggest {
	return Completer.Complete(d)
}

func RegisterEnv() (*environments.Runtime, error) {
	if c, err := kube.NewCompleter(); err != nil {
		return nil, err
	} else {
		Completer = c

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

		Completer.KubernetesRuntime = &environments.Runtime{
			Prefix: "kubectl",
			//Completer:      Complete,
			Executor:       environments.GetDefaultExecutor("kubectl", nil, krew.GetBinPath()...),
			Commands:       ExtraCommands,
			MainSuggestion: kube.Commands,
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > kubectl ", c.Namespace), true
			},
		}
		return Completer.KubernetesRuntime, nil
	}
}
