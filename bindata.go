// Code generated by go-bindata.
// sources:
// prom-templates/deployment.yaml
// prom-templates/kustomization.yaml
// prom-templates/route.yaml
// prom-templates/service.yaml
// DO NOT EDIT!

package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _promTemplatesDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x53\x3d\x73\x13\x31\x10\xed\xfd\x2b\xb6\xa0\x3d\xfb\x9c\x0c\x45\xd4\x31\x64\x06\x0a\x02\x9e\xe0\xa1\x5f\xeb\x36\x3e\x91\xd3\x07\xbb\x2b\x27\x26\xc9\x7f\x67\x14\xc7\x77\x32\x84\x86\x02\x55\xf2\x7b\x7a\x6f\xf7\xad\xf7\x30\xb9\x6f\xc4\xe2\x62\x30\x40\xf7\x4a\xa1\x5c\x65\xb1\x5b\x6e\x48\x71\x39\xbb\x75\xa1\x33\x70\x49\x69\x88\x7b\x4f\x41\x67\x9e\x14\x3b\x54\x34\x33\x80\x80\x9e\x0c\x24\x8e\x7e\x26\x89\x6c\x81\x98\xd2\xe0\x2c\x8a\x81\xe5\x0c\x40\xc9\xa7\x01\x95\x0a\x03\x50\x4b\xcb\xa9\xe4\x07\x60\xc0\x0d\x0d\x72\xa4\x01\x30\xa5\x8a\x3f\x96\x28\xc7\x05\xa7\xef\x63\x50\x74\x81\x78\x54\x34\x2f\x96\xd6\x35\x37\xa4\xb6\x27\x1e\xad\x9c\xc7\x2d\x19\x60\xda\x3a\x51\xde\xcf\x6f\xa8\x8b\x8c\x89\xe3\x77\xb2\x3a\x8f\xbc\x5d\x1c\x10\x73\xde\x8e\x1a\x1b\xbd\xc7\xd0\x4d\xfd\x34\xb0\xd8\xb8\xb0\xd8\xa0\xf4\x15\xd6\xd8\xea\xc7\xe3\x78\x07\x10\x52\x68\xf2\x7d\x84\xe4\x12\xdd\xa0\x1b\x2a\x2e\x7b\x94\x5b\x68\xdb\xb6\xad\x40\x9b\x79\x80\x46\x3e\xc1\x9b\x87\xd5\xf5\x97\xab\xf5\xbb\xeb\x27\x78\x04\x45\x86\xfb\xdd\x4f\x68\xfc\xf8\xf4\x2e\xf2\xad\x0b\xdb\x4b\xc7\x06\x16\x65\x40\xa4\x3d\x65\x59\x8c\x0f\x76\x71\xc8\x9e\xae\x62\x0e\x2a\x75\xff\xd3\xc8\x0f\x8a\x46\x34\x32\x6e\xa9\x39\x08\xaa\x5e\x7c\xd1\xae\x50\xfb\x57\x2b\xd8\xbf\xce\x7e\x7a\xfb\xfb\xec\x0b\x53\x59\x99\xdd\xd9\x7c\x79\x36\x9f\xf2\xa7\xc8\xaf\x35\x7b\x47\x9b\xec\xea\x21\x1d\x4b\xaf\x22\xab\x81\x8b\xf6\x62\xb2\x60\xc2\xce\x05\x12\x59\x71\xdc\x90\xa9\x54\xbd\x6a\xfa\x40\x5a\x43\x00\xe9\x10\xef\x14\xfb\xd3\xf5\xf9\xbf\xb4\x3d\x95\x6e\x3e\xae\xd7\xab\x8a\x50\xe7\x29\x66\xfd\x4a\x36\x86\xee\x65\xe9\x47\x27\x62\x17\xbb\x89\xaa\x0d\x25\x5b\x4b\x22\xeb\x9e\x49\xfa\x38\x74\xa7\xca\xb2\x2b\x99\xa9\x62\xcf\xab\x88\x12\x33\x5b\x92\x3a\x0a\xd3\x8f\x4c\xa2\x72\x1a\xcf\xa6\x5c\xca\xb6\xfe\x04\xf5\xe4\x23\xef\x0d\xbc\x6d\xdb\x2b\xf7\xdf\x16\x46\x7a\x64\x5a\x71\x2c\xa9\x3f\xa3\x27\x49\x68\xc9\x80\x72\x3e\x7a\x1c\x0c\xff\xb1\x38\xf9\xa4\xfb\xe7\xcf\xe1\xe1\x69\xf6\x2b\x00\x00\xff\xff\x95\x4d\xdc\x4e\xd3\x04\x00\x00")

func promTemplatesDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesDeploymentYaml,
		"prom-templates/deployment.yaml",
	)
}

func promTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := promTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/deployment.yaml", size: 1235, mode: os.FileMode(420), modTime: time.Unix(1570199473, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesKustomizationYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcd\x4e\xc3\x30\x10\x84\xef\x7e\x8a\x55\xef\x06\x71\xf5\xad\x40\x6f\x09\xa9\x0a\xf4\x5a\x2d\xce\xb4\x58\xc4\x3f\x5a\x3b\x11\x95\x78\x78\x94\x34\x6a\x5a\xdf\x66\x3e\xef\xee\x67\xa3\xf7\x31\x54\xfc\x85\x2e\x1b\x45\xc4\x29\x19\x7a\x69\xea\xba\x79\x3b\x54\xeb\xe7\x4d\xa5\x02\x7b\x6c\x05\x47\xf7\x7b\x0f\xb4\x12\xe4\xd8\x8b\xc5\x34\xa8\xa9\x45\xea\xe2\xd9\x23\x94\x87\x33\xfb\x6e\xea\x24\xf6\x05\x4b\xcc\x90\xc1\xd9\xb9\x48\x5c\xec\x37\xf2\x7b\x11\x2e\x38\x39\x5b\x43\x4e\x30\x4a\xd3\x9f\x9e\x44\xdc\x1e\x92\x5d\x0c\x66\x94\xca\x8f\xc3\x93\x22\xfa\x71\xa1\x35\xf4\x7a\xbd\xa4\x88\x3c\x0a\xb7\x5c\x78\x94\x20\x1a\x6d\x0d\x25\x89\x5e\x11\xe5\x04\x7b\xa9\x0b\x7c\xea\xb8\xe0\x92\x6e\xc9\xf8\x6c\x0c\x85\x5d\x80\xe4\xa5\xd3\xf3\x2a\xeb\xf4\x11\xa3\xa8\x5c\x11\x11\xc2\x60\x6e\xe2\xf2\x7b\xbb\x6b\xea\x8f\xf5\xee\x8e\x11\x0d\xdc\xf5\x30\xb4\x9a\xe9\x61\xbf\xae\x3e\x37\x2b\xf5\x1f\x00\x00\xff\xff\x2a\x85\xc0\x74\x7c\x01\x00\x00")

func promTemplatesKustomizationYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesKustomizationYaml,
		"prom-templates/kustomization.yaml",
	)
}

func promTemplatesKustomizationYaml() (*asset, error) {
	bytes, err := promTemplatesKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/kustomization.yaml", size: 380, mode: os.FileMode(420), modTime: time.Unix(1570199613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesRouteYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x4e\x03\x31\x0c\x44\xef\xf9\x0a\x7f\x41\xd9\x1e\x9b\x3b\x57\x54\x15\xc4\xdd\x4a\x86\xd4\x62\x13\x47\x8e\xdb\x8a\xbf\x47\xd9\x05\xc1\x25\xd2\xbc\x37\x19\xf9\x53\x5a\x8e\x74\xd1\x9b\x23\x70\x97\x77\xd8\x10\x6d\x91\x6c\x92\x83\x76\xb4\x71\x95\x0f\x3f\x88\x3e\xdd\x8f\xa1\xc2\x39\xb3\x73\x0c\x44\x8d\x2b\x22\x75\xd3\x1a\x46\x47\x9a\xc8\x75\xbe\x44\xfb\xe8\x2b\xec\x2e\x09\x1b\xf9\x57\x9e\xf1\x01\x29\x57\x8f\x74\x5c\x96\x40\xd4\xd5\x7c\xff\xe8\x6c\x05\x7e\x9e\x99\x4e\xcb\x69\x4a\x5f\xc7\x8f\x83\x55\x69\xec\xdb\x79\xc8\x65\x1f\x96\x36\x90\x6e\x86\xe7\x5c\xf0\xf6\xd7\x38\xeb\x2a\xe9\x2b\xd2\x05\x59\x0c\xc9\x03\xd1\x43\xd6\x9c\xd8\xf2\xaf\x7a\xd1\x86\xf0\x1d\x00\x00\xff\xff\x84\xf0\x9a\x83\xff\x00\x00\x00")

func promTemplatesRouteYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesRouteYaml,
		"prom-templates/route.yaml",
	)
}

func promTemplatesRouteYaml() (*asset, error) {
	bytes, err := promTemplatesRouteYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/route.yaml", size: 255, mode: os.FileMode(420), modTime: time.Unix(1570192158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\x31\xae\xc2\x30\x0c\xc6\xf1\x3d\xa7\xf8\x2e\xf0\xa4\xbe\xb1\x5e\xb9\x00\x12\x88\xdd\x4d\x3d\x44\x24\xb1\x95\x98\x72\x7d\x94\x92\x89\xd1\x3f\x7f\xfa\xb3\xa5\x87\xb4\x9e\xb4\x12\x8e\xff\xf0\x4c\x75\x27\xdc\xa4\x1d\x29\x4a\x28\xe2\xbc\xb3\x33\x05\xa0\x72\x11\x82\x35\x2d\x01\xc8\xbc\x49\xee\x83\x01\x36\x9b\xde\x4d\xe2\x30\xd3\xe6\xf3\xf9\x77\x1e\x84\x75\x59\x97\x13\x30\xa6\xae\x51\x33\xe1\x7e\xb9\x4e\xfb\xc6\xdf\xb2\xbd\x52\x00\xba\x64\x89\xae\xed\xb7\xff\x09\x00\x00\xff\xff\x6d\xe1\xdb\x34\xac\x00\x00\x00")

func promTemplatesServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesServiceYaml,
		"prom-templates/service.yaml",
	)
}

func promTemplatesServiceYaml() (*asset, error) {
	bytes, err := promTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/service.yaml", size: 172, mode: os.FileMode(420), modTime: time.Unix(1570192162, 0)}
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
	"prom-templates/deployment.yaml": promTemplatesDeploymentYaml,
	"prom-templates/kustomization.yaml": promTemplatesKustomizationYaml,
	"prom-templates/route.yaml": promTemplatesRouteYaml,
	"prom-templates/service.yaml": promTemplatesServiceYaml,
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
	"prom-templates": &bintree{nil, map[string]*bintree{
		"deployment.yaml": &bintree{promTemplatesDeploymentYaml, map[string]*bintree{}},
		"kustomization.yaml": &bintree{promTemplatesKustomizationYaml, map[string]*bintree{}},
		"route.yaml": &bintree{promTemplatesRouteYaml, map[string]*bintree{}},
		"service.yaml": &bintree{promTemplatesServiceYaml, map[string]*bintree{}},
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

