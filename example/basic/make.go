package main

import (
	"fmt"
	"os"

	"github.com/targodan/go-make"
)

type V struct{}

func (v V) String() string {
	return "1.0.0"
}

func main() {
	var err error

	all := make.PlatformSet{
		make.LinuxAmd64,
		make.WindowsAmd64,
	}
	suite := make.NewBuildSuite(all)

	buildTargets := make.MultiPlatformBuild(&make.BuildTarget{
		ExecutableName:      make.DefaultNameTemplate("test"),
		Version:             make.VersionFromGit("."),
		VersionVariableName: "main.version",
	}, all)

	err = suite.Execute(make.Parallelize(make.ConvertBuildTargetSlice(buildTargets)...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-2)
	}
}
