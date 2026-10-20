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

type MultiplayerBehavior struct {
}

// The Java class MultiplayerBehavior has subclasses so its variables are this interface
type IMultiplayerBehavior interface {
    MultiplayerBehavior_() *MultiplayerBehavior
    StartTick()
    EndTick()
    GetApproximateElapsed() int64
}

func (this *MultiplayerBehavior) MultiplayerBehavior_() *MultiplayerBehavior {
    return this
}

var MultiplayerBehavior_NULL_MULTIPLAYER_BEHAVIOR IMultiplayerBehavior = NewMultiplayerBehavior(StringUtil_GetInstance().EMPTY_STRING)


func NewMultiplayerBehavior(username string) *MultiplayerBehavior {
    this := &MultiplayerBehavior{}
    return this
}


func (this *MultiplayerBehavior) StartTick() {
}


func (this *MultiplayerBehavior) EndTick() {
}


func (this *MultiplayerBehavior) GetApproximateElapsed() int64 {
        return 0
}

