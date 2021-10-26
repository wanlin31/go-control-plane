package sotw

import (
	"context"
	"reflect"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

// watches for all xDS resource types
type watches struct {
	responders map[string]*watch

	// indexes is a list of indexes for each dynamic select case which match to a watch
	cases []reflect.SelectCase
}

// newWatches creates and initializes watches.
func newWatches() watches {
	return watches{
		responders: make(map[string]*watch, int(types.UnknownType)),
		cases:      make([]reflect.SelectCase, 0),
	}
}

// Cancel all watches
func (w *watches) Cancel() {
	for _, watch := range w.responders {
		watch.Cancel()
	}
}

// recomputeWatches rebuilds the known list of dynamic channels if needed
func (w *watches) RecomputeWatches(ctx context.Context, req <-chan *discovery.DiscoveryRequest) {
	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ctx.Done()),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(req),
		},
	}

	index := len(cases)
	for _, watch := range w.responders {
		cases = append(cases, w.cases[watch.index])
		watch.index = index
		index++
	}

	w.cases = cases
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	cancel func()
	nonce  string

	// Index is used to track the location of this channel in watches. This allows us
	// to update the channel used at this slot without recomputing the entire list of select
	// statements.
	index int
}

// Cancel calls terminate and cancel
func (w *watch) Cancel() {
	if w.cancel != nil {
		w.cancel()
	}
}
