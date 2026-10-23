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


func (this *SubDirectory) Search_1(fileFilter *FileFilter, file *AbFile) IBasicArrayList {
        return this.Directory.Search(fileFilter, file, true)
}

// Java overloads of SubDirectory.search
func (this *SubDirectory) Search(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*FileFilter](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_1(Java_Arg[*FileFilter](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return this.Search_2(Java_Arg[*AbFile](args[0]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_3(Java_Arg[string](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_4(Java_Arg[int](args[0]), Java_Arg[*AbFile](args[1]))
        }
    }
    panic("No Java overload of SubDirectory.search for the arguments")
}


func (this *SubDirectory) Search_2(file *AbFile) IBasicArrayList {
        return this.Directory.Search(file, true)
}


func (this *SubDirectory) Search_3(searchValue string, file *AbFile) IBasicArrayList {
        return this.Directory.Search(searchValue, file, true)
}


func (this *SubDirectory) Search_4(level int, file *AbFile) IBasicArrayList {
        return this.Directory.Search(level, file, true)
}


func SubDirectory_ToString(files IBasicArrayList) string {
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

