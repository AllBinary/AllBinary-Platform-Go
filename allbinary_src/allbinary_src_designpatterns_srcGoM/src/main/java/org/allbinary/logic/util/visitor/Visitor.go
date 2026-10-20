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

type Visitor struct {
    NullUtil *NullUtil
}

// Implements VisitorInterface

// The Java class Visitor has subclasses so its variables are this interface
type IVisitor interface {
    Visitor_() *Visitor
    Visit(anyType any) any
}

func (this *Visitor) Visitor_() *Visitor {
    return this
}

var Visitor_Instance IVisitor = NewVisitor()

func (this *Visitor) initFields() {
    this.NullUtil = NullUtil_GetInstance()
}

func NewVisitor() *Visitor {
    this := &Visitor{}
    this.initFields()
    return this
}


func (this *Visitor) Visit(anyType any) any {
        return this.NullUtil.NULL_OBJECT
}


func Visitor_GetInstance() IVisitor {
        return Visitor_Instance
}
