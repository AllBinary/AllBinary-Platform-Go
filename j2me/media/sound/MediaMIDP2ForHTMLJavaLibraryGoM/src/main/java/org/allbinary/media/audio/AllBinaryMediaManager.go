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

//MediaMIDP2ForHTMLJavaLibrary
type AllBinaryMediaManager struct {
}

var AllBinaryMediaManager_THIS string = "AllBinaryMediaManagerHTML5"


func NewAllBinaryMediaManager() *AllBinaryMediaManager {
    this := &AllBinaryMediaManager{}
    return this
}


func AllBinaryMediaManager_IsMuted() bool {
        return false
}


func AllBinaryMediaManager_SetMuted(aMuted bool) {
}


func AllBinaryMediaManager_Update() bool {
        return true
}


// throws Exception
func AllBinaryMediaManager_Init(soundsFactoryInterface SoundsFactoryInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonString *CommonStrings = CommonStrings_GetInstance(); _ = commonString
        logUtil.PutF(commonString.START, AllBinaryMediaManager_THIS, commonString.INIT)
        ProgressCanvasFactory_GetInstance().AddNormalPortion(50, "Media Manager")
        NewSounds(soundsFactoryInterface).Init()
}


// throws Exception
func AllBinaryMediaManager_Shutdown(soundsFactoryInterface SoundsFactoryInterface) {
        NewSounds(soundsFactoryInterface).StopAll()
        NewSounds(soundsFactoryInterface).CloseAll()
        TsUtil_GetInstance().Gc()
}


// throws Exception
func AllBinaryMediaManager_CreatePlayer(resource string) *Player {
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().SOUND) {
            if TsUtil_GetInstance().CompareTo(resource, Manager_TONE_DEVICE_LOCATOR) == 0 {
                return Manager_CreatePlayer(resource)
            } else {
                return Manager_CreatePlayer(resource[0:len(resource) - 4])
            }
        } else {
            return NoPlayer_NO_PLAYER
        }
}


// throws MediaException
// synchronized
func AllBinaryMediaManager_PlayTone(frequency int, time int, volume int) {
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().SOUND) {
            Manager_PlayTone(frequency, time, volume)
        }
}

