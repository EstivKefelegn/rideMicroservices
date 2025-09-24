package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	PickUP      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
