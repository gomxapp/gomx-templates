package api

import (
	"bufio"
	"github.com/gomxapp/gomx"
	"github.com/gomxapp/gomx/config"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	gomx.RegisterOnPath("/template/{file}", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.PathValue("file") + ".tmpl"
		fp := filepath.Join(config.ApiRootDir, "templates", requestedFile)
		header, body, err := readFile(fp)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			_, _ = w.Write([]byte("Internal server error"))
			return
		}
		for key, value := range header {
			h := w.Header()
			h.Add(key, value)
		}
		w.WriteHeader(200)
		w.Write(body)
	}))
}

type FileHeader map[string]string
type FileBody []byte

// readFile gets a file based on the name without .gotmpl
func readFile(filepath string) (FileHeader, FileBody, error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, nil, err
	}
	parts := [2]string{"", ""}
	s := bufio.NewScanner(file)
	i := 0
	for s.Scan() {
		if s.Text() == "---" {
			i++
			continue
		}
		parts[i] += s.Text() + "\n"
	}
	// read header
	fh := make(FileHeader)
	headerString := parts[0]
	headers := strings.Split(headerString, "\n")
	for _, header := range headers {
		headerKV := strings.SplitN(header, ":", 2)
		if len(headerKV) != 2 {
			continue
		}
		k := strings.Trim(headerKV[0], " \r\n")
		v := strings.Trim(headerKV[1], " \r\n")
		fh[k] = v
	}
	// read body
	body := FileBody(parts[1])

	return fh, body, nil
}
