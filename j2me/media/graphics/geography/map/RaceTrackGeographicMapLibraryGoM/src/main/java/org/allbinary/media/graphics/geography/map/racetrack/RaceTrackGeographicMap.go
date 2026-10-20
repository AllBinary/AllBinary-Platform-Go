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

type RaceTrackGeographicMap struct {
    BaseRaceTrackGeographicMap
    CustomMapGenerator ICustomMapGeneratorBase
}


func (this *RaceTrackGeographicMap) initFields() {
    
}


func NewRaceTrackGeographicMap(raceTrackInfo *RaceTrackInfo, raceTrackData *RaceTrackData, tiledLayerFactoryInterface AllBinaryTiledLayerFactoryInterface, geographicMapCellPositionFactoryInterface GeographicMapCellPositionFactoryInterface, geographicMapCellPositionBaseFactory IGeographicMapCellPositionBaseFactory, geographicMapCellTypeFactory IGeographicMapCellTypeFactory, customMapGeneratorBaseFactory ICustomMapGeneratorBaseFactory) *RaceTrackGeographicMap {
    this := &RaceTrackGeographicMap{}
        this.BaseRaceTrackGeographicMap = *NewBaseRaceTrackGeographicMap(raceTrackInfo, raceTrackData, tiledLayerFactoryInterface.GetInstance(raceTrackInfo, raceTrackData), NewBasicGeographicMapFactory().GetInstance(raceTrackInfo, raceTrackData, tiledLayerFactoryInterface, geographicMapCellPositionFactoryInterface, geographicMapCellTypeFactory), geographicMapCellPositionFactoryInterface, geographicMapCellPositionBaseFactory, geographicMapCellTypeFactory)
        this.initFields()
        this.CustomMapGenerator = customMapGeneratorBaseFactory.Create(this)
    return this
}


func (this *RaceTrackGeographicMap) GetCustomMapGenerator() ICustomMapGeneratorBase {
        return this.CustomMapGenerator
}


// throws Exception
func (this *RaceTrackGeographicMap) IsValid(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var customMapArray [][]int = this.CustomMapGenerator.GetCustomMapArray(); _ = customMapArray
        if geographicMapCellPosition.GetColumn() >= len(customMapArray[0]) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Column: ")
            stringBuffer.Appendint(geographicMapCellPosition.GetColumn())
            stringBuffer.Append(" not in: ")
            stringBuffer.Appendint(len(customMapArray[0]))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            if geographicMapCellPosition.GetColumn() == len(customMapArray[0]) {
                return true
            } else {
                return false
            }
        } else if geographicMapCellPosition.GetRow() > len(customMapArray) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Row: ")
            stringBuffer.Appendint(geographicMapCellPosition.GetRow())
            stringBuffer.Append(" not in: ")
            stringBuffer.Appendint(len(customMapArray))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            if geographicMapCellPosition.GetRow() == len(customMapArray) {
                return true
            } else {
                return false
            }
        }
        return true
}

