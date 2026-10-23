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

type CustomizerUtil struct {
    LogUtil *LogUtil
}

var CustomizerUtil_Instance *CustomizerUtil = NewCustomizerUtil()

func (this *CustomizerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func CustomizerUtil_GetInstance() *CustomizerUtil {
        return CustomizerUtil_Instance
}


func NewCustomizerUtil() *CustomizerUtil {
    this := &CustomizerUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *CustomizerUtil) GetTransformInfoInterfaceToCustomize(customizerTransformInfoInterface TransformInfoInterface, transformInfoInterface TransformInfoInterface) TransformInfoInterface {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + customizerTransformInfoInterface.GetName(), this, "getTransformInfoInterfaceToCustomize()")
        }
        var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
        var viewName string = transformInfoInterface.GetName(); _ = viewName
        var httpTransformInfoInterface TransformInfoHttpInterface = any(customizerTransformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        var customizedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewName, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = customizedTransformInfoInterface
        return customizedTransformInfoInterface
}


// throws Exception
func (this *CustomizerUtil) GetTransformInfoObjectConfigComponentNodesToCustomize(customizerTransformInfoInterface TransformInfoInterface, transformInfoInterface TransformInfoInterface) IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + customizerTransformInfoInterface.GetName(), this, "getTransformInfoObjectConfigComponentNodesToCustomize()")
        }
        var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
        var viewName string = transformInfoInterface.GetName(); _ = viewName
        var httpTransformInfoInterface TransformInfoHttpInterface = any(customizerTransformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        var customizedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewName, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = customizedTransformInfoInterface
        return this.GetTransformInfoObjectConfigComponentNodes(customizedTransformInfoInterface)
}


// throws Exception
func (this *CustomizerUtil) GetTransformInfoObjectConfigComponentNodes(transformInfoInterface TransformInfoInterface) IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "getTransformInfoObjectConfigComponentNodes()")
        }
        var transformInfoObjectConfigInterface TransformInfoObjectConfigInterface = transformInfoInterface.GetObjectConfigInterface(); _ = transformInfoObjectConfigInterface
        return transformInfoObjectConfigInterface.GetTransforms()
}


// throws Exception
func (this *CustomizerUtil) GetTransformInfoObjectConfigGroupComponentNodes(transformInfoInterface TransformInfoInterface) IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "getTransformInfoObjectConfigComponentNodes()")
        }
        var transformInfoObjectConfigInterface TransformInfoObjectConfigInterface = transformInfoInterface.GetObjectConfigInterface(); _ = transformInfoObjectConfigInterface
        return transformInfoObjectConfigInterface.GetGroupTransforms()
}


// throws Exception
func (this *CustomizerUtil) Insert(transformInfoInterface TransformInfoInterface, domNodeInterface DomNodeInterface) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "insert()")
        }
        var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
        var allViewsToBeModified IBasicArrayList = this.GetTransformInfoObjectConfigGroupComponentNodes(transformInfoInterface); _ = allViewsToBeModified
        var document *Document = DomDocumentHelper_Create(); _ = document
        document.AppendChild(domNodeInterface.ToXmlNode(document))
        var documentString string = XmlDocumentHelper_ToString(document); _ = documentString
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            stringBuffer.Append("Total Views For Modification: ")
            stringBuffer.Appendint(allViewsToBeModified.Size())
            stringBuffer.Append(" New Xml/Data: ")
            stringBuffer.Append(documentString)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
        }
        transformInfoInterface.SetData(documentString)
        var size int = allViewsToBeModified.Size(); _ = size
        for index := int(0); index < size; index++ {
            var transformInfo ITransformInfo = any(allViewsToBeModified.Get(index)).(ITransformInfo); _ = transformInfo
            var viewNameOfViewToBeModified string = transformInfo.GetName(); _ = viewNameOfViewToBeModified
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(transformInfoInterface.GetName())
                stringBuffer.Append(" is modifying view: ")
                stringBuffer.Append(viewNameOfViewToBeModified)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
            }
            var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var specifiedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewNameOfViewToBeModified, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = specifiedTransformInfoInterface
            if specifiedTransformInfoInterface == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append(transformInfoInterface.GetName())
                    stringBuffer.Append(" is unable to add data to view: ")
                    stringBuffer.Append(viewNameOfViewToBeModified)
                    stringBuffer.Append(" since it does not exist")
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF(transformInfoInterface.GetName() + " is adding data to view: " + viewNameOfViewToBeModified, this, "insert()")
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append(viewNameOfViewToBeModified)
                    stringBuffer.Append(" is changing data in ")
                    stringBuffer.Append(fmt.Sprint(specifiedTransformInfoInterface.GetDataFilePath()))
                    stringBuffer.Append(" to the following data:\n")
                    stringBuffer.Append(documentString)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "insert()")
                }
                this.Write(specifiedTransformInfoInterface, documentString)
            }
        }
}


// throws Exception
func (this *CustomizerUtil) Write(specifiedTransformInfoInterface TransformInfoInterface, documentString string) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + specifiedTransformInfoInterface.GetName(), this, "write()")
        }
        var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
        specifiedTransformInfoInterface.SetData(documentString)
        transformInfoEntityInterface.Update(specifiedTransformInfoInterface.ToHashMap())
}


// throws Exception
func (this *CustomizerUtil) GetViewDataForComponentsInObjectConfig(transformInfoInterface TransformInfoInterface) *Document {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "getViewDataForComponentsInObjectConfig()")
        }
        var transformInfoObjectConfigInterface TransformInfoObjectConfigInterface = transformInfoInterface.GetObjectConfigInterface(); _ = transformInfoObjectConfigInterface
        var allViewsToBeModified IBasicArrayList = transformInfoObjectConfigInterface.GetGroupTransforms(); _ = allViewsToBeModified
        var transformInfoEntityInterface *TransformInfoEntity = TransformInfoEntityBuilder_GetInstance(); _ = transformInfoEntityInterface
        var size int = allViewsToBeModified.Size(); _ = size
        for index := int(0); index < size; index++ {
            var nextTransformInfoInterface ITransformInfo = any(allViewsToBeModified.Get(index)).(ITransformInfo); _ = nextTransformInfoInterface
            var viewNameOfViewToBeModified string = nextTransformInfoInterface.GetName(); _ = viewNameOfViewToBeModified
            var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var specifiedTransformInfoInterface TransformInfoInterface = transformInfoEntityInterface.Get(viewNameOfViewToBeModified, httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = specifiedTransformInfoInterface
            return specifiedTransformInfoInterface.GetDataDocument()
        }
        panic(errors.New(fmt.Sprint("No Components Specified in ObjectConfig")))
}

