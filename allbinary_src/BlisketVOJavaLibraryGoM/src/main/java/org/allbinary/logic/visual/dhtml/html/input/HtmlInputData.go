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

type HtmlInputData struct {
    TEXT string
    RADIO string
    CHECKBOX string
    HIDDEN string
    IMAGE string
    PASSWORD string
    RESET string
    SUBMIT string
    TYPES []string
}

var HtmlInputData_Instance *HtmlInputData = NewHtmlInputData()

func (this *HtmlInputData) initFields() {
    this.TEXT = "text"
    this.RADIO = "radio"
    this.CHECKBOX = "checkbox"
    this.HIDDEN = "hidden"
    this.IMAGE = "image"
    this.PASSWORD = "password"
    this.RESET = "reset"
    this.SUBMIT = "submit"
    this.TYPES = []string{this.TEXT, this.RADIO, this.CHECKBOX, this.HIDDEN, this.IMAGE, this.PASSWORD, this.RADIO, this.RESET, this.SUBMIT}
}

func NewHtmlInputData() *HtmlInputData {
    this := &HtmlInputData{}
    this.initFields()
    return this
}


func HtmlInputData_GetInstance() *HtmlInputData {
        return HtmlInputData_Instance
}

