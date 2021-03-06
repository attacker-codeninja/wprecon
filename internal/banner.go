package internal

import (
	"fmt"
	"time"

	"github.com/blackcrw/wprecon/pkg/printer"
)

// Sbanner :: A banner that will only be executed if the scan is started correctly.
func SBanner(target string) {
	Banner()
	printer.Done("Target:", target)
	printer.Done("Starting:", time.Now().Format(("02/jan/2006 15:04:05")), "\n")
}

// Banner :: A simple banner.
func Banner() {
	fmt.Print("—————————————————————————————————————————————————————————————————————\n")
	fmt.Print(`
___       ______________________________________________   __
__ |     / /__  __ \__  __ \__  ____/_  ____/_  __ \__  | / /
__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / 
__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  
____/|__/  /_/     /_/ |_| /_____/  \____/  \____/ /_/ |_/   

`)
	fmt.Print("Github: ", "https://github.com/blackcrw/wprecon\n")
	fmt.Print("Version: ", "0.0.1a\n")
	fmt.Print("—————————————————————————————————————————————————————————————————————\n")
}
