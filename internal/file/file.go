package file

import "context"

type Identity struct {
	Hash Hash
}

type Handle interface {
	Identity() Identity
	// SameContentsOnDisk reports whether the file has the same content on disk.
	// It is false for files open on an editor with unsaved edits.
	SameContentsOnDisk() bool
	// Version returns the file version, as defined by the LSP client.
	// for on-disk file handles, Version returns 0.
	Version() int32
	// Content returns the contents of a file.
	// If the file is not available, retuns a nil slice and and error.
	Content() ([]byte, error)
}

type Source interface {
	ReadFile(ctx context.Context) (Handle, error)
}
