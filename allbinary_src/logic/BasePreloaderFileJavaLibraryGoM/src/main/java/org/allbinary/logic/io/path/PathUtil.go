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

//Used to correct user input for paths
type PathUtil struct {
    AbPathData *AbPathData
    FilePathData *FilePathData
}

var PathUtil_Instance *PathUtil = NewPathUtil()

func (this *PathUtil) initFields() {
    this.AbPathData = AbPathData_GetInstance()
    this.FilePathData = FilePathData_GetInstance()
}


func PathUtil_GetInstance() *PathUtil {
        return PathUtil_Instance
}


func NewPathUtil() *PathUtil {
    this := &PathUtil{}
    this.initFields()
    return this
}


func (this *PathUtil) GetExtensionWithAbFilePath(abFilePath *AbFilePath) string {
        return this.AbPathData.GetExtension(fmt.Sprint(abFilePath))
}


func (this *PathUtil) GetExtensionWithAbPath(abPath IAbPath) string {
        return this.AbPathData.GetExtension(fmt.Sprint(abPath))
}


// throws Exception
func (this *PathUtil) GetWithoutExtensionWithAbFilePath(abFilePath *AbFilePath) string {
        return this.GetWithoutExtension(fmt.Sprint(abFilePath))
}


// throws Exception
func (this *PathUtil) GetWithoutExtensionWithAbPath(abPath IAbPath) string {
        return this.GetWithoutExtension(fmt.Sprint(abPath))
}


// throws Exception
func (this *PathUtil) GetWithoutExtension(filePath string) string {
        var indexOfFileExtensionDelmiter int = this.AbPathData.GetExtensionIndex(filePath); _ = indexOfFileExtensionDelmiter
        var pathWithoutExtension string = filePath; _ = pathWithoutExtension
        if indexOfFileExtensionDelmiter >= 0 {
            pathWithoutExtension = filePath[0:indexOfFileExtensionDelmiter]
        }
        return pathWithoutExtension
}


func (this *PathUtil) IsValid(path string) bool {
        if this.IsValidStart(path) {
            if this.IsValidEnd(path) {
                return true
            }
        }
        return false
}


func (this *PathUtil) IsValidStart(path string) bool {
        if rune(path[0]) == this.AbPathData.SEPARATORCHAR || rune(path[0]) == '.' {
            return true
        } else {
            return false
        }
}


func (this *PathUtil) IsValidEnd(path string) bool {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(path) {
            return false
        }
        if rune(path[len(path) - 1]) == this.AbPathData.SEPARATORCHAR {
            return true
        } else {
            return false
        }
}


func (this *PathUtil) AdjustStart(path string) string {
        if !this.IsValidStart(path) {
            path = fmt.Sprint(this.AbPathData.SEPARATOR) + path
        }
        return path
}


func (this *PathUtil) AdjustEnd(path string) string {
        if !this.IsValidEnd(path) {
            path = path + fmt.Sprint(this.AbPathData.SEPARATOR)
        }
        return path
}


func (this *PathUtil) Adjust(path string) string {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(path) {
            return this.AbPathData.SEPARATOR
        }
        if strings.Index(path, string("\\")) >= 0 {
            path = strings.ReplaceAll(path, string('\\'), string(this.AbPathData.SEPARATORCHAR))
        }
        return path
}


// throws Exception
func (this *PathUtil) RemoveNameFromPath(path string) IAbPath {
        var endIndex int = strings.LastIndex(path, string(this.AbPathData.SEPARATOR)); _ = endIndex
        if endIndex < 0 {
            endIndex = strings.LastIndex(path, string(this.FilePathData.SEPARATORCHAR))
        }
        if endIndex < 0 {
            return AbPath_CreateAbPath()
        }
        if len(path) == endIndex + 1 {
            return this.RemoveNameFromPath(path[0:endIndex - 1])
        } else {
            var pathWithoutName string = path[0:endIndex]; _ = pathWithoutName
            return NewAbPath(pathWithoutName, StringUtil_GetInstance().EMPTY_STRING)
        }
}

