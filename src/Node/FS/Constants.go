package Node_FS_Constants

import "gopurs/output/gopurs_runtime"

// The values are: F_OK = 0, R_OK = 4, W_OK = 2, X_OK = 1
var F_OK = gopurs_runtime.Box(0)
var R_OK = gopurs_runtime.Box(4)
var W_OK = gopurs_runtime.Box(2)
var X_OK = gopurs_runtime.Box(1)

var CopyFile_EXCL = gopurs_runtime.Box(1)
var CopyFile_FICLONE = gopurs_runtime.Box(2)
var CopyFile_FICLONE_FORCE = gopurs_runtime.Box(4)

var O_RDONLY = gopurs_runtime.Box(0)
var O_WRONLY = gopurs_runtime.Box(1)
var O_RDWR = gopurs_runtime.Box(2)
var O_CREAT = gopurs_runtime.Box(64)
var O_EXCL = gopurs_runtime.Box(128)
var O_NOCTTY = gopurs_runtime.Box(256)
var O_TRUNC = gopurs_runtime.Box(512)
var O_APPEND = gopurs_runtime.Box(1024)
var O_DIRECTORY = gopurs_runtime.Box(65536)
var O_NOATIME = gopurs_runtime.Box(262144)
var O_NOFOLLOW = gopurs_runtime.Box(131072)
var O_SYNC = gopurs_runtime.Box(1052672)
var O_SYMLINK = gopurs_runtime.Box(2097152)
var O_DIRECT = gopurs_runtime.Box(16384)
var O_NONBLOCK = gopurs_runtime.Box(2048)

var S_IFMT = gopurs_runtime.Box(61440)
var S_IFREG = gopurs_runtime.Box(32768)
var S_IFDIR = gopurs_runtime.Box(16384)
var S_IFCHR = gopurs_runtime.Box(8192)
var S_IFBLK = gopurs_runtime.Box(24576)
var S_IFIFO = gopurs_runtime.Box(4096)
var S_IFLNK = gopurs_runtime.Box(40960)
var S_IFSOCK = gopurs_runtime.Box(49152)

var S_IRWXU = gopurs_runtime.Box(448)
var S_IRUSR = gopurs_runtime.Box(256)
var S_IWUSR = gopurs_runtime.Box(128)
var S_IXUSR = gopurs_runtime.Box(64)
var S_IRWXG = gopurs_runtime.Box(56)
var S_IRGRP = gopurs_runtime.Box(32)
var S_IWGRP = gopurs_runtime.Box(16)
var S_IXGRP = gopurs_runtime.Box(8)
var S_IRWXO = gopurs_runtime.Box(7)
var S_IROTH = gopurs_runtime.Box(4)
var S_IWOTH = gopurs_runtime.Box(2)
var S_IXOTH = gopurs_runtime.Box(1)

var UV_FS_SYMLINK_DIR = gopurs_runtime.Box(1)
var UV_FS_SYMLINK_JUNCTION = gopurs_runtime.Box(2)
