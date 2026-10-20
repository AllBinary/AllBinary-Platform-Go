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

//BasicPlayerMIDP2JavaLibrary
type BasicPlayer struct {
    ListenersList IBasicArrayList
    TimeBase TimeBase
    StateP int
    LoopCountP int
}

// Implements *Player, TimeBaseInterface, Controllable2

// The Java class BasicPlayer has subclasses so its variables are this interface
type IBasicPlayer interface {
    BasicPlayer_() *BasicPlayer
    GetContentType() string
    GetControl(controlType string) Control
    GetControls() []Control
    Close()
    AddPlayerListener(playerListener PlayerListener)
    RemovePlayerListener(playerListener PlayerListener)
    GetState() int
    SetState(state int)
    GetDuration() int64
    GetMediaTime() int64
    GetTimeBase() TimeBase
    SetTimeBase(timeBase TimeBase)
    Deallocate()
    Prefetch()
    Realize()
    SetLoopCount(count int)
    GetLoopCount() int
    SetMediaTime(now int64) int64
    Start()
    Stop()
    SetVolume(leftVolume int, rightVolume int)
}

func (this *BasicPlayer) BasicPlayer_() *BasicPlayer {
    return this
}

var BasicPlayer_CONTROL_TYPE string = "ToneControl"

func (this *BasicPlayer) initFields() {
    this.ListenersList = NewBasicArrayListD()
    
    
    
}


func NewBasicPlayer() *BasicPlayer {
    this := &BasicPlayer{}
    this.initFields()
        this.SetLoopCount(0)
        this.SetState(Player_UNREALIZED)
    return this
}


func (this *BasicPlayer) GetContentType() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *BasicPlayer) GetControl(controlType string) Control {
        return NewNullControl()
}


func (this *BasicPlayer) GetControls() []Control {
        return make([]Control, 0)
}


func (this *BasicPlayer) Close() {
}


// synchronized
func (this *BasicPlayer) AddPlayerListener(playerListener PlayerListener) {
        if !this.ListenersList.Contains(playerListener) {
            this.ListenersList.Add(playerListener)
        }
}


func (this *BasicPlayer) RemovePlayerListener(playerListener PlayerListener) {
        this.ListenersList.Remove(playerListener)
}


func (this *BasicPlayer) GetState() int {
        return this.StateP
}


// synchronized
func (this *BasicPlayer) SetState(state int) {
        this.StateP = state
}


func (this *BasicPlayer) GetDuration() int64 {
        return 0
}


func (this *BasicPlayer) GetMediaTime() int64 {
        return 0
}


func (this *BasicPlayer) GetTimeBase() TimeBase {
        return this.TimeBase
}


// synchronized
func (this *BasicPlayer) SetTimeBase(timeBase TimeBase) {
        this.TimeBase = timeBase
}


func (this *BasicPlayer) Deallocate() {
}


// throws MediaException
func (this *BasicPlayer) Prefetch() {
}


// throws MediaException
func (this *BasicPlayer) Realize() {
}


// synchronized
func (this *BasicPlayer) SetLoopCount(count int) {
        this.LoopCountP = count
}


func (this *BasicPlayer) GetLoopCount() int {
        return this.LoopCountP
}


// throws MediaException
// synchronized
func (this *BasicPlayer) SetMediaTime(now int64) int64 {
        return 0
}


// throws MediaException
// synchronized
func (this *BasicPlayer) Start() {
        this.SetState(Player_STARTED)
}


// throws MediaException
// synchronized
func (this *BasicPlayer) Stop() {
        this.SetState(Player_PREFETCHED)
}


func (this *BasicPlayer) SetVolume(leftVolume int, rightVolume int) {
}

