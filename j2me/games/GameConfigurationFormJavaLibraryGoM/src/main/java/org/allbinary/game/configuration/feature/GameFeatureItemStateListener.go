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

type GameFeatureItemStateListener struct {
    LogUtil *LogUtil
    GameOptionsForm *GameOptionsForm
}

// Implements ItemStateListener

var GameFeatureItemStateListener_ToggleList IBasicArrayList = NewBasicArrayListD()

func (this *GameFeatureItemStateListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewGameFeatureItemStateListener(gameOptionsForm *GameOptionsForm) *GameFeatureItemStateListener {
    this := &GameFeatureItemStateListener{}
    this.initFields()
        this.GameOptionsForm = gameOptionsForm
        var gameFeatureFactory *GameFeatureFactory = GameFeatureFactory_GetInstance(); _ = gameFeatureFactory
        GameFeatureItemStateListener_Add(gameFeatureFactory.ARTIFICIAL_INTELLEGENCE_PROCESSOR)
        GameFeatureItemStateListener_Add(gameFeatureFactory.COLLIDABLE_INTERFACE_LAYER_PROCESSOR)
        GameFeatureItemStateListener_Add(gameFeatureFactory.DAMAGE_FLOATERS)
        GameFeatureItemStateListener_Add(gameFeatureFactory.DROPPED_ITEMS)
        GameFeatureItemStateListener_Add(gameFeatureFactory.GAME_INPUT_LAYER_PROCESSOR)
        GameFeatureItemStateListener_Add(gameFeatureFactory.HEALTH_BARS)
        GameFeatureItemStateListener_Add(gameFeatureFactory.SOUND)
        GameFeatureItemStateListener_Add(gameFeatureFactory.TICKABLE_LAYER_PROCESSOR)
    return this
}


func GameFeatureItemStateListener_Add(gameFeature IGameFeature) {
        if !GameFeatureItemStateListener_ToggleList.Contains(gameFeature) {
            GameFeatureItemStateListener_ToggleList.Add(gameFeature)
        }
}


func (this *GameFeatureItemStateListener) ItemStateChanged(item IItem) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "itemStateChanged", e)
                }
            }()
            var itemLabel string = item.GetLabel(); _ = itemLabel
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().ITEM_LABEL).Append(itemLabel)), this, "itemStateChanged")
            if func() bool { _, ok := any(item).(*GameConfigurationGauge); return ok }() {
                GameConfigurationUtil_GetInstance().Change(this.GameOptionsForm, any(item).(*GameConfigurationGauge))
            } else if func() bool { _, ok := any(item).(IChoiceGroup); return ok }() {
                var gameFeatureUtil *GameFeatureUtil = GameFeatureUtil_GetInstance(); _ = gameFeatureUtil
                if gameFeatureUtil.IsExclusive(itemLabel) {
                    gameFeatureUtil.UpdateExclusiveForChoiceGroup(any(item).(IChoiceGroup))
                } else {
                    gameFeatureUtil.UpdateMultiple(any(item).(IChoiceGroup))
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

