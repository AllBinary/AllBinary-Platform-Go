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

type MyCommandsFactory struct {
    NO_COMMAND ICommand
    GAUGE_CHANGE ICommand
    SET_DISPLAYABLE ICommand
    PAUSE_COMMAND ICommand
    RESUME_COMMAND ICommand
}

var MyCommandsFactory_Instance *MyCommandsFactory = NewMyCommandsFactory()

func (this *MyCommandsFactory) initFields() {
    this.NO_COMMAND = NewCommand(StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING, -1, -1)
    this.GAUGE_CHANGE = NewCommand("GAUGE_CHANGE", StringUtil_GetInstance().EMPTY_STRING, Command_ITEM, 1)
    this.SET_DISPLAYABLE = NewCommand("Display", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)
    this.PAUSE_COMMAND = NewCommand("Pause", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 0)
    this.RESUME_COMMAND = NewCommand("Resume", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 0)
}


func MyCommandsFactory_GetInstance() *MyCommandsFactory {
        return MyCommandsFactory_Instance
}


func NewMyCommandsFactory() *MyCommandsFactory {
    this := &MyCommandsFactory{}
    this.initFields()
    return this
}

