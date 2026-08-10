package Node_FS_Stats

import (
	"os"
	"gopurs/output/gopurs_runtime"
)

func unboxStats(arg0 interface{}) os.FileInfo {
	val := arg0.(gopurs_runtime.Value)
	return (*(*any)(val.UnsafePtr)).(os.FileInfo)
}

func IsFileImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return info.Mode().IsRegular()
}
func IsDirectoryImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return info.IsDir()
}
func IsBlockDeviceImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return (info.Mode() & os.ModeDevice) != 0 && (info.Mode() & os.ModeCharDevice) == 0
}
func IsCharacterDeviceImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return (info.Mode() & os.ModeCharDevice) != 0
}
func IsFIFOImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return (info.Mode() & os.ModeNamedPipe) != 0
}
func IsSocketImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return (info.Mode() & os.ModeSocket) != 0
}
func IsSymbolicLinkImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return (info.Mode() & os.ModeSymlink) != 0
}
func ModifiedTimeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func AccessedTimeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func StatusChangedTimeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func BirthTimeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func DevImpl(arg0 interface{}) interface{} { return 0.0 }
func ModeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.Mode())
}
func NlinkImpl(arg0 interface{}) interface{} { return 0.0 }
func UidImpl(arg0 interface{}) interface{} { return 0.0 }
func GidImpl(arg0 interface{}) interface{} { return 0.0 }
func RdevImpl(arg0 interface{}) interface{} { return 0.0 }
func BlkSizeImpl(arg0 interface{}) interface{} { return 0.0 }
func InodeImpl(arg0 interface{}) interface{} { return 0.0 }
func SizeImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.Size())
}
func BlocksImpl(arg0 interface{}) interface{} { return 0.0 }
func AccessedTimeMsImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func ModifiedTimeMsImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func StatusChangedTimeMsImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func BirthtimeMsImpl(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return float64(info.ModTime().UnixMilli())
}
func ShowStatsObj(arg0 interface{}) interface{} {
	info := unboxStats(arg0)
	return info.Name()
}
