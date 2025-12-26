package fsm

import (
	"strings"
)

// +vectorsigma:guard:ComponentDiffers
func (fsm *BifrostOperator) ComponentDiffersGuard(_ ...string) bool {
	return fsm.ExtendedState.DesiredComponent.DiffersSemanticallyFrom(fsm.ExtendedState.ActualComponent)
}

// +vectorsigma:guard:ComponentDoesNotExist
func (fsm *BifrostOperator) ComponentDoesNotExistGuard(_ ...string) bool {
	return fsm.ExtendedState.ActualComponent == nil
}

// +vectorsigma:guard:Dirty
func (fsm *BifrostOperator) DirtyGuard(_ ...string) bool {
	return fsm.ExtendedState.Dirty
}

// +vectorsigma:guard:HasFlag
func (fsm *BifrostOperator) HasFlagGuard(params ...string) bool {
	return fsm.ExtendedState.Flags[params[0]]
}

// +vectorsigma:guard:IsError
func (fsm *BifrostOperator) IsErrorGuard(params ...string) bool {
	if fsm.ExtendedState.Error == nil {
		return false
	}

	if len(params) == 0 {
		return true
	}

	if strings.Contains(fsm.ExtendedState.Error.Error(), params[0]) {
		return true
	}

	return false
}

// +vectorsigma:guard:MoreComponents
func (fsm *BifrostOperator) MoreComponentsGuard(_ ...string) bool {
	return len(fsm.ExtendedState.DesiredComponents) > 0
}

// +vectorsigma:guard:ReconcileNotPlanned
func (fsm *BifrostOperator) ReconcileNotPlannedGuard(_ ...string) bool {
	return fsm.ExtendedState.Result.RequeueAfter == 0
}

// +vectorsigma:guard:MissingFlag
func (fsm *BifrostOperator) MissingFlagGuard(params ...string) bool {
	return !fsm.ExtendedState.Flags[params[0]]
}

// +vectorsigma:guard:ObservedComponentsNotReady
func (fsm *BifrostOperator) ObservedComponentsNotReadyGuard(_ ...string) bool {
	for _, object := range fsm.ExtendedState.ObservedComponents {
		if _, _, ready := object.IsReady(); !ready {
			return true
		}
	}

	return false
}

// +vectorsigma:guard:WasWave
func (fsm *BifrostOperator) WasWaveGuard(params ...string) bool {
	return fsm.ExtendedState.CurrentWave == params[0]
}

// +vectorsigma:guard:WaveNotDone
func (fsm *BifrostOperator) WaveNotDoneGuard(_ ...string) bool {
	return fsm.ExtendedState.CurrentWave != "Done"
}
