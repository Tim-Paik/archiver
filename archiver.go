package archiver

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
)

var (
	magicXz      = []byte{0xfd, 0x37, 0x7a, 0x58, 0x5a, 0x00}
	magiTar      = []byte{0x75, 0x73, 0x74, 0x61, 0x72} //offset by 257 bytes
	magicZstd    = []byte{0x28, 0xb5, 0x2f, 0xfd}
	magicZip     = []byte{0x50, 0x4b, 0x03, 0x04}
	magicLz4     = []byte{0x04, 0x22, 0x4d, 0x18}
	magicLz47zip = []byte{0x50, 0x2a, 0x4d, 0x18}
	magicLz4Java = []byte{0xf0, 0x1b, 0x7b, 0x22}
	magicGz      = []byte{0x1f, 0x8b}
	magicBzip2   = []byte{0x42, 0x5a}
)

//Format 返回所读文件的格式，如果找不到则依靠 fileName 的后缀确定，如果可以肯定有魔术数字则可以省略 fileName。offset 为魔术数字偏移值。
//Format returns the format of the file read. If it cannot be found, it depends on the suffix of fileName.
//If there is a magic number, fileName can be omitted. offset is the magic number offset value.
func Format(reader *bufio.Reader, offset int, fileName string) (format string, err error) {
	var (
		headerBytes []byte
		magic       []byte
	)
	if headerBytes, err = reader.Peek(offset + 6); err != nil {
		return "", err
	}
	magic = headerBytes[offset : offset+6]
	if bytes.Equal(magicXz, magic) {
		return "xz", nil
	}
	if bytes.Equal(magiTar, magic[0:5]) {
		return "tar", nil
	}
	if bytes.Equal(magicZstd, magic[0:4]) {
		return "zst", nil
	}
	if bytes.Equal(magicZip, magic[0:4]) {
		return "zip", nil
	}
	if bytes.Equal(magicLz4, magic[0:4]) ||
		bytes.Equal(magicLz47zip, magic[0:4]) ||
		bytes.Equal(magicLz4Java, magic[0:4]) {
		return "lz4", nil
	}
	if bytes.Equal(magicGz, magic[0:2]) {
		return "gz", nil
	}
	if bytes.Equal(magicBzip2, magic[0:2]) {
		return "bz2", nil
	}
	if filepath.Ext(fileName) == ".br" {
		return "br", nil
	}
	if filepath.Ext(fileName) == ".sz" {
		return "sz", nil
	}
	return "", fmt.Errorf("unknown compression package format")
}
