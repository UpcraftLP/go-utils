package info

import (
	"fmt"
	"os"
	"strings"
)

var Name = "NULL"
var Version = "DEVELOPMENT"
var BuildTime = "INVALID"

type AppInfo interface {
	Name() string
	Version() string
	BuildTime() string
	Print()
}

func GetAppInfo() AppInfo {
	return getAppInfoInternal()
}

func printInfo(info AppInfo) {
	// TODO dynamically adjust spaces and line length
	fmt.Println(info.Name())
	fmt.Println("------------------------------")
	fmt.Println("Version:          " + info.Version())
	fmt.Println("Build Timestamp:  " + info.BuildTime())
}

func CheckPrintInfo() {
	args := os.Args[1:]
	if len(args) == 1 {
		arg := strings.ReplaceAll(args[0], "-", "")
		if arg == "v" || arg == "version" {
			GetAppInfo().Print()
			os.Exit(0)
		}
	}

}
