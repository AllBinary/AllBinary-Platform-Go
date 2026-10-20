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

type HtmlInput struct {
    HtmlTag
    END string
    STARTINPUT string
    ENDINPUT string
    TYPE string
    NAME string
    Before string
    Type string
    Name string
    After string
}

// The Java class HtmlInput has subclasses so its variables are this interface
type IHtmlInput interface {
    IHtmlTag
    HtmlInput_() *HtmlInput
    ToString() string
}

func (this *HtmlInput) HtmlInput_() *HtmlInput {
    return this
}


func (this *HtmlInput) initFields() {
    this.END = ">"
    this.STARTINPUT = "<INPUT "
    this.ENDINPUT = "</INPUT>"
    this.TYPE = "TYPE=\""
    this.NAME = "NAME=\""
    
    
    
    
}


func NewHtmlInput(before string, type_ string, name string, after string) *HtmlInput {
    this := &HtmlInput{}
    this.HtmlTag = *NewHtmlTag()
    this.initFields()
        this.Before = before
        this.Type = type_
        this.Name = name
        this.After = after
    return this
}


func (this *HtmlInput) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var attributeKeys []any = this.OtherAttributes.KeySet().ToArray(); _ = attributeKeys
        var attributeSize int = len(attributeKeys); _ = attributeSize
        result = this.Before
        result += this.STARTINPUT
        result += this.TYPE
        result += this.Type
        result += "\" "
        result += this.NAME
        result += this.Name
        result += "\" "
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
        result += this.ENDINPUT
        result += this.After
        return result
}

func (this *HtmlInput) String() string {
    return this.ToString()
}

