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

func (c *EchoController) Echo(req *carrot.Request, br *carrot.Broadcast) {
	res, err := carrot.CreateDefaultResponse(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	br.Broadcast(res)
}

func (c *EchoController) PrintParams(req *carrot.Request, res *carrot.Broadcast) {
	fmt.Println(req.Params)
}

type EchoStreamController struct {
	count int
}

func (c *EchoStreamController) EchoStreamer(req *carrot.Request, res *carrot.Broadcast) {
	fmt.Println(c.count)
	c.count += 1
}

func main() {

	// Register endpoints here in the form of endpoint, controller, method
	carrot.Add("echo", EchoController{}, "Echo", true)
	carrot.Add("print_params", EchoController{}, "PrintParams", false)
	carrot.Add("echo_streamer", EchoController{}, "EchoStreamer", true)

	// Run the server and serve traffic
	carrot.Run()
}
