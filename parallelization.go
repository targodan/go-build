package make

// ParallelTargets contains Targets that will be
// executed in paralell.
type ParallelTargets struct {
	SubTargets []Target
}

// Execute executes all sub targets in parallel.
func (t *ParallelTargets) Execute(suite *Suite) error {
	errors := make(chan error, len(t.SubTargets))

	for _, st := range t.SubTargets {
		go func(st Target, errors chan<- error) {
			errors <- st.Execute(suite)
		}(st, errors)
	}

	errs := make([]error, 0)
	for i := 0; i < len(t.SubTargets); i++ {
		err := <-errors
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return &MultiError{Errors: errs}
	}
	return nil
}

// Parallelize combines Targets to be executed in parallel.
func Parallelize(targets ...Target) Target {
	return &ParallelTargets{SubTargets: targets}
}

// SequentialTargets contains targets that will be executed
// sequentailly. If AbortOnFirstError is false any errors will
// be recorded and returned but later Targets will still be
// executed.
type SequentialTargets struct {
	SubTargets        []Target
	AbortOnFirstError bool
}

// Execute executes all sub targets sequentailly.
func (t *SequentialTargets) Execute(suite *Suite) error {
	errs := make([]error, 0)

	for _, st := range t.SubTargets {
		err := st.Execute(suite)
		if err != nil {
			if t.AbortOnFirstError {
				return err
			}

			errs = append(errs, err)
		}
	}

	if !t.AbortOnFirstError && len(errs) > 0 {
		return &MultiError{Errors: errs}
	}
	return nil
}

// Concatenate concatenates Targets to be executed sequentailly.
// If AbortOnFirstError is false any errors will
// be recorded and returned but later Targets will still be
// executed.
func Concatenate(abortOnFirstError bool, targets ...Target) Target {
	return &SequentialTargets{SubTargets: targets, AbortOnFirstError: abortOnFirstError}
}
