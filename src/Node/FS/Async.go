package Node_FS_Async

import (
    "os"
    "fmt"
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

func AccessImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: accessImpl"); return nil }
func CopyFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: copyFileImpl"); return nil }
func MkdtempImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: mkdtempImpl"); return nil }
func RenameImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: renameImpl"); return nil }
func TruncateImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: truncateImpl"); return nil }
func ChownImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: chownImpl"); return nil }
func ChmodImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: chmodImpl"); return nil }
func StatImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: statImpl"); return nil }
func LstatImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: lstatImpl"); return nil }
func LinkImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: linkImpl"); return nil }
func SymlinkImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: symlinkImpl"); return nil }
func ReadlinkImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: readlinkImpl"); return nil }
func RealpathImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: realpathImpl"); return nil }
func UnlinkImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: unlinkImpl"); return nil }
func RmdirImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: rmdirImpl"); return nil }
func RmImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} { panic("Not implemented: rmImpl"); return nil }
func ReaddirImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: readdirImpl"); return nil }
func UtimesImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: utimesImpl"); return nil }
func AppendFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: appendFileImpl"); return nil }
func OpenImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} { panic("Not implemented: openImpl"); return nil }
func ReadImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}) interface{} { panic("Not implemented: readImpl"); return nil }
func WriteImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}) interface{} { panic("Not implemented: writeImpl"); return nil }
func CloseImpl(arg0 interface{}, arg1 interface{}) interface{} { panic("Not implemented: closeImpl"); return nil }

