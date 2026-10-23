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

type NewWorkFlowFactory struct {
    LogUtil *LogUtil
}

var NewWorkFlowFactory_Instance *NewWorkFlowFactory = NewNewWorkFlowFactory()

func (this *NewWorkFlowFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewWorkFlowFactory_GetInstance() *NewWorkFlowFactory {
        return NewWorkFlowFactory_Instance
}


func NewNewWorkFlowFactory() *NewWorkFlowFactory {
    this := &NewWorkFlowFactory{}
    this.initFields()
    return this
}


// throws Exception, LicensingException
func (this *NewWorkFlowFactory) GetInstance_1(abeClientInformation AbeClientInformationInterface, hashMap *HashMap, pageContext *PageContext) WorkFlowInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 WorkFlowInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "getInstance(HashMap)", e)
            }
            panic(e)
                }
            }()
            var className string = any(hashMap.Get(DynamicObjectData_NAME)).(string); _ = className
            var params []any = make([]any, 2); _ = params
            var classes []reflect.Type = make([]reflect.Type, 2); _ = classes
            classes[0] = reflect.TypeOf(hashMap)
            classes[1] = AbeFactory_GetInstance().GetClass(abeClientInformation, "javax.servlet.jsp.PageContext")
            params[0] = any(hashMap).(any)
            params[1] = any(pageContext).(any)
            var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, className, classes, params); _ = anyType
            tryControl1, tryResult1 = 1, any(NewWorkFlowWrapper(anyType)).(WorkFlowInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of NewWorkFlowFactory.getInstance
func (this *NewWorkFlowFactory) GetInstance(args ...any) WorkFlowInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) && Java_IsArg[*PageContext](args[2], pass) {
            return this.GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*HashMap](args[1]), Java_Arg[*PageContext](args[2]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of NewWorkFlowFactory.getInstance for the arguments")
}


// throws Exception, LicensingException
func (this *NewWorkFlowFactory) GetInstance_2(abeClientInformation AbeClientInformationInterface, className string) WorkFlowInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 WorkFlowInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            if className != "" && strings.Compare(className, stringUtil.EMPTY_STRING) != 0 {
                var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, className, nil, nil); _ = anyType
                tryControl1, tryResult1 = 1, any(NewWorkFlowWrapper(anyType)).(WorkFlowInterface); return
            } else {
                panic(errors.New(fmt.Sprint("No WorkFlow ClassName")))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

