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

type InputAutomationBundleActivator struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements BundleActivator, InputAutomationBundleActivatorListenerInterface

var InputAutomationBundleActivator_ModuleManagementThread IThread
var InputAutomationBundleActivator_BundleContext BundleContext

func (this *InputAutomationBundleActivator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewInputAutomationBundleActivator() *InputAutomationBundleActivator {
    this := &InputAutomationBundleActivator{}
    this.initFields()
    return this
}


func InputAutomationBundleActivator_GetBundleContext() BundleContext {
        return InputAutomationBundleActivator_BundleContext
}


// throws Exception
func (this *InputAutomationBundleActivator) Start(bundleContext BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.START, e)
            panic(e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.START)
            InputAutomationBundleActivator_BundleContext = bundleContext
            InputAutomationBundleActivator_ModuleManagementThread = NewThread(NewInputAutomationNewBundleRunnable(this))
            InputAutomationBundleActivator_ModuleManagementThread.Start()
            InputAutomationJFrame_Create(this)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputAutomationBundleActivator) RegisterAsService() {
        OSGIActivatorUtil_RegisterAsService(this.GetBundleContext(), InputAutomationServiceFactory_GetInstance(), reflect.TypeOf((*InputAutomationConfigurationModuleChangeListener)(nil)).Elem().GetName())
        OSGIActivatorUtil_RegisterAsService(this.GetBundleContext(), InputAutomationServiceFactory_GetInstance(), reflect.TypeOf((*InputAutomationRobotChangeListener)(nil)).Elem().GetName())
}


// throws Exception
func (this *InputAutomationBundleActivator) UseServices() {
        NewInputAutomationRobotServiceConsumer(this.GetBundleContext()).Process()
        NewInputAutomationModuleServiceConsumer(this.GetBundleContext()).Process()
}


// throws Exception
func (this *InputAutomationBundleActivator) Stop(context BundleContext) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "stop")
        if InputAutomationJFrame_GetInstance() != nil {
            InputAutomationJFrame_Destroy()
        } else {
            this.LogUtil.PutF("Nothing to stop", this, "stop")
        }
}

