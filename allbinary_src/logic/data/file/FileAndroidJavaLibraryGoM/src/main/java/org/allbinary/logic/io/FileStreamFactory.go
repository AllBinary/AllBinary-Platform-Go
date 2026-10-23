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

type FileStreamFactory struct {
    Context any
}

var FileStreamFactory_SINGLETON any = NullUtil_GetInstance().NULL_OBJECT

func (this *FileStreamFactory) initFields() {
    
}


func FileStreamFactory_GetInstance() *FileStreamFactory {
        if FileStreamFactory_SINGLETON == NullUtil_GetInstance().NULL_OBJECT {
            FileStreamFactory_SINGLETON = NewFileStreamFactory(ResourceUtil_GetInstance().GetContext())
        }
        return any(FileStreamFactory_SINGLETON).(*FileStreamFactory)
}


func NewFileStreamFactory(context *Context) *FileStreamFactory {
    this := &FileStreamFactory{}
    this.initFields()
        this.Context = context
    return this
}


// throws Exception
func (this *FileStreamFactory) GetFileInputStreamInstance(path string, fileName string) IAbFileInputStream {
        return NewAbFileInputStream(this.GetContext().OpenFileInput(fileName))
}


// throws Exception
func (this *FileStreamFactory) GetFileOutputStreamInstance(path string, fileName string) *AbFileOutputStream {
        var context *Context = this.GetContext(); _ = context
        context.DeleteFile(fileName)
        return NewAbFileOutputStream(context.OpenFileOutput(fileName, 0))
}


// throws Exception
func (this *FileStreamFactory) Delete(path string, fileName string) {
        this.GetContext().DeleteFile(fileName)
}


func (this *FileStreamFactory) GetContext() *Context {
        return any(this.Context).(*Context)
}

