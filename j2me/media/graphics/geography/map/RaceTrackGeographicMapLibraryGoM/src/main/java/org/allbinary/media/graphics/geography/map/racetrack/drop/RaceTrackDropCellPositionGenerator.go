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

type RaceTrackDropCellPositionGenerator struct {
    BaseDropCellPositionGenerator
    List IBasicArrayList
    TimeDelayHelper ITimeDelayHelper
    STRAIGHTAWAY int
    RaceTrackGeographicMap RaceTrackGeographicMapInterface
    RaceTrackGeographicMapCellTypeFactory IGeographicMapCellTypeFactory
    BasicArrayListUtil *BasicArrayListUtil
}

// The Java class RaceTrackDropCellPositionGenerator has subclasses so its variables are this interface
type IRaceTrackDropCellPositionGenerator interface {
    IBaseDropCellPositionGenerator
    RaceTrackDropCellPositionGenerator_() *RaceTrackDropCellPositionGenerator
    Init()
    IsDropAllowedAt(geographicMapCellPosition IGeographicMapCellPosition) bool
    Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap)
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    Drop(allBinaryLayerManager IAllBinaryLayerManager, index int)
}

func (this *RaceTrackDropCellPositionGenerator) RaceTrackDropCellPositionGenerator_() *RaceTrackDropCellPositionGenerator {
    return this
}


func (this *RaceTrackDropCellPositionGenerator) initFields() {
    this.List = NewBasicArrayListD()
    
    this.STRAIGHTAWAY = 4
    this.RaceTrackGeographicMap = NullRaceTrackGeographicMap_NULL_RACE_TRACK_GEOGRAPHIC_MAP
    this.RaceTrackGeographicMapCellTypeFactory = GeographicMapCellTypeFactory_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewRaceTrackDropCellPositionGenerator() *RaceTrackDropCellPositionGenerator {
    this := &RaceTrackDropCellPositionGenerator{}
    this.BaseDropCellPositionGenerator = *NewBaseDropCellPositionGenerator()
    this.initFields()
        this.TimeDelayHelper = NewTimeDelayHelper(10000)
    return this
}


func (this *RaceTrackDropCellPositionGenerator) Init() {
        this.List.Clear()
}


// throws Exception
func (this *RaceTrackDropCellPositionGenerator) IsDropAllowedAt(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var raceTrackGeographicMapCellType *RaceTrackGeographicMapCellType = any(this.RaceTrackGeographicMap.GetCellTypeAt(geographicMapCellPosition)).(*RaceTrackGeographicMapCellType); _ = raceTrackGeographicMapCellType
        var raceTrackGeographicMapCellTypeFactory IRaceTrackGeographicMapCellTypeFactory = any(this.RaceTrackGeographicMapCellTypeFactory).(IRaceTrackGeographicMapCellTypeFactory); _ = raceTrackGeographicMapCellTypeFactory
        if raceTrackGeographicMapCellType == raceTrackGeographicMapCellTypeFactory.GeographicMapCellTypeFactory_().HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE || raceTrackGeographicMapCellType == raceTrackGeographicMapCellTypeFactory.GeographicMapCellTypeFactory_().VERTICAL_STRAIGHT_ROAD_CELL_TYPE {
            return true
        }
        return false
}


// throws Exception
func (this *RaceTrackDropCellPositionGenerator) Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap) {
        this.Init()
        var baseRaceTrackGeographicMap IBaseRaceTrackGeographicMap = any(geographicMapInterface).(IBaseRaceTrackGeographicMap); _ = baseRaceTrackGeographicMap
        this.RaceTrackGeographicMap = baseRaceTrackGeographicMap
        this.RaceTrackGeographicMapCellTypeFactory = any(baseRaceTrackGeographicMap.GetGeographicMapCellTypeFactory()).(IRaceTrackGeographicMapCellTypeFactory)
        var roadGeographicMapCellHistory *GeographicMapCellHistory = RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance(); _ = roadGeographicMapCellHistory
        var trackedList IBasicArrayList = roadGeographicMapCellHistory.GetTracked(); _ = trackedList
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        var lastIndex int = trackedList.Size() - 1; _ = lastIndex
        var total int; _ = total
        var ahead int; _ = ahead
        for index := int(lastIndex); index > 0; index-- {
            geographicMapCellPosition = any(trackedList.Get(index)).(IGeographicMapCellPosition)
            total = 0
            ahead = index
            for this.IsDropAllowedAt(any(trackedList.Get(ahead)).(IGeographicMapCellPosition)) {
                total++
                if total > this.STRAIGHTAWAY {
                    this.List.Add(geographicMapCellPosition)
                    break
                }
                ahead++
                if ahead > lastIndex {
                    break
                }
            }
        }
        var size int = this.List.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            this.Drop(allBinaryGameLayerManager, index)
        }
}


// throws Exception
func (this *RaceTrackDropCellPositionGenerator) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.TimeDelayHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            var index int = this.BasicArrayListUtil.GetRandomIndex(this.List); _ = index
            this.Drop(allBinaryLayerManager, index)
        }
}


// throws Exception
func (this *RaceTrackDropCellPositionGenerator) Drop(allBinaryLayerManager IAllBinaryLayerManager, index int) {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}

