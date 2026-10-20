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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type Timer struct {
    StartTime int64
    HundredthTime int64
    MaxTime int64
    Modifier int64
    HundrethsString []rune
    TotalDigits int
    PrimitiveLongUtil *PrimitiveLongUtil
    MAX []rune
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
}

// The Java class Timer has subclasses so its variables are this interface
type ITimer interface {
    Timer_() *Timer
    Start()
    Get() int64
    Update()
    Reduce(value int)
    Add(value int)
    GetTimeChars() []rune
    GetCurrentTotalDigits() int
}

func (this *Timer) Timer_() *Timer {
    return this
}


func (this *Timer) initFields() {
    this.StartTime = int64(-1)
    
    
    
    this.HundrethsString = PrimitiveLongSingleton_GetInstance().ZERO
    this.TotalDigits = 1
    
    this.MAX = []rune{'L', 'O', 'L'}
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
}


func NewTimer(maxTime int) *Timer {
    this := &Timer{}
    this.initFields()
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(maxTime)
        this.MaxTime = int64(maxTime)
    return this
}


func (this *Timer) Start() {
        this.HundredthTime = 0
        this.StartTime = this.GameTickTimeDelayHelper.StartTime
}


func (this *Timer) Get() int64 {
        return this.HundredthTime
}


func (this *Timer) Update() {
        var lastLowerTime int64 = this.HundredthTime; _ = lastLowerTime
        var currentTime int64 = ((this.Modifier + this.GameTickTimeDelayHelper.StartTime) - this.StartTime); _ = currentTime
        if currentTime < 0 {
            currentTime = 0
        }
        this.HundredthTime = currentTime / 100
        if this.HundredthTime >= this.MaxTime {
            if this.StartTime > 0 {
                this.HundrethsString = this.MAX
            }
        } else if lastLowerTime != this.HundredthTime {
            this.HundrethsString = this.PrimitiveLongUtil.GetCharArray(int(this.HundredthTime))
            this.TotalDigits = this.PrimitiveLongUtil.GetCurrentTotalDigits()
        }
}


func (this *Timer) Reduce(value int) {
        this.Modifier -= int64(value)
}


func (this *Timer) Add(value int) {
        this.Modifier += int64(value)
}


func (this *Timer) GetTimeChars() []rune {
        return this.HundrethsString
}


func (this *Timer) GetCurrentTotalDigits() int {
        return this.TotalDigits
}

