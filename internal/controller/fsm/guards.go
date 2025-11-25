package fsm

// +vectorsigma:guard:ComponentDiffers
func (fsm *BifrostOperator) ComponentDiffersGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:ComponentDoesNotExist
func (fsm *BifrostOperator) ComponentDoesNotExistGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:Dirty
func (fsm *BifrostOperator) DirtyGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:HasFlag
func (fsm *BifrostOperator) HasFlagGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:IsError
func (fsm *BifrostOperator) IsErrorGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:MoreComponents
func (fsm *BifrostOperator) MoreComponentsGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:ReconcileNotPlanned
func (fsm *BifrostOperator) ReconcileNotPlannedGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:MissingFlag
func (fsm *BifrostOperator) MissingFlagGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:ObservedComponentsNotReady
func (fsm *BifrostOperator) ObservedComponentsNotReadyGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:WasWave
func (fsm *BifrostOperator) WasWaveGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}

// +vectorsigma:guard:WaveNotDone
func (fsm *BifrostOperator) WaveNotDoneGuard(_ ...string) bool {
	// TODO: Implement me!
	return false
}
