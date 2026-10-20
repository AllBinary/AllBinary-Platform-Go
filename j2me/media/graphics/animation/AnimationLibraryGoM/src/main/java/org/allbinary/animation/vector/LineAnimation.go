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

type LineAnimation struct {
    Animation
    NewLine [][]int
    TempLine [][]int
    Line [][]int
}

// Implements ColorCompositeInterface, PositionInterface


func (this *LineAnimation) initFields() {
    this.NewLine = func() [][]int { array := make([][]int, 2); for index := range array { array[index] = make([]int, 2) }; return array }()
    this.TempLine = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY
    this.Line = func() [][]int { array := make([][]int, 2); for index := range array { array[index] = make([]int, 2) }; return array }()
}


func NewLineAnimation(basicColor *BasicColor) *LineAnimation {
    this := &LineAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.SetBasicColorP(basicColor)
    return this
}


func (this *LineAnimation) Init(x int, y int) {
        this.NewLine[0][0] = x
        this.NewLine[0][1] = y
        this.NewLine[1][0] = x
        this.NewLine[1][1] = y
        this.TempLine = this.Line
        this.Line = this.NewLine
        this.NewLine = this.TempLine
}


func (this *LineAnimation) SetPosition(x int, y int, z int) {
        this.NewLine[0][0] = this.Line[1][0]
        this.NewLine[0][1] = this.Line[1][1]
        this.NewLine[1][0] = x
        this.NewLine[1][1] = y
        this.TempLine = this.Line
        this.Line = this.NewLine
        this.NewLine = this.TempLine
}


func (this *LineAnimation) PaintXY(graphics IGraphics, x int, y int) {
        var currentLine [][]int = this.Line; _ = currentLine
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        graphics.DrawLine(currentLine[0][0], currentLine[0][1], currentLine[1][0], currentLine[1][1])
}


func (this *LineAnimation) NextFrame() {
}

