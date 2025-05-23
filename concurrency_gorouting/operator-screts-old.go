package controller

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SecretReconciler watches for Secret changes
type SecretReconciler struct {
	client.Client
}

// Reconcile is triggered when a Secret is created, updated, or deleted
func (r *SecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Add ReconcileID to the logger for structured logging
	reconcileID := req.Name + "-" + req.Namespace // or any custom logic to generate ReconcileID
	logger = logger.WithValues("reconcileID", reconcileID)

	// Fetch the Secret
	var secret corev1.Secret
	if err := r.Get(ctx, req.NamespacedName, &secret); err != nil {
		logger.Error(err, "❌ Unable to fetch Secret")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Log Secret detected
	logger.Info("🔑 Secret detected", "Secret", secret.Name, "namespace", secret.Namespace)

	// Find Deployments that mount this Secret
	var deployments appsv1.DeploymentList
	if err := r.List(ctx, &deployments, client.InNamespace(secret.Namespace)); err != nil {
		logger.Error(err, "❌ Failed to list Deployments")
		return ctrl.Result{}, err
	}

	// Loop through each Deployment
	for _, deployment := range deployments.Items {
		usesSecret := false
		for _, volume := range deployment.Spec.Template.Spec.Volumes {
			if volume.Secret != nil && volume.Secret.SecretName == secret.Name {
				usesSecret = true
				break
			}
		}

		if usesSecret {
			// Update annotation to trigger a rolling restart
			if deployment.Spec.Template.Annotations == nil {
				deployment.Spec.Template.Annotations = make(map[string]string)
			}
			deployment.Spec.Template.Annotations["secret-update-timestamp"] = time.Now().Format(time.RFC3339)

			// Apply update to Deployment
			if err := r.Update(ctx, &deployment); err != nil {
				logger.Error(err, fmt.Sprintf("❌ Failed to update Deployment %s", deployment.Name))
				return ctrl.Result{}, err
			}

			// Log deployment restart
			logger.Info("🔄 Deployment restarted due to Secret change", "Deployment", deployment.Name, "Secret", secret.Name)

			// Find ReplicaSets for the Deployment
			var replicaSets appsv1.ReplicaSetList
			// Adjust label matching to use the `app` label from the Deployment's spec template labels
			if err := r.List(ctx, &replicaSets, client.InNamespace(secret.Namespace), client.MatchingLabels{"app": deployment.Spec.Template.Labels["app"]}); err != nil {
				logger.Error(err, fmt.Sprintf("❌ Failed to list ReplicaSets for Deployment %s", deployment.Name))
				return ctrl.Result{}, err
			}

			// Log number of ReplicaSets found
			logger.Info(fmt.Sprintf("🔄 Found %d ReplicaSets for Deployment %s", len(replicaSets.Items), deployment.Name))

			// Loop through each ReplicaSet
			for _, replicaSet := range replicaSets.Items {
				logger.Info("🔄 Found ReplicaSet", "ReplicaSet", replicaSet.Name, "Deployment", deployment.Name, "namespace", secret.Namespace)

				// Find Pods managed by the ReplicaSet
				var pods corev1.PodList
				if err := r.List(ctx, &pods, client.InNamespace(secret.Namespace), client.MatchingLabels{"app": deployment.Spec.Template.Labels["app"]}); err != nil {
					logger.Error(err, fmt.Sprintf("❌ Failed to list Pods for ReplicaSet %s", replicaSet.Name))
					return ctrl.Result{}, err
				}

				// Log number of Pods found
				logger.Info(fmt.Sprintf("🔄 Found %d Pods for ReplicaSet %s", len(pods.Items), replicaSet.Name))

				// Loop through each Pod
				for _, pod := range pods.Items {
					logger.Info("🔄 Found Pod", "Pod", pod.Name, "ReplicaSet", replicaSet.Name, "Deployment", deployment.Name, "namespace", secret.Namespace)
				}
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager registers the controller with the manager
func (r *SecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}). // Watch Secret events
		Complete(r)
}

//it is good
