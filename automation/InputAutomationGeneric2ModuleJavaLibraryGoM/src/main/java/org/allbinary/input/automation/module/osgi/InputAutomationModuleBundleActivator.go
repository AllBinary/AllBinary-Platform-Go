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

type InputAutomationModuleBundleActivator struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputAutomationModuleInterface []IInputAutomationModuleFactoryInterface
    ServiceFactory ServiceFactory
}

// Implements BundleActivator

// The Java class InputAutomationModuleBundleActivator has subclasses so its variables are this interface
type IInputAutomationModuleBundleActivator interface {
    InputAutomationModuleBundleActivator_() *InputAutomationModuleBundleActivator
    Init()
    GetInputAutomationConfigurationModuleChangeListener(context BundleContext) InputAutomationConfigurationModuleChangeListener
    AddModules(context BundleContext)
    RegisterAsService(bundleContext BundleContext)
    RemoveModules(context BundleContext)
    Start(context BundleContext)
    Stop(context BundleContext)
    GetInputAutomationModuleInterface() []IInputAutomationModuleFactoryInterface
    SetInputAutomationModuleInterface(inputAutomationModuleInterface []IInputAutomationModuleFactoryInterface)
    GetServiceFactory() any
    SetServiceFactory(serviceObject ServiceFactory)
}

func (this *InputAutomationModuleBundleActivator) InputAutomationModuleBundleActivator_() *InputAutomationModuleBundleActivator {
    return this
}


func (this *InputAutomationModuleBundleActivator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewInputAutomationModuleBundleActivator() *InputAutomationModuleBundleActivator {
    this := &InputAutomationModuleBundleActivator{}
    this.initFields()
        this.Init()
    return this
}


// throws Exception
func (this *InputAutomationModuleBundleActivator) Init() {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *InputAutomationModuleBundleActivator) GetInputAutomationConfigurationModuleChangeListener(context BundleContext) InputAutomationConfigurationModuleChangeListener {
        var serviceReference ServiceReference = context.GetServiceReference(reflect.TypeOf((*InputAutomationConfigurationModuleChangeListener)(nil)).Elem().GetName()); _ = serviceReference
        if serviceReference != nil {
            var inputAutomationConfigurationModuleChangeListener InputAutomationConfigurationModuleChangeListener = any(context.GetService(serviceReference)).(InputAutomationConfigurationModuleChangeListener); _ = inputAutomationConfigurationModuleChangeListener
            if inputAutomationConfigurationModuleChangeListener == nil {
                panic(errors.New(fmt.Sprint("No Service For ServiceReference")))
            }
            return inputAutomationConfigurationModuleChangeListener
        } else {
            this.LogUtil.PutF("No ServiceReference: " + reflect.TypeOf((*InputAutomationConfigurationModuleChangeListener)(nil)).Elem().GetName(), this, "addModules")
            return nil
        }
}


func (this *InputAutomationModuleBundleActivator) AddModules(context BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, "addModules")
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "addModules")
            var inputAutomationConfigurationModuleChangeListener InputAutomationConfigurationModuleChangeListener = this.GetInputAutomationConfigurationModuleChangeListener(context); _ = inputAutomationConfigurationModuleChangeListener
            if inputAutomationConfigurationModuleChangeListener != nil {
                for index := int(0); index < len(this.GetInputAutomationModuleInterface()); index++ {
                    var inputAutomationConfigurationModuleChangeEvent *InputAutomationConfigurationModuleChangeEvent = InputAutomationConfigurationUtil_GetChangeEvent(this.GetInputAutomationModuleInterface()[index]); _ = inputAutomationConfigurationModuleChangeEvent
                    inputAutomationConfigurationModuleChangeListener.OnAdd(inputAutomationConfigurationModuleChangeEvent)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputAutomationModuleBundleActivator) RegisterAsService(bundleContext BundleContext) {
        OSGIActivatorUtil_RegisterAsService(bundleContext, this.GetServiceFactory(), reflect.TypeOf((*InputAutomationModuleServiceInterface)(nil)).Elem().GetName())
}


func (this *InputAutomationModuleBundleActivator) RemoveModules(context BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, "removeModules")
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "removeModules")
            var inputAutomationConfigurationModuleChangeListener InputAutomationConfigurationModuleChangeListener = this.GetInputAutomationConfigurationModuleChangeListener(context); _ = inputAutomationConfigurationModuleChangeListener
            if inputAutomationConfigurationModuleChangeListener != nil {
                for index := int(0); index < len(this.GetInputAutomationModuleInterface()); index++ {
                    var inputAutomationConfigurationModuleChangeEvent *InputAutomationConfigurationModuleChangeEvent = InputAutomationConfigurationUtil_GetChangeEvent(this.GetInputAutomationModuleInterface()[index]); _ = inputAutomationConfigurationModuleChangeEvent
                    inputAutomationConfigurationModuleChangeListener.OnRemove(inputAutomationConfigurationModuleChangeEvent)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputAutomationModuleBundleActivator) Start(context BundleContext) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.START)
        this.AddModules(context)
        this.RegisterAsService(context)
}


// throws Exception
func (this *InputAutomationModuleBundleActivator) Stop(context BundleContext) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "stop")
        this.RemoveModules(context)
}


func (this *InputAutomationModuleBundleActivator) GetInputAutomationModuleInterface() []IInputAutomationModuleFactoryInterface {
        return this.InputAutomationModuleInterface
}


func (this *InputAutomationModuleBundleActivator) SetInputAutomationModuleInterface(inputAutomationModuleInterface []IInputAutomationModuleFactoryInterface) {
        this.InputAutomationModuleInterface = inputAutomationModuleInterface
}


func (this *InputAutomationModuleBundleActivator) GetServiceFactory() any {
        return this.ServiceFactory
}


func (this *InputAutomationModuleBundleActivator) SetServiceFactory(serviceObject ServiceFactory) {
        this.ServiceFactory = serviceObject
}

