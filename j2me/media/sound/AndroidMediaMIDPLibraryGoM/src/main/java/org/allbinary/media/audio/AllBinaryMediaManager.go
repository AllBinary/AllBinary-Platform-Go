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

//AndroidMediaMIDPLibrary
type AllBinaryMediaManager struct {
}

var AllBinaryMediaManager_THIS string = "AllBinaryMediaManagerAndroid"
var AllBinaryMediaManager_MostUsedTotal int = 0


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
        AllBinaryMediaManager_Shutdown(soundsFactoryInterface)
        ProgressCanvasFactory_GetInstance().AddNormalPortion(50, "Media Manager")
        TsUtil_GetInstance().Gc()
        NewSounds(soundsFactoryInterface).Init()
        logUtil.PutF(commonString.END, AllBinaryMediaManager_THIS, commonString.INIT)
}


// throws Exception
func AllBinaryMediaManager_Shutdown(soundsFactoryInterface SoundsFactoryInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonString *CommonStrings = CommonStrings_GetInstance(); _ = commonString
        logUtil.PutF(commonString.START, AllBinaryMediaManager_THIS, "shutdown")
        if soundsFactoryInterface.IsInitialized() {
            NewSounds(soundsFactoryInterface).StopAll()
            var soundInterfaceArray []ISound = soundsFactoryInterface.GetSoundInterfaceArray(); _ = soundInterfaceArray
            var player *Player; _ = player
            var player2 *Player; _ = player2
            var androidMediaPlayerWrapper *AndroidMediaPlayerWrapper; _ = androidMediaPlayerWrapper
            for i := int(0); i < len(soundInterfaceArray); i++ {
                if soundInterfaceArray[i] != nil {
                    player = soundInterfaceArray[i].GetPlayerP()
                    if player != nil {
                        player2 = (any(player).(*PlayerComposite)).GetPlayerP()
                        if func() bool { _, ok := any(player2).(*AndroidMediaPlayerWrapper); return ok }() {
                            androidMediaPlayerWrapper = any(player2).(*AndroidMediaPlayerWrapper)
                            MediaPlayerUtil_GetInstance().WaitForMediaPlayer(androidMediaPlayerWrapper.GetMediaPlayer())
                        } else {
                            panic(errors.New(fmt.Sprint("Unknown Property Player: " + reflect.TypeOf(player).String())))
                        }
                    }
                }
            }
            NewSounds(soundsFactoryInterface).CloseAll()
            TsUtil_GetInstance().Gc()
            soundsFactoryInterface.SetInitialized(false)
            AllBinaryMediaManager_MostUsedTotal = 0
        }
        logUtil.PutF(commonString.START, AllBinaryMediaManager_THIS, "shutdown")
}


// throws Exception
func AllBinaryMediaManager_CreatePlayer(resource string) *Player {
        AllBinaryMediaManager_MostUsedTotal++
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().SOUND) {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Player) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.Put("Could not create AndroidMediaPlayerWrapper using NoPlayer at " + fmt.Sprint(CommonLabels_GetInstance().TOTAL_LABEL) + fmt.Sprint(AllBinaryMediaManager_MostUsedTotal), AllBinaryMediaManager_THIS, "createPlayer", e)
                tryControl1, tryResult1 = 1, NoPlayer_NO_PLAYER; return
                    }
                }()
                tryControl1, tryResult1 = 1, NewAndroidMediaPlayerWrapper(resource); return
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        } else {
            return NoPlayer_NO_PLAYER
        }
}


// throws IOException, MediaException
func AllBinaryMediaManager_CreatePlayerFromInputStream(stream IInputStream, type_ string) *Player {
        panic(errors.New(fmt.Sprint("No Input Stream Player")))
}


// throws MediaException
// synchronized
func AllBinaryMediaManager_PlayTone(frequency int, time int, volume int) {
        panic(errors.New(fmt.Sprint("No Tone Player")))
}


func AllBinaryMediaManager_IsInitialized() bool {
        return true
}

