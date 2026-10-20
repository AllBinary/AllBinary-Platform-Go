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

type FormTypeFactory struct {
    UNK string
    NULL_FORM_TYPE *FormType
    TEMP_HORIZONTAL_FORM *FormType
    HORIZONTAL_FORM *FormType
    VERTICAL_CENTER_FORM *FormType
}

var FormTypeFactory_Instance *FormTypeFactory = NewFormTypeFactory()

func (this *FormTypeFactory) initFields() {
    this.UNK = "Unknown FormType"
    this.NULL_FORM_TYPE = NewFormType(StringUtil_GetInstance().NULL_STRING)
    this.TEMP_HORIZONTAL_FORM = NewFormType("TempHorizontal")
    this.HORIZONTAL_FORM = NewFormType("Horizontal")
    this.VERTICAL_CENTER_FORM = NewFormType("Vertical")
}

func NewFormTypeFactory() *FormTypeFactory {
    this := &FormTypeFactory{}
    this.initFields()
    return this
}


func FormTypeFactory_GetInstance() *FormTypeFactory {
        return FormTypeFactory_Instance
}


func (this *FormTypeFactory) GetFormType() *FormType {
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var formType *FormType; _ = formType
        if displayInfo.IsPortrait() {
            formType = this.VERTICAL_CENTER_FORM
        } else {
            formType = this.HORIZONTAL_FORM
        }
        return formType
}

