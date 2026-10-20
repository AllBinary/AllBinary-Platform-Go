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

type FREEBLISKET_PATH_GLOBALS struct {
    INSTALLPATH string
    XSLPATH string
    TEMPLATEPATH string
    STYLEPATH string
    THEMEPATH string
    DBINITPATH string
    LINEDATAPATH string
    CONFIGURATIONPATH string
    USERCONFIGURATIONPATH string
    CONTEXTCONFIGURATIONPATH string
    LICENSEERRORPAGE string
    ERRORPAGE string
}

var FREEBLISKET_PATH_GLOBALS_Instance *FREEBLISKET_PATH_GLOBALS = NewFREEBLISKET_PATH_GLOBALS()

func (this *FREEBLISKET_PATH_GLOBALS) initFields() {
    
    
    this.TEMPLATEPATH = "template" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)
    this.STYLEPATH = this.TEMPLATEPATH + "style" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)
    this.THEMEPATH = this.STYLEPATH + "theme" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)
    this.DBINITPATH = fmt.Sprint(PATH_GLOBALS.GetInstance().INIT_PATH) + "db" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)
    this.LINEDATAPATH = fmt.Sprint(PATH_GLOBALS.GetInstance().INIT_PATH) + "lineData" + fmt.Sprint(AbPathData_GetInstance().SEPARATOR)
    
    
    
    this.LICENSEERRORPAGE = "LicensingError.jsp"
    this.ERRORPAGE = "Error.jsp"
}


func FREEBLISKET_PATH_GLOBALS_GetInstance() *FREEBLISKET_PATH_GLOBALS {
        return FREEBLISKET_PATH_GLOBALS_Instance
}


func NewFREEBLISKET_PATH_GLOBALS() *FREEBLISKET_PATH_GLOBALS {
    this := &FREEBLISKET_PATH_GLOBALS{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("install")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.INSTALLPATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(PATH_GLOBALS.GetInstance().INIT_PATH)
        stringBuffer.Append("views")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.XSLPATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.XSLPATH)
        stringBuffer.Append("configuration")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.CONFIGURATIONPATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.CONFIGURATIONPATH)
        stringBuffer.Append("user")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.USERCONFIGURATIONPATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(this.CONFIGURATIONPATH)
        stringBuffer.Append("context")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.CONTEXTCONFIGURATIONPATH = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
    return this
}

