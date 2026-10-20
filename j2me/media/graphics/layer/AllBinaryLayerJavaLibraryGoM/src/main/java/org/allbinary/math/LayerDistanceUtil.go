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

type LayerDistanceUtil struct {
    MathUtil *MathUtil
}

var LayerDistanceUtil_Instance *LayerDistanceUtil = NewLayerDistanceUtil()

func (this *LayerDistanceUtil) initFields() {
    this.MathUtil = MathUtil_GetInstance()
}

func NewLayerDistanceUtil() *LayerDistanceUtil {
    this := &LayerDistanceUtil{}
    this.initFields()
    return this
}


func LayerDistanceUtil_GetInstance() *LayerDistanceUtil {
        return LayerDistanceUtil_Instance
}


func (this *LayerDistanceUtil) GetDistance(layerInterface IAllBinaryLayer, layerInterface2 IAllBinaryLayer) int {
        var dx int = (layerInterface.GetXP() + layerInterface.GetHalfWidth()) - (layerInterface2.GetXP() + layerInterface2.GetHalfWidth()); _ = dx
        var dy int = (layerInterface.GetYP() + layerInterface.GetHalfHeight()) - (layerInterface2.GetYP() + layerInterface2.GetHalfHeight()); _ = dy
        return this.MathUtil.Sqrt((dx * dx) + (dy * dy))
}


func (this *LayerDistanceUtil) GetDistance3D(layerInterface IAllBinaryLayer, layerInterface2 IAllBinaryLayer) int {
        var dx int = (layerInterface.GetXP() + layerInterface.GetHalfWidth()) - (layerInterface2.GetXP() + layerInterface2.GetHalfWidth()); _ = dx
        var dy int = (layerInterface.GetYP() + layerInterface.GetHalfHeight()) - (layerInterface2.GetYP() + layerInterface2.GetHalfHeight()); _ = dy
        var dz int = (layerInterface.GetZP() + layerInterface.GetHalfDepth()) - (layerInterface2.GetZP() + layerInterface2.GetHalfDepth()); _ = dz
        return this.MathUtil.Sqrt((dx * dx) + (dy * dy) + (dz * dz))
}


func (this *LayerDistanceUtil) GetDistanceAt(layerInterface IAllBinaryLayer, point IGPoint) int {
        var dx int = layerInterface.GetXP() - point.GetX(); _ = dx
        var dy int = layerInterface.GetYP() - point.GetY(); _ = dy
        return this.MathUtil.Sqrt((dx * dx) + (dy * dy))
}

