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

type HtmlTable struct {
    HtmlTag
    END string
    START string
    ENDTAG string
    CLASS string
    DIRECTION string
    ID string
    LANG string
    UNITS string
    BORDER string
    CELLPADDING string
    CELLSPACING string
    COLS string
    FLOAT string
    FRAME string
    RULES string
    STYLE string
    BORDERCOLLAPSE string
    NumberOfColumns int
    Before string
    Border string
    Cellpadding string
    Cellspacing string
    Style string
    Name string
    After string
    HtmlRowsVector IBasicArrayList
}


func (this *HtmlTable) initFields() {
    this.END = ">"
    this.START = "<table "
    this.ENDTAG = "</table>"
    this.CLASS = "class"
    this.DIRECTION = "dir"
    this.ID = "id"
    this.LANG = "lang"
    this.UNITS = "units"
    this.BORDER = "border"
    this.CELLPADDING = "cellpadding"
    this.CELLSPACING = "cellspacing"
    this.COLS = "cols"
    this.FLOAT = "float"
    this.FRAME = "frame"
    this.RULES = "rules"
    this.STYLE = "style"
    this.BORDERCOLLAPSE = "border-collapse: collapse"
    this.NumberOfColumns = 0
    
    
    
    
    
    
    
    
}


func NewHtmlTable(before string, cellpadding string, cellspacing string, after string) *HtmlTable {
    this := &HtmlTable{}
    this.HtmlTag = *NewHtmlTag()
    this.initFields()
        this.Before = before
        this.Cellpadding = cellpadding
        this.Cellspacing = cellspacing
        this.After = after
        this.HtmlRowsVector = NewBasicArrayListD()
        this.AddAttribute(this.BORDER, CommonPhoneStrings_GetInstance().ZERO)
        this.AddAttribute(this.CELLPADDING, this.Cellpadding)
        this.AddAttribute(this.CELLSPACING, this.Cellspacing)
        this.AddAttribute(this.STYLE, this.BORDERCOLLAPSE)
    return this
}


func (this *HtmlTable) AddRow(htmlRow *HtmlRow) {
        if this.NumberOfColumns < htmlRow.GetNumberOfColumns() {
            this.NumberOfColumns = htmlRow.GetNumberOfColumns()
        }
        this.HtmlRowsVector.Add(htmlRow)
}


func (this *HtmlTable) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var attributeKeys []any = this.OtherAttributes.KeySet().ToArray(); _ = attributeKeys
        result = this.Before
        result += this.START
        result += " "
        for i := int(0); i < len(attributeKeys); i++ {
            var key string = any(attributeKeys[i]).(string); _ = key
            var value string = any(this.OtherAttributes.Get(key)).(string); _ = value
            if value != "" && strings.Compare(value, stringUtil.EMPTY_STRING) != 0 {
                result += key
                result += "=\""
                result += value
                result += "\" "
            }
        }
        result += this.END
        var size int = this.HtmlRowsVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            result += fmt.Sprint(this.HtmlRowsVector.Get(i))
            result += " "
        }
        result += this.ENDTAG
        result += this.After
        return result
}

func (this *HtmlTable) String() string {
    return this.ToString()
}

