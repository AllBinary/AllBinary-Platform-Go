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

type VectorExplosionGenerator struct {
    RANDOM *VectorExplosionType
    ROTATION *VectorExplosionType
    VectorAnimationUtil *VectorAnimationUtil
    RandomRotationFactory *RandomRotationFactory
}

var VectorExplosionGenerator_Instance *VectorExplosionGenerator = NewVectorExplosionGenerator()

func (this *VectorExplosionGenerator) initFields() {
    this.RANDOM = NewVectorExplosionType()
    this.ROTATION = NewVectorExplosionType()
    this.VectorAnimationUtil = VectorAnimationUtil_GetInstance()
    this.RandomRotationFactory = RandomRotationFactory_GetInstance()
}


func VectorExplosionGenerator_GetInstance() *VectorExplosionGenerator {
        return VectorExplosionGenerator_Instance
}


func NewVectorExplosionGenerator() *VectorExplosionGenerator {
    this := &VectorExplosionGenerator{}
    this.initFields()
    return this
}


// throws Exception
func (this *VectorExplosionGenerator) CreateList(list IBasicArrayList, howMuch int, type_ *VectorExplosionType) IBasicArrayList {
        var point IGPoint; _ = point
        var size int = list.Size(); _ = size
        var points [][]int = func() [][]int { array := make([][]int, size); for index := range array { array[index] = make([]int, 2) }; return array }(); _ = points
        for index := int(0); index < size; index++ {
            point = any(list.BasicArrayList_().ObjectArray[index]).(IGPoint)
            points[index][0] = point.GetX()
            points[index][1] = point.GetY()
        }
        return this.CreateListFromPoints(list, points, howMuch, type_)
}


// throws Exception
func (this *VectorExplosionGenerator) GetInstance(points [][]int, howMuch int, type_ *VectorExplosionType) [][][]int {
        var pointsBasicArrayList IBasicArrayList = this.GetInstanceStartFrame(points, howMuch, type_, true); _ = pointsBasicArrayList
        var tempBasicArrayList IBasicArrayList = any(pointsBasicArrayList.BasicArrayList_().ObjectArray[0]).(IBasicArrayList); _ = tempBasicArrayList
        var newPoints [][][]int = this.VectorAnimationUtil.ToAnimationArrayFromListOfPointListWithPointsPerFrame(pointsBasicArrayList, tempBasicArrayList.Size()); _ = newPoints
        return newPoints
}


// throws Exception
func (this *VectorExplosionGenerator) GetInstanceStartFrame(points [][]int, howMuch int, type_ *VectorExplosionType, startFrame bool) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var pointsBasicArrayList IBasicArrayList = NewBasicArrayListS(howMuch); _ = pointsBasicArrayList
            pointsBasicArrayList.Add(this.CreatePointsBasicArrayList(points))
            var frameIndex int = 0; _ = frameIndex
            var tempBasicArrayList IBasicArrayList; _ = tempBasicArrayList
            var pointBasicArrayList IBasicArrayList; _ = pointBasicArrayList
            for frameIndex < howMuch {
                tempBasicArrayList = any(pointsBasicArrayList.BasicArrayList_().ObjectArray[frameIndex]).(IBasicArrayList)
                pointBasicArrayList = this.CreateListFromPoints(tempBasicArrayList, points, howMuch, type_)
                pointsBasicArrayList.Add(pointBasicArrayList)
                frameIndex++
            }
            if !startFrame {
                pointsBasicArrayList.RemoveAt(0)
            }
            tryControl1, tryResult1 = 1, pointsBasicArrayList; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *VectorExplosionGenerator) CreateListFromPoints(tempBasicArrayList IBasicArrayList, points [][]int, howMuch int, type_ *VectorExplosionType) IBasicArrayList {
        var index int = 0; _ = index
        var pointBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = pointBasicArrayList
        var sectionBasicArrayList IBasicArrayList; _ = sectionBasicArrayList
        for index < len(points) {
            sectionBasicArrayList = NewBasicArrayListD()
            for points[index][0] != 1000 {
                sectionBasicArrayList.Add(tempBasicArrayList.BasicArrayList_().ObjectArray[index])
                index++
                if index >= len(points) {
                    break
                }
            }
            sectionBasicArrayList = RandomTranslation_GetInstance(sectionBasicArrayList, howMuch)
            if type_ == this.ROTATION {
                sectionBasicArrayList = this.RandomRotationFactory.GetInstanceList(sectionBasicArrayList, howMuch)
            }
            var size int = sectionBasicArrayList.Size(); _ = size
            for index2 := int(0); index2 < size; index2++ {
                pointBasicArrayList.Add(sectionBasicArrayList.BasicArrayList_().ObjectArray[index2])
            }
            if index >= len(points) {
                break
            } else if points[index][0] == 1000 {
                pointBasicArrayList.Add(tempBasicArrayList.BasicArrayList_().ObjectArray[index])
                index++
            }
        }
        return pointBasicArrayList
}


// throws Exception
func (this *VectorExplosionGenerator) CreatePointsBasicArrayList(points [][]int) IBasicArrayList {
        if len(points) == 0 {
            panic(errors.New(fmt.Sprint("Not Points Provided")))
        }
        var firstPointBasicArrayList IBasicArrayList = NewBasicArrayListS(len(points)); _ = firstPointBasicArrayList
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        for index := int(0); index < len(points); index++ {
            firstPointBasicArrayList.Add(pointFactory.CreateXY(points[index][0], points[index][1]))
        }
        return firstPointBasicArrayList
}

