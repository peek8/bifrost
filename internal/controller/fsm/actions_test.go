package fsm_test

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/peek8/bifrost/internal/controller/fsm"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	bifrostv1alpha1 "github.com/peek8/bifrost/api/v1alpha1"
)

// silentLogger creates a logger that discards all output
func silentLogger() logr.Logger {
	return logr.Discard()
}

// testContext returns a fully configured test context
func testContext() *fsm.Context {
	// Create a new scheme and register all types we might need in tests
	testScheme := scheme.Scheme
	_ = bifrostv1alpha1.AddToScheme(testScheme)

	// Create a fake client with the comprehensive scheme and status subresource support
	fakeClient := fake.NewClientBuilder().
		WithScheme(testScheme).
		WithStatusSubresource(&bifrostv1alpha1.LogSpace{}).
		Build()

	return &fsm.Context{
		Logger: silentLogger(),
		Client: fakeClient,
		Ctx:    context.TODO(),
	}
}

// +vectorsigma:action:ClearCondition
func TestBifrostOperator_ClearConditionAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.ClearConditionAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.ClearConditionAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:ClearError
func TestBifrostOperator_ClearErrorAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.ClearErrorAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.ClearErrorAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:CollectCommonStatusInformation
func TestBifrostOperator_CollectCommonStatusInformationAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.CollectCommonStatusInformationAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.CollectCommonStatusInformationAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:CollectStatusInformation
func TestBifrostOperator_CollectStatusInformationAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.CollectStatusInformationAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.CollectStatusInformationAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:CreateComponent
func TestBifrostOperator_CreateComponentAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.CreateComponentAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.CreateComponentAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateAlloyConfig
func TestBifrostOperator_GenerateAlloyConfigAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateAlloyConfigAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateAlloyConfigAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateAlloyRbac
func TestBifrostOperator_GenerateAlloyRbacAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateAlloyRbacAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateAlloyRbacAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateAlloyStorage
func TestBifrostOperator_GenerateAlloyStorageAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateAlloyStorageAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateAlloyStorageAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateAlloyWorkloads
func TestBifrostOperator_GenerateAlloyWorkloadsAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateAlloyWorkloadsAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateAlloyWorkloadsAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateGrafanaConfig
func TestBifrostOperator_GenerateGrafanaConfigAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateGrafanaConfigAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateGrafanaConfigAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateGrafanaStorage
func TestBifrostOperator_GenerateGrafanaStorageAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateGrafanaStorageAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateGrafanaStorageAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateGrafanaWorkloads
func TestBifrostOperator_GenerateGrafanaWorkloadsAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateGrafanaWorkloadsAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateGrafanaWorkloadsAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateLokiConfig
func TestBifrostOperator_GenerateLokiConfigAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateLokiConfigAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateLokiConfigAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateLokiStorage
func TestBifrostOperator_GenerateLokiStorageAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateLokiStorageAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateLokiStorageAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:GenerateLokiWorkloads
func TestBifrostOperator_GenerateLokiWorkloadsAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.GenerateLokiWorkloadsAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.GenerateLokiWorkloadsAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:InitializeContext
func TestBifrostOperator_InitializeContextAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.InitializeContextAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.InitializeContextAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:InitializeWave
func TestBifrostOperator_InitializeWaveAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.InitializeWaveAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.InitializeWaveAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:LoadComponent
func TestBifrostOperator_LoadComponentAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.LoadComponentAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.LoadComponentAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:LoadSubject
func TestBifrostOperator_LoadSubjectAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.LoadSubjectAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.LoadSubjectAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:PickNextComponent
func TestBifrostOperator_PickNextComponentAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.PickNextComponentAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.PickNextComponentAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:RaiseCondition
func TestBifrostOperator_RaiseConditionAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.RaiseConditionAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.RaiseConditionAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:ReconcileIn
func TestBifrostOperator_ReconcileInAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.ReconcileInAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.ReconcileInAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:SendEvent
func TestBifrostOperator_SendEventAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.SendEventAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.SendEventAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:SetFlag
func TestBifrostOperator_SetFlagAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.SetFlagAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.SetFlagAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:SetPhase
func TestBifrostOperator_SetPhaseAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.SetPhaseAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.SetPhaseAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:SetWave
func TestBifrostOperator_SetWaveAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.SetWaveAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.SetWaveAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:UpdateComponent
func TestBifrostOperator_UpdateComponentAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.UpdateComponentAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.UpdateComponentAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:UpdateStatus
func TestBifrostOperator_UpdateStatusAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.UpdateStatusAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.UpdateStatusAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// +vectorsigma:action:DoneWithComponent
func TestBifrostOperator_DoneWithComponentAction(t *testing.T) {
	type fields struct {
		context       *fsm.Context
		currentState  fsm.StateName
		stateConfigs  map[fsm.StateName]fsm.StateConfig
		ExtendedState *fsm.ExtendedState
	}

	type args struct {
		params []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fsm := &fsm.BifrostOperator{
				Context:       tt.fields.context,
				CurrentState:  tt.fields.currentState,
				StateConfigs:  tt.fields.stateConfigs,
				ExtendedState: tt.fields.ExtendedState,
			}
			if err := fsm.DoneWithComponentAction(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("BifrostOperator.DoneWithComponentAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
