package all_

import (
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ne(e error) {
	if e != nil {
		panic(e)
	}
}

func path() string {
	pkg, e := build.Import("e8vm.net/leaf/leaf/tests", "", build.FindOnly)
	ne(e)
	return pkg.Dir
}

func listDirs() []string {
	p := path()
	dir, e := os.Open(p)
	ne(e)

	list, e := dir.Readdir(0)
	ne(e)

	var ret []string
	for _, f := range list {
		if !f.IsDir() {
			continue
		}
		name := f.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}
		if strings.HasSuffix(name, "_") {
			continue
		}
		ret = append(ret, filepath.Join(p, name))
	}

	return ret
}

func TestAll(t *testing.T) {
	dirs := listDirs()
	for _, dir := range dirs {
		testCase(t, dir)
	}
}

func testCase(t *testing.T, dir string) {
	src, e := ioutil.ReadFile(filepath.Join(dir, "main.l"))
	ne(e)

	out, e := ioutil.ReadFile(filepath.Join(dir, "out"))
	if e != nil {
		out = nil
	}

	if out != nil {
		testOut(t, src, out)
	}
}
