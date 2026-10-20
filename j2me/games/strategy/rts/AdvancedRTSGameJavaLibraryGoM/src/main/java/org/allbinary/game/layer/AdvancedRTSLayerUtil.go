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

type AdvancedRTSLayerUtil struct {
}

var AdvancedRTSLayerUtil_Instance *AdvancedRTSLayerUtil = NewAdvancedRTSLayerUtil()

func NewAdvancedRTSLayerUtil() *AdvancedRTSLayerUtil {
    this := &AdvancedRTSLayerUtil{}
    return this
}


func AdvancedRTSLayerUtil_GetInstance() *AdvancedRTSLayerUtil {
        return AdvancedRTSLayerUtil_Instance
}


func (this *AdvancedRTSLayerUtil) CreateFakeRTSLayerHashtable(baseRaceTrackGeographicMap IBaseRaceTrackGeographicMap) *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = RTSLayerUtil_GetInstance().CreateFakeRTSLayerHashtable(baseRaceTrackGeographicMap); _ = hashtable
        hashtable.Put(WorkWaypoint_ID, SmallIntegerSingletonFactory_GetInstance().GetAt(50))
        return hashtable
}

