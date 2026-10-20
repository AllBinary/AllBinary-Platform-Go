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

type LapInfo struct {
    NumberOfLaps int
    TotalLaps int
}

var LapInfo_NULL_LAP_INFO *LapInfo = NewLapInfo(-1)

func (this *LapInfo) initFields() {
    
    
}


func NewLapInfo(totalLaps int) *LapInfo {
    this := &LapInfo{}
    this.initFields()
        this.TotalLaps = totalLaps
        this.NumberOfLaps = 1
    return this
}


func (this *LapInfo) Add() {
        this.NumberOfLaps++
}


func (this *LapInfo) GetCurrentLap() int {
        return this.NumberOfLaps
}


func (this *LapInfo) GetTotalLaps() int {
        return this.TotalLaps
}


func (this *LapInfo) IsDone() bool {
        if this.NumberOfLaps >= this.TotalLaps {
            return true
        }
        return false
}


func (this *LapInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("isDone: ")
        stringBuffer.Appendboolean(this.IsDone())
        stringBuffer.Append(" CurrentLap: ")
        stringBuffer.Appendint(this.GetCurrentLap())
        stringBuffer.Append(" Total Laps: ")
        stringBuffer.Appendint(this.GetTotalLaps())
        return fmt.Sprint(stringBuffer)
}

func (this *LapInfo) String() string {
    return this.ToString()
}

