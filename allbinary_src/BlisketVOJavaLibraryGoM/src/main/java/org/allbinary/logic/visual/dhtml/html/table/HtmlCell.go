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

type HtmlCell struct {
    HtmlTag
    END string
    START string
    ENDTAG string
    HEIGHT string
    WIDTH string
    Height string
    Width string
    Before string
    Body string
    After string
}


func (this *HtmlCell) initFields() {
    this.END = ">"
    this.START = "<td "
    this.ENDTAG = "</td>"
    this.HEIGHT = "height"
    this.WIDTH = "width"
    
    
    
    
    
}


func NewHtmlCell(before string, width string, height string, after string) *HtmlCell {
    this := &HtmlCell{}
    this.HtmlTag = *NewHtmlTag()
    this.initFields()
        this.Before = before
        this.After = after
        this.Height = height
        this.Width = width
        this.AddAttribute(this.HEIGHT, this.Height)
        this.AddAttribute(this.WIDTH, this.Width)
    return this
}


func (this *HtmlCell) AddText(body string) {
        this.Body = body
}


func (this *HtmlCell) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var attributeKeys []any = this.OtherAttributes.KeySet().ToArray(); _ = attributeKeys
        var attributeSize int = len(attributeKeys); _ = attributeSize
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
        result += this.Body
        result += this.ENDTAG
        result += this.After
        return result
}

func (this *HtmlCell) String() string {
    return this.ToString()
}

