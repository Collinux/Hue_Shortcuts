/*
* hueshortcuts.go
* Uses a Philips Hue go library to enable command
* line shortcuts for light functions.
* Copyright (C) 2016 Collin Guarino (Collinux) collin.guarino@gmail.com
* License: GPL version 2 or higher http://www.gnu.org/licenses/gpl.html
*/
// Example: go run hueshortcuts.go all_on
// Example: go run hueshortcuts.go all_off
// Example: go run hueshortcuts.go writing_code_scene


package main

import (
    "fmt"
    "os"
    "github.com/collinux/gohue"
)

func main() {
    bridge, _ := hue.NewBridge("192.168.1.128") // Enter your own bridge IP
    bridge.Username = "427de8bd6d49f149c8398e4fc08f" // Enter your own user

    lights, _ := bridge.GetAllLights()
    switch os.Args[1] { // Index 0 arg would be the file name
        case "all_on":
            fmt.Println("calling all_on")
            for _, light := range lights {
                light.On()
                light.SetColor(hue.WHITE)
            }
            break
        case "all_off":
            fmt.Println("calling all off")
            for _, light := range lights {
                light.Off()
            }
            break
        case "writing_code_scene":
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
