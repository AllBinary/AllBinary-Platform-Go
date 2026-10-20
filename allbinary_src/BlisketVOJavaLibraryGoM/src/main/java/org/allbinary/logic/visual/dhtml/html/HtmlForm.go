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

type HtmlForm struct {
    FORMBEGIN string
    END string
    FORMEND string
    ACTION string
    GET string
    POST string
    METHODS []string
    METHOD string
    Method string
    Action string
    Inputs IBasicArrayList
}

// The Java class HtmlForm has subclasses so its variables are this interface
type IHtmlForm interface {
    HtmlForm_() *HtmlForm
    SetPost()
    SetGet()
    AddInput(htmlInput IHtmlInput)
    AddTable(htmlTable *HtmlTable)
    ToString() string
}

func (this *HtmlForm) HtmlForm_() *HtmlForm {
    return this
}


func (this *HtmlForm) initFields() {
    this.FORMBEGIN = "<form "
    this.END = ">"
    this.FORMEND = "</form>"
    this.ACTION = "ACTION=\""
    this.GET = "GET\""
    this.POST = "POST\""
    this.METHODS = []string{this.GET, this.POST}
    this.METHOD = "METHOD=\""
    
    
    
}


func NewHtmlForm(action string) *HtmlForm {
    this := &HtmlForm{}
    this.initFields()
        this.Action = action
        this.Inputs = NewBasicArrayListD()
    return this
}


func (this *HtmlForm) SetPost() {
        this.Method = this.POST
}


func (this *HtmlForm) SetGet() {
        this.Method = this.GET
}


func (this *HtmlForm) AddInput(htmlInput IHtmlInput) {
        this.Inputs.Add(htmlInput)
}


func (this *HtmlForm) AddTable(htmlTable *HtmlTable) {
        this.Inputs.Add(htmlTable)
}


func (this *HtmlForm) ToString() string {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var result string = stringUtil.EMPTY_STRING; _ = result
        var inputArray []any = this.Inputs.ToArray(); _ = inputArray
        var inputSize int = len(inputArray); _ = inputSize
        result = this.FORMBEGIN
        result += this.METHOD
        result += this.Method
        result += "\" "
        result += this.ACTION
        result += this.Action
        result += "\" "
        result += this.END
        for i := int(0); i < inputSize; i++ {
            result += fmt.Sprint(inputArray[i])
            result += " "
        }
        result += this.FORMEND
        return result
}

func (this *HtmlForm) String() string {
    return this.ToString()
}

