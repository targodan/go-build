package build

import "os"

// CleanTargetNamePrefix is the prefix all CleanTargets
// will have in theire name.
const CleanTargetNamePrefix = "clean_"

// CleanTarget removes the file with the given filename.
type CleanTarget struct {
	// The name of the file to be removed.
	Filename string
	// The optional Platform the file was created for.
	Platform *Platform
}

// MakeCleanTargets creates clean targets from NamedOutputTargets.
func MakeCleanTargets(targets ...OutputTarget) []*CleanTarget {
	cleans := make([]*CleanTarget, len(targets))

	for i, t := range targets {
		cleans[i] = &CleanTarget{Filename: t.OutputName()}
	}

	return cleans
}

// Execute removes the file with the given filename.
func (t *CleanTarget) Execute(suite *Suite) error {
	err := os.Remove(t.Filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Name returns the name of this Target.
// The name will consist of the CleanTargetNamePrefix
// and the Platform name if present or the filename otherwise.
func (t *CleanTarget) Name() string {
	var postfix string
	if t.Platform == nil {
		postfix = t.Filename
	} else {
		postfix = t.Platform.String()
	}
	return CleanTargetNamePrefix + postfix
}
