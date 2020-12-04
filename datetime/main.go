package main

import (
	"fmt"
	"time"
	"utils/src/info"
)

func main() {
	info.CheckPrintInfo()
	fmt.Print(time.Now().Unix())
}
