package build

// PlatformSet is a set of Platforms.
type PlatformSet []*Platform

// Contains returns true and the contained platform if the set
// contains a Platform equivalent to the given one.
func (s PlatformSet) Contains(p *Platform) (bool, *Platform) {
	for _, pl := range s {
		if pl.Equals(p) {
			return true, pl
		}
	}
	return false, nil
}
