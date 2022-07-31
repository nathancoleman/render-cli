package main

import "github.com/nathancoleman/render-sdk-go"

func main() {
	_, err := render.NewClient(render.DefaultConfig())
	if err != nil {
		panic(err)
	}
}
