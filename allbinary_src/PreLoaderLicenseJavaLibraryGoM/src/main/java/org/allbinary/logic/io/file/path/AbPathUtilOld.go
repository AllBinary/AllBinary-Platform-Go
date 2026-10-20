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
type AbPathUtilOld struct {
    LogUtil *LogUtil
    AbPathData *AbPathData
    FilePathData *FilePathData
}

var AbPathUtilOld_Instance *AbPathUtilOld = NewAbPathUtilOld()
var AbPathUtilOld_MIN int = 4

func (this *AbPathUtilOld) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbPathData = AbPathData_GetInstance()
    this.FilePathData = FilePathData_GetInstance()
}


func AbPathUtilOld_GetInstance() *AbPathUtilOld {
        return AbPathUtilOld_Instance
}


func NewAbPathUtilOld() *AbPathUtilOld {
    this := &AbPathUtilOld{}
    this.initFields()
    return this
}


// throws Exception
func (this *AbPathUtilOld) GetExtension(abFilePath *AbFilePath) string {
        return this.GetExtension(fmt.Sprint(abFilePath))
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *AbPathUtilOld) GetExtension(abPath *AbPath) string {
        return this.GetExtension(fmt.Sprint(abPath))
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *AbPathUtilOld) GetExtension(filePath string) string {
        if len(filePath) < AbPathUtilOld_MIN {
            panic(errors.New(fmt.Sprint("Could not be a file path since its less that 4 characters")))
        }
        var beginIndex int = len(filePath) - AbPathUtilOld_MIN; _ = beginIndex
        var extension string = filePath[beginIndex + 1:]; _ = extension
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("FileExtension: " + extension, this, "getExtension()")
        }
        return extension
}


// throws Exception
func (this *AbPathUtilOld) GetWithoutExtension(abFilePath *AbFilePath) string {
        return this.GetWithoutExtension(fmt.Sprint(abFilePath))
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *AbPathUtilOld) GetWithoutExtension(abPath *AbPath) string {
        return this.GetWithoutExtension(fmt.Sprint(abPath))
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *AbPathUtilOld) GetWithoutExtension(filePath string) string {
        if len(filePath) < AbPathUtilOld_MIN {
            panic(errors.New(fmt.Sprint("Could not be a file path since its less that 4 characters")))
        }
        var endIndex int = len(filePath) - AbPathUtilOld_MIN; _ = endIndex
        var pathWithoutExtension string = filePath[0:endIndex]; _ = pathWithoutExtension
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("FileWithoutExtension: " + pathWithoutExtension, this, "getWithoutExtension()")
        }
        return pathWithoutExtension
}


func (this *AbPathUtilOld) IsValid(path string) bool {
        if this.IsValidStart(path) {
            if this.IsValidEnd(path) {
                return true
            }
        }
        return false
}


func (this *AbPathUtilOld) IsValidStart(path string) bool {
        if rune(path[0]) == this.AbPathData.SEPARATORCHAR {
            return true
        } else {
            return false
        }
}


func (this *AbPathUtilOld) IsValidEnd(path string) bool {
        if StringValidationUtil_GetInstance().IsEmpty(path) {
            return false
        }
        if rune(path[len(path) - 1]) == this.AbPathData.SEPARATORCHAR {
            return true
        } else {
            return false
        }
}


func (this *AbPathUtilOld) AdjustStart(path string) string {
        if !this.IsValidStart(path) {
            path = fmt.Sprint(this.AbPathData.SEPARATOR) + path
        }
        return path
}


func (this *AbPathUtilOld) AdjustEnd(path string) string {
        if !this.IsValidEnd(path) {
            path = path + fmt.Sprint(this.AbPathData.SEPARATOR)
        }
        return path
}


func (this *AbPathUtilOld) Adjust(path string) string {
        if StringValidationUtil_GetInstance().IsEmpty(path) {
            return this.AbPathData.SEPARATOR
        }
        if strings.Index(path, string("\\")) >= 0 {
            path = strings.ReplaceAll(path, string('\\'), string(this.AbPathData.SEPARATORCHAR))
        }
        return path
}


// synchronized
func (this *AbPathUtilOld) GetNameFromPath(categoryPath string) string {
        var endIndex int = strings.LastIndex(categoryPath, string(this.AbPathData.SEPARATOR)); _ = endIndex
        if endIndex < 0 {
            endIndex = strings.LastIndex(categoryPath, string(this.FilePathData.SEPARATORCHAR))
        }
        if endIndex < 0 {
            return categoryPath
        }
        if len(categoryPath) == endIndex + 1 {
            var categoryName string = categoryPath[0:endIndex]; _ = categoryName
            return this.GetNameFromPath(categoryName)
        } else {
            var categoryName string = categoryPath[endIndex + 1:]; _ = categoryName
            return categoryName
        }
}


// throws Exception
// synchronized
func (this *AbPathUtilOld) RemoveNameFromPath(categoryPath string) *AbPath {
        var endIndex int = strings.LastIndex(categoryPath, string(this.AbPathData.SEPARATOR)); _ = endIndex
        if endIndex < 0 {
            endIndex = strings.LastIndex(categoryPath, string(this.FilePathData.SEPARATORCHAR))
        }
        if endIndex < 0 {
            return AbPath_CreateAbPath()
        }
        if len(categoryPath) == endIndex + 1 {
            return this.RemoveNameFromPath(categoryPath[0:endIndex - 1])
        } else {
            var categoryName string = categoryPath[0:endIndex]; _ = categoryName
            return NewAbPath(categoryName, StringUtil_GetInstance().EMPTY_STRING)
        }
}

