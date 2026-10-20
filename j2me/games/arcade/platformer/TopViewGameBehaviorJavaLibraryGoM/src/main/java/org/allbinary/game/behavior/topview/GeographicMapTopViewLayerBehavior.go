/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type GeographicMapTopViewLayerBehavior struct {
    TopViewGameLayerBehavior
}

// The Java class GeographicMapTopViewLayerBehavior has subclasses so its variables are this interface
type IGeographicMapTopViewLayerBehavior interface {
    ITopViewGameLayerBehavior
    GeographicMapTopViewLayerBehavior_() *GeographicMapTopViewLayerBehavior
    GetGeographicMapCellPositionIfNotSolidBlockOrOffMapLocation(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) IGeographicMapCellPosition
    GetGeographicMapCellPositionIfNotSolidBlockOrOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition, velocityProperties *VelocityProperties, layer IAllBinaryLayer) IGeographicMapCellPosition
    GetGeographicMapCellPositionFromListIfNotSolidBlockOrOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPositionList IBasicArrayList, velocityProperties *VelocityProperties, layer IAllBinaryLayer) IGeographicMapCellPosition
    Gravity(velocityProperties *VelocityProperties, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition)
    Left(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer)
    Move(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) bool
    MoveAndLand(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int)
    Right(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer)
    HasSolidBlock(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool
    IsOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool
}

func (this *GeographicMapTopViewLayerBehavior) GeographicMapTopViewLayerBehavior_() *GeographicMapTopViewLayerBehavior {
    return this
}



func NewGeographicMapTopViewLayerBehavior(maxGravityActionIndex int) *GeographicMapTopViewLayerBehavior {
    this := &GeographicMapTopViewLayerBehavior{}
        this.TopViewGameLayerBehavior = *NewTopViewGameLayerBehavior(maxGravityActionIndex)
    return this
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) GetGeographicMapCellPositionIfNotSolidBlockOrOffMapLocation(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) IGeographicMapCellPosition {
        return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) GetGeographicMapCellPositionIfNotSolidBlockOrOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition, velocityProperties *VelocityProperties, layer IAllBinaryLayer) IGeographicMapCellPosition {
        return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) GetGeographicMapCellPositionFromListIfNotSolidBlockOrOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPositionList IBasicArrayList, velocityProperties *VelocityProperties, layer IAllBinaryLayer) IGeographicMapCellPosition {
        return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) Gravity(velocityProperties *VelocityProperties, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition) {
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) Left(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer) {
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) Move(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) bool {
        return false
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) MoveAndLand(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) {
}


// throws Exception
func (this *GeographicMapTopViewLayerBehavior) Right(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer) {
}


func (this *GeographicMapTopViewLayerBehavior) HasSolidBlock(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool {
        var size int = len(geographicMapInterfaceArray); _ = size
        var basicTopViewGeographicMapCellTypeFactory *BasicTopViewGeographicMapCellTypeFactory; _ = basicTopViewGeographicMapCellTypeFactory
        for index := int(0); index < size; index++ {
            basicTopViewGeographicMapCellTypeFactory = any(geographicMapInterfaceArray[index].GetGeographicMapCellTypeFactory()).(*BasicTopViewGeographicMapCellTypeFactory)
            if basicTopViewGeographicMapCellTypeFactory.BLOCK_CELL_TYPE.IsType(geographicMapCellTypeArray[index]) {
                return true
            }
        }
        return false
}


func (this *GeographicMapTopViewLayerBehavior) IsOffMap(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool {
        var size int = len(geographicMapInterfaceArray); _ = size
        var basicTopViewGeographicMapCellTypeFactory *BasicTopViewGeographicMapCellTypeFactory; _ = basicTopViewGeographicMapCellTypeFactory
        for index := int(0); index < size; index++ {
            basicTopViewGeographicMapCellTypeFactory = any(geographicMapInterfaceArray[index].GetGeographicMapCellTypeFactory()).(*BasicTopViewGeographicMapCellTypeFactory)
            if basicTopViewGeographicMapCellTypeFactory.OFF_MAP_CELL_TYPE.IsType(geographicMapCellTypeArray[index]) {
                return true
            }
        }
        return false
}

