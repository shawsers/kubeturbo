package util

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	// Kubernetes workload controller types
	KindCronJob               = "CronJob"
	KindDaemonSet             = "DaemonSet"
	KindDeployment            = "Deployment"
	KindJob                   = "Job"
	KindReplicaSet            = "ReplicaSet"
	KindDeploymentConfig      = "DeploymentConfig"
	KindReplicationController = "ReplicationController"
	KindStatefulSet           = "StatefulSet"

	K8sExtensionsGroupName = "extensions"
	K8sAppsGroupName       = "apps"
	OpenShiftAppsGroupName = "apps.openshift.io"

	ReplicationControllerResName = "replicationcontrollers"
	ReplicaSetResName            = "replicasets"
	DeploymentResName            = "deployments"
	DeploymentConfigResName      = "deploymentconfigs"
)

// The API group version under which deployments and replicasets are exposed by the k8s cluster as of today
var K8sAPIDeploymentReplicasetDefaultGV = schema.GroupVersion{Group: K8sAppsGroupName, Version: "v1"}

// The API group version under which deployments are exposed by the k8s cluster
var K8sAPIDeploymentGV = schema.GroupVersion{Group: K8sAppsGroupName, Version: "v1"}

// The API group version under which replicasets are exposed by the k8s cluster
var K8sAPIReplicasetGV = schema.GroupVersion{Group: K8sAppsGroupName, Version: "v1"}

// The API group under which replicationcontrollers are exposed by the k8s server
// We do not discover the latest GV for this as we know that it has matured under core/v1
var K8sAPIReplicationControllerGV = schema.GroupVersion{Group: "", Version: "v1"}

// The API group under which openshifts deploymentconfig resource is exposed by the server
var OpenShiftAPIDeploymentConfigGV = schema.GroupVersion{Group: OpenShiftAppsGroupName, Version: "v1"}
