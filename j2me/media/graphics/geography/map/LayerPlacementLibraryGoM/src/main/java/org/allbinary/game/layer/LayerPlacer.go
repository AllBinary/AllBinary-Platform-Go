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

type LayerPlacer struct {
    EnumerationUtil *EnumerationUtil
    LayerInterfaceVisitor LayerInterfaceVisitor
    Dimension IGPoint
}


func (this *LayerPlacer) initFields() {
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    
    
}


func NewLayerPlacer(layerInterfaceVisitor LayerInterfaceVisitor, dimension IGPoint) *LayerPlacer {
    this := &LayerPlacer{}
    this.initFields()
        this.LayerInterfaceVisitor = layerInterfaceVisitor
        this.Dimension = dimension
    return this
}


// throws Exception
func (this *LayerPlacer) ProcessList(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        var layerPlacementInterface LayerPlacementInterface; _ = layerPlacementInterface
        for index := int(0); index < size; index++ {
            layerPlacementInterface = any(list.BasicArrayList_().ObjectArray[index]).(LayerPlacementInterface)
            this.Process(layerPlacementInterface)
        }
}


// throws Exception
func (this *LayerPlacer) Process(layerPlacementInterface LayerPlacementInterface) {
        var relativePoint IGPoint = this.GetPoint(layerPlacementInterface); _ = relativePoint
        var layerInterfaceFactory *LayerInterfaceFactory = LayerInterfaceFactory_GetInstance(); _ = layerInterfaceFactory
        var hashtable *ABHashtable[any, any] = layerPlacementInterface.GetInstance(); _ = hashtable
        var enumeration Enumeration = hashtable.Keys(); _ = enumeration
        var point IGPoint; _ = point
        var layerHashtable *ABHashtable[any, any]; _ = layerHashtable
        var x int; _ = x
        var y int; _ = y
        var z int; _ = z
        var layerInterface IAllBinaryLayer; _ = layerInterface
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            point = any(this.EnumerationUtil.NextElement(enumeration)).(IGPoint)
            layerHashtable = any(hashtable.Get(point)).(*ABHashtable[any, any])
            x = point.GetX() + relativePoint.GetX()
            y = point.GetY() + relativePoint.GetY()
            z = point.GetZ() + relativePoint.GetZ()
            layerInterface = layerInterfaceFactory.GetNexInstance(layerHashtable, x, y, z)
            this.LayerInterfaceVisitor.Visit(layerInterface)
        }
}


// throws Exception
func (this *LayerPlacer) GetPoint(layerPlacementInterface LayerPlacementInterface) IGPoint {
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var layerPlacementType *LayerPlacementType = layerPlacementInterface.GetLayerType(); _ = layerPlacementType
        if layerPlacementType == LayerPlacementTypeFactory_GetInstance().MAP {
            var width int = layerPlacementInterface.GetWidth(); _ = width
            var height int = layerPlacementInterface.GetHeight(); _ = height
            var x int = ((this.Dimension.GetX() - width) / 2); _ = x
            var y int = ((this.Dimension.GetY() - height) / 2); _ = y
            return pointFactory.CreateXY(x, y)
        } else if layerPlacementType == LayerPlacementTypeFactory_GetInstance().UP {
            var width int = layerPlacementInterface.GetWidth(); _ = width
            var height int = layerPlacementInterface.GetHeight(); _ = height
            var x int = ((this.Dimension.GetX() - width) / 2); _ = x
            var y int = -height; _ = y
            return pointFactory.CreateXY(x, y)
        } else {
            panic(errors.New(fmt.Sprint("PlacementType Not Recognized")))
        }
}

