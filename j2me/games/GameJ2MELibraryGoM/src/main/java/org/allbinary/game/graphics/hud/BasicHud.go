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

type BasicHud struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CanvasStrings *CanvasStrings
    BasicSetColorUtil IBasicColorSetUtil
    DisplayInfo *DisplayInfoSingleton
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    Location int
    Direction int
    BufferZone int
    BufferZoneY int
    HudGraphicsPosition *HudGraphicsPosition
    X int
    Y int
    BasicColor *BasicColor
    Color int
    UpdateMaxWidth int
    UpdateMaxHeight int
    OffsetY int
}

// Implements UpdateMyFontInterface

// The Java class BasicHud has subclasses so its variables are this interface
type IBasicHud interface {
    BasicHud_() *BasicHud
    UpdateMeasurement(graphics IGraphics)
    GetLocation() int
    GetDirection() int
    GetHudGraphicsPositionWH(width int, height int, maxWidth int, maxHeight int) *HudGraphicsPosition
    OnEvent(eventObject IAllBinaryEventObject)
    OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent)
    GetPoint(x int, y int) IGPoint
    GetBufferZone() int
    SetBufferZone(bufferZone int)
    SetLocation(location int)
    SetDirection(direction int)
    GetHudGraphicsPosition() *HudGraphicsPosition
    GetColor() int
    GetX() int
    SetX(x int)
    SetBasicColorP(basicColor *BasicColor)
    GetBasicColorP() *BasicColor
    SetY(y int)
    GetY() int
    PaintSSO(graphics IGraphics, string_ string, string2 string, offset int)
    PaintSSOO(graphics IGraphics, string_ string, string2 string, offset int, offset2 int)
    PaintDXY(graphics IGraphics, charArray []rune, offset int, len_ int, charArray2 []rune, offset2 int, len2 int, xOffset int, xOffset2 int)
    PaintDX(graphics IGraphics, charArray []rune, offset int, len_ int, charArray2 []rune, offset2 int, len2 int, xOffset int)
    PaintOffsetAndLength(graphics IGraphics, charArray []rune, offset int, len_ int)
    PaintString(graphics IGraphics, string_ string)
}

func (this *BasicHud) BasicHud_() *BasicHud {
    return this
}


func (this *BasicHud) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    
    
    
    
    this.HudGraphicsPosition = HudGraphicsPosition_NULL_HUD_GRAPHICS_POSITION
    
    
    this.BasicColor = BasicColorFactory_GetInstance().BLACK
    
    
    
    
}


func NewBasicHud(location int, direction int, bufferZone int, basicColor *BasicColor) *BasicHud {
    this := &BasicHud{}
    this.initFields()
        this.SetLocation(location)
        this.SetDirection(direction)
        this.SetBufferZone(bufferZone)
        this.OnDisplayChangeEvent(DisplayInfoSingleton_GetInstance().DisplayChangeEvent)
        this.SetBasicColorP(basicColor)
        this.Color = basicColor.IntValue()
    return this
}


func (this *BasicHud) UpdateMeasurement(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT, e)
                }
            }()
            this.HudGraphicsPosition = this.GetHudGraphicsPositionWH(this.DisplayInfo.GetLastWidth(), this.DisplayInfo.GetLastHeight(), this.UpdateMaxWidth, this.UpdateMaxHeight)
            this.X = this.HudGraphicsPosition.GetPoint().GetX()
            this.SetY(this.HudGraphicsPosition.GetPoint().GetY())
            return
        }(); tryControl1 == 1 {
            return
        }
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *BasicHud) GetLocation() int {
        return this.Location
}


func (this *BasicHud) GetDirection() int {
        return this.Direction
}


// throws Exception
func (this *BasicHud) GetHudGraphicsPositionWH(width int, height int, maxWidth int, maxHeight int) *HudGraphicsPosition {
        var x int = 0; _ = x
        var y int = 0; _ = y
        var anchor int = 0; _ = anchor
        var basicHudFactory *BasicHudFactory = BasicHudFactory_GetInstance(); _ = basicHudFactory
        if basicHudFactory.BOTTOMLEFT == this.GetLocation() {
            x = this.BufferZone + 2
            y = height - maxHeight - this.BufferZone
            anchor = Graphics_BOTTOM & Graphics_LEFT
        } else if basicHudFactory.BOTTOMRIGHT == this.GetLocation() {
            x = width - maxWidth
            y = height - maxHeight - this.BufferZone
            anchor = Graphics_BOTTOM & Graphics_RIGHT
        } else if basicHudFactory.TOPLEFT == this.GetLocation() {
            x = this.BufferZone + 2
            y = this.BufferZoneY
            anchor = Anchor_TOP_LEFT
        } else if basicHudFactory.TOPRIGHT == this.GetLocation() {
            x = width - maxWidth
            y = this.BufferZoneY
            anchor = Graphics_TOP & Graphics_RIGHT
        } else if basicHudFactory.TOPCENTER == this.GetLocation() {
            x = ((width - maxWidth) / 2)
            y = this.BufferZoneY
            anchor = Graphics_TOP & Graphics_HCENTER
        } else if basicHudFactory.BOTTOMCENTER == this.GetLocation() {
            x = ((width - maxWidth) / 2)
            y = height - maxHeight - this.BufferZone
            anchor = Graphics_BOTTOM & Graphics_HCENTER
        } else if basicHudFactory.ABSOLUTE == this.GetLocation() {
            x = maxHeight
            y = maxWidth
            anchor = 0
        }
        return NewHudGraphicsPosition(this.GetPoint(x, y), anchor)
}


func (this *BasicHud) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *BasicHud) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *BasicHud) GetPoint(x int, y int) IGPoint {
        return PointFactory_GetInstance().CreateXY(x, y)
}


func (this *BasicHud) GetBufferZone() int {
        return this.BufferZone
}


func (this *BasicHud) SetBufferZone(bufferZone int) {
        this.BufferZone = bufferZone
        if J2MEUtil_IsHTML() {
            this.BufferZoneY = this.BufferZone
        } else {
            this.BufferZoneY = this.BufferZone + 5
        }
}


func (this *BasicHud) SetLocation(location int) {
        this.Location = location
}


func (this *BasicHud) SetDirection(direction int) {
        this.Direction = direction
}


func (this *BasicHud) GetHudGraphicsPosition() *HudGraphicsPosition {
        return this.HudGraphicsPosition
}


func (this *BasicHud) GetColor() int {
        return this.Color
}


func (this *BasicHud) GetX() int {
        return this.X
}


func (this *BasicHud) SetX(x int) {
        this.X = x
}


func (this *BasicHud) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *BasicHud) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *BasicHud) SetY(y int) {
        this.Y = y
}


func (this *BasicHud) GetY() int {
        return this.Y
}


func (this *BasicHud) PaintSSO(graphics IGraphics, string_ string, string2 string, offset int) {
        this.MyFontProcessor.Process(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        graphics.DrawString(string_, this.X, this.GetY(), this.HudGraphicsPosition.GetAnchor())
        graphics.DrawString(string2, this.X + offset, this.GetY(), this.HudGraphicsPosition.GetAnchor())
}


func (this *BasicHud) PaintSSOO(graphics IGraphics, string_ string, string2 string, offset int, offset2 int) {
        this.MyFontProcessor.Process(graphics)
        graphics.SetColor(this.GetColor())
        graphics.DrawString(string_, this.X + offset, this.GetY(), this.HudGraphicsPosition.GetAnchor())
        graphics.DrawString(string2, this.X + offset2, this.GetY(), this.HudGraphicsPosition.GetAnchor())
}


func (this *BasicHud) PaintDXY(graphics IGraphics, charArray []rune, offset int, len_ int, charArray2 []rune, offset2 int, len2 int, xOffset int, xOffset2 int) {
        this.MyFontProcessor.Process(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        var y int = this.GetY(); _ = y
        graphics.DrawChars(charArray, offset, len_, this.X + xOffset, y, this.HudGraphicsPosition.GetAnchor())
        graphics.DrawChars(charArray2, offset2, len2, this.X + xOffset2, y, this.HudGraphicsPosition.GetAnchor())
}


func (this *BasicHud) PaintDX(graphics IGraphics, charArray []rune, offset int, len_ int, charArray2 []rune, offset2 int, len2 int, xOffset int) {
        this.MyFontProcessor.Process(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        var y int = this.GetY(); _ = y
        graphics.DrawChars(charArray, offset, len_, this.X, y, this.HudGraphicsPosition.GetAnchor())
        graphics.DrawChars(charArray2, offset2, len2, this.X + xOffset, y, this.HudGraphicsPosition.GetAnchor())
}


func (this *BasicHud) PaintOffsetAndLength(graphics IGraphics, charArray []rune, offset int, len_ int) {
        this.MyFontProcessor.Process(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        var y int = this.GetY(); _ = y
        graphics.DrawChars(charArray, offset, len_, this.X, y, this.HudGraphicsPosition.GetAnchor())
}


func (this *BasicHud) PaintString(graphics IGraphics, string_ string) {
        this.MyFontProcessor.Process(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        var y int = this.GetY() + this.OffsetY; _ = y
        graphics.DrawString(string_, this.X, y, this.HudGraphicsPosition.GetAnchor())
}

