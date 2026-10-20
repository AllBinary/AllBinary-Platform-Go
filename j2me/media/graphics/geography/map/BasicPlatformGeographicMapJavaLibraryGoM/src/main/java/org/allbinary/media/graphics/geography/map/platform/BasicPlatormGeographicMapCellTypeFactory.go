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

type BasicPlatormGeographicMapCellTypeFactory struct {
    GeographicMapCellTypeFactory
    EnumerationUtil *EnumerationUtil
    BLOCK_CELL_TYPE *BasicPlatormGeographicMapCellType
    JUMP_THRU_CELL_TYPE *BasicPlatormGeographicMapCellType
    LADDER_CELL_TYPE *BasicPlatormGeographicMapCellType
    MaxTileId int
}


func (this *BasicPlatormGeographicMapCellTypeFactory) initFields() {
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    
    
    
    
}


func NewBasicPlatormGeographicMapCellTypeFactory(tileTypeToTileIdsMap *ABHashtable[any, any], maxTileId int) *BasicPlatormGeographicMapCellTypeFactory {
    this := &BasicPlatormGeographicMapCellTypeFactory{}
    this.GeographicMapCellTypeFactory = *NewGeographicMapCellTypeFactory()
    this.initFields()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(commonStrings.START, this, commonStrings.INIT)
        this.MaxTileId = maxTileId
        NewGeographicMapCellType(0, 0)
        var BLOCK_CELL_TYPE *BasicPlatormGeographicMapCellType = BasicPlatormGeographicMapCellType_CreateType(1); _ = BLOCK_CELL_TYPE
        var JUMP_THRU_CELL_TYPE *BasicPlatormGeographicMapCellType = BLOCK_CELL_TYPE; _ = JUMP_THRU_CELL_TYPE
        var LADDER_CELL_TYPE *BasicPlatormGeographicMapCellType = BLOCK_CELL_TYPE; _ = LADDER_CELL_TYPE
        var PLATFORM string = "Platform"; _ = PLATFORM
        var JUMP_TRHU string = "JumpThru"; _ = JUMP_TRHU
        var LADDER string = "Ladder"; _ = LADDER
        var enumeration Enumeration = tileTypeToTileIdsMap.Keys(); _ = enumeration
        var idsWithTypeList IBasicArrayList; _ = idsWithTypeList
        var key string; _ = key
        var basicPlatormGeographicMapCellType *BasicPlatormGeographicMapCellType; _ = basicPlatormGeographicMapCellType
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            key = any(this.EnumerationUtil.NextElement(enumeration)).(string)
            this.LogUtil.PutF(key, this, commonStrings.INIT)
            idsWithTypeList = any(tileTypeToTileIdsMap.Get(key)).(IBasicArrayList)
            basicPlatormGeographicMapCellType = NewBasicPlatormGeographicMapCellType(idsWithTypeList)
            if (key == PLATFORM) {
                BLOCK_CELL_TYPE = basicPlatormGeographicMapCellType
            }
            if (key == JUMP_TRHU) {
                JUMP_THRU_CELL_TYPE = basicPlatormGeographicMapCellType
            }
            if (key == LADDER) {
                LADDER_CELL_TYPE = basicPlatormGeographicMapCellType
            }
        }
        this.BLOCK_CELL_TYPE = BLOCK_CELL_TYPE
        this.JUMP_THRU_CELL_TYPE = JUMP_THRU_CELL_TYPE
        this.LADDER_CELL_TYPE = LADDER_CELL_TYPE
        NewGeographicMapCellType(this.MaxTileId - 1, 0)
        NewGeographicMapCellType(this.MaxTileId - 2, 0)
    return this
}


func (this *BasicPlatormGeographicMapCellTypeFactory) GetStartType() int {
        return this.MaxTileId - 1
}


func (this *BasicPlatormGeographicMapCellTypeFactory) GetEndType() int {
        return this.MaxTileId - 2
}


func (this *BasicPlatormGeographicMapCellTypeFactory) GetEmptyType() int {
        return 0
}


func (this *BasicPlatormGeographicMapCellTypeFactory) IsPath(cellType IGeographicMapCellType) bool {
        if cellType.GetType() == 0 {
            return true
        }
        return false
}

