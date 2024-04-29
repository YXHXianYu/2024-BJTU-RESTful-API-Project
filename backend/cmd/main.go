package main

import (
	"restfulapi/internal/utils"
)

func main() {
	r := utils.Setup()
    r.Run()
}