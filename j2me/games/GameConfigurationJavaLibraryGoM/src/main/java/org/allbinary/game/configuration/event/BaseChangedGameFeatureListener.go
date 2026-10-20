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

type BaseChangedGameFeatureListener struct {
    LogUtil *LogUtil
    GameFeatureUtil *GameFeatureUtil
    List IBasicArrayList
    Changed bool
}

// Implements GameFeatureListenerInterface

// The Java class BaseChangedGameFeatureListener has subclasses so its variables are this interface
type IBaseChangedGameFeatureListener interface {
    BaseChangedGameFeatureListener_() *BaseChangedGameFeatureListener
    OnEvent(eventObject IAllBinaryEventObject)
    OnGameFeatureChange(gameFeatureEvent *GameFeatureEvent)
    Add(gameFeature IFeature)
    Remove(gameFeature IFeature)
    SetChanged(initialized bool)
    IsChangedFeature(gameFeature IFeature) bool
    IsChanged() bool
}

func (this *BaseChangedGameFeatureListener) BaseChangedGameFeatureListener_() *BaseChangedGameFeatureListener {
    return this
}


func (this *BaseChangedGameFeatureListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GameFeatureUtil = GameFeatureUtil_GetInstance()
    this.List = NewBasicArrayListD()
    this.Changed = true
}

func NewBaseChangedGameFeatureListener() *BaseChangedGameFeatureListener {
    this := &BaseChangedGameFeatureListener{}
    this.initFields()
    return this
}


func (this *BaseChangedGameFeatureListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(CommonStrings_GetInstance().NOT_IMPLEMENTED, this)
}


func (this *BaseChangedGameFeatureListener) OnGameFeatureChange(gameFeatureEvent *GameFeatureEvent) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameFeatureUtil.GAME_FEATURE_CHANGED).Append(gameFeatureEvent.GetWhatChanged())), this, this.GameFeatureUtil.ON_GAME_FEATURE_CHANGE)
        this.List.Add(gameFeatureEvent.GetGameOption())
        this.SetChanged(true)
}


func (this *BaseChangedGameFeatureListener) Add(gameFeature IFeature) {
        this.List.Add(gameFeature)
}


func (this *BaseChangedGameFeatureListener) Remove(gameFeature IFeature) {
        this.List.Remove(gameFeature)
}


func (this *BaseChangedGameFeatureListener) SetChanged(initialized bool) {
        this.Changed = initialized
        if !this.IsChanged() {
            this.List.Clear()
        }
}


func (this *BaseChangedGameFeatureListener) IsChangedFeature(gameFeature IFeature) bool {
        var isChanged bool = this.List.Contains(gameFeature); _ = isChanged
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("GameFeature: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(gameFeature))
        stringBuffer.Append(" isFeature: ")
        stringBuffer.Appendboolean(Features_GetInstance().IsFeature(gameFeature))
        stringBuffer.Append(" isChanged: ")
        stringBuffer.Appendboolean(isChanged)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isChanged")
        return isChanged
}


func (this *BaseChangedGameFeatureListener) IsChanged() bool {
        return this.Changed
}

