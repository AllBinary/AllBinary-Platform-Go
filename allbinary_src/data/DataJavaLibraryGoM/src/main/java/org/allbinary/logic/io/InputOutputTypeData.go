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

type InputOutputTypeData struct {
    NAME string
    FILE string
    RESPONSE string
    DB string
    JSP string
    JSP_FRAGMENT string
    PHP string
    ASP string
    PL string
    DEFAULT string
    DEFAULT_FRAGMENT string
}

var InputOutputTypeData_Instance *InputOutputTypeData = NewInputOutputTypeData()

func (this *InputOutputTypeData) initFields() {
    this.NAME = "INPUT_OUTPUT_TYPE_NAME"
    this.FILE = "INPUT_OUTPUT_TYPE_FILE"
    this.RESPONSE = "INPUT_OUTPUT_TYPE_RESPONSE"
    this.DB = "INPUT_OUTPUT_TYPE_DB"
    this.JSP = "jsp"
    this.JSP_FRAGMENT = "jspf"
    this.PHP = "php"
    this.ASP = "asp"
    this.PL = "pl"
    this.DEFAULT = this.JSP
    this.DEFAULT_FRAGMENT = this.JSP_FRAGMENT
}


func InputOutputTypeData_GetInstance() *InputOutputTypeData {
        return InputOutputTypeData_Instance
}


func NewInputOutputTypeData() *InputOutputTypeData {
    this := &InputOutputTypeData{}
    this.initFields()
    return this
}

