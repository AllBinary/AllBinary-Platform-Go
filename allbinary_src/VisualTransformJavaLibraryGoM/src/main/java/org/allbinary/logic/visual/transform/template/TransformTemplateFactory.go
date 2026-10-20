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

type TransformTemplateFactory struct {
    LogUtil *LogUtil
}

var TransformTemplateFactory_Instance *TransformTemplateFactory = NewTransformTemplateFactory()

func (this *TransformTemplateFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformTemplateFactory_GetInstance() *TransformTemplateFactory {
        return TransformTemplateFactory_Instance
}


func NewTransformTemplateFactory() *TransformTemplateFactory {
    this := &TransformTemplateFactory{}
    this.initFields()
    return this
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformTemplateFactory) GetInstance(abeClientInformation AbeClientInformationInterface, templateName string, propertiesHashMap *HashMap, pageContext *PageContext) TransformTemplateInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformTemplateInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Get Instance"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(error_ + templateName + "->TemplateFactory", this, "getInstance(viewName, HashMap, PageContext)", e)
            }
            panic(e)
                }
            }()
            var transformInfoEntity *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntity
            var transformInfoInterface TransformInfoInterface = transformInfoEntity.Get(templateName, propertiesHashMap, pageContext); _ = transformInfoInterface
            tryControl1, tryResult1 = 1, any(TransformInfoObjectFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface)).(TransformTemplateInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

