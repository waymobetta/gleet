package main

import (
	"flag"

	"github.com/waymobetta/gleet"
)

func main() {
	delayFlag := flag.Int("delay", 125, "time delay")

	flag.Parse()

	leetWord := gleet.Profile{
		Data:  flag.Args()[0],
		Delay: *delayFlag,
	}

	// fmt.Println("Running with delay: ", *delayFlag)

	leetWord.Render()
}
