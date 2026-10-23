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

// The Java class CssPropertyView has subclasses so its variables are this interface
type ICssPropertyView interface {
    CssPropertyView_() *CssPropertyView
    GetProperty() *CssProperty
    ToXmlNode(document *Document) *Node
}

func (this *CssPropertyView) CssPropertyView_() *CssPropertyView {
    return this
}


func (this *CssPropertyView) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new CssPropertyView
func NewCssPropertyView(args ...any) *CssPropertyView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewCssPropertyView_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCssPropertyView_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new CssPropertyView for the arguments")
}
func NewCssPropertyView_1(name string) *CssPropertyView {
    this := &CssPropertyView{}
    this.initFields()
        this.CssProperty = NewCssProperty(name)
    return this
}


func NewCssPropertyView_2(node *Node) *CssPropertyView {
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

