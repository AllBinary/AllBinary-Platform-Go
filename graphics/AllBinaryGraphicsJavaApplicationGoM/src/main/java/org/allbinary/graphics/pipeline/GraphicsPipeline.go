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

type GraphicsPipeline struct {
    BasicGraphicsPipeline
}



func NewGraphicsPipeline(points IBasicArrayList) *GraphicsPipeline {
    this := &GraphicsPipeline{}
        this.BasicGraphicsPipeline = *NewBasicGraphicsPipeline(points)
    return this
}


// throws Exception
func (this *GraphicsPipeline) RotateTheta(aTheta float64) {
        var newVector IBasicArrayList = NewBasicArrayListD(); _ = newVector
        var size int = this.PointBasicArrayList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var secondPoint IGPoint = any(this.PointBasicArrayList.Get(index)).(IGPoint); _ = secondPoint
            var secondX float64 = (secondPoint.GetX() * math.Cos(aTheta)) - (secondPoint.GetY() * math.Sin(aTheta)); _ = secondX
            var secondY float64 = (secondPoint.GetX() * math.Sin(aTheta)) + (secondPoint.GetY() * math.Cos(aTheta)); _ = secondY
            var point IGPoint = PointFactory_GetInstance().CreateXY((int(secondX)), (int(secondY))); _ = point
            newVector.Add(point)
        }
        this.PointBasicArrayList = newVector
}

