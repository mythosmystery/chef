package workspace

import "io/fs"

// Walk walks the project tree respecting gitignore rules.
func (w *Workspace) Walk(fn fs.WalkDirFunc) error {
	panic("not implemented")
}
