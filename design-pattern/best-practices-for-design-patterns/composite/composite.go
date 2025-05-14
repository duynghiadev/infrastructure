package composite

import "fmt"

// FileSystemNode File system node interface
type FileSystemNode interface {
	List() string
	Add(node FileSystemNode)
	Remove(node FileSystemNode)
}

// File Leaf node: file
type File struct {
	name string
	size int
}

func (f *File) List() string {
	return fmt.Sprintf("File: %s (%dKB)", f.name, f.size)
}

func (f *File) Add(node FileSystemNode)    {}
func (f *File) Remove(node FileSystemNode) {}

// Directory Container node: directory
type Directory struct {
	name     string
	children []FileSystemNode
}

func (d *Directory) List() string {
	list := fmt.Sprintf("Directory: %s\n", d.name)
	for _, child := range d.children {
		list += fmt.Sprintf("  ├─ %s\n", child.List())
	}
	return list
}

func (d *Directory) Add(node FileSystemNode) {
	d.children = append(d.children, node)
}

func (d *Directory) Remove(node FileSystemNode) {
	for i, child := range d.children {
		if child == node {
			d.children = append(d.children[:i], d.children[i+1:]...)
			return
		}
	}
}
