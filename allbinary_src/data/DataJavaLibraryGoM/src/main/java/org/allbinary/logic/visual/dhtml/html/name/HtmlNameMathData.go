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

type HtmlNameMathData struct {
    EQUALS string
    PLUS string
    MINUS string
    DIVIDE string
    MULTIPLY string
}

var HtmlNameMathData_Instance *HtmlNameMathData = NewHtmlNameMathData()

func (this *HtmlNameMathData) initFields() {
    this.EQUALS = "_EQUALS_"
    this.PLUS = "_PLUS_"
    this.MINUS = "_MINUS_"
    this.DIVIDE = "_DIVIDE_"
    this.MULTIPLY = "_MULTIPLY_"
}


func HtmlNameMathData_GetInstance() *HtmlNameMathData {
        return HtmlNameMathData_Instance
}


func NewHtmlNameMathData() *HtmlNameMathData {
    this := &HtmlNameMathData{}
    this.initFields()
    return this
}

