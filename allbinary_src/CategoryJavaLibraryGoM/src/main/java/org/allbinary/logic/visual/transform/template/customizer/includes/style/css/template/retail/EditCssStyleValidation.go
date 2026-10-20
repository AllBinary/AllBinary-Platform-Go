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

type EditCssStyleValidation struct {
    CssStyleValidation
}

// Implements ValidationInterface, DomNodeInterface





// Java overloads of new EditCssStyleValidation
func NewEditCssStyleValidation(args ...any) *EditCssStyleValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewEditCssStyleValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return NewEditCssStyleValidation_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewEditCssStyleValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new EditCssStyleValidation for the arguments")
}
func NewEditCssStyleValidation_1() *EditCssStyleValidation {
    this := &EditCssStyleValidation{}
    this.CssStyleValidation = *NewCssStyleValidation()
    return this
}


func NewEditCssStyleValidation_2(document *Document) *EditCssStyleValidation {
    this := &EditCssStyleValidation{}
        this.CssStyleValidation = *NewCssStyleValidation(document)
    return this
}


func NewEditCssStyleValidation_3(hashMap *HashMap) *EditCssStyleValidation {
    this := &EditCssStyleValidation{}
        this.CssStyleValidation = *NewCssStyleValidation(hashMap)
    return this
}


func (this *EditCssStyleValidation) IsValid() bool {
        return this.CssStyleValidation.IsValid()
}


func (this *EditCssStyleValidation) ValidationInfo() string {
        return this.CssStyleValidation.ValidationInfo()
}


func (this *EditCssStyleValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditCssStyleValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *EditCssStyleValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(StylesData_GetInstance().NAME); _ = node
        var styleNode *Node = document.CreateElement(StyleData_GetInstance().NAME); _ = styleNode
        node.AppendChild(styleNode)
        styleNode.AppendChild(NewColorsView().ToXmlNode(document))
        var size int = this.CssStyleElementVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var styleDomNodeInterface DomNodeInterface = any(this.CssStyleElementVector.Get(i)).(DomNodeInterface); _ = styleDomNodeInterface
            styleNode.AppendChild(styleDomNodeInterface.ToXmlNode(document))
        }
        return node
}

