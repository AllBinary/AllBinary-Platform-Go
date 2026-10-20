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

type ThemeData struct {
    NAME string
    PATH string
    PREVIEW_IMAGE_NAME string
    PREVIEW_IMAGE_PATH string
}

var ThemeData_Instance *ThemeData = NewThemeData()

func (this *ThemeData) initFields() {
    this.NAME = "THEME_NAME"
    this.PATH = "THEME_PATH"
    this.PREVIEW_IMAGE_NAME = "THEME_PREVIEW_IMAGE_NAME"
    this.PREVIEW_IMAGE_PATH = "THEME_PREVIEW_IMAGE_PATH"
}


func ThemeData_GetInstance() *ThemeData {
        return ThemeData_Instance
}


func NewThemeData() *ThemeData {
    this := &ThemeData{}
    this.initFields()
    return this
}

