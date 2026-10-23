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
}

var FileFactory_Instance *FileFactory = NewFileFactory()

func NewFileFactory() *FileFactory {
    this := &FileFactory{}
    return this
}


func FileFactory_GetInstance() *FileFactory {
        return FileFactory_Instance
}


// throws Exception
func (this *FileFactory) GetInstance_1(filePath string) *AbFile {
        return AbFile_CreateAbFileFromRawPath(filePath)
}

// Java overloads of FileFactory.getInstance
func (this *FileFactory) GetInstance(args ...any) *AbFile {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetInstance_1(Java_Arg[string](args[0]))
        }
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetInstance_2(Java_Arg[*AbFile](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of FileFactory.getInstance for the arguments")
}


// throws Exception
func (this *FileFactory) GetInstance_2(file *AbFile, childFilePath string) *AbFile {
        return AbFile_CreateAbFileWithChild(file, childFilePath)
}


// throws Exception
func (this *FileFactory) IsFile(path string) bool {
        var file *AbFile = AbFile_CreateAbFile(path); _ = file
        if file.Exists() {
            return true
        } else {
            return false
        }
}

