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

type TestInputForMotionRectanglesResultsWorker struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Index int64
    Running bool
    MotionRectanglesVector IBasicArrayList
}

// Implements MotionRectanglesResultsListener


func (this *TestInputForMotionRectanglesResultsWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewTestInputForMotionRectanglesResultsWorker() *TestInputForMotionRectanglesResultsWorker {
    this := &TestInputForMotionRectanglesResultsWorker{}
    this.initFields()
        this.MotionRectanglesVector = NewBasicArrayListD()
    return this
}


func (this *TestInputForMotionRectanglesResultsWorker) GetMotionRectanglesVector() IBasicArrayList {
        return this.MotionRectanglesVector
}


func (this *TestInputForMotionRectanglesResultsWorker) OnMotionRectanglesImageComparisonResultsEvent(motionRectanglesImageComparisonResultsEvent *MotionRectanglesResultsEvent) {
        this.GetMotionRectanglesVector().Add(motionRectanglesImageComparisonResultsEvent.GetMotionRectangles())
        this.Run()
}


func (this *TestInputForMotionRectanglesResultsWorker) OnEvent(allBinaryEventObject IAllBinaryEventObject) {
        this.OnMotionRectanglesImageComparisonResultsEvent(any(allBinaryEventObject).(*MotionRectanglesResultsEvent))
}


// synchronized
func (this *TestInputForMotionRectanglesResultsWorker) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *TestInputForMotionRectanglesResultsWorker) SetRunning(running bool) {
        this.Running = running
}


func (this *TestInputForMotionRectanglesResultsWorker) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.RUN)
            this.SetRunning(true)
            var timeHelper ITimeDelayHelper = NewTimeDelayHelper(1000); _ = timeHelper
            timeHelper.SetStartTimeTNT()
            var motionRectangles IMotionRectangles = any(this.GetMotionRectanglesVector().Get(0)).(IMotionRectangles); _ = motionRectangles
            var motionRectangleVector IBasicArrayList = motionRectangles.GetVector(); _ = motionRectangleVector
            if motionRectangleVector.Size() > 0 {
                var rectangle *Rectangle = any(motionRectangleVector.Get(0)).(*Rectangle); _ = rectangle
                var robotHashtable *ABHashtable[any, any] = InputRobotFactory_GetInstance().Get().(*ABHashtable); _ = robotHashtable
                var inputTypeNameArray []any = robotHashtable.KeySet().ToArray(); _ = inputTypeNameArray
                var size int = len(inputTypeNameArray); _ = size
                for index := int(0); index < size; index++ {
                    var inputTypeNameString string = any(inputTypeNameArray[index]).(string); _ = inputTypeNameString
                    var inputRobotInterface InputRobotInterface = any(robotHashtable.Get(inputTypeNameString)).(InputRobotInterface); _ = inputRobotInterface
                    inputRobotInterface.MouseMoveToTarget(rectangle)
                    if strings.Compare(inputRobotInterface.GetName(), TempInputRobotNames_DX_NAME) == 0 {
                        inputRobotInterface.MousePress(InputEvent_BUTTON1_MASK)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.MouseRelease(InputEvent_BUTTON1_MASK)
                        inputRobotInterface.MousePress(InputEvent_BUTTON2_MASK)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.MouseRelease(InputEvent_BUTTON2_MASK)
                        inputRobotInterface.MousePress(InputEvent_BUTTON3_MASK)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.MouseRelease(InputEvent_BUTTON3_MASK)
                        inputRobotInterface.KeyPress(KeyEvent_VK_1)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.KeyRelease(KeyEvent_VK_1)
                        inputRobotInterface.KeyPress(KeyEvent_VK_2)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.KeyRelease(KeyEvent_VK_2)
                        inputRobotInterface.KeyPress(KeyEvent_VK_3)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.KeyRelease(KeyEvent_VK_3)
                        inputRobotInterface.KeyPress(KeyEvent_VK_4)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.KeyRelease(KeyEvent_VK_4)
                        inputRobotInterface.KeyPress(KeyEvent_VK_F1)
                        time.Sleep(time.Duration(300) * time.Millisecond)
                        inputRobotInterface.KeyRelease(KeyEvent_VK_F1)
                    }
                }
            }
            this.Index++
            this.GetMotionRectanglesVector().Remove(motionRectangles)
            this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

