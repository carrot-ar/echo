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

type request struct {
	SessionToken string  `json:"session_token"`
	Endpoint     string  `json:"endpoint"`
	Origin       Origin  `json:"origin"`
	Payload      Payload `json:"payload"`
}

type Origin struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Altitude  float64 `json:"altitude"`
}

type Payload struct {
	Offset Offset            `json:"offset"`
	Params map[string]string `json:"params"`
}

type Offset struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Controller implementation
type EchoController struct{}

func (c *EchoController) Echo(req *carrot.Request, res *carrot.Broadcast) {
	responseData := request{
		SessionToken: string(req.SessionToken),
		Endpoint:     "echo",
		Origin: Origin{
			Latitude:  req.Origin.Latitude,
			Longitude: req.Origin.Longitude,
			Altitude:  req.Origin.Altitude,
		},
		Payload: Payload{
			Offset: Offset{
				X: req.Offset.X,
				Y: req.Offset.Y,
				Z: req.Offset.Z,
			},
			Params: req.Params,
		},
	}

	jsonData, err := json.Marshal(&responseData)
	if err != nil {
		fmt.Println("COULD NOT UNMARSHAL")
		fmt.Println(err)
		return
	}
	res.Send(jsonData)
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
