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

type TimeLastModifiedView struct {
    TimeLastModified *TimeLastModified
}

// Implements DomNodeInterface


func (this *TimeLastModifiedView) initFields() {
    
}


func NewTimeLastModifiedView(timeLastModified *TimeLastModified) *TimeLastModifiedView {
    this := &TimeLastModifiedView{}
    this.initFields()
        this.SetTimeLastModified(timeLastModified)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTimeLastModifiedView(node *Node) *TimeLastModifiedView {
    this := &TimeLastModifiedView{}
    this.initFields()
        this.SetTimeLastModified(NewTimeLastModified(DomNodeHelper_GetTextNodeValue(node)))
    return this
}


// throws Exception
func (this *TimeLastModifiedView) ToXmlNode(document *Document) *Node {
        return ModDomHelper_CreateTextNode(document, EntryData_GetInstance().TIMECREATED, fmt.Sprint(this.GetTimeLastModified()))
}


func (this *TimeLastModifiedView) GetTimeLastModified() *TimeLastModified {
        return this.TimeLastModified
}


func (this *TimeLastModifiedView) SetTimeLastModified(timeLastModified *TimeLastModified) {
        this.TimeLastModified = timeLastModified
}

