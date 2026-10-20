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

type HtmlRadioInput struct {
    HtmlValueInput
    CHECKED string
}


func (this *HtmlRadioInput) initFields() {
    this.CHECKED = "checked"
}


func NewHtmlRadioInput(before string, name string, value string, after string) *HtmlRadioInput {
    this := &HtmlRadioInput{}
        this.HtmlValueInput = *NewHtmlValueInput(before, HtmlInputData_GetInstance().RADIO, name, value, after)
        this.initFields()
    return this
}


func (this *HtmlRadioInput) SetSelected() {
        this.HtmlValueInput.AddAttribute(this.CHECKED, BooleanFactory_GetInstance().TRUE_STRING)
}

