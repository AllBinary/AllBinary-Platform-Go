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
    RawRectangle *Rectangle
    XBorder int
    YBorder int
    Rectangle *Rectangle
    CellPosition ICellPosition
    AnimationX int
    AnimationY int
}

// The Java class TouchButton has subclasses so its variables are this interface
type ITouchButton interface {
    IPaintable
    TouchButton_() *TouchButton
    Paint(graphics IGraphics)
    UpdateRectangle()
    GetCellPositionP() ICellPosition
    GetRectangleP() *Rectangle
    GetTouchButtonInput() *TouchButtonInput
    ToString() string
}

func (this *TouchButton) TouchButton_() *TouchButton {
    return this
}


func (this *TouchButton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    this.Rectangle = RectangleFactory_SINGLETON
    
    
    
}


func NewTouchButton(touchButtonInput *TouchButtonInput, animationInterface AnimationInterface, rawRectangle *Rectangle, cellPosition ICellPosition, xBorder int, yBorder int) *TouchButton {
    this := &TouchButton{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.TouchButtonInput = touchButtonInput
        this.AnimationInterface = any(animationInterface).(IAnimation)
        this.RawRectangle = rawRectangle
        this.CellPosition = cellPosition
        this.XBorder = xBorder
        this.YBorder = yBorder
        this.UpdateRectangle()
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Created: ").Append(this.ToString())), this, this.CommonStrings.CONSTRUCTOR)
    return this
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
            var x int = this.RawRectangle.GetWidth() * this.CellPosition.GetColumn(); _ = x
            var y int = this.RawRectangle.GetHeight() * this.CellPosition.GetRow(); _ = y
            var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
            this.Rectangle = NewRectangle(pointFactory.CreateXY(x + this.XBorder, y + this.YBorder), this.RawRectangle.GetWidth(), this.RawRectangle.GetHeight())
            var point IGPoint = this.Rectangle.GetPoint(); _ = point
            this.AnimationX = point.GetX()
            this.AnimationY = point.GetY()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchButton) GetCellPositionP() ICellPosition {
        return this.CellPosition
}


func (this *TouchButton) GetRectangleP() *Rectangle {
        return this.Rectangle
}


func (this *TouchButton) GetTouchButtonInput() *TouchButtonInput {
        return this.TouchButtonInput
}


func (this *TouchButton) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append("TouchButton: ")
        stringBuffer.Append(stringUtil.ToString(this.Rectangle))
        stringBuffer.Append(" CellPosition: ")
        stringBuffer.Append(stringUtil.ToString(this.CellPosition))
        stringBuffer.Append(" xBorder: ")
        stringBuffer.Appendint(this.XBorder)
        stringBuffer.Append(" yBorder: ")
        stringBuffer.Appendint(this.YBorder)
        return fmt.Sprint(stringBuffer)
}

func (this *TouchButton) String() string {
    return this.ToString()
}

