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

type AbstractInputRobot struct {
    LogUtil *LogUtil
    HelpSet *HelpSet
}

// The Java class AbstractInputRobot has subclasses so its variables are this interface
type IAbstractInputRobot interface {
    AbstractInputRobot_() *AbstractInputRobot
    GetHelpSet() *HelpSet
    MouseMove(args ...any)
    MouseMoveToTarget(args ...any)
    SetHelpSet(helpSet *HelpSet)
}

func (this *AbstractInputRobot) AbstractInputRobot_() *AbstractInputRobot {
    return this
}


func (this *AbstractInputRobot) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewAbstractInputRobot(helpSet *HelpSet) *AbstractInputRobot {
    this := &AbstractInputRobot{}
    this.initFields()
        this.SetHelpSet(helpSet)
    return this
}


func (this *AbstractInputRobot) GetHelpSet() *HelpSet {
        return this.HelpSet
}


func (this *AbstractInputRobot) MouseMove_1(point *Point) {
        this.MouseMove(point.X, point.Y)
        this.LogUtil.PutF("Moved Mouse To: x: " + fmt.Sprint(point.X) + " y: " + fmt.Sprint(point.Y), this, "moveMouse")
}

// Java overloads of AbstractInputRobot.mouseMove
func (this *AbstractInputRobot) MouseMove(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Point](args[0], pass) {
            this.MouseMove_1(Java_Arg[*Point](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            this.MouseMove_2(Java_Arg[int](args[0]), Java_Arg[int](args[1]))
            return
        }
    }
    panic("No Java overload of AbstractInputRobot.mouseMove for the arguments")
}


// throws Exception
func (this *AbstractInputRobot) MouseMoveToTarget_1(rectangle *Rectangle, x int, y int) {
        var point *Point = PointHelper_GetCenterPoint(rectangle); _ = point
        this.MouseMove(point.X + x, point.Y + y)
        this.LogUtil.PutF("Moved Mouse To: x: " + fmt.Sprint(point.X) + " y: " + fmt.Sprint(point.Y) + " in the middle of: " + fmt.Sprint(rectangle), this, "moveMouseToTarget")
}

// Java overloads of AbstractInputRobot.mouseMoveToTarget
func (this *AbstractInputRobot) MouseMoveToTarget(args ...any) {
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
    panic("No Java overload of AbstractInputRobot.mouseMoveToTarget for the arguments")
}


// throws Exception
func (this *AbstractInputRobot) MouseMoveToTarget_2(rectangle *Rectangle) {
        this.MouseMoveToTarget(rectangle, 0, 0)
}


func (this *AbstractInputRobot) MouseMove_2(x int, y int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *AbstractInputRobot) SetHelpSet(helpSet *HelpSet) {
        this.HelpSet = helpSet
}

