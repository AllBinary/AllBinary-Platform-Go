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

type KeyboardInputAutomationProcessor struct {
}



func NewKeyboardInputAutomationProcessor() *KeyboardInputAutomationProcessor {
    this := &KeyboardInputAutomationProcessor{}
    return this
}


// throws Exception
func KeyboardInputAutomationProcessor_Process(keyActionScriptInputInterface KeyboardActionScriptInputInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        logUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + fmt.Sprint(keyActionScriptInputInterface), "KeyInputAutomationProcessor", commonStrings.PROCESS)
        var inputRobotInterface InputRobotInterface = keyActionScriptInputInterface.GetInputRobotInterface(); _ = inputRobotInterface
        var integer []int = keyActionScriptInputInterface.GetKeyArray(); _ = integer
        if keyActionScriptInputInterface.GetDelayBetweenKeys() == 0 {
            for index := int(0); index < len(keyActionScriptInputInterface.GetKeyArray()); index++ {
                if keyActionScriptInputInterface.IsPress() {
                    inputRobotInterface.KeyPress(integer[index])
                }
            }
            time.Sleep(time.Duration(keyActionScriptInputInterface.GetTime()) * time.Millisecond)
            for index := int(len(keyActionScriptInputInterface.GetKeyArray()) - 1); index >= 0; index-- {
                if keyActionScriptInputInterface.IsRelease() {
                    inputRobotInterface.KeyRelease(integer[index])
                }
            }
        } else {
            for index := int(0); index < len(keyActionScriptInputInterface.GetKeyArray()); index++ {
                if keyActionScriptInputInterface.IsPress() {
                    inputRobotInterface.KeyPress(integer[index])
                }
                time.Sleep(time.Duration(keyActionScriptInputInterface.GetTime()) * time.Millisecond)
                if keyActionScriptInputInterface.IsRelease() {
                    inputRobotInterface.KeyRelease(integer[index])
                }
                var timeHelper ITimeDelayHelper = NewTimeDelayHelper(keyActionScriptInputInterface.GetDelayBetweenKeys()); _ = timeHelper
                var toLong int = 0; _ = toLong
                for !timeHelper.IsTimeTNT() || toLong > 6000 {
                    toLong++
                    time.Sleep(time.Duration(50) * time.Millisecond)
                }
            }
        }
}

