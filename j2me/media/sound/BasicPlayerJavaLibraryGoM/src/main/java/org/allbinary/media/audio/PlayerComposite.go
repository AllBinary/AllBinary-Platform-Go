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

//BasicPlayer
type PlayerComposite struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Player *Player
    TimeElapsedHelper ITimeDelayHelper
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
}

// Implements Controllable, *Player


func (this *PlayerComposite) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.TimeElapsedHelper = NewTimeDelayHelper(0)
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
}


func NewPlayerComposite(player *Player, repeatTime int) *PlayerComposite {
    this := &PlayerComposite{}
    this.initFields()
        this.Player = player
        this.TimeElapsedHelper.TimeDelayHelper_().Delay = repeatTime
    return this
}


// synchronized
func (this *PlayerComposite) AddPlayerListener(playerListener PlayerListener) {
        this.Player.AddPlayerListener(playerListener)
}


func (this *PlayerComposite) RemovePlayerListener(playerListener PlayerListener) {
        this.Player.RemovePlayerListener(playerListener)
}


func (this *PlayerComposite) Close() {
        this.Player.Close()
}


func (this *PlayerComposite) Deallocate() {
        this.Player.Deallocate()
}


func (this *PlayerComposite) GetContentType() string {
        return this.Player.GetContentType()
}


func (this *PlayerComposite) GetDuration() int64 {
        return this.Player.GetDuration()
}


func (this *PlayerComposite) GetMediaTime() int64 {
        return this.Player.GetMediaTime()
}


func (this *PlayerComposite) Prefetch() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "prefetch", e)
                }
            }()
            this.Player.Prefetch()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerComposite) Realize() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "realize", e)
                }
            }()
            this.Player.Realize()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerComposite) GetState() int {
        return this.Player.GetState()
}


func (this *PlayerComposite) SetLoopCount(count int) {
        this.Player.SetLoopCount(count)
}


func (this *PlayerComposite) SetMediaTime(now int64) int64 {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int64) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setMediaTime", e)
            tryControl1, tryResult1 = 1, int64(-1); return
                }
            }()
            tryControl1, tryResult1 = 1, this.Player.SetMediaTime(now); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PlayerComposite) Start() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.START_METHOD_NAME, e)
                }
            }()
            if this.TimeElapsedHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
                this.Player.Start()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerComposite) Stop() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "stop", e)
                }
            }()
            this.Player.Stop()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerComposite) GetControl(controlType string) Control {
        return this.Player.GetControl(controlType)
}


func (this *PlayerComposite) GetControls() []Control {
        panic(errors.New(fmt.Sprint()))
}


func (this *PlayerComposite) SetVolume(leftVolume int, rightVolume int) {
        var controllable2 Controllable2 = (any(this.Player).(Controllable2)); _ = controllable2
        controllable2.SetVolume(leftVolume, rightVolume)
}


func (this *PlayerComposite) GetPlayerP() *Player {
        return this.Player
}

