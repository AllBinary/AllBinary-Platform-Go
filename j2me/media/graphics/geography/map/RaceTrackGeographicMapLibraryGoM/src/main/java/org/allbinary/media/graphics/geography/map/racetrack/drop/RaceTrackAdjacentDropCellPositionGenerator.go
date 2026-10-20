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

type RaceTrackAdjacentDropCellPositionGenerator struct {
    RaceTrackDropCellPositionGenerator
    LayerCoveringCellPositionsUtil *LayerCoveringCellPositionsUtil
    GroupCommonFactory *GroupCommonFactory
    SurroundingCellPositions []IGeographicMapCellPosition
    Hashtable *ABHashtable[any, any]
    GeographicMapDirectionUtil *GeographicMapDirectionUtil
}

var RaceTrackAdjacentDropCellPositionGenerator_SINGLETON *RaceTrackAdjacentDropCellPositionGenerator = NewRaceTrackAdjacentDropCellPositionGenerator()

func (this *RaceTrackAdjacentDropCellPositionGenerator) initFields() {
    this.LayerCoveringCellPositionsUtil = LayerCoveringCellPositionsUtil_GetInstance()
    this.GroupCommonFactory = GroupCommonFactory_GetInstance()
    this.SurroundingCellPositions = make([]IGeographicMapCellPosition, 4)
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    this.GeographicMapDirectionUtil = GeographicMapDirectionUtil_GetInstance()
}


func RaceTrackAdjacentDropCellPositionGenerator_GetInstance() DropCellPositionGeneratorInterface {
        return RaceTrackAdjacentDropCellPositionGenerator_SINGLETON
}


func NewRaceTrackAdjacentDropCellPositionGenerator() *RaceTrackAdjacentDropCellPositionGenerator {
    this := &RaceTrackAdjacentDropCellPositionGenerator{}
    this.RaceTrackDropCellPositionGenerator = *NewRaceTrackDropCellPositionGenerator()
    this.initFields()
    return this
}


// throws Exception
func (this *RaceTrackAdjacentDropCellPositionGenerator) GetFirstNonRoadAdjacentCellPosition(column int, row int) IGeographicMapCellPosition {
        var nonRoadGeographicMapCellPosition IGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION; _ = nonRoadGeographicMapCellPosition
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.RaceTrackGeographicMap.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        this.SurroundingCellPositions[0] = geographicMapCellPositionFactory.GetAt(column, row - 1)
        this.SurroundingCellPositions[1] = geographicMapCellPositionFactory.GetAt(column, row + 1)
        this.SurroundingCellPositions[2] = geographicMapCellPositionFactory.GetAt(column - 1, row)
        this.SurroundingCellPositions[3] = geographicMapCellPositionFactory.GetAt(column + 1, row)
        var baseRaceTrackGeographicMap IBaseRaceTrackGeographicMap = any(this.RaceTrackGeographicMap).(IBaseRaceTrackGeographicMap); _ = baseRaceTrackGeographicMap
        var raceTrackGeographicMapCellTypeFactory IRaceTrackGeographicMapCellTypeFactory = any(baseRaceTrackGeographicMap.GetGeographicMapCellTypeFactory()).(IRaceTrackGeographicMapCellTypeFactory); _ = raceTrackGeographicMapCellTypeFactory
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        var raceTrackGeographicMapCellType *RaceTrackGeographicMapCellType; _ = raceTrackGeographicMapCellType
        for index := int(len(this.SurroundingCellPositions)); Java_PreDecrement(&index) >= 0;  {
            geographicMapCellPosition = this.SurroundingCellPositions[index]
            raceTrackGeographicMapCellType = any(this.RaceTrackGeographicMap.GetCellTypeAt(geographicMapCellPosition)).(*RaceTrackGeographicMapCellType)
            if !raceTrackGeographicMapCellTypeFactory.IsPath(raceTrackGeographicMapCellType) {
                nonRoadGeographicMapCellPosition = geographicMapCellPosition
                break
            }
        }
        return nonRoadGeographicMapCellPosition
}


// throws Exception
func (this *RaceTrackAdjacentDropCellPositionGenerator) Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap) {
        this.Hashtable.Put(AllBinaryGameLayerManager_ID, allBinaryGameLayerManager)
        this.RaceTrackDropCellPositionGenerator.Update(allBinaryGameLayerManager, geographicMapInterface)
}


// throws Exception
func (this *RaceTrackAdjacentDropCellPositionGenerator) Drop(allBinaryLayerManager IAllBinaryLayerManager, index int) {
        var geographicMapCellPosition IGeographicMapCellPosition = any(this.List.Get(index)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
        var randomGeographicMapCellPosition IGeographicMapCellPosition = this.GetFirstNonRoadAdjacentCellPosition(geographicMapCellPosition.GetColumn(), geographicMapCellPosition.GetRow()); _ = randomGeographicMapCellPosition
        if randomGeographicMapCellPosition == nil {
            return
        }
        var dropCellPositionHistory *DropCellPositionHistory = DropCellPositionHistory_GetInstance(); _ = dropCellPositionHistory
        if !dropCellPositionHistory.IsCellPositionWithDrop(geographicMapCellPosition) {
            var direction *Direction = this.GeographicMapDirectionUtil.GetDirectionFromCellPositionToAdjacentCellPosition(randomGeographicMapCellPosition, geographicMapCellPosition); _ = direction
            var objectCanBeNull any = this.Hashtable.Get(this.GroupCommonFactory.ID); _ = objectCanBeNull
            if objectCanBeNull != nil {
            } else {
                this.Hashtable.Put(this.GroupCommonFactory.ID, BasicGroupFactory_GetInstance().NONE_ARRAY)
            }
            this.Hashtable.Put(DirectionFactory_GetInstance().NAME, direction)
            var point IGPoint = randomGeographicMapCellPosition.GetPoint(); _ = point
            var x int = point.GetX(); _ = x
            var y int = point.GetY(); _ = y
            var z int = point.GetZ(); _ = z
            var layerInterface IAllBinaryLayer = RaceTrackAdjacentDropLayerFactory_GetInstance().GetRandomInstance().GetNextInstance(this.Hashtable, x, y, z); _ = layerInterface
            var baseRaceTrackGeographicMap IBaseRaceTrackGeographicMap = any(this.RaceTrackGeographicMap).(IBaseRaceTrackGeographicMap); _ = baseRaceTrackGeographicMap
            var list IBasicArrayList = this.LayerCoveringCellPositionsUtil.GetAll(baseRaceTrackGeographicMap, randomGeographicMapCellPosition, layerInterface, NewBasicArrayListD()); _ = list
            if dropCellPositionHistory.AnyCellPositionWithDrop(list) {
                return
            }
            var geographicMapCellType IGeographicMapCellType; _ = geographicMapCellType
            for index2 := int(list.Size()); Java_PreDecrement(&index2) >= 0;  {
                geographicMapCellType = this.RaceTrackGeographicMap.GetCellTypeAt(any(list.Get(index2)).(IGeographicMapCellPosition))
                if this.RaceTrackGeographicMapCellTypeFactory.IsPath(geographicMapCellType) {
                    return
                }
            }
            dropCellPositionHistory.AddAll(list, layerInterface)
            allBinaryLayerManager.Append(layerInterface)
        }
}

