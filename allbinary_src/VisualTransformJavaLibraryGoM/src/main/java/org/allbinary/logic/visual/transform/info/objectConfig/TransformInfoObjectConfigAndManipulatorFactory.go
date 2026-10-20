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

type TransformInfoObjectConfigAndManipulatorFactory struct {
    TransformInfoObjectConfigAndManipulatorFactoryBase
    LogUtil *LogUtil
}

var TransformInfoObjectConfigAndManipulatorFactory_Instance *TransformInfoObjectConfigAndManipulatorFactory = NewTransformInfoObjectConfigAndManipulatorFactory()

func (this *TransformInfoObjectConfigAndManipulatorFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformInfoObjectConfigAndManipulatorFactory_GetInstance() *TransformInfoObjectConfigAndManipulatorFactory {
        return TransformInfoObjectConfigAndManipulatorFactory_Instance
}


func NewTransformInfoObjectConfigAndManipulatorFactory() *TransformInfoObjectConfigAndManipulatorFactory {
    this := &TransformInfoObjectConfigAndManipulatorFactory{}
    this.TransformInfoObjectConfigAndManipulatorFactoryBase = *NewTransformInfoObjectConfigAndManipulatorFactoryBase()
    this.initFields()
    return this
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, objectConfigFileAbPath *AbPath) TransformInfoObjectConfigInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInfoObjectConfigInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Could Not Load Object Config", TransformInfoObjectConfigAndManipulatorFactory_GetInstance(), this.CommonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var data string = NewCryptFileReader(TransformInfoObjectConfigData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(objectConfigFileAbPath); _ = data
            tryControl1, tryResult1 = 1, this.GetInstance(abeClientInformation, transformInfoInterface, DomDocumentHelper_Create(data)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) TransformInfoObjectConfigInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInfoObjectConfigInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Get Instance: "; _ = error_
            if transformInfoInterface != nil {
                error_ = error_ + transformInfoInterface.GetName()
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, TransformInfoObjectConfigAndManipulatorFactory_GetInstance(), this.CommonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if transformInfoInterface != nil && !stringValidationUtil.IsEmpty(transformInfoInterface.GetStoreName()) {
                tryControl1, tryResult1 = 1, any(NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface)).(TransformInfoObjectConfigInterface); return
            } else {
                tryControl1, tryResult1 = 1, any(NewTransformInfoObjectConfig(transformInfoInterface)).(TransformInfoObjectConfigInterface); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) TransformInfoObjectConfigInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInfoObjectConfigInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Get Instance: "; _ = error_
            if transformInfoInterface != nil {
                error_ = error_ + transformInfoInterface.GetName()
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, TransformInfoObjectConfigAndManipulatorFactory_GetInstance(), "getInstance(document)", e)
            }
            panic(e)
                }
            }()
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if transformInfoInterface != nil && !stringValidationUtil.IsEmpty(transformInfoInterface.GetStoreName()) {
                tryControl1, tryResult1 = 1, any(NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface, document)).(TransformInfoObjectConfigInterface); return
            } else {
                tryControl1, tryResult1 = 1, any(NewTransformInfoObjectConfig(transformInfoInterface, document)).(TransformInfoObjectConfigInterface); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

