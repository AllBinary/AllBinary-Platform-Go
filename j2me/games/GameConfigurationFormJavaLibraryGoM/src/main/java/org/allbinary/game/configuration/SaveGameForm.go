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

type SaveGameForm struct {
    CommandForm
    MeUtil *MEUtil
}

var SaveGameForm_FORM ICommandForm = CommandForm_GetNullCommandForm()

func (this *SaveGameForm) initFields() {
    this.MeUtil = MEUtil_GetInstance()
}


// throws Exception
func SaveGameForm_CreateForm(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) ICommandForm {
        SaveGameForm_FORM = NewSaveGameForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        return SaveGameForm_FORM
}


func SaveGameForm_GetInstance() ICommandForm {
        return SaveGameForm_FORM
}


func NewSaveGameForm(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *SaveGameForm {
    this := &SaveGameForm{}
        this.CommandForm = *NewCommandForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        var timeString string = fmt.Sprint(ABSystemWrapper_GetInstance().CurrentTimeMillis()); _ = timeString
        this.MeUtil.AppendItem(this, NewTextField("Name: ", timeString, 30, TextField_ANY))
        this.InitCommands(commandListener)
    return this
}


func (this *SaveGameForm) InitCommands(cmdListener CommandListener) {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        this.RemoveAllCommands()
        this.AddCommand(gameCommandsFactory.SAVE)
        this.AddCommand(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS)
        this.SetCommandListener(cmdListener)
}

