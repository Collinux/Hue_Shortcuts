/*
*
* hueshortcuts.go
* Uses a Philips Hue go library to enable command shortcuts for light functions.
*
* Examples: go run hueshortcuts.go all_on
*           go run hueshortcuts.go all_off
*           go run hueshortcuts.go writing_code_scene
 */

package main

import (
	"github.com/collinux/gohue"
	"log"
	"os"
)

func main() {
	// Find your hue bridge or enter the IP
	bridges, err := hue.FindBridges()
	if err != nil {
		log.Fatal("Could not find any bridges. ", err)
	}
	bridge := bridges[0]

	// Enter your own username token from bridge.Login
	// or see the godoc from https://github.com/Collinux/GoHue for details.
	bridge.Username = "427de8bd6d49f149c8398e4fc08f"

	lights, _ := bridge.GetAllLights()
	switch os.Args[1] { // Index 0 arg would be the file name
	case "all_on":
		log.Print("Calling all_on")
		for _, light := range lights {
			light.On()
			light.SetColor(hue.WHITE)
		}
		break
	case "all_off":
		log.Print("Calling all_off")
		for _, light := range lights {
			light.Off()
		}
		break
	case "writing_code_scene":
		log.Print("Calling writing_code_scene")
		for _, light := range lights {
			if light.Name == "Desk Light" {
				light.On()
				light.SetColor(hue.WHITE)
			} else if light.Name == "Bedroom Lamp" {
				light.On()
				light.SetColor(hue.BLUE)
				light.SetBrightness(70)
			} else {
				light.Off()
			}
		}
		break
	}
}
