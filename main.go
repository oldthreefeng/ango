package main

import (
	"github.com/oldthreefeng/ango/cmd"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
