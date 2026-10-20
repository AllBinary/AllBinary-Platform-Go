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

type ShippingMethodsFactory struct {
    DefaultShippingMethodInterface ShippingInterface
    Document *Document
    ShippingVector IBasicArrayList
}

var ShippingMethodsFactory_SHIPPINGMETHODSFILEPATHSTRING string
var ShippingMethodsFactory_SHIPPINGMETHODSFILENAME string = "shippingMethods.xml"

func (this *ShippingMethodsFactory) initFields() {
    
    
    
}


func init() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var sep string = AbPathData_GetInstance().SEPARATOR; _ = sep
        stringBuffer.Append(sep)
        stringBuffer.Append("generic")
        stringBuffer.Append(sep)
        stringBuffer.Append("shipping")
        stringBuffer.Append(sep)
        ShippingMethodsFactory_SHIPPINGMETHODSFILEPATHSTRING = fmt.Sprint(stringBuffer)
}


func NewShippingMethodsFactory(abeClientInformation AbeClientInformationInterface, storeFrontInterface StoreFrontInterface) *ShippingMethodsFactory {
    this := &ShippingMethodsFactory{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var sep string = AbPathData_GetInstance().SEPARATOR; _ = sep
        stringBuffer.Append(URLGLOBALS.GetMainPath())
        stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(storeFrontInterface.GetName())
        stringBuffer.Append(sep)
        stringBuffer.Append(ShippingMethodsFactory_SHIPPINGMETHODSFILEPATHSTRING)
        var abPath IAbPath = any(NewAbPath(fmt.Sprint(stringBuffer), ShippingMethodsFactory_SHIPPINGMETHODSFILENAME)).(IAbPath); _ = abPath
        var data string = NewCryptFileReader(TransformInfoObjectConfigData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(abPath); _ = data
        this.Document = DomDocumentHelper_Create(data)
        this.ShippingVector = NewBasicArrayListD()
        var nodeList *NodeList = this.Document.GetElementsByTagName(ShippingMethodsData_NAME); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var node *Node = nodeList.Item(index); _ = node
            var shippingMethodNodeChildren *NodeList = node.GetChildNodes(); _ = shippingMethodNodeChildren
            var shippingMethodNameNode *Node = DomSearchHelper_GetNode(ShippingMethodData_NAME, shippingMethodNodeChildren); _ = shippingMethodNameNode
            var classNameNode *Node = DomSearchHelper_GetNode(DynamicObjectData_NAME, shippingMethodNameNode.GetChildNodes()); _ = classNameNode
            var shippingMethodClassName string = DomNodeHelper_GetTextNodeValue(classNameNode); _ = shippingMethodClassName
            var shippingMethodInterface ShippingInterface = any(AbeFactory_GetInstance().GetInstance(abeClientInformation, shippingMethodClassName)).(ShippingInterface); _ = shippingMethodInterface
            this.ShippingVector.Add(shippingMethodInterface)
            var defaultShippingMethodNameNode *Node = DomSearchHelper_GetNode(ShippingMethodData_DEFAULT, shippingMethodNameNode.GetChildNodes()); _ = defaultShippingMethodNameNode
            var defaultShippingMethodNameNodeValue string = DomNodeHelper_GetTextNodeValue(defaultShippingMethodNameNode); _ = defaultShippingMethodNameNodeValue
            if defaultShippingMethodNameNodeValue != "" && strings.Compare(defaultShippingMethodNameNodeValue, ShippingMethodData_DEFAULT) == 0 {
                this.DefaultShippingMethodInterface = shippingMethodInterface
            }
        }
        if this.ShippingVector.Size() < 1 {
            this.DefaultShippingMethodInterface = NewBasicWeightShippingModuleView()
            this.ShippingVector.Add(this.DefaultShippingMethodInterface)
            this.ShippingVector.Add(NewNoShippingModuleView())
        }
    return this
}


// throws Exception
func (this *ShippingMethodsFactory) GetInstance() IBasicArrayList {
        return this.ShippingVector
}


func (this *ShippingMethodsFactory) GetDefaultInstance() ShippingInterface {
        return this.DefaultShippingMethodInterface
}

