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

type StartIntermissionProcessor struct {
    Processor
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameCanvas IAllBinaryGameCanvas
    WAIT int64
}


func (this *StartIntermissionProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.WAIT = 5000
}


func NewStartIntermissionProcessor(gameCanvas IAllBinaryGameCanvas) *StartIntermissionProcessor {
    this := &StartIntermissionProcessor{}
    this.Processor = *NewProcessor()
    this.initFields()
        this.GameCanvas = gameCanvas
    return this
}


// throws Exception
func (this *StartIntermissionProcessor) Process() {
        if this.GameCanvas.GetStartIntermissionInterface().GetTimeDelayHelper().IsElapsedTNT(this.WAIT) {
            this.LogUtil.PutF("Intermission End", this, this.CommonStrings.PROCESS)
            this.GameCanvas.GetStartIntermissionInterface().SetEnabled(false)
        }
}

