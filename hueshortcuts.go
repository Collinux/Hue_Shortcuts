package main

import (
    "fmt"
    "os"
    "github.com/collinux/gohue"
)

func main() {
    bridge, _ := hue.NewBridge("192.168.1.128")
    bridge.Username = "427de8bd6d49f149c8398e4fc08f"

    switch os.Args[1] { // Index 0 arg would be the file name
        case "all_on":
            fmt.Println("calling all_on")
            for index := 0; index <= 10; index++ {
                light, err := bridge.GetLightByIndex(index)
                if err != nil {
                    continue
                }
                light.On()
                light.SetColor(hue.WHITE)
            }
            break
        case "all_off":
            fmt.Println("calling all off")
            for index := 0; index <= 10; index++ {
                light, err := bridge.GetLightByIndex(index)
                if err != nil {
                    continue
                }
                light.Off()
            }
            break
        case "writing_code_scene":
            for index := 0; index <= 10; index++ {
                light, err := bridge.GetLightByIndex(index)
                if err != nil {
                    continue
                }
                if light.Name == "Desk Light" {
                    light.On()
                    light.SetColor(hue.WHITE)
                } else if light.Name == "Bedroom Lamp" {
                    light.On()
                    light.SetColor(hue.BLUE)
                    light.Dim(20)
                } else {
                    light.Off()
                }
            }
    }
}
