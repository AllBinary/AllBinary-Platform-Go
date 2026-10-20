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

type BasicTopViewGeographicMapCellTypeFactory struct {
    GeographicMapCellTypeFactory
    BLOCK_CELL_TYPE *BasicTopViewGeographicMapCellType
    OFF_MAP_CELL_TYPE *BasicTopViewGeographicMapCellType
    FLOOR_CELL_TYPE *BasicTopViewGeographicMapCellType
    DOOR_CELL_TYPE *BasicTopViewGeographicMapCellType
    STAIRS_UP_CELL_TYPE *BasicTopViewGeographicMapCellType
    STAIRS_DOWN_CELL_TYPE *BasicTopViewGeographicMapCellType
    OTHER_CELL_TYPE *BasicTopViewGeographicMapCellType
    MaxTileId int
}


func (this *BasicTopViewGeographicMapCellTypeFactory) initFields() {
    
    
    
    
    
    
    
    
}


func NewBasicTopViewGeographicMapCellTypeFactory(tileTypeToTileIdsMap *ABHashtable[any, any], maxTileId int) *BasicTopViewGeographicMapCellTypeFactory {
    this := &BasicTopViewGeographicMapCellTypeFactory{}
    this.GeographicMapCellTypeFactory = *NewGeographicMapCellTypeFactory()
    this.initFields()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(commonStrings.START, this, commonStrings.CONSTRUCTOR)
        this.MaxTileId = maxTileId
        var basicTopViewGeographicMapStrings *BasicTopViewGeographicMapStrings = BasicTopViewGeographicMapStrings_GetInstance(); _ = basicTopViewGeographicMapStrings
        var geographicMapCellTypeFactory IGeographicMapCellTypeFactory = GeographicMapCellTypeFactory_GetInstance(); _ = geographicMapCellTypeFactory
        var geographicMapCellTypeArray []IGeographicMapCellType = geographicMapCellTypeFactory.GetGeographicMapCellTypeArray(); _ = geographicMapCellTypeArray
        var type_ int = 0; _ = type_
        if geographicMapCellTypeArray[type_] == nil {
            NewRaceTrackGeographicMapCellType(CommonStrings_GetInstance().UNKNOWN, type_, 999, 0)
        } else {
        }
        var BLOCK_CELL_TYPE *BasicTopViewGeographicMapCellType = BasicTopViewGeographicMapCellType_CreateType(basicTopViewGeographicMapStrings.DEFAULT, 1, 1); _ = BLOCK_CELL_TYPE
        var OFF_MAP_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = OFF_MAP_CELL_TYPE
        var FLOOR_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = FLOOR_CELL_TYPE
        var DOOR_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = DOOR_CELL_TYPE
        var STAIRS_UP_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = STAIRS_UP_CELL_TYPE
        var STAIRS_DOWN_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = STAIRS_DOWN_CELL_TYPE
        var OTHER_CELL_TYPE *BasicTopViewGeographicMapCellType = BLOCK_CELL_TYPE; _ = OTHER_CELL_TYPE
        var keyArray []any = HashtableUtil_GetInstance().GetKeysAsArray(tileTypeToTileIdsMap); _ = keyArray
        var size int = len(keyArray); _ = size
        var idsWithTypeList IBasicArrayList; _ = idsWithTypeList
        var key string; _ = key
        var basicPlatormGeographicMapCellType *BasicTopViewGeographicMapCellType; _ = basicPlatormGeographicMapCellType
        for index := int(0); index < size; index++ {
            key = any(keyArray[index]).(string)
            idsWithTypeList = any(tileTypeToTileIdsMap.Get(key)).(IBasicArrayList)
            if (key == basicTopViewGeographicMapStrings.WALL) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.WALL, idsWithTypeList, 1000)
                BLOCK_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.OFF_MAP) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.OFF_MAP, idsWithTypeList, 1001)
                OFF_MAP_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.FLOOR) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.FLOOR, idsWithTypeList, 1)
                FLOOR_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.DOOR) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.DOOR, idsWithTypeList, 1)
                DOOR_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.STAIRS_UP) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.STAIRS_UP, idsWithTypeList, 1)
                STAIRS_UP_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.STAIRS_DOWN) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.STAIRS_DOWN, idsWithTypeList, 1)
                STAIRS_DOWN_CELL_TYPE = basicPlatormGeographicMapCellType
            } else if (key == basicTopViewGeographicMapStrings.OTHER) {
                basicPlatormGeographicMapCellType = NewBasicTopViewGeographicMapCellType(basicTopViewGeographicMapStrings.OTHER, idsWithTypeList, 1)
                OTHER_CELL_TYPE = basicPlatormGeographicMapCellType
            }
        }
        this.BLOCK_CELL_TYPE = BLOCK_CELL_TYPE
        this.OFF_MAP_CELL_TYPE = OFF_MAP_CELL_TYPE
        this.FLOOR_CELL_TYPE = FLOOR_CELL_TYPE
        this.DOOR_CELL_TYPE = DOOR_CELL_TYPE
        this.STAIRS_UP_CELL_TYPE = STAIRS_UP_CELL_TYPE
        this.STAIRS_DOWN_CELL_TYPE = STAIRS_DOWN_CELL_TYPE
        this.OTHER_CELL_TYPE = OTHER_CELL_TYPE
        type_ = this.MaxTileId - 1
        if geographicMapCellTypeArray[type_] == nil {
            NewRaceTrackGeographicMapCellType(commonStrings.START, type_, 1, 0)
        } else {
        }
        type_ = this.MaxTileId - 2
        if geographicMapCellTypeArray[type_] == nil {
            NewRaceTrackGeographicMapCellType(commonStrings.START, type_, 1, 0)
        } else {
        }
    return this
}


func (this *BasicTopViewGeographicMapCellTypeFactory) GetStartType() int {
        return this.MaxTileId - 1
}


func (this *BasicTopViewGeographicMapCellTypeFactory) GetEndType() int {
        return this.MaxTileId - 2
}


func (this *BasicTopViewGeographicMapCellTypeFactory) GetEmptyType() int {
        return this.FLOOR_CELL_TYPE.GetTypes()[0]
}


func (this *BasicTopViewGeographicMapCellTypeFactory) IsPath(cellType IGeographicMapCellType) bool {
        if this.FLOOR_CELL_TYPE.IsType(cellType) {
            return true
        }
        return false
}


func (this *BasicTopViewGeographicMapCellTypeFactory) ToString() string {
        return fmt.Sprint(NewStringMaker().Append("key: WALL/BLOCK_CELL_TYPE: ").Append(fmt.Sprint(this.BLOCK_CELL_TYPE)).Append("key: FLOOR_CELL_TYPE: ").Append(fmt.Sprint(this.FLOOR_CELL_TYPE)).Append("key: OTHER_CELL_TYPE: ").Append(fmt.Sprint(this.OTHER_CELL_TYPE)).Append("key: OFF_MAP_CELL_TYPE: ").Append(fmt.Sprint(this.OFF_MAP_CELL_TYPE)).Append("key: DOOR_CELL_TYPE: ").Append(fmt.Sprint(this.DOOR_CELL_TYPE)).Append("key: STAIRS_DOWN_CELL_TYPE: ").Append(fmt.Sprint(this.STAIRS_DOWN_CELL_TYPE)).Append("key: STAIRS_UP_CELL_TYPE: ").Append(fmt.Sprint(this.STAIRS_UP_CELL_TYPE)))
}

func (this *BasicTopViewGeographicMapCellTypeFactory) String() string {
    return this.ToString()
}

