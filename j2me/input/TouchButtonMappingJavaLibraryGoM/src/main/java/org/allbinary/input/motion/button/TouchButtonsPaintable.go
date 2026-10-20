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

type TouchButtonsPaintable struct {
    Paintable
    ForegroundBasicColor *BasicColor
    ForegroundColor int
    BasicSetColorUtil IBasicColorSetUtil
    TouchButtonFactory *TouchButtonFactory
}


func (this *TouchButtonsPaintable) initFields() {
    
    
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.TouchButtonFactory = TouchButtonFactory_GetInstance()
}


func NewTouchButtonsPaintable(basicColor *BasicColor) *TouchButtonsPaintable {
    this := &TouchButtonsPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.ForegroundBasicColor = basicColor
        this.ForegroundColor = basicColor.IntValue()
    return this
}


func (this *TouchButtonsPaintable) Paint(graphics IGraphics) {
        var list IBasicArrayList = this.TouchButtonFactory.GetList(); _ = list
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.ForegroundBasicColor, this.ForegroundColor)
        var touchButton ITouchButton; _ = touchButton
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            touchButton = any(list.BasicArrayList_().ObjectArray[index]).(ITouchButton)
            touchButton.Paint(graphics)
        }
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            touchButton = any(list.BasicArrayList_().ObjectArray[index]).(ITouchButton)
            touchButton.PaintHint(graphics)
        }
}

