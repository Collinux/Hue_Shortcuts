# Hue Shortcuts
Bind custom keyboard shortcuts to custom Philips Hue actions.

## Installation
Clone or download this repo then run the command `go run hueshortcuts.go all_on` or `go run hueshortcuts.go all_off`

## Usage
Use the [GoHue Library](https://github.com/Collinux/GoHue) to custom code any light functions you need.

### Binary Method (Recommended)
The `hueshortcuts` binary is included in this repo which can be built by running `go build hueshortcuts.go`.
* Note: These instructions are for Linux only, though it's also possible on other operating system that support Go

1. Go into System Settings -> Keyboard -> Shortcuts -> Add 
2. Add a new shortcut with "Name": (anything), "Command": "`hueshortcuts` `command`"
  * Make sure `hueshortcuts` includes the path to the binary (/home/collin/Hue_Shortcuts/hueshortcuts)
  * `command` will be a command option from hueshortcuts.go such as "all_on", "all_off", etc.

### Other Method
A shortcut can be added to run the go file (instead of the binary) by using the command `go run /path_to/hueshortcuts.go` with an command option from hueshortcuts.go such as "all_on", "all_off"
