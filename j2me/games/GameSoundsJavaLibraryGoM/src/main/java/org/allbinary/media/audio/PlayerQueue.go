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

type PlayerQueue struct {
    CommonStrings *CommonStrings
    List IBasicArrayList
    Max int
}

var PlayerQueue_TOTAL string = "Sounds In Queue: "

func (this *PlayerQueue) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.List = NewBasicArrayListD()
    
}


func NewPlayerQueue(max_ int) *PlayerQueue {
    this := &PlayerQueue{}
    this.initFields()
        this.Max = max_
    return this
}


func (this *PlayerQueue) Add(sound ISound) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.ADD)
                }
            }()
            if !this.List.Contains(sound) {
                this.List.Add(sound)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerQueue) Process() bool {
        var sound ISound = NoSound_GetInstance(); _ = sound
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var resource string = StringUtil_GetInstance().EMPTY_STRING; _ = resource
            if sound != nil {
                resource = sound.GetResource()
            }
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(resource)), this, this.CommonStrings.PROCESS)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var played bool = false; _ = played
            if this.List.Size() > 0 {
                sound = any(this.List.RemoveAt(0)).(ISound)
                if sound != nil {
                    sound.GetPlayerP().Start()
                }
                for this.List.Size() > this.Max {
                    this.List.RemoveAt(0)
                }
                played = true
            }
            tryControl1, tryResult1 = 1, played; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PlayerQueue) Clear() {
        this.List.Clear()
}


func (this *PlayerQueue) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(PlayerQueue_TOTAL).Appendint(this.List.Size()))
}

func (this *PlayerQueue) String() string {
    return this.ToString()
}

