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

type OSGIActivatorUtil struct {
}



func NewOSGIActivatorUtil() *OSGIActivatorUtil {
    this := &OSGIActivatorUtil{}
    return this
}


// throws Exception
func OSGIActivatorUtil_RegisterAsService_1(bundleContext BundleContext, anyType any, serviceName string) {
        OSGIActivatorUtil_RegisterAsService(bundleContext, anyType, serviceName, StdUtil_GetInstance().CreateHashtable())
}

// Java overloads of OSGIActivatorUtil.registerAsService
func OSGIActivatorUtil_RegisterAsService(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[BundleContext](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) {
            OSGIActivatorUtil_RegisterAsService_1(Java_Arg[BundleContext](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]))
            return
        }
        if len(args) == 4 && Java_IsArg[BundleContext](args[0], pass) && Java_IsArg[any](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[*ABHashtable[any, any]](args[3], pass) {
            OSGIActivatorUtil_RegisterAsService_2(Java_Arg[BundleContext](args[0]), Java_Arg[any](args[1]), Java_Arg[string](args[2]), Java_Arg[*ABHashtable[any, any]](args[3]))
            return
        }
    }
    panic("No Java overload of OSGIActivatorUtil.registerAsService for the arguments")
}


// throws Exception
func OSGIActivatorUtil_RegisterAsService_2(bundleContext BundleContext, anyType any, serviceName string, properties *ABHashtable[any, any]) {
        var serviceRegistration ServiceRegistration = bundleContext.RegisterService(serviceName, anyType, properties); _ = serviceRegistration
        var serviceReference ServiceReference = serviceRegistration.GetReference(); _ = serviceReference
        serviceReference = bundleContext.GetServiceReference(serviceName)
        if serviceReference == NullServiceReferenceFactory_GetInstance().NULL_SERVICE_REFERENCE {
            panic(errors.New(fmt.Sprint("No Such Service Reference")))
        }
}

