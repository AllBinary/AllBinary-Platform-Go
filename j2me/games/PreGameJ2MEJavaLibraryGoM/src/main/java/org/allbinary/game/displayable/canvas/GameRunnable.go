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

type GameRunnable struct {
    WAIT int
    FAST int
    WaitInMillis int64
}

// Implements Runnable

// The Java class GameRunnable has subclasses so its variables are this interface
type IGameRunnable interface {
    GameRunnable_() *GameRunnable
    Run()
    ProcessLoopSleep()
}

func (this *GameRunnable) GameRunnable_() *GameRunnable {
    return this
}


func (this *GameRunnable) initFields() {
    this.WAIT = 240
    this.FAST = 60
    this.WaitInMillis = int64(this.WAIT)
}


func NewGameRunnable() *GameRunnable {
    this := &GameRunnable{}
    this.initFields()
    return this
}


func (this *GameRunnable) Run() {
        DisplayInfoSingleton_GetInstance().Process()
}


// throws Exception
func (this *GameRunnable) ProcessLoopSleep() {
        time.Sleep(time.Duration(this.WaitInMillis) * time.Millisecond)
}

