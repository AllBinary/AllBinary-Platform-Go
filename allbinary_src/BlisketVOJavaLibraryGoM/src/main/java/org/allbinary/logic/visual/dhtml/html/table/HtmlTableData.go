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

type HtmlTableData struct {
    LEFT string
    RIGHT string
    NONE string
    TOP string
    BOTTOM string
    TOPBOT string
    SIDES string
    ALL string
    BORDERED string
    BASIC string
    ROWS string
    FRAMES []string
}

var HtmlTableData_Instance *HtmlTableData = NewHtmlTableData()

func (this *HtmlTableData) initFields() {
    this.LEFT = CommonStrings_GetInstance().LEFT
    this.RIGHT = CommonStrings_GetInstance().RIGHT
    this.NONE = "none"
    this.TOP = "top"
    this.BOTTOM = "bottom"
    this.TOPBOT = "topbot"
    this.SIDES = "sides"
    this.ALL = "all"
    this.BORDERED = "border"
    this.BASIC = "basic"
    this.ROWS = "rows"
    this.FRAMES = []string{this.NONE, this.TOP, this.BOTTOM, this.TOPBOT, this.SIDES, this.ALL, this.BORDERED}
}

func NewHtmlTableData() *HtmlTableData {
    this := &HtmlTableData{}
    this.initFields()
    return this
}


func HtmlTableData_GetInstance() *HtmlTableData {
        return HtmlTableData_Instance
}

