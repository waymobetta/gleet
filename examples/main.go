package main

import (
	"github.com/waymobetta/gleet"
)

func main() {
	data := "world domination is imminent"
	delay := 111

	leetWord := gleet.Profile{
		Data:  data,
		Delay: delay,
	}

	leetWord.Render()
}
