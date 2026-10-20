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

type VectorAnimation struct {
    IndexedAnimation
    CurrentPoints [][][]int
    CircularIndexUtil *CircularIndexUtil
}

// Implements VectorAnimationInterface

// The Java class VectorAnimation has subclasses so its variables are this interface
type IVectorAnimation interface {
    IIndexedAnimation
    VectorAnimation_() *VectorAnimation
    SetBasicColorP(basicColor *BasicColor)
    ChangeBasicColor(basicColor *BasicColor)
    SetAlpha(alpha int)
    SetScale(scaleX float32, scaleY float32)
    SetMaxScale(maxScaleX float32, maxScaleY float32)
    UpdateModifiers()
    GetAnimationSize() int
    GetSize() int
    SetSequence(sequence []int)
    GetSequence() []int
    NextFrame()
    PreviousFrame()
    PaintVectors(graphics IGraphics, x int, y int)
    PaintXY(graphics IGraphics, x int, y int)
    GetFrame() int
    SetFrame(index int)
    GetPoints(frame int) [][]int
    SetPoints(currentPoints [][][]int)
}

func (this *VectorAnimation) VectorAnimation_() *VectorAnimation {
    return this
}


func (this *VectorAnimation) initFields() {
    this.CurrentPoints = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY_ARRAY
    this.CircularIndexUtil = CircularIndexUtil_CreateInstance(0)
}


func VectorAnimation_CreateVectorAnimation(currentPoints2 [][]int, basicColor *BasicColor, animationBehavior IAnimationBehavior) IVectorAnimation {
        var currentPoints [][][]int = func() [][][]int { array := make([][][]int, 1); for index := range array { array[index] = make([][]int, len(currentPoints2)); for index2 := range array[index] { array[index][index2] = make([]int, 2) } }; return array }(); _ = currentPoints
        var size int = len(currentPoints2); _ = size
        for index := int(0); index < size; index++ {
            currentPoints[0][index][0] = currentPoints2[index][0]
            currentPoints[0][index][1] = currentPoints2[index][1]
        }
        return NewVectorAnimation(currentPoints, basicColor, animationBehavior)
}


func NewVectorAnimation(currentPoints [][][]int, basicColor *BasicColor, animationBehavior IAnimationBehavior) *VectorAnimation {
    this := &VectorAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.SetPoints(currentPoints)
        this.SetBasicColorP(basicColor)
    return this
}


func (this *VectorAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.IndexedAnimation.SetBasicColorP(basicColor)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *VectorAnimation) ChangeBasicColor(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetChangeBasicColor() == nil || this.GetChangeBasicColor().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.IndexedAnimation.ChangeBasicColor(basicColor)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *VectorAnimation) SetAlpha(alpha int) {
        var changed bool = false; _ = changed
        if this.AlphaP != alpha {
            changed = true
        }
        this.IndexedAnimation.SetAlpha(alpha)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *VectorAnimation) SetScale(scaleX float32, scaleY float32) {
}


func (this *VectorAnimation) SetMaxScale(maxScaleX float32, maxScaleY float32) {
}


func (this *VectorAnimation) UpdateModifiers() {
}


// throws Exception
func (this *VectorAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *VectorAnimation) GetSize() int {
        return len(this.CurrentPoints)
}


func (this *VectorAnimation) SetSequence(sequence []int) {
}


func (this *VectorAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *VectorAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
}


func (this *VectorAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
}


func (this *VectorAnimation) PaintVectors(graphics IGraphics, x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "paintVectors", e)
                }
            }()
            var nextPointX int = 0; _ = nextPointX
            var nextPointY int = 0; _ = nextPointY
            var nextPoint []int; _ = nextPoint
            var point []int; _ = point
            var currentPointsFrame [][]int = this.CurrentPoints[this.CircularIndexUtil.GetIndex()]; _ = currentPointsFrame
            var size int = len(currentPointsFrame); _ = size
            var index int = size - 2; _ = index
            for Java_PreDecrement(&index) >= 0 {
                nextPoint = currentPointsFrame[index]
                point = currentPointsFrame[index + 1]
                nextPointX = nextPoint[0]
                nextPointY = nextPoint[1]
                if nextPointX != 1000 {
                    graphics.DrawLine(point[0] + x, point[1] + y, nextPointX + x, nextPointY + y)
                } else {
                    index--
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *VectorAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.BasicColor)
        this.PaintVectors(graphics, x, y)
}


func (this *VectorAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *VectorAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
}


func (this *VectorAnimation) GetPoints(frame int) [][]int {
        return this.CurrentPoints[frame]
}


func (this *VectorAnimation) SetPoints(currentPoints [][][]int) {
        this.CurrentPoints = currentPoints
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(currentPoints))
}

