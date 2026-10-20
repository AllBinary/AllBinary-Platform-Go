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

type OSGIServiceUtil struct {
    LogUtil *LogUtil
}

var OSGIServiceUtil_Instance *OSGIServiceUtil = NewOSGIServiceUtil()

func (this *OSGIServiceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func OSGIServiceUtil_GetInstance() *OSGIServiceUtil {
        return OSGIServiceUtil_Instance
}


func NewOSGIServiceUtil() *OSGIServiceUtil {
    this := &OSGIServiceUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *OSGIServiceUtil) GetServicesObjectVector(bundleContext BundleContext, serviceReferences []ServiceReference) IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        if serviceReferences != nil {
            var size int = len(serviceReferences); _ = size
            this.LogUtil.PutF("Service References: " + fmt.Sprint(size), this, "getServicesObjectVector")
            for index := int(0); index < size; index++ {
                var serviceReference ServiceReference = serviceReferences[index]; _ = serviceReference
                if serviceReference != nil {
                    var osgiServiceInterface OSGIServiceInterface = any(bundleContext.GetService(serviceReference)).(OSGIServiceInterface); _ = osgiServiceInterface
                    if osgiServiceInterface == nil {
                        panic(errors.New(fmt.Sprint("No Service For Reference")))
                    }
                    vector.Add(osgiServiceInterface)
                }
            }
        }
        return vector
}

