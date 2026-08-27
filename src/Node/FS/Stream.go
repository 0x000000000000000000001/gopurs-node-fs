package Node_FS_Stream

import (
	"os"
	"gopurs/output/gopurs_runtime"
)

func CreateWriteStreamImpl(filepath string) interface{} {
	f, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	return Node_EventEmitter_NewImpl(f)
}

func CreateWriteStreamOptsImpl(filepath string, opts interface{}) interface{} {
	f, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	return Node_EventEmitter_NewImpl(f)
}

func FdCreateWriteStreamImpl(fd interface{}) interface{} {
	fdFloat := 0.0
	if val, ok := fd.(gopurs_runtime.Value); ok {
		fdFloat = val.FloatVal()
	} else if v, ok := fd.(float64); ok {
		fdFloat = v
	}
	f := os.NewFile(uintptr(fdFloat), "")
	return Node_EventEmitter_NewImpl(f)
}

func FdCreateWriteStreamOptsImpl(fd interface{}, opts interface{}) interface{} {
	fdFloat := 0.0
	if val, ok := fd.(gopurs_runtime.Value); ok {
		fdFloat = val.FloatVal()
	} else if v, ok := fd.(float64); ok {
		fdFloat = v
	}
	f := os.NewFile(uintptr(fdFloat), "")
	return Node_EventEmitter_NewImpl(f)
}

func CreateReadStreamImpl(filepath string) interface{} {
	f, _ := os.Open(filepath)
	return Node_EventEmitter_NewImpl(f)
}

func CreateReadStreamOptsImpl(filepath string, opts interface{}) interface{} {
	f, _ := os.Open(filepath)
	return Node_EventEmitter_NewImpl(f)
}

func FdCreateReadStreamImpl(fd interface{}) interface{} {
	fdFloat := 0.0
	if val, ok := fd.(gopurs_runtime.Value); ok {
		fdFloat = val.FloatVal()
	} else if v, ok := fd.(float64); ok {
		fdFloat = v
	}
	f := os.NewFile(uintptr(fdFloat), "")
	return Node_EventEmitter_NewImpl(f)
}

func FdCreateReadStreamOptsImpl(fd interface{}, opts interface{}) interface{} {
	fdFloat := 0.0
	if val, ok := fd.(gopurs_runtime.Value); ok {
		fdFloat = val.FloatVal()
	} else if v, ok := fd.(float64); ok {
		fdFloat = v
	}
	f := os.NewFile(uintptr(fdFloat), "")
	return Node_EventEmitter_NewImpl(f)
}
