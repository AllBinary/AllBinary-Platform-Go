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

type InputRobot struct {
    LogUtil *LogUtil
    Robot *Robot
}

// Implements InputRobotInterface

var InputRobot_NAME string = "Java Robot"

func (this *InputRobot) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewInputRobot(graphicsDevice *GraphicsDevice) *InputRobot {
    this := &InputRobot{}
    this.initFields()
        this.Robot = NewRobot(graphicsDevice)
    return this
}


func (this *InputRobot) GetName() string {
        return InputRobot_NAME
}


func (this *InputRobot) GetHelpSet() *HelpSet {
        return nil
}


func (this *InputRobot) GetMousePoint() *Point {
        var point *Point = MouseInfo_GetPointerInfo().GetLocation(); _ = point
        return point
}


// throws Exception
func (this *InputRobot) MouseMoveToTarget_1(rectangle *Rectangle, x int, y int) {
        var point *Point = PointHelper_GetCenterPoint(rectangle); _ = point
        this.Robot.MouseMove(point.X + x, point.Y + y)
        var message string = fmt.Sprint(NewStringMaker().Append("Moved Mouse To: x: ").Appendint(point.X).Append(" y: ").Appendint(point.Y).Append(" in the middle of: ").Append(StringUtil_GetInstance().ToString(rectangle))); _ = message
        this.LogUtil.PutF(message, this, "moveMouseToTarget")
}

// Java overloads of InputRobot.mouseMoveToTarget
func (this *InputRobot) MouseMoveToTarget(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[*Rectangle](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            this.MouseMoveToTarget_1(Java_Arg[*Rectangle](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]))
            return
        }
        if len(args) == 1 && Java_IsArg[*Rectangle](args[0], pass) {
            this.MouseMoveToTarget_2(Java_Arg[*Rectangle](args[0]))
            return
        }
    }
    panic("No Java overload of InputRobot.mouseMoveToTarget for the arguments")
}


func (this *InputRobot) MouseMove_1(point *Point) {
        this.MouseMove(point.X, point.Y)
}

// Java overloads of InputRobot.mouseMove
func (this *InputRobot) MouseMove(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Point](args[0], pass) {
            this.MouseMove_1(Java_Arg[*Point](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.MouseMove_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.MouseMove_3(Java_Arg[int](args[0]), Java_Arg[int](args[1]))
            return
        }
    }
    panic("No Java overload of InputRobot.mouseMove for the arguments")
}


func (this *InputRobot) MouseMove_2(x int, y int) {
        this.Robot.MouseMove(x, y)
        var message string = fmt.Sprint(NewStringMaker().Append("Moved Mouse To: x: ").Appendint(x).Append(" y: ").Appendint(y)); _ = message
        this.LogUtil.PutF(message, this, "moveMouse")
}


// throws Exception
func (this *InputRobot) MouseMoveToTarget_2(rectangle *Rectangle) {
        this.MouseMoveToTarget(rectangle, 0, 0)
}


func (this *InputRobot) CreateScreenCapture(screenRect *Rectangle) *BufferedImage {
        return this.Robot.CreateScreenCapture(screenRect)
}


func (this *InputRobot) Delay(ms int) {
        this.Robot.Delay(ms)
}


func (this *InputRobot) GetAutoDelay() int {
        return this.Robot.GetAutoDelay()
}


func (this *InputRobot) GetPixelColor(x int, y int) *Color {
        return this.Robot.GetPixelColor(x, y)
}


func (this *InputRobot) IsAutoWaitForIdle() bool {
        return this.Robot.IsAutoWaitForIdle()
}


func (this *InputRobot) KeyPress(keycode int) {
        this.Robot.KeyPress(keycode)
}


func (this *InputRobot) KeyRelease(keycode int) {
        this.Robot.KeyRelease(keycode)
}


func (this *InputRobot) MouseMove_3(x int, y int) {
        this.Robot.MouseMove(x, y)
}


func (this *InputRobot) MousePress(buttons int) {
        this.Robot.MousePress(buttons)
}


func (this *InputRobot) MouseRelease(buttons int) {
        this.Robot.MouseRelease(buttons)
}


func (this *InputRobot) MouseWheel(wheelAmt int) {
        this.Robot.MouseWheel(wheelAmt)
}


func (this *InputRobot) SetAutoDelay(ms int) {
        this.Robot.SetAutoDelay(ms)
}


func (this *InputRobot) SetAutoWaitForIdle(isOn bool) {
        this.Robot.SetAutoWaitForIdle(isOn)
}


func (this *InputRobot) ToString() string {
        return fmt.Sprint(this.Robot)
}

func (this *InputRobot) String() string {
    return this.ToString()
}


func (this *InputRobot) WaitForIdle() {
        this.Robot.WaitForIdle()
}


func (this *InputRobot) On() {
}


func (this *InputRobot) Off() {
}

