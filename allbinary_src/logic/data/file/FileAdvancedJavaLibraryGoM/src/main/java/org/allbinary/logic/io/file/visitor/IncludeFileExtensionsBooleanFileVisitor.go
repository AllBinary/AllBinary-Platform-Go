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

type IncludeFileExtensionsBooleanFileVisitor struct {
    BooleanFileVisitor
}



func NewIncludeFileExtensionsBooleanFileVisitor(filterStringBasicArrayList *BasicArrayList) *IncludeFileExtensionsBooleanFileVisitor {
    this := &IncludeFileExtensionsBooleanFileVisitor{}
        this.BooleanFileVisitor = *NewBooleanFileVisitor(filterStringBasicArrayList)
    return this
}


func (this *IncludeFileExtensionsBooleanFileVisitor) Visit(file *AbFile) bool {
        if this.GetFilterStringBasicArrayList().Size() == 0 {
            return false
        }
        return this.BooleanFileVisitor.Visit(file)
}


// Java overloaded method: Go needs a different name
func (this *IncludeFileExtensionsBooleanFileVisitor) Visit(file *AbFile, fileNameString string) bool {
        var fileNameAndExtension string = file.GetPath(); _ = fileNameAndExtension
        if len(fileNameAndExtension) > len(fileNameString) {
            var startIndex int = len(fileNameAndExtension) - len(fileNameString); _ = startIndex
            var fileName string = fileNameAndExtension[startIndex:]; _ = fileName
            if strings.Compare(fileName, fileNameString) == 0 {
                return true
            }
        }
        return false
}

