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

type TransformsGeneratorUtil struct {
    LogUtil *LogUtil
}

var TransformsGeneratorUtil_Instance *TransformsGeneratorUtil = NewTransformsGeneratorUtil()

func (this *TransformsGeneratorUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func TransformsGeneratorUtil_GetInstance() *TransformsGeneratorUtil {
        return TransformsGeneratorUtil_Instance
}


func NewTransformsGeneratorUtil() *TransformsGeneratorUtil {
    this := &TransformsGeneratorUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformsGeneratorUtil) GenerateComponentsFromObjectConfig(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, group string) string {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Started Group: " + group, this, "generateComponentsFromObjectConfig(2)")
        }
        var transformInfoObjectConfigInterface TransformInfoObjectConfigInterface = transformInfoInterface.GetObjectConfigInterface(); _ = transformInfoObjectConfigInterface
        var result string = this.GenerateComponentsFromObjectConfig(abeClientInformation, transformInfoObjectConfigInterface, transformInfoInterface, group); _ = result
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("End Group: " + group, this, "generateComponentsFromObjectConfig(2)")
        }
        return result
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *TransformsGeneratorUtil) GenerateComponentsFromObjectConfig(abeClientInformation AbeClientInformationInterface, transformInfoObjectConfigInterface TransformInfoObjectConfigInterface, transformInfoInterface TransformInfoInterface, group string) string {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Started Group: " + group, this, "generateComponentsFromObjectConfig()")
        }
        var transformInfoObjectConfigComponentVector *BasicArrayList = transformInfoObjectConfigInterface.GetTransformsGroup(group); _ = transformInfoObjectConfigComponentVector
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Processing ")
            stringBuffer.Appendint(transformInfoObjectConfigComponentVector.Size())
            stringBuffer.Append(" Components Group: ")
            stringBuffer.Append(group)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "generateComponentsFromObjectConfig()")
        }
        if transformInfoObjectConfigComponentVector.Size() < 1 {
            panic(errors.New(fmt.Sprint("No Pages Generated For Template.")))
        }
        var size int = transformInfoObjectConfigComponentVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var transformInfoObjectConfigComponent *TransformInfoDomNode = any(transformInfoObjectConfigComponentVector.Get(index)).(*TransformInfoDomNode); _ = transformInfoObjectConfigComponent
            TransformGeneratorUtil_GetInstance().Generate(abeClientInformation, transformInfoObjectConfigComponent, transformInfoInterface)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Done Group: " + group, this, "generateComponentsFromObjectConfig()")
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("<!-- Generated all template views successfully for View: ")
        stringBuffer.Append(transformInfoInterface.GetName())
        stringBuffer.Append("-->")
        return fmt.Sprint(stringBuffer)
}

