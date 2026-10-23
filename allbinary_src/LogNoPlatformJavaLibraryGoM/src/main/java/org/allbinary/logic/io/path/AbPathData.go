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

type AbPathData struct {
    EXTENSION_SEP string
    SEPARATORCHAR rune
    SEPARATOR string
}

var AbPathData_Instance *AbPathData = NewAbPathData()

func (this *AbPathData) initFields() {
    this.EXTENSION_SEP = CommonSeps_GetInstance().PERIOD
    this.SEPARATORCHAR = '/'
    this.SEPARATOR = CommonSeps_GetInstance().FORWARD_SLASH
}


func AbPathData_GetInstance() *AbPathData {
        return AbPathData_Instance
}


func NewAbPathData() *AbPathData {
    this := &AbPathData{}
    this.initFields()
    return this
}


func (this *AbPathData) GetExtensionIndex(filePath string) int {
        var indexOfFileExtensionDelmiter int = strings.LastIndex(filePath, string(this.EXTENSION_SEP)); _ = indexOfFileExtensionDelmiter
        var indexOfLatDelimiter int = strings.LastIndex(filePath, string(this.SEPARATORCHAR)); _ = indexOfLatDelimiter
        if indexOfFileExtensionDelmiter < 0 {
            return -1
        }
        if indexOfFileExtensionDelmiter < indexOfLatDelimiter {
            return -1
        }
        return indexOfFileExtensionDelmiter
}


func (this *AbPathData) GetExtensionWithDot(filePath string) string {
        var indexOfFileExtensionDelmiter int = this.GetExtensionIndex(filePath); _ = indexOfFileExtensionDelmiter
        var extension string = StringUtil_GetInstance().EMPTY_STRING; _ = extension
        if indexOfFileExtensionDelmiter >= 0 {
            extension = filePath[indexOfFileExtensionDelmiter:]
        }
        return extension
}


func (this *AbPathData) GetExtension(filePath string) string {
        var indexOfFileExtensionDelmiter int = this.GetExtensionIndex(filePath); _ = indexOfFileExtensionDelmiter
        var extension string = StringUtil_GetInstance().EMPTY_STRING; _ = extension
        if indexOfFileExtensionDelmiter >= 0 {
            extension = filePath[indexOfFileExtensionDelmiter + 1:]
        }
        return extension
}


func (this *AbPathData) GetNameFromPath(path string) string {
        var endIndex int = strings.LastIndex(path, string(this.SEPARATORCHAR)); _ = endIndex
        if endIndex < 0 {
            endIndex = strings.LastIndex(path, string(FilePathData_GetInstance().SEPARATORCHAR))
        }
        if endIndex < 0 {
            return path
        }
        if len(path) == endIndex + 1 {
            var categoryName string = path[0:endIndex]; _ = categoryName
            return this.GetNameFromPath(categoryName)
        } else {
            var categoryName string = path[endIndex + 1:]; _ = categoryName
            return categoryName
        }
}


func (this *AbPathData) RemoveNameFromPath(path string) string {
        var endIndex int = strings.LastIndex(path, string(this.SEPARATORCHAR)); _ = endIndex
        if endIndex < 0 {
            endIndex = strings.LastIndex(path, string(FilePathData_GetInstance().SEPARATORCHAR))
        }
        if endIndex < 0 {
            return StringUtil_GetInstance().EMPTY_STRING
        }
        if len(path) == endIndex + 1 {
            return this.RemoveNameFromPath(path[0:endIndex - 1])
        } else {
            var pathWithoutName string = path[0:endIndex]; _ = pathWithoutName
            return pathWithoutName
        }
}

