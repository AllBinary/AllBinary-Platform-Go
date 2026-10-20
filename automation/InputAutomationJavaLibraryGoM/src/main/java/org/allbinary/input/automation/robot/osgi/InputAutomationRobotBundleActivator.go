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

type InputAutomationRobotBundleActivator struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputRobotInterface []InputRobotInterface
    ServiceFactory ServiceFactory
}

// Implements BundleActivator

// The Java class InputAutomationRobotBundleActivator has subclasses so its variables are this interface
type IInputAutomationRobotBundleActivator interface {
    InputAutomationRobotBundleActivator_() *InputAutomationRobotBundleActivator
    Init()
    RegisterAsService(bundleContext BundleContext)
    GetInputAutomationRobotChangeListener(context BundleContext) InputAutomationRobotChangeListener
    AddRobots(context BundleContext)
    RemoveRobots(context BundleContext)
    Start(context BundleContext)
    Stop(context BundleContext)
    GetServiceFactory() any
    SetServiceFactory(serviceObject ServiceFactory)
    GetInputRobotInterface() []InputRobotInterface
    SetInputRobotInterface(inputRobotInterface []InputRobotInterface)
}

func (this *InputAutomationRobotBundleActivator) InputAutomationRobotBundleActivator_() *InputAutomationRobotBundleActivator {
    return this
}


func (this *InputAutomationRobotBundleActivator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewInputAutomationRobotBundleActivator() *InputAutomationRobotBundleActivator {
    this := &InputAutomationRobotBundleActivator{}
    this.initFields()
        this.Init()
    return this
}


// throws Exception
func (this *InputAutomationRobotBundleActivator) Init() {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *InputAutomationRobotBundleActivator) RegisterAsService(bundleContext BundleContext) {
        OSGIActivatorUtil_RegisterAsService(bundleContext, this.GetServiceFactory(), reflect.TypeOf((*InputAutomationRobotServiceInterface)(nil)).Elem().GetName())
}


// throws Exception
func (this *InputAutomationRobotBundleActivator) GetInputAutomationRobotChangeListener(context BundleContext) InputAutomationRobotChangeListener {
        var serviceReference ServiceReference = context.GetServiceReference(reflect.TypeOf((*InputAutomationRobotChangeListener)(nil)).Elem().GetName()); _ = serviceReference
        if serviceReference != nil {
            var inputAutomationRobotChangeListener InputAutomationRobotChangeListener = any(context.GetService(serviceReference)).(InputAutomationRobotChangeListener); _ = inputAutomationRobotChangeListener
            if inputAutomationRobotChangeListener == nil {
                panic(errors.New(fmt.Sprint("No Service For ServiceReference")))
            }
            return inputAutomationRobotChangeListener
        } else {
            this.LogUtil.PutF("No ServiceReference: " + reflect.TypeOf((*InputAutomationRobotChangeListener)(nil)).Elem().GetName(), this, "getInputAutomationRobotChangeListener")
            return nil
        }
}


func (this *InputAutomationRobotBundleActivator) AddRobots(context BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, "addModules")
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "addRobots")
            var inputAutomationRobotChangeListener InputAutomationRobotChangeListener = this.GetInputAutomationRobotChangeListener(context); _ = inputAutomationRobotChangeListener
            if inputAutomationRobotChangeListener != nil {
                for index := int(0); index < len(this.GetInputRobotInterface()); index++ {
                    var inputAutomationRobotChangeEvent *InputAutomationRobotChangeEvent = InputAutomationRobotUtil_GetChangeEvent(this.GetInputRobotInterface()[index]); _ = inputAutomationRobotChangeEvent
                    inputAutomationRobotChangeListener.OnAdd(inputAutomationRobotChangeEvent)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InputAutomationRobotBundleActivator) RemoveRobots(context BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, "removeRobots")
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "removeRobots")
            var inputAutomationRobotChangeListener InputAutomationRobotChangeListener = this.GetInputAutomationRobotChangeListener(context); _ = inputAutomationRobotChangeListener
            if inputAutomationRobotChangeListener != nil {
                for index := int(0); index < len(this.GetInputRobotInterface()); index++ {
                    var inputAutomationRobotChangeEvent *InputAutomationRobotChangeEvent = InputAutomationRobotUtil_GetChangeEvent(this.GetInputRobotInterface()[index]); _ = inputAutomationRobotChangeEvent
                    inputAutomationRobotChangeListener.OnRemove(inputAutomationRobotChangeEvent)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputAutomationRobotBundleActivator) Start(context BundleContext) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.START)
        this.AddRobots(context)
        this.RegisterAsService(context)
}


// throws Exception
func (this *InputAutomationRobotBundleActivator) Stop(context BundleContext) {
        this.LogUtil.PutF("Stop", this, this.CommonStrings.START)
        this.RemoveRobots(context)
}


func (this *InputAutomationRobotBundleActivator) GetServiceFactory() any {
        return this.ServiceFactory
}


func (this *InputAutomationRobotBundleActivator) SetServiceFactory(serviceObject ServiceFactory) {
        this.ServiceFactory = serviceObject
}


func (this *InputAutomationRobotBundleActivator) GetInputRobotInterface() []InputRobotInterface {
        return this.InputRobotInterface
}


func (this *InputAutomationRobotBundleActivator) SetInputRobotInterface(inputRobotInterface []InputRobotInterface) {
        this.InputRobotInterface = inputRobotInterface
}

