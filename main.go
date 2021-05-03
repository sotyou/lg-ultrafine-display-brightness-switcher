package main

import (
	_ "embed"
	"fmt"

	"github.com/felicianotech/go-lguf/lguf"
	"github.com/wailsapp/wails"
)

var conn *lguf.Connection

func basic(value uint16) uint16 {
	var err error
	if value == 0 {
		value, _ = conn.Brightness()
		return value / 540
	}
	err = conn.SetBrightness(uint16(value * 540))
	if err != nil {
		fmt.Println(err)
	}
	return value
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	var err error
	conn, err = lguf.NewConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 700,
		Title:  "Ultrafine-Brightness-Slider",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
