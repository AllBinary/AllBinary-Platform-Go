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

type VectorRotationGenerator struct {
    LogUtil *LogUtil
    VectorAnimationUtil *VectorAnimationUtil
    AngleFactory *AngleFactory
}

var VectorRotationGenerator_Instance *VectorRotationGenerator = NewVectorRotationGenerator()

func (this *VectorRotationGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.VectorAnimationUtil = VectorAnimationUtil_GetInstance()
    this.AngleFactory = AngleFactory_GetInstance()
}


func VectorRotationGenerator_GetInstance() *VectorRotationGenerator {
        return VectorRotationGenerator_Instance
}


func NewVectorRotationGenerator() *VectorRotationGenerator {
    this := &VectorRotationGenerator{}
    this.initFields()
    return this
}


func (this *VectorRotationGenerator) GetInstance(vectorRotationInfo *VectorInfo) [][][]int {
        return this.GetInstanceFrames(vectorRotationInfo.GetWidth(), vectorRotationInfo.GetHeight(), vectorRotationInfo.GetPoints(), vectorRotationInfo.GetTotalFrames())
}


func (this *VectorRotationGenerator) GetInstanceFrames(width int, height int, points [][]int, frames int) [][][]int {
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
            var totalAngle int16 = this.AngleFactory.TOTAL_ANGLE; _ = totalAngle
            var angleIncrement int = int(totalAngle) / frames; _ = angleIncrement
            var pointsBasicArrayList IBasicArrayList = NewBasicArrayListS(int(totalAngle) / angleIncrement); _ = pointsBasicArrayList
            var pointBasicArrayList IBasicArrayList = NewBasicArrayListS(len(points)); _ = pointBasicArrayList
            var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
            var size int = len(points); _ = size
            for index := int(0); index < size; index++ {
                pointBasicArrayList.Add(pointFactory.CreateXY(points[index][0], points[index][1]))
            }
            var index int = 0; _ = index
            for index < int(totalAngle) {
                pointsBasicArrayList.Add(this.GetListAtAngle(width, height, pointBasicArrayList, this.AngleFactory.GetAt(index)))
                index += angleIncrement
            }
            var newPoints [][][]int = this.VectorAnimationUtil.ToAnimationArrayFromListOfPointListWithPointsPerFrame(pointsBasicArrayList, pointBasicArrayList.Size()); _ = newPoints
            tryControl1, tryResult1 = 1, newPoints; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *VectorRotationGenerator) GetListAtAngle(width int, height int, pointBasicArrayList IBasicArrayList, angle IAngle) IBasicArrayList {
        var graphicsPipe IBasicGraphicsPipeline = NewBasicGraphicsPipeline(pointBasicArrayList); _ = graphicsPipe
        graphicsPipe.CreateMatrix()
        graphicsPipe.Translate(-(width >> 1), -(height >> 1))
        graphicsPipe.RotateToAngle(angle)
        graphicsPipe.Translate((width >> 1), (height >> 1))
        return graphicsPipe.GetMatrix()
}

