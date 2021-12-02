package reflector

import (
	"bytes"
	"filesystem"
	"fmt"
	"io"
	"strings"
)

const tmpl = "|___%s\n"
const space = `    `

func MakeView(fs *filesystem.Tree) io.Reader {
	buf := bytes.NewBuffer([]byte{})
	makeView(buf,fs,0)
	return buf
}

func makeView(buffer *bytes.Buffer, fs *filesystem.Tree, d int) {
	space := strings.Repeat(space, d)
	buffer.Write([]byte(space + fmt.Sprintf(tmpl, fs.Name)))
	if f := fs.GetChildren(); f != nil {
		makeView(buffer,f,d+1)
	}
	if f:= fs.GetNext();f != nil {
		makeView(buffer,f,d)
	}
}
