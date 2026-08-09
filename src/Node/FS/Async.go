package Node_FS_Async

import (
    "os"
	"gopurs/output/gopurs_runtime"
)

func unboxString(v interface{}) string {
    return v.(string)
}

func MkdirImpl(path string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    err := os.MkdirAll(path, 0755)
    if err != nil {
        cb(err, nil)
    } else {
        cb(nil, nil)
    }
    return nil
}

func ReadFileImpl(filepath string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
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
    return nil
}

func WriteFileImpl(filepath string, content string, opts interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    data := []byte(content)
    err := os.WriteFile(filepath, data, 0644)
    if err != nil {
        cb(err, nil)
    } else {
        cb(nil, nil)
    }
    return nil
}

func AccessImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func CopyFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func MkdtempImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func RenameImpl(oldPath string, newPath string, cb func(interface{}, interface{}) interface{}) interface{} {
	
	err := os.Rename(oldPath, newPath)
	if err != nil {
		cb(err, nil)
	} else {
		cb(nil, nil)
	}
	return nil
}
func TruncateImpl(path string, len interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, nil)
    return nil
}
func ChownImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func ChmodImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func StatImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, nil)
    return nil
}
func LstatImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, nil)
    return nil
}
func LinkImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func SymlinkImpl(target string, path string, type_ string, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, nil)
    return nil
}
func ReadlinkImpl(arg0 interface{}, arg1 interface{}) interface{} { ; return nil }
func RealpathImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func UnlinkImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
	err := os.Remove(path)
	if err != nil {
		cb(err, nil)
	} else {
		cb(nil, nil)
	}
	return nil
}
func RmdirImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func RmImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func ReaddirImpl(path string, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, []interface{}{})
    return nil
}
func UtimesImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func AppendFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func OpenImpl(path string, flags interface{}, mode interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, 0)
    return nil
}
func ReadImpl(fd interface{}, buf interface{}, off interface{}, len interface{}, pos interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, 0)
    return nil
}
func WriteImpl(fd interface{}, buf interface{}, off interface{}, len interface{}, pos interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, 0)
    return nil
}
func CloseImpl(fd interface{}, cb func(interface{}, interface{}) interface{}) interface{} {
    cb(nil, nil)
    return nil
}

