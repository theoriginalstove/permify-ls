package cache

// A Snapshot represents the current for a given view.
//
// It is an implementation of file.Source where ReadFile
// method returns consistent information about the existence of and
// contents of each file throughout its lifetime.
type Snapshot struct {
	// sequenceID is the monotonically increasing sequence ID of this
	// snapshot within its view.
	// They cannot be compared from different views.
	sequenceID uint64
}
