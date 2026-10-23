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

type TransformGeneratorUtil struct {
    LogUtil *LogUtil
}

var TransformGeneratorUtil_Instance *TransformGeneratorUtil = NewTransformGeneratorUtil()

func (this *TransformGeneratorUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformGeneratorUtil_GetInstance() *TransformGeneratorUtil {
        return TransformGeneratorUtil_Instance
}


func NewTransformGeneratorUtil() *TransformGeneratorUtil {
    this := &TransformGeneratorUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformGeneratorUtil) Generate_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, ownerTransformInfoInterface TransformInfoInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to generate a view", this, "generate()", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Generating View: " + transformInfoInterface.GetName(), this, "generate()")
            }
            var componentInterface TransformInterface = TransformFactory_GetInstance().GetInstance(abeClientInformation, transformInfoInterface.GetName(), ownerTransformInfoInterface); _ = componentInterface
            var result string = componentInterface.View(); _ = result
            var httpTransformInfoInterface TransformInfoHttpInterface = any(componentInterface.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var inputOutputTypeData *InputOutputTypeData = InputOutputTypeData_GetInstance(); _ = inputOutputTypeData
            if strings.Index(result, string("<HTML>")) >= 0 {
                httpTransformInfoInterface.GetPropertiesHashMap().Put(inputOutputTypeData.NAME, inputOutputTypeData.DEFAULT)
            } else {
                httpTransformInfoInterface.GetPropertiesHashMap().Put(inputOutputTypeData.NAME, inputOutputTypeData.DEFAULT_FRAGMENT)
            }
            NewStoreFileGenerator(componentInterface.GetTransformInfoInterface()).Process(result)
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of TransformGeneratorUtil.generate
func (this *TransformGeneratorUtil) Generate(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[TransformInfoInterface](args[2], pass) {
            this.Generate_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[TransformInfoInterface](args[2]))
            return
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[ITransformInfoDomNode](args[1], pass) && Java_IsArg[TransformInfoInterface](args[2], pass) {
            this.Generate_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[ITransformInfoDomNode](args[1]), Java_Arg[TransformInfoInterface](args[2]))
            return
        }
    }
    panic("No Java overload of TransformGeneratorUtil.generate for the arguments")
}


// throws Exception
func (this *TransformGeneratorUtil) Generate_2(abeClientInformation AbeClientInformationInterface, transformInfoDomNode ITransformInfoDomNode, ownerTransformInfoInterface TransformInfoInterface) {
        this.Generate(abeClientInformation, transformInfoDomNode.GetTransformInfoInterface(), ownerTransformInfoInterface)
}

