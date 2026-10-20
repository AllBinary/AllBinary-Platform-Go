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

type HtmlTags struct {
    HtmlTagVector IBasicArrayList
}


func (this *HtmlTags) initFields() {
    
}


func NewHtmlTags() *HtmlTags {
    this := &HtmlTags{}
    this.initFields()
        this.HtmlTagVector = NewBasicArrayListD()
    return this
}


func (this *HtmlTags) Get(index int) IHtmlTag {
        return any(this.HtmlTagVector.Get(index)).(IHtmlTag)
}


func (this *HtmlTags) Add(htmlTag IHtmlTag) {
        this.HtmlTagVector.Add(htmlTag)
}


func (this *HtmlTags) Size() int {
        return this.HtmlTagVector.Size()
}

