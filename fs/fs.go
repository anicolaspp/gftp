package fs

import (
	"fmt"
	"os"
	"path"
)

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

	MountPoint() string
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
	Mount string
	User  string

	currentPath string
}

// NewMemFileSystem creates a instance of the in memory file system mounted
// for a specific user.
func NewMemFileSystem(mountPoint, user string) (*Mem, error) {
	vfs := &Mem{
		Mount: mountPoint,
		User:  user,
	}

	if err := os.Mkdir(vfs.Root(), os.ModePerm); err != nil {
		return nil, err
	}

	vfs.currentPath = vfs.Root()
	return vfs, nil
}

// Root is the root path for the virtual file system and the given users.
func (m *Mem) Root() string {
	return path.Join(m.MountPoint(), m.User)
}

// MountPoint is the mount point in the underlaying file system.
func (m *Mem) MountPoint() string {
	return m.Mount
}

func (m *Mem) Create(meta *DirMeta) (string, error) {
	return "", fmt.Errorf("implement me")
}
