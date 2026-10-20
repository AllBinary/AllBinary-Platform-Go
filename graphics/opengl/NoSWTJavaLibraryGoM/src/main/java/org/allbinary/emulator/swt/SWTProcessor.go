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

type SWTProcessor struct {
    Runnable Runnable
}

// The Java class SWTProcessor has subclasses so its variables are this interface
type ISWTProcessor interface {
    SWTProcessor_() *SWTProcessor
    Process(display *Display)
}

func (this *SWTProcessor) SWTProcessor_() *SWTProcessor {
    return this
}

var SWTProcessor_Instance ISWTProcessor = NewSWTProcessor()

func (this *SWTProcessor) initFields() {
    this.Runnable = NullRunnable_GetInstance()
}

func NewSWTProcessor() *SWTProcessor {
    this := &SWTProcessor{}
    this.initFields()
    return this
}


func SWTProcessor_GetInstance() ISWTProcessor {
        return SWTProcessor_Instance
}


func (this *SWTProcessor) Process(display *Display) {
        display.Sleep()
}

