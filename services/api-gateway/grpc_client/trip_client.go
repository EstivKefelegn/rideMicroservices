package grpcclient

import pb "ride-sharing/shared/proto/trip"


type tripServiceClient struct {
	CLient pb.TripServiceClient
}