package main

import (
	"github.com/NooBeeID/go-logging/messages"
)

func main() {
	messages.SysInfo("Server running on port: %v", "3000")
}
