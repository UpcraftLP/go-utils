package main

import (
	"fmt"
	"github.com/karrick/godirwalk"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"utils/pkg/sysutils"
)

var (
	goCmd         = "D:\\Files\\GoLang\\go1.15.5\\bin\\go.exe"
	version       = "development-preview"
	buildFlags    = []string{"-s", "-w"}
	buildPath     = "./bin"
	verboseOutput = false
)

var buildTargets = []BuildTarget{
	{
		Arch: ARCH_i386,
		OS:   OS_WINDOWS,
	},
	{
		Arch: ARCH_amd64,
		OS:   OS_WINDOWS,
	},
	// FIXME npipe does not exist for linux
	//{
	//	Arch: ARCH_amd64,
	//	OS:   OS_LINUX,
	//},
	//{
	//	Arch: ARCH_i386,
	//	OS:   OS_LINUX,
	//},
	//{
	//	Arch: ARCH_ARM,
	//	OS:   OS_LINUX,
	//},
}

type BuildInfo struct {
	Name     string
	FullPath string
}

type BuildTarget struct {
	Arch Architecture    `json:"arch"`
	OS   OperatingSystem `json:"os"`
}

type Architecture string

//goland:noinspection ALL
var (
	ARCH_amd64 Architecture = "amd64" // 64-bit
	ARCH_i386  Architecture = "386"   // 32-bit
	ARCH_ARM   Architecture = "arm"
)

type OperatingSystem struct {
	Name        string
	PathBuilder func(binaryName string, arch Architecture) string
}

//goland:noinspection ALL
var (
	OS_WINDOWS = OperatingSystem{
		Name: "windows",
		PathBuilder: func(binaryName string, arch Architecture) string {
			return fmt.Sprintf("windows-%v/%v.exe", arch.name(), binaryName)
		},
	}
	OS_LINUX = OperatingSystem{
		Name: "linux",
		PathBuilder: func(binaryName string, arch Architecture) string {
			return fmt.Sprintf("linux-%v/%v", arch.name(), binaryName)
		},
	}
)

func main() {
	buildTime := fmt.Sprint(time.Now().Unix())
	log.Printf("Build time is %v.\n", buildTime)
	appVersion, prefixedAppVersion := checkVersionFromArgs()

	//fmt.Println(appVersion, prefixedAppVersion) // TODO formatting
	//fmt.Printf("Build targets:\n\t%v\n", strings.Join(buildTargets, "\n\t"))

	if err := os.RemoveAll(fmt.Sprintf("%v/%v", buildPath, appVersion)); err != nil {
		panic(err)
	}

	root := "cmd"
	children, err := godirwalk.ReadDirents(root, nil) // read direct children of root dir
	if err != nil {
		panic(err)
	}
	var buildInfos []BuildInfo
	for _, path := range children {
		if path.IsDir() {
			resolved, err := filepath.Abs(filepath.Join(root, path.Name()))
			if err != nil {
				panic(err)
			}
			buildInfos = append(buildInfos, BuildInfo{
				Name:     path.Name(),
				FullPath: resolved,
			})
		}
	}

	total := len(buildTargets)

	for i, target := range buildTargets {
		fmt.Printf("Building target %v/%v: %v\n", i+1, total, target.name())
		for _, info := range buildInfos {
			fmt.Println("Building", info.Name)

			if err := os.Setenv("GOOS", target.OS.Name); err != nil {
				panic(err)
			}
			if err := os.Setenv("GOARCH", target.Arch.name()); err != nil {
				panic(err)
			}



			buildCmd := getBuildCmd(info, target, prefixedAppVersion, buildTime)
			//goland:noinspection GoBoolExpressions
			if verboseOutput {
				fmt.Println("RUN", strings.Join(buildCmd, " "))
			}

			process, err := sysutils.Start(buildCmd...)
			if err != nil {
				panic(err)
			}

			state, err := process.Wait()
			if err != nil {
				panic(err)
			}
			if !state.Success() {
				panic(fmt.Sprintf("unknown error: %v", state.ExitCode()))
			}
		}
	}
}

func getBuildCmd(info BuildInfo, target BuildTarget, prefixedAppVersion, buildTime string) []string {
	var variables = make(map[string]string)
	variables["utils/internal/info.Name"] = info.Name
	variables["utils/internal/info.Version"] = prefixedAppVersion
	variables["utils/internal/info.BuildTime"] = buildTime
	var ldFlags = strings.Join(buildFlags, " ")
	if len(buildFlags) > 0 && len(variables) > 0 {
		ldFlags += " "
	}
	var variableFlags []string = nil
	for k, v := range variables {
		variableFlags = append(variableFlags, fmt.Sprintf("-X '%v=%v'", k, v))
	}
	ldFlags += strings.Join(variableFlags, " ")


return []string{goCmd, "build", "-o", fmt.Sprintf("%v/%v/%v", buildPath, version, target.OS.PathBuilder(info.Name, target.Arch)), "-ldflags", ldFlags, info.FullPath}
}

func checkVersionFromArgs() (string, string) {
	args := os.Args[1:]
	if len(args) > 0 {
		version = args[0]
	}
	if version == "development-preview" {
		log.Println("No version defined, using 'development-preview'!")
		return version, version
	} else {
		log.Printf("Version is %v\n", version)
		return version, fmt.Sprintf("v%v", version)
	}
}

func (target BuildTarget) name() string {
	return fmt.Sprintf("%v/%v", target.OS.Name, target.Arch)
}

func (arch Architecture) name() string {
	return string(arch)
}
