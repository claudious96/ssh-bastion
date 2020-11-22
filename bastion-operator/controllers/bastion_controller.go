/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	tenantv1alpha1 "example.com/m/api/v1alpha1"
)

// BastionReconciler reconciles a Bastion object
type BastionReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tenant.crownlabs.polito.it,resources=bastions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tenant.crownlabs.polito.it,resources=bastions/status,verbs=get;update;patch

func (r *BastionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("bastion", req.NamespacedName)
	log.Info("reconciling bastion")

	// Get tenants resources
	var list tenantv1alpha1.TenantList
	if err := r.List(ctx, &list); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Collect public keys of tenants
	var keys []string
	for _, tenant := range list.Items {
		keys = append(keys, tenant.Spec.PublicKeys...)
	}

	authorizedKeys := strings.Join(keys[:], "\n")

	log.Info(authorizedKeys)

	cm := &corev1.ConfigMap{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: "default",
		Name:      "auth-keys-bastion",
	}, cm)
	if err != nil {
		err = r.Create(ctx, &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "auth-keys-bastion",
			},
			Data: map[string]string{"authorized_keys": authorizedKeys},
		}, &client.CreateOptions{})
		if err != nil && !errors.IsAlreadyExists(err) {
			log.Error(err, "unable to create configmap auth-keys-bastion")
			return ctrl.Result{}, nil
		}
	} else {

		cm.Data = map[string]string{"authorized_keys": authorizedKeys}
		if err := r.Update(ctx, cm); err != nil && !errors.IsAlreadyExists(err) {
			log.Error(err, "unable to update configmap auth-keys-bastion")
			return ctrl.Result{}, nil
		}

	}
	return ctrl.Result{}, nil
}

func (r *BastionReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&tenantv1alpha1.Bastion{}).
		Watches(
			&source.Kind{Type: &tenantv1alpha1.Tenant{}},
			&handler.EnqueueRequestForObject{},
		).
		Complete(r)
}
