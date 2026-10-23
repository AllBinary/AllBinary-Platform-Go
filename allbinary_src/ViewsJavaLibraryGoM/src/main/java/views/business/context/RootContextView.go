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

type RootContextView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}

// Implements TransformTemplateInterface


func (this *RootContextView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewRootContextView(transformInfoInterface TransformInfoInterface) *RootContextView {
    this := &RootContextView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "RootContextView()")
        }
    return this
}


// throws Exception
func (this *RootContextView) View() string {
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
            var viewName string = this.GetName(); _ = viewName
            var componentInterface TransformInterface = TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, viewName, this.GetTransformInfoInterface()); _ = componentInterface
            tryControl1, tryResult1 = 1, componentInterface.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *RootContextView) GetName() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getName()", e)
            }
            panic(e)
                }
            }()
            var objectConfig INoTemplateTransformInfoObjectConfig = NewNoTemplateTransformInfoObjectConfig(this.GetTransformInfoInterface(), this.GetTransformInfoInterface().GetObjectConfigInterface().ToXmlDoc()); _ = objectConfig
            var componentsVector IBasicArrayList = objectConfig.GetTransforms(); _ = componentsVector
            if componentsVector.Size() == 0 {
                panic(errors.New(fmt.Sprint("Template View Not Set - No Components")))
            }
            if componentsVector.Size() > 1 {
                panic(errors.New(fmt.Sprint("To Many Root Template Views: " + fmt.Sprint(componentsVector.Size()))))
            }
            var transformInfoInterface ITransformInfo = any(componentsVector.Get(0)).(ITransformInfo); _ = transformInfoInterface
            var viewName string = transformInfoInterface.GetName(); _ = viewName
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Root Template With View Name: " + viewName, this, "getName()")
            }
            if StringValidationUtil_GetInstance().IsEmpty(viewName) {
                panic(errors.New(fmt.Sprint("Template View Not Set - No View Name")))
            }
            tryControl1, tryResult1 = 1, viewName; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

