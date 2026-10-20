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
    FileOutputStream
}



// throws FileNotFoundException
func AbFileOutputStream_CreateFromAbFile(file *AbFile) *AbFileOutputStream {
        return NewAbFileOutputStream(file.GetPath(), false)
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromAbFileAppend(file *AbFile, append_ bool) *AbFileOutputStream {
        return NewAbFileOutputStream(file.GetPath(), append_)
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromFilePath(name string) *AbFileOutputStream {
        return NewAbFileOutputStream(name, false)
}


// throws FileNotFoundException
func AbFileOutputStream_CreateFromFilePathAppend(name string, append_ bool) *AbFileOutputStream {
        return NewAbFileOutputStream(name, append_)
}


func NewAbFileOutputStream(name string, append_ bool) *AbFileOutputStream {
    this := &AbFileOutputStream{}
        this.FileOutputStream = *NewFileOutputStream(name, append_)
    return this
}

