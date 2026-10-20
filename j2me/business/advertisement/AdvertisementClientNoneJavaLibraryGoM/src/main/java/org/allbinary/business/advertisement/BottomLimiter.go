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

type BottomLimiter struct {
    Bottom int
    DeltaY int
    Lower int
    Higher int
}

var BottomLimiter_Instance *BottomLimiter = NewBottomLimiter()

func (this *BottomLimiter) initFields() {
    
    
    
    
}


func BottomLimiter_GetInstance() *BottomLimiter {
        return BottomLimiter_Instance
}


func NewBottomLimiter() *BottomLimiter {
    this := &BottomLimiter{}
    this.initFields()
    return this
}


func (this *BottomLimiter) Init(lower int, higher int) {
        this.Lower = lower
        this.Higher = higher
}


func (this *BottomLimiter) Update() {
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var gameAdState *GameAdState = GameAdStateFactory_GetInstance().GetCurrentInstance(); _ = gameAdState
        if gameAdState.IsShowingAt() {
            this.DeltaY = this.Higher
        } else {
            this.DeltaY = this.Lower
        }
        this.Bottom = displayInfo.GetLastHeight() - this.DeltaY
}


func (this *BottomLimiter) GetDeltaY() int {
        return this.DeltaY
}


func (this *BottomLimiter) GetBottom() int {
        return this.Bottom
}

