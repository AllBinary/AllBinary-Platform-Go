/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
 *  
 * By agreeing to this license you and any business entity you represent are 
 * legally bound to the AllBinary Open License Version 1 legal agreement. 
 *  
 * You may obtain the AllBinary Open License Version 1 legal agreement from 
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository. 
 *  
 * Created By: Travis Berthelot  
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

type FakeLayerManager struct {
    AllBinaryGameLayerManager
    GeographicMapInterfaceArray []IBasicGeographicMap
    GeographicMapCellTypeArray []IGeographicMapCellType
}

// Implements GeographicMapCompositeInterface


func (this *FakeLayerManager) initFields() {
    this.GeographicMapInterfaceArray = BasicGeographicMap_NULL_BASIC_GEOGRAPHIC_MAP_ARRAY
    this.GeographicMapCellTypeArray = GeographicMapCellType_NULL_GEOGRAPHIC_MAP_CELL_TYPE_ARRAY
}


func NewFakeLayerManager(gameInfo *GameInfo) *FakeLayerManager {
    this := &FakeLayerManager{}
        this.AllBinaryGameLayerManager = *NewAllBinaryGameLayerManager(BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE, gameInfo)
        this.initFields()
    return this
}


func (this *FakeLayerManager) GetGeographicMapInterface() []IBasicGeographicMap {
        return this.GeographicMapInterfaceArray
}


func (this *FakeLayerManager) SetGeographicMapInterface(geographicMapInterfaceArray []IBasicGeographicMap) {
        this.GeographicMapInterfaceArray = geographicMapInterfaceArray
        this.GeographicMapCellTypeArray = make([]IGeographicMapCellType, len(this.GeographicMapInterfaceArray))
}


func (this *FakeLayerManager) GeographicMapCellTypeArray() []IGeographicMapCellType {
        return this.GeographicMapCellTypeArray
}

