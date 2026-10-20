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

type TransformFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var TransformFactory_Instance *TransformFactory = NewTransformFactory()

func (this *TransformFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func TransformFactory_GetInstance() *TransformFactory {
        return TransformFactory_Instance
}


func NewTransformFactory() *TransformFactory {
    this := &TransformFactory{}
    this.initFields()
    return this
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformFactory) GetInstance(abeClientInformation AbeClientInformationInterface, viewName string, ownerTransformInfoInterface TransformInfoInterface) TransformInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Creating Transform: " + viewName, this, this.CommonStrings.GET_INSTANCE)
            }
            var ownerTransformInfoHttpInterface TransformInfoHttpInterface = any(ownerTransformInfoInterface).(TransformInfoHttpInterface); _ = ownerTransformInfoHttpInterface
            var transformInfoEntity *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntity
            var transformInfoInterface TransformInfoInterface = transformInfoEntity.Get(viewName, ownerTransformInfoHttpInterface.GetPropertiesHashMap(), ownerTransformInfoHttpInterface.GetPageContext()); _ = transformInfoInterface
            if transformInfoInterface == nil {
                panic(errors.New(fmt.Sprint("No Such View In DB: " + viewName)))
            }
            transformInfoInterface.SetChild()
            var anyType any = TransformInfoObjectFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface); _ = anyType
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Created Transform: " + viewName, this, this.CommonStrings.GET_INSTANCE)
            }
            tryControl1, tryResult1 = 1, any(anyType).(TransformInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformFactory) GetInstance(abeClientInformation AbeClientInformationInterface, propertiesHashMap *HashMap, pageContext *PageContext) TransformInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Failed To Get Instance: ")
                stringBuffer.Append(any(propertiesHashMap.Get(TransformInfoData_GetInstance().NAME)).(string))
                stringBuffer.Append("->TransformFactory")
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, this.CommonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Creating Transform: " + fmt.Sprint(propertiesHashMap.Get(transformInfoData.NAME)), this, this.CommonStrings.GET_INSTANCE)
            }
            var transformInfoEntity *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntity
            var transformInfoInterface TransformInfoInterface = transformInfoEntity.Get(any(propertiesHashMap.Get(transformInfoData.NAME)).(string), propertiesHashMap, pageContext); _ = transformInfoInterface
            if transformInfoInterface != nil {
                transformInfoInterface.Override(propertiesHashMap)
            } else {
                transformInfoInterface = any(TransformInfoHttpFactory_GetInstance(propertiesHashMap, pageContext)).(TransformInfoInterface)
            }
            var anyType any = TransformInfoObjectFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface); _ = anyType
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Created Transform: " + fmt.Sprint(propertiesHashMap.Get(transformInfoData.NAME)), this, this.CommonStrings.GET_INSTANCE)
            }
            tryControl1, tryResult1 = 1, any(anyType).(TransformInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformFactory) GetInstance(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) TransformInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(TransformInfoInterface)", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Creating Transform: " + transformInfoInterface.GetName(), this, "getInstance(TransformInfoInterface)")
            }
            var anyType any = TransformInfoObjectFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface); _ = anyType
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Created Transform: " + transformInfoInterface.GetName(), this, "getInstance(TransformInfoInterface)")
            }
            tryControl1, tryResult1 = 1, any(anyType).(TransformInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

