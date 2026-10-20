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

type InGameOptionsForm struct {
    CommandForm
}

var InGameOptionsForm_DISPLAY ICommand = NewCommand("Options In Game", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
var InGameOptionsForm_SAVE ICommand = NewCommand("Save", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
var InGameOptionsForm_DEFAULT ICommand = NewCommand("Default", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)


func NewInGameOptionsForm(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *InGameOptionsForm {
    this := &InGameOptionsForm{}
        this.CommandForm = *NewCommandForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        var gameFeatureFormUtil *GameFeatureFormUtil = GameFeatureFormUtil_GetInstance(); _ = gameFeatureFormUtil
        gameFeatureFormUtil.AddChoiceGroup(this, InGameFeatureChoiceGroups_GetExclusiveInstance().Get(), Choice_EXCLUSIVE)
        gameFeatureFormUtil.AddChoiceGroup(this, InGameFeatureChoiceGroups_GetMultipleInstance().Get(), Choice_MULTIPLE)
        this.InitCommands(commandListener)
    return this
}


func (this *InGameOptionsForm) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameCommandsFactory_GetInstance().CLOSE_AND_SHOW_GAME_CANVAS)
        this.AddCommand(InGameOptionsForm_DEFAULT)
        this.SetCommandListener(cmdListener)
}

