package prototype

import (
	"fmt"
)

type Inode interface {
	print(string)
	clone() Inode
}

func Run() {
	fmt.Println("---------- Prototype Pattern ----------")

	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "folder2",
	}

	fmt.Println("\n Printing heirarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\n Printing heirarchy for cloned folder")
	cloneFolder.print("  ")
}
