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

type GeographicMapCellPositionFactory struct {
    GeographicMapCellPositionBaseFactory
    LogUtil *LogUtil
}

var GeographicMapCellPositionFactory_Hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable()

func (this *GeographicMapCellPositionFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewGeographicMapCellPositionFactory() *GeographicMapCellPositionFactory {
    this := &GeographicMapCellPositionFactory{}
    this.GeographicMapCellPositionBaseFactory = *NewGeographicMapCellPositionBaseFactory()
    this.initFields()
    return this
}


// throws Exception
func (this *GeographicMapCellPositionFactory) GetInstance(geographicMapInterface IBasicGeographicMap) IBasicGeographicMapCellPositionFactory {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var geographicMapCellPositionFactoryCanBeNull any = GeographicMapCellPositionFactory_Hashtable.Get(allBinaryTiledLayer.GetDataId()); _ = geographicMapCellPositionFactoryCanBeNull
        if geographicMapCellPositionFactoryCanBeNull == nil {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Creating GeographicMapCellPositionFactory for TileLayer: ").Appendint(int(allBinaryTiledLayer.GetDataId()))), this, commonStrings.GET_INSTANCE)
            geographicMapCellPositionFactoryCanBeNull = NewBasicGeographicMapCellPositionFactory(geographicMapInterface)
            GeographicMapCellPositionFactory_Hashtable.Put(allBinaryTiledLayer.GetDataId(), geographicMapCellPositionFactoryCanBeNull)
            return any(geographicMapCellPositionFactoryCanBeNull).(IBasicGeographicMapCellPositionFactory)
        } else {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Reusing GeographicMapCellPositionFactory for TileLayer: ").Appendint(int(allBinaryTiledLayer.GetDataId()))), this, commonStrings.GET_INSTANCE)
            return any(geographicMapCellPositionFactoryCanBeNull).(IBasicGeographicMapCellPositionFactory)
        }
}


func GeographicMapCellPositionFactory_GetHashtable() *ABHashtable[any, any] {
        return GeographicMapCellPositionFactory_Hashtable
}

