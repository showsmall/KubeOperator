// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\xcb\x72\xe3\xba\x11\xdd\xf3\x2b\x60\xaa\xbc\x9b\xdc\xca\xda\x3b\x9a\xa2\x6d\x66\x28\x92\x45\x52\xbe\x71\x36\x2c\x88\x6c\x49\x88\x29\x80\x05\x40\xe3\xe8\xee\xf2\x5f\xf9\xa7\xfc\x42\xaa\x01\xf0\xa1\x87\x33\x9e\xba\xab\xb1\xa6\xd0\xdd\x07\xfd\x38\x7d\xc0\x45\x23\x0e\x07\xc1\xbd\x34\x58\x45\x75\xf4\xf7\xb8\xac\xca\x07\xe2\xa7\xf4\x00\x84\x76\x12\x68\x7b\x22\xf0\x2f\xa6\xb4\xf2\xbd\x38\xaf\xd3\xac\x9a\x0e\xe5\x1d\x50\x05\x64\xcb\xba\x8e\x30\x4e\xf4\x1e\x88\x84\x1d\x53\x5a\x9e\x48\x9c\x13\x61\xff\x4b\x9d\x94\x86\x03\x51\xa0\x35\xe3\x3b\xd2\xd3\x1d\xf8\x9e\xe7\x2d\x9a\xee\xa8\x34\x48\x2f\x4c\xd6\x65\x15\x15\xf5\x32\x4a\xa2\x2a\xaa\x9f\x82\x38\x89\x96\x0f\xc4\x6f\x28\x27\x5c\x68\xd2\x42\x07\x1a\x88\x3b\x8e\x81\x9a\xa3\x94\xc0\x35\x51\x9a\x6a\xf0\x47\x07\x71\x69\xe0\x15\xeb\x34\x8d\xd3\xe7\x07\xe2\x57\xfb\x99\x99\x32\xce\xe4\x91\x73\xc6\x77\x57\x46\x49\x16\x06\xc9\x03\xf1\xe3\x43\x2f\xa4\x1e\xad\x1a\xca\xd1\x6a\x03\xe4\xd8\xef\x24\x6d\xa1\xf5\x11\xb9\x84\x16\xb8\x66\xb4\xf3\xce\x40\xd7\x45\x54\x66\xeb\x22\x8c\x1e\x88\xff\x44\x59\x07\x2d\xd1\xc2\xe1\xbf\x23\xd5\x1e\x24\x20\x0e\xca\x09\x55\x4a\x34\x8c\x6a\x68\xc9\x5e\x28\x4d\x8e\xbc\x05\x49\xf4\x9e\x29\xf2\x0e\x27\xff\x13\xb7\xf5\x3f\xb2\xf4\x97\x7c\xff\x21\x38\xdc\xf0\xfd\x14\xac\x93\xaa\x0e\x8b\x68\x19\xa5\x55\x1c\x24\x75\x18\xa4\x26\x0b\x36\xec\x03\xf1\x97\xb0\xa5\xc7\x4e\x93\xe9\xa6\xb3\x54\xd8\xa0\xad\x6f\x5b\x26\x7c\x89\xc2\xef\x53\xd5\x4c\xce\x27\x2b\x8e\x7d\x34\x99\x9a\x7e\x30\xad\xa5\xcc\xdf\x47\x05\xd2\x9c\xf1\xbd\x3c\x28\xcb\xdf\xb3\x62\x39\x82\x49\xd7\x09\x56\xa4\xa7\x4a\x7d\x08\xd9\x92\xa1\x1f\x36\x40\x36\x1d\xe5\xef\xff\xfd\xcf\xbf\x7d\x2f\x2f\xe2\xd7\xa0\x8a\xea\xef\xd1\xdb\xa5\x21\x22\xe9\x25\xfb\x41\x35\xe0\xc5\x67\x28\x26\x73\x6f\x81\xe9\xf7\x5e\xb2\xb2\xaa\x83\xa4\x88\x82\xe5\xdb\xd4\xde\x2f\x58\x99\xcb\x19\x70\x95\x31\x16\xe3\xa5\x6f\x16\xc4\x56\x16\x6b\xe2\x5c\xcc\x0a\xf3\xc1\xf4\xde\x24\xc0\x35\xda\x2d\xbf\xf5\xe3\x5b\x9d\x17\xd9\xdf\xa2\xb0\xfa\x53\x21\x7a\x29\xfe\x09\x8d\xf6\xbd\xf2\xad\xac\xa2\x55\xed\xa6\xf8\x29\x5b\xa7\xcb\xdb\x43\xdc\x89\x86\x76\x38\xc1\x5b\x26\x95\x36\x89\x4a\x33\xb4\x0b\x5e\x83\x38\x09\x1e\x13\x6c\x91\x54\x90\xb8\x27\xf4\x07\x65\x1d\xdd\x74\xe0\x7b\x71\x69\xa7\xc8\xdc\x61\x36\xbf\xcc\x8e\x94\x75\x8a\x80\x7d\x9b\xef\x78\x95\x67\x45\x55\x47\x45\x91\x15\x43\xcd\x52\x41\x5a\xaa\x29\x5e\xd3\x99\x7d\x50\x45\xb6\xe2\xc8\xdb\x3b\xe2\x90\x36\x7b\x68\xde\x0d\x4e\x77\x64\xcb\x3a\xb8\x3b\x77\x8a\xee\xea\xd7\x20\x59\x23\xd2\xe8\xd0\xeb\x93\xf5\x2b\x38\xe9\x18\x07\x72\xaf\xce\xcf\xff\x5e\x64\xe9\x73\xfd\x94\x15\xab\x00\xa1\xc7\xbc\x11\x52\x42\xa3\x89\x0d\x20\xe4\x81\xea\x4f\x8d\x67\x83\x34\x4f\x6c\x38\x9b\x02\xa1\xed\x25\x3e\xf5\xe1\x4a\x9e\xae\x57\x0f\xc4\x0f\x88\x16\x9a\x76\x44\x6c\xc9\xbd\x22\x52\x7c\x28\xfc\xd3\x5c\x80\x4a\x20\x74\xc3\x11\x50\xf7\x8d\xa8\x77\xd6\xdf\xf4\x53\x06\xaf\xe7\x44\xa1\xe8\x0f\xd7\x2d\xf7\xea\x1b\x91\x40\x95\xe0\x0f\x08\xc2\xf3\x16\x38\x84\xd3\xf8\xe1\x15\x56\x41\x15\xbe\x0c\x23\x34\xcc\x1f\x53\x84\x0d\x79\xf1\xbd\xac\x88\x9f\xe3\xd4\x5d\x79\x7e\x5e\x48\xb6\x63\x9c\x76\x9f\x19\xae\xcb\x89\x75\x83\xb0\x8a\x0d\xd0\x6a\xe0\x02\x47\xd3\xc0\xb1\xa7\x66\x35\x17\x5c\xd3\x46\x9b\xaa\xd3\xf6\xc0\x38\x2e\x19\xaa\x85\xbc\x73\x0e\xe7\x89\x4f\x05\x51\xc7\x66\x6f\x1c\x9a\xe6\x0d\x96\xab\x38\xbd\x26\x39\x0c\xda\x3a\xa2\x33\x4e\x2d\x84\x2b\xa2\xbb\x3b\x07\x5d\x44\x49\x50\x45\xcb\xd9\x6c\xae\xd1\x6c\x4f\x11\xfa\x7c\x02\xdd\xe0\x19\x08\xc9\x32\xc8\x47\x04\xeb\x7c\x19\x8c\x08\xba\x96\xf6\x97\x81\xa1\x65\x36\xee\x6b\x54\xc4\x4f\x6f\x75\x98\x2d\x67\x8b\xf1\x15\x24\xdb\xb2\x86\x6a\x26\x38\x69\x44\x0b\x04\xa4\x14\xd2\xf7\xa2\x55\x10\x27\xf5\x32\x2e\xdd\x88\xae\x28\xeb\x86\xbd\xab\x4c\xeb\xb4\x4c\x7d\x31\xb1\x83\xb7\x79\x79\xa3\x03\x3a\x3c\x50\xdd\xec\xc9\xd6\xb4\x96\xe5\x06\x5c\x03\x63\xff\x94\xf8\x6b\xc4\x8a\xa9\xf9\x3f\x3b\x60\xe8\x91\x4b\x27\x86\x14\x1e\x88\xff\x21\x05\xdf\x4d\x5b\x82\x08\x39\x33\xb1\x00\x0d\x5d\x8f\xe0\x2e\xe9\xda\x5b\xa0\x22\x11\x7c\xe0\xd7\x22\x7a\x8e\xb3\xf4\xab\xfb\x9a\x58\xe3\x9f\x31\x2c\xae\x59\x0c\x85\xff\x0e\x81\x70\x55\x7f\x39\x8c\xd9\xd3\x3f\xa3\xf1\x8e\xf2\x73\xd9\x62\x29\x33\xb4\x89\xdd\x81\x9e\x2f\x94\x1b\x6c\xd9\x08\xbe\x65\xbb\xa3\x34\x7d\x63\x0a\x17\xaf\x82\xe7\xe8\x73\x57\xec\x40\x77\xf0\x35\x47\x79\x5d\xbe\x64\x85\xa5\x4e\x75\xdc\x6e\x59\xc3\x50\xa0\xc5\x3d\xa6\x45\xf4\xc0\x95\xa6\xcd\xbb\xf7\x1c\x55\x43\x05\x86\x0a\xa7\x62\x48\xb2\x21\x48\x3c\xef\xe6\x66\x05\x87\x0d\xc8\x71\xf4\x82\x25\xf6\xd3\xbd\x22\xe3\xb4\x6d\x00\x38\xa1\xad\x11\x65\xf3\x01\x1d\x78\xe0\x5e\x9d\x71\x8a\xf1\xef\xb6\xbe\x0b\x31\x6a\x21\x37\xcb\x9f\x0b\x21\x67\x70\x4b\x05\x0d\xb6\x2f\x41\x59\xbb\xf2\xe0\x5a\x75\x06\xb3\x52\x8e\xa5\x09\x6f\x30\x8c\xb7\xe0\xa2\x05\x2f\xc5\x49\x1f\x94\x88\x53\xb2\x75\x15\x94\xdf\x71\x2d\xb4\x2d\xc1\x43\x38\x05\x4e\x14\x9b\x9f\x43\xd7\x38\x6d\xfb\xad\xb7\x05\xfb\xa0\x4c\x13\xa6\x49\x2b\x38\xfc\x86\x01\x36\xb4\x79\x3f\xf6\x41\xd3\x88\x23\xd7\x5e\x1e\x14\xc1\xaa\x8e\x56\x79\xf5\x76\x59\xb6\x9e\x4a\x7a\x00\x0d\x52\xe1\xe2\xaf\xea\x72\x9d\xe7\xb6\xba\x6b\xae\x8e\x3d\xee\x44\xec\xe1\x53\x8f\xe2\xfb\x5c\xfe\x9d\x71\x93\xe5\x88\x51\xdb\x3c\x06\xe1\xf7\x75\x5e\x07\x61\x98\xad\xd3\x5f\x51\x39\x67\xc0\xbf\x2c\x77\xbc\x05\x8e\xcc\x85\x94\xfe\x42\x34\xb4\xfa\x85\x20\xae\xaa\x8f\x06\xa3\xe7\xee\xf8\x14\x27\x51\x39\x17\xa3\x6e\x88\x5c\xd9\xf4\x78\x29\x23\x5e\xc8\x06\xb6\x42\x02\x51\x1f\x4c\x37\x7b\x7c\x24\xcd\x0e\x50\x7b\xed\xb3\xd1\xb7\x51\xae\x5f\x38\x1b\x40\x63\x34\x84\x96\x1c\x7b\xd3\xec\x33\xb3\x22\x2a\xab\xac\x88\xae\xed\x24\x28\x2d\x24\xe3\xbb\xf3\xf1\x28\x40\x89\xa3\x6c\xe0\x3a\x87\xb3\x6b\xfe\xf4\x72\x93\x88\xbc\xad\x71\xa7\x99\x19\x15\xed\x90\xfa\x0d\x74\x02\x57\x97\x16\xe7\xdc\x56\xe1\x63\x46\xf4\x20\xdd\x0e\x1c\xe7\xa9\x07\x89\x62\xcd\x4c\x94\x95\x45\x57\x94\xf0\xe2\x24\xf3\x48\x09\xbe\x37\xbe\xaf\x2c\xcd\x44\xee\xe4\x70\x7d\xd3\x0c\x96\x66\xf2\x24\x48\x6f\xf8\xcc\x5d\xcb\xcc\x7c\x5e\x34\xfb\xb5\xcd\xe3\x65\x53\xcf\x8c\xbd\x05\xaa\x02\x2b\x1a\x46\x66\xb3\x8d\x54\x9e\x78\xb3\x97\x82\xb3\x3f\xb0\xc4\x0a\xa4\x5d\xec\x7f\x75\x12\x23\xc9\x9e\xe3\xf4\xd2\x66\x3d\x57\x56\xb8\x18\xef\xdc\xe9\x49\x2a\x54\xd3\x43\xbd\x97\x62\xcf\x36\x4c\x2b\x82\x67\x5c\x8c\xad\x14\x07\xd2\x89\xdd\x0e\x1b\x8c\xf1\xdf\xbe\x22\xcc\xbc\x45\xc3\x94\x17\xc6\xa5\x21\xb0\x4b\x56\xc3\x17\x02\x53\x44\x53\xf5\x7e\xc9\x60\x68\xfa\xe3\x10\x9a\x35\xe3\xbd\xae\xea\x30\x4b\x9f\xe2\xe7\xe9\x61\x16\xce\x17\xd0\xd5\xca\x9f\x0c\x2e\xbf\x28\x54\x97\xcb\xeb\xb3\x46\x6b\xa1\xef\xc4\xe9\x60\xb8\xb0\xa3\xfc\x8b\x0d\xe7\x2d\x58\x8f\xab\x70\xc4\x89\xf1\x80\x6b\x90\xd0\xe2\x6b\x4a\xc1\xce\xb8\x44\x08\x1d\x6b\xb4\x9a\xf8\xc4\x60\xc7\xd4\x4e\xc7\xbe\x91\xfe\xfc\x65\x46\x77\x94\xf1\xf1\x23\xcc\xfc\x29\x16\xe7\xf8\x4e\x40\x84\x4d\x73\xec\x19\xb4\x84\xf2\x76\x06\x52\x82\xf1\xd4\x5a\x88\x7f\xe9\x85\xe8\xd0\x4d\x9e\x65\xc9\xcd\x1c\xc5\x39\xc1\x33\xb3\x25\x7b\x83\x02\xc7\x97\x9f\x55\x2f\xe7\x11\xc7\xe5\x68\xf5\x97\xd2\xf2\xe4\xe1\xe2\x2f\xab\xe2\xed\xfa\xa1\x5d\xcd\x3f\x1c\x89\xad\xfd\x56\x41\x65\xb3\x67\x1a\x1a\x7d\x94\x60\x76\xcd\xb5\xb6\xfb\x5f\x00\x00\x00\xff\xff\xe1\x15\x70\x08\xbb\x12\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 4795, mode: os.FileMode(420), modTime: time.Unix(1617263005, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\x5f\x57\xe2\x58\x12\x7f\xcf\xa7\xe0\xc0\x99\xb7\xdd\x3d\xfb\xdc\x6f\xb7\x43\xd4\x6c\x87\x24\x27\x09\xce\xba\x2f\x39\x0e\xcd\xce\xba\xa3\xc0\x11\xdd\x73\x76\x9f\x86\xf6\x0f\xe0\x10\x71\xbb\xd5\x6d\x95\x69\xc1\xd6\x6e\x47\x07\xf0\xcf\xb4\xd0\x80\xfa\x65\x72\x6f\xc2\x93\x5f\x61\xcf\xbd\x37\x09\x17\x68\x7b\x7c\x14\xeb\x57\x55\xb7\xaa\x7e\xbf\xaa\x44\x12\xe9\x85\x85\x74\x8a\x93\x41\x4c\x30\x85\xbf\x8a\xba\xa1\x3f\x0b\x85\xe1\x96\xe5\x7c\xbc\x80\xad\x2b\x58\x7f\x0b\x2b\xa7\x61\x4e\x54\x4d\x59\x31\x06\x06\x6e\xb3\x05\x2b\xa7\xce\x75\xd7\xe9\x1e\xba\x8d\x3b\xe7\xb6\xd1\xaf\x7d\xea\xff\x7c\x04\x6b\xe7\x70\x7d\xcf\xee\xbe\x81\x9d\x37\xa2\x1a\xe6\xb8\x48\x62\x7e\x39\xbb\x94\x5c\xe4\x78\x29\xae\x1b\x82\x66\x46\x05\x49\x30\x04\x73\x02\x88\x92\x10\x7d\x16\x0a\xa3\xff\x55\xd1\xf5\x0e\x2c\x54\xfb\x7b\xc7\xf0\xf6\x0d\x2c\x5a\xce\xc6\x0d\xfa\x31\xe7\xec\xaf\xf6\x0f\xd6\x9d\xbb\xe3\x70\x00\x15\x75\x92\x84\x16\x97\x65\x51\x9e\x7c\x16\x0a\x53\x03\xbb\x6d\xc1\xca\xa9\x7b\xbf\xe5\xd6\x4a\x76\xbb\xfe\xd0\xcb\x8d\x41\x24\x85\x07\x12\x7e\x57\xb3\x07\xd7\x4e\x28\xcc\x0b\x6c\xe5\x9d\xce\x47\x92\xe8\x62\xf2\x65\x32\xb5\x34\x37\x3b\xcf\x0d\xe5\x68\x6a\x82\xae\xc4\x35\x5e\xc0\x78\x9a\xe6\xf1\xa5\xfb\xdb\xc9\x43\x2f\xe7\x36\x4f\x9c\x8f\x6f\xfb\xaf\x4f\xec\xf6\x4f\xa8\x52\x84\x6b\xd7\x6e\x6e\xdb\x6e\x77\x51\xa5\x13\x7e\xc4\x89\xf9\x37\x45\x7e\xaa\x27\x58\x6e\x3a\xdb\xa7\xb0\x44\x9c\x4d\x80\xb8\x64\x98\xbc\x26\x44\x05\xd9\x10\x81\x64\xf2\x40\x26\x6f\xa3\x71\x70\x35\xba\x6f\xdd\xc6\x31\xcc\xd7\x91\xd5\xb0\xdb\x96\xbb\x72\x4b\x83\x90\x82\x90\xfe\xf2\x53\x02\xff\x62\x50\x7a\x1a\x91\xf6\x9a\x02\xec\xf6\xa6\xb3\x7d\x8a\x0a\x2d\xfc\xe3\x41\x1b\x6e\x95\xc2\x9c\x0a\x74\xfd\x5b\x45\x8b\x06\x01\xe5\xb8\x44\x6b\xb9\xee\x54\x73\x3e\xae\xe3\xfc\xd2\x21\x81\x54\x4d\x9c\x06\x86\x60\xbe\x10\x66\x46\x11\xfe\x0b\x47\x10\x1c\x17\xf9\x47\x3a\xbb\xc4\x4d\x29\xba\x61\x02\x49\x13\x40\x74\x66\x30\x69\xb4\x9c\xcc\x28\x7a\x75\x25\xd6\xc1\x53\xc6\xcb\x19\xe0\x9c\x6e\x99\x96\xd3\x1f\xa7\x71\x07\xe6\xf3\x19\x53\xd5\x94\xbf\x08\xbc\xf1\x54\x5f\xb5\xcf\xce\x41\x83\xa4\xaf\xcf\xe8\x86\x10\x33\x3d\x86\x4c\x28\x71\x39\xea\x11\x64\xad\x40\xe9\x80\x2a\xbf\xa2\x4a\x47\x54\x69\x23\x14\x6c\x0a\xa6\x81\x28\x81\xe7\x12\xee\x9b\xa8\x86\xdc\x4f\xab\xa8\xb3\x85\x2b\x73\x73\x1d\xe6\x44\x9d\x0e\x2c\x49\x91\xf8\xf2\x32\xb0\xdb\x9b\x94\x5a\x21\x51\x0d\xc1\xf5\x2b\xe7\x2c\xf7\xd0\x2b\xa1\xca\x99\xbb\x72\x8b\x0a\x5b\x70\xe3\x10\xb5\xba\x70\xa3\x1a\xa6\xb5\x14\x63\xaa\xa2\x19\xa6\xa0\x69\x8a\xe6\xf7\x00\x55\xce\x50\xf1\x0e\x16\x2e\x60\xb9\x49\xd9\xe0\xec\xaf\xa2\x9d\x0b\x64\x35\xc8\x5b\x5b\xe8\xfd\x8f\xe8\xf0\x84\xfe\x0b\xed\xe6\xed\xee\x0d\x49\x9b\x75\x88\x5d\x99\xd3\x40\x8a\xe3\xec\xbf\xc9\x86\xdc\x5a\x09\x55\x8a\xce\x2f\x1d\xea\x67\xd8\xf8\x5b\x4d\x91\x27\xcd\x09\x45\x8b\x01\x23\x30\x77\xce\x9b\xb0\xfc\x1e\x55\x7b\xb0\x57\xb6\xdb\x16\xaa\xbf\x77\x6a\x23\x38\x66\xd2\xd9\xba\x7a\xe1\x8a\x77\x98\xf5\x85\x0b\xd8\x5c\xef\xbf\x3e\x19\x46\x7a\x3d\x95\xe3\x31\xdc\xcc\xb5\xcb\x90\x87\x21\xc9\xd1\x68\xb0\xdd\x7e\xe8\x95\xdc\xd6\xb5\x7b\x9f\xff\x22\x58\x07\xd3\x42\x30\x7c\xd8\x81\x7d\xff\x33\x1e\x40\x6f\x28\x4a\x70\xf3\x10\x1e\x54\x1f\x7a\xfb\xdf\x64\xc9\x00\x2f\x67\x93\x8b\x03\xa6\xe0\x84\x63\xc0\xe0\xa7\x02\x9a\xf4\xb7\xf7\xdc\x66\x33\xcc\x29\x9a\x38\x29\xca\xde\x9b\x02\x93\xcd\xc3\x61\xab\xb8\x3e\x90\x2e\xc0\x1b\x22\xc9\x85\x12\x13\x55\xce\xe0\x16\x96\x05\xda\x2d\x37\xb7\x8d\x45\xb8\x51\x73\xb6\xd6\xe1\x7f\xdf\x92\x56\x11\x34\x5b\x33\xac\x74\xf5\x63\x8a\x27\x16\x20\x1a\x13\xe5\xc7\x04\x24\x34\xfb\x72\x61\x2e\x15\xa2\xe6\x94\xac\xee\xd1\x39\x23\x25\x6c\x76\x9a\x20\x01\x43\x88\x32\xec\xf1\xd2\xbc\xaa\x05\x32\x46\xb9\x12\xe6\xa4\x28\x50\x83\xa0\x71\x35\x0a\x48\x50\xfc\xeb\x50\x30\xfb\xbe\x81\xb6\x3f\x93\x48\xd3\x82\x26\x4e\xcc\x98\xbc\x12\x65\x16\x46\xff\xac\xe4\x36\x73\x4c\xb5\x84\x18\x10\x25\x33\x2a\xea\x1e\x9f\xfa\xaf\x1a\x76\xf7\x86\x6e\x25\xf7\xe8\xdc\xf9\x90\x7b\xac\x5c\x3e\x96\x6d\x06\x45\xc3\xd2\xe7\xfe\x9a\x15\x68\x80\xa7\x9f\x41\x83\x75\xfc\xd7\x40\x47\x7d\xc9\x0c\x44\x94\xb6\xd3\x57\xd0\x61\x2c\xa1\x23\x8b\x42\x85\xdd\xe1\xf6\xd3\xa4\x88\x02\xd2\x84\x9c\xc6\x25\x23\x80\x1c\x17\x59\x4c\x7e\x3f\x97\x4e\xf9\x4a\xa6\x09\x93\xa2\x22\x3f\x69\x5d\xc1\x52\x07\x1e\x1e\xb2\x4a\xc6\x2c\x19\x2e\xf2\x9f\x74\x2a\xe9\x7b\xc5\x8b\xea\x69\x3e\x7d\x0f\x43\x02\xb9\x72\xea\xdc\x5e\xb9\x8d\x1a\x2c\xbc\x1e\x5e\xc7\x54\x87\xdc\xcd\x16\x2c\xef\x52\x12\x53\x61\x66\xe5\xa7\xbf\x66\x39\xb7\x54\x5e\xc5\x18\x98\x14\x1e\x03\xee\x54\xe0\x4a\xf9\x31\xa0\x6a\xea\x53\x8a\x86\x4b\x28\x66\x42\xbe\xb6\x72\x5c\x24\x9d\x49\xa6\xb2\x4b\xb3\x89\x1f\xb8\x49\xc1\xf0\x8b\xe7\x77\x65\xa0\x2c\xa4\x52\xb8\x28\x99\xc5\xf4\x3f\x93\x89\xa5\x58\x72\xe1\xbb\xe4\x62\x30\xfd\x20\xea\xc9\x91\xd7\x47\xf2\x76\x5f\x7e\x59\x8a\x30\xca\x15\x50\x98\x4a\x30\xdd\x7e\xbe\xff\x60\xc7\x7b\x5c\x7a\x8c\x9f\x94\x4e\x63\x0b\xde\x47\x4d\x01\xdd\xf4\xca\x8d\x21\xc4\x98\xdd\x7f\x0f\xbd\xdc\x18\x96\x8b\xa4\xd2\x2f\x93\x9c\x8c\x79\xe6\x2f\x60\xef\xca\x32\x0d\xa0\xbf\x20\x3a\x7a\x63\x77\x77\xdd\x8d\x57\xce\xab\xcf\x68\xe7\xa2\x9f\x2f\xa3\x37\x96\x7d\x5b\xc1\x62\x5a\x39\x45\xc5\x8f\x6e\xad\xf4\x87\x90\xdb\x6c\x39\xf5\x22\xbc\x5b\x83\x8d\x15\xbb\xfb\x2b\xfd\x19\x36\x4a\xa8\xb9\xf3\x27\x1c\xe6\xbb\xd9\xc4\x0f\xcb\x19\x90\x48\xa4\x97\x53\x4b\x9c\x0a\x34\x10\x33\x85\x98\x6a\xcc\xe0\x08\xe5\x57\x68\xe7\xc2\xef\x13\x7e\xb8\x1e\x57\x55\xda\x40\xac\xda\xdb\x4d\x54\xc2\x57\xa1\x73\xd9\x85\xef\x7e\x0a\x73\x23\xc7\x0c\xaa\xd6\xfa\x67\x25\x86\xb1\xde\x18\x3f\x07\xfc\x8b\xb8\x6a\x02\x9e\x57\xe2\xf2\x53\x17\x3e\x3c\xce\xdb\xdd\x5b\xf7\xb7\x0f\xb0\xdc\x7a\x64\xed\x73\x91\xcc\xfc\x6c\x6a\xe4\xca\xfb\x1d\xb7\x2c\x29\xc6\xdd\x32\xd7\xf2\x73\x52\x27\xce\xcb\x7d\x42\x94\x04\x9d\x3d\xa4\xbc\xe3\xc2\xf3\x8f\x53\xa5\x9b\x1a\xae\x5b\xb0\x90\x47\xd6\x11\x9b\xff\x10\xfb\xa8\xc7\xe0\x74\xa6\xdd\xa3\xd6\x5f\x38\x9d\x35\x41\x37\x14\x4d\x18\x31\x47\xb9\x23\x78\x6c\xf9\xe6\xc1\xfc\x6a\xc9\x6c\x7a\x79\x31\x91\x1c\x2f\x09\xf3\x8c\xaf\x24\xcf\x36\x6d\xe4\x36\x1b\xcc\xf3\xd0\x25\x76\xf9\xce\xee\x6c\x8e\x4c\xb5\x7b\x7f\x80\x77\x7d\xfd\x98\x8e\x27\x73\xbb\x8c\xb1\xd1\xbf\xa9\x2c\x5f\x58\x83\x1b\x9d\xd2\x5b\xf0\xef\x0d\x72\x9c\xc1\xd6\x95\xcf\x6e\x55\x02\xf2\x17\xfc\xb1\xdd\x65\xbc\x8e\x0c\xe0\x38\x8e\x6d\x16\x83\xe3\x22\xf3\x2f\x67\x33\x74\x6b\x06\x7a\xe2\x9d\xde\x5b\x25\x54\x3f\x81\x85\x0b\x4c\x07\x6f\x75\x76\xfe\xec\x6d\x58\x49\x99\x14\xe5\x51\x44\xb0\x60\xa9\x77\x52\x16\x62\x3d\x58\x9b\xf4\x63\xce\xf9\x90\x43\xf5\x23\xfc\x2f\x0a\x71\xf6\xba\xfd\xbd\xf5\xd0\x23\xa7\x06\x17\x49\xcc\x65\x39\x5e\xd4\x89\x4c\x8c\x6a\x07\xce\xd4\x5f\x5b\xa8\x78\x8e\xca\x65\xbb\x5d\x77\xf6\x57\xed\x6e\x17\x6e\xd4\x30\xfa\x5f\x0b\x7c\x3a\xf5\xf7\xb9\xef\xb9\xe9\x98\xc9\x2b\xf2\x84\x38\x39\x38\xfc\xa9\x94\x33\x7b\x6f\x60\x33\xfa\x0d\x19\x98\x0e\xa6\x82\xe9\xc5\x57\x67\x83\x8b\xcc\x65\xf0\xaa\x18\x7c\xfa\x92\x93\xdd\xd9\x5f\x15\x55\xd4\xf8\x84\x77\x79\xeb\x0a\x55\x8a\xe4\xaf\xe0\xe4\x76\x9b\xad\x7e\xde\x42\xbb\x17\xd4\x3a\xf8\x3e\x1e\x39\xeb\xf1\xf0\x1c\x9d\x43\xab\x4a\x0e\x90\x52\x70\x4c\xf5\xf3\x1b\x68\xfb\x8e\x06\xff\x63\x26\x9d\x9e\xc7\x78\x55\x51\xa4\xb1\x97\x89\x99\x10\xba\xac\x7e\x71\x5f\xe3\xcb\x9f\xf9\x76\x0e\x7b\x67\x41\x76\x69\xf1\xdf\x1c\x5e\x6a\xba\xa1\xcd\x8c\x7f\x4e\xb9\xcd\x13\xf4\xee\x06\xbd\xf3\x64\x94\x74\x03\x7f\x51\x78\x21\xfc\xd9\xfb\x7f\x00\x00\x00\xff\xff\xbd\x48\x19\x09\x17\x10\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 4119, mode: os.FileMode(420), modTime: time.Unix(1617262989, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
