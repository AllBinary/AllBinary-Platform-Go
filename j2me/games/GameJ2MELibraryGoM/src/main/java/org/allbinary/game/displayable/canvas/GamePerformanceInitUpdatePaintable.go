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

type GamePerformanceInitUpdatePaintable struct {
    InitUpdatePaintable
    HalfHeight int
    YArray []int
    MyFontProcessor IMyFontProcessor
    BaseRefreshHelperCharArray [][]rune
    DefaultStringWidth int
    RED int
}

// Implements UpdateMyFontInterface


func (this *GamePerformanceInitUpdatePaintable) initFields() {
    this.HalfHeight = DisplayInfoSingleton_GetInstance().GetLastHalfHeight()
    this.YArray = []int{this.HalfHeight + 30, this.HalfHeight + 30, this.HalfHeight + 30 + 15, this.HalfHeight + 30 + 15, this.HalfHeight + 30 + 30, this.HalfHeight + 30 + 30, this.HalfHeight + 30 + 45, this.HalfHeight + 30 + 45}
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.BaseRefreshHelperCharArray = func() [][]rune { array := make([][]rune, 0); for index := range array { array[index] = make([]rune, 0) }; return array }()
    
    this.RED = BasicColorFactory_GetInstance().RED.IntValue()
}

func NewGamePerformanceInitUpdatePaintable() *GamePerformanceInitUpdatePaintable {
    this := &GamePerformanceInitUpdatePaintable{}
    this.InitUpdatePaintable = *NewInitUpdatePaintable()
    this.initFields()
    return this
}


func (this *GamePerformanceInitUpdatePaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.DefaultStringWidth = MyFontProcessor_DefaultStringWidth(font, 2)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *GamePerformanceInitUpdatePaintable) Init() {
}


func (this *GamePerformanceInitUpdatePaintable) Update() {
        this.BaseRefreshHelperCharArray = GameStatisticsFactory_GetInstance().To2DCharArray()
}


func (this *GamePerformanceInitUpdatePaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        graphics.SetColor(this.RED)
        var charArray []rune; _ = charArray
        var charArray2 []rune; _ = charArray2
        var size2 int; _ = size2
        var size3 int; _ = size3
        var size int = len(this.BaseRefreshHelperCharArray) - 2; _ = size
        for index := int(size); index >= 0; index -= 2 {
            charArray = this.BaseRefreshHelperCharArray[index]
            charArray2 = this.BaseRefreshHelperCharArray[index + 1]
            size2 = len(charArray)
            size3 = len(charArray2)
            graphics.DrawChars(charArray, 0, size2, 0, this.YArray[index], 0)
            graphics.DrawChars(charArray2, 0, size3, size2 * this.DefaultStringWidth, this.YArray[index + 1], 0)
        }
}

