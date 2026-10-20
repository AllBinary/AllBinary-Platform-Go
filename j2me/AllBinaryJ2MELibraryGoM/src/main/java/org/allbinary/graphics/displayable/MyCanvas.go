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

type MyCanvas struct {
    Canvas
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CanvasStrings *CanvasStrings
    StringUtil *StringUtil
    DisplayInfo *DisplayInfoSingleton
    Name string
    ChildNameList IBasicArrayList
    CommandStack *ABStack[any]
    CommandListener CommandListener
    Paused bool
    TouchME *TouchJ2ME
}

// Implements DisplayableInterface, MyCommandInterface

// The Java class MyCanvas has subclasses so its variables are this interface
type IMyCanvas interface {
    ICanvas
    MyCanvas_() *MyCanvas
    SetFullScreenMode(mode bool)
    SizeChanged(w int, h int)
    GetCommandStack() *ABStack[any]
    IsCommand(command ICommand) bool
    AddCommand(command ICommand)
    RemoveCommand(command ICommand)
    RemoveAllCommands()
    SetCommandListener(l CommandListener)
    GetCustomCommandListener() CommandListener
    IsPaused() bool
    RemovePauseCommand()
    Pause()
    UnPause()
    Process()
    Paint(graphics IGraphics)
    Draw(graphics IGraphics)
    HasChild(displayable IMyCanvas) bool
    Destroy()
    SetPaused(isPaused bool)
    KeyPressedByDevice(keyCode int, deviceId int)
    KeyRepeatedByDevice(keyCode int, deviceId int)
    KeyReleasedByDevice(keyCode int, deviceId int)
    PointerDragged(x int, y int)
    PointerPressed(x int, y int)
    PointerReleased(x int, y int)
    NextSong(nextSongSound ISound, leftVolume int, rightVolume int)
}

func (this *MyCanvas) MyCanvas_() *MyCanvas {
    return this
}


func (this *MyCanvas) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    
    
    
    this.CommandListener = NullCommandListener_NULL_COMMAND_LISTENER
    
    this.TouchME = NewTouchJ2ME()
}


func NewMyCanvas(name string, childNameList IBasicArrayList) *MyCanvas {
    this := &MyCanvas{}
    this.Canvas = *NewCanvas()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        this.DisplayInfo.Update(this, this.CanvasStrings.CONSTRUCTOR)
        this.Name = name
        this.ChildNameList = childNameList
        this.CommandStack = StdUtil_GetInstance().CreateStack()
    return this
}


func (this *MyCanvas) SetFullScreenMode(mode bool) {
        this.Canvas.SetFullScreenMode(mode)
        this.DisplayInfo.Update(this, this.CanvasStrings.SET_FULL_SCREEN_MODE)
}


func (this *MyCanvas) SizeChanged(w int, h int) {
        this.DisplayInfo.Update(this, this.CanvasStrings.SIZE_CHANGED)
}


func (this *MyCanvas) GetCommandStack() *ABStack[any] {
        return this.CommandStack
}


// synchronized
func (this *MyCanvas) IsCommand(command ICommand) bool {
        return this.CommandStack.Contains(command)
}


// synchronized
func (this *MyCanvas) AddCommand(command ICommand) {
        if !this.CommandStack.Contains(command) {
            this.CommandStack.Push(command)
            this.Canvas.AddCommand(command)
        }
}


// synchronized
func (this *MyCanvas) RemoveCommand(command ICommand) {
        this.CommandStack.RemoveElement(command)
        this.Canvas.RemoveCommand(command)
}


// synchronized
func (this *MyCanvas) RemoveAllCommands() {
        var size int = this.CommandStack.Size(); _ = size
        for index := int(0); index < size; index++ {
            this.Canvas.RemoveCommand(any(this.CommandStack.Pop()).(ICommand))
        }
}


func (this *MyCanvas) SetCommandListener(l CommandListener) {
        this.Canvas.SetCommandListener(l)
        this.CommandListener = l
}


func (this *MyCanvas) GetCustomCommandListener() CommandListener {
        return this.CommandListener
}


// synchronized
func (this *MyCanvas) IsPaused() bool {
        return this.Paused
}


func (this *MyCanvas) RemovePauseCommand() {
        this.RemoveCommand(MyCommandsFactory_GetInstance().PAUSE_COMMAND)
}


// synchronized
func (this *MyCanvas) Pause() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CanvasStrings.PAUSE)
        this.RemovePauseCommand()
        this.AddCommand(MyCommandsFactory_GetInstance().RESUME_COMMAND)
        this.SetPaused(true)
}


// synchronized
func (this *MyCanvas) UnPause() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CanvasStrings.UN_PAUSE)
        this.RemoveCommand(MyCommandsFactory_GetInstance().RESUME_COMMAND)
        this.AddCommand(MyCommandsFactory_GetInstance().PAUSE_COMMAND)
        this.SetPaused(false)
}


// throws Exception
func (this *MyCanvas) Process() {
        this.DisplayInfo.Process()
}


func (this *MyCanvas) Paint(graphics IGraphics) {
}


func (this *MyCanvas) Draw(graphics IGraphics) {
}


func (this *MyCanvas) HasChild(displayable IMyCanvas) bool {
        return this.ChildNameList.Contains(displayable.MyCanvas_().Name)
}


func (this *MyCanvas) Destroy() {
        this.LogUtil.PutF("Destroyed MyCanvas", this, "destroy")
}


func (this *MyCanvas) SetPaused(isPaused bool) {
        this.Paused = isPaused
}


func (this *MyCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
}


func (this *MyCanvas) KeyRepeatedByDevice(keyCode int, deviceId int) {
}


func (this *MyCanvas) KeyReleasedByDevice(keyCode int, deviceId int) {
}


func (this *MyCanvas) PointerDragged(x int, y int) {
        this.TouchME.PointerDragged(x, y)
}


func (this *MyCanvas) PointerPressed(x int, y int) {
        this.TouchME.PointerPressed(x, y)
}


func (this *MyCanvas) PointerReleased(x int, y int) {
        this.TouchME.PointerReleased(x, y)
}


func (this *MyCanvas) NextSong(nextSongSound ISound, leftVolume int, rightVolume int) {
}

