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

type HighScoreTextBox struct {
    CustomTextBox
    VirtualKeyboardEventHandler *VirtualKeyboardEventHandler
    HighScoreUtil *HighScoreUtil
    PleaseWait IPaintable
    Paintable IPaintable
    Submitted bool
}


func (this *HighScoreTextBox) initFields() {
    this.VirtualKeyboardEventHandler = VirtualKeyboardEventHandler_GetInstance()
    
    this.PleaseWait = NewSimpleTextPaintable(this.CommonStrings.PLEASE_WAIT, BasicColorFactory_GetInstance().WHITE)
    this.Paintable = NullPaintable_GetInstance()
    this.Submitted = false
}


func NewHighScoreTextBox(highScoresFactoryInterface HighScoresFactoryInterface, highScoresHelper HighScoresHelperBaseInterface, abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, cmdListener CommandListener, name string, highScore *HighScore, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *HighScoreTextBox {
    this := &HighScoreTextBox{}
        this.CustomTextBox = *NewCustomTextBox(cmdListener, "New High Score Enter Name:", name, 12, TextField_ANY, Font_GetDefaultFont(), backgrounBasicColor, foregroundBasicColor)
        this.initFields()
        this.HighScoreUtil = NewHighScoreUtil(highScoresFactoryInterface, highScoresHelper, abeClientInformation, gameInfo, cmdListener, name, highScore)
        var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
        if operatingSystemInterface.IsOverScan() {
            var username string = InApplicationPurchaseFactory_GetInstance().GetUserName(); _ = username
            if username != "" {
                this.GetTextFieldItem().SetString(username)
            } else {
                this.GetTextFieldItem().SetString("NoGamerInfo")
            }
        }
    return this
}


func (this *HighScoreTextBox) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        PostGameCommandFactory_GetInstance().AddCommands(this)
        this.AddCommand(HighScoreUtil_SUBMIT_TEXTBOX_COMMAND)
        this.SetCommandListener(cmdListener)
}


func (this *HighScoreTextBox) Open() {
        this.VirtualKeyboardEventHandler.Open()
        this.Paintable = NullPaintable_GetInstance()
        this.CustomTextBox.Open()
}


// throws Exception
func (this *HighScoreTextBox) Close() {
        this.VirtualKeyboardEventHandler.Close()
        this.Paintable = this.PleaseWait
        this.RepaintBehavior.OnChangeRepaint(this)
        this.CustomTextBox.Close()
        this.RemoveCommand(HighScoreUtil_SUBMIT_TEXTBOX_COMMAND)
        this.Update()
        if this.Submitted {
            this.HighScoreUtil.SaveHighScore()
        } else {
        }
        this.Paintable = NullPaintable_GetInstance()
        this.RepaintBehavior.OnChangeRepaint(this)
}


// throws Exception
func (this *HighScoreTextBox) Update() {
        var name string = this.GetTextFieldItem().GetString(); _ = name
        this.HighScoreUtil.Update(name)
        this.CustomTextBox.Update()
}


func (this *HighScoreTextBox) Paint(graphics IGraphics) {
        this.CustomTextBox.Paint(graphics)
        this.Paintable.Paint(graphics)
}


func (this *HighScoreTextBox) SaveHighScore() {
        this.HighScoreUtil.SaveHighScore()
}


func (this *HighScoreTextBox) Submit() {
        this.HighScoreUtil.Submit(this)
}

