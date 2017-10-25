package main

import (
	"encoding/json"
	"fmt"
	"github.com/carrot-ar/carrot"
)

/*
  Requests are in the form of

	{
		"session_token": "KjIQhKUPNrvHkUHv1VySBg==",
		"endpoint": "test_endpoint",
		"origin": {
			"longitude": 45.501689,
			"latitude": -73.567256
		},
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

func (c *EchoController) Echo(req *carrot.Request, res *carrot.Broadcast) {
	jsonData, err := json.Marshal(req)
	fmt.Println(err)
	res.Send([]byte([]byte(jsonData)))
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
