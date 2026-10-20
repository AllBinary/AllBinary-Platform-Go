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

type FullTouchButton struct {
    TouchButton
}



// throws Exception
func FullTouchButton_CreateButton(touchButtonInput *TouchButtonInput, touchButtonResource ITouchButtonResource, rawRectangle *Rectangle, cellPosition ICellPosition, xBorder int, yBorder int) *FullTouchButton {
        return NewFullTouchButton(touchButtonInput, FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(touchButtonResource.TouchButtonResource_().RESOURCE).GetInstance(0), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(touchButtonResource.TouchButtonResource_().HINT).GetInstance(0), rawRectangle, cellPosition, xBorder, yBorder)
}


func NewFullTouchButton(touchButtonInput *TouchButtonInput, animationInterface IAnimation, hintAnimationInterface IAnimation, rawRectangle *Rectangle, cellPosition ICellPosition, xBorder int, yBorder int) *FullTouchButton {
    this := &FullTouchButton{}
        this.TouchButton = *NewTouchButton(touchButtonInput, animationInterface, hintAnimationInterface, rawRectangle, cellPosition, xBorder, yBorder)
    return this
}


func (this *FullTouchButton) UpdateRectangle() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "updateRectangle", e)
                }
            }()
            var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
            var x int = this.RawRectangle.GetWidth() * this.CellPositionP.GetColumn(); _ = x
            var y int = this.RawRectangle.GetHeight() * this.CellPositionP.GetRow(); _ = y
            this.RectangleP = NewRectangle(pointFactory.CreateXY(x, y), this.RawRectangle.GetWidth() + (2 * this.XBorder), this.RawRectangle.GetHeight() + (2 * this.YBorder))
            var point IGPoint = this.RectangleP.GetPoint(); _ = point
            this.AnimationX = point.GetX() + this.XBorder
            this.AnimationY = point.GetY() + this.YBorder
            this.HintAnimationY = this.AnimationY - this.RectangleP.GetHeight() >> 1
            return
        }(); tryControl1 == 1 {
            return
        }
}

