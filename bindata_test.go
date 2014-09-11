package migrate

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

func test_migrations_1_initial_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0xcd,
		0x41, 0x0e, 0x82, 0x40, 0x0c, 0x05, 0xd0, 0x7d, 0x4f, 0xf1, 0x77, 0x68,
		0x0c, 0x27, 0x60, 0x85, 0xc2, 0x8e, 0x44, 0x45, 0x38, 0x00, 0x42, 0x03,
		0x8d, 0xe3, 0x4c, 0x33, 0x8c, 0xc1, 0xe3, 0x3b, 0x9a, 0x48, 0xe2, 0xce,
		0x5d, 0xfb, 0xdb, 0xfc, 0x97, 0xa6, 0xd8, 0x8d, 0xce, 0x2b, 0x5a, 0xa5,
		0x38, 0x5f, 0xce, 0x15, 0xc4, 0x62, 0xe6, 0x3e, 0x88, 0xb3, 0x48, 0x5a,
		0x4d, 0x20, 0x33, 0xf8, 0xc9, 0xfd, 0x23, 0xf0, 0x80, 0x65, 0x62, 0x8b,
		0x30, 0xc5, 0xe8, 0x2e, 0xa3, 0xef, 0x3e, 0x4f, 0x71, 0xe9, 0x54, 0x8d,
		0xf0, 0x40, 0x87, 0xba, 0xcc, 0x9b, 0x12, 0x4d, 0xbe, 0xaf, 0x4a, 0x28,
		0x3b, 0x35, 0x8c, 0x8d, 0x0c, 0xb1, 0x32, 0x6c, 0x33, 0x22, 0x5a, 0xb5,
		0xc2, 0x2d, 0xf6, 0xeb, 0xad, 0xd8, 0x3b, 0xfc, 0x8b, 0xf3, 0xce, 0x98,
		0x78, 0xbd, 0x76, 0xfd, 0x8d, 0x8a, 0xfa, 0x78, 0xfa, 0x01, 0x33, 0x7a,
		0x05, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x2f, 0x94, 0xfa, 0xd4, 0x00, 0x00,
		0x00,
	},
		"test-migrations/1_initial.sql",
	)
}

func test_migrations_2_record_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd2, 0xd5,
		0x55, 0xd0, 0x4e, 0xcf, 0x2f, 0x2a, 0x50, 0x08, 0x2d, 0xe0, 0xf2, 0xf4,
		0x0b, 0x76, 0x0d, 0x0a, 0x51, 0xf0, 0xf4, 0x0b, 0xf1, 0x57, 0x28, 0x48,
		0xcd, 0x2f, 0xc8, 0x49, 0x55, 0xd0, 0xc8, 0x4c, 0xd1, 0x54, 0x08, 0x73,
		0xf4, 0x09, 0x75, 0x0d, 0x56, 0xd0, 0x30, 0xd4, 0xb4, 0xe6, 0xe2, 0x82,
		0x6b, 0x70, 0xc9, 0x2f, 0xcf, 0xe3, 0x72, 0x71, 0xf5, 0x71, 0x0d, 0x71,
		0x55, 0x70, 0x0b, 0xf2, 0xf7, 0x85, 0x69, 0x09, 0xf7, 0x70, 0x0d, 0x72,
		0x55, 0xc8, 0x4c, 0xb1, 0x35, 0xb4, 0xe6, 0x02, 0x04, 0x00, 0x00, 0xff,
		0xff, 0x37, 0xa3, 0x5f, 0x55, 0x5e, 0x00, 0x00, 0x00,
	},
		"test-migrations/2_record.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"test-migrations/1_initial.sql": test_migrations_1_initial_sql,
	"test-migrations/2_record.sql": test_migrations_2_record_sql,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test-migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"1_initial.sql": &_bintree_t{test_migrations_1_initial_sql, map[string]*_bintree_t{
		}},
		"2_record.sql": &_bintree_t{test_migrations_2_record_sql, map[string]*_bintree_t{
		}},
	}},
}}