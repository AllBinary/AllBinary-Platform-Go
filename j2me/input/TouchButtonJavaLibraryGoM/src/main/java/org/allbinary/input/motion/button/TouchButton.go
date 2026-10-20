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

type TouchButton struct {
    Paintable
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TouchButtonInput *TouchButtonInput
    AnimationInterface IAnimation
    HintAnimationInterface IAnimation
    RawRectangle *Rectangle
    XBorder int
    YBorder int
    RectangleP *Rectangle
    CellPositionP ICellPosition
    AnimationX int
    AnimationY int
    HintAnimationY int
}

// The Java class TouchButton has subclasses so its variables are this interface
type ITouchButton interface {
    IPaintable
    TouchButton_() *TouchButton
    PaintHint(graphics IGraphics)
    Paint(graphics IGraphics)
    UpdateRectangle()
    GetCellPosition() ICellPosition
    GetRectangle() *Rectangle
    GetTouchButtonInput() *TouchButtonInput
    ToString() string
}

func (this *TouchButton) TouchButton_() *TouchButton {
    return this
}


func (this *TouchButton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    this.RectangleP = RectangleFactory_SINGLETON
    
    
    
    
}


// throws Exception
func TouchButton_CreateButton(touchButtonInput *TouchButtonInput, touchButtonResource ITouchButtonResource, rawRectangle *Rectangle, cellPosition ICellPosition, xBorder int, yBorder int) ITouchButton {
        return NewTouchButton(touchButtonInput, FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(touchButtonResource.TouchButtonResource_().RESOURCE).GetInstance(0), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(touchButtonResource.TouchButtonResource_().HINT).GetInstance(0), rawRectangle, cellPosition, xBorder, yBorder)
}


func NewTouchButton(touchButtonInput *TouchButtonInput, animationInterface IAnimation, hintAnimationInterface IAnimation, rawRectangle *Rectangle, cellPosition ICellPosition, xBorder int, yBorder int) *TouchButton {
    this := &TouchButton{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.TouchButtonInput = touchButtonInput
        this.AnimationInterface = animationInterface
        this.HintAnimationInterface = hintAnimationInterface
        this.RawRectangle = rawRectangle
        this.CellPositionP = cellPosition
        this.XBorder = xBorder
        this.YBorder = yBorder
        this.UpdateRectangle()
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Created: ").Append(this.ToString())), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *TouchButton) PaintHint(graphics IGraphics) {
        this.HintAnimationInterface.PaintXY(graphics, this.AnimationX, this.HintAnimationY)
}


func (this *TouchButton) Paint(graphics IGraphics) {
        this.AnimationInterface.PaintXY(graphics, this.AnimationX, this.AnimationY)
}


func (this *TouchButton) UpdateRectangle() {
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
            this.RectangleP = NewRectangle(pointFactory.CreateXY(x + this.XBorder, y + this.YBorder), this.RawRectangle.GetWidth(), this.RawRectangle.GetHeight())
            var point IGPoint = this.RectangleP.GetPoint(); _ = point
            this.AnimationX = point.GetX()
            this.AnimationY = point.GetY()
            this.HintAnimationY = this.AnimationY - 32
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchButton) GetCellPosition() ICellPosition {
        return this.CellPositionP
}


func (this *TouchButton) GetRectangle() *Rectangle {
        return this.RectangleP
}


func (this *TouchButton) GetTouchButtonInput() *TouchButtonInput {
        return this.TouchButtonInput
}


func (this *TouchButton) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append("TouchButton: ")
        stringBuffer.Append(stringUtil.ToString(this.GetRectangle()))
        stringBuffer.Append(" CellPosition: ")
        stringBuffer.Append(stringUtil.ToString(this.CellPositionP))
        stringBuffer.Append(" xBorder: ")
        stringBuffer.Appendint(this.XBorder)
        stringBuffer.Append(" yBorder: ")
        stringBuffer.Appendint(this.YBorder)
        return fmt.Sprint(stringBuffer)
}

func (this *TouchButton) String() string {
    return this.ToString()
}

