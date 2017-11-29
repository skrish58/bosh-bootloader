// Code generated by go-bindata.
// sources:
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x5d\x6a\xc5\x20\x10\x85\xdf\x5d\xc5\x30\xf4\x35\xd9\x41\x57\x52\x8a\x18\x9d\xb6\xd2\x44\xc3\xf8\xd3\xd2\xe0\xde\x8b\x21\xb9\x10\x6f\xc2\xbd\xbe\xfa\xcd\x78\xbe\x23\x53\xf0\x89\x35\x01\xaa\xbf\xc4\xc4\x93\xcc\x96\x63\x52\xa3\x74\x14\x7f\x3c\x7f\x23\xe0\xe0\xc3\x17\xc2\x22\x00\x9c\x9a\x08\x9a\xf3\x0a\xf8\xb2\x64\xc5\x3d\xb9\x2c\xad\x29\x5d\xc5\xbb\xec\x50\x00\x28\x63\x98\x42\x90\x61\x56\x9a\x6e\xfc\xdb\x36\xb0\xbd\x20\xb5\x35\x5c\xf0\x5d\x00\x8c\x5e\xab\x68\xbd\x3b\xdd\xbf\x5f\x96\xba\x79\xcf\x2d\x3f\xd9\xa7\x59\xae\xc1\x56\x72\xd7\x38\x02\x7d\x0d\xd5\x57\xaa\xa0\x28\x42\xdc\x6b\x87\x34\x38\x8a\x0f\x6d\x2f\x74\xc3\x41\x77\x66\xfa\xb0\xbf\xed\xc0\x51\xf7\xc2\xe1\x69\x09\x80\xe6\xa3\x4e\x3a\x68\x88\xa6\x84\xff\x00\x00\x00\xff\xff\xfe\x40\x76\xae\xfb\x01\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 507, mode: os.FileMode(480), modTime: time.Unix(1511971757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcd\x6e\x9b\x40\x10\x80\xef\x3c\xc5\x68\xd5\x53\x24\x5b\x29\x81\xc8\x97\x1c\x7a\xec\xbd\x77\xb4\xde\x1d\xf0\xaa\x78\x07\xcd\x2e\x4e\xdb\x88\x77\xaf\x16\x8a\x13\x1c\x1c\x61\xd4\xaa\xc2\xb2\xaf\xfe\xe6\x6f\xe7\xd3\x08\x46\x47\x35\x2b\x04\x21\x7f\xd5\x8c\xbc\xcf\x2c\xfa\x67\xe2\xef\x99\x43\x55\xb3\xf1\x3f\xb3\x82\xa9\xae\x04\x88\x2d\xb9\x9d\x80\x97\x08\xc0\xca\x3d\xc2\xc9\xef\x09\xc4\xa7\x97\x83\xe4\x35\xda\x43\x66\x74\xb3\x6a\xf1\x08\xa0\x24\x25\xbd\x21\x3b\x0a\xf7\x7f\x36\x81\xec\x7b\xe9\x2a\x66\x6d\x95\x96\xec\x5b\x1b\x02\xeb\x50\x61\x1d\xa8\x46\x44\x11\x80\x97\x85\x6b\xdb\x03\x40\x7b\x30\x4c\x76\x8f\xd6\xbf\x6b\x2c\x54\x6a\xa2\x26\x8a\x2e\x18\x5d\xe5\x17\x0c\xae\xf2\xa5\x8f\xcd\x75\x89\x02\x84\xfb\x68\xdf\x67\xc7\x77\xdd\xda\x2b\x36\x14\x92\x8d\xc7\xc4\xf7\xf7\x11\x80\x36\x8c\xea\xf4\x91\x5e\xf3\x7e\xb5\x5b\xaa\xad\x0e\xd9\xa4\x52\xe8\xdc\xd9\x0e\xbe\x94\x25\x3d\x77\x55\xc9\x93\xa2\xf2\x0c\xf7\x4d\x55\x81\xfa\xf3\x9c\x15\xb1\xcf\x58\xda\x02\x87\xd4\x5d\x60\x34\x3a\x6f\x6c\xbb\xa5\x77\xe0\x13\x88\x38\x7e\x93\x48\x6a\xcd\xe8\x5c\x56\x31\xe6\xe6\xc7\x07\x89\x4e\xc1\x9e\x19\x53\x60\xf0\xc0\x13\x54\x00\x18\xd7\x77\x44\xa8\x71\x70\x90\xed\x22\x51\x42\xe0\x4a\x16\x68\xfd\x0c\x5f\xde\x04\x4f\xd0\xe6\xf3\xb2\xb5\x79\xdc\x3c\x6e\x6e\xe2\x0c\xc5\xe9\xd6\x49\x3c\xd7\x9d\x63\xfc\x04\x7d\xe2\x65\xeb\x13\xa7\x69\x9a\xde\xfc\x39\xfa\xa3\xad\x9b\x61\x4d\x88\x9a\xe0\xca\xc3\xff\x70\xe5\xee\x2f\x99\x92\x3e\xdc\x34\x39\x6a\xa2\x18\xf5\xae\xde\xce\x50\xa5\x8f\x9c\xa0\x4b\xb2\xec\xd3\xb2\xd9\x24\xc9\x4d\x99\x57\x65\xf2\xd5\xce\xfb\xea\x1f\x9e\x97\x85\x7f\xc9\x24\xc9\x15\x5e\x18\x95\xcf\x95\xa5\xa4\x62\xce\x79\xe9\x02\xaf\xff\xc3\x25\xb9\x7a\x5d\x7e\x07\x00\x00\xff\xff\x4c\xca\x07\xe0\x49\x11\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 4425, mode: os.FileMode(480), modTime: time.Unix(1511971755, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcb\x8e\xa3\x30\x10\x45\xf7\x7c\x85\x85\x66\x4d\x4b\x91\xd8\x44\x9a\x6f\xb1\x8c\xa9\xee\xd4\xb4\xb1\xad\x7a\x30\x9d\x89\xf2\xef\x23\x42\x88\xe2\x24\x30\x0f\x96\xf8\x9e\x53\x57\x65\x27\x95\xac\x62\xea\x2e\xf1\xc1\x46\x90\x9f\x89\x3e\x6d\x74\x03\xd4\xe6\x54\x19\x33\xba\xa0\x60\xbe\x9b\xfa\xdb\xc9\xfd\x52\x02\x1a\xec\x88\x24\xea\xc2\x12\x6e\x26\xb2\x99\x88\x73\x5d\x9d\xab\xaa\x10\xb2\x76\x11\x64\xdb\x37\x67\x36\x35\x04\x9c\x94\x3c\xd8\x0f\x4a\x9a\xb7\x75\x65\x76\xbb\x9d\x24\x72\x1f\x60\x9d\xf7\x49\xe3\x9f\x6a\x96\xe1\x4d\x71\x0f\xef\x4e\x83\x58\x06\xaf\x84\x72\x9c\xbb\xac\xaa\x97\xb5\x97\xf1\xb5\x09\xf0\x25\x40\xd1\x05\x8b\xeb\xc6\xac\x5d\x40\x6f\xf1\x2a\xc1\x6c\x5d\xdf\x13\x30\x97\xaa\x1e\x09\xbc\x24\x5a\x4e\x1f\x7c\x07\x91\xcc\xfb\xb7\xb7\xbf\xf1\xee\x77\x6d\xdb\xb6\xcf\xab\x18\x07\xb6\x99\x70\x74\x02\xf6\x13\x8e\xf7\x13\xa6\xef\xd2\x5a\x42\x91\x69\x16\xb0\xb9\xfb\x69\x33\x0c\xe7\xba\x32\x86\x21\x32\x0a\x8e\x53\x43\x21\x85\xd7\x13\xe7\x9e\xff\x3e\xf0\xc6\xd9\x94\x21\x32\x1f\x9e\x66\xbe\xbb\xc0\xc5\xd0\x1f\x3a\xe4\x2e\x7d\x59\xa5\xf0\x1f\xf7\xb1\xdf\xed\x8a\xa5\x2d\x6f\xc1\x63\x4f\x4f\xba\xd1\x51\x73\x1f\x28\x6f\x13\xe3\xf5\x61\xac\xb2\x45\xa2\x84\x59\x3b\xf6\x84\x59\x30\x45\x8b\xfd\x4b\xfc\x21\x53\x0a\x04\xa2\x8b\xb2\x86\xde\x4e\x4b\xc8\x07\x84\x75\xe8\x76\xfa\x12\x62\xf0\x04\xb2\x05\xce\x89\x0b\xfc\x3b\x00\x00\xff\xff\xe0\xed\x1c\x75\xe7\x04\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1255, mode: os.FileMode(480), modTime: time.Unix(1511971749, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x6a\xc5\x30\x0c\x44\xf7\x3e\xc5\x60\xba\x6d\x6e\xd0\xb3\x18\xc5\x11\xa9\x21\xb1\x83\x6c\x67\xd1\xe0\xbb\x17\x1b\xd2\x26\x24\xff\x7f\xf8\x5a\x8a\xd1\x68\xe6\x09\xc7\x90\xc5\x32\x34\xfd\x64\x61\x99\xcd\xbe\x31\xa3\x84\xbc\x68\xe8\x3e\xc4\x6f\x8d\x4d\x01\x9e\x66\x46\x9d\x2f\xe8\x8f\x6d\x25\xe9\xd8\xaf\xc6\x0d\xe5\xb3\x69\x14\x30\x05\x4b\xc9\x05\xff\xaf\xd8\x37\x45\x2b\x05\x24\x1a\x63\xb3\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x8b\x5f\xb5\x2a\xaa\x28\x75\x8d\xb7\xe4\x7e\x72\xd6\xb8\x07\xc9\xee\xe6\x75\xda\xa7\x57\x87\x06\xc0\x99\x8e\x39\xff\x6d\x27\xf7\x1c\xbb\xfa\xb3\xab\xf2\x66\xf3\xd7\xc2\xd0\x30\x08\xc7\x68\x68\x3a\xb2\x8b\x89\x92\xb3\xef\x20\xfb\x0d\x00\x00\xff\xff\xe9\x3c\x7f\x17\xd1\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 465, mode: os.FileMode(480), modTime: time.Unix(1511971760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6e\xc3\x20\x0c\xc7\xef\x3c\x85\x85\x76\xee\x1b\xec\xbc\xfb\xfa\x00\xc8\x21\x56\x86\x44\x20\xb2\x4d\xa6\xad\xca\xbb\x4f\xb0\x36\x6b\x91\xaa\xf6\xb8\x5c\xf3\xf3\xff\x0b\x26\xc9\x85\x3d\x81\xc5\xef\xc2\xc4\xb3\x13\xcd\x8c\x13\x39\xf4\x3e\x97\xa4\x16\xec\x90\xe5\xc3\xc2\xc9\x00\x24\x9c\x09\xba\xef\x15\xec\xcb\x69\x45\x3e\x48\x98\x97\x48\x8e\xd2\xea\xc2\xb8\x59\x03\x70\x11\x77\x13\xe7\xb2\xb8\x76\xdd\xf0\x8b\xd7\x2d\x70\xa8\x46\x87\x4a\x6d\xd6\x18\x80\x98\x3d\x6a\xc8\xa9\x77\xac\x1a\x9f\x24\x5a\xa4\x9a\x9c\x73\x3a\x0d\xc4\x3d\x75\x54\x4c\x23\xf2\x78\xcd\x31\x2d\x31\xfc\x0a\x3b\xfd\x5a\x5a\xa2\xb7\xf7\x63\x73\x54\x9c\xa4\x15\x05\xa0\xb4\x06\xce\x69\xa6\xa4\x7f\x15\xaf\xba\x6d\x66\x33\xe6\xfe\x7a\x3e\x27\xc5\x90\x88\x1f\xee\xd7\x82\x36\xe4\xce\x62\xf0\xf4\x66\x00\xdd\xe3\x9d\x05\x6e\xee\x3b\xa4\x13\xd8\x73\xd7\xff\x24\xb2\x4f\xb4\x70\x58\x51\xc9\x3e\x5f\x5b\x94\x66\x4f\x31\x3e\xa8\xbe\x63\xff\xba\xfe\x10\xf3\x50\xbb\xff\x04\x00\x00\xff\xff\x40\x93\x87\xbe\x30\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 816, mode: os.FileMode(480), modTime: time.Unix(1511971764, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1511971762, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xd1\x6a\xc5\x20\x0c\x40\xdf\xfd\x8a\x20\x7b\xee\xd6\x31\xf6\xb6\x6f\x29\xa9\x66\x10\x66\xb5\x44\xeb\x60\xa5\xff\x3e\x5a\xc1\xed\x8a\xdc\xdb\xbe\xe5\x9c\x03\x26\x19\x85\x71\x76\x04\x9a\x7c\x9e\xd8\x6a\xd8\x0f\xa5\xfe\xa6\x2e\x18\x4c\x1c\x7c\x3b\x8f\xbc\xac\x8e\xa6\x7e\x14\xb7\x39\x1a\xe1\xf5\x0c\x3b\x38\x91\x47\x9f\x3a\xc0\x38\xa6\x7b\x20\x92\x11\x4a\x2d\xf4\x94\xbe\x83\x7c\x4d\x86\xad\x68\xd8\x15\x80\xa5\x4f\xdc\x5c\x82\x0f\xd0\xe3\xcb\x70\xfd\xcf\xe3\xbb\x56\x37\x19\xfb\x44\xe2\xd1\x3d\xe8\x5e\xdf\xae\x6e\x95\x90\xd9\x92\x80\xc6\x9f\x4d\x48\x96\x52\x34\x9b\x9e\xe5\xd3\x9e\x51\x86\x06\x1c\x5a\x01\xd4\xbd\xa1\x7c\x55\xae\xe0\xd2\xea\x15\x5a\xad\x82\xff\x5a\xb9\x49\x47\x2b\xe0\x38\x5f\xff\x1b\x00\x00\xff\xff\xbf\x05\x53\xff\xe5\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 485, mode: os.FileMode(480), modTime: time.Unix(1511971767, 0)}
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
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

