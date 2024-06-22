package api

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_readFile(t *testing.T) {
	cwd, err := os.Executable()
	if err != nil {
		t.Error(err)
	}
	fp := filepath.Join(cwd, "templates", "go.mod", ".gotmpl")
	headers, body, err := readFile(fp)
	if err != nil {
		t.Error(err)
	}
	v := headers["version"]
	t.Log("Got header v: " + v)
	t.Log(body)
}
