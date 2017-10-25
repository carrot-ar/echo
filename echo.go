package main

import (
	"fmt"
	"github.com/carrot-ar/carrot"
)

/*
  Requests are in the form of

  {
			"session_token: "de0u21ejdd12d",
			"endpoint": "test",
      "params": {
          "test1": "result1",
          "test2": "result2"
      }
  }

*/

// Controller implementation
type EchoController struct{}

func (c *EchoController) Echo(req *carrot.Request, res *carrot.Broadcast) {
	res.Send([]byte([]byte(fmt.Sprintf("%v", req.Params["rick"]))))
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
