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

type NullInputRobot struct {
    AbstractInputRobot
    LogUtil *LogUtil
}

// Implements InputRobotInterface

var NullInputRobot_LIBRARY_NAME string = "null"
var NullInputRobot_NAME string = "NULL"

func (this *NullInputRobot) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewNullInputRobot() *NullInputRobot {
    this := &NullInputRobot{}
        this.AbstractInputRobot = *NewAbstractInputRobot(nil)
        this.initFields()
    return this
}


func (this *NullInputRobot) GetName() string {
        return NullInputRobot_NAME
}


func (this *NullInputRobot) GetHelpSet() *HelpSet {
        return nil
}


func (this *NullInputRobot) GetMousePoint() *Point {
        return PointFactory_GetInstance(0, 0)
}


func (this *NullInputRobot) KeyPress(keycode int) {
        this.LogUtil.PutF("KeyCode: " + fmt.Sprint(keycode), this, "keyPress")
}


func (this *NullInputRobot) KeyRelease(keycode int) {
        this.LogUtil.PutF("KeyCode: " + fmt.Sprint(keycode), this, "keyRelease")
}


func (this *NullInputRobot) MouseMove(x int, y int) {
        this.LogUtil.PutF("X: " + fmt.Sprint(x) + " Y: " + fmt.Sprint(y), this, "mouseMove")
}


func (this *NullInputRobot) MousePress(buttons int) {
        this.LogUtil.PutF("Buttons: " + fmt.Sprint(buttons), this, "mousePress")
}


func (this *NullInputRobot) MouseRelease(buttons int) {
        this.LogUtil.PutF("Buttons: " + fmt.Sprint(buttons), this, "mouseRelease")
}


func (this *NullInputRobot) CreateScreenCapture(screenRect *Rectangle) *BufferedImage {
        return nil
}


func (this *NullInputRobot) Delay(ms int) {
}


func (this *NullInputRobot) GetAutoDelay() int {
        return -1
}


func (this *NullInputRobot) GetPixelColor(x int, y int) *Color {
        return nil
}


func (this *NullInputRobot) IsAutoWaitForIdle() bool {
        return false
}


func (this *NullInputRobot) MouseWheel(wheelAmt int) {
}


func (this *NullInputRobot) SetAutoDelay(ms int) {
}


func (this *NullInputRobot) SetAutoWaitForIdle(isOn bool) {
}


func (this *NullInputRobot) ToString() string {
        return "DirectXWindowsInputRobot"
}

func (this *NullInputRobot) String() string {
    return this.ToString()
}


func (this *NullInputRobot) WaitForIdle() {
}

