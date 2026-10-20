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

type Life struct {
    Lives int16
    StartLives int16
    Maxlives int16
}

// Implements LifeInterface

var Life_NO_LIFE *Life = Life_CreateLife()

func (this *Life) initFields() {
    
    
    
}


func Life_CreateLife() *Life {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Life) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(errors.New(fmt.Sprint()))
                }
            }()
            tryControl1, tryResult1 = 1, NewLife(int16(0), int16(0)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func NewLife(lives int16, maxlives int16) *Life {
    this := &Life{}
    this.initFields()
        this.Maxlives = maxlives
        if int(lives) > int(this.Maxlives) {
            panic(errors.New(fmt.Sprint("To Many Lives")))
        }
        this.Lives = lives
        this.StartLives = lives
    return this
}


func (this *Life) Get() int16 {
        return this.Lives
}


func (this *Life) Set(lives int16) {
        this.Lives = lives
}


func (this *Life) Reset() {
        this.Lives = this.StartLives
}


func (this *Life) Add(lives int16) {
        if !this.ReachedLimit() {
            if int(this.Lives) + int(lives) <= int(this.GetMaxlives()) {
                this.Lives = int16((int(this.Lives) + int(lives)))
            } else {
                this.Lives = this.GetMaxlives()
            }
        }
}


func (this *Life) IsAlive() bool {
        if int(this.Lives) <= 0 {
            return false
        } else {
            return true
        }
}


func (this *Life) GetMaxlives() int16 {
        return this.Maxlives
}


func (this *Life) ReachedLimit() bool {
        if int(this.Lives) > int(this.Maxlives) || int(this.Lives) < 0 {
            return true
        } else {
            return false
        }
}


func (this *Life) Deaths(lives int16) {
        if !this.ReachedLimit() {
            if int(this.Lives) - int(lives) >= 0 {
                this.Lives = int16((int(this.Lives) - int(lives)))
            } else {
                this.Lives = 0
            }
        }
}


func (this *Life) GetStartLives() int16 {
        return this.StartLives
}


func (this *Life) SetStartLives(startLives int16) {
        this.StartLives = startLives
}

