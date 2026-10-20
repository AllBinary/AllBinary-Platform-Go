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

type InputAutomationModuleServiceFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputAutomationModuleInterfaceArray []IInputAutomationModuleFactoryInterface
}

// Implements ServiceFactory

// The Java class InputAutomationModuleServiceFactory has subclasses so its variables are this interface
type IInputAutomationModuleServiceFactory interface {
    InputAutomationModuleServiceFactory_() *InputAutomationModuleServiceFactory
    GetService(bundle Bundle, registration ServiceRegistration) any
    UngetService(bundle Bundle, registration ServiceRegistration, service any)
    GetInputAutomationModuleInterfaceArray() []IInputAutomationModuleFactoryInterface
    SetInputAutomationModuleInterfaceArray(inputAutomationModuleInterfaceArray []IInputAutomationModuleFactoryInterface)
}

func (this *InputAutomationModuleServiceFactory) InputAutomationModuleServiceFactory_() *InputAutomationModuleServiceFactory {
    return this
}


func (this *InputAutomationModuleServiceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewInputAutomationModuleServiceFactory() *InputAutomationModuleServiceFactory {
    this := &InputAutomationModuleServiceFactory{}
    this.initFields()
    return this
}


func (this *InputAutomationModuleServiceFactory) GetService(bundle Bundle, registration ServiceRegistration) any {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getService")
        return NewInputAutomationModuleService(this.GetInputAutomationModuleInterfaceArray())
}


func (this *InputAutomationModuleServiceFactory) UngetService(bundle Bundle, registration ServiceRegistration, service any) {
}


func (this *InputAutomationModuleServiceFactory) GetInputAutomationModuleInterfaceArray() []IInputAutomationModuleFactoryInterface {
        return this.InputAutomationModuleInterfaceArray
}


func (this *InputAutomationModuleServiceFactory) SetInputAutomationModuleInterfaceArray(inputAutomationModuleInterfaceArray []IInputAutomationModuleFactoryInterface) {
        this.InputAutomationModuleInterfaceArray = inputAutomationModuleInterfaceArray
}

