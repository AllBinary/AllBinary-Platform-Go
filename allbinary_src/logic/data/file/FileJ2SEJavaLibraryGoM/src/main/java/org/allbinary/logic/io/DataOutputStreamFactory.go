/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type DataOutputStreamFactory struct {
}

var DataOutputStreamFactory_Instance *DataOutputStreamFactory = NewDataOutputStreamFactory()

func NewDataOutputStreamFactory() *DataOutputStreamFactory {
    this := &DataOutputStreamFactory{}
    return this
}


func DataOutputStreamFactory_GetInstance() *DataOutputStreamFactory {
        return DataOutputStreamFactory_Instance
}


// throws Exception
func (this *DataOutputStreamFactory) GetInstanceForAbFile(file *AbFile) *AbDataOutputStream {
        var idFileOutputStream *AbFileOutputStream = AbFileOutputStream_CreateFromAbFile(file); _ = idFileOutputStream
        var idOutData *AbDataOutputStream = NewAbDataOutputStream(idFileOutputStream); _ = idOutData
        return idOutData
}


// throws Exception
func (this *DataOutputStreamFactory) GetInstance(filePath string, fileName string) *AbDataOutputStream {
        var fileStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileStreamFactory
        var fileOutputStream *AbFileOutputStream = fileStreamFactory.GetFileOutputStreamInstance(filePath, fileName); _ = fileOutputStream
        var idOutData *AbDataOutputStream = NewAbDataOutputStream(fileOutputStream); _ = idOutData
        return idOutData
}

