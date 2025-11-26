/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"
	bifrostv1alpha1 "github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/controller/fsm"
)

// LogSpaceReconciler reconciles a LogSpace object
type LogSpaceReconciler struct {
	client.Client
	Log       logr.Logger
	Scheme    *runtime.Scheme
	Recorder  record.EventRecorder
	Namespace string
}

//+kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=get;list;create;patch;watch
//+kubebuilder:rbac:groups=bifrost.peek8.io,resources=logspaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bifrost.peek8.io,resources=logspaces/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bifrost.peek8.io,resources=logspaces/finalizers,verbs=update
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=create;update;delete;list;watch;get
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=secrets;configmaps;persistentvolumeclaims;services;serviceaccounts,verbs=create;update;delete;list;watch;get
//+kubebuilder:rbac:groups="gateway.networking.k8s.io",resources=httproutes,verbs=create;update;delete;list;watch;get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LogSpace object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *LogSpaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Ignore requests for other namespaces, if specified
	if r.Namespace != "" && req.Namespace != r.Namespace {
		return ctrl.Result{}, nil
	}

	//initialize and run the state machine
	// Initialize the state machine
	stateMachine := fsm.New()

	// Configure the state machine context
	stateMachine.Context = &fsm.Context{
		Logger:   r.Log,
		Client:   r.Client,
		Ctx:      ctx,
		Recorder: r.Recorder,
	}
	// Set the resource name in the extended state
	stateMachine.ExtendedState.ResourceName = req.NamespacedName

	// Run the state machine
	res, err := stateMachine.Run()

	if err != nil {
		return res, fmt.Errorf("%w", err)
	}

	return res, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LogSpaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bifrostv1alpha1.LogSpace{}).
		Named("logspace").
		Complete(r)
}
