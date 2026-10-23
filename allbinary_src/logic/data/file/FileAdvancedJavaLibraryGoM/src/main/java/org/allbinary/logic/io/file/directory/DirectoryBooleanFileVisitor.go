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

type DirectoryBooleanFileVisitor struct {
}

// Implements VisitorInterface



func NewDirectoryBooleanFileVisitor() *DirectoryBooleanFileVisitor {
    this := &DirectoryBooleanFileVisitor{}
    return this
}


func (this *DirectoryBooleanFileVisitor) Visit_1(anyType any) any {
        return any(this.Visit(any(anyType).(*AbFile)).(any)).(any)
}

// Java overloads of DirectoryBooleanFileVisitor.visit
func (this *DirectoryBooleanFileVisitor) Visit(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[any](args[0], pass) {
            return this.Visit_1(Java_Arg[any](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return this.Visit_2(Java_Arg[*AbFile](args[0]))
        }
    }
    panic("No Java overload of DirectoryBooleanFileVisitor.visit for the arguments")
}


func (this *DirectoryBooleanFileVisitor) Visit_2(file *AbFile) bool {
        if file.IsDirectory() {
            return true
        }
        return false
}

