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

type BaseRaceTrackGeographicMap struct {
    BasicGeographicMap
    LogUtil *LogUtil
    RaceTrackInfo *RaceTrackInfo
    RaceTrackData *RaceTrackData
    MiniGeographicMap IBasicGeographicMap
}

// Implements RaceTrackGeographicMapInterface

// The Java class BaseRaceTrackGeographicMap has subclasses so its variables are this interface
type IBaseRaceTrackGeographicMap interface {
    IBasicGeographicMap
    BaseRaceTrackGeographicMap_() *BaseRaceTrackGeographicMap
    GetRaceTrackInfo() *RaceTrackInfo
    SetRaceTrackInfo(aRaceTrackInfo *RaceTrackInfo)
    SetMiniBasicGeographicMap(miniBasicGeographicMap IBasicGeographicMap)
    GetMiniBasicGeographicMap() IBasicGeographicMap
    GetRaceTrackData() *RaceTrackData
    SetRaceTrackData(raceTrackData *RaceTrackData)
}

func (this *BaseRaceTrackGeographicMap) BaseRaceTrackGeographicMap_() *BaseRaceTrackGeographicMap {
    return this
}


func (this *BaseRaceTrackGeographicMap) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewBaseRaceTrackGeographicMap(raceTrackInfo *RaceTrackInfo, raceTrackData *RaceTrackData, tiledLayer IAllBinaryTiledLayer, miniGeographicMap IBasicGeographicMap, geographicMapCellPositionFactoryInterface GeographicMapCellPositionFactoryInterface, geographicMapCellPositionBaseFactory IGeographicMapCellPositionBaseFactory, geographicMapCellTypeFactory IGeographicMapCellTypeFactory) *BaseRaceTrackGeographicMap {
    this := &BaseRaceTrackGeographicMap{}
        this.BasicGeographicMap = *NewBasicGeographicMap(raceTrackInfo.GetId(), raceTrackInfo.GetName(), raceTrackData.GetCellTypeIdToGeographicMapCellTypeArray(), tiledLayer, raceTrackInfo.GetForegroundBasicColor(), raceTrackInfo.GetBackgroundBasicColor(), geographicMapCellPositionFactoryInterface, geographicMapCellPositionBaseFactory, geographicMapCellTypeFactory)
        this.initFields()
        this.RaceTrackInfo = raceTrackInfo
        this.RaceTrackData = raceTrackData
        this.MiniGeographicMap = miniGeographicMap
        var miniTiledLayer IAllBinaryTiledLayer = this.MiniGeographicMap.GetAllBinaryTiledLayer(); _ = miniTiledLayer
        if miniTiledLayer.GetColumns() != this.GetGeographicMapCellPositionFactory().GetColumns() {
            var error_ string = fmt.Sprint(NewStringMaker().Append("RaceTrackMap has incorrect Mini Map columns: ").Appendint(miniTiledLayer.GetColumns()).Append(" != ").Appendint(this.GetGeographicMapCellPositionFactory().GetColumns())); _ = error_
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR)
            panic(errors.New(fmt.Sprint(error_)))
        }
        if miniTiledLayer.GetRows() != this.GetGeographicMapCellPositionFactory().GetRows() {
            var error_ string = "RaceTrackMap has incorrect Mini Map rows"; _ = error_
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR)
            panic(errors.New(fmt.Sprint(error_)))
        }
    return this
}


func (this *BaseRaceTrackGeographicMap) GetRaceTrackInfo() *RaceTrackInfo {
        return this.RaceTrackInfo
}


func (this *BaseRaceTrackGeographicMap) SetRaceTrackInfo(aRaceTrackInfo *RaceTrackInfo) {
        this.RaceTrackInfo = aRaceTrackInfo
}


func (this *BaseRaceTrackGeographicMap) SetMiniBasicGeographicMap(miniBasicGeographicMap IBasicGeographicMap) {
        this.MiniGeographicMap = miniBasicGeographicMap
}


func (this *BaseRaceTrackGeographicMap) GetMiniBasicGeographicMap() IBasicGeographicMap {
        return this.MiniGeographicMap
}


func (this *BaseRaceTrackGeographicMap) GetRaceTrackData() *RaceTrackData {
        return this.RaceTrackData
}


func (this *BaseRaceTrackGeographicMap) SetRaceTrackData(raceTrackData *RaceTrackData) {
        this.RaceTrackData = raceTrackData
}

