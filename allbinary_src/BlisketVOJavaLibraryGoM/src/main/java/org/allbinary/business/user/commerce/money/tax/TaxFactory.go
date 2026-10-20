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

type TaxFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TAXATIONFILENAME string
}

var TaxFactory_Instance *TaxFactory = NewTaxFactory()

func (this *TaxFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.TAXATIONFILENAME = "taxationMethod.xml"
}


func TaxFactory_GetInstance() *TaxFactory {
        return TaxFactory_Instance
}


func NewTaxFactory() *TaxFactory {
    this := &TaxFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *TaxFactory) GetInstance(abeClientInformation AbeClientInformationInterface, storeFrontInterface StoreFrontInterface) TaxModuleInterface {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(URLGLOBALS.GetMainPath())
        stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(storeFrontInterface.GetName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append("generic")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append("taxes")
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        var abPath IAbPath = NewAbPath(fmt.Sprint(stringBuffer), this.TAXATIONFILENAME); _ = abPath
        var transformInfoObjectConfigData *TransformInfoObjectConfigData = TransformInfoObjectConfigData_GetInstance(); _ = transformInfoObjectConfigData
        var data string = NewCryptFileReader(transformInfoObjectConfigData.UNCRYPTED_EXTENSION, transformInfoObjectConfigData.ENCRYPTED_EXTENSION).Get(abPath); _ = data
        var document *Document = DomDocumentHelper_Create(data); _ = document
        var taxNameNodeList *NodeList = document.GetElementsByTagName(TaxData_NAME); _ = taxNameNodeList
        for index := int(0); index < taxNameNodeList.GetLength(); index++ {
            var node *Node = taxNameNodeList.Item(index); _ = node
            var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
            if nodeList != nil {
                var classNameNode *Node = DomSearchHelper_GetNode(DynamicObjectData_NAME, nodeList); _ = classNameNode
                if classNameNode != nil {
                    var className string = DomNodeHelper_GetTextNodeValue(classNameNode); _ = className
                    return any(AbeFactory_GetInstance().GetInstance(abeClientInformation, className)).(TaxModuleInterface)
                } else {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAX) {
                        this.LogUtil.PutF("Class Node Null", this, this.CommonStrings.GET_INSTANCE)
                    }
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAX) {
                    this.LogUtil.PutF("Tax Name Node Node Children", this, this.CommonStrings.GET_INSTANCE)
                }
            }
        }
        return any(NewSimpleStateTaxModule()).(TaxModuleInterface)
}

