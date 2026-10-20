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

type VectorBaseRotationAnimation struct {
    RotationAnimation
    CurrentPoints [][][]int
}

// Implements VectorAnimationInterface

// The Java class VectorBaseRotationAnimation has subclasses so its variables are this interface
type IVectorBaseRotationAnimation interface {
    IRotationAnimation
    VectorBaseRotationAnimation_() *VectorBaseRotationAnimation
    GetAnimationSize() int
    GetFrame() int
    SetFrame(index int)
    NextFrame()
    PreviousFrame()
    GetSize() int
    SetSequence(sequence []int)
    GetSequence() []int
    PaintXY(graphics IGraphics, x int, y int)
    GetPoints(frame int) [][]int
    SetPoints(currentPoints [][][]int)
}

func (this *VectorBaseRotationAnimation) VectorBaseRotationAnimation_() *VectorBaseRotationAnimation {
    return this
}


func (this *VectorBaseRotationAnimation) initFields() {
    this.CurrentPoints = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY_ARRAY
}


func NewVectorBaseRotationAnimation(angleInfo *AngleInfo, currentPoints [][][]int, basicColor *BasicColor, animationBehavior IAnimationBehavior) *VectorBaseRotationAnimation {
    this := &VectorBaseRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(angleInfo, CircularIndexUtil_CreateInstance(360 / int(angleInfo.GetAngleIncrementInfo().GetAngleIncrement())), animationBehavior)
        this.initFields()
        this.SetPoints(currentPoints)
        this.SetBasicColorP(basicColor)
    return this
}


// throws Exception
func (this *VectorBaseRotationAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *VectorBaseRotationAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *VectorBaseRotationAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
}


func (this *VectorBaseRotationAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
}


func (this *VectorBaseRotationAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
}


func (this *VectorBaseRotationAnimation) GetSize() int {
        return len(this.CurrentPoints)
}


func (this *VectorBaseRotationAnimation) SetSequence(sequence []int) {
}


func (this *VectorBaseRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *VectorBaseRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.BasicColor)
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


func (this *VectorBaseRotationAnimation) GetPoints(frame int) [][]int {
        return this.CurrentPoints[frame]
}


func (this *VectorBaseRotationAnimation) SetPoints(currentPoints [][][]int) {
        this.CurrentPoints = currentPoints
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(this.CurrentPoints))
}

