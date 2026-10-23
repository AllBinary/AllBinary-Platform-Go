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




// Java overloads of new TimeLastModifiedView
func NewTimeLastModifiedView(args ...any) *TimeLastModifiedView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*TimeLastModified](args[0], pass) {
            return NewTimeLastModifiedView_1(Java_Arg[*TimeLastModified](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewTimeLastModifiedView_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new TimeLastModifiedView for the arguments")
}
func NewTimeLastModifiedView_1(timeLastModified *TimeLastModified) *TimeLastModifiedView {
    this := &TimeLastModifiedView{}
    this.initFields()
        this.SetTimeLastModified(timeLastModified)
    return this
}


func NewTimeLastModifiedView_2(node *Node) *TimeLastModifiedView {
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

