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

type CssPropertyView struct {
    CommonStrings *CommonStrings
    CssProperty *CssProperty
}

// Implements DomNodeInterface


func (this *CssPropertyView) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewCssPropertyView(name string) *CssPropertyView {
    this := &CssPropertyView{}
    this.initFields()
        this.CssProperty = NewCssProperty(name)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewCssPropertyView(node *Node) *CssPropertyView {
    this := &CssPropertyView{}
    this.initFields()
        this.CssProperty = NewCssProperty(node)
    return this
}


func (this *CssPropertyView) GetProperty() *CssProperty {
        return this.CssProperty
}


// throws Exception
func (this *CssPropertyView) ToXmlNode(document *Document) *Node {
        return this.CssProperty.ToXmlNode(document)
}

