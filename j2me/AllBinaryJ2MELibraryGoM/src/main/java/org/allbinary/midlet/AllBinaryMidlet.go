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

type AllBinaryMidlet struct {
    MIDlet
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SET_DISPLAY string
    SETTING_ string
    _DISPLAY_ string
    SETTING_NO_TITLE string
    Hashtable *ABHashtable[any, any]
    MidletDestroyed bool
}

// Implements CommandListener

// The Java class AllBinaryMidlet has subclasses so its variables are this interface
type IAllBinaryMidlet interface {
    IMIDlet
    AllBinaryMidlet_() *AllBinaryMidlet
    SetDisplay(newDisplay IDisplayable)
    GetDisplay() *Display
    GetCurrentDisplayable() IDisplayable
    SetDestroyed(destroyed bool)
    IsDestroyed() bool
    StartApp()
    PauseApp()
    DestroyAppInRunnable(unconditional bool, isProgress bool)
    DestroyApp(unconditional bool)
    SetStartStateHashtable(hashtable *ABHashtable[any, any])
    GetStartStateHashtable() *ABHashtable[any, any]
    GetCurrentStateHashtable() *ABHashtable[any, any]
    CommandAction(command ICommand, displayable IDisplayable)
}

func (this *AllBinaryMidlet) AllBinaryMidlet_() *AllBinaryMidlet {
    return this
}

var AllBinaryMidlet_NULL_ALLBINARY_MIDLET any = NullUtil_GetInstance().NULL_OBJECT

func (this *AllBinaryMidlet) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SET_DISPLAY = "setDisplay"
    this.SETTING_ = "Setting: "
    this._DISPLAY_ = " Display: "
    this.SETTING_NO_TITLE = "Setting: No Title, Display: "
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    
}


func AllBinaryMidlet_GetNullInstance() IAllBinaryMidlet {
        if AllBinaryMidlet_NULL_ALLBINARY_MIDLET == NullUtil_GetInstance().NULL_OBJECT {
            AllBinaryMidlet_NULL_ALLBINARY_MIDLET = NewAllBinaryMidlet()
        }
        return any(AllBinaryMidlet_NULL_ALLBINARY_MIDLET).(IAllBinaryMidlet)
}


func NewAllBinaryMidlet() *AllBinaryMidlet {
    this := &AllBinaryMidlet{}
    this.MIDlet = *NewMIDlet()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, "AllBinaryMidlet::AllBinaryMidlet")
    return this
}


func (this *AllBinaryMidlet) SetDisplay(newDisplay IDisplayable) {
        var title string = StringUtil_GetInstance().EMPTY_STRING; _ = title
        if newDisplay != NullCanvas_NULL_CANVAS {
            title = newDisplay.GetTitle()
            if title != "" {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.SETTING_).Append(title).Append(this._DISPLAY_).Append(StringUtil_GetInstance().ToString(newDisplay))), this, this.SET_DISPLAY)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.SETTING_NO_TITLE).Append(StringUtil_GetInstance().ToString(newDisplay))), this, this.SET_DISPLAY)
            }
        }
        var display *Display = this.GetDisplay(); _ = display
        display.SetCurrent(newDisplay)
}


func (this *AllBinaryMidlet) GetDisplay() *Display {
        return Display_GetDisplay(this)
}


func (this *AllBinaryMidlet) GetCurrentDisplayable() IDisplayable {
        return Display_GetDisplay(this).GetCurrent()
}


func (this *AllBinaryMidlet) SetDestroyed(destroyed bool) {
        this.MidletDestroyed = destroyed
}


func (this *AllBinaryMidlet) IsDestroyed() bool {
        return this.MidletDestroyed
}


// throws MIDletStateChangeException
func (this *AllBinaryMidlet) StartApp() {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *AllBinaryMidlet) PauseApp() {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *AllBinaryMidlet) DestroyAppInRunnable(unconditional bool, isProgress bool) {
}


func (this *AllBinaryMidlet) DestroyApp(unconditional bool) {
        var METHOD_NAME string = "AllBinaryMidlet::destroyApp"; _ = METHOD_NAME
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, METHOD_NAME, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, METHOD_NAME)
            PreLogUtil_Put(Memory_GetInfo(), this, METHOD_NAME)
            this.SetDestroyed(true)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryMidlet) SetStartStateHashtable(hashtable *ABHashtable[any, any]) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(StringUtil_GetInstance().ToString(hashtable))), this, "setStartStateHashtable")
        this.Hashtable = hashtable
}


// throws Exception
func (this *AllBinaryMidlet) GetStartStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(StringUtil_GetInstance().ToString(this.Hashtable))), this, "getStartStateHashtable")
        return this.Hashtable
}


// throws Exception
func (this *AllBinaryMidlet) GetCurrentStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getStateHashtable")
        return StdUtil_GetInstance().NULL_TABLE
}


func (this *AllBinaryMidlet) CommandAction(command ICommand, displayable IDisplayable) {
}

