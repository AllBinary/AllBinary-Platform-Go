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

type VectorCenterGenerator struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    MinX int
    MinY int
    MaxX int
    MaxY int
    Dx int
    Dy int
    Width int
    Height int
}


func (this *VectorCenterGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.MinX = Integer_MAX_VALUE
    this.MinY = Integer_MAX_VALUE
    this.MaxX = Integer_MIN_VALUE
    this.MaxY = Integer_MIN_VALUE
    this.Dx = 0
    this.Dy = 0
    
    
}


func NewVectorCenterGenerator() *VectorCenterGenerator {
    this := &VectorCenterGenerator{}
    this.initFields()
    return this
}


// throws Exception
func (this *VectorCenterGenerator) Calculate(hashMap *HashMap) {
        var graphicItemArray []any = hashMap.KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(hashMap.Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            if item.GetName() == LinesGraphicItem_GetStaticName() {
                this.GetInstance(item.GetPointsInterface().GetPoints())
            }
        }
        this.LogUtil.PutF("minX: " + fmt.Sprint(this.MinX) + " minY: " + fmt.Sprint(this.MinY) + " maxX: " + fmt.Sprint(this.MaxX) + " maxY: " + fmt.Sprint(this.MaxY), this, this.CommonStrings.GET_INSTANCE)
        this.SetWidth(this.MaxX - this.MinX)
        this.SetHeight(this.MaxY - this.MinY)
        var max_ int = this.GetWidth(); _ = max_
        if this.GetHeight() > max_ {
            max_ = this.GetHeight()
        }
        var middle int = max_ / 2; _ = middle
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var s string = fmt.Sprint(NewStringMaker().Append(commonLabels.WIDTH_LABEL).Appendint(this.GetWidth()).Append(commonLabels.HEIGHT_LABEL).Appendint(this.GetHeight()).Append(" max: ").Appendint(max_).Append(" middle: ").Appendint(middle)); _ = s
        this.LogUtil.PutF(s, this, this.CommonStrings.GET_INSTANCE)
        var currentMiddleX int = this.MinX + this.GetWidth() / 2; _ = currentMiddleX
        var currentMiddleY int = this.MinY + this.GetHeight() / 2; _ = currentMiddleY
        this.Dx = middle - currentMiddleX
        this.Dy = middle - currentMiddleY
}


// throws Exception
func (this *VectorCenterGenerator) Transform(hashMap *HashMap) {
        this.Calculate(hashMap)
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(" dx: ").Appendint(this.Dx).Append(" dy: ").Appendint(this.Dy)), this, this.CommonStrings.GET_INSTANCE)
        var graphicItemArray []any = hashMap.KeySet().ToArray(); _ = graphicItemArray
        var size int = len(graphicItemArray); _ = size
        for index := int(0); index < size; index++ {
            var item GraphicItemInterface = any(hashMap.Get(graphicItemArray[index])).(GraphicItemInterface); _ = item
            var list IBasicArrayList = VectorCenterGenerator_Translate(item.GetPointsInterface().GetPoints(), this.Dx, this.Dy); _ = list
            var newPoints *Points = NewPoints(); _ = newPoints
            newPoints.AddPoints(list)
            item.SetPointsInterface(newPoints)
        }
}


// throws Exception
func (this *VectorCenterGenerator) GetInstance(pointVector IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_INSTANCE, e)
            panic(e)
                }
            }()
            var size int = pointVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var point IGPoint = any(pointVector.Get(index)).(IGPoint); _ = point
                if point.GetX() < this.MinX {
                    this.MinX = point.GetX()
                }
                if point.GetX() > this.MaxX {
                    this.MaxX = point.GetX()
                }
                if point.GetY() < this.MinY {
                    this.MinY = point.GetY()
                }
                if point.GetY() > this.MaxY {
                    this.MaxY = point.GetY()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func VectorCenterGenerator_Translate(list IBasicArrayList, x int, y int) IBasicArrayList {
        var graphicsPipe IBasicGraphicsPipeline = NewBasicGraphicsPipeline(list); _ = graphicsPipe
        graphicsPipe.CreateMatrix()
        graphicsPipe.Translate(x, y)
        return graphicsPipe.GetMatrix()
}


func (this *VectorCenterGenerator) GetWidth() int {
        return this.Width
}


func (this *VectorCenterGenerator) SetWidth(width int) {
        this.Width = width
}


func (this *VectorCenterGenerator) GetHeight() int {
        return this.Height
}


func (this *VectorCenterGenerator) SetHeight(height int) {
        this.Height = height
}

