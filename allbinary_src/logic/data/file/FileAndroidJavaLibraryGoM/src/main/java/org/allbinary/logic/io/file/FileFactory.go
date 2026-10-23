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

type FileFactory struct {
    Context *Context
}

var FileFactory_SINGLETON *FileFactory = NewFileFactory(ResourceUtil_GetInstance().GetContext())

func (this *FileFactory) initFields() {
    
}


func NewFileFactory(context *Context) *FileFactory {
    this := &FileFactory{}
    this.initFields()
        this.Context = context
    return this
}


func FileFactory_GetInstance() *FileFactory {
        return FileFactory_SINGLETON
}


// throws Exception
func (this *FileFactory) IsFile(path string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.Context.OpenFileInput(path)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *FileFactory) GetContext() *Context {
        return this.Context
}


// throws Exception
func (this *FileFactory) GetFile(filePath string) *AbFile {
        return AbFile_CreateAbFileFromRawPath(filePath)
}


// throws Exception
func (this *FileFactory) GetInstance(file *AbFile, childFilePath string) *AbFile {
        return AbFile_CreateAbFileWithChild(file, childFilePath)
}

