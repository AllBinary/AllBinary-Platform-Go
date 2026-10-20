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

type GeographicMapTopViewGameLayerBehavior1 struct {
    GeographicMapTopViewLayerBehavior
    LogUtil *LogUtil
    GravityUtil *GravityUtil
    GeographicMapBehavior *MultiGeographicMapBehavior
    AutoStepBlocks bool
    OffsetY int
    PreviousGeographicMapCellPosition IGeographicMapCellPosition
}


func (this *GeographicMapTopViewGameLayerBehavior1) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GravityUtil = GravityUtil_GetInstance()
    this.GeographicMapBehavior = NewMultiGeographicMapBehavior()
    
    
    this.PreviousGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


func NewGeographicMapTopViewGameLayerBehavior1(maxGravityActionIndex int, autoStepBlocks bool, offsetY int) *GeographicMapTopViewGameLayerBehavior1 {
    this := &GeographicMapTopViewGameLayerBehavior1{}
        this.GeographicMapTopViewLayerBehavior = *NewGeographicMapTopViewLayerBehavior(maxGravityActionIndex)
        this.initFields()
        this.AutoStepBlocks = autoStepBlocks
        this.OffsetY = offsetY
    return this
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) Gravity(velocityProperties *VelocityProperties, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition) {
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            this.GeographicMapBehavior.GetCellTypeAt(geographicMapInterfaceArray, geographicMapCellTypeArray, geographicMapCellPosition)
            var hasSolidBlock bool = this.HasSolidBlock(geographicMapInterfaceArray, geographicMapCellTypeArray); _ = hasSolidBlock
            if !hasSolidBlock {
                this.GravityUtil.Process(velocityProperties, this.GravityUtil.GAME_GRAVITY_VELOCITY)
                velocityProperties.LimitXYToForwardAndReverseMaxVelocity()
                this.Gravity()
            } else {
            }
        }
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) GetPosition(geographicMapInterfaceArray []IBasicGeographicMap, layer IAllBinaryLayer, x int, y int) IGeographicMapCellPosition {
        var xCellPosition int = layer.GetXP() + -x + layer.GetWidth(); _ = xCellPosition
        var yCellPosition int = layer.GetYP() + -y + layer.GetHeight(); _ = yCellPosition
        return geographicMapInterfaceArray[0].GetCellPositionAtXYNoThrow(xCellPosition, yCellPosition)
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) GetLeftPosition(geographicMapInterfaceArray []IBasicGeographicMap, layer IAllBinaryLayer) IGeographicMapCellPosition {
        var xCellPosition int = layer.GetXP(); _ = xCellPosition
        var yCellPosition int = layer.GetYP() + layer.GetHeight(); _ = yCellPosition
        return geographicMapInterfaceArray[0].GetCellPositionAtXYNoThrow(xCellPosition, yCellPosition)
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) GetRightPosition(geographicMapInterfaceArray []IBasicGeographicMap, layer IAllBinaryLayer) IGeographicMapCellPosition {
        var xCellPosition int = layer.GetXP() + layer.GetWidth(); _ = xCellPosition
        var yCellPosition int = layer.GetYP() + layer.GetHeight(); _ = yCellPosition
        return geographicMapInterfaceArray[0].GetCellPositionAtXYNoThrow(xCellPosition, yCellPosition)
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) GetGeographicMapCellPositionIfNotSolidBlockOrOffMapLocation(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) IGeographicMapCellPosition {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetPosition(geographicMapInterfaceArray, layer, x, y); _ = geographicMapCellPosition
        if this.PreviousGeographicMapCellPosition != geographicMapCellPosition && geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var possibleStepGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPosition; _ = possibleStepGeographicMapCellPosition
            var tiledLayer IAllBinaryTiledLayer = geographicMapInterfaceArray[0].GetAllBinaryTiledLayer(); _ = tiledLayer
            if possibleStepGeographicMapCellPosition.GetColumn() > 0 && possibleStepGeographicMapCellPosition.GetRow() > 0 && possibleStepGeographicMapCellPosition.GetColumn() < tiledLayer.GetColumns() && possibleStepGeographicMapCellPosition.GetRow() < tiledLayer.GetRows() {
                this.GeographicMapBehavior.GetCellTypeAt(geographicMapInterfaceArray, geographicMapCellTypeArray, possibleStepGeographicMapCellPosition)
                var hasSolidBlock bool = this.HasSolidBlock(geographicMapInterfaceArray, geographicMapCellTypeArray); _ = hasSolidBlock
                if hasSolidBlock {
                    velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
                    this.PreviousGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                    return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                } else {
                    this.PreviousGeographicMapCellPosition = possibleStepGeographicMapCellPosition
                    return possibleStepGeographicMapCellPosition
                }
            }
        }
        this.PreviousGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
        return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) MoveAndLand(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) {
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var topViewCharacterInterface TopViewCharacterInterface = (any(layer).(TopViewCharacterInterface)); _ = topViewCharacterInterface
            topViewCharacterInterface.TerrainMove(geographicMapInterfaceArray, geographicMapCellTypeArray, x, y)
        } else {
        }
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) Move(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer, x int, y int) bool {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetGeographicMapCellPositionIfNotSolidBlockOrOffMapLocation(geographicMapInterfaceArray, geographicMapCellTypeArray, velocityProperties, layer, x, y); _ = geographicMapCellPosition
        this.MoveAndLand(geographicMapInterfaceArray, geographicMapCellTypeArray, geographicMapCellPosition, velocityProperties, layer, x, y)
        if geographicMapCellPosition == SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            return false
        } else {
            return true
        }
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) Left(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer) {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetLeftPosition(geographicMapInterfaceArray, layer); _ = geographicMapCellPosition
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var possibleStepGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterfaceArray[0].GetGeographicMapCellPositionFactory().GetAt(geographicMapCellPosition.GetColumn(), geographicMapCellPosition.GetRow() - 1); _ = possibleStepGeographicMapCellPosition
            this.GeographicMapBehavior.GetCellTypeAt(geographicMapInterfaceArray, geographicMapCellTypeArray, possibleStepGeographicMapCellPosition)
            var hasSolidBlock bool = this.HasSolidBlock(geographicMapInterfaceArray, geographicMapCellTypeArray); _ = hasSolidBlock
            if hasSolidBlock {
                if this.AutoStepBlocks {
                    var topViewCharacterInterface TopViewCharacterInterface = (any(layer).(TopViewCharacterInterface)); _ = topViewCharacterInterface
                    topViewCharacterInterface.Leftp()
                } else {
                    velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
                }
            } else {
                var topViewCharacterInterface TopViewCharacterInterface = (any(layer).(TopViewCharacterInterface)); _ = topViewCharacterInterface
                topViewCharacterInterface.Leftp()
            }
        }
}


// throws Exception
func (this *GeographicMapTopViewGameLayerBehavior1) Right(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, velocityProperties *VelocityProperties, layer IAllBinaryLayer) {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetRightPosition(geographicMapInterfaceArray, layer); _ = geographicMapCellPosition
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var possibleStepGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterfaceArray[0].GetGeographicMapCellPositionFactory().GetAt(geographicMapCellPosition.GetColumn(), geographicMapCellPosition.GetRow() - 1); _ = possibleStepGeographicMapCellPosition
            this.GeographicMapBehavior.GetCellTypeAt(geographicMapInterfaceArray, geographicMapCellTypeArray, possibleStepGeographicMapCellPosition)
            var hasSolidBlock bool = this.HasSolidBlock(geographicMapInterfaceArray, geographicMapCellTypeArray); _ = hasSolidBlock
            if hasSolidBlock {
                if this.AutoStepBlocks {
                    var topViewCharacterInterface TopViewCharacterInterface = (any(layer).(TopViewCharacterInterface)); _ = topViewCharacterInterface
                    topViewCharacterInterface.Rightp()
                } else {
                    velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
                }
            } else {
                var topViewCharacterInterface TopViewCharacterInterface = (any(layer).(TopViewCharacterInterface)); _ = topViewCharacterInterface
                topViewCharacterInterface.Rightp()
            }
        }
}

