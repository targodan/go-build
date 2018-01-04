package build

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// BuildTargetNamePrefix is the prefix all BuildTargets
// will have in theire name.
const BuildTargetNamePrefix = "build_"

// BuildTarget is an implementation for a build target that will
// build an executable from the main-package in the current
// working directory.
type BuildTarget struct {
	// ExecutableName defines the templated name of the resulting
	// executable. It will receive the Platform as ".", so take
	// a look at the Platform struct for usable variables.
	ExecutableName *template.Template

	Version Version
	// Full name of the Variable holding the version string.
	// E. g. "main.version".
	VersionVariableName string

	Platform             *Platform
	AdditionalBuildFlags []string

	// Where to redirect the build commands stdout. If nil
	// stdout will be redirected to this precesses stdout.
	Stdout io.Writer
	// Where to redirect the build commands stderr. If nil
	// stderr will be redirected to this precesses stderr.
	Stderr io.Writer
}

// Copy copies a BuildTarget.
func (t *BuildTarget) Copy() *BuildTarget {
	copy := *t
	return &copy
}

// MultiPlatform returns one BuildTarget based on
// the current build target for each platform.
func (t *BuildTarget) MultiPlatform(platforms PlatformSet) []*BuildTarget {
	newTargets := make([]*BuildTarget, len(platforms))
	for i, platform := range platforms {
		newTargets[i] = t.Copy()
		newTargets[i].Platform = platform
	}
	return newTargets
}

// MultiPlatformBuild returns one BuildTarget based on
// the base build target for each given platform.
func MultiPlatformBuild(baseTarget *BuildTarget, platforms PlatformSet) []*BuildTarget {
	return baseTarget.MultiPlatform(platforms)
}

// ConvertBuildTargetSlice converts a *BuildTarget slice to a Target slice.
func ConvertBuildTargetSlice(buildTargets []*BuildTarget) []Target {
	ret := make([]Target, len(buildTargets))
	for i := range buildTargets {
		ret[i] = buildTargets[i]
	}
	return ret
}

// Execute build the executable.
func (t *BuildTarget) Execute(suite *Suite) error {
	if err := suite.CheckPlatform(t.Platform); err != nil {
		return err
	}

	executableName := t.OutputName()

	cmd := t.makeCommand(executableName)

	fmt.Println("Building binary:", executableName)
	if err := cmd.Run(); err != nil {
		log.Fatalln("Error running go build:", err)
	}
	return nil
}

func (t *BuildTarget) makeCommand(executableName string) (cmd *exec.Cmd) {
	if t.VersionVariableName != "" && t.Version != nil {
		ldflags := fmt.Sprintf("-ldflags=-X %s=%s", t.VersionVariableName, t.Version)
		cmd = exec.Command("go", "build", ldflags, "-o", executableName)
	} else {
		cmd = exec.Command("go", "build", "-o", executableName)
	}

	if t.Stdout != nil {
		cmd.Stdout = t.Stdout
	} else {
		cmd.Stdout = os.Stdout
	}
	if t.Stderr != nil {
		cmd.Stderr = t.Stderr
	} else {
		cmd.Stderr = os.Stderr
	}

	cmd.Env = os.Environ()
	cmd.Env = setEnv(cmd.Env, "GOOS", t.Platform.OS.String())
	cmd.Env = setEnv(cmd.Env, "GOARCH", t.Platform.Arch.String())

	return cmd
}

// OutputName returns the name of the output file.
func (t *BuildTarget) OutputName() string {
	buf := &bytes.Buffer{}
	err := t.ExecutableName.Execute(buf, t.Platform)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// Name returns the name of this Target.
func (t *BuildTarget) Name() string {
	return BuildTargetNamePrefix + t.Platform.String()
}

// setEnv searches in a slice of environment variables with the form key=value
// for a key and if found it sets its value, otherwise it adds the pair.
func setEnv(environ []string, key, value string) []string {
	for i, env := range environ {
		if strings.Split(env, "=")[0] == key {
			environ[i] = fmt.Sprintf("%s=%s", key, value)
			return environ
		}
	}
	return append(environ, fmt.Sprintf("%s=%s", key, value))
}
