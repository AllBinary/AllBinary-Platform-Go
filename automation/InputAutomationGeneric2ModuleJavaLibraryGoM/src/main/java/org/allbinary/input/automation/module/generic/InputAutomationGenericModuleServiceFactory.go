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

type InputAutomationGenericModuleServiceFactory struct {
    InputAutomationModuleServiceFactory
    LogUtil *LogUtil
}

var InputAutomationGenericModuleServiceFactory_InputAutomationGenericModuleServiceFactory *InputAutomationGenericModuleServiceFactory = NewInputAutomationGenericModuleServiceFactory()

func (this *InputAutomationGenericModuleServiceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewInputAutomationGenericModuleServiceFactory() *InputAutomationGenericModuleServiceFactory {
    this := &InputAutomationGenericModuleServiceFactory{}
    this.InputAutomationModuleServiceFactory = *NewInputAutomationModuleServiceFactory()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            var inputAutomationModuleInterfaceArray []IInputAutomationModuleFactoryInterface = make([]IInputAutomationModuleFactoryInterface, 1); _ = inputAutomationModuleInterfaceArray
            inputAutomationModuleInterfaceArray[0] = NewGenericModuleFactory()
            this.SetInputAutomationModuleInterfaceArray(inputAutomationModuleInterfaceArray)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func InputAutomationGenericModuleServiceFactory_GetInstance() *InputAutomationGenericModuleServiceFactory {
        return InputAutomationGenericModuleServiceFactory_InputAutomationGenericModuleServiceFactory
}

