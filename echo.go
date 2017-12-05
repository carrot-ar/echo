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
	res, err := carrot.NewResponse(token, "Print", payload)
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

func (c *EchoController) Print(req *carrot.Request, br *carrot.Broadcast) {
	fmt.Printf("The params are:\t%v\n", req.Params)
}

func main() {

	// Register endpoints here in the form of endpoint, controller, method
	carrot.Add("echoSimple", EchoController{}, "EchoSimple", true)
	carrot.Add("echoExtendable", EchoController{}, "EchoExtendable", true)
	carrot.Add("print", EchoController{}, "Print", true)

	// Run the server and serve traffic
	carrot.Run()
}
