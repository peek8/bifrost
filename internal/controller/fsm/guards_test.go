package fsm_test

import (
	"github.com/peek8/bifrost/internal/controller/fsm"
	"testing"
)

// +vectorsigma:guard:ComponentDiffers
func TestBifrostOperator_ComponentDiffersGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.ComponentDiffersGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.ComponentDiffersGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:ComponentDoesNotExist
func TestBifrostOperator_ComponentDoesNotExistGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.ComponentDoesNotExistGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.ComponentDoesNotExistGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:Dirty
func TestBifrostOperator_DirtyGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.DirtyGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.DirtyGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:HasFlag
func TestBifrostOperator_HasFlagGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.HasFlagGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.HasFlagGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:IsError
func TestBifrostOperator_IsErrorGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.IsErrorGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.IsErrorGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:MoreComponents
func TestBifrostOperator_MoreComponentsGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.MoreComponentsGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.MoreComponentsGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:ReconcileNotPlanned
func TestBifrostOperator_ReconcileNotPlannedGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.ReconcileNotPlannedGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.ReconcileNotPlannedGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:MissingFlag
func TestBifrostOperator_MissingFlagGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.MissingFlagGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.MissingFlagGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:ObservedComponentsNotReady
func TestBifrostOperator_ObservedComponentsNotReadyGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.ObservedComponentsNotReadyGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.ObservedComponentsNotReadyGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:WasWave
func TestBifrostOperator_WasWaveGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.WasWaveGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.WasWaveGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// +vectorsigma:guard:WaveNotDone
func TestBifrostOperator_WaveNotDoneGuard(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   bool
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
			if got := fsm.WaveNotDoneGuard(tt.args.params...); got != tt.want {
				t.Errorf("BifrostOperator.WaveNotDoneGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}
