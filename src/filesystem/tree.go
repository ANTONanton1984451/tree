package filesystem

type Tree struct {
	IsDir    bool
	Name     string
	Next     *Tree
	Children *Tree
}

func (t *Tree) SetNext(fs *Tree) *Tree {
	t.Next = fs
	return fs
}
func (t *Tree) SetChildren(fs *Tree) *Tree {
	t.Children = fs
	return fs
}

func (t Tree) GetNext() *Tree  {
	return t.Next
}

func (t Tree) GetChildren() *Tree {
	return t.Children
}
