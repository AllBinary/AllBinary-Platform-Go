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

type HtmlSelect struct {
    HtmlTag
    END string
    STARTSELECT string
    ENDSELECT string
    SIZE string
    NAME string
    STARTOPTION string
    SELETED string
    ENDOPTION string
    Before string
    Name string
    After string
    Size string
    Selected string
    Options IBasicArrayList
    Multiple string
    IsMultipleSelect bool
}


func (this *HtmlSelect) initFields() {
    this.END = ">"
    this.STARTSELECT = "<SELECT "
    this.ENDSELECT = "</SELECT>"
    this.SIZE = "SIZE=\""
    this.NAME = "NAME=\""
    this.STARTOPTION = "<OPTION"
    this.SELETED = "SELECTED"
    this.ENDOPTION = "</OPTION>"
    
    
    
    
    
    
    this.Multiple = "multiple"
    this.IsMultipleSelect = false
}


func NewHtmlSelect(before string, size string, name string, after string) *HtmlSelect {
    this := &HtmlSelect{}
    this.HtmlTag = *NewHtmlTag()
    this.initFields()
        this.Before = before
        this.Size = size
        this.Name = name
        this.After = after
        this.Options = NewBasicArrayListD()
    return this
}


func (this *HtmlSelect) AddOption(key string) {
        this.Options.Add(key)
}


func (this *HtmlSelect) AddOptions(fileName string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.AddOption("Error")
                }
            }()
            var lineReader *LineReader = NewLineReader(fileName); _ = lineReader
            for lineReader.HasNext() {
                var option string = lineReader.Next(); _ = option
                this.AddOption(option)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *HtmlSelect) SetSelected(selected string) {
        this.Selected = selected
}


func (this *HtmlSelect) IsMultiple() bool {
        return this.IsMultipleSelect
}


func (this *HtmlSelect) SetMultiple(value bool) {
        this.IsMultipleSelect = value
}


func (this *HtmlSelect) GetOptions() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var optionsArray []any = this.Options.ToArray(); _ = optionsArray
        var optionsSize int = len(optionsArray); _ = optionsSize
        for i := int(0); i < optionsSize; i++ {
            var value string = any(optionsArray[i]).(string); _ = value
            if value != nil && value.CompareTo(stringUtil.EMPTY_STRING) != 0 {
                result += this.STARTOPTION
                result += this.END
                result += value
                result += this.ENDOPTION
            }
        }
        return result
}


func (this *HtmlSelect) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var attributeKeys []any = this.OtherAttributes.KeySet().ToArray(); _ = attributeKeys
        var attributeSize int = len(attributeKeys); _ = attributeSize
        result = this.Before
        result += this.STARTSELECT
        result += this.NAME
        result += this.Name
        result += "\" "
        result += this.SIZE
        result += this.Size
        result += "\" "
        for i := int(0); i < attributeSize; i++ {
            var key string = any(attributeKeys[i]).(string); _ = key
            var value string = any(this.OtherAttributes.Get(key)).(string); _ = value
            if value != nil && value.CompareTo(stringUtil.EMPTY_STRING) != 0 {
                result += key
                result += "=\""
                result += value
                result += "\" "
            }
        }
        if this.IsMultiple() {
            result += " " + this.Multiple + " "
        }
        result += this.END
        result += this.GetOptions()
        result += this.ENDSELECT
        result += this.After
        return result
}

func (this *HtmlSelect) String() string {
    return this.ToString()
}

