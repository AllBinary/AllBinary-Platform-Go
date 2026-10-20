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

type AllBinaryPreloaderActivator struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements BundleActivator

var AllBinaryPreloaderActivator_Context BundleContext
var AllBinaryPreloaderActivator_CRYPT_REGISTRY_NAME string = "bundle.abcs.logic.system.loader.CryptService"

func (this *AllBinaryPreloaderActivator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewAllBinaryPreloaderActivator() *AllBinaryPreloaderActivator {
    this := &AllBinaryPreloaderActivator{}
    this.initFields()
    return this
}


func AllBinaryPreloaderActivator_GetBundleContext() BundleContext {
        return AllBinaryPreloaderActivator_Context
}


// throws Exception
func (this *AllBinaryPreloaderActivator) Start(context BundleContext) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "start", e)
            panic(e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "start")
            AllBinaryPreloaderActivator_Context = context
            this.RegisterAsService()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryPreloaderActivator) RegisterAsService() {
        var properties *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = properties
        var serviceRegistration ServiceRegistration = AllBinaryPreloaderActivator_Context.RegisterService(AllBinaryPreloaderActivator_CRYPT_REGISTRY_NAME, NewCryptServiceFactory(), properties); _ = serviceRegistration
        var serviceReference ServiceReference = serviceRegistration.GetReference(); _ = serviceReference
        serviceReference = AllBinaryPreloaderActivator_Context.GetServiceReference(AllBinaryPreloaderActivator_CRYPT_REGISTRY_NAME)
        if serviceReference == nil {
            panic(errors.New(fmt.Sprint("No Such Service Reference")))
        }
}


// throws Exception
func (this *AllBinaryPreloaderActivator) Stop(context BundleContext) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "stop")
}

