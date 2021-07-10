package main

import "common"

func main() {
	var logger = common.NewLogger("service01")
	logger.InvalidArgValue("client", "nil")
}