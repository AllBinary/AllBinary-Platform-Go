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

type CompoundContextTransformInfoObjectConfig struct {
    GenericStoreTransformInfoObjectConfig
    LogUtil *LogUtil
}


func (this *CompoundContextTransformInfoObjectConfig) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new CompoundContextTransformInfoObjectConfig
func NewCompoundContextTransformInfoObjectConfig(args ...any) *CompoundContextTransformInfoObjectConfig {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) {
            return NewCompoundContextTransformInfoObjectConfig_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[*Document](args[2], pass) {
            return NewCompoundContextTransformInfoObjectConfig_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[*Document](args[2]))
        }
        if len(args) == 4 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return NewCompoundContextTransformInfoObjectConfig_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
    }
    panic("No Java overload of new CompoundContextTransformInfoObjectConfig for the arguments")
}
func NewCompoundContextTransformInfoObjectConfig_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) *CompoundContextTransformInfoObjectConfig {
    this := &CompoundContextTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface)
        this.initFields()
    return this
}


func NewCompoundContextTransformInfoObjectConfig_2(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) *CompoundContextTransformInfoObjectConfig {
    this := &CompoundContextTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface, document)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


func NewCompoundContextTransformInfoObjectConfig_3(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, name string, type_ string) *CompoundContextTransformInfoObjectConfig {
    this := &CompoundContextTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface, name, type_)
        this.initFields()
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// throws Exception
func (this *CompoundContextTransformInfoObjectConfig) Get() string {
        var storeName string = this.GetTransformInfoInterface().GetStoreName(); _ = storeName
        var httpTransformInfoInterface TransformInfoHttpInterface = any(this.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        var templateNameOverride string = StringUtil_GetInstance().GetNonNull(any(httpTransformInfoInterface.GetPropertiesHashMap().Get(TransformInfoData_GetInstance().PARTIAL)).(string)); _ = templateNameOverride
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(storeName)
        stringBuffer.Append(templateNameOverride)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(RootTransformInfoData_NAME)
        var rootComponentInterface TransformInterface = TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, fmt.Sprint(stringBuffer), this.GetTransformInfoInterface()); _ = rootComponentInterface
        var viewVector IBasicArrayList = this.GetTransformDomNodes(); _ = viewVector
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Root View Called with: " + fmt.Sprint(viewVector.Size()), this, "get(transformInfoInterface)")
        }
        var rootView string = rootComponentInterface.View(); _ = rootView
        var startXMLHeader string = "<xsl:text disable-output-escaping=\"yes\" ><![CDATA["; _ = startXMLHeader
        var endXMLHeader string = "]]></xsl:text>"; _ = endXMLHeader
        var size int = viewVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var objectConfigTransformInfoDomNode ITransformInfoDomNode = any(viewVector.Get(index)).(ITransformInfoDomNode); _ = objectConfigTransformInfoDomNode
            var templateKey string = objectConfigTransformInfoDomNode.GetReplaceKey(); _ = templateKey
            var componentInterface TransformInterface = TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, objectConfigTransformInfoDomNode.GetTransformInfoInterface().GetName(), this.GetTransformInfoInterface()); _ = componentInterface
            var replacementViewString string = componentInterface.View(); _ = replacementViewString
            replacementViewString = DocumentToNode_ConvertDocumentToNodeString(replacementViewString)
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(startXMLHeader)
            stringBuffer.Append(replacementViewString)
            stringBuffer.Append(endXMLHeader)
            var templateValue string = fmt.Sprint(stringBuffer); _ = templateValue
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Replacing: ")
                stringBuffer.Append(templateKey)
                stringBuffer.Append(" with ")
                stringBuffer.Append(templateValue)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "get()")
            }
            var replace *Replace = NewReplace(templateKey, templateValue); _ = replace
            rootView = replace.All(rootView)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("End: Result of a compound view: " + rootView, this, "get()")
        }
        return rootView
}


// throws Exception
func (this *CompoundContextTransformInfoObjectConfig) CreateInputStream() IInputStream {
        var completeTemplateViewBytes []int8 = Java_GetBytes(this.Get()); _ = completeTemplateViewBytes
        var bais *ByteArrayInputStream = NewByteArrayInputStream(completeTemplateViewBytes); _ = bais
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Template: " + string(completeTemplateViewBytes), this, "createInputStream()")
        }
        return bais
}

