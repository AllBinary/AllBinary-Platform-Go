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

type DynamicRaceTracksPathPrinter struct {
    RaceTracksPathPrinter
    LogUtil *LogUtil
}


func (this *DynamicRaceTracksPathPrinter) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewDynamicRaceTracksPathPrinter() *DynamicRaceTracksPathPrinter {
    this := &DynamicRaceTracksPathPrinter{}
    this.RaceTracksPathPrinter = *NewRaceTracksPathPrinter()
    this.initFields()
    return this
}


func (this *DynamicRaceTracksPathPrinter) PrintAllPaths(raceTrackGeographicMapInterfaceFactoryInterface RaceTrackGeographicMapInterfaceFactoryInterface, pathFindingInfo *PathFindingInfo, size int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.PROCESS, e)
                }
            }()
            for index := int(0); index < size; index++ {
                var geographicMapInterface IBasicGeographicMap = raceTrackGeographicMapInterfaceFactoryInterface.GetTrackInstance(SmallIntegerSingletonFactory_GetInstance().GetAt(index)); _ = geographicMapInterface
                var list IBasicArrayList = PathGenerator_GetInstance().CreateList(geographicMapInterface, RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance(), pathFindingInfo, 10); _ = list
                fmt.Println("Track " + fmt.Sprint(index))
                this.PrintPaths(list)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

