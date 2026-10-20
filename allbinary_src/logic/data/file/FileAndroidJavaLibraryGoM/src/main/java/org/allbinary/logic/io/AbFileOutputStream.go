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

type AbFileOutputStream struct {
    OutputStream
    FileOutputStream *FileOutputStream
}


func (this *AbFileOutputStream) initFields() {
    
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromAbFile(file *AbFile) *AbFileOutputStream {
        return NewAbFileOutputStream(NewFileOutputStream(AbFileNativeUtil_Get(file)))
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromAbFileAppend(file *AbFile, append_ bool) *AbFileOutputStream {
        return NewAbFileOutputStream(NewFileOutputStream(file.GetPath()))
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromFilePath(name string) *AbFileOutputStream {
        return NewAbFileOutputStream(NewFileOutputStream(name))
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromFilePathAppend(name string, append_ bool) *AbFileOutputStream {
        return NewAbFileOutputStream(NewFileOutputStream(name))
}


func NewAbFileOutputStream(fileOutputStream *FileOutputStream) *AbFileOutputStream {
    this := &AbFileOutputStream{}
    this.OutputStream = *NewOutputStream()
    this.initFields()
        this.FileOutputStream = fileOutputStream
    return this
}


// throws IOException
func (this *AbFileOutputStream) Close() {
        this.FileOutputStream.Close()
}


func (this *AbFileOutputStream) GetChannel() *FileChannel {
        return this.FileOutputStream.GetChannel()
}


// throws IOException
func (this *AbFileOutputStream) Write(buffer []int8) {
        this.FileOutputStream.Write(buffer)
}


// throws IOException
// Java overloaded method: Go needs a different name
func (this *AbFileOutputStream) Write(buffer []int8, offset int, count int) {
        this.FileOutputStream.Write(buffer, offset, count)
}


// throws IOException
// Java overloaded method: Go needs a different name
func (this *AbFileOutputStream) Write(b int) {
        this.FileOutputStream.Write(b)
}

