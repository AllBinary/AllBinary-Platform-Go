/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type GameTickTimeDelayHelper struct {
    TimeDelta int64
    StartTime int64
    LastStartTime int64
}


func (this *GameTickTimeDelayHelper) initFields() {
    
    
    
}


func NewGameTickTimeDelayHelper(startTime int64) *GameTickTimeDelayHelper {
    this := &GameTickTimeDelayHelper{}
    this.initFields()
        this.StartTime = startTime
        this.LastStartTime = 0
    return this
}


func (this *GameTickTimeDelayHelper) SetStartTime() int64 {
        this.StartTime = time.Now().UnixMilli()
        return this.StartTime
}


func (this *GameTickTimeDelayHelper) Loop() {
        if this.LastStartTime == Long_MIN_VALUE {
            this.TimeDelta = 0
        } else {
            this.TimeDelta = time.Now().UnixMilli() - this.LastStartTime
        }
}


func (this *GameTickTimeDelayHelper) GetTimeFromStart() int64 {
        return this.StartTime - this.LastStartTime
}
