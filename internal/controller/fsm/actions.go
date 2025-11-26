package fsm

import (
	"reflect"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	k8stypes "k8s.io/apimachinery/pkg/types"

	bifrostv1alpha1 "github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/components"
	ctrl "sigs.k8s.io/controller-runtime"
)

// +vectorsigma:action:ClearCondition
func (fsm *BifrostOperator) ClearConditionAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:ClearError
func (fsm *BifrostOperator) ClearErrorAction(_ ...string) error {
	fsm.ExtendedState.LastError = nil

	return nil
}

// +vectorsigma:action:CollectCommonStatusInformation
func (fsm *BifrostOperator) CollectCommonStatusInformationAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:CollectStatusInformation
func (fsm *BifrostOperator) CollectStatusInformationAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:CreateComponent
func (fsm *BifrostOperator) CreateComponentAction(_ ...string) error {
	err := ctrl.SetControllerReference(fsm.ExtendedState.Instance, fsm.ExtendedState.DesiredComponent, fsm.Context.Client.Scheme())
	if err != nil {
		return err
	}

	// TODO: Set annotations and labels

	err = fsm.Context.Client.Create(fsm.Context.Ctx, fsm.ExtendedState.DesiredComponent.GetClientObject())
	if err != nil {
		return err
	}
	fsm.ExtendedState.ActualComponent = fsm.ExtendedState.DesiredComponent

	return nil
}

// +vectorsigma:action:GenerateAlloyConfig
func (fsm *BifrostOperator) GenerateAlloyConfigAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateAlloyRbac
func (fsm *BifrostOperator) GenerateAlloyRbacAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateAlloyStorage
func (fsm *BifrostOperator) GenerateAlloyStorageAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateAlloyWorkloads
func (fsm *BifrostOperator) GenerateAlloyWorkloadsAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateGrafanaConfig
func (fsm *BifrostOperator) GenerateGrafanaConfigAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateGrafanaStorage
func (fsm *BifrostOperator) GenerateGrafanaStorageAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateGrafanaWorkloads
func (fsm *BifrostOperator) GenerateGrafanaWorkloadsAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateLokiConfig
func (fsm *BifrostOperator) GenerateLokiConfigAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateLokiStorage
func (fsm *BifrostOperator) GenerateLokiStorageAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:GenerateLokiWorkloads
func (fsm *BifrostOperator) GenerateLokiWorkloadsAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:InitializeContextAction
func (fsm *BifrostOperator) InitializeContextAction(_ ...string) error {
	fsm.ExtendedState.Waves = make(map[string]components.Components)
	fsm.ExtendedState.Flags = make(map[string]bool)

	return nil
}

// +vectorsigma:action:InitializeWave
func (fsm *BifrostOperator) InitializeWaveAction(_ ...string) error {
	fsm.ExtendedState.DesiredComponents = fsm.ExtendedState.Waves[fsm.ExtendedState.CurrentWave]

	return nil
}

// +vectorsigma:action:LoadComponent
func (fsm *BifrostOperator) LoadComponentAction(_ ...string) error {
	nn := k8stypes.NamespacedName{
		Name:      fsm.ExtendedState.DesiredComponent.GetName(),
		Namespace: fsm.ExtendedState.ResourceName.Namespace,
	}

	copy := reflect.New(reflect.TypeOf(fsm.ExtendedState.DesiredComponent).Elem()).Interface()
	fsm.ExtendedState.ActualComponent = copy.(components.Component)

	err := fsm.Context.Client.Get(fsm.Context.Ctx, nn, fsm.ExtendedState.ActualComponent.GetClientObject())
	if k8serrors.IsNotFound(err) {
		fsm.ExtendedState.ActualComponent = nil

		return nil
	}

	if err != nil {
		fsm.ExtendedState.ActualComponent = nil

		return err
	}

	return nil
}

// +vectorsigma:action:LoadSubject
func (fsm *BifrostOperator) LoadSubjectAction(_ ...string) error {
	subject := new(bifrostv1alpha1.LogSpace)

	err := fsm.Context.Client.Get(fsm.Context.Ctx, fsm.ExtendedState.ResourceName, subject)
	if err != nil {
		return err
	}

	fsm.ExtendedState.Instance = subject

	return nil
}

// +vectorsigma:action:PickNextComponent
func (fsm *BifrostOperator) PickNextComponentAction(_ ...string) error {
	// remove first element from list and place it in state
	fsm.ExtendedState.DesiredComponent = fsm.ExtendedState.DesiredComponents[0]

	fsm.ExtendedState.DesiredComponents = fsm.ExtendedState.DesiredComponents[1:]

	return nil
}

// +vectorsigma:action:RaiseCondition
func (fsm *BifrostOperator) RaiseConditionAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}

// +vectorsigma:action:ReconcileIn
func (fsm *BifrostOperator) ReconcileInAction(params ...string) error {

	duration, err := time.ParseDuration(params[0])
	if err != nil {
		panic(err)
	}

	fsm.ExtendedState.Result.RequeueAfter = duration

	return nil
}

// +vectorsigma:action:SendEvent
func (fsm *BifrostOperator) SendEventAction(_ ...string) error {

	return nil
}

// +vectorsigma:action:SetFlag
func (fsm *BifrostOperator) SetFlagAction(params ...string) error {
	fsm.ExtendedState.Flags[params[0]] = true

	return nil
}

// +vectorsigma:action:SetPhase
func (fsm *BifrostOperator) SetPhaseAction(params ...string) error {
	if fsm.ExtendedState.Instance.Status.Phase != params[0] {
		fsm.ExtendedState.Instance.Status.Phase = params[0]

		fsm.ExtendedState.Dirty = true
	}

	return nil
}

// +vectorsigma:action:SetWave
func (fsm *BifrostOperator) SetWaveAction(params ...string) error {
	fsm.ExtendedState.CurrentWave = params[0]

	return nil
}

// +vectorsigma:action:UpdateComponent
func (fsm *BifrostOperator) UpdateComponentAction(_ ...string) error {
	fsm.ExtendedState.DesiredComponent.DeepCopySpecInto(fsm.ExtendedState.ActualComponent)

	err := fsm.Context.Client.Update(fsm.Context.Ctx, fsm.ExtendedState.ActualComponent.GetClientObject())
	if err != nil {
		return err
	}

	return nil
}

// +vectorsigma:action:UpdateStatus
func (fsm *BifrostOperator) UpdateStatusAction(_ ...string) error {
	err := fsm.Context.Client.Status().Update(fsm.Context.Ctx, fsm.ExtendedState.Instance)
	if err != nil {
		return err
	}

	return nil
}

// +vectorsigma:action:DoneWithComponent
func (fsm *BifrostOperator) DoneWithComponentAction(_ ...string) error {
	// TODO: Implement me!
	return nil
}
