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

//In general allow scrolling of the menu and selection the center item
type CommandFormInputProcessor struct {
    BasicMenuInputProcessor
    MOTION_GESTURE_SOURCE_ID int
    CLICK_DELAY int
    ClickTimeHelper ITimeDelayHelper
    DOUBLE_CLICK_DELAY int
    DoubleClickTimeHelper ITimeDelayHelper
    IsSingleKeyProcessing bool
    Form IScrollSelectionForm
    HasPressed bool
    PROCESS_COMMAND string
    NAME_LABEL string
}

// The Java class CommandFormInputProcessor has subclasses so its variables are this interface
type ICommandFormInputProcessor interface {
    IBasicMenuInputProcessor
    CommandFormInputProcessor_() *CommandFormInputProcessor
    ProcessInput(key int) int
    ProcessCommand() int
    ProcessInputList() int
    ProcessMotionInputs() int
    ProcessMotionInput(motionGestureEvent *MotionGestureEvent)
    ToString() string
}

func (this *CommandFormInputProcessor) CommandFormInputProcessor_() *CommandFormInputProcessor {
    return this
}


func (this *CommandFormInputProcessor) initFields() {
    this.MOTION_GESTURE_SOURCE_ID = GameKeyEventFactory_GetInstance().MOTION_GESTURE_SOURCE_ID
    this.CLICK_DELAY = 150
    this.ClickTimeHelper = NewTimeDelayHelper(this.CLICK_DELAY)
    this.DOUBLE_CLICK_DELAY = 1200
    this.DoubleClickTimeHelper = NewTimeDelayHelper(this.DOUBLE_CLICK_DELAY)
    this.IsSingleKeyProcessing = InputFeatureFactory_GetInstance().IsSingleKeyProcessing()
    
    this.HasPressed = false
    this.PROCESS_COMMAND = "processCommand"
    this.NAME_LABEL = " ScrollSelectionForm: "
}


func NewCommandFormInputProcessor(gameKeyEventList IBasicArrayList, playerInputId int, gameCanvas IMyCanvas, form IScrollSelectionForm) *CommandFormInputProcessor {
    this := &CommandFormInputProcessor{}
        this.BasicMenuInputProcessor = *NewBasicMenuInputProcessor(gameKeyEventList, playerInputId, gameCanvas)
        this.initFields()
        this.Form = form
    return this
}


// throws Exception
func (this *CommandFormInputProcessor) ProcessInput(key int) int {
        if key == Canvas_LEFT || key == Canvas_RIGHT || key == Canvas_UP || key == Canvas_DOWN {
            PrimaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
            this.Form.ProcessInputKey(key)
            return 1
        } else if key == Canvas_FIRE {
            PrimaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
            return this.ProcessCommand()
        }
        return 0
}


func (this *CommandFormInputProcessor) ProcessCommand() int {
        var commandCurrentSelectionForm *CommandCurrentSelectionForm = any(this.Form).(*CommandCurrentSelectionForm); _ = commandCurrentSelectionForm
        var command ICommand = commandCurrentSelectionForm.GetSelectedCommand(); _ = command
        this.LogUtil.PutF(fmt.Sprint(command), this, this.PROCESS_COMMAND)
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if SWTUtil_IsSWT && features.IsFeature(openGLFeatureFactory.OPENGL) && command != GameCommandsFactory_GetInstance().EXIT_COMMAND {
            NewCommandRunnable(this, command).Run()
        } else {
            PrimaryThreadPool_GetInstance().RunTask(NewCommandRunnable(this, command))
        }
        if command == GameCommandsFactory_GetInstance().QUIT_COMMAND {
            return -1
        } else {
            return 1
        }
}


// throws Exception
func (this *CommandFormInputProcessor) ProcessInputList() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.PROCESS_INPUT, e)
            tryControl1, tryResult1 = 1, -1; return
                }
            }()
            var motionInputsIndex int = this.ProcessMotionInputs(); _ = motionInputsIndex
            var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
            var size int = list.Size(); _ = size
            var key int = 0; _ = key
            var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
            for index := int(0); index < size; index++ {
                gameKeyEvent = any(list.BasicArrayList_().ObjectArray[index]).(*GameKeyEvent)
                key = gameKeyEvent.GetKey()
                if gameKeyEvent.GetSourceId() != this.MOTION_GESTURE_SOURCE_ID {
                    if this.ProcessInput(key) == 1 {
                        break
                    }
                }
            }
            this.Clear()
            if size > 0 || motionInputsIndex >= 0 {
                tryControl1, tryResult1 = 1, 1; return
            } else {
                tryControl1, tryResult1 = 1, -1; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *CommandFormInputProcessor) ProcessMotionInputs() int {
        var lastIndex int = this.MotionGestureEventList.Size() - 1; _ = lastIndex
        if lastIndex >= 0 {
            var motionGestureEvent *MotionGestureEvent = any(this.MotionGestureEventList.BasicArrayList_().ObjectArray[lastIndex]).(*MotionGestureEvent); _ = motionGestureEvent
            this.ProcessMotionInput(motionGestureEvent)
        }
        this.MotionGestureEventList.Clear()
        return lastIndex
}


// throws Exception
func (this *CommandFormInputProcessor) ProcessMotionInput(motionGestureEvent *MotionGestureEvent) {
        var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
        var motionGestureInput *MotionGestureInput = motionGestureEvent.GetMotionGesture(); _ = motionGestureInput
        if motionGestureInput == touchMotionGestureFactory.RELEASED {
            var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
            if this.Form.IsInForm(point) {
                var index int = this.Form.GetSelectedIndexForPoint(point); _ = index
                if index != -1 {
                    PrimaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
                    if index == this.Form.GetSelectedIndex() {
                        if this.ClickTimeHelper.IsTimeTNT() {
                            this.ProcessCommand()
                        }
                    } else {
                        this.Form.SetSelectedIndex(index)
                    }
                } else {
                }
            }
            if this.HasPressed {
                if !this.DoubleClickTimeHelper.IsTimeTNT() {
                    this.LogUtil.PutF("Double Press", this, this.GameInputStrings.PROCESS_MOTION_INPUT)
                    this.ProcessCommand()
                }
                this.DoubleClickTimeHelper.TimeDelayHelper_().Delay = this.DOUBLE_CLICK_DELAY
                this.DoubleClickTimeHelper.SetStartTimeTNT()
            }
            this.HasPressed = false
        } else if motionGestureInput == touchMotionGestureFactory.PRESSED {
            this.DoubleClickTimeHelper.TimeDelayHelper_().Delay = 0
            this.HasPressed = true
        }
}


func (this *CommandFormInputProcessor) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.BasicMenuInputProcessor.ToString()).Append(this.NAME_LABEL).Append(fmt.Sprint(this.Form)))
}

func (this *CommandFormInputProcessor) String() string {
    return this.ToString()
}

