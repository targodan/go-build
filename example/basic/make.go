package main

import (
	"fmt"
	"os"

	"github.com/targodan/go-build"
)

type V struct{}

func (v V) String() string {
	return "1.0.0"
}

func main() {
	var err error

	all := build.PlatformSet{
		build.LinuxAmd64,
		build.WindowsAmd64,
	}
	suite := build.NewBuildSuite(all)

	buildTargets := build.MultiPlatformBuild(&build.BuildTarget{
		ExecutableName:      build.DefaultNameTemplate("test"),
		Version:             build.VersionFromGit("."),
		VersionVariableName: "main.version",
	}, all)

	err = suite.Execute(build.Parallelize(build.ConvertBuildTargetSlice(buildTargets)...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-2)
	}
}
