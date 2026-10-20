/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type TechnologyRTSInterfaceImageItem struct {
    ABCustomImageItem
    RtsInterface RTSInterface
    PrimitiveLongUtil *PrimitiveLongUtil
    LEVEL string
    COST string
    DOLLAR string
    MyFontProcessor IMyFontProcessor
    CostString []rune
    CostLength int
    LevelString []rune
    LevelLength int
    AdjustedCostLabelY int
    AdjustedCostX int
    AdjustedCostY int
    AdjustedLevelX int
    AdjustedLevelY int
}

// Implements UpdateMyFontInterface


func (this *TechnologyRTSInterfaceImageItem) initFields() {
    
    this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(10000)
    this.LEVEL = "Level"
    this.COST = "Cost"
    this.DOLLAR = "$"
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.CostString = CharArrayFactory_GetInstance().GetZeroCharArray()
    
    this.LevelString = CharArrayFactory_GetInstance().GetZeroCharArray()
    
    
    
    
    
    
}


func NewTechnologyRTSInterfaceImageItem(label string, img IImage, layout int, altText string, basicColor *BasicColor, rtsInterface RTSInterface) *TechnologyRTSInterfaceImageItem {
    this := &TechnologyRTSInterfaceImageItem{}
        this.ABCustomImageItem = *NewABCustomImageItem(label, img, layout, altText, basicColor, 0)
        this.initFields()
        this.RtsInterface = rtsInterface
        this.Update()
    return this
}


func (this *TechnologyRTSInterfaceImageItem) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        var fontHeight int = font.GetHeight(); _ = fontHeight
        var imageHeight int = 0; _ = imageHeight
        var image IImage = this.GetImage(); _ = image
        if image != NullImage_NULL_IMAGE {
            imageHeight = image.GetHeight()
        }
        this.AdjustedCostLabelY = -this.YOffset + imageHeight - (3 * fontHeight)
        this.AdjustedCostY = -this.YOffset + imageHeight - (2 * fontHeight)
        this.AdjustedCostX = 2 + (len(this.DOLLAR) * (fontHeight - 1))
        this.AdjustedLevelY = -this.YOffset + imageHeight - fontHeight
        this.AdjustedLevelX = 2 + (len(this.LEVEL) * (fontHeight - 1))
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *TechnologyRTSInterfaceImageItem) GetRtsInterface() RTSInterface {
        return this.RtsInterface
}


func (this *TechnologyRTSInterfaceImageItem) Update() {
        this.CostString = this.PrimitiveLongUtil.GetCharArray(this.GetRtsInterface().GetUpgradeCost())
        this.LevelString = this.PrimitiveLongUtil.GetCharArray(this.GetRtsInterface().GetLevel())
}


func (this *TechnologyRTSInterfaceImageItem) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        this.ABCustomImageItem.PaintXY(graphics, x, y)
        var xa int = x + 2; _ = xa
        graphics.DrawString(this.COST, xa, y + this.AdjustedCostLabelY, 0)
        graphics.DrawString(this.DOLLAR, xa, y + this.AdjustedCostY, 0)
        graphics.DrawChars(this.CostString, 0, this.CostLength, x + this.AdjustedCostX, y + this.AdjustedCostY, 0)
        graphics.DrawString(this.LEVEL, xa, y + this.AdjustedLevelY, 0)
        graphics.DrawChars(this.LevelString, 0, this.LevelLength, x + this.AdjustedLevelX, y + this.AdjustedLevelY, 0)
}

