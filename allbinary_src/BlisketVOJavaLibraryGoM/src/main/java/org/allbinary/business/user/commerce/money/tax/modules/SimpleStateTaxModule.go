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

type SimpleStateTaxModule struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    StreetAddress *StreetAddress
    StoreFrontInterface StoreFrontInterface
    ALL string
    STATETAXPATH string
    STATETAXFILE string
    Document *Document
}

// Implements TaxModuleInterface


func (this *SimpleStateTaxModule) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.ALL = CommonPhoneStrings_GetInstance().STAR
    this.STATETAXPATH = "/generic/taxes/"
    this.STATETAXFILE = "stateTaxation.xml"
    
}


func NewSimpleStateTaxModule() *SimpleStateTaxModule {
    this := &SimpleStateTaxModule{}
    this.initFields()
    return this
}


// throws Exception
func (this *SimpleStateTaxModule) GetTaxRate(streetAddress *StreetAddress, storeFrontInterface StoreFrontInterface) float32 {
        this.StreetAddress = streetAddress
        this.StoreFrontInterface = storeFrontInterface
        var fileAbPath IAbPath = NewAbPath(fmt.Sprint(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH) + this.StoreFrontInterface.GetName() + this.STATETAXPATH, this.STATETAXFILE); _ = fileAbPath
        var data string = NewCryptFileReader(TransformInfoObjectConfigData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(fileAbPath); _ = data
        this.Document = DomDocumentHelper_Create(data)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAX) {
            this.LogUtil.PutF("Tax Doc: " + XmlDocumentHelper_ToString(this.Document), this, "getTaxRate")
        }
        var state string = this.StreetAddress.GetState(); _ = state
        var country string = this.StreetAddress.GetCountry(); _ = country
        state = strings.ToUpper(state)
        var nodeList *NodeList = this.Document.GetElementsByTagName(TaxData_NAME); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var node *Node = nodeList.Item(index); _ = node
            var taxNodeChildren *NodeList = node.GetChildNodes(); _ = taxNodeChildren
            var streetAddressNode *Node = DomSearchHelper_GetNode(StreetAddressData_NAME, taxNodeChildren); _ = streetAddressNode
            var taxableStreetAddress *StreetAddress = NewStreetAddress(streetAddressNode); _ = taxableStreetAddress
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if !stringValidationUtil.IsEmpty(country) && !stringValidationUtil.IsEmpty(taxableStreetAddress.GetCountry()) && strings.Compare(country, taxableStreetAddress.GetCountry()) == 0 {
                if state != "" && taxableStreetAddress.GetState() != nil && strings.Compare(state, taxableStreetAddress.GetState()) == 0 {
                    var rateNode *Node = DomSearchHelper_GetNode(TaxData_RATE, taxNodeChildren); _ = rateNode
                    var taxRate string = DomNodeHelper_GetTextNodeValue(rateNode); _ = taxRate
                    var taxRateFloat float32 = float32(float64(float32(float32(taxRate))) * .01); _ = taxRateFloat
                    return taxRateFloat
                }
            }
        }
        return float32(0)
}


func (this *SimpleStateTaxModule) IsValid(streetAddress *StreetAddress, storeFrontInterface StoreFrontInterface) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.StreetAddress = streetAddress
            this.StoreFrontInterface = storeFrontInterface
            var fileAbPath IAbPath = NewAbPath(fmt.Sprint(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH) + this.StoreFrontInterface.GetName() + this.STATETAXPATH, this.STATETAXFILE); _ = fileAbPath
            var data string = NewCryptFileReader(TransformInfoObjectConfigData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(fileAbPath); _ = data
            this.Document = DomDocumentHelper_Create(data)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAX) {
                this.LogUtil.PutF("Tax Doc: " + XmlDocumentHelper_ToString(this.Document), this, this.CommonStrings.IS_VALID)
            }
            var country string = this.StreetAddress.GetCountry(); _ = country
            var rootNode *Node = this.Document.GetElementsByTagName(TransformInfoObjectConfigData_GetInstance().NAME).Item(0); _ = rootNode
            var nodeList *NodeList = this.Document.GetElementsByTagName(TaxData_NAME); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var node *Node = nodeList.Item(index); _ = node
                var taxNodeChildren *NodeList = node.GetChildNodes(); _ = taxNodeChildren
                var streetAddressNode *Node = DomSearchHelper_GetNode(StreetAddressData_NAME, taxNodeChildren); _ = streetAddressNode
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAX) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("StreetAddress Node: ")
                    stringBuffer.Append(streetAddressNode.GetNodeName())
                    stringBuffer.Append(" Value: ")
                    stringBuffer.Append(streetAddressNode.GetNodeValue())
                    stringBuffer.Append(" Number Of Children: ")
                    stringBuffer.Appendint(int(streetAddressNode.GetChildNodes().GetLength()))
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                }
                var taxableStreetAddress *StreetAddress = NewStreetAddress(streetAddressNode); _ = taxableStreetAddress
                var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
                if !stringValidationUtil.IsEmpty(country) && !stringValidationUtil.IsEmpty(taxableStreetAddress.GetCountry()) && strings.Compare(taxableStreetAddress.GetCountry(), this.ALL) == 0 {
                    tryControl1, tryResult1 = 1, true; return
                }
                if !stringValidationUtil.IsEmpty(country) && !stringValidationUtil.IsEmpty(taxableStreetAddress.GetCountry()) && strings.Compare(country, taxableStreetAddress.GetCountry()) == 0 {
                    var state string = this.StreetAddress.GetState(); _ = state
                    if state != "" {
                        tryControl1, tryResult1 = 1, true; return
                    }
                }
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

