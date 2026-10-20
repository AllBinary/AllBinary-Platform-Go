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

type GameFeatureItemCommandListener struct {
    LogUtil *LogUtil
    GameOptionsForm *GameOptionsForm
}

// Implements ItemCommandListener


func (this *GameFeatureItemCommandListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewGameFeatureItemCommandListener(gameOptionsForm *GameOptionsForm) *GameFeatureItemCommandListener {
    this := &GameFeatureItemCommandListener{}
    this.initFields()
        this.GameOptionsForm = gameOptionsForm
    return this
}


func (this *GameFeatureItemCommandListener) CommandAction(command ICommand, item IItem) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, MidletStrings_GetInstance().COMMAND_ACTION, e)
                }
            }()
            var itemLabel string = item.GetLabel(); _ = itemLabel
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().ITEM_LABEL).Append(itemLabel)), this, MidletStrings_GetInstance().COMMAND_ACTION)
            if func() bool { _, ok := any(item).(*GameConfigurationGauge); return ok }() {
                GameConfigurationUtil_GetInstance().Change(this.GameOptionsForm, any(item).(*GameConfigurationGauge))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

