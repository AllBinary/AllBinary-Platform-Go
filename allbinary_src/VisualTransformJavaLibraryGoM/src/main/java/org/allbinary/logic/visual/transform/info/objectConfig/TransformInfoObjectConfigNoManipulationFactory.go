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

type TransformInfoObjectConfigNoManipulationFactory struct {
    LogUtil *LogUtil
}

var TransformInfoObjectConfigNoManipulationFactory_Instance *TransformInfoObjectConfigNoManipulationFactory = NewTransformInfoObjectConfigNoManipulationFactory()

func (this *TransformInfoObjectConfigNoManipulationFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformInfoObjectConfigNoManipulationFactory_GetInstance() *TransformInfoObjectConfigNoManipulationFactory {
        return TransformInfoObjectConfigNoManipulationFactory_Instance
}


func NewTransformInfoObjectConfigNoManipulationFactory() *TransformInfoObjectConfigNoManipulationFactory {
    this := &TransformInfoObjectConfigNoManipulationFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformInfoObjectConfigNoManipulationFactory) GetInstance(transformInfoInterface TransformInfoInterface, objectConfigFileAbPath IAbPath) TransformInfoObjectConfigInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransformInfoObjectConfigInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Could Not Load Object Config", this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var transformInfoObjectConfigData *TransformInfoObjectConfigData = TransformInfoObjectConfigData_GetInstance(); _ = transformInfoObjectConfigData
            var data string = NewCryptFileReader(transformInfoObjectConfigData.UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(objectConfigFileAbPath); _ = data
            tryControl1, tryResult1 = 1, NewTransformInfoObjectConfig(transformInfoInterface, DomDocumentHelper_Create(data)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

