package main

import (
	"fmt"
	"time"
	"utils/internal/info"
)

func main() {
	info.CheckPrintInfo()
	fmt.Print(time.Now().Unix())
}
