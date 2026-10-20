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

type RuntimeEnvironmentVariables struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}


func (this *RuntimeEnvironmentVariables) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewRuntimeEnvironmentVariables() *RuntimeEnvironmentVariables {
    this := &RuntimeEnvironmentVariables{}
    this.initFields()
    return this
}


func (this *RuntimeEnvironmentVariables) GetTest1() string {
        return System_GetProperty("java.java.home")
}


func (this *RuntimeEnvironmentVariables) GetVariable(var_ string) string {
        var environmentVariables IProperties = this.Get(); _ = environmentVariables
        return environmentVariables.GetProperty(var_)
}


func (this *RuntimeEnvironmentVariables) Get() IProperties {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IProperties) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "get()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var process *Process; _ = process
            var environmentVariables IProperties = NewProperties(); _ = environmentVariables
            var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
            var osName string = operatingSystemInterface.GetName(); _ = osName
            var linuxCommand string = "env"; _ = linuxCommand
            var windowsCommand string = "cmd.exe /c set"; _ = windowsCommand
            var solarisCommand string = linuxCommand; _ = solarisCommand
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            if strings.Compare(osName, operatingSystems.LINUX) == 0 {
                process = Runtime_GetRuntime().Exec(linuxCommand)
            } else if strings.Index(osName, string(operatingSystems.WINDOWS)) >= 0 {
                process = Runtime_GetRuntime().Exec(windowsCommand)
            } else if strings.Compare(osName, operatingSystems.SOLARIS) == 0 {
                process = Runtime_GetRuntime().Exec(solarisCommand)
            } else {
                panic(errors.New(fmt.Sprint()))
            }
            var processInputStreamReader *InputStreamReader = NewInputStreamReader(process.GetInputStream()); _ = processInputStreamReader
            var bufferedReader *BufferedReader = NewBufferedReader(processInputStreamReader); _ = bufferedReader
            var keyValuePair string; _ = keyValuePair
            for (func() any { keyValuePair = bufferedReader.ReadLine(); return keyValuePair }()) != nil {
                var index int = strings.Index(keyValuePair, string('=')); _ = index
                var key string = keyValuePair[0:index]; _ = key
                var value string = keyValuePair[index + 1:]; _ = value
                environmentVariables.Put(key, value)
            }
            tryControl1, tryResult1 = 1, environmentVariables; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) GetTempDir() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getTempDir()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, System_GetProperty("java.io.tmpdir"); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) GetLibraryPath() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getLibraryPath()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, System_GetProperty("java.library.path"); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) GetClassPath() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getClassPath()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, System_GetProperty("java.class.path"); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) UpdateLibraryPath(libPath string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "updateLibraryPath()", e)
            }
                }
            }()
            var existingPath string = this.GetLibraryPath(); _ = existingPath
            System_SetProperty("java.library.path", libPath + ":" + existingPath)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *RuntimeEnvironmentVariables) GetUserHome() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getUserHome()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, System_GetProperty("user.home"); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) GetUserDir() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getUserDir()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, System_GetProperty("user.dir"); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RuntimeEnvironmentVariables) Set() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "get()", e)
            }
                }
            }()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *RuntimeEnvironmentVariables) SetVariable(var_ string, val string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().OSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setVariable()", e)
            }
                }
            }()
            var process *Process; _ = process
            var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            var osName string = operatingSystemInterface.GetName(); _ = osName
            var windowsCommand string = "cmd.exe /c set" + var_ + "=" + val; _ = windowsCommand
            if strings.Compare(osName, operatingSystems.LINUX) == 0 {
                panic(errors.New(fmt.Sprint()))
            } else if strings.Index(osName, string(operatingSystems.WINDOWS)) >= 0 {
                process = Runtime_GetRuntime().Exec(windowsCommand)
            } else if strings.Compare(osName, operatingSystems.SOLARIS) == 0 {
                panic(errors.New(fmt.Sprint()))
            } else {
                panic(errors.New(fmt.Sprint()))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

