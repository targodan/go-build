package make

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
)

// Version holds information about the version.
type Version interface {
	fmt.Stringer
}

// BasicVersion is a basic implementation of Version.
type BasicVersion string

func (v BasicVersion) String() string {
	return string(v)
}

// VersionFromGit returns a version from the git repository
// in the given path. This is either the tag name if present
// or short commit hash if not.
func VersionFromGit(path string) Version {
	cmd := exec.Command("git", "describe", "--tags", "--always")
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running git describe: %v", err)
	}
	versionString := strings.TrimSpace(string(out))

	version, err := version.NewVersion(versionString)
	if err != nil {
		return version
	}
	return BasicVersion(versionString)
}
