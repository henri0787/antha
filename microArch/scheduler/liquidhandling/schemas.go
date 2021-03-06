// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schemas/actions.schema.json (15.986kB)
// schemas/layout.schema.json (8.11kB)

package liquidhandling

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _actionsSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x6f\x1c\xb9\x11\x3e\x4b\xbf\xa2\xd0\xbb\x08\x6c\x64\xf4\xf0\x29\x88\x6e\x06\xf6\xb2\x41\x10\x1b\xd8\x4d\x2e\x86\x33\xe0\x74\xd7\xa8\xb9\x66\x93\xbd\x24\x5b\xe3\x89\xa1\xff\x1e\xf0\xd5\x4f\xf6\x4b\x33\x12\xbc\x58\xe9\x60\xcf\x34\x8b\xc5\x62\x3d\xbe\x2a\x16\x7b\xbe\x5d\x5e\x24\x3f\xd2\x2c\xb9\x83\x24\xd7\xba\x54\x77\x37\x37\x84\xeb\x9c\x5c\xa7\xa2\xb8\x21\xa9\xa6\x82\xab\x2b\x95\xe6\x58\x90\x64\x63\x68\xfd\x67\x4f\x7f\x77\x73\xf3\x9b\x12\xdc\x53\x5c\x0b\x79\x7f\x93\x49\xb2\xd7\x57\xb7\x7f\xbb\x71\xcf\x7e\xb0\xd3\x32\x54\xa9\xa4\xa5\x61\x67\xa6\xfe\xe3\x97\x0f\xff\x82\x5f\xec\x38\xec\x85\x04\x37\xbc\xa3\xfc\x1e\xfc\x9a\x90\x12\x29\x29\x66\x20\x2a\x0d\x59\x25\xcd\x10\xa3\xbf\x57\x34\xcb\x09\xcf\x18\xe5\xf7\xc9\xe6\x12\x00\x20\xd1\xc7\x12\x0d\x4f\xb1\xfb\x0d\x53\x1d\x9e\x4a\xfc\xbd\xa2\x12\xcd\xc6\x3e\x25\x0f\x28\x95\x59\x79\x03\x89\x67\x9f\x7c\xf6\x74\xa5\x14\x25\x4a\x4d\x51\x25\x77\xf0\xcd\x3e\xb3\xcf\xc3\x94\xf6\x43\x3b\xd0\xdb\x89\xce\x11\x3c\x2d\x88\x3d\x98\xaf\x6e\xdf\x1b\xbb\xb1\x07\xc2\x68\x46\x2c\xf1\xa6\xcb\x27\x15\x5c\x69\xc3\xe1\xdd\xf5\x6d\x52\x0f\x3d\x36\x54\xb5\xa8\x4b\x44\x98\xd0\x9a\x19\xee\x6a\x0e\xcc\x96\xa3\x42\x05\x5d\x12\x29\xc9\xb1\x3f\x58\x50\xfe\xb3\xc6\xc2\x08\xf4\xae\x37\x44\xfd\xf3\xae\xa0\x76\x48\x70\xfc\xb0\x37\x56\x18\x0c\x99\xbf\x6f\x90\xfc\x28\xd1\x8c\x27\x3f\xdc\x64\xb8\xa7\x9c\xda\x8d\xdc\x68\x49\xb8\xda\xa3\x7c\x6f\x37\x96\xb4\x15\xb3\x68\x7e\x29\x45\x51\xea\xa7\xce\x3e\x10\xda\xcc\x1d\x4c\xfd\xdc\x79\xd2\x8c\xbb\x4f\x8f\xce\xdf\x6b\x66\x56\x2d\x17\x17\x89\xb3\x81\xff\x36\x88\x88\x9f\x5c\x04\x20\x10\x6f\x2c\xa0\x1c\x08\x94\x8c\x68\x84\x03\x32\x66\xc2\xe8\xe2\x62\xe8\xed\xe6\x61\xc7\xd9\x39\x29\xd0\x78\xba\x16\x9a\xb0\xed\x83\x60\x55\x81\xc6\xdd\x0d\x61\xcf\xdb\x2f\xcc\x33\x4b\x1f\xbe\x0d\xe4\xfa\x35\x47\xc8\xa8\x2a\x19\x39\x82\xa1\x0c\x4e\xee\x77\xb3\xf1\xb3\x82\x58\x4a\x1b\x9f\x4b\xec\xd3\x47\x37\xd8\x15\x64\x72\x21\x4b\x09\x8e\x12\x4a\x89\x0a\xb9\x1e\x59\x30\x6e\xb7\x02\x89\xaa\x24\x16\xc8\x75\x57\x86\x54\x14\xa5\xe0\xc8\xb5\x9a\x96\x40\x55\xbb\xab\x86\x16\x0e\x39\x4d\x73\x28\xc8\x17\x84\xaa\x04\x9d\x53\x35\xb6\xf1\x10\x31\xee\x69\x13\x0f\x17\xd1\xa5\xde\x9b\x85\xa0\x5e\x28\xcc\x8b\x9b\x77\xd4\xc2\xa9\xe0\x29\x72\x1d\x42\x19\x92\x8a\x53\xed\x4d\x3d\x62\xed\xa1\xc1\x23\xd2\xe9\x11\x9b\x77\x94\x53\xcb\x36\x62\xfd\x46\xf9\x4e\xff\x6d\x51\xe7\x16\xef\x50\x2f\x5e\x9d\x57\xc5\x0e\x65\x7b\xa4\xa0\x9c\x16\x55\x91\xdc\xc1\xed\xf5\x6d\x44\x2a\xab\xaf\x39\x61\x0c\x91\x32\xd1\xe8\x9c\x61\x28\x1f\x55\x80\x5f\x8d\xb7\x2a\xcc\x96\x68\xe5\xb2\x27\x47\x42\xb2\xcc\xfa\x2f\x61\x1f\xdb\x16\xdb\x13\xa6\xf0\xb2\x43\x3b\x4f\x6a\xb9\x3b\xf2\x39\x62\x4b\x95\xb4\x43\x66\x04\x9b\xde\x43\x8b\x68\x03\xfa\x58\xd2\x94\x30\x76\x34\x96\x21\xf0\x1f\x17\xd8\x4b\xe1\xe9\x81\xb0\x0a\x7b\xce\x1a\xc5\x25\x47\x38\x19\xad\x96\x24\xf8\x47\x7b\x23\xfd\xe8\xf4\xae\xd1\x81\x84\xae\xf1\x63\xec\x23\x96\x6f\x2d\x12\xb5\x7b\x1c\x08\x2f\xfd\x3f\xed\xec\x6e\x30\xfd\x9f\x22\x6d\x02\xa2\x93\x35\x23\x7e\xc8\x3c\xb1\xd3\xba\x99\x0e\x07\xaa\x73\x9b\x26\x7a\xf9\x5d\x8a\x9d\xd0\x63\xb9\xbd\x53\x27\xd5\xa3\x6d\x1b\x99\xe5\xd3\x2f\x5b\x03\x62\x5b\x03\x75\x90\x48\x71\x70\x80\xc3\x12\xf8\xdc\x9b\x39\x52\x41\xb5\xb6\xd2\xe2\x15\xa3\xe8\x48\xe7\xf5\x16\xcf\xd7\x31\x23\xfd\xfc\x93\x55\x08\x07\xb3\x04\xbc\x29\x25\x15\x12\xb4\xe8\xa9\xe4\x6d\x32\x60\x18\xa9\x09\xec\x3e\x67\x45\xac\x71\x66\x89\x88\xc6\x72\x52\x1c\xea\x2c\x16\x2c\x3e\x32\xbb\xa8\x98\xa6\x25\x73\x25\xd3\xbb\xeb\xdb\x31\xb2\x0e\xb0\x2d\xd9\x99\x31\xdd\xf9\x77\x96\x9a\xd8\xe7\x2f\xbb\xb9\xcb\x89\xad\x4e\x43\x5e\x64\x92\x49\x4c\x1a\xb9\xfe\x77\x99\x11\x8d\xb3\x71\x58\x59\x32\x15\x92\x80\xb6\x45\x42\x1d\x8f\xa7\x44\x1c\x13\xa9\x89\x30\x8e\x87\xad\x67\xbc\x3e\xd2\x0c\x8f\xbb\x89\xa2\xb6\x8d\x38\x51\x27\x69\xaf\x3e\xc1\xc8\x97\x40\xe7\x36\x45\x8e\xf4\x3e\xd7\xb3\x36\x30\x41\xef\x48\x83\xdf\x21\xcf\xc2\x47\x4d\x4b\x90\xc8\x88\xa6\x0f\xa6\x9c\x04\x25\x0a\x34\x18\xd9\x47\x94\x55\xb6\x91\xb8\x47\x89\x3c\xb5\x69\xab\x9b\xbf\x56\xdb\xa8\xe1\x35\x1a\x8d\xb1\x0a\x9c\xe8\xaa\xb0\x07\x4a\x5d\xef\x7e\x2c\xca\x90\xdb\xd8\xf9\xe4\x52\xcc\x76\x27\xb4\x16\x85\x11\xd8\x7e\xd5\xa2\x34\x9f\x9d\x09\xb7\x0c\x1f\xd0\x40\x7a\x0c\x42\x9e\x94\x84\x6b\xd1\x9e\x29\xff\x7a\xc3\xaf\x48\xbd\x67\x75\x51\x45\xf9\x3d\xc3\x5f\xfd\xd9\x74\xd6\x55\x5b\x87\x3a\x37\x13\xd2\x9c\x70\x8e\x0c\xc2\xf1\xf6\x14\xb7\xdc\x4b\x67\x57\x2d\xac\x5f\xfa\x32\x0c\x92\x03\x51\x1a\x6d\xda\x2e\x05\xa3\xe9\xd1\xb6\x3d\x94\x35\x7b\xe6\xfe\xcb\x91\x64\xeb\x5d\xd7\xae\xb7\xd4\x6b\x6d\xc5\x2e\x2a\x99\x76\x0b\x17\x1b\xa3\xf1\xbd\xd7\x9c\xe2\xa0\xd3\x05\xea\x45\x39\x4f\x8b\x55\xe2\x66\xa8\x34\xe5\x56\xd4\x37\xea\x6d\x5f\xda\x0d\x84\xf4\x05\xb4\x28\x19\x45\xe5\x1e\x5c\x99\x93\x12\x72\x85\x63\xdb\x99\xea\xa9\xd4\x44\xf5\x81\x71\xd9\xf6\xc7\x1a\x1a\x9d\x1e\xcd\x22\x1d\xb5\x8e\xe5\x4b\xf5\xe4\xcf\xe7\x62\x1f\x3a\x15\x41\x45\x12\x6d\xd7\xc2\x1e\x93\xc7\x5a\x4c\x35\xe3\x05\x07\xf8\x25\xf2\x7b\x67\x3f\x49\x7e\xc7\xe3\xc5\x45\xf7\xd1\xb9\x46\xf4\x61\x07\xc6\x9c\xc6\xd0\x67\x06\xaa\x66\x63\xab\x8f\x91\x4b\xe4\x34\xd8\xb1\x54\xc8\x8f\x44\x92\x02\x35\x4a\xe5\xd2\x30\x66\x26\x0b\x1b\x71\x0f\xe4\xd8\x40\x39\x51\x25\xf5\xa7\xe7\x03\x51\x50\xa2\xdc\x0b\x59\x58\x2c\x5f\xa1\xf1\x92\x96\xa8\x35\xe5\xf7\x1f\x4a\xd7\x69\x5b\xb4\x9d\xec\xfc\xdb\xf1\x10\x60\x8e\x5e\x8b\xb6\x43\x18\x9b\xe8\x89\xc2\x64\x5f\xb3\xbf\xe7\x31\x28\xb0\x5c\x46\x47\x60\x36\xd9\x0c\xa8\xbb\xc9\x47\x8b\x2a\xcd\xc5\x7e\x3f\x48\x23\x83\x79\x65\xad\xc3\x51\xb5\x77\xe8\x77\x4c\x1c\x44\x35\xac\x03\x47\x27\xf4\x6c\x56\x46\x6d\x46\xc0\xf3\x9d\x0d\xef\x01\xff\x55\x6a\xaa\x67\x75\xd5\xd5\x64\xe7\x50\x21\x41\xb2\x67\xe2\xb0\x95\x16\xce\x67\x74\x58\x33\x9d\x49\xd1\xa3\xf3\x66\x50\x7e\x74\xde\x38\x7a\x56\x0a\x33\x0b\x3b\xb5\x62\x37\x40\xf7\xb3\xae\x3f\xba\xd2\x13\x30\x75\xec\x6f\x22\x20\x06\xcb\x8e\x9c\x39\x66\xe7\x45\x14\xe3\x4b\x53\x53\x25\x85\x82\x55\x04\x7d\x34\x5a\x3a\x8f\x5e\xbc\xd8\xcf\xa2\x92\xc6\x2d\xcf\xa1\x15\xc3\x0d\x0c\xb7\xfa\xf4\xf2\xd4\x38\xac\x57\x39\xa7\xab\x2c\xa2\x9c\xa7\x5a\xa0\xdf\x06\x31\x9f\x0a\x6c\x87\x1c\x75\x8e\x12\x84\x04\x2e\x34\x10\x08\x1c\xcd\x71\x68\x6d\xe0\xd5\xa8\xb6\x13\x82\x21\xe1\xf3\x3a\x9b\xd6\xc2\xf8\x68\x7c\x24\x7a\xe4\x1c\x08\x69\x0f\x2a\x6b\x0a\x24\x33\xc1\x76\xff\x84\xf9\x9f\x96\xca\x7c\x11\x12\x2a\xde\x3c\x71\xe7\xa6\x53\x3a\x61\xdf\x63\x23\x6b\x50\x18\xcc\x1d\x4e\xfd\x9d\x3b\x9a\x82\xd7\x5d\x1e\x93\x9d\x89\xcb\x5c\x1c\xa0\x66\x66\xeb\x99\xd6\xa5\xf2\x29\xe7\xd5\x58\xe6\xeb\xb5\x22\xb6\x7b\xc1\x98\x38\xac\x3f\x9b\x36\x38\x3e\x03\x99\x51\x37\xeb\x40\xde\x12\x80\x89\x72\x89\x6d\x64\xa9\xef\xd2\x3d\x68\x59\xe1\x06\xdc\xbc\x76\x81\x6f\xf9\xb5\x2f\xf3\x67\xcf\x27\xb3\xa1\x1d\x93\xbe\xa0\x5f\xcd\x89\x60\xa9\xc0\xaa\x2a\x0a\x22\xe9\xff\xd0\x8a\xd4\x98\xc7\x37\xe3\x9d\x4b\x11\x06\x8e\xed\x72\x99\x27\x8b\xac\x9e\x43\xa5\xc7\x94\xa1\xea\x36\x3f\x4e\x70\xb2\x7a\x95\x85\x55\x56\x58\x7f\x0e\xcf\x63\xad\x2d\x87\x2f\x46\x59\x05\xfd\x0a\x8e\x11\xe8\x9c\x68\x38\xa0\xc4\x89\x88\x1b\x55\xdc\x24\x62\xd5\xd4\xcb\x90\xab\x21\x9f\x44\xb0\xf0\x37\x91\xfd\x5a\x95\xe7\x53\xe3\xaa\x66\x75\x52\x8c\xd7\x5c\xce\x12\xeb\x35\xb7\x35\x31\x5f\x4f\x7a\xc6\xd8\xaf\xd7\x58\x9c\xde\xe3\xe9\x79\xac\xbf\xb4\x28\x2d\xc5\x59\x3f\xc6\xd3\x16\x91\x84\x31\x64\x8b\x7b\xaa\x12\xfd\x8b\x21\xca\xde\x7b\x2a\x77\x0b\xe0\x67\xfb\x18\x22\xdd\x10\x02\x45\x8d\xdf\x13\x8e\xa2\x52\xec\x08\xbb\x23\x10\x50\x68\x67\xfa\x86\xac\x3a\x25\xb1\x35\x3c\x20\xf9\x42\x79\x96\xc0\x06\x12\x4d\x0b\xdc\xa2\xd2\xb4\xf0\x10\x94\x56\x45\xe5\x2e\x26\xb6\xdd\xb1\xb5\xb9\xce\x2e\x61\x9d\xb7\x7e\x87\x2c\x68\x71\x5b\xf7\x7f\xe2\x97\x80\x41\xd0\x35\x25\x95\xed\x5b\xb6\x55\xb5\x01\xca\x33\xfc\x8a\x99\x51\x64\x68\x68\x4f\x17\x4c\xcb\xd0\xbd\x24\x5a\xa3\xe4\x1f\x17\xc2\xef\x7f\x3f\xdd\x5e\xfd\xfd\xf3\x5f\xa7\xe2\xb8\xd7\xaf\x7f\x56\x67\x0f\x6f\x3b\x75\x6c\x0b\xe1\x82\x23\xa6\xdd\x40\x65\x3d\x98\x16\x08\x9a\x7c\x41\xde\x34\xf3\x28\x57\x5a\x56\xf6\x35\x34\xa3\x73\x50\x98\x0a\x9e\xf5\x5d\x75\xa0\xe4\x99\xda\x75\xf0\x4e\x4c\xfd\x8e\xd4\x98\x83\x2e\xdf\x84\x7d\x7f\x6b\xb0\x15\x04\x59\x71\x50\x02\xf6\x44\x02\xd9\x6b\x1c\xee\x0f\x72\x53\x67\x8a\xa2\x64\xa8\x31\x7b\xd6\xdd\x9e\xb7\xec\xd6\xb4\xf4\x2f\x0a\x2e\xad\xb7\xa1\x79\x5f\xd3\xf6\xa0\x05\xc9\x6c\x85\x14\x0e\x29\x16\xea\x69\x79\x5e\x44\x0a\x97\x3e\x2f\x02\x4c\x71\x53\x34\xf7\x92\x66\x97\xee\x06\xd5\x7e\x7a\x3d\x0b\x9e\x0f\xac\x4b\xa1\x68\xb8\xc0\xea\x6c\xdd\xb6\xa5\x84\x84\x4c\x8a\xb2\x56\x86\x09\x4f\x24\x69\x1e\x50\xfc\xfa\x7b\xc3\xef\xde\x2b\x0b\xaf\xe8\xfd\x8a\xde\x1d\xab\xc5\x21\xb9\xfb\xea\xf8\x62\x5c\x26\x1c\x90\x6b\x2a\x9b\x3b\x5f\x0f\xd3\x1b\xdf\xcf\xa5\xed\x92\xd3\x84\xd6\x26\x60\xc9\x5f\xa0\x5f\xc7\x36\x95\x7a\xab\x26\x65\xe2\xde\xbf\x39\x79\x2f\x45\x55\x0e\x9a\x75\xab\x30\x3e\xe0\x7a\x9a\x53\x96\x49\xe4\x2f\x83\xed\xad\xa2\x73\xae\xd6\xf4\x62\xad\x81\x2f\x3b\x69\xfa\xc7\x0c\x2b\x2e\x39\x97\xdd\xb8\x8f\x03\xd4\xf4\x4f\x17\xc2\xdf\xf8\x4f\x18\xea\xd2\x60\xae\x3f\x3c\x7e\xdb\xd7\x3f\x1d\x4d\xf4\x59\x87\x59\x14\x26\xd0\x71\xee\x45\x81\x57\x34\xfc\xe3\xa1\xe1\xc9\x2d\xe4\xf6\x6f\x66\x96\xb7\x8f\x09\x07\xfb\x0b\x30\x89\xc8\xc1\xf1\xf0\x80\x79\x20\x54\x2b\xab\xb8\x4a\xa1\x04\xca\xcb\xd3\x3a\xc8\x01\xf2\x0a\x54\x8a\xdc\xe3\x3a\xc4\x5b\xdd\x53\x8e\x9d\xb3\xed\xee\xe2\x80\x17\x84\x5a\x83\x77\x7e\x8e\xcb\x10\x54\x81\xca\xc5\x81\x87\x57\x0c\x8c\xce\x4e\x78\x8f\xe3\x35\x82\xff\x7c\x11\xdc\xfa\xdd\xda\xf7\x1c\xbf\x59\xe5\x4a\xa3\x6d\xad\xd5\x17\x2e\x5d\xcc\xc6\xe2\x51\x3c\x10\x6d\x69\x38\xe7\xe2\x00\x4c\x98\x12\x45\x58\xbd\x19\xb5\x3d\x97\xe7\xf4\x87\x5f\xa3\xfd\x0f\x12\xed\x67\x4a\x1b\x6f\xc2\xf5\xda\xdb\xf9\x0c\x62\x42\x99\xa1\x75\xc9\x89\x1f\xd6\xcc\xe6\x93\x33\xe0\xd4\xe5\xc5\x23\x5c\x3e\x5e\xfe\x3f\x00\x00\xff\xff\x04\x97\xee\x91\x72\x3e\x00\x00")

func actionsSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_actionsSchemaJson,
		"actions.schema.json",
	)
}

func actionsSchemaJson() (*asset, error) {
	bytes, err := actionsSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "actions.schema.json", size: 15986, mode: os.FileMode(0644), modTime: time.Unix(1554136481, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x87, 0xea, 0xb8, 0x3a, 0x6c, 0xd8, 0xdc, 0x6a, 0x56, 0x8f, 0x31, 0xcf, 0xe5, 0x3a, 0x48, 0x9c, 0x5d, 0x12, 0x2c, 0x36, 0xb4, 0x6b, 0x38, 0x89, 0x4c, 0x5b, 0x6f, 0x15, 0xc4, 0xb7, 0x77, 0x45}}
	return a, nil
}

var _layoutSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\x5b\x6f\xdb\xb6\x17\x7f\x96\x3f\x05\xa1\x16\xc8\xc3\xdf\x89\xd3\x7f\x1f\x86\xe5\xad\x40\x5f\xba\x0d\x6b\xb1\x0e\xdb\x43\x90\x05\xb4\x74\x14\xb3\x95\x48\x95\xa4\xe2\xb8\x85\xbf\xfb\x70\x48\x4a\x22\x25\x4a\xbe\xb4\x1d\x30\x2c\xc0\x66\x59\x3e\x97\xdf\xb9\xf2\x1c\xf6\xcb\x22\x49\x9f\xb3\x3c\xbd\x21\xe9\x46\xeb\x5a\xdd\xac\x56\x94\xeb\x0d\xbd\xca\x44\xb5\x2a\xe9\x4e\x34\xfa\x52\x65\x1b\xa8\x68\xba\x44\x52\xf7\xec\xc8\x6f\x56\xab\x0f\x4a\x70\x47\x71\x25\xe4\xc3\x2a\x97\xb4\xd0\x97\xd7\x3f\xac\xec\xbb\x67\x86\x2d\x07\x95\x49\x56\x6b\x26\x38\xb2\xfe\xf4\xfe\xed\xaf\xe4\xbd\xf9\x9d\x14\x42\x92\x1c\xb2\x8f\xc4\x2a\x23\x3e\x29\xb2\xea\x5d\x0d\xc8\x23\xd6\x1f\x20\xd3\xe6\x95\x84\x4f\x0d\x93\x80\xa0\x6f\xd3\x35\x14\x42\x42\xba\x24\x29\x2d\x34\x48\x7c\xe0\xb0\xbd\x67\xb9\xc2\xc7\x47\x90\x0a\x25\xdd\x21\x5f\x2d\x45\x0d\x52\x33\x50\xe9\x0d\xf9\xb2\x20\xee\xaf\x23\xf2\x5f\x9a\x1f\x06\xb0\xf5\x06\x88\xa3\x25\xa2\x20\x7a\xc3\x14\x51\xbd\x15\x8f\xb4\x64\x39\x75\xc0\x03\x39\x99\xe0\x4a\xa3\x84\x17\x57\xd7\x69\xf7\xd3\x7e\xb9\x48\x92\x16\x3f\xea\x4e\x92\x24\x7d\x2e\xa1\x40\xca\x67\xab\x1c\x0a\xc6\x19\x8a\x53\x2b\x74\xd0\xfb\xa6\xaa\xa8\xdc\xa5\x8b\x24\xb1\x9c\xd6\xde\x33\x18\x5b\xff\xb4\xac\x63\x17\x27\xa3\x90\x55\xb4\x56\xa4\x90\xa2\x22\xe8\x05\x4b\x49\xde\xbc\x56\x84\x71\x72\x61\x6d\xb8\x20\x5a\x90\x0b\x83\xea\xc2\x09\xa1\x79\x6e\x80\xd0\xf2\x5d\xe8\xfb\xc4\xd7\xab\xb4\x64\xfc\xc1\xb2\x8c\x14\xbf\xe2\xe4\xcd\x6b\xa3\xc5\x49\x46\xaa\xfd\xc2\xfc\xb7\xb7\xb9\xd5\x99\xeb\x44\xa7\xbe\xd5\xa7\x18\xf9\xda\x7c\x5d\x83\x32\x46\xba\x84\x14\x05\xa1\xa4\x64\x9f\x1a\x96\x6f\x28\xcf\x4b\xc6\x1f\x6c\xc2\x52\x4d\x28\x79\x60\x8f\xc0\x49\x2d\x18\xd7\x08\x52\xb3\x0a\x9c\xec\x20\x4b\x6b\xa1\x1c\xc4\x3b\xfb\x6b\x1d\xf1\x47\x4f\xd4\xbe\x8a\xc3\x1e\x03\xff\x19\x76\x8a\x50\x09\x06\x37\xa7\x15\x28\x84\xdd\xc9\x23\x82\x9b\x5f\x0c\x6c\x93\xb9\x10\xb3\xa8\x93\x3e\x1b\xb6\xd9\x5c\x7b\xe7\x54\xa6\x96\x74\xbf\xe8\xfe\xbf\xef\xd2\x2f\xa0\x3b\x25\x3c\xb4\xb3\xa8\x35\x68\xd2\x88\xb8\xf7\xb1\x23\x28\xf6\x19\x8e\x09\x82\x17\x83\xb8\xb9\x99\x10\x32\x67\x9c\x6a\x50\x2f\xad\xb9\x7b\x97\xc1\x46\xc5\x09\xec\xff\x0f\xd9\x99\x86\xea\x70\x06\x08\x0e\x6f\x51\xec\x6d\x1b\x95\x29\x45\x75\x49\x35\xb8\x78\xb4\x3a\x92\xe4\x10\x97\x66\xf5\x5a\x3c\x9d\xc3\xb6\xa5\xca\xd3\x67\x3f\xef\xbc\x3c\x98\xed\x0c\x05\x2d\x15\x74\x99\x12\xf8\xe8\x94\x4c\xf9\x7d\x2b\x48\xce\x2a\xe0\xca\xa8\x20\x9e\xa0\x58\x7e\x3c\xdd\x57\x15\xe6\xc6\x0e\x3f\x67\x72\xc3\xd0\x8d\x23\xc3\x9b\x6a\x8d\x27\x4f\xbc\x36\x9f\x3c\xed\xd8\x22\xaa\x2a\x0c\xf7\xee\x1c\xa1\xbb\x29\xa1\x5f\xe1\xe2\x97\xa7\xb9\x78\x23\x01\xfe\xa3\x4e\x76\x42\x3f\x9f\x23\xf4\xf3\xb7\x8a\x9c\xdf\x26\xe2\x11\x0b\xdc\xcf\x72\x33\x1a\xd1\xca\xcc\x4a\x86\x7e\x89\xc7\x3a\x6f\x0a\x9a\xe9\x46\xda\xd1\x29\x98\xbe\x48\x2a\xc5\xd6\x8c\x51\x99\x28\x9b\x8a\x9b\xc7\x2e\xe2\xe6\xdb\x16\xca\xf2\x3e\xf2\x4a\x69\x2a\x75\xf7\xcd\x68\x9b\x09\xb9\x99\x3f\x87\x6e\x0c\xa6\x82\x48\xfa\x01\x69\x38\xfb\xd4\x00\x8e\x07\xee\x4c\x73\xd6\x07\x41\x32\x16\x9f\x23\x1d\x19\xbb\x31\xaf\xa6\x52\xb3\xac\x29\xa9\x8c\x2a\x71\x52\x4f\x57\x82\x94\xa1\x26\x2b\x7e\x49\xe0\xea\xe1\x8a\x5c\xd4\x99\x34\xfd\xfb\x22\xd4\x17\x84\xed\xeb\x8c\x03\xe2\x0b\x1b\xc0\x08\xb5\x86\x72\x4e\x53\xfa\x8a\xe0\x29\xbd\xf3\xc7\xfb\x59\x5d\x26\xf3\x4e\xad\x2d\x63\x99\xa1\x40\xd9\x28\x02\x3f\x31\x05\x15\x11\x92\x68\x56\x9b\x79\xd5\x57\xda\x4a\xaa\x9a\x52\xb3\xba\xb4\x67\xea\x8b\xab\xeb\xee\x3d\xe3\xac\x6a\xb0\xd0\xae\xaf\xae\x03\x84\x6d\x4d\x7c\x1d\x48\x27\xe5\x3b\xe2\xf4\xaa\xb3\x87\x1a\xc1\x84\x53\xcb\x30\x24\xcb\x33\x27\xa0\x61\x5b\x38\x52\x31\x1c\xe7\x83\xb3\xd0\xd8\x8e\x34\x0b\xa4\x9f\x2e\x2d\x98\x0c\xb8\x6e\x4b\x02\x48\x53\xd7\x20\x49\x09\x85\xbe\xac\x84\xd2\x06\xaa\x43\xfa\x4d\x81\x0e\x7a\xc9\x44\xd3\x58\xa2\x6a\x49\xf5\x06\xa4\x41\xa7\x36\xd4\xbc\xf4\xdd\xd8\xa2\xb2\x13\xe0\x72\xbe\x5e\x81\x9b\xec\xb9\x4d\xb3\x5d\xc9\x78\x6e\x8f\x83\xac\x59\x0b\x96\xa7\x77\xb1\x31\xde\x4a\x6d\x0f\xa0\x01\xcc\xdf\xa0\x96\xa0\x80\xe3\x8a\x64\x08\xfd\x15\xa4\x5d\x0e\xcb\xd2\x4e\xb0\xc1\xb2\x8c\x7f\x5f\xc8\x84\xeb\xcc\x99\x47\x9c\xbb\x5a\x0f\x45\x4e\x94\x24\x49\x3f\x32\xee\x9d\x2a\x63\x84\x8d\x82\x1c\x17\xd6\x47\x90\xac\xd8\x19\x6c\xf6\x50\xec\x18\xba\xad\x7d\x38\x41\x0f\xe1\x76\x7b\xbe\x04\xc5\xf2\x86\x96\xf7\x8f\x58\xd5\x30\xba\x50\x18\x31\x0c\x30\x59\x36\xb2\xdd\xb0\x6c\x43\x32\xca\xb9\xd0\x64\x0d\x44\x42\x25\x1e\x21\xef\xd7\x6f\x57\x25\x45\xeb\xdc\x74\x1a\x93\x51\x13\x77\x66\x05\x54\x35\x12\x2a\xe0\x3a\x9d\xe4\xef\x26\x7f\x74\x87\x06\xae\x55\xe0\xd4\x9a\x6a\x0d\x92\x47\xd7\xc4\x24\x49\xff\xba\xbd\xbe\xfc\xf1\xee\x7f\xcf\x83\xb7\x91\xbd\xce\x76\xc1\xae\x09\xf6\x41\x98\xdc\x7f\x8e\xd0\x3e\xa9\x7f\x8c\xc0\x54\xbb\xb3\xcf\xaf\x21\xdc\xf1\xf1\xb9\xdd\xf2\x5d\x7f\x60\x36\x97\x2d\x6a\x1f\xd1\x64\xcd\xdb\x45\x35\xf5\x28\xf7\xfd\xf3\xde\xb7\xe9\xc0\xec\x37\x60\xee\x59\x8f\x61\x74\x6c\x2d\x53\x38\x20\x9a\x82\x59\x8e\x93\x98\x78\xd5\x7f\xd7\x55\xbf\xdb\x0f\x8f\x28\x7f\x4b\x39\x57\xff\xf3\x8b\xb2\xa9\x79\xbf\x4d\xfe\x63\x75\x3f\xb1\x03\xa7\x15\x53\x8a\xf1\x87\x7b\x3c\xa3\xe6\x54\xfd\xc2\x94\xb6\x57\x49\xfd\x4d\x0c\xd5\xae\xbc\xcd\x01\x47\x25\x10\x2c\x72\xe7\x2d\x0f\x44\x9b\xf5\x54\x4a\xba\xf3\xde\xa3\x37\x06\x45\x36\xd0\xfa\x27\xe6\xad\x77\xc6\x74\x63\x9e\x45\x8d\x8a\xbd\x94\x9d\x2e\xaf\x30\x3b\xa4\xd8\xba\x45\xc0\x0d\xf2\x33\x31\x70\xf3\xdb\xb0\x18\xe3\xe3\x91\x73\x69\x7c\xac\x69\xfd\x3d\x9c\x6d\xc6\x65\x83\xc0\xbe\xb7\xbe\x45\x44\xf3\x51\xf5\xba\x3f\xba\xfe\xa6\x6a\xcd\x5e\xaa\x1c\x57\x6d\x86\xf6\x5f\x5a\x6f\xe1\xe5\x51\x57\x71\x39\x53\x19\x95\x39\xe4\xe3\x9a\x9b\x8a\xf2\xfc\xd8\xdd\x09\xf4\x06\x4d\x20\x56\x7b\x2f\x62\x32\x4b\xa2\x29\x72\x66\x74\xdd\xd1\x30\x11\xdb\xf6\x4a\xba\xbb\x83\x46\xac\xed\x4c\x85\x47\x94\x0b\xee\xe1\x0b\x80\x6e\xf5\x17\xba\x6f\xee\x33\x4b\xf9\x60\x71\x8e\xb8\x33\x67\xaa\x2e\xe9\x2e\x58\x26\x9d\x35\x13\xc3\x66\x30\xed\x06\x40\xe6\x07\x5e\xa4\x24\x6e\x3a\x6a\x73\x3d\xae\xf0\x88\x51\x27\x09\x96\xb8\xaa\x16\x3c\x1c\x6b\xa2\x3b\x4a\xb3\xbe\xec\x69\x5d\x0b\xaf\xe8\x47\x5c\x0b\xec\x90\x3d\x61\x78\xd0\xbf\x47\xdd\x7b\xbc\x25\xab\x66\x4d\x3a\x45\x5d\x2a\x4e\x75\xe9\x78\x84\x33\xc1\x71\x75\x91\xb4\xbd\xc6\x69\x38\xd3\x7d\xdb\x9e\x68\xda\x83\x80\x4f\x4c\x49\xb1\x98\x07\xce\x89\x9d\x2d\x7e\xf4\xc3\xa1\x25\x84\x7a\x48\x79\x40\x7d\xb4\xf6\x51\xff\x8f\x37\x78\x0f\x95\xf1\xd7\x21\x30\x48\x64\x3a\x87\x3b\xcf\x47\xf8\x98\x22\xf0\x84\xd9\xaa\x20\x3f\xc6\x2b\xa3\xae\x77\xf8\x68\xd9\x1f\xf8\x17\x9b\x9e\xf4\xe4\xeb\x45\xbf\x64\x26\x7a\xd3\x2b\xe2\x11\x2d\xb1\xad\xb3\x8c\x96\xe5\xce\x6e\x26\x7f\xd8\xc2\x3e\xb6\x3d\x3d\xd2\xb2\x81\x41\xb2\x46\xfb\x92\x25\x9c\xad\x56\x43\xd2\xcd\x3d\x9e\x21\xc3\xea\x74\xa9\x11\xb4\x84\x30\xf8\x13\xf7\x8e\xc3\xc8\x7b\x4a\xa2\x71\x8f\x37\xc2\x13\x62\xb2\x48\xf6\x8b\xfd\xe2\xef\x00\x00\x00\xff\xff\x94\x3f\x47\xf0\xae\x1f\x00\x00")

func layoutSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_layoutSchemaJson,
		"layout.schema.json",
	)
}

func layoutSchemaJson() (*asset, error) {
	bytes, err := layoutSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.schema.json", size: 8110, mode: os.FileMode(0644), modTime: time.Unix(1553246720, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1d, 0x9e, 0x59, 0x28, 0x2a, 0x44, 0x5c, 0xb1, 0xc3, 0xd6, 0x5a, 0xf2, 0xeb, 0x2b, 0x69, 0xed, 0xf5, 0x8a, 0x34, 0x89, 0x50, 0xdc, 0xfe, 0xa, 0x6e, 0xf5, 0x86, 0x6e, 0x23, 0x1e, 0x87, 0xf8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"actions.schema.json": actionsSchemaJson,

	"layout.schema.json": layoutSchemaJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"actions.schema.json": &bintree{actionsSchemaJson, map[string]*bintree{}},
	"layout.schema.json":  &bintree{layoutSchemaJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
