package Node_FS_Stream

import (
    "gopurs/output/Node.EventEmitter"
)

func CreateWriteStreamImpl(filepath string) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func CreateWriteStreamOptsImpl(filepath string, opts interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func FdCreateWriteStreamImpl(fd interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func FdCreateWriteStreamOptsImpl(fd interface{}, opts interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func CreateReadStreamImpl(filepath string) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func CreateReadStreamOptsImpl(filepath string, opts interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func FdCreateReadStreamImpl(fd interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}

func FdCreateReadStreamOptsImpl(fd interface{}, opts interface{}) interface{} {
    return Node_EventEmitter.NewImpl(nil)
}
