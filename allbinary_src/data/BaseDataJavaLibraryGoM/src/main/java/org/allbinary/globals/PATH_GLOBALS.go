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

type PATH_GLOBALS struct {
    DATA_PATH string
    INIT_PATH string
    VIEWS_PATH string
    LOG_PATH string
    BACKUP_PATH string
}

var PATH_GLOBALS_Instance *PATH_GLOBALS = NewPATH_GLOBALS()

func (this *PATH_GLOBALS) initFields() {
    
    
    
    
    
}


func PATH_GLOBALS_GetInstance() *PATH_GLOBALS {
        return PATH_GLOBALS_Instance
}


func NewPATH_GLOBALS() *PATH_GLOBALS {
    this := &PATH_GLOBALS{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var abPathData *AbPathData = AbPathData_GetInstance(); _ = abPathData
        stringBuffer.Append("data")
        stringBuffer.Append(abPathData.SEPARATOR)
        this.DATA_PATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.DATA_PATH)
        stringBuffer.Append(CommonStrings_GetInstance().INIT)
        stringBuffer.Append(abPathData.SEPARATOR)
        this.INIT_PATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.DATA_PATH)
        stringBuffer.Append("log")
        stringBuffer.Append(abPathData.SEPARATOR)
        this.LOG_PATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.DATA_PATH)
        stringBuffer.Append("backup")
        stringBuffer.Append(abPathData.SEPARATOR)
        this.BACKUP_PATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.INIT_PATH)
        stringBuffer.Append("views")
        stringBuffer.Append(abPathData.SEPARATOR)
        this.VIEWS_PATH = fmt.Sprint(stringBuffer)
    return this
}
