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
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, objectConfigFileAbPath IAbPath) TransformInfoObjectConfigInterface {
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

// Java overloads of TransformInfoObjectConfigAndManipulatorFactory.getInstance
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance(args ...any) TransformInfoObjectConfigInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[IAbPath](args[2], pass) {
            return this.GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[IAbPath](args[2]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) {
            return this.GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[*Document](args[2], pass) {
            return this.GetInstance_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[*Document](args[2]))
        }
    }
    panic("No Java overload of TransformInfoObjectConfigAndManipulatorFactory.getInstance for the arguments")
}


// throws Exception
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance_2(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) TransformInfoObjectConfigInterface {
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
func (this *TransformInfoObjectConfigAndManipulatorFactory) GetInstance_3(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) TransformInfoObjectConfigInterface {
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

