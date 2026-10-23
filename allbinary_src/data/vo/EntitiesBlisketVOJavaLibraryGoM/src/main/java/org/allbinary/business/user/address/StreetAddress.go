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

type StreetAddress struct {
    LogUtil *LogUtil
    Id string
    Name string
    Street string
    City string
    State string
    Code string
    Country string
    IsDefault bool
}


func (this *StreetAddress) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    this.IsDefault = false
}




// Java overloads of new StreetAddress
func NewStreetAddress(args ...any) *StreetAddress {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewStreetAddress_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewStreetAddress_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*StreetAddress](args[0], pass) {
            return NewStreetAddress_3(Java_Arg[*StreetAddress](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return NewStreetAddress_4(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 6 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) && Java_IsArg[string](args[5], pass) {
            return NewStreetAddress_5(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]), Java_Arg[string](args[5]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewStreetAddress_6(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new StreetAddress for the arguments")
}
func NewStreetAddress_1() *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
    return this
}


func NewStreetAddress_2(node *Node) *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
        var nameNode *Node = DomSearchHelper_GetNode(StreetAddressData_NAME, node.GetChildNodes()); _ = nameNode
        this.Name = DomNodeHelper_GetTextNodeValue(nameNode)
        var streetNode *Node = DomSearchHelper_GetNode(StreetAddressData_STREET, node.GetChildNodes()); _ = streetNode
        this.Street = DomNodeHelper_GetTextNodeValue(streetNode)
        var cityNode *Node = DomSearchHelper_GetNode(StreetAddressData_CITY, node.GetChildNodes()); _ = cityNode
        this.City = DomNodeHelper_GetTextNodeValue(cityNode)
        var stateNode *Node = DomSearchHelper_GetNode(StreetAddressData_STATE, node.GetChildNodes()); _ = stateNode
        this.State = DomNodeHelper_GetTextNodeValue(stateNode)
        var codeNode *Node = DomSearchHelper_GetNode(StreetAddressData_CODE, node.GetChildNodes()); _ = codeNode
        this.Code = DomNodeHelper_GetTextNodeValue(codeNode)
        var countryNode *Node = DomSearchHelper_GetNode(StreetAddressData_COUNTRY, node.GetChildNodes()); _ = countryNode
        this.Country = DomNodeHelper_GetTextNodeValue(countryNode)
        this.Log()
    return this
}


func NewStreetAddress_3(streetAddress *StreetAddress) *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
        this.Id = streetAddress.GetId()
        this.Name = streetAddress.GetName()
        this.Street = streetAddress.GetStreet()
        this.City = streetAddress.GetCity()
        this.State = streetAddress.GetState()
        this.Code = streetAddress.GetCode()
        this.Country = streetAddress.GetCountry()
        this.IsDefault = streetAddress.IsDefault()
        this.Log()
    return this
}


func NewStreetAddress_4(request *HttpServletRequest) *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
        this.SetId(request.GetParameter(StreetAddressData_ID))
        this.SetName(request.GetParameter(StreetAddressData_NAME))
        this.SetStreet(request.GetParameter(StreetAddressData_STREET))
        this.SetCity(request.GetParameter(StreetAddressData_CITY))
        this.SetState(request.GetParameter(StreetAddressData_STATE))
        this.SetCode(request.GetParameter(StreetAddressData_CODE))
        this.SetCountry(request.GetParameter(StreetAddressData_COUNTRY))
        this.Log()
    return this
}


func NewStreetAddress_5(name string, street string, city string, state string, code string, country string) *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
        this.Name = name
        this.Street = street
        this.City = city
        this.State = state
        this.Code = code
        this.Country = country
        this.Log()
    return this
}


func NewStreetAddress_6(address *HashMap) *StreetAddress {
    this := &StreetAddress{}
    this.initFields()
        this.Id = any(address.Get(StreetAddressData_ID)).(string)
        this.Name = any(address.Get(StreetAddressData_NAME)).(string)
        this.Street = any(address.Get(StreetAddressData_STREET)).(string)
        this.City = any(address.Get(StreetAddressData_CITY)).(string)
        this.State = any(address.Get(StreetAddressData_STATE)).(string)
        this.Code = any(address.Get(StreetAddressData_CODE)).(string)
        this.Country = any(address.Get(StreetAddressData_COUNTRY)).(string)
        var def string = any(address.Get(StreetAddressData_DEFAULT)).(string); _ = def
        if def != "" && strings.Compare(def, StreetAddressData_DEFAULT) == 0 {
            this.IsDefault = true
        } else {
            this.IsDefault = false
        }
        this.Log()
    return this
}


func (this *StreetAddress) Log() {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("Created Address: \n" + fmt.Sprint(this.ToHashMap()), this, "log")
        }
}


func (this *StreetAddress) SetId(id string) {
        this.Id = id
}


func (this *StreetAddress) SetName(name string) {
        this.Name = name
}


func (this *StreetAddress) SetStreet(value string) {
        this.Street = value
}


func (this *StreetAddress) SetCity(value string) {
        this.City = value
}


func (this *StreetAddress) SetState(value string) {
        this.State = value
}


func (this *StreetAddress) SetCode(value string) {
        this.Code = value
}


func (this *StreetAddress) SetCountry(value string) {
        this.Country = value
}


func (this *StreetAddress) GetName() string {
        return this.Name
}


func (this *StreetAddress) GetStreet() string {
        return this.Street
}


func (this *StreetAddress) GetCity() string {
        return this.City
}


func (this *StreetAddress) GetState() string {
        return this.State
}


func (this *StreetAddress) GetCode() string {
        return this.Code
}


func (this *StreetAddress) GetCountry() string {
        return this.Country
}


func (this *StreetAddress) GetId() string {
        return this.Id
}


func (this *StreetAddress) IsDefault() bool {
        return this.IsDefault
}


func (this *StreetAddress) IsEmpty() bool {
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        if this.GetName() == "" && strings.Compare(this.GetName(), EMPTY_STRING) == 0 && this.GetStreet() == nil && strings.Compare(this.GetStreet(), EMPTY_STRING) == 0 && this.GetCity() == nil && strings.Compare(this.GetCity(), EMPTY_STRING) == 0 && this.GetState() == nil && strings.Compare(this.GetState(), EMPTY_STRING) == 0 && this.GetCode() == nil && strings.Compare(this.GetCode(), EMPTY_STRING) == 0 && this.GetCountry() == nil && strings.Compare(this.GetCountry(), EMPTY_STRING) == 0 {
            return true
        }
        return false
}


func (this *StreetAddress) IsValid() bool {
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        if this.GetName() == "" || strings.Compare(this.GetName(), EMPTY_STRING) == 0 || this.GetStreet() == nil || strings.Compare(this.GetStreet(), EMPTY_STRING) == 0 || this.GetCity() == nil || strings.Compare(this.GetCity(), EMPTY_STRING) == 0 || this.GetState() == nil || strings.Compare(this.GetState(), EMPTY_STRING) == 0 || this.GetCode() == nil || strings.Compare(this.GetCode(), EMPTY_STRING) == 0 || this.GetCountry() == nil || strings.Compare(this.GetCountry(), EMPTY_STRING) == 0 {
            return false
        }
        return true
}


func (this *StreetAddress) ValidationInfo() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Address Failed To Validate - All fields must contain data.<br></br>")
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.GetName(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid name")
            stringBuffer.Append("<br></br>")
        }
        if !stringValidationUtil.IsValidRequired(this.GetStreet(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid street")
            stringBuffer.Append("<br></br>")
        }
        if !stringValidationUtil.IsValidRequired(this.GetCity(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid city")
            stringBuffer.Append("<br></br>")
        }
        if !stringValidationUtil.IsValidRequired(this.GetState(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid state")
            stringBuffer.Append("<br></br>")
        }
        if !stringValidationUtil.IsValidRequired(this.GetCode(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid zipcode")
            stringBuffer.Append("<br></br>")
        }
        if !stringValidationUtil.IsValidRequired(this.GetCountry(), StreetAddressData_MIN, StreetAddressData_MAX) {
            stringBuffer.Append("Please enter a valid country")
            stringBuffer.Append("<br></br>")
        }
        return fmt.Sprint(stringBuffer)
}


func (this *StreetAddress) ToValidationInfoDoc() *Document {
        return nil
}


func (this *StreetAddress) ToValidationInfoNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "toValidationInfoNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var node *Node = document.CreateElement(StreetAddressData_FORM); _ = node
            node.AppendChild(this.ToXmlNode(document))
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StreetAddress) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(StreetAddressData_ID, this.Id)
        hashMap.Put(StreetAddressData_NAME, this.Name)
        hashMap.Put(StreetAddressData_STREET, this.Street)
        hashMap.Put(StreetAddressData_CITY, this.City)
        hashMap.Put(StreetAddressData_STATE, this.State)
        hashMap.Put(StreetAddressData_CODE, this.Code)
        hashMap.Put(StreetAddressData_COUNTRY, this.Country)
        return hashMap
}


// throws Exception
func (this *StreetAddress) ToXmlNode(document *Document) *Node {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        var keySet Set = hashMap.KeySet(); _ = keySet
        var keyArray []any = keySet.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        var node *Node = document.CreateElement(StreetAddressData_ADDRESS); _ = node
        for i := int(0); i < size; i++ {
            var name string = any(keyArray[i]).(string); _ = name
            var value string = stringUtil.GetNonNull(any(hashMap.Get(name)).(string)); _ = value
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        return node
}

