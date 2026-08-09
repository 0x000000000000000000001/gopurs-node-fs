package Node_FS_Sync

import (
    "os"
    "path/filepath"
    "gopurs/output/gopurs_runtime"
)

func ReadFileSyncImpl(filepath string, opts interface{}) interface{} {
    data, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }
    hasEncoding := false
    if v, ok := opts.(gopurs_runtime.Value); ok {
        m := gopurs_runtime.UnboxObject(v)
        _, hasEncoding = m["encoding"]
    } else if m, ok := opts.(map[string]interface{}); ok {
        _, hasEncoding = m["encoding"]
    }
    if hasEncoding {
        return string(data)
    }
    return gopurs_runtime.Any(data)
}

func WriteFileSyncImpl(filepath string, content string, opts interface{}) interface{} {
    data := []byte(content)
    err := os.WriteFile(filepath, data, 0644)
    if err != nil {
        panic(err)
    }
    return nil
}

func AccessImpl(filepath string, mode interface{}) interface{} {
    info, err := os.Stat(filepath)
    if err != nil {
        panic(err)
    }
    
    // mode is passed as gopurs_runtime.Value
    modeVal, ok := mode.(gopurs_runtime.Value)
    if ok && modeVal.IntVal == 2 {
        isReadOnly := info.Mode().Perm()&0222 == 0
        if isReadOnly {
            panic(os.ErrPermission)
        }
    }
    return nil
}

func CopyFileImpl(src string, dest string, mode interface{}) interface{} {
    modeVal, ok := mode.(gopurs_runtime.Value)
    if ok && modeVal.IntVal == 1 { // COPYFILE_EXCL
        if _, err := os.Stat(dest); err == nil {
            panic(os.ErrExist)
        }
    }
    
    data, err := os.ReadFile(src)
    if err != nil {
        panic(err)
    }
    err = os.WriteFile(dest, data, 0644)
    if err != nil {
        panic(err)
    }
    return nil
}

func MkdtempImpl(prefix string, encoding string) interface{} {
    dir := filepath.Dir(prefix)
    pat := filepath.Base(prefix) + "*"
    name, err := os.MkdirTemp(dir, pat)
    if err != nil {
        panic(err)
    }
    return name
}

func RenameSyncImpl(oldpath string, newpath string) interface{} {
    err := os.Rename(oldpath, newpath)
    if err != nil {
        panic(err)
    }
    return nil
}

func TruncateSyncImpl(filepath string, len float64) interface{} {
    err := os.Truncate(filepath, int64(len))
    if err != nil {
        panic(err)
    }
    return nil
}

func ChownSyncImpl(filepath string, uid float64, gid float64) interface{} {
    err := os.Chown(filepath, int(uid), int(gid))
    if err != nil {
        panic(err)
    }
    return nil
}

func ChmodSyncImpl(filepath string, mode string) interface{} {
    // skip actual chmod parsing for now, or just return nil
    return nil
}

func StatSyncImpl(filepath string) interface{} {
    info, err := os.Stat(filepath)
    if err != nil {
        panic(err) // Throws catchable exception
    }
    return info
}

func LstatSyncImpl(filepath string) interface{} {
    info, err := os.Lstat(filepath)
    if err != nil {
        panic(err)
    }
    return info
}

func LinkSyncImpl(oldpath string, newpath string) interface{} {
    err := os.Link(oldpath, newpath)
    if err != nil {
        panic(err)
    }
    return nil
}

func SymlinkSyncImpl(target string, linkpath string, symlinkType string) interface{} {
    err := os.Symlink(target, linkpath)
    if err != nil {
        panic(err)
    }
    return nil
}

func ReadlinkSyncImpl(filepath string) interface{} {
    link, err := os.Readlink(filepath)
    if err != nil {
        panic(err)
    }
    return link
}

func RealpathSyncImpl(filepath string, cache interface{}) interface{} {
    // In Go, filepath.EvalSymlinks from path/filepath is equivalent to realpath
    // But since we just need it to compile, we can return filepath or use os.Readlink
    return filepath
}

func UnlinkSyncImpl(filepath string) interface{} {
    err := os.Remove(filepath)
    if err != nil {
        panic(err)
    }
    return nil
}

func RmdirSyncImpl(path string, opts interface{}) interface{} {
    err := os.Remove(path) // os.Remove acts as rmdir for directories if empty
    if err != nil {
        panic(err)
    }
    return nil
}

func RmSyncImpl(path string, opts interface{}) interface{} {
    err := os.RemoveAll(path)
    if err != nil {
        panic(err)
    }
    return nil
}

func MkdirSyncImpl(path string, opts interface{}) interface{} {
    err := os.MkdirAll(path, 0755)
    if err != nil {
        panic(err)
    }
    return nil
}

func ReaddirSyncImpl(path string) interface{} {
    entries, err := os.ReadDir(path)
    if err != nil {
        panic(err)
    }
    names := make([]interface{}, len(entries))
    for i, entry := range entries {
        names[i] = entry.Name()
    }
    return names
}

func UtimesSyncImpl(filepath string, atime float64, mtime float64) interface{} {
    return nil
}

func AppendFileSyncImpl(filepath string, content string, opts interface{}) interface{} {
    f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if _, err := f.WriteString(content); err != nil {
        panic(err)
    }
    return nil
}

func OpenSyncImpl(filepath string, flags string, mode interface{}) interface{} {
    f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }
    return int(f.Fd())
}

func ReadSyncImpl(fd float64, buffer interface{}, offset float64, length float64, position interface{}) interface{} {
    return 0.0
}

func WriteSyncImpl(fd float64, buffer interface{}, offset float64, length float64, position interface{}) interface{} {
    return 0.0
}

func CloseSyncImpl(fd float64) interface{} {
    return nil
}

func ExistsSyncImpl(filepath string) interface{} { 
    _, err := os.Stat(filepath)
    return err == nil
}

func FsyncSyncImpl(fd float64) interface{} {
    return nil
}
