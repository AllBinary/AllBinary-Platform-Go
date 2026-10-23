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

type DirectoryOrIncludeFileExtensionBooleanFileVisitor struct {
    IncludeFileExtensionsBooleanFileVisitor
    IncludesString string
}


func (this *DirectoryOrIncludeFileExtensionBooleanFileVisitor) initFields() {
    
}




// Java overloads of new DirectoryOrIncludeFileExtensionBooleanFileVisitor
func NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(args ...any) *DirectoryOrIncludeFileExtensionBooleanFileVisitor {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IBasicArrayList](args[0], pass) {
            return NewDirectoryOrIncludeFileExtensionBooleanFileVisitor_1(Java_Arg[IBasicArrayList](args[0]))
        }
        if len(args) == 2 && Java_IsArg[IBasicArrayList](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewDirectoryOrIncludeFileExtensionBooleanFileVisitor_2(Java_Arg[IBasicArrayList](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of new DirectoryOrIncludeFileExtensionBooleanFileVisitor for the arguments")
}
func NewDirectoryOrIncludeFileExtensionBooleanFileVisitor_1(filterStringBasicArrayList IBasicArrayList) *DirectoryOrIncludeFileExtensionBooleanFileVisitor {
    this := &DirectoryOrIncludeFileExtensionBooleanFileVisitor{}
        this.IncludeFileExtensionsBooleanFileVisitor = *NewIncludeFileExtensionsBooleanFileVisitor(filterStringBasicArrayList)
        this.initFields()
        this.IncludesString = nil
    return this
}


func NewDirectoryOrIncludeFileExtensionBooleanFileVisitor_2(filterStringBasicArrayList IBasicArrayList, includesString string) *DirectoryOrIncludeFileExtensionBooleanFileVisitor {
    this := &DirectoryOrIncludeFileExtensionBooleanFileVisitor{}
        this.IncludeFileExtensionsBooleanFileVisitor = *NewIncludeFileExtensionsBooleanFileVisitor(filterStringBasicArrayList)
        this.initFields()
        this.IncludesString = includesString
    return this
}


func (this *DirectoryOrIncludeFileExtensionBooleanFileVisitor) Visit(file *AbFile) bool {
        if file.IsDirectory() {
            return true
        }
        if this.IncludesString == "" || strings.Index(file.GetAbsolutePath(), string(this.IncludesString)) >= 0 {
            return this.IncludeFileExtensionsBooleanFileVisitor.Visit(file)
        } else {
            return false
        }
}

