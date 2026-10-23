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

type StoreTagWorkFlowFactory struct {
    LogUtil *LogUtil
}

var StoreTagWorkFlowFactory_Instance *StoreTagWorkFlowFactory = NewStoreTagWorkFlowFactory()

func (this *StoreTagWorkFlowFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func StoreTagWorkFlowFactory_GetInstance() *StoreTagWorkFlowFactory {
        return StoreTagWorkFlowFactory_Instance
}


func NewStoreTagWorkFlowFactory() *StoreTagWorkFlowFactory {
    this := &StoreTagWorkFlowFactory{}
    this.initFields()
    return this
}


// throws Exception, LicensingException
func (this *StoreTagWorkFlowFactory) GetInstance(abeClientInformation AbeClientInformationInterface, propertiesHashMap *HashMap, pageContext *PageContext) StoreWorkFlowInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 StoreWorkFlowInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var workFlowClassName string = any(propertiesHashMap.Get(DynamicObjectData_NAME)).(string); _ = workFlowClassName
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Creating WorkFlow: ")
                stringBuffer.Append(workFlowClassName)
                stringBuffer.Append("->")
                stringBuffer.Append(reflect.TypeOf(StoreTagWorkFlowFactory_Instance).String())
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var workFlowClassName string = any(propertiesHashMap.Get(DynamicObjectData_NAME)).(string); _ = workFlowClassName
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Creating WorkFlow: ")
                stringBuffer.Append(workFlowClassName)
                stringBuffer.Append("->")
                stringBuffer.Append(reflect.TypeOf(StoreTagWorkFlowFactory_Instance).String())
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.GET_INSTANCE)
            }
            var params []any = make([]any, 2); _ = params
            var classes []reflect.Type = make([]reflect.Type, 2); _ = classes
            classes[0] = reflect.TypeOf(propertiesHashMap)
            classes[1] = AbeFactory_GetInstance().GetClass(abeClientInformation, "javax.servlet.jsp.PageContext")
            params[0] = any(propertiesHashMap).(any)
            params[1] = any(pageContext).(any)
            tryControl1, tryResult1 = 1, any(AbeFactory_GetInstance().GetInstance(abeClientInformation, workFlowClassName, classes, params)).(StoreWorkFlowInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

