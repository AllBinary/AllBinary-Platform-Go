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

type InputAutomationRobotOSGIServiceVisitor struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements OSGIServiceVisitorInterface


func (this *InputAutomationRobotOSGIServiceVisitor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewInputAutomationRobotOSGIServiceVisitor() *InputAutomationRobotOSGIServiceVisitor {
    this := &InputAutomationRobotOSGIServiceVisitor{}
    this.initFields()
    return this
}


func (this *InputAutomationRobotOSGIServiceVisitor) Visit_1(anyType any) any {
        return this.Visit(any(anyType).(OSGIServiceInterface)).(any)
}

// Java overloads of InputAutomationRobotOSGIServiceVisitor.visit
func (this *InputAutomationRobotOSGIServiceVisitor) Visit(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[any](args[0], pass) {
            return this.Visit_1(Java_Arg[any](args[0]))
        }
        if len(args) == 1 && Java_IsArg[OSGIServiceInterface](args[0], pass) {
            return this.Visit_2(Java_Arg[OSGIServiceInterface](args[0]))
        }
    }
    panic("No Java overload of InputAutomationRobotOSGIServiceVisitor.visit for the arguments")
}


func (this *InputAutomationRobotOSGIServiceVisitor) Visit_2(osgiServiceInterface OSGIServiceInterface) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "visit", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "visit")
            var inputAutomationRobotServiceInterface InputAutomationRobotServiceInterface = any(osgiServiceInterface).(InputAutomationRobotServiceInterface); _ = inputAutomationRobotServiceInterface
            var inputRobotInterfaceArray []InputRobotInterface = inputAutomationRobotServiceInterface.GetInputRobotInterfaceArray(); _ = inputRobotInterfaceArray
            for index := int(0); index < len(inputRobotInterfaceArray); index++ {
                this.LogUtil.PutF("Adding: " + inputRobotInterfaceArray[index].GetName(), this, "visit")
                InputRobotFactory_GetInstance().Add(inputRobotInterfaceArray[index])
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

