package build

// Target represents a target that can be executed.
// Targets can be whatever you want, like build-, clean-
// or other targets.
type Target interface {
	Execute(suite *Suite) error
}

// OutputTarget represents a Target that knows the name of
// its output filename.
type OutputTarget interface {
	Target

	// OutputName returns the name of the output file.
	OutputName() string
}

// NamedTarget is a Target with a name.
type NamedTarget interface {
	Target

	// Name returns the name of the Target.
	Name() string
}

// NamedOutputTarget represents a Target that has a name and
// knows the name of its output filename.
type NamedOutputTarget interface {
	Target

	// OutputName returns the name of the output file.
	OutputName() string
	// Name returns the name of the Target.
	Name() string
}
