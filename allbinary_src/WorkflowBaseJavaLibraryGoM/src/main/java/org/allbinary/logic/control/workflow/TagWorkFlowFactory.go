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

type TagWorkFlowFactory struct {
    LogUtil *LogUtil
}

var TagWorkFlowFactory_Instance *TagWorkFlowFactory = NewTagWorkFlowFactory()

func (this *TagWorkFlowFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TagWorkFlowFactory_GetInstance() *TagWorkFlowFactory {
        return TagWorkFlowFactory_Instance
}


func NewTagWorkFlowFactory() *TagWorkFlowFactory {
    this := &TagWorkFlowFactory{}
    this.initFields()
    return this
}


// throws Exception, LicensingException
// Java overloaded method: Go needs a different name
func (this *TagWorkFlowFactory) GetInstance(abeClientInformation AbeClientInformationInterface, propertiesHashMap *HashMap, pageContext *PageContext) WorkFlowInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 WorkFlowInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var workFlowClassName string = any(propertiesHashMap.Get(DynamicObjectData_NAME)).(string); _ = workFlowClassName
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Failed To Get Instance: ")
                stringBuffer.Append(workFlowClassName)
                stringBuffer.Append("->")
                stringBuffer.Append(reflect.TypeOf(TagWorkFlowFactory_Instance).String())
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var workFlowClassName string = any(propertiesHashMap.Get(DynamicObjectData_NAME)).(string); _ = workFlowClassName
            var params []any = make([]any, 2); _ = params
            var classes []reflect.Type = make([]reflect.Type, 2); _ = classes
            classes[0] = reflect.TypeOf(propertiesHashMap)
            classes[1] = AbeFactory_GetInstance().GetClass(abeClientInformation, "javax.servlet.jsp.PageContext")
            params[0] = any(propertiesHashMap).(any)
            params[1] = any(pageContext).(any)
            tryControl1, tryResult1 = 1, any(AbeFactory_GetInstance().GetInstance(abeClientInformation, workFlowClassName, classes, params)).(WorkFlowInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

