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

type FontFamilyView struct {
    CssPropertyView
}

// The Java class FontFamilyView has subclasses so its variables are this interface
type IFontFamilyView interface {
    ICssPropertyView
    FontFamilyView_() *FontFamilyView
}

func (this *FontFamilyView) FontFamilyView_() *FontFamilyView {
    return this
}





// Java overloads of new FontFamilyView
func NewFontFamilyView(args ...any) *FontFamilyView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewFontFamilyView_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewFontFamilyView_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new FontFamilyView for the arguments")
}
func NewFontFamilyView_1() *FontFamilyView {
    this := &FontFamilyView{}
        this.CssPropertyView = *NewCssPropertyView(FontFamilyData_GetInstance().VALUE)
    return this
}


func NewFontFamilyView_2(node *Node) *FontFamilyView {
    this := &FontFamilyView{}
        this.CssPropertyView = *NewCssPropertyView(node)
    return this
}

