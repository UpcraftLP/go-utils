package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"utils/internal/info"
	"utils/pkg/sysutils"
)

func main() {
	info.CheckPrintInfo()
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("at least 1 argument required")
	}

	startTime := time.Now()
	process, err := sysutils.Start(args...)
	if err != nil {
		log.Fatalln(err)
	}
	state, err := process.Wait()
	stopTime := time.Now()
	if err != nil {
		log.Fatalln(err)
	}
	duration := stopTime.Sub(startTime)

	fmt.Printf("Process finished after %v with exit code %v\n", duration, formatCode(state.ExitCode()))
}

func formatCode(exitCode int) string {
	result := strconv.Itoa(exitCode)
	if len(result) > 1 {
		return fmt.Sprintf("0x%v", strconv.FormatInt(int64(exitCode), 16))
	}
	return result
}
