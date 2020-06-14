package archiver

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	files := []string{"testdata/test.tar", "testdata/test.tar.br", "testdata/test.tar.bz2", "testdata/test.tar.gz", "testdata/test.tar.lz4", "testdata/test.tar.xz", "testdata/test.tar.zst", "testdata/test.zip"}
	for _, fileName := range files {
		var (
			file   *os.File
			r      *bufio.Reader
			format string
			offset int
			err    error
		)
		if file, err = os.Open(fileName); err != nil {
			t.Errorf("open %s error!", fileName)
			t.Errorf("%v", err)
		}
		r = bufio.NewReader(file)
		if fileName == "testdata/test.tar" {
			offset = 257
		}
		if format, err = Format(r, offset, fileName); err != nil {
			t.Errorf("%v", err)
		}
		if err := file.Close(); err != nil {
			t.Errorf("%v", err)
		}
		if !strings.HasSuffix(fileName, format) {
			t.Errorf("Is file %s %s ?", fileName, format)
		}
	}
	return
}
