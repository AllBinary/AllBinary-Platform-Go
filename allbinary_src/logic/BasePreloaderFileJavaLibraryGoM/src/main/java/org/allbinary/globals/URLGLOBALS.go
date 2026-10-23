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

type URLGLOBALS struct {
}

var URLGLOBALS_UrlGlobalsInterface IUrlGlobalsInterface = AppUrlGlobals_NULL_APP_URL_GLOBALS

func NewURLGLOBALS() *URLGLOBALS {
    this := &URLGLOBALS{}
    return this
}


func URLGLOBALS_Init(urlGlobalsInterface IUrlGlobalsInterface) {
        URLGLOBALS_UrlGlobalsInterface = urlGlobalsInterface
}


func URLGLOBALS_IsTestingMode() bool {
        return URLGLOBALS_UrlGlobalsInterface.IsTestingMode()
}


func URLGLOBALS_GetWebappPath() string {
        return URLGLOBALS_UrlGlobalsInterface.GetWebappPath()
}


func URLGLOBALS_GetTestHtmlPath() string {
        return URLGLOBALS_UrlGlobalsInterface.GetTestHtmlPath()
}


// throws Exception
func URLGLOBALS_GetMainPath() string {
        return URLGLOBALS_UrlGlobalsInterface.GetMainPath()
}

