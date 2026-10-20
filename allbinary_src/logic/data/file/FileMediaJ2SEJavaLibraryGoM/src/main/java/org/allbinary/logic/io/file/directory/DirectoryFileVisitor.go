/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type DirectoryFileVisitor struct {
}

// Implements VisitorInterface, *FileFilter



func NewDirectoryFileVisitor() *DirectoryFileVisitor {
    this := &DirectoryFileVisitor{}
    return this
}


func (this *DirectoryFileVisitor) Accept(arg0 *File) bool {
        return (any(this.Visit(any(arg0).(any))).(bool))
}


func (this *DirectoryFileVisitor) Visit(anyType any) any {
        var file *File = any(anyType).(*File); _ = file
        var index int = 1; _ = index
        var aFile *File = file.GetParentFile(); _ = aFile
        for (func() any { aFile = aFile.GetParentFile(); return aFile }()) != nil {
            index++
        }
        if file.IsDirectory() && index < 5 {
            return true
        }
        return false
}

