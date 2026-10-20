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

type AbPath struct {
    CommonSeps *CommonSeps
    AbPathData *AbPathData
    AbPathUtil *PathUtil
    EMPTY_STRING string
    Schema string
    Path string
    NameP string
    HasSchemaP bool
    NumberOfSeps int
}

var AbPath_NETWORK_SEP string = ":/"

func (this *AbPath) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    this.AbPathData = AbPathData_GetInstance()
    this.AbPathUtil = PathUtil_GetInstance()
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    this.Schema = this.EMPTY_STRING
    this.Path = this.EMPTY_STRING
    this.NameP = this.EMPTY_STRING
    this.HasSchemaP = false
    this.NumberOfSeps = 0
}


func AbPath_CreateAbPath() *AbPath {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *AbPath) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    ex := Java_ToError(recovered)
                    _ = ex
            panic(errors.New(fmt.Sprint()))
                }
            }()
            var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
            tryControl1, tryResult1 = 1, NewAbPath(EMPTY_STRING, EMPTY_STRING); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func NewAbPath(aPath string, name string) *AbPath {
    this := &AbPath{}
    this.initFields()
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsEmpty(aPath) {
            this.Schema = this.GetSchema(aPath)
            this.NameP = name
            this.Path = this.AbPathUtil.AdjustEnd(this.AbPathUtil.Adjust(this.GetPathFromPath(aPath)))
        } else {
            this.Init()
            if !stringValidationUtil.IsEmpty(name) {
                this.NameP = name
            }
        }
    return this
}


func (this *AbPath) Init() {
        this.Schema = this.EMPTY_STRING
        this.SetPath(this.AbPathData.SEPARATOR)
        this.NameP = this.EMPTY_STRING
}


func (this *AbPath) GetSchema(aPath string) string {
        var beginIndex int = strings.Index(aPath, string(this.CommonSeps.COLON)); _ = beginIndex
        if beginIndex >= 0 {
            this.HasSchemaP = true
            return aPath[0:beginIndex]
        }
        this.HasSchemaP = false
        return this.EMPTY_STRING
}


func (this *AbPath) HasSchema() bool {
        return this.HasSchemaP
}


// throws Exception
func (this *AbPath) GetPathFromPath(aPath string) string {
        var filePathData *FilePathData = FilePathData_GetInstance(); _ = filePathData
        var tempPath string = aPath; _ = tempPath
        if !this.HasSchema() {
            tempPath = this.AbPathUtil.AdjustStart(tempPath)
        } else {
            var beginIndex int = strings.Index(tempPath, string(this.CommonSeps.COLON)); _ = beginIndex
            if beginIndex >= 0 {
                beginIndex++
                for rune(tempPath[beginIndex]) == this.AbPathData.SEPARATORCHAR || rune(tempPath[beginIndex]) == filePathData.SEPARATORCHAR {
                    beginIndex++
                    this.NumberOfSeps++
                    if this.NumberOfSeps > 2 {
                        panic(errors.New(fmt.Sprint("Should Not Have More Than Two Seps")))
                    }
                }
                tempPath = tempPath[beginIndex:len(tempPath)]
            }
        }
        return tempPath
}


func (this *AbPath) SetName(name string) {
        this.NameP = name
}


func (this *AbPath) GetName() string {
        return this.NameP
}


func (this *AbPath) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if this.HasSchema() {
            stringBuffer.Append(this.Schema)
            stringBuffer.Append(AbPath_NETWORK_SEP)
        }
        stringBuffer.Append(this.GetPath())
        stringBuffer.Append(this.NameP)
        return fmt.Sprint(stringBuffer)
}

func (this *AbPath) String() string {
    return this.ToString()
}


func (this *AbPath) ToFileSystemString() string {
        if this.HasSchema() {
            var filePathData *FilePathData = FilePathData_GetInstance(); _ = filePathData
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.Schema)
            stringBuffer.Append(this.CommonSeps.COLON)
            stringBuffer.Append(filePathData.PATH_START)
            stringBuffer.Append(FilePathUtil_Adjust(this.GetPath()))
            stringBuffer.Append(this.NameP)
            return fmt.Sprint(stringBuffer)
        } else {
            return fmt.Sprint(FilePathUtil_Adjust(this.GetPath())) + this.NameP
        }
}


func (this *AbPath) GetPath() string {
        return this.Path
}


func (this *AbPath) SetPath(path string) {
        this.Path = path
}

