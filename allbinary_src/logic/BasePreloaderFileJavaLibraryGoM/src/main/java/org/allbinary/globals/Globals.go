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

type Globals struct {
    LogUtil *LogUtil
}

var Globals_Instance *Globals = NewGlobals()

func (this *Globals) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func Globals_GetInstance() *Globals {
        return Globals_Instance
}


func NewGlobals() *Globals {
    this := &Globals{}
    this.initFields()
    return this
}


func (this *Globals) Init(classLoader *ClassLoader, pathString string) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var appUrlGlobals *AppUrlGlobals = NewAppUrlGlobals(); _ = appUrlGlobals
        appUrlGlobals.SetWebappPath(pathString)
        URLGLOBALS.Init(appUrlGlobals)
        WebappClassLoaderInfo_SetLoader(classLoader)
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Setting Up AllBinary System Configuration Args: ")
        stringBuffer.Append(pathString)
        stringBuffer.Append(" ClassLoader: ")
        stringBuffer.Append(reflect.TypeOf(classLoader).String())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.INIT)
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(pathString) {
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("No Path Provided. Using Process Path: ")
            stringBuffer.Append(pathString)
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            stringBuffer.Append(CommonSeps_GetInstance().EQUALS)
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.INIT)
        }
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append("Webapp Path Set To: ")
        stringBuffer.Append(URLGLOBALS.GetWebappPath())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.INIT)
}

