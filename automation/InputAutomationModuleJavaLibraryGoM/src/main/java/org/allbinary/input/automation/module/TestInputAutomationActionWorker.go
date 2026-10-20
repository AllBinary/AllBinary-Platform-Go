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

type TestInputAutomationActionWorker struct {
    AbstractInputAutomationWorker
    LogUtil *LogUtil
    Sleep int
}


func (this *TestInputAutomationActionWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewTestInputAutomationActionWorker(sleep int, inputAutomationActionInterface InputAutomationActionInterface, imageComparatorConstraintsInterface ImageComparatorConstraintsInterface, motionRectangleConstraintsInterface MotionRectangleConstraintsInterface) *TestInputAutomationActionWorker {
    this := &TestInputAutomationActionWorker{}
        this.AbstractInputAutomationWorker = *NewAbstractInputAutomationWorker(inputAutomationActionInterface)
        this.initFields()
        this.LogUtil.PutF("BasicTestAttackWorder", this, this.CommonStrings.CONSTRUCTOR)
        this.Sleep = sleep
    return this
}


// throws Exception
func (this *TestInputAutomationActionWorker) ProcessBetweenWorkers() {
        time.Sleep(time.Duration(2000) * time.Millisecond)
}


// throws Exception
func (this *TestInputAutomationActionWorker) Process() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
        var hashtable *ABHashtable[any, any] = InputRobotFactory_GetInstance().Get().(*ABHashtable); _ = hashtable
        var set Set = hashtable.KeySet(); _ = set
        var inputTypeNameArray []any = set.ToArray(); _ = inputTypeNameArray
        var size int = len(inputTypeNameArray); _ = size
        for index := int(0); index < size; index++ {
            var inputTypeNameString string = any(inputTypeNameArray[index]).(string); _ = inputTypeNameString
            var robot InputRobotInterface = any(hashtable.Get(inputTypeNameString)).(InputRobotInterface); _ = robot
            if this.Sleep > 0 {
                time.Sleep(time.Duration(this.Sleep) * time.Millisecond)
            }
        }
}

