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

type StoreCustomizerComponentUtil struct {
    LogUtil *LogUtil
}

var StoreCustomizerComponentUtil_Instance *StoreCustomizerComponentUtil = NewStoreCustomizerComponentUtil()

func (this *StoreCustomizerComponentUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func StoreCustomizerComponentUtil_GetInstance() *StoreCustomizerComponentUtil {
        return StoreCustomizerComponentUtil_Instance
}


func NewStoreCustomizerComponentUtil() *StoreCustomizerComponentUtil {
    this := &StoreCustomizerComponentUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *StoreCustomizerComponentUtil) Generate_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) string {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Generating Modified Views For: " + transformInfoInterface.GetName(), this, "generateModifiedViews()")
        }
        var customizerUtil *CustomizerUtil = CustomizerUtil_GetInstance(); _ = customizerUtil
        var allCustomizedViews IBasicArrayList = customizerUtil.GetTransformInfoObjectConfigGroupComponentNodes(transformInfoInterface); _ = allCustomizedViews
        var size int = allCustomizedViews.Size(); _ = size
        for index := int(0); index < size; index++ {
            var nextTransformInfoInterface ITransformInfo = any(allCustomizedViews.Get(index)).(ITransformInfo); _ = nextTransformInfoInterface
            var allViewsToBeModified IBasicArrayList = customizerUtil.GetTransformInfoObjectConfigComponentNodesToCustomize(transformInfoInterface, nextTransformInfoInterface); _ = allViewsToBeModified
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("View or component that was modified by a Customizer: " + nextTransformInfoInterface.GetName(), this, "generateModifiedViews()")
            }
            this.Generate(abeClientInformation, transformInfoInterface, allViewsToBeModified)
        }
        return StringUtil_GetInstance().EMPTY_STRING
}

// Java overloads of StoreCustomizerComponentUtil.generate
func (this *StoreCustomizerComponentUtil) Generate(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) {
            return this.Generate_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[IBasicArrayList](args[2], pass) {
            return this.Generate_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[IBasicArrayList](args[2]))
        }
    }
    panic("No Java overload of StoreCustomizerComponentUtil.generate for the arguments")
}


// throws Exception
func (this *StoreCustomizerComponentUtil) Generate_2(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, allViewsToBeModifiedVector IBasicArrayList) {
        var customizerUtil *CustomizerUtil = CustomizerUtil_GetInstance(); _ = customizerUtil
        var size int = allViewsToBeModifiedVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var nextTransformInfoInterface ITransformInfo = any(allViewsToBeModifiedVector.Get(index)).(ITransformInfo); _ = nextTransformInfoInterface
            var specifiedTransformInfoInterface TransformInfoInterface = customizerUtil.GetTransformInfoInterfaceToCustomize(transformInfoInterface, nextTransformInfoInterface); _ = specifiedTransformInfoInterface
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Retrieved: " + specifiedTransformInfoInterface.GetName(), this, "generateModifiedViews()")
            }
            TransformGeneratorUtil_GetInstance().Generate(abeClientInformation, nextTransformInfoInterface, specifiedTransformInfoInterface)
        }
}

