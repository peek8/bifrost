package fsm

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	bifrostv1alpha1 "github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/alloy"
	"github.com/peek8/bifrost/internal/components"
)

// A struct that holds the items needed for the actions to do their work.
// Things like client libraries and loggers, go here.
type Context struct {
	Logger   logr.Logger
	Client   client.Client
	Ctx      context.Context
	Recorder record.EventRecorder
}

// A struct that holds the "extended state" of the state machine, including data
// being fetched and read. This should only be modified by actions, while guards
// should only read the extended state to assess their value.
type ExtendedState struct {
	LastError    error
	Result       ctrl.Result
	ResourceName types.NamespacedName
	Instance     *bifrostv1alpha1.LogSpace

	// Waves and componentes
	// DesiredComponents is a list of objects that we want to create at cluster.
	DesiredComponents components.Components

	Waves map[string]components.Components

	CurrentWave string
	// DesiredComponent is the first component from  DesiredComponents
	DesiredComponent components.Component

	// ActualComponent is the current loaded kubernetes. It is nil if its not created yet.
	ActualComponent components.Component

	// ObservedComponents holds all of the objects that were desired and processed
	ObservedComponents components.Components

	Flags map[string]bool

	Dirty bool

	// builders
	AlloyBuilder components.Builder[alloy.Alloy, alloy.Data]
}
