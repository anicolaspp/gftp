package fs

import (
	"fmt"
	"path"
	"strings"
)

// Virtual is a way to abstracts multiple file systems.
// We need this since the underlaying file system will change depending on the
// cloud provider.
type Virtual interface {
	Ls(path string) ([]string, error)
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
	Path string
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

	dirs map[string][]string
}

// NewMemFileSystem creates a instance of the in memory file system mounted
// for a specific user.
func NewMemFileSystem(mountPoint, user string) (*Mem, error) {
	vfs := &Mem{
		Mount: mountPoint,
		User:  user,
		dirs:  map[string][]string{},
	}

	vfs.currentPath = vfs.Root()
	vfs.dirs[vfs.currentPath] = make([]string, 0)
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
	paths := strings.Split(meta.Path, "/")

	for i := 0; i < len(paths)-1; i++ {
		m.dirs[paths[i]] = append(m.dirs[paths[i]], paths[i+1])
	}

	m.dirs[paths[len(paths)-1]] = make([]string, 0)

	return meta.Path, nil
}

func (m *Mem) Ls(path string) ([]string, error) {
	if path == "" {
		return m.dirs[m.currentPath], nil
	}
	p := fmt.Sprintf("%s/%s", m.currentPath, path)
	paths, ok := m.dirs[p]
	if !ok {
		return nil, fmt.Errorf("path %s not found", p)
	}
	return paths, nil
}
