package fsutil

import (
	"io"
	"os"
	"path"
)

// some commonly flag consts for open file
const (
	FsCWAFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND // create, append write-only
	FsCWTFlags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC  // create, override write-only
	FsCWFlags  = os.O_CREATE | os.O_WRONLY               // create, write-only
	FsRFlags   = os.O_RDONLY                             // read-only
)

// IsFile reports whether the named file or directory exists.
func IsFile(path string) bool {
	if path == "" || len(path) > 468 {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}

// MustReadFile read file contents, will panic on error
func MustReadFile(filePath string) []byte {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return bs
}

// OpenFile like os.OpenFile, but will auto create dir.
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)
	if err := os.MkdirAll(fileDir, 0o775); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// WriteOSFile write data to give os.File, then close file.
// data type allow: string, []byte, io.Reader
func WriteOSFile(f *os.File, data any) (n int, err error) {
	switch typData := data.(type) {
	case []byte:
		n, err = f.Write(typData)
	case string:
		n, err = f.WriteString(typData)
	case io.Reader: // eg: buffer
		var n64 int64
		n64, err = io.Copy(f, typData)
		n = int(n64)
	default:
		_ = f.Close()
		panic("WriteFile: data type only allow: []byte, string, io.Reader")
	}

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return n, err
}

// ExtName is ...
func ExtName(fpath string) string {
	if ext := path.Ext(fpath); len(ext) > 0 {
		return ext[1:]
	}
	return ""
}

// CopyFile copy a file to another file path.
func CopyFile(srcPath, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, FsRFlags, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := OpenFile(dstPath, FsCWTFlags, 0o665)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
