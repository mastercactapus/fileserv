package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xc1\x6e\x9c\x30\x10\x3d\x6f\xbe\xc2\xf5\xa1\xb7\x2c\x3f\x80\x89\xd4\xa4\x39\xa5\xed\xaa\x49\x0f\x3d\x3a\x78\x76\x19\xd5\x98\x95\x3d\x8a\x4a\x11\xff\xde\x31\x06\x76\xb3\x11\x55\xe8\xc9\x1e\xcf\xbc\x79\xef\x19\x3c\xf9\x87\xbb\x6f\xb7\x4f\x3f\x77\x9f\x45\x45\xb5\x2d\xae\xf2\x69\x01\x6d\x8a\xab\x4d\x4e\x48\x16\x8a\x3b\xf4\x50\x52\xe3\x5b\xf1\x80\x81\xd0\x1d\xc4\xb5\xe8\xba\xed\x4e\x53\xd5\xf7\x79\x96\x8a\xb8\xda\xa2\xfb\x25\x3c\x58\x25\x03\xb5\x16\x42\x05\x40\x52\x50\x7b\x04\x25\x09\x7e\x53\x56\x86\x20\x45\xe5\x61\xaf\x64\x76\xa3\x43\x00\x52\xb5\x46\xb7\x8d\xe7\xcc\x9a\x25\xda\xfc\xb9\x31\x2d\x2f\xa4\x9f\x2d\x08\x34\x4a\x96\x8d\x23\x70\x14\x8b\x58\xd3\x28\x2e\xee\x84\xb6\x78\x70\x4a\x5a\xd8\x53\x4c\x6e\xba\x0e\xf7\x62\x7b\x3b\xd6\x6f\x1f\x1b\x4f\x9f\xda\xaf\xba\x86\xbe\xe7\xec\x26\xd7\xa2\xb4\x4c\xac\xa4\x2e\x09\x5f\x60\x52\x33\xbb\xb9\x09\x8c\x50\x8e\x01\x1f\x3d\xbc\x80\x0f\xa0\x38\x37\xf7\xfb\x9e\xce\xee\x2d\x1e\xfb\x5e\x16\xb1\x71\x9e\xe9\x44\x0c\x36\x9c\x58\x16\xdb\x5e\x82\x9c\x19\x30\x7c\x8b\xd5\x2a\x4f\x5f\x1a\x83\x7b\x04\xb3\xca\x57\x3d\x82\xde\xe3\xed\x41\x07\x12\x13\x60\x85\xc9\x09\xb2\xd8\xe1\x7f\x1d\x3f\xe2\x9f\x75\x5f\x31\x30\xe0\x3d\x4e\x63\xe3\x15\x06\x63\xdb\x4b\xd0\x6b\x4f\x71\x99\x1e\x50\xfa\x97\x47\x47\x3b\xed\x59\xc0\x0f\x6f\x53\x35\xf9\x62\x20\x22\x53\xbc\x26\x9b\xab\x64\x91\x02\x31\x3f\xc1\xc8\xc9\xfd\xcd\x09\x79\x1e\x5c\x4f\x11\xaf\x7e\xa0\x1d\xa5\x75\x9d\xd7\xee\x00\x67\x97\x7a\x8f\xfc\x42\xff\xa9\x63\x54\xc0\xbb\xf4\x80\xde\x50\x73\xe6\xf4\x1b\x5e\x66\xd2\xe7\x7a\xab\x67\xba\xdf\x73\xd6\xa7\x0a\x83\x30\xf3\x94\xe1\x00\xea\x23\xb5\xcb\x3e\x97\x6c\x72\x3c\x0e\x8f\x6c\x98\x1e\x71\x33\x1d\x0c\xa3\xed\x6f\x00\x00\x00\xff\xff\x63\xe7\x83\x51\xf1\x04\x00\x00")

func assets_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_assets_index_html,
		"assets/index.html",
	)
}

func assets_index_html() (*asset, error) {
	bytes, err := assets_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/index.html", size: 1265, mode: os.FileMode(436), modTime: time.Unix(1423441555, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assets_main_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x5d\x8a\xe3\x30\x10\x84\x9f\xad\x53\x08\xc2\xbe\x64\x51\x22\xb3\x6b\x58\x94\xd3\xc8\x56\xdb\x6e\xa2\x1f\x23\x75\x36\x3f\x43\xee\x3e\x92\x13\x4f\x42\x60\x98\x3c\x18\x44\xf5\xd7\xe5\x72\x59\xa4\x5b\x0b\xfc\x83\x55\x47\x34\x34\x2a\x5e\x4b\xf9\x6b\xc7\xaa\x36\x44\x03\x51\x74\xc1\x5a\x3d\x25\x50\x7c\x39\xed\xd8\x95\x31\xb6\x5d\x7b\xed\xa0\x88\xeb\x2d\xa3\x91\x3f\x19\xfc\x2d\xfb\xd7\x4c\xb8\x60\xb0\x47\x30\x0f\xea\x37\xbf\xa1\xf3\x38\xe1\x05\x5e\x46\xcb\x7c\xb1\x6a\x66\x27\x46\xb1\x68\xe1\x40\x16\x3d\x28\xf9\x95\x6e\x3e\x4e\xda\x18\xf4\x83\xe2\x72\xd3\x80\xe3\xe5\x91\x9b\x7f\xe0\xca\x26\x45\x35\x86\xff\x10\x39\x99\xe2\xd0\xea\x6e\x3f\xc4\x70\xf0\x46\xad\x00\xe0\x46\xe4\x99\xea\x31\x26\x12\xa1\x17\x74\x9e\xe6\x2e\xee\xa6\xc2\x42\x4f\xaa\x96\xd3\xe9\xf1\x22\x11\x71\x18\x17\x31\x1b\x98\x27\x5e\x35\xd3\x89\x97\x50\xdf\x64\xbd\x97\xda\x06\xa2\xe0\x54\x9d\xe1\x14\x2c\x1a\xbe\x02\xd3\xd7\x7d\x93\x89\x1c\x96\xb0\xd3\x56\x68\x8b\x83\x57\x0e\x8d\xb1\x39\x68\x45\x70\xa2\xbb\x56\x32\x65\x45\x1c\xa1\xdd\x23\x09\x8a\xda\x27\x24\x0c\x5e\x3d\xbe\x8f\xff\x91\xd2\x25\x0e\xba\xfc\xaf\x4a\xb8\x70\x79\x8f\x4b\x6f\x61\xe1\x1d\xea\x67\x64\x2e\x4f\x97\xfa\x0c\xa6\xc9\xea\xb3\xe2\xad\x0d\xdd\x7e\xf7\x72\x17\xaf\xec\x33\x00\x00\xff\xff\x20\x3e\xfe\xa6\xa4\x02\x00\x00")

func assets_main_css_bytes() ([]byte, error) {
	return bindata_read(
		_assets_main_css,
		"assets/main.css",
	)
}

func assets_main_css() (*asset, error) {
	bytes, err := assets_main_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/main.css", size: 676, mode: os.FileMode(436), modTime: time.Unix(1423440212, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"assets/index.html": assets_index_html,
	"assets/main.css": assets_main_css,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{assets_index_html, map[string]*_bintree_t{
		}},
		"main.css": &_bintree_t{assets_main_css, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

