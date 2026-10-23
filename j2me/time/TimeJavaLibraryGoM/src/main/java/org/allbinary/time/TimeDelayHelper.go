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

type TimeDelayHelper struct {
    StartTime int64
    ElapsedTimeAtPause int64
    Delay int
}

// The Java class TimeDelayHelper has subclasses so its variables are this interface
type ITimeDelayHelper interface {
    TimeDelayHelper_() *TimeDelayHelper
    IsTimeTNT() bool
    IsTimeSinceTNT(delay int) bool
    IsTime(currentTime int64) bool
    IsTimeWithoutReset(currentTime int64) bool
    IsTimeSince(delay int, currentTime int64) bool
    GetElapsedTNT() int64
    GetElapsed(currentTime int64) int64
    IsElapsed(currentTime int64, time int64) bool
    IsElapsedTNT(time int64) bool
    SetStartTime(startTime int64)
    GetStartTime() int64
    SetStartTimeTNT()
    Pause()
    Resume() bool
    ToStringAt(currentTime int64) string
    ToString() string
}

func (this *TimeDelayHelper) TimeDelayHelper_() *TimeDelayHelper {
    return this
}


func (this *TimeDelayHelper) initFields() {
    this.StartTime = int64(-1)
    this.ElapsedTimeAtPause = Long_MIN_VALUE
    
}


func NewTimeDelayHelper(delay int) *TimeDelayHelper {
    this := &TimeDelayHelper{}
    this.initFields()
        this.Delay = delay
        this.SetStartTimeTNT()
    return this
}


func (this *TimeDelayHelper) IsTimeTNT() bool {
        var currentTime int64 = time.Now().UnixMilli(); _ = currentTime
        if currentTime - this.StartTime > int64(this.Delay) {
            this.StartTime = currentTime
            return true
        }
        return false
}


func (this *TimeDelayHelper) IsTimeSinceTNT(delay int) bool {
        var currentTime int64 = time.Now().UnixMilli(); _ = currentTime
        if currentTime - this.StartTime > int64(delay) {
            this.StartTime = currentTime
            return true
        }
        return false
}


func (this *TimeDelayHelper) IsTime(currentTime int64) bool {
        if currentTime - this.StartTime > int64(this.Delay) {
            this.StartTime = currentTime
            return true
        }
        return false
}


func (this *TimeDelayHelper) IsTimeWithoutReset(currentTime int64) bool {
        if currentTime - this.StartTime > int64(this.Delay) {
            return true
        }
        return false
}


func (this *TimeDelayHelper) IsTimeSince(delay int, currentTime int64) bool {
        if currentTime - this.StartTime > int64(this.Delay) {
            this.StartTime = currentTime
            return true
        }
        return false
}


func (this *TimeDelayHelper) GetElapsedTNT() int64 {
        return time.Now().UnixMilli() - this.StartTime
}


func (this *TimeDelayHelper) GetElapsed(currentTime int64) int64 {
        return currentTime - this.StartTime
}


func (this *TimeDelayHelper) IsElapsed(currentTime int64, time int64) bool {
        if this.GetElapsed(currentTime) > time {
            return true
        } else {
            return false
        }
}


func (this *TimeDelayHelper) IsElapsedTNT(time int64) bool {
        if this.GetElapsedTNT() > time {
            return true
        } else {
            return false
        }
}


func (this *TimeDelayHelper) SetStartTime(startTime int64) {
        this.StartTime = startTime
}


func (this *TimeDelayHelper) GetStartTime() int64 {
        return this.StartTime
}


func (this *TimeDelayHelper) SetStartTimeTNT() {
        this.StartTime = time.Now().UnixMilli()
}


func (this *TimeDelayHelper) Pause() {
        this.ElapsedTimeAtPause = this.GetElapsedTNT()
        this.StartTime = Long_MAX_VALUE
}


func (this *TimeDelayHelper) Resume() bool {
        if this.ElapsedTimeAtPause != Long_MIN_VALUE {
            this.StartTime = time.Now().UnixMilli() + this.ElapsedTimeAtPause
            this.ElapsedTimeAtPause = Long_MIN_VALUE
            this.SetStartTimeTNT()
            return true
        }
        return false
}


func (this *TimeDelayHelper) ToStringAt(currentTime int64) string {
        var elapsed int64 = this.GetElapsed(currentTime); _ = elapsed
        if elapsed > 0 {
            var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
            var elapsedAsString string = fmt.Sprint(elapsed); _ = elapsedAsString
            return fmt.Sprint(NewStringMaker().Append(commonLabels.START).Appendlong(this.StartTime).Append(commonLabels.CURRENT).Appendlong(currentTime).Append(commonLabels.ELAPSED).Append(elapsedAsString))
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}


func (this *TimeDelayHelper) ToString() string {
        var currentTime int64 = time.Now().UnixMilli(); _ = currentTime
        var elapsed int64 = this.GetElapsed(currentTime); _ = elapsed
        if elapsed > 0 {
            var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
            var elapsedAsString string = fmt.Sprint(elapsed); _ = elapsedAsString
            return fmt.Sprint(NewStringMaker().Append(commonLabels.START).Appendlong(this.StartTime).Append(commonLabels.CURRENT).Appendlong(currentTime).Append(commonLabels.ELAPSED).Append(elapsedAsString))
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}

func (this *TimeDelayHelper) String() string {
    return this.ToString()
}

