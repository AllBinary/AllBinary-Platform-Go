/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2019 AllBinary 
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

type AndroidServicesUtil struct {
    LogUtil *LogUtil
    ResourceUtil *ResourceUtil
    SERVICE_LIMIT_MAX int
    IS_SERVICE_RUNNING string
    SERVICE_NOT_FOUND_RUNNING string
}

var AndroidServicesUtil_Instance *AndroidServicesUtil = NewAndroidServicesUtil()

func (this *AndroidServicesUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ResourceUtil = ResourceUtil_GetInstance()
    this.SERVICE_LIMIT_MAX = 1000
    this.IS_SERVICE_RUNNING = "isServiceRunning"
    this.SERVICE_NOT_FOUND_RUNNING = "Service not found Running: "
}

func NewAndroidServicesUtil() *AndroidServicesUtil {
    this := &AndroidServicesUtil{}
    this.initFields()
    return this
}


func AndroidServicesUtil_GetInstance() *AndroidServicesUtil {
        return AndroidServicesUtil_Instance
}


func (this *AndroidServicesUtil) IsServiceRunning(name string) bool {
        var activityManager *ActivityManager = any(this.ResourceUtil.GetContext().GetSystemService(Context_ACTIVITY_SERVICE)).(*ActivityManager); _ = activityManager
        var runningServicesList IList[*RunningServiceInfo] = activityManager.GetRunningServices(this.SERVICE_LIMIT_MAX); _ = runningServicesList
        var runningServiceInfo *RunningServiceInfo; _ = runningServiceInfo
        var serviceComponent *ComponentName; _ = serviceComponent
        var serviceName string; _ = serviceName
        var size int = runningServicesList.Size(); _ = size
        for index := int(0); index < size; index++ {
            runningServiceInfo = runningServicesList.Get(index)
            serviceComponent = runningServiceInfo.Service
            serviceName = fmt.Sprint(serviceComponent)
            if strings.Index(serviceName, string(name)) >= 0 {
                return true
            }
        }
        this.LogUtil.PutF(this.SERVICE_NOT_FOUND_RUNNING + name, this, this.IS_SERVICE_RUNNING)
        return false
}

