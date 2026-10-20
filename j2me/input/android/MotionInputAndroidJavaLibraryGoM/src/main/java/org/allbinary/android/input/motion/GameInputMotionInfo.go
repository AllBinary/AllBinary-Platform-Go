/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2009 AllBinary
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

type GameInputMotionInfo struct {
    MAX_API int
    MAX_POINTERS int
}

// The Java class GameInputMotionInfo has subclasses so its variables are this interface
type IGameInputMotionInfo interface {
    GameInputMotionInfo_() *GameInputMotionInfo
}

func (this *GameInputMotionInfo) GameInputMotionInfo_() *GameInputMotionInfo {
    return this
}


func (this *GameInputMotionInfo) initFields() {
    
    
}


func NewGameInputMotionInfo(maxApi int, maxPointers int) *GameInputMotionInfo {
    this := &GameInputMotionInfo{}
    this.initFields()
        this.MAX_API = maxApi
        this.MAX_POINTERS = maxPointers
    return this
}

