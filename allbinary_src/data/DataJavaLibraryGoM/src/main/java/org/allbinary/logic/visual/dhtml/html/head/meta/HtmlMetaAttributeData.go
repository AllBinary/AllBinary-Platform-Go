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

type HtmlMetaAttributeData struct {
    Id int
    Name string
}


func (this *HtmlMetaAttributeData) initFields() {
    
    
}


func NewHtmlMetaAttributeData(id int, value string) *HtmlMetaAttributeData {
    this := &HtmlMetaAttributeData{}
    this.initFields()
        this.Name = value
        this.Id = id
    return this
}


func (this *HtmlMetaAttributeData) Equals(htmlMetaAttributeData *HtmlMetaAttributeData) bool {
        if htmlMetaAttributeData.Id == this.Id {
            return true
        } else {
            return false
        }
}


func (this *HtmlMetaAttributeData) ToString() string {
        return this.Name
}

func (this *HtmlMetaAttributeData) String() string {
    return this.ToString()
}

