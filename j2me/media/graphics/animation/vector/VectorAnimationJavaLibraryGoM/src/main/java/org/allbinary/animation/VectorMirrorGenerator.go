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

type VectorMirrorGenerator struct {
    LogUtil *LogUtil
    VectorAnimationUtil *VectorAnimationUtil
}

var VectorMirrorGenerator_Instance *VectorMirrorGenerator = NewVectorMirrorGenerator()

func (this *VectorMirrorGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.VectorAnimationUtil = VectorAnimationUtil_GetInstance()
}


func VectorMirrorGenerator_GetInstance() *VectorMirrorGenerator {
        return VectorMirrorGenerator_Instance
}


func NewVectorMirrorGenerator() *VectorMirrorGenerator {
    this := &VectorMirrorGenerator{}
    this.initFields()
    return this
}


func (this *VectorMirrorGenerator) GetInstance_1(points [][][]int, width int) [][][]int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 [][][]int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY_ARRAY; return
                }
            }()
            var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
            var pointBasicArrayList IBasicArrayList; _ = pointBasicArrayList
            var nextPoints [][]int; _ = nextPoints
            var size2 int; _ = size2
            var size int = len(points); _ = size
            var pointsBasicArrayList IBasicArrayList = NewBasicArrayListS(size); _ = pointsBasicArrayList
            for frame := int(0); frame < size; frame++ {
                nextPoints = points[frame]
                size2 = len(nextPoints)
                pointBasicArrayList = NewBasicArrayListS(size2)
                for index := int(0); index < size2; index++ {
                    pointBasicArrayList.Add(pointFactory.CreateXY(nextPoints[index][0], nextPoints[index][1]))
                }
                pointsBasicArrayList.Add(pointBasicArrayList)
            }
            for index := int(0); index < len(points); index++ {
                pointsBasicArrayList.Add(this.GetInstance(any(pointsBasicArrayList.BasicArrayList_().ObjectArray[index]).(IBasicArrayList), width))
            }
            var newPoints [][][]int = this.VectorAnimationUtil.ToAnimationArrayFromListOfPointList(pointsBasicArrayList); _ = newPoints
            tryControl1, tryResult1 = 1, newPoints; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of VectorMirrorGenerator.getInstance
func (this *VectorMirrorGenerator) GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[[][][]int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.GetInstance_1(Java_Arg[[][][]int](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 2 && Java_IsArg[IBasicArrayList](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.GetInstance_2(Java_Arg[IBasicArrayList](args[0]), Java_Arg[int](args[1]))
        }
    }
    panic("No Java overload of VectorMirrorGenerator.getInstance for the arguments")
}


// throws Exception
func (this *VectorMirrorGenerator) GetInstance_2(pointBasicArrayList IBasicArrayList, width int) IBasicArrayList {
        var graphicsPipe IBasicGraphicsPipeline = NewBasicGraphicsPipeline(pointBasicArrayList); _ = graphicsPipe
        graphicsPipe.CreateMatrix()
        graphicsPipe.Mirror(width)
        return graphicsPipe.GetMatrix()
}

