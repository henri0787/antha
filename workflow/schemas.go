// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schemas/workflow.schema.json (12.184kB)

package workflow

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

var _workflowSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5b\x73\x9b\xb8\x17\x7f\xcf\xa7\xd0\xd0\xcc\xf4\xf2\xc7\x76\xe7\xff\xb2\xb3\x79\x6b\x1d\x6f\x9b\x9d\x4d\xeb\x8d\x3d\xbd\x4e\xda\x91\xe1\xd8\xa8\x01\x89\x0a\x91\x86\x76\xf3\xdd\x77\xc0\x20\xb0\x2d\x81\x94\xe0\x26\x5b\x3f\x25\xe8\x1c\x9d\xfb\xd1\x8f\x23\x7e\x1c\x20\x84\x90\x73\x48\x7c\xe7\x08\x39\x81\x10\x71\x72\x34\x1a\x61\x2a\x02\x3c\xf4\x58\x34\xfa\xc6\xf8\xc5\x32\x64\xdf\x92\x41\xe2\x05\x10\x61\xc7\x2d\x19\xca\x7f\x4b\xa6\xa3\xd1\xe8\x4b\xc2\x68\x49\x34\x64\x7c\x35\xf2\x39\x5e\x8a\xc1\xd3\xdf\x46\xeb\x67\x0f\x2a\x4e\x1f\x12\x8f\x93\x58\x10\x46\x73\xee\x3f\x67\xaf\x5f\xa1\x59\x41\x82\x96\x8c\xa3\xf5\xf2\x82\xd0\x15\x92\xb2\x2b\x56\x91\xc5\x90\xf3\xb0\xc5\x17\xf0\x44\xf5\x94\xc3\xd7\x94\x70\xc8\x0d\xf8\x58\x3c\x29\x9e\xae\xb7\x7c\x03\x3c\xc9\x05\x15\xcf\xcf\x4b\x06\xec\xfb\x24\x17\x8f\xc3\x29\x67\x31\x70\x41\x20\x71\x8e\xd0\x12\x87\x09\x94\x24\x71\x73\xe1\x47\xbd\xeb\xdb\x52\xa5\x13\x7f\xe3\xf9\x86\x76\x89\xe0\x84\xae\x4a\xed\xe4\x6a\x8c\x85\x00\x5e\x98\xfc\xe9\x23\x1e\x7c\x7f\x36\xf8\xf0\x74\xf0\xfb\x67\x34\x38\x7f\x72\xe8\x48\xd2\xeb\x9a\xcb\x39\x05\x81\xb5\x52\x3e\x4a\x27\x20\x87\xa6\x61\xe8\x9c\x6f\xcb\x53\x5b\x20\xd7\x5f\xe1\x08\x94\x2b\x2a\x5b\x76\x88\xae\x0f\xd4\xff\x35\xf5\x3f\xa1\x97\x40\x05\xe3\xd9\xfe\x8c\x98\x86\x58\xc0\x3c\x8b\x35\xeb\x48\x9b\x34\x3b\x54\x65\x78\xa6\xed\x02\x25\xf9\xf0\x49\xeb\x7a\x41\x73\xc8\x61\x99\x0b\x7e\x30\xf2\x61\x49\x68\x91\x72\xc9\x28\xae\x54\xde\x75\x6a\xf5\xbb\x56\xae\xec\x3e\x35\x09\xc2\x66\x19\x58\xe6\x2c\xd0\x34\x2a\xc2\xf4\xff\xe1\x53\xe7\x5c\xb9\xff\x19\xc4\x2c\x21\x82\x71\x95\xcb\xda\x7d\x5f\x45\x37\xcb\x53\x51\x13\xe0\x66\xd5\x3c\xe2\xb9\xac\xa1\xec\x4d\xff\xac\x88\x08\xd2\x45\xfe\xe7\xe3\xd1\xa6\x37\xaf\xd5\xc5\xd7\x11\xdd\xb6\xa8\x6a\xa2\xc9\x2b\xf3\xb3\x9b\xd5\xc8\x24\x84\x08\xa8\xb8\xa9\xeb\xb4\x96\x18\xd6\x04\xe6\x1c\x67\xba\x92\x20\x02\xa2\x8e\x32\x50\x3b\x05\xd6\x46\xe9\x93\x5c\x91\xca\xbb\x3a\x38\x27\x34\x11\x98\x7a\x3d\xd4\x76\x67\x9e\xd5\x06\x79\x2c\xca\x75\xcf\x37\x9d\x07\x24\x41\x24\x41\x22\x00\x44\x71\x04\x88\x2d\x11\xa6\xa8\xb4\x0e\x91\x52\x3d\x17\xc1\x70\x35\x44\x0f\xa3\xec\x33\x0e\xc9\xd7\x94\x89\x87\x1a\x45\x50\xfb\x39\xf0\xbf\x43\x8d\xb7\xee\xba\x65\x95\x16\x57\xf1\xe8\xa1\x71\xb5\x45\x7b\xcc\x28\x05\xaf\x90\x7c\xc7\x09\x5c\xa9\x54\x6b\x64\x9c\xce\xad\xdd\xa8\x1b\x7f\x48\xd2\x26\xb6\x51\x37\xe0\x31\xa3\x4b\xb2\xda\x43\xeb\xdd\x2c\x85\x46\x09\x20\x1f\x2e\x89\x07\xc8\x0b\x71\x92\x28\x7c\x2f\x8f\x0d\x75\x54\xc6\xd9\x73\xc2\x74\x21\x7b\x41\xc2\x84\xd1\x29\x89\x41\x9c\xe2\x2b\x2d\x55\xc8\x16\x38\x3c\x25\x57\xc0\x75\x24\x2f\x71\x44\x42\xc1\xa8\x6e\xfd\x2f\xbc\xf0\x32\x01\xba\xe5\x02\x58\x9c\x01\xf6\xf5\x12\xfe\x9e\x8e\xcf\x74\x6b\xb3\x00\x5f\x00\x3f\xa1\x5e\xba\xc0\x82\x69\xb7\x98\x83\x87\x15\x39\x75\xde\xc7\x59\xf6\xa9\xe1\xa4\x43\xe3\x0e\xda\x4f\x26\xab\xb2\x49\x59\x68\xeb\x4c\x1a\x17\x89\xd4\x2a\xa8\xbb\x0e\xe6\x90\x88\x1c\xc2\xf4\x81\x35\x91\xd2\x84\x53\x72\x35\xc7\xc9\xc5\x38\x00\xef\x42\x43\x53\x0b\x44\x66\xcd\x49\xbb\x4b\x8b\xd3\xa2\x86\x1e\x37\x6c\x48\x6a\xa8\x48\xa2\x34\xc4\xa2\x0d\x27\xde\x1a\xb0\xd7\x32\x14\x6f\x50\xdb\xd2\xd4\xa8\xb4\x96\x65\xf4\x46\xa5\xb0\x54\xee\xd0\x80\xc5\x3b\x42\x95\x0c\x33\x81\xb9\x30\x27\x9f\x50\xdf\x9c\xf8\x84\x1e\x13\x6e\x4e\xfe\x3a\x15\x56\xf4\x13\xce\x19\xef\xed\x44\x55\xc8\x34\xf2\xb8\x16\xef\xee\xe8\x61\x04\xe9\x3a\x41\x4f\x3b\x0e\xb6\x13\x5a\x0b\xb7\x02\x5e\x92\xcd\x00\x80\x49\x5a\x75\xe5\x27\xeb\xea\x01\x7f\xd2\x85\xb0\x9b\x3f\x35\x2a\xeb\x5e\xd5\xe0\x4e\x64\x86\xcd\x25\xa9\xa5\x6f\x8d\xb1\xba\x64\x69\x02\x95\xc9\x16\x38\x47\xc4\x57\x41\x94\x5d\xb1\xca\x5e\xa2\x85\xe4\xd5\xaf\xc5\x45\xe8\x5e\x65\x4a\x37\x76\x97\x26\xdd\x34\x5b\xac\xde\x5d\xda\x20\xf0\xcf\x41\xd3\x9b\x47\x91\x2b\xcf\x82\x06\xba\x38\x68\x88\x73\x1a\xde\xdd\x1c\x09\x36\x46\x00\x3d\xbf\xc4\x1f\x13\x0e\x9e\x72\xe3\x1d\x01\xda\x39\x9d\xa2\x07\x3f\xe7\x98\x7a\x41\xbf\x7b\x8e\x59\x14\x11\xd1\xdb\x3c\xb1\xa7\x10\xd7\x0e\x54\x63\xc6\xe6\xa0\xa2\xe7\xd8\xc9\xb9\x58\x76\xcb\x41\xab\xfe\x10\x9d\x62\x71\x9b\x28\xee\xc5\xe3\x5b\x66\xbb\x9b\xca\xb6\x46\x41\xb6\xa8\x3d\x8c\xc2\xfa\x8f\x81\x72\x38\xbf\xbd\x63\x17\x5c\x36\xb5\x41\xd2\x5d\x99\x1f\xb7\x34\x8d\x16\xc0\x5b\x66\x35\x2d\xa7\xbb\xbe\xe1\xd8\x8b\xe9\xe3\x58\x10\x3c\x05\x0d\x87\xfa\xc5\xb0\x45\x8e\x33\xc5\x1c\x47\x20\xc0\x04\x0c\xf7\xfa\x56\xdc\x59\x3c\x32\x53\x8d\xea\xa4\x31\x95\xea\xb9\x62\x66\x2c\xe5\x8a\x3a\x94\xeb\x6a\xc8\xe1\x49\x7d\xe6\xc0\x23\x42\x71\x68\x56\x49\x73\xcc\x57\xd0\x72\x74\xdc\x42\xda\x7e\x70\xc3\xda\x3b\xae\xd4\x5c\x1d\x2d\x85\x82\x3d\x87\x69\x1b\xda\xf5\xda\xdf\x64\x95\xf4\x7a\x4f\xd8\x53\x08\xb6\x2d\x77\xb7\xf5\x55\x87\xa4\x39\x6d\xea\x1b\xaa\x15\x5b\xdf\xab\x9b\x8a\x6a\x4c\xbb\x7d\x4f\x71\xcc\x38\x49\xcc\xaf\x28\x06\xf5\x64\xe5\xd1\x63\x74\x2f\x2e\x29\xb6\x06\xd3\x75\x6c\xda\xec\x3e\xa1\x71\x2a\xea\xfb\xe2\x21\x1a\xa7\x9c\x03\x15\x61\xe6\x22\x4c\x33\x74\x01\xd9\xe8\x12\x87\x29\x24\x08\x73\x40\x97\x38\x24\x3e\x0a\x80\xeb\x46\xc4\x52\x9d\xae\xb3\x42\x7a\x48\xed\xb7\x9f\x53\x32\xea\x82\xd8\x98\x24\xf6\x5c\x11\x79\x71\xf2\xd4\xf4\x1e\x47\x7b\xd0\xaa\x27\x5f\x71\x2a\x3a\x27\xb0\x96\xbb\xce\x49\x04\x93\x44\x90\x08\x0b\x83\x7e\xa7\x83\x3d\x26\x73\xd6\xfa\x0b\x80\x9e\x5d\x5e\xb5\x6a\xb3\x81\xe0\x29\xa6\xe9\x12\x7b\x22\xe5\x60\x31\x46\x7c\x0b\x61\x38\x0b\x70\x6c\x21\x28\x67\x79\xb9\x49\x5e\xba\x4f\x4b\xfe\xd6\x8e\xfc\xd8\x9c\xfc\x14\x5f\xbd\x61\xa1\x05\x3d\xa1\x56\xf4\xcf\x99\x10\x2c\xaa\xa2\x6b\xc3\x63\xe9\xa2\x77\x76\xe4\xef\xed\xc8\x3f\x98\x93\x8f\x59\x38\x23\xdf\x2d\xcc\x3d\x63\xdf\xec\x18\x5e\x02\x59\x05\xc2\xd2\x3d\xaf\x97\xcb\x04\x2c\x99\xde\xdf\x84\xe9\x9d\xe2\x52\xa0\x53\xd0\x0d\x78\x3e\x58\xf2\x4c\xae\x04\xc7\x9b\xe4\x65\x5b\x31\xea\x52\xda\x09\xe2\x9d\x36\xad\xe6\xcb\xbc\x19\xc7\x14\xe7\xa7\xfc\xbe\x6e\x9b\x8c\xaf\x9a\x8a\xeb\x17\xab\x8b\x26\x01\xd5\x80\xa7\xeb\xb2\xa5\x2f\x58\x5d\x8d\x6c\x3a\xde\x43\x95\xb7\x10\x7b\x1f\x9f\x19\xde\x72\x6d\x8e\xc6\xcc\x98\x5e\xb0\x19\xf7\xec\x58\x9e\x51\x11\x60\x6b\xae\x02\x83\x26\x7f\x10\x08\xfd\xbb\xfb\x68\x71\x5b\xd3\xdb\x7c\xb4\x53\x42\xb1\x5f\xc8\xa2\x7a\x3a\xf4\x0b\x19\x75\x8c\x05\xfe\xaf\x9a\xb3\x97\x4e\xd7\x3e\xa6\x76\x9b\x2d\xc1\xdd\x2a\x76\x57\x51\xc6\xae\xaa\x12\x5c\x4d\x32\xb9\x3b\xf1\xd8\xb9\x72\x3a\xb8\x3e\xf8\x37\x00\x00\xff\xff\x37\x1d\xea\x38\x98\x2f\x00\x00")

func workflowSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_workflowSchemaJson,
		"workflow.schema.json",
	)
}

func workflowSchemaJson() (*asset, error) {
	bytes, err := workflowSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "workflow.schema.json", size: 12184, mode: os.FileMode(0640), modTime: time.Unix(1558105158, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6, 0xe8, 0x25, 0x89, 0x1a, 0xcc, 0x8d, 0xe, 0x3f, 0xeb, 0x19, 0x89, 0x77, 0x27, 0x62, 0x1a, 0x69, 0xf2, 0x5a, 0x2c, 0x4a, 0xb6, 0xbe, 0x6, 0xc1, 0xda, 0x6c, 0xa6, 0x5b, 0x32, 0xa5, 0xc0}}
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
	"workflow.schema.json": workflowSchemaJson,
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
	"workflow.schema.json": &bintree{workflowSchemaJson, map[string]*bintree{}},
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
