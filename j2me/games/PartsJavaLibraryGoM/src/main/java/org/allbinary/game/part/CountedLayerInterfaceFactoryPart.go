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

type CountedLayerInterfaceFactoryPart struct {
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    AnimationInterface IAnimation
    Total int
    TotalString []rune
    XOffset int
    CountedPickedUpLayerInterfaceFactory ICountedPickedUpLayerInterfaceFactory
    PrimitiveLongUtil *PrimitiveLongUtil
}

// Implements PartInterface, UpdateMyFontInterface

var CountedLayerInterfaceFactoryPart_NULL_COUNTED_LAYER_INTERFACE_FACTORY *CountedLayerInterfaceFactoryPart = NewCountedLayerInterfaceFactoryPart(0, CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY)

func (this *CountedLayerInterfaceFactoryPart) initFields() {
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
    
    this.TotalString = NullUtil_GetInstance().NULL_CHAR_ARRAY
    
    this.CountedPickedUpLayerInterfaceFactory = CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY
    
}


func NewCountedLayerInterfaceFactoryPart(total int, countedPickedUpLayerInterfaceFactory ICountedPickedUpLayerInterfaceFactory) *CountedLayerInterfaceFactoryPart {
    this := &CountedLayerInterfaceFactoryPart{}
    this.initFields()
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(1000)
        this.Init(total, countedPickedUpLayerInterfaceFactory)
    return this
}


func (this *CountedLayerInterfaceFactoryPart) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.SetXOffset(font.CharsWidth(this.TotalString, 0, this.PrimitiveLongUtil.GetCurrentTotalDigits()) + (font.GetSize() >> 1))
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *CountedLayerInterfaceFactoryPart) Init(total int, countedPickedUpLayerInterfaceFactory ICountedPickedUpLayerInterfaceFactory) {
        this.SetCountedPickedUpLayerInterfaceFactory(countedPickedUpLayerInterfaceFactory)
        this.Total = total
}


// throws Exception
func (this *CountedLayerInterfaceFactoryPart) GetInstance(hashtable *ABHashtable[any, any], x int, y int, z int) IAllBinaryLayer {
        if this.Total > 0 {
            this.SetTotal(this.Total - 1)
            return this.GetCountedPickedUpLayerInterfaceFactory().GetNextInstance(hashtable, x, y, z)
        } else {
            panic(errors.New(fmt.Sprint("No more left. Could use a listener to automatically remove")))
        }
}


func (this *CountedLayerInterfaceFactoryPart) GetTotal() int {
        return this.Total
}


func (this *CountedLayerInterfaceFactoryPart) SetTotal(total int) {
        this.Total = total
        this.TotalString = this.PrimitiveLongUtil.GetCharArray(total)
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *CountedLayerInterfaceFactoryPart) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
}


func (this *CountedLayerInterfaceFactoryPart) PaintThreed(graphics IGraphics) {
}


func (this *CountedLayerInterfaceFactoryPart) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *CountedLayerInterfaceFactoryPart) SetAnimationInterface(animationInterface IAnimation) {
        this.AnimationInterface = animationInterface
}


func (this *CountedLayerInterfaceFactoryPart) GetXOffset() int {
        return this.XOffset
}


func (this *CountedLayerInterfaceFactoryPart) SetXOffset(totalStringWidth int) {
        this.XOffset = totalStringWidth
}


func (this *CountedLayerInterfaceFactoryPart) GetTotalString() []rune {
        return this.TotalString
}


func (this *CountedLayerInterfaceFactoryPart) GetCountedPickedUpLayerInterfaceFactory() ICountedPickedUpLayerInterfaceFactory {
        return this.CountedPickedUpLayerInterfaceFactory
}


func (this *CountedLayerInterfaceFactoryPart) SetCountedPickedUpLayerInterfaceFactory(countedPickedUpLayerInterfaceFactory ICountedPickedUpLayerInterfaceFactory) {
        this.CountedPickedUpLayerInterfaceFactory = countedPickedUpLayerInterfaceFactory
}

