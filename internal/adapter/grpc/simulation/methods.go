/*
Package simulation provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/model/model.proto).

All commands for services to model management are provided by this adapter.
*/
package simulation

import (
	api "github.com/Goboolean/command-server/api/grpc/simulation"
)

// SimulationModel
//
// This method requests a model based on the model information received in the request.
// The return type is stream, which provides a model state.
// When the model is finished, return the access token to the join server at the end of the stream and close the stream.
func (a *Adapter) SimulationModel(in *api.SimulationModelRequest, stream api.SimulationService_SimulationModelServer) error {

	ch := make(chan *api.SimulationStatus)

	// ==============================================================================================================
	go func() {
		ch <- &api.SimulationStatus{
			Status:      1,
			AccessToken: "access_token",
		}
		close(ch)
	}()

	for s := range ch {
		if err := stream.Send(s); err != nil {
			return err
		}
	}
	// ==============================================================================================================

	return nil
}
