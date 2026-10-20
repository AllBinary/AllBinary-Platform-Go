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

type BasicGraphicsPipeline struct {
    PointBasicArrayList IBasicArrayList
    NoDecimalTrigTable *NoDecimalTrigTable
}

// The Java class BasicGraphicsPipeline has subclasses so its variables are this interface
type IBasicGraphicsPipeline interface {
    BasicGraphicsPipeline_() *BasicGraphicsPipeline
    SetInitMatrix(points IBasicArrayList)
    CreateMatrix()
    Translate(x int, y int)
    RotateToAngle(angle IAngle)
    Rotate(angle int)
    Mirror(width int)
    GetMatrix() IBasicArrayList
}

func (this *BasicGraphicsPipeline) BasicGraphicsPipeline_() *BasicGraphicsPipeline {
    return this
}


func (this *BasicGraphicsPipeline) initFields() {
    
    this.NoDecimalTrigTable = NoDecimalTrigTable_GetInstance()
}


func NewBasicGraphicsPipeline(points IBasicArrayList) *BasicGraphicsPipeline {
    this := &BasicGraphicsPipeline{}
    this.initFields()
        this.PointBasicArrayList = points
    return this
}


func (this *BasicGraphicsPipeline) SetInitMatrix(points IBasicArrayList) {
        this.PointBasicArrayList = points
}


func (this *BasicGraphicsPipeline) CreateMatrix() {
}


// throws Exception
func (this *BasicGraphicsPipeline) Translate(x int, y int) {
        var newBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = newBasicArrayList
        var size int = this.PointBasicArrayList.Size(); _ = size
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var secondPoint IGPoint; _ = secondPoint
        for index := int(0); index < size; index++ {
            secondPoint = any(this.PointBasicArrayList.BasicArrayList_().ObjectArray[index]).(IGPoint)
            if secondPoint.GetX() != 1000 {
                newBasicArrayList.Add(pointFactory.CreateXY(secondPoint.GetX() + x, secondPoint.GetY() + y))
            } else {
                newBasicArrayList.Add(secondPoint)
            }
        }
        this.PointBasicArrayList = newBasicArrayList
}


// throws Exception
func (this *BasicGraphicsPipeline) RotateToAngle(angle IAngle) {
        this.Rotate(int(angle.GetValue()))
}


// throws Exception
func (this *BasicGraphicsPipeline) Rotate(angle int) {
        var newBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = newBasicArrayList
        var secondPoint IGPoint; _ = secondPoint
        var y int64; _ = y
        var secondX int64; _ = secondX
        var secondY int64; _ = secondY
        var point IGPoint; _ = point
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var size int = this.PointBasicArrayList.Size(); _ = size
        var sin int64; _ = sin
        var cos int64; _ = cos
        for index := int(0); index < size; index++ {
            secondPoint = any(this.PointBasicArrayList.BasicArrayList_().ObjectArray[index]).(IGPoint)
            if secondPoint.GetX() != 1000 {
                sin = this.NoDecimalTrigTable.Sin(angle)
                cos = this.NoDecimalTrigTable.Cos(angle)
                y = int64(secondPoint.GetY()) * sin
                secondX = (int64(secondPoint.GetX()) * cos) - y
                secondY = (int64(secondPoint.GetX()) * sin) + (int64(secondPoint.GetY()) * cos)
                point = pointFactory.CreateXY(int(secondX) / 10000, int(secondY) / 10000)
                newBasicArrayList.Add(point)
            } else {
                newBasicArrayList.Add(secondPoint)
            }
        }
        this.PointBasicArrayList = newBasicArrayList
}


// throws Exception
func (this *BasicGraphicsPipeline) Mirror(width int) {
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var halfWidth int = (width >> 1); _ = halfWidth
        var newBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = newBasicArrayList
        var secondPoint IGPoint; _ = secondPoint
        var newX int; _ = newX
        var point IGPoint; _ = point
        var size int = this.PointBasicArrayList.Size(); _ = size
        for index := int(0); index < size; index++ {
            secondPoint = any(this.PointBasicArrayList.BasicArrayList_().ObjectArray[index]).(IGPoint)
            if secondPoint.GetX() != 1000 {
                newX = 0
                if secondPoint.GetX() > halfWidth {
                    newX = halfWidth - (secondPoint.GetX() - halfWidth)
                } else {
                    newX = halfWidth + (halfWidth - secondPoint.GetX())
                }
                point = pointFactory.CreateXY(newX, secondPoint.GetY())
                newBasicArrayList.Add(point)
            } else {
                newBasicArrayList.Add(secondPoint)
            }
        }
        this.PointBasicArrayList = newBasicArrayList
}


func (this *BasicGraphicsPipeline) GetMatrix() IBasicArrayList {
        return this.PointBasicArrayList
}

