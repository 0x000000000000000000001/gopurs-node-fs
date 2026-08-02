package Node_FS_Sync

import (
    "os"
)

func ReadFileSyncImpl(filepath string, opts interface{}) interface{} {
    data, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func WriteFileSyncImpl(filepath string, content string, opts interface{}) interface{} {
    data := []byte(content)
    err := os.WriteFile(filepath, data, 0644)
    if err != nil {
        panic(err)
    }
    return nil
}

func AccessImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: accessImpl")
    return nil
}

func CopyFileImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: copyFileImpl")
    return nil
}

func MkdtempImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: mkdtempImpl")
    return nil
}

func RenameSyncImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: renameSyncImpl")
    return nil
}

func TruncateSyncImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: truncateSyncImpl")
    return nil
}

func ChownSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: chownSyncImpl")
    return nil
}

func ChmodSyncImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: chmodSyncImpl")
    return nil
}

func StatSyncImpl(arg0 interface{}) interface{} {
    panic("Not implemented: statSyncImpl")
    return nil
}

func LstatSyncImpl(arg0 interface{}) interface{} {
    panic("Not implemented: lstatSyncImpl")
    return nil
}

func LinkSyncImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: linkSyncImpl")
    return nil
}

func SymlinkSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: symlinkSyncImpl")
    return nil
}

func ReadlinkSyncImpl(arg0 interface{}) interface{} {
    panic("Not implemented: readlinkSyncImpl")
    return nil
}

func RealpathSyncImpl(arg0 interface{}, arg1 interface{}) interface{} {
    panic("Not implemented: realpathSyncImpl")
    return nil
}

func UnlinkSyncImpl(arg0 interface{}) interface{} {
    panic("Not implemented: unlinkSyncImpl")
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

func UtimesSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: utimesSyncImpl")
    return nil
}

func AppendFileSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: appendFileSyncImpl")
    return nil
}

func OpenSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
    panic("Not implemented: openSyncImpl")
    return nil
}

func ReadSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}) interface{} {
    panic("Not implemented: readSyncImpl")
    return nil
}

func WriteSyncImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}) interface{} {
    panic("Not implemented: writeSyncImpl")
    return nil
}

func CloseSyncImpl(arg0 interface{}) interface{} {
    panic("Not implemented: closeSyncImpl")
    return nil
}

func ExistsSyncImpl(filepath string) interface{} {
    _, err := os.Stat(filepath)
    return err == nil
}
