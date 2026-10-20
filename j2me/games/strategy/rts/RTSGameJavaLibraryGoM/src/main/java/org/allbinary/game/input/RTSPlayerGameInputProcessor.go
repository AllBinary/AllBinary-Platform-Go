/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSPlayerGameInputProcessor struct {
    GameInputProcessor
    RtsPlayerGameInput IRTSPlayerGameInput
}

// The Java class RTSPlayerGameInputProcessor has subclasses so its variables are this interface
type IRTSPlayerGameInputProcessor interface {
    IGameInputProcessor
    RTSPlayerGameInputProcessor_() *RTSPlayerGameInputProcessor
}

func (this *RTSPlayerGameInputProcessor) RTSPlayerGameInputProcessor_() *RTSPlayerGameInputProcessor {
    return this
}


func (this *RTSPlayerGameInputProcessor) initFields() {
    
}


func NewRTSPlayerGameInputProcessor(rtsPlayerGameInput IRTSPlayerGameInput) *RTSPlayerGameInputProcessor {
    this := &RTSPlayerGameInputProcessor{}
    this.GameInputProcessor = *NewGameInputProcessor()
    this.initFields()
        this.RtsPlayerGameInput = rtsPlayerGameInput
    return this
}

