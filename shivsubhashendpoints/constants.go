package shivsubhashendpoints

import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

const clientId = "522011478849-a1thjunen0u36rs4qoksoih24vdlabtn.apps.googleusercontent.com"

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.APIExplorerClientID}
	// in case we'll want to use TicTacToe API from an Android app
	audiences = []string{clientId}
)
