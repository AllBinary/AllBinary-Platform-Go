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

type SubDirectory struct {
    Directory *Directory
}

var SubDirectory_Instance *SubDirectory = NewSubDirectory()

func (this *SubDirectory) initFields() {
    this.Directory = Directory_GetInstance()
}


func SubDirectory_GetInstance() *SubDirectory {
        return SubDirectory_Instance
}


func NewSubDirectory() *SubDirectory {
    this := &SubDirectory{}
    this.initFields()
    return this
}


func (this *SubDirectory) Search(fileFilter *FileFilter, file *AbFile) *BasicArrayList {
        return this.Directory.Search(fileFilter, file, true)
}


// Java overloaded method: Go needs a different name
func (this *SubDirectory) Search(file *AbFile) *BasicArrayList {
        return this.Directory.Search(file, true)
}


// Java overloaded method: Go needs a different name
func (this *SubDirectory) Search(searchValue string, file *AbFile) *BasicArrayList {
        return this.Directory.Search(searchValue, file, true)
}


// Java overloaded method: Go needs a different name
func (this *SubDirectory) Search(level int, file *AbFile) *BasicArrayList {
        return this.Directory.Search(level, file, true)
}


func SubDirectory_ToString(files *BasicArrayList) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var NEW_LINE string = CommonSeps_GetInstance().NEW_LINE; _ = NEW_LINE
        var size int = files.Size(); _ = size
        for index := int(0); index < size; index++ {
            var file *AbFile = any(files.Get(index)).(*AbFile); _ = file
            stringBuffer.Append(file.GetPath())
            stringBuffer.Append(NEW_LINE)
        }
        return fmt.Sprint(stringBuffer)
}

