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

type TransformInfoObjectFactory struct {
    LogUtil *LogUtil
}

var TransformInfoObjectFactory_Instance *TransformInfoObjectFactory = NewTransformInfoObjectFactory()

func (this *TransformInfoObjectFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformInfoObjectFactory_GetInstance() *TransformInfoObjectFactory {
        return TransformInfoObjectFactory_Instance
}


func NewTransformInfoObjectFactory() *TransformInfoObjectFactory {
    this := &TransformInfoObjectFactory{}
    this.initFields()
    return this
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformInfoObjectFactory) GetInstance(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) TransformInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Get Instance: "; _ = error_
            if transformInfoInterface != nil {
                error_ = error_ + transformInfoInterface.GetName()
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(error_, this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.PutF("Creating View: " + transformInfoInterface.GetName(), this, commonStrings.GET_INSTANCE)
            }
            var params []any = make([]any, 1); _ = params
            var classes []reflect.Type = make([]reflect.Type, 1); _ = classes
            classes[0] = reflect.TypeOf((*TransformInfoInterface)(nil)).Elem()
            params[0] = any(transformInfoInterface).(any)
            var anyType TransformInterface = any(AbeFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface.GetObjectFile(), classes, params)).(TransformInterface); _ = anyType
            if anyType == nil {
                if transformInfoInterface != nil {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("\nNo Such View Name: ")
                    stringBuffer.Append(transformInfoInterface.GetName())
                    stringBuffer.Append("\nNo Such View Object: ")
                    stringBuffer.Append(transformInfoInterface.GetObjectFile())
                    panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
                } else {
                    panic(errors.New(fmt.Sprint("No Such View Object since transformInfo is null")))
                }
            }
            tryControl1, tryResult1 = 1, anyType; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

