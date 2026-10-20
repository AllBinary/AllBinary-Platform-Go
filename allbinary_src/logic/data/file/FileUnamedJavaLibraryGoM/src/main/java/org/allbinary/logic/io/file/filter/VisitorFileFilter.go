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

type VisitorFileFilter struct {
    BooleanVisitorInterface VisitorInterface
}

// Implements *FileFilter


func (this *VisitorFileFilter) initFields() {
    
}


func NewVisitorFileFilter(booleanVisitorInterface VisitorInterface) *VisitorFileFilter {
    this := &VisitorFileFilter{}
    this.initFields()
        this.BooleanVisitorInterface = booleanVisitorInterface
    return this
}


func (this *VisitorFileFilter) Accept(file *File) bool {
        var abFile *AbFile = FileWrapperUtil_WrapFile(file); _ = abFile
        var isFileAcceptable bool = (any(this.BooleanVisitorInterface.Visit(abFile)).(bool)); _ = isFileAcceptable
        return isFileAcceptable || file.IsDirectory()
}


// Java overloaded method: Go needs a different name
func (this *VisitorFileFilter) Accept(file *AbFile) bool {
        var isFileAcceptable bool = (any(this.BooleanVisitorInterface.Visit(file)).(bool)); _ = isFileAcceptable
        return isFileAcceptable || file.IsDirectory()
}

