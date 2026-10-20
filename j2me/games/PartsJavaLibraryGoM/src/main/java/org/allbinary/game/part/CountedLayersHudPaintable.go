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

type CountedLayersHudPaintable struct {
    DisplayInfoSingleton *DisplayInfoSingleton
    PartInterfaceArray []PartInterface
    CountedTotalStringColor int
    CountedPartsBorder int
    StartIndex int
    DropSize int
    MyFontProcessor IMyFontProcessor
    Height int
}

// Implements PaintableInterface, UpdateMyFontInterface

var CountedLayersHudPaintable_XXStringWidth int = 0

func (this *CountedLayersHudPaintable) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    
    
    
    
    
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    
}


func NewCountedLayersHudPaintable(partInterfaceArray []PartInterface, dropSize int, startIndex int, countedTotalStringColor int, countedPartsBorder int) *CountedLayersHudPaintable {
    this := &CountedLayersHudPaintable{}
    this.initFields()
        this.PartInterfaceArray = partInterfaceArray
        this.StartIndex = startIndex
        this.CountedTotalStringColor = countedTotalStringColor
        this.CountedPartsBorder = countedPartsBorder
        this.DropSize = dropSize
    return this
}


func (this *CountedLayersHudPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.Height = font.GetHeight()
        if this.DropSize > font.GetHeight() {
            this.Height = this.DropSize
        }
        if CountedLayersHudPaintable_XXStringWidth == 0 {
            var XXString string = "XX"; _ = XXString
            CountedLayersHudPaintable_XXStringWidth = font.StringWidth(XXString)
        }
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *CountedLayersHudPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var lastWidth int = this.DisplayInfoSingleton.GetLastWidth(); _ = lastWidth
        var widthEdge int = lastWidth - this.DropSize; _ = widthEdge
        var size int = len(this.PartInterfaceArray); _ = size
        var count int = 0; _ = count
        var y int; _ = y
        var countedLayerInterfaceFactory *CountedLayerInterfaceFactoryPart; _ = countedLayerInterfaceFactory
        var pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface; _ = pickedUpLayerInterfaceFactoryInterface
        var layerInterface IAllBinaryLayer; _ = layerInterface
        var charArray []rune; _ = charArray
        for index := int(this.StartIndex); index < size; index++ {
            countedLayerInterfaceFactory = any(this.PartInterfaceArray[index]).(*CountedLayerInterfaceFactoryPart)
            if countedLayerInterfaceFactory.GetTotal() > 0 {
                pickedUpLayerInterfaceFactoryInterface = countedLayerInterfaceFactory.GetCountedPickedUpLayerInterfaceFactory()
                layerInterface = pickedUpLayerInterfaceFactoryInterface.GetIconLayer()
                y = 40 + (count * this.Height)
                layerInterface.SetPosition(widthEdge, y, layerInterface.GetZP())
                layerInterface.Paint(graphics)
                graphics.SetColor(this.CountedTotalStringColor)
                countedLayerInterfaceFactory.Paint(graphics)
                charArray = countedLayerInterfaceFactory.GetTotalString()
                graphics.DrawChars(charArray, 0, len(charArray), widthEdge - countedLayerInterfaceFactory.GetXOffset(), y, 0)
                count++
            }
        }
        if count > 0 {
            graphics.SetColor(this.CountedPartsBorder)
            graphics.DrawRect(lastWidth - (CountedLayersHudPaintable_XXStringWidth + this.DropSize), 40, CountedLayersHudPaintable_XXStringWidth + this.DropSize, (count * this.Height) + 3)
        }
}


func (this *CountedLayersHudPaintable) PaintThreed(graphics IGraphics) {
}

