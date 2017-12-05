package main

import (
	"fmt"
	"github.com/carrot-ar/carrot"
)

/*
  Requests and responses are in the form of:

	{
		"session_token": "KjIQhKUPNrvHkUHv1VySBg==",
		"endpoint": "test_endpoint",
		"payload": {
			"offset": {
				"x": 3.2,
				"y": 1.3,
				"z": 4.0
			},
			"params": {
				"foo": "bar"
			}
		}
	}

*/

// Controller implementation
type EchoController struct{}

func (c *EchoController) EchoSimple(req *carrot.Request, br *carrot.Broadcast) {
	message, err := carrot.CreateDefaultResponse(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	br.Broadcast(message)
}

func (c *EchoController) EchoExtendable(req *carrot.Request, br *carrot.Broadcast) {
	token := string(req.SessionToken)
	payload, err := carrot.NewPayload(token, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := carrot.NewResponse(token, "Echo", payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.AddParam("someKey", "someValue")
	message, err := res.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	br.Broadcast(message)
}

func main() {

	// Register endpoints here in the form of endpoint, controller, method
	carrot.Add("echo_simple", EchoController{}, "EchoSimple", true)
	carrot.Add("echo_extendable", EchoController{}, "EchoExtendable", true)

	// Run the server and serve traffic
	carrot.Run()
}
