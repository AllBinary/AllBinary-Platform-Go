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

type NonPlayerTopViewCharacterBehavior struct {
    TopViewCharacterBehavior
    GeographicMapBehavior *MultiGeographicMapBehavior
    CLIFF *BasicTerrainInfo
}


func (this *NonPlayerTopViewCharacterBehavior) initFields() {
    this.GeographicMapBehavior = NewMultiGeographicMapBehavior()
    this.CLIFF = NewBasicTerrainInfo(AngleFactory_GetInstance().DOWN)
}

func NewNonPlayerTopViewCharacterBehavior() *NonPlayerTopViewCharacterBehavior {
    this := &NonPlayerTopViewCharacterBehavior{}
    this.TopViewCharacterBehavior = *NewTopViewCharacterBehavior()
    this.initFields()
    return this
}


// throws Exception
func (this *NonPlayerTopViewCharacterBehavior) TerrainEvent(layer IAllBinaryLayer, direction *Direction, x int, y int, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType, geographicMapCellPosition IGeographicMapCellPosition) {
        var maxColumns int = geographicMapInterfaceArray[0].GetAllBinaryTiledLayer().GetColumns(); _ = maxColumns
        if geographicMapCellPosition.GetColumn() > 0 && geographicMapCellPosition.GetColumn() < maxColumns {
            var nextTerrainGeographicMapCellPosition IGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION; _ = nextTerrainGeographicMapCellPosition
            var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterfaceArray[0].GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
            if direction == DirectionFactory_GetInstance().LEFT {
                nextTerrainGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(geographicMapCellPosition.GetColumn() - 1, geographicMapCellPosition.GetRow())
            } else if direction == DirectionFactory_GetInstance().RIGHT {
                nextTerrainGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(geographicMapCellPosition.GetColumn() + 1, geographicMapCellPosition.GetRow())
            }
            this.GeographicMapBehavior.GetCellTypeAt(geographicMapInterfaceArray, geographicMapCellTypeArray, nextTerrainGeographicMapCellPosition)
            var hasSolidBlock bool = this.HasSolidBlock(geographicMapInterfaceArray, geographicMapCellTypeArray); _ = hasSolidBlock
            if !hasSolidBlock {
                var terrainEvent *TerrainEvent = TerrainEventCircularStaticPool_GetInstance().GetNext(this.CLIFF); _ = terrainEvent
                TerrainEventHandler_GetInstance(layer).FireEvent(terrainEvent)
            }
        }
}


func (this *NonPlayerTopViewCharacterBehavior) TerrainMove(layer IAllBinaryLayer, geographicMapInterfaceArray []IBasicGeographicMap, x int, y int) {
        layer.MoveDXY(x, y)
}

