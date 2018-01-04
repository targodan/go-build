package make

import (
	"fmt"
	"strings"
)

// Suite represents the build suite of a product.
type Suite struct {
	SupportedPlatforms PlatformSet

	registeredTargets map[string]Target
}

// NewBuildSuite creates a new Suite.
func NewBuildSuite(supportedPlatforms PlatformSet) *Suite {
	return &Suite{
		SupportedPlatforms: supportedPlatforms,
		registeredTargets:  make(map[string]Target),
	}
}

// CheckPlatform returns nil if the given Platform is supported
// by the build suite. Otherwise an error is returned.
func (s *Suite) CheckPlatform(p *Platform) error {
	if ok, _ := s.SupportedPlatforms.Contains(p); !ok {
		return fmt.Errorf("platform %s is not supported by this software", p)
	}
	return nil
}

// Execute runs the given Target in the context of this build suite.
func (s *Suite) Execute(t Target) error {
	return t.Execute(s)
}

// RegisterTarget registers a NamedTarget that can later be
// executed via ExecuteNamedTarget.
func (s *Suite) RegisterTarget(target NamedTarget) {
	s.registeredTargets[target.Name()] = target
}

// RegisterTargets registers NamedTargets that can later be
// executed via ExecuteNamedTarget.
func (s *Suite) RegisterTargets(targets ...NamedTarget) {
	for _, target := range targets {
		s.RegisterTarget(target)
	}
}

// Lookup returns the previously registered target by name or
// nil if no such Target exists.
func (s *Suite) Lookup(targetName string) Target {
	return s.registeredTargets[targetName]
}

// LookupPrefix returns all registered Targets starting with the
// given prefix.
func (s *Suite) LookupPrefix(namePrefix string) []Target {
	ret := make([]Target, 0)
	for name, target := range s.registeredTargets {
		if strings.HasPrefix(name, namePrefix) {
			ret = append(ret, target)
		}
	}
	return ret
}

// LookupBuildTargets returns all registerd build targets.
func (s *Suite) LookupBuildTargets() []Target {
	return s.LookupPrefix(BuildTargetNamePrefix)
}

// LookupCleanTargets returns all registerd clean targets.
func (s *Suite) LookupCleanTargets() []Target {
	return s.LookupPrefix(CleanTargetNamePrefix)
}

type targetNotFoundError struct {
	name string
}

func (e *targetNotFoundError) Error() string {
	return "target \"" + e.name + "\" could not be found"
}

// IsNotFound returns true if the given error represents that the
// target was not found.
func IsNotFound(err error) bool {
	_, ok := err.(*targetNotFoundError)
	return ok
}

// ExecuteNamedTarget executes the previously registered target by name.
// Use the IsNotFound function to check if the returned error was an error
// during the execution or if the target name was not found.
func (s *Suite) ExecuteNamedTarget(targetName string) error {
	target := s.Lookup(targetName)
	if target == nil {
		return &targetNotFoundError{name: targetName}
	}
	return target.Execute(s)
}
