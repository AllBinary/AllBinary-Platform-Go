/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSLayerCostAnimation struct {
    Animation
    LogUtil *LogUtil
    PrimitiveLongUtil *PrimitiveLongUtil
    DOLLAR string
    Image IImage
    LayerInterfaceFactoryInterface CostLayerInterfaceFactoryInterface
    MyFontProcessor IMyFontProcessor
    CostString []rune
    Len int
    AdjustedCostX int
    FontHeight int
}

// Implements TechEventListenerInterface, UpdateMyFontInterface


func (this *RTSLayerCostAnimation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(10000)
    this.DOLLAR = "$"
    
    
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.CostString = NullUtil_GetInstance().NULL_CHAR_ARRAY
    
    
    this.FontHeight = 0
}


func NewRTSLayerCostAnimation(image IImage, layerInterfaceFactoryInterface CostLayerInterfaceFactoryInterface) *RTSLayerCostAnimation {
    this := &RTSLayerCostAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.Image = image
        this.LayerInterfaceFactoryInterface = layerInterfaceFactoryInterface
        this.Update()
    return this
}


func (this *RTSLayerCostAnimation) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.AdjustedCostX = font.StringWidth(this.DOLLAR)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *RTSLayerCostAnimation) OnEvent(event IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *RTSLayerCostAnimation) OnTechEvent(event IAllBinaryEventObject) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "onTechEvent", e)
                }
            }()
            this.Update()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *RTSLayerCostAnimation) Update() {
        this.Len = 0
        this.CostString = this.PrimitiveLongUtil.GetCharArray(this.LayerInterfaceFactoryInterface.GetCost())
        this.Len = this.PrimitiveLongUtil.GetCurrentTotalDigits()
}


func (this *RTSLayerCostAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        this.Animation.PaintXY(graphics, x, y)
        var adjustedCostY int = this.Image.GetHeight() - this.FontHeight; _ = adjustedCostY
        var xa int = x + 2; _ = xa
        graphics.DrawString(this.DOLLAR, xa, y + adjustedCostY, 0)
        graphics.DrawChars(this.CostString, 0, this.Len, x + this.AdjustedCostX, y + adjustedCostY, 0)
}

