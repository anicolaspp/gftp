package fs

// Virtual is a way to abstracts multiple file systems.
// We need this since the underlaying file system will change depending on the
// cloud provider.
type Virtual interface {
	Ls(path string) ([]*Dir, error)
	Read(path string) ([]byte, error)
	Write(path string, buffer []byte) (int, error)

	Create(meta *DirMeta) (string, error)

	// Root is a mapping from real file system to the virtual file system.
	Root() string
}

// Dir represents a directory on the virtual file system.
type Dir struct {
}

// DirMeta represents the associated metadata of a directory.
type DirMeta struct {
}

// Mem is an implementation of a virtual file system that is completely in
// memory.
//
// We can use for testing most of the functionality related to file
// manipulation.
type Mem struct {
}
