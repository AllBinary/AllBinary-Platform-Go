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

type ThemesValidation struct {
    Validation
    ThemeVector IBasicArrayList
}

// Implements DomNodeInterface


func (this *ThemesValidation) initFields() {
    
}




// Java overloads of new ThemesValidation
func NewThemesValidation(args ...any) *ThemesValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewThemesValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewThemesValidation_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewThemesValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new ThemesValidation for the arguments")
}
func NewThemesValidation_1() *ThemesValidation {
    this := &ThemesValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.ThemeVector = NewBasicArrayListD()
    return this
}


func NewThemesValidation_2(node *Node) *ThemesValidation {
    this := &ThemesValidation{}
    this.Validation = *NewValidation()
    this.initFields()
    return this
}


func NewThemesValidation_3(hashMap *HashMap) *ThemesValidation {
    this := &ThemesValidation{}
    this.Validation = *NewValidation()
    this.initFields()
    return this
}


func (this *ThemesValidation) IsValid() bool {
        return true
}


func (this *ThemesValidation) ValidationInfo() string {
        return nil
}


func (this *ThemesValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *ThemesValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *ThemesValidation) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(ThemesData_GetInstance().NAME); _ = node
        var size int = this.ThemeVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var themeDomNodeInterface DomNodeInterface = any(this.ThemeVector.Get(i)).(DomNodeInterface); _ = themeDomNodeInterface
            node.AppendChild(themeDomNodeInterface.ToXmlNode(document))
        }
        return node
}

