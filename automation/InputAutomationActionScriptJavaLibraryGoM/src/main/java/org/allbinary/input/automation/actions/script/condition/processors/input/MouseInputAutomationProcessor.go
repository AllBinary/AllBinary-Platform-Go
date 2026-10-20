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

type MouseInputAutomationProcessor struct {
}



func NewMouseInputAutomationProcessor() *MouseInputAutomationProcessor {
    this := &MouseInputAutomationProcessor{}
    return this
}


// throws Exception
func MouseInputAutomationProcessor_Process(mouseActionScriptInputInterface MouseActionScriptInputInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        logUtil.PutF(fmt.Sprint(NewStringMaker().Append("Start - Processing Mouse Input at: ").Append(fmt.Sprint(mouseActionScriptInputInterface.GetPoint())).Append(" clicks: ").Appendint(mouseActionScriptInputInterface.GetButtonClicks())), "MouseInputAutomationProcessor", commonStrings.PROCESS)
        var inputRobotInterface InputRobotInterface = mouseActionScriptInputInterface.GetInputRobotInterface(); _ = inputRobotInterface
        if mouseActionScriptInputInterface.GetPoint().X != -1 && mouseActionScriptInputInterface.GetPoint().Y != -1 {
            inputRobotInterface.MouseMove(mouseActionScriptInputInterface.GetPoint())
        }
        time.Sleep(time.Duration(mouseActionScriptInputInterface.GetTime()) * time.Millisecond)
        if mouseActionScriptInputInterface.GetButtonClicks() != 0 {
            inputRobotInterface.MousePress(mouseActionScriptInputInterface.GetButtonClicks())
            time.Sleep(time.Duration(mouseActionScriptInputInterface.GetTime()) * time.Millisecond)
            inputRobotInterface.MouseRelease(mouseActionScriptInputInterface.GetButtonClicks())
        }
        time.Sleep(time.Duration(mouseActionScriptInputInterface.GetTime()) * time.Millisecond)
}

