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

type HtmlRow struct {
    HtmlTag
    END string
    START string
    ENDTAG string
    NumberOfColumns int
    Before string
    After string
    HtmlCellsVector IBasicArrayList
}


func (this *HtmlRow) initFields() {
    this.END = ">"
    this.START = "<tr "
    this.ENDTAG = "</tr>"
    this.NumberOfColumns = 0
    
    
    
}


func NewHtmlRow(before string, after string) *HtmlRow {
    this := &HtmlRow{}
    this.HtmlTag = *NewHtmlTag()
    this.initFields()
        this.Before = before
        this.After = after
        this.HtmlCellsVector = NewBasicArrayListD()
    return this
}


func (this *HtmlRow) AddCell(htmlCell *HtmlCell) {
        this.HtmlCellsVector.Add(htmlCell)
        this.NumberOfColumns++
}


func (this *HtmlRow) GetNumberOfColumns() int {
        return this.NumberOfColumns
}


func (this *HtmlRow) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var attributeKeys []any = this.OtherAttributes.KeySet().ToArray(); _ = attributeKeys
        var attributeSize int = len(attributeKeys); _ = attributeSize
        var cellSize int = this.HtmlCellsVector.Size(); _ = cellSize
        result = this.Before
        result += this.START
        result += " "
        for i := int(0); i < attributeSize; i++ {
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
        for i := int(0); i < cellSize; i++ {
            result += fmt.Sprint(this.HtmlCellsVector.Get(i))
            result += " "
        }
        result += this.ENDTAG
        result += this.After
        return result
}

func (this *HtmlRow) String() string {
    return this.ToString()
}

