package crawler

import (
	"filesystem"
	"io/ioutil"
	"os"
)

func Scan(root string, d int, h, s bool) (*filesystem.Tree, error) {
	var fs filesystem.Tree
	scanRoot(&fs, root, d)
	return &fs, nil
}

func scanRoot(fs *filesystem.Tree, f string, d int) error {
	fi, err := os.Stat(f)
	if err != nil {
		return err
	}
	fs.Name = fi.Name()
	fs.IsDir = fi.IsDir()
	if !fi.IsDir() {
		return NoDirErr
	}
	fs.SetChildren(scan(f, d))
	return nil
}
func scan(f string, d int) *filesystem.Tree {
	if d <= 0 {
		return nil
	}
	fi, err := ioutil.ReadDir(f)
	if err != nil {
		return nil
	}
	if len(fi) == 0 {
		return nil
	}
	var fs filesystem.Tree
	scanNext(&fs, f, fi, d)

	return &fs
}
func scanNext(fs *filesystem.Tree, root string, fi []os.FileInfo, d int) {
	if len(fi) == 0 {
		return
	}
	fs.Name = fi[0].Name()
	fs.IsDir = fi[0].IsDir()
	if fs.IsDir {
		fs.SetChildren(scan(root+"/"+fs.Name, d-1))
	}
	var fsnew filesystem.Tree
	if len(fi) > 1 {
		scanNext(fs.SetNext(&fsnew), root, fi[1:], d)
	}
}
