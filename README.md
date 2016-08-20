# Hue Shortcuts
Bind custom keyboard shortcuts to custom Philips Hue actions.

## Installation
```
go get github.com/collinux/gohue
```

## Usage
Use the [GoHue Library](https://github.com/Collinux/GoHue) to custom code any light functions you need.

### Binary Method (Recommended)
The `hueshortcuts` binary is included in this repo which can be built by running `go build hueshortcuts.go`.
Note: These instructions are for Linux only, though it's also possible on Windows and Mac
1. Go into System Settings -> Keyboard -> Shortcuts -> Add 
2. Add a new shortcut with "Name": (anything), "Command": "`hueshortcuts` `command`"
Make sure `hueshortcuts` includes the path to the binary (/home/collin/Hue_Shortcuts/hueshortcuts)

### Other Method
You can also run the go file by using `go run /path_to/hueshortcuts.go`, though this does not work on most systems.
