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

type VectorAnimationUtil struct {
}

var VectorAnimationUtil_Instance *VectorAnimationUtil = NewVectorAnimationUtil()


func VectorAnimationUtil_GetInstance() *VectorAnimationUtil {
        return VectorAnimationUtil_Instance
}


func NewVectorAnimationUtil() *VectorAnimationUtil {
    this := &VectorAnimationUtil{}
    return this
}


func (this *VectorAnimationUtil) ToAnimationArrayFromListOfPointListWithPointsPerFrame(vector IBasicArrayList, pointsPerFrame int) [][][]int {
        var size int = vector.Size(); _ = size
        var points [][][]int = func() [][][]int { array := make([][][]int, size); for index := range array { array[index] = make([][]int, pointsPerFrame); for index2 := range array[index] { array[index][index2] = make([]int, 2) } }; return array }(); _ = points
        var nextBasicArrayList IBasicArrayList; _ = nextBasicArrayList
        var framePoints [][]int; _ = framePoints
        for index := int(0); index < size; index++ {
            nextBasicArrayList = any(vector.BasicArrayList_().ObjectArray[index]).(IBasicArrayList)
            framePoints = this.ToFrameArrayFromPointBasicArrayList(nextBasicArrayList)
            for pointIndex := int(0); pointIndex < len(framePoints); pointIndex++ {
                points[index][pointIndex][0] = framePoints[pointIndex][0]
                points[index][pointIndex][1] = framePoints[pointIndex][1]
            }
        }
        return points
}


func (this *VectorAnimationUtil) ToAnimationArrayFromListOfPointList(vector IBasicArrayList) [][][]int {
        var size int = vector.Size(); _ = size
        var points [][][]int = func() [][][]int { array := make([][][]int, size); for index := range array { array[index] = make([][]int, 0); for index2 := range array[index] { array[index][index2] = make([]int, 0) } }; return array }(); _ = points
        var nextBasicArrayList IBasicArrayList; _ = nextBasicArrayList
        var framePoints [][]int; _ = framePoints
        for index := int(0); index < size; index++ {
            nextBasicArrayList = any(vector.BasicArrayList_().ObjectArray[index]).(IBasicArrayList)
            framePoints = this.ToFrameArrayFromPointBasicArrayList(nextBasicArrayList)
            points[index] = func() [][]int { array := make([][]int, len(framePoints)); for index := range array { array[index] = make([]int, 2) }; return array }()
            for pointIndex := int(0); pointIndex < len(framePoints); pointIndex++ {
                points[index][pointIndex][0] = framePoints[pointIndex][0]
                points[index][pointIndex][1] = framePoints[pointIndex][1]
            }
        }
        return points
}


func (this *VectorAnimationUtil) ToFrameArrayFromPointBasicArrayList(list IBasicArrayList) [][]int {
        var points [][]int = func() [][]int { array := make([][]int, list.Size()); for index := range array { array[index] = make([]int, 2) }; return array }(); _ = points
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var nextPoint IGPoint = any(list.BasicArrayList_().ObjectArray[index]).(IGPoint); _ = nextPoint
            points[index][0] = nextPoint.GetX()
            points[index][1] = nextPoint.GetY()
        }
        return points
}

