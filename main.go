package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/file/internal"
)

func main() {
	drone.New(internal.NewPlugin).Boot()
}
