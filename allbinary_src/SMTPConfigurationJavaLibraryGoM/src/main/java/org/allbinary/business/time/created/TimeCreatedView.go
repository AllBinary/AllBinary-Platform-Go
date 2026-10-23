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

type TimeCreatedView struct {
    TimeCreated *TimeCreated
}

// Implements DomNodeInterface


func (this *TimeCreatedView) initFields() {
    
}




// Java overloads of new TimeCreatedView
func NewTimeCreatedView(args ...any) *TimeCreatedView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*TimeCreated](args[0], pass) {
            return NewTimeCreatedView_1(Java_Arg[*TimeCreated](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewTimeCreatedView_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new TimeCreatedView for the arguments")
}
func NewTimeCreatedView_1(timeCreated *TimeCreated) *TimeCreatedView {
    this := &TimeCreatedView{}
    this.initFields()
        this.SetTimeCreated(timeCreated)
    return this
}


func NewTimeCreatedView_2(node *Node) *TimeCreatedView {
    this := &TimeCreatedView{}
    this.initFields()
        this.SetTimeCreated(NewTimeCreated(DomNodeHelper_GetTextNodeValue(node)))
    return this
}


// throws Exception
func (this *TimeCreatedView) ToXmlNode(document *Document) *Node {
        return ModDomHelper_CreateTextNode(document, EntryData_GetInstance().TIMECREATED, fmt.Sprint(this.GetTimeCreated()))
}


func (this *TimeCreatedView) GetTimeCreated() *TimeCreated {
        return this.TimeCreated
}


func (this *TimeCreatedView) SetTimeCreated(timeCreated *TimeCreated) {
        this.TimeCreated = timeCreated
}

