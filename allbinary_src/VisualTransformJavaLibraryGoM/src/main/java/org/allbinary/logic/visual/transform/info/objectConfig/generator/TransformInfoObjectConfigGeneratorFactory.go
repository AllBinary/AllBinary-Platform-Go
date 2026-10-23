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

type TransformInfoObjectConfigGeneratorFactory struct {
    LogUtil *LogUtil
}

// Implements TransformInfoObjectConfigGeneratorFactoryInterface

var TransformInfoObjectConfigGeneratorFactory_Instance *TransformInfoObjectConfigGeneratorFactory = NewTransformInfoObjectConfigGeneratorFactory()

func (this *TransformInfoObjectConfigGeneratorFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformInfoObjectConfigGeneratorFactory_GetInstance() *TransformInfoObjectConfigGeneratorFactory {
        return TransformInfoObjectConfigGeneratorFactory_Instance
}


func NewTransformInfoObjectConfigGeneratorFactory() *TransformInfoObjectConfigGeneratorFactory {
    this := &TransformInfoObjectConfigGeneratorFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformInfoObjectConfigGeneratorFactory) GetInstance(transformInfoInterface TransformInfoInterface) ITransformInfoObjectConfigGenerator {
        var type_ string = any(transformInfoInterface.GetObjectConfigInterface().GetOutputTypeName()).(string); _ = type_
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF("Generating Instance for type: " + type_, this, commonStrings.GET_INSTANCE)
        }
        if type_ == "" || strings.Compare(type_, InputOutputTypeData_GetInstance().RESPONSE) == 0 || transformInfoInterface.IsChild() {
            return NewBasicGenerator()
        } else if strings.Compare(type_, InputOutputTypeData_GetInstance().FILE) == 0 {
            return NewStoreFileGenerator(transformInfoInterface)
        } else {
            panic(errors.New(fmt.Sprint("No Such TransformInfoObjectConfigGenerator Type")))
        }
}

