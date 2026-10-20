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

type BasicPopupMenuPaintable struct {
    Paintable
    BasicSetColorUtil IBasicColorSetUtil
    DrawStringUtil *DrawVerticalStringUtil
    Label string
    ForegroundBasicColor *BasicColor
    MyFontProcessor IMyFontProcessor
    Rectangle *Rectangle
    BORDER int
    HeightOffset int
    Offset int
    AnimationInterface IAnimation
}

// Implements UpdateMyFontInterface

var BasicPopupMenuPaintable_NAME string = "MENU"

func (this *BasicPopupMenuPaintable) initFields() {
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.DrawStringUtil = DrawVerticalStringUtil_GetInstance()
    
    
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    
    
    
    
    
}


func NewBasicPopupMenuPaintable(rectangle *Rectangle, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *BasicPopupMenuPaintable {
    this := &BasicPopupMenuPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.ForegroundBasicColor = foregroundBasicColor
        this.Label = BasicPopupMenuPaintable_NAME
        this.Rectangle = rectangle
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
            this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
        } else {
            this.AnimationInterface = NewRectangleFilledAnimation(this.Rectangle.GetWidth(), this.Rectangle.GetHeight(), BasicColorFactory_GetInstance().TRANSPARENT_GREY)
        }
        this.Init(rectangle)
    return this
}


func (this *BasicPopupMenuPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        var features *Features = Features_GetInstance(); _ = features
        var isOpenGL bool = features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL); _ = isOpenGL
        this.DrawStringUtil.UpdateMeasurement(graphics, this.Label)
        var BORDER int = 0; _ = BORDER
        if J2MEUtil_IsHTML() || isOpenGL {
            BORDER = MyFontProcessor_DefaultCharWidth(font) / 2
        } else if AndroidUtil_IsAndroid() || J2MEUtil_IsJ2SE() || SWTUtil_IsSWT {
            BORDER = MyFontProcessor_DefaultCharWidth(font)
        } else {
            BORDER = MyFontProcessor_DefaultCharWidth(font) * 2
        }
        this.BORDER = BORDER
        this.HeightOffset = this.Rectangle.GetHeight() - (font.GetHeight() * len(BasicPopupMenuPaintable_NAME))
        if OpenGLFeatureUtil_GetInstance().IsAnyThreed() {
            this.HeightOffset -= font.GetHeight() + 2
            if AndroidUtil_IsAndroid() {
                this.HeightOffset = font.GetHeight()
            } else {
                this.HeightOffset -= font.GetHeight() + 2
            }
        }
        this.Offset = (this.HeightOffset >> 1) + (func() int { if SWTJOGLProcessor_GetInstance().IsJOGL() { return 4 }; return 0 }())
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


// throws Exception
func (this *BasicPopupMenuPaintable) Init(rectangle *Rectangle) {
        this.Rectangle = rectangle
        var width int = this.Rectangle.GetWidth(); _ = width
        var height int = this.Rectangle.GetHeight(); _ = height
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
        } else {
            var rectangleFilledAnimation IRectangleFilledAnimation = any(this.AnimationInterface).(IRectangleFilledAnimation); _ = rectangleFilledAnimation
            rectangleFilledAnimation.SetWidth(width)
            rectangleFilledAnimation.SetHeight(height)
        }
}


func (this *BasicPopupMenuPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var point IGPoint = this.Rectangle.GetPoint(); _ = point
        var x int = point.GetX(); _ = x
        var y int = point.GetY(); _ = y
        var width int = this.Rectangle.GetWidth(); _ = width
        var height int = this.Rectangle.GetHeight(); _ = height
        this.AnimationInterface.PaintXY(graphics, x, y)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.ForegroundBasicColor)
        this.DrawStringUtil.PaintVerticle(graphics, this.Label, x + this.BORDER, y + this.Offset, 0)
        graphics.DrawRect(x, y, width, height)
}

