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

type LicenseRegistrationTextBox struct {
    CustomTextBox
    LogUtil *LogUtil
    PleaseWaitPaintable *Paintable
    Paintable *Paintable
}

var LicenseRegistrationTextBox_DISPLAY_COMMAND *Command = NewCommand("Register", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)
var LicenseRegistrationTextBox_SUBMIT_COMMAND *Command = NewCommand("Submit", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)

func (this *LicenseRegistrationTextBox) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PleaseWaitPaintable = NewSimpleTextPaintable(this.CommonStrings.PLEASE_WAIT, BasicColorFactory_GetInstance().WHITE)
    this.Paintable = NullPaintable_GetInstance()
}


func NewLicenseRegistrationTextBox(cmdListener *CommandListener, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *LicenseRegistrationTextBox {
    this := &LicenseRegistrationTextBox{}
        this.CustomTextBox = *NewCustomTextBox(cmdListener, "Enter Registration Code:", StringUtil_GetInstance().EMPTY_STRING, 60, TextField_ANY, Font_GetDefaultFont(), backgrounBasicColor, foregroundBasicColor)
        this.initFields()
    return this
}


func (this *LicenseRegistrationTextBox) InitCommands(cmdListener *CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameCommandsFactory_GetInstance().CLOSE_AND_SHOW_GAME_CANVAS)
        this.AddCommand(GameCommandsFactory_GetInstance().TOGGLE_KEYBOARD)
        this.AddCommand(LicenseRegistrationTextBox_SUBMIT_COMMAND)
        this.SetCommandListener(cmdListener)
}


func (this *LicenseRegistrationTextBox) Open() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "open", e)
                }
            }()
            var virtualKeyboardEventHandler *VirtualKeyboardEventHandler = VirtualKeyboardEventHandler_GetInstance(); _ = virtualKeyboardEventHandler
            virtualKeyboardEventHandler.FireEvent(virtualKeyboardEventHandler.SHOW_EVENT)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.CustomTextBox.Open()
        this.Paintable = NullPaintable_GetInstance()
        this.Repaint()
}


// throws Exception
func (this *LicenseRegistrationTextBox) Close() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "open", e)
                }
            }()
            var virtualKeyboardEventHandler *VirtualKeyboardEventHandler = VirtualKeyboardEventHandler_GetInstance(); _ = virtualKeyboardEventHandler
            virtualKeyboardEventHandler.FireEvent(virtualKeyboardEventHandler.HIDE_EVENT)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.Paintable = this.PleaseWaitPaintable
        this.Repaint()
        this.CustomTextBox.Close()
}


func (this *LicenseRegistrationTextBox) Paint(graphics *Graphics) {
        this.CustomTextBox.Paint(graphics)
        this.Paintable.Paint(graphics)
}


func (this *LicenseRegistrationTextBox) Submit() {
        var commandListener *CommandListener = this.GetCustomCommandListener(); _ = commandListener
        commandListener.CommandAction(LicenseRegistrationTextBox_SUBMIT_COMMAND, this)
}

