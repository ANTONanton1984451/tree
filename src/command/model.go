package command

import (
	"crawler"
	"flag"
	"fmt"
	"os"
	"reflector"
)

type Command struct {
	maxDepth      int
	includeHidden bool
	includeSystem bool
	root          string
}

func (c Command) Exec() error {
	err := c.parseFlags()
	if err != nil {
		return err
	}
	tree, err := crawler.Scan(c.root, c.maxDepth, c.includeHidden, c.includeSystem)
	if err != nil {
		return err
	}
	v := reflector.MakeView(tree)
	fmt.Fprint(os.Stdout, v)
	return nil
}

func (c *Command) parseFlags() error {
	flag.IntVar(&c.maxDepth, "d", -1, "max depth")
	flag.BoolVar(&c.includeHidden, "H", false, "include hidden files")
	flag.BoolVar(&c.includeSystem, "S", false, "include system files")
	flag.StringVar(&c.root, "root", "", "root directory")
	flag.Parse()
	if c.root == "" {
		return EmptyRootErr
	}
	return nil
}

func New() Command {
	return Command{}
}
