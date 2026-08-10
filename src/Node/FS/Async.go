package Node_FS_Async

import (
	"os"
	"path/filepath"
	"gopurs/output/gopurs_runtime"
)

func unboxString(v interface{}) string {
	return v.(string)
}

func MkdirImpl(path string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func ReadFileImpl(filepath string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		data, err := os.ReadFile(filepath)
		if err != nil {
			cb(err, nil)
		} else {
			hasEncoding := false
			if v, ok := opts.(gopurs_runtime.Value); ok {
				m := gopurs_runtime.UnboxObject(v)
				_, hasEncoding = m["encoding"]
			} else if m, ok := opts.(map[string]interface{}); ok {
				_, hasEncoding = m["encoding"]
			}
			if hasEncoding {
				cb(nil, string(data))
			} else {
				cb(nil, gopurs_runtime.Any(data))
			}
		}
	}()
	return nil
}

func WriteFileImpl(filepath string, content string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		data := []byte(content)
		err := os.WriteFile(filepath, data, 0644)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func AccessImpl(filepath string, mode interface{}, cb func(interface{}) interface{}) interface{} {
	go func() {
		info, err := os.Stat(filepath)
		if err != nil {
			cb(err)
			return
		}
		
		modeVal, ok := mode.(gopurs_runtime.Value)
		if ok && modeVal.IntVal == 2 {
			isReadOnly := info.Mode().Perm()&0222 == 0
			if isReadOnly {
				cb(os.ErrPermission)
				return
			}
		}
		cb(nil)
	}()
	return nil
}

func CopyFileImpl(src string, dest string, mode interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		modeVal, ok := mode.(gopurs_runtime.Value)
		if ok && modeVal.IntVal == 1 { // COPYFILE_EXCL
			if _, err := os.Stat(dest); err == nil {
				cb(os.ErrExist, nil)
				return
			}
		}
		
		data, err := os.ReadFile(src)
		if err != nil {
			cb(err, nil)
			return
		}
		err = os.WriteFile(dest, data, 0644)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func MkdtempImpl(prefix string, encoding string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		dir := filepath.Dir(prefix)
		pat := filepath.Base(prefix) + "*"
		name, err := os.MkdirTemp(dir, pat)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, name)
		}
	}()
	return nil
}

func RenameImpl(oldPath string, newPath string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func TruncateImpl(path string, len interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		l := 0.0
		if val, ok := len.(gopurs_runtime.Value); ok {
			l = val.FloatVal()
		}
		err := os.Truncate(path, int64(l))
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func ChownImpl(path string, uid interface{}, gid interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		cb(nil, nil)
	}()
	return nil
}

func ChmodImpl(path string, mode interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		cb(nil, nil)
	}()
	return nil
}

func StatImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		info, err := os.Stat(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, info)
		}
	}()
	return nil
}

func LstatImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		info, err := os.Lstat(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, info)
		}
	}()
	return nil
}

func LinkImpl(oldpath string, newpath string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.Link(oldpath, newpath)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func SymlinkImpl(target string, path string, type_ string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.Symlink(target, path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func ReadlinkImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		link, err := os.Readlink(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, link)
		}
	}()
	return nil
}

func RealpathImpl(p string, cache interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		link, err := filepath.EvalSymlinks(p)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, link)
		}
	}()
	return nil
}

func UnlinkImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.Remove(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func RmdirImpl(path string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.Remove(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func RmImpl(path string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		err := os.RemoveAll(path)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func ReaddirImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		entries, err := os.ReadDir(path)
		if err != nil {
			cb(err, nil)
			return
		}
		names := make([]interface{}, len(entries))
		for i, entry := range entries {
			names[i] = entry.Name()
		}
		cb(nil, names)
	}()
	return nil
}

func UtimesImpl(path string, atime interface{}, mtime interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() { cb(nil, nil) }()
	return nil
}

func AppendFileImpl(filepath string, content string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			cb(err, nil)
			return
		}
		defer f.Close()
		if _, err := f.WriteString(content); err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}

func OpenImpl(filepath string, flags interface{}, mode interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, float64(f.Fd()))
		}
	}()
	return nil
}

func ReadImpl(fd interface{}, buffer interface{}, offset interface{}, length interface{}, position interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		fdFloat := 0.0
		if val, ok := fd.(gopurs_runtime.Value); ok {
			fdFloat = val.FloatVal()
		} else if v, ok := fd.(float64); ok {
			fdFloat = v
		}
		f := os.NewFile(uintptr(fdFloat), "")
		l := 0.0
		if val, ok := length.(gopurs_runtime.Value); ok {
			l = val.FloatVal()
		} else if v, ok := length.(float64); ok {
			l = v
		}
		
		buf := make([]byte, int(l))
		var n int
		var err error
		
		if position == nil {
			n, err = f.Read(buf)
		} else {
			posVal, ok := position.(gopurs_runtime.Value)
			if ok && posVal.Type == gopurs_runtime.TypeInt {
				n, err = f.ReadAt(buf, posVal.IntVal)
			} else {
				n, err = f.Read(buf)
			}
		}
		
		if err != nil && err.Error() != "EOF" {
			cb(err, nil)
		} else {
			cb(nil, float64(n))
		}
	}()
	return nil
}

func WriteImpl(fd interface{}, buffer interface{}, offset interface{}, length interface{}, position interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		fdFloat := 0.0
		if val, ok := fd.(gopurs_runtime.Value); ok {
			fdFloat = val.FloatVal()
		} else if v, ok := fd.(float64); ok {
			fdFloat = v
		}
		f := os.NewFile(uintptr(fdFloat), "")
		
		var data []byte
		if v, ok := buffer.(gopurs_runtime.Value); ok {
			if ptr := v.UnsafePtr; ptr != nil {
				if b, ok2 := (*(*any)(ptr)).([]byte); ok2 {
					data = b
				}
			}
		} else if str, ok := buffer.(string); ok {
			data = []byte(str)
		}
		
		l := 0.0
		if val, ok := length.(gopurs_runtime.Value); ok {
			l = val.FloatVal()
		} else if v, ok := length.(float64); ok {
			l = v
		}
		
		if data == nil {
			data = make([]byte, int(l))
		}
		
		var n int
		var err error
		if position == nil {
			n, err = f.Write(data[:int(l)])
		} else {
			posVal, ok := position.(gopurs_runtime.Value)
			if ok && posVal.Type == gopurs_runtime.TypeInt {
				n, err = f.WriteAt(data[:int(l)], posVal.IntVal)
			} else {
				n, err = f.Write(data[:int(l)])
			}
		}
		
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, float64(n))
		}
	}()
	return nil
}

func CloseImpl(fd interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
	go func() {
		fdFloat := 0.0
		if val, ok := fd.(gopurs_runtime.Value); ok {
			fdFloat = val.FloatVal()
		} else if v, ok := fd.(float64); ok {
			fdFloat = v
		}
		f := os.NewFile(uintptr(fdFloat), "")
		err := f.Close()
		if err != nil {
			cb(err, nil)
		} else {
			cb(nil, nil)
		}
	}()
	return nil
}
