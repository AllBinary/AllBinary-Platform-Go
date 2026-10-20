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

type PlayerTopViewCharacterBehavior struct {
    TopViewCharacterBehavior
    LogUtil *LogUtil
}


func (this *PlayerTopViewCharacterBehavior) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewPlayerTopViewCharacterBehavior() *PlayerTopViewCharacterBehavior {
    this := &PlayerTopViewCharacterBehavior{}
    this.TopViewCharacterBehavior = *NewTopViewCharacterBehavior()
    this.initFields()
    return this
}


// throws Exception
func (this *PlayerTopViewCharacterBehavior) TerrainEvent(layer IAllBinaryLayer, direction *Direction, x int, y int, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellPosition IGeographicMapCellPosition) {
        TrackingEventHandler_GetInstance().FireEvent(TrackingEventCircularStaticPool_GetInstance().GetNextInstance(layer))
}


func (this *PlayerTopViewCharacterBehavior) TerrainMove(layer IAllBinaryLayer, geographicMapInterfaceArray []IBasicGeographicMap, x int, y int) {
        var basicGeographicMapUtil *BasicGeographicMapUtil = BasicGeographicMapUtil_GetInstance(); _ = basicGeographicMapUtil
        basicGeographicMapUtil.SetPosition(geographicMapInterfaceArray, x, y)
}


func (this *PlayerTopViewCharacterBehavior) MoveIfOnScreen(layer IAllBinaryLayer, ax int, ay int) {
}


func (this *PlayerTopViewCharacterBehavior) IsTiledLayerMoveable(terrainTiledLayer IAllBinaryTiledLayer, x int, y int) bool {
        return (terrainTiledLayer.GetXP() + DisplayInfoSingleton_GetInstance().GetLastWidth() < terrainTiledLayer.GetWidth() || x < 0) && (terrainTiledLayer.GetXP() > 0 || x > 0)
}

