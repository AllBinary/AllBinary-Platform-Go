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

type Sounds struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SoundStrings *SoundStrings
    SoundsFactoryInterface SoundsFactoryInterface
}


func (this *Sounds) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SoundStrings = SoundStrings_GetInstance()
    
}


func NewSounds(soundsFactoryInterface SoundsFactoryInterface) *Sounds {
    this := &Sounds{}
    this.initFields()
        this.SoundsFactoryInterface = soundsFactoryInterface
    return this
}


// throws Exception
func (this *Sounds) Init() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.INIT)
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        this.SoundsFactoryInterface.Init()
        var soundInterfaceArray []ISound = this.SoundsFactoryInterface.GetSoundInterfaceArray(); _ = soundInterfaceArray
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var indexInteger int; _ = indexInteger
        var indexString string; _ = indexString
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for i := int(0); i < len(soundInterfaceArray); i++ {
            var soundInterface ISound = soundInterfaceArray[i]; _ = soundInterface
            if soundInterface != nil {
                indexInteger = smallIntegerSingletonFactory.GetAt(i)
                indexString = fmt.Sprint(indexInteger)
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(commonLabels.INDEX_LABEL)
                stringBuffer.Append(indexString)
                stringBuffer.Append(this.SoundStrings.SOUND)
                stringBuffer.Append(reflect.TypeOf(soundInterface).String())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.INIT)
                soundInterface.Init()
                progressCanvas.AddNormalPortion(100, fmt.Sprint(NewStringMaker().Append(this.SoundStrings.INIT_SOUND).Append(indexString)))
            }
        }
        this.SoundsFactoryInterface.SetInitialized(true)
        this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.INIT)
}


// throws Exception
func (this *Sounds) StopAll() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.SoundStrings.STOP_ALL)
        var soundInterfaceArray []ISound = this.SoundsFactoryInterface.GetSoundInterfaceArray(); _ = soundInterfaceArray
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        var player *Player; _ = player
        for index := int(0); index < len(soundInterfaceArray); index++ {
            if soundInterfaceArray[index] != nil {
                player = soundInterfaceArray[index].GetPlayerP()
                if player != nil {
                    player.Stop()
                    progressCanvas.AddPortion(100, this.SoundStrings.STOPPING_SOUND, index)
                }
            }
        }
}


// throws Exception
func (this *Sounds) CloseAll() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.SoundStrings.CLOSE_ALL)
        var soundInterfaceArray []ISound = this.SoundsFactoryInterface.GetSoundInterfaceArray(); _ = soundInterfaceArray
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        var player *Player; _ = player
        for index := int(0); index < len(soundInterfaceArray); index++ {
            if soundInterfaceArray[index] != nil {
                player = soundInterfaceArray[index].GetPlayerP()
                if player != nil {
                    player.Close()
                    progressCanvas.AddPortion(100, this.SoundStrings.CLOSING_SOUND, index)
                }
            }
        }
}

