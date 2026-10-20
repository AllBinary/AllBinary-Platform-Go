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

type InputAutomationServiceConsumer struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    OsgiServiceUtil *OSGIServiceUtil
    RegistryName string
    BundleContext BundleContext
    OsgiServiceVisitorInterface OSGIServiceVisitorInterface
}

// The Java class InputAutomationServiceConsumer has subclasses so its variables are this interface
type IInputAutomationServiceConsumer interface {
    InputAutomationServiceConsumer_() *InputAutomationServiceConsumer
    Process()
    GetServiceReferences() []ServiceReference
    GetBundleContext() BundleContext
    SetBundleContext(aBundleContext BundleContext)
    GetRegistryName() string
    SetRegistryName(registryName string)
    GetOsgiServiceVisitorInterface() OSGIServiceVisitorInterface
    SetOsgiServiceVisitorInterface(osgiServiceVisitorInterface OSGIServiceVisitorInterface)
}

func (this *InputAutomationServiceConsumer) InputAutomationServiceConsumer_() *InputAutomationServiceConsumer {
    return this
}


func (this *InputAutomationServiceConsumer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.OsgiServiceUtil = OSGIServiceUtil_GetInstance()
    
    
    
}


func NewInputAutomationServiceConsumer(registryName string, bundleContext BundleContext, osgiServiceVisitorInterface OSGIServiceVisitorInterface) *InputAutomationServiceConsumer {
    this := &InputAutomationServiceConsumer{}
    this.initFields()
        this.SetRegistryName(registryName)
        this.SetBundleContext(bundleContext)
        this.SetOsgiServiceVisitorInterface(osgiServiceVisitorInterface)
    return this
}


// throws Exception
func (this *InputAutomationServiceConsumer) Process() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
        var vector IBasicArrayList = this.OsgiServiceUtil.GetServicesObjectVector(this.GetBundleContext(), this.GetServiceReferences()); _ = vector
        this.LogUtil.PutF("Processing " + fmt.Sprint(vector.Size()) + " Services", this, this.CommonStrings.PROCESS)
        var size int = vector.Size(); _ = size
        var osgiServiceInterface OSGIServiceInterface; _ = osgiServiceInterface
        for index := int(0); index < size; index++ {
            osgiServiceInterface = any(vector.Get(index)).(OSGIServiceInterface)
            if !this.GetOsgiServiceVisitorInterface().Visit(osgiServiceInterface) {
                panic(errors.New(fmt.Sprint("Unable to process service: " + fmt.Sprint(osgiServiceInterface))))
            }
        }
}


// throws Exception
func (this *InputAutomationServiceConsumer) GetServiceReferences() []ServiceReference {
        return this.GetBundleContext().GetServiceReferences(this.GetRegistryName(), nil)
}


func (this *InputAutomationServiceConsumer) GetBundleContext() BundleContext {
        return this.BundleContext
}


func (this *InputAutomationServiceConsumer) SetBundleContext(aBundleContext BundleContext) {
        this.BundleContext = aBundleContext
}


func (this *InputAutomationServiceConsumer) GetRegistryName() string {
        return this.RegistryName
}


func (this *InputAutomationServiceConsumer) SetRegistryName(registryName string) {
        this.RegistryName = registryName
}


func (this *InputAutomationServiceConsumer) GetOsgiServiceVisitorInterface() OSGIServiceVisitorInterface {
        return this.OsgiServiceVisitorInterface
}


func (this *InputAutomationServiceConsumer) SetOsgiServiceVisitorInterface(osgiServiceVisitorInterface OSGIServiceVisitorInterface) {
        this.OsgiServiceVisitorInterface = osgiServiceVisitorInterface
}

