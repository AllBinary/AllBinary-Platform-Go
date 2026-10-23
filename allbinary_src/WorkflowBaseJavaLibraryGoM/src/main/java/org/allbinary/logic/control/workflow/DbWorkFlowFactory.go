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

type DbWorkFlowFactory struct {
    LogUtil *LogUtil
}

var DbWorkFlowFactory_Instance *DbWorkFlowFactory = NewDbWorkFlowFactory()

func (this *DbWorkFlowFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func DbWorkFlowFactory_GetInstance() *DbWorkFlowFactory {
        return DbWorkFlowFactory_Instance
}


func NewDbWorkFlowFactory() *DbWorkFlowFactory {
    this := &DbWorkFlowFactory{}
    this.initFields()
    return this
}


// throws Exception, LicensingException
func (this *DbWorkFlowFactory) GetInstance(abeClientInformation AbeClientInformationInterface, hashMap *HashMap) WorkFlowInterface {
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
            var params []any = make([]any, 1); _ = params
            var classes []reflect.Type = make([]reflect.Type, 1); _ = classes
            classes[0] = reflect.TypeOf(hashMap)
            params[0] = any(hashMap).(any)
            var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, className, classes, params); _ = anyType
            tryControl1, tryResult1 = 1, any(NewWorkFlowWrapper(anyType)).(WorkFlowInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

