package main

import (
	"github.com/targodan/go-make"
)

func main() {
	all := make.PlatformSet{
		make.LinuxX386,
		make.LinuxAmd64,
		make.WindowsX386,
		make.WindowsAmd64,
	}
	suite := make.NewBuildSuite(all)

	buildTargets := make.MultiPlatformBuild(&make.BuildTarget{
		ExecutableName:      make.DefaultNameTemplate("test"),
		Version:             make.VersionFromGit("."),
		VersionVariableName: "main.version",
	}, all)

	for _, target := range buildTargets {
		suite.RegisterTarget(target)
		suite.RegisterTarget(make.CleanTargetsFromOutputTargets(target)[0])
	}

	app := make.CLIApp(suite)
	app.RunAndExitOnError()
}
