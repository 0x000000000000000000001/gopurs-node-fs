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
        cb(nil, string(data)) // assuming encoding is utf8 for now
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
func TruncateImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func ChownImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func ChmodImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func StatImpl(arg0 interface{}, arg1 interface{}) interface{} { ; return nil }
func LstatImpl(arg0 interface{}, arg1 interface{}) interface{} { ; return nil }
func LinkImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { ; return nil }
func SymlinkImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
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
func ReaddirImpl(arg0 interface{}, arg1 interface{}) interface{} { ; return nil }
func UtimesImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func AppendFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func OpenImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { ; return nil }
func ReadImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}) interface{} { ; return nil }
func WriteImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}) interface{} { ; return nil }
func CloseImpl(arg0 interface{}, arg1 interface{}) interface{} { ; return nil }

