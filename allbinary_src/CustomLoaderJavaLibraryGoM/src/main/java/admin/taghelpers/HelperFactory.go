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

type HelperFactory struct {
    LogUtil *LogUtil
}


func (this *HelperFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewHelperFactory() *HelperFactory {
    this := &HelperFactory{}
    this.initFields()
    return this
}


// throws LicensingException
func HelperFactory_GetInstance_1(abeClientInformation AbeClientInformationInterface, factoryName string, className string, hashMap *HashMap, httpServletRequest *HttpServletRequest) any {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Get Instance Args: HashMap=" + fmt.Sprint(hashMap) + " HttpServletRequest=" + fmt.Sprint(httpServletRequest); _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                logUtil.Put(error_, factoryName + "->HelperFactory", "getInstance(String, String, HashMap, PageContext)", e)
            }
            panic(e)
                }
            }()
            var params []any = make([]any, 2); _ = params
            var classes []reflect.Type = make([]reflect.Type, 2); _ = classes
            var abeFactory *AbeFactory = AbeFactory_GetInstance(); _ = abeFactory
            classes[0] = reflect.TypeOf(hashMap)
            classes[1] = abeFactory.GetClass(abeClientInformation, "javax.servlet.http.HttpServletRequest")
            params[0] = any(hashMap).(any)
            params[1] = any(httpServletRequest).(any)
            var anyType any = abeFactory.GetInstance(abeClientInformation, className, classes, params); _ = anyType
            tryControl1, tryResult1 = 1, anyType; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of HelperFactory.getInstance
func HelperFactory_GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 5 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[*HashMap](args[3], pass) && Java_IsArg[*HttpServletRequest](args[4], pass) {
            return HelperFactory_GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[*HashMap](args[3]), Java_Arg[*HttpServletRequest](args[4]))
        }
        if len(args) == 5 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[*HashMap](args[3], pass) && Java_IsArg[*PageContext](args[4], pass) {
            return HelperFactory_GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[*HashMap](args[3]), Java_Arg[*PageContext](args[4]))
        }
        if len(args) == 6 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[*HashMap](args[3], pass) && Java_IsArg[*HashMap](args[4], pass) && Java_IsArg[*PageContext](args[5], pass) {
            return HelperFactory_GetInstance_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[*HashMap](args[3]), Java_Arg[*HashMap](args[4]), Java_Arg[*PageContext](args[5]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return HelperFactory_GetInstance_4(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of HelperFactory.getInstance for the arguments")
}


// throws LicensingException
func HelperFactory_GetInstance_2(abeClientInformation AbeClientInformationInterface, factoryName string, className string, hashMap *HashMap, pageContext *PageContext) any {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var error_ string = "Failed To Get Instance Args: HashMap=" + fmt.Sprint(hashMap) + " PageContext=" + fmt.Sprint(pageContext); _ = error_
                logUtil.Put(error_, factoryName + "->HelperFactory", "getInstance(String, String, HashMap, PageContext)", e)
            }
            panic(e)
                }
            }()
            var params []any = make([]any, 2); _ = params
            var classes []reflect.Type = make([]reflect.Type, 2); _ = classes
            var abeFactory *AbeFactory = AbeFactory_GetInstance(); _ = abeFactory
            classes[0] = reflect.TypeOf(hashMap)
            classes[1] = abeFactory.GetClass(abeClientInformation, "javax.servlet.jsp.PageContext")
            params[0] = any(hashMap).(any)
            params[1] = any(pageContext).(any)
            var anyType any = abeFactory.GetInstance(abeClientInformation, className, classes, params); _ = anyType
            tryControl1, tryResult1 = 1, anyType; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func HelperFactory_GetInstance_3(abeClientInformation AbeClientInformationInterface, factoryName string, className string, hashMap *HashMap, specialhashMap *HashMap, pageContext *PageContext) any {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var error_ string = "Failed To Get Instance Args: HashMap=" + fmt.Sprint(hashMap) + " PageContext=" + fmt.Sprint(pageContext); _ = error_
                logUtil.Put(error_, factoryName + "->HelperFactory", "getInstance(String, String, HashMap, PageContext)", e)
            }
            panic(e)
                }
            }()
            var params []any = make([]any, 3); _ = params
            var classes []reflect.Type = make([]reflect.Type, 3); _ = classes
            var abeFactory *AbeFactory = AbeFactory_GetInstance(); _ = abeFactory
            classes[0] = reflect.TypeOf(hashMap)
            classes[1] = reflect.TypeOf(specialhashMap)
            classes[2] = abeFactory.GetClass(abeClientInformation, "javax.servlet.jsp.PageContext")
            params[0] = any(hashMap).(any)
            params[1] = any(specialhashMap).(any)
            params[2] = any(pageContext).(any)
            var anyType any = abeFactory.GetInstance(abeClientInformation, className, classes, params); _ = anyType
            tryControl1, tryResult1 = 1, anyType; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func HelperFactory_GetInstance_4(abeClientInformation AbeClientInformationInterface, factoryName string, className string) any {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.Put(commonStrings.EXCEPTION, factoryName + "->HelperFactory", "getInstance(String, String, HashMap, PageContext)", e)
            }
            panic(e)
                }
            }()
            var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, className); _ = anyType
            tryControl1, tryResult1 = 1, anyType; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

