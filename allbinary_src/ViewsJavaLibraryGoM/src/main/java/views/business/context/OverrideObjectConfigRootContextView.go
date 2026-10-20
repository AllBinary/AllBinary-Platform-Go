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

type OverrideObjectConfigRootContextView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}


func (this *OverrideObjectConfigRootContextView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewOverrideObjectConfigRootContextView(transformInfoInterface TransformInfoInterface) *OverrideObjectConfigRootContextView {
    this := &OverrideObjectConfigRootContextView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "OverrideObjectRootContextView()")
        }
    return this
}


// throws Exception
func (this *OverrideObjectConfigRootContextView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("View Name: " + this.GetTransformInfoInterface().GetName(), this, "view()")
            }
            var rootTransformInfoInterface TransformInfoInterface = any(NewTransformInfoHttp(any(this.GetTransformInfoInterface()).(*TransformInfoHttp))).(TransformInfoInterface); _ = rootTransformInfoInterface
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.GetTransformInfoInterface().GetStoreName())
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            var nameStart string = fmt.Sprint(stringBuffer); _ = nameStart
            if strings.Index(this.GetTransformInfoInterface().GetName(), string(GeneratorTransformInfoData_NAME)) > 0 {
                rootTransformInfoInterface.SetName(nameStart + fmt.Sprint(GeneratorTransformInfoData_NAME))
            } else if strings.Index(this.GetTransformInfoInterface().GetName(), string(RootTransformInfoData_NAME)) > 0 {
                rootTransformInfoInterface.SetName(nameStart + fmt.Sprint(RootTransformInfoData_NAME))
            } else {
                panic(errors.New(fmt.Sprint("Override Currently Not Supported")))
            }
            var overrideObjectConfigDocument *Document = this.GetTransformInfoInterface().GetObjectConfigInterface().ToXmlDoc(); _ = overrideObjectConfigDocument
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Retrieved OverrideObjectConfig: " + XmlDocumentHelper_ToString(overrideObjectConfigDocument), this, "view()")
            }
            var transformFactory *TransformFactory = TransformFactory_GetInstance(); _ = transformFactory
            var rootComponentInterface TransformInterface = transformFactory.GetInstance(this.AbeClientInformation, rootTransformInfoInterface.GetName(), rootTransformInfoInterface); _ = rootComponentInterface
            var transformTemplateInterface TransformTemplateInterface = any(rootComponentInterface).(TransformTemplateInterface); _ = transformTemplateInterface
            var componentInterface TransformInterface = transformFactory.GetInstance(this.AbeClientInformation, transformTemplateInterface.GetName(), rootComponentInterface.GetTransformInfoInterface()); _ = componentInterface
            var transformInfoObjectConfigInterface TransformInfoObjectConfigInterface = TransformInfoObjectConfigAndManipulatorFactory_GetInstance().GetInstance(this.AbeClientInformation, componentInterface.GetTransformInfoInterface(), overrideObjectConfigDocument); _ = transformInfoObjectConfigInterface
            componentInterface.GetTransformInfoInterface().SetObjectConfigInterface(transformInfoObjectConfigInterface)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Set OverrideObjectConfig: " + fmt.Sprint(componentInterface.GetTransformInfoInterface().GetObjectConfigInterface()), this, "view()")
            }
            tryControl1, tryResult1 = 1, componentInterface.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

