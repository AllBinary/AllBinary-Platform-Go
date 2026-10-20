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

type Payment struct {
    Name string
    Expiration string
    Encryption string
    Number string
    Id string
    Type string
    TenderType string
    TransactionType string
    Aba string
    Account string
    AccountType string
    AuthorizationCode string
    CheckNumber string
    CheckType string
    DriversLicense string
    MagneticInkCheckReader string
    IsDefault bool
}

// Implements PaymentInterface, DomNodeInterface


func (this *Payment) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    this.IsDefault = false
}


func NewPayment(payment *HashMap) *Payment {
    this := &Payment{}
    this.initFields()
        this.Name = any(payment.Get(PaymentData_NAME)).(string)
        this.Type = any(payment.Get(PaymentData_TYPE)).(string)
        this.Expiration = any(payment.Get(PaymentData_EXPIRATION)).(string)
        this.Encryption = any(payment.Get(EntryData_GetInstance().ENCRYPTION)).(string)
        this.Number = any(NewSuperCrypt(int(int(this.Encryption))).Decrypt(any(payment.Get(PaymentData_NUMBER)).(string))).(string)
        this.Id = any(payment.Get(PaymentData_ID)).(string)
        var def string = any(payment.Get(EntryData_GetInstance().DEFAULT)).(string); _ = def
        if def != "" && strings.Compare(def, EntryData_GetInstance().DEFAULT) == 0 {
            this.IsDefault = true
        } else {
            this.IsDefault = false
        }
        this.TenderType = any(payment.Get(PaymentData_TENDERTYPE)).(string)
        this.TransactionType = any(payment.Get(PaymentData_TRANSACTIONTYPE)).(string)
        this.Aba = any(payment.Get(PaymentData_ABA)).(string)
        this.Account = any(payment.Get(PaymentData_ACCOUNT)).(string)
        this.AccountType = any(payment.Get(PaymentData_ACCOUNTTYPE)).(string)
        this.AuthorizationCode = any(payment.Get(PaymentData_AUTHORIZATIONCODE)).(string)
        this.CheckNumber = any(payment.Get(PaymentData_CHECKNUMBER)).(string)
        this.CheckType = any(payment.Get(PaymentData_CHECKTYPE)).(string)
        this.DriversLicense = any(payment.Get(PaymentData_DRIVERSLICENSE)).(string)
        this.MagneticInkCheckReader = any(payment.Get(PaymentData_MAGNETICINKCHECKREADER)).(string)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewPayment(request *ServletRequest) *Payment {
    this := &Payment{}
    this.initFields()
        this.Name = request.GetParameter(PaymentData_NAME)
        this.Type = request.GetParameter(PaymentData_TYPE)
        this.Expiration = request.GetParameter(PaymentData_EXPIRATION)
        this.Number = request.GetParameter(PaymentData_NUMBER)
        this.TenderType = request.GetParameter(PaymentData_TENDERTYPE)
        this.TransactionType = request.GetParameter(PaymentData_TRANSACTIONTYPE)
        this.Aba = request.GetParameter(PaymentData_ABA)
        this.Account = request.GetParameter(PaymentData_ACCOUNT)
        this.AccountType = request.GetParameter(PaymentData_ACCOUNTTYPE)
        this.AuthorizationCode = request.GetParameter(PaymentData_AUTHORIZATIONCODE)
        this.CheckNumber = request.GetParameter(PaymentData_CHECKNUMBER)
        this.CheckType = request.GetParameter(PaymentData_CHECKTYPE)
        this.DriversLicense = request.GetParameter(PaymentData_DRIVERSLICENSE)
        this.MagneticInkCheckReader = request.GetParameter(PaymentData_MAGNETICINKCHECKREADER)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewPayment() *Payment {
    this := &Payment{}
    this.initFields()
    return this
}


func (this *Payment) IsValid() bool {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var aFalse bool = false; _ = aFalse
        if this.Name == "" || strings.Compare(this.Name, stringUtil.EMPTY_STRING) == 0 {
            return aFalse
        }
        if this.Type == "" || strings.Compare(this.Type, stringUtil.EMPTY_STRING) == 0 {
            return aFalse
        }
        if this.Expiration == "" || strings.Compare(this.Expiration, stringUtil.EMPTY_STRING) == 0 {
            return aFalse
        }
        if this.Number == "" || strings.Compare(this.Number, stringUtil.EMPTY_STRING) == 0 {
            return aFalse
        }
        if len(this.Number) > org_Allbinary.Business.User.Commerce.Money.Payment.Types.CreditCardType.MAXLEN {
            return aFalse
        }
        return true
}


func (this *Payment) SetTenderType(value string) {
        this.TenderType = value
}


func (this *Payment) SetTransactionType(value string) {
        this.TransactionType = value
}


func (this *Payment) SetAuthorizationCode(value string) {
        this.AuthorizationCode = value
}


func (this *Payment) SetCheckType(value string) {
        this.CheckType = value
}


func (this *Payment) SetCheckNumber(value string) {
        this.CheckNumber = value
}


func (this *Payment) SetDriversLicense(value string) {
        this.DriversLicense = value
}


func (this *Payment) SetMagneticInkCheckReader(value string) {
        this.MagneticInkCheckReader = value
}


func (this *Payment) SetAba(value string) {
        this.Aba = value
}


func (this *Payment) SetAccountType(value string) {
        this.AccountType = value
}


func (this *Payment) SetName(value string) {
        this.Name = value
}


func (this *Payment) SetType(value string) {
        this.Type = value
}


func (this *Payment) SetExpiration(value string) {
        this.Expiration = value
}


func (this *Payment) SetNumber(value string) {
        this.Number = value
}


func (this *Payment) GetId() string {
        return this.Id
}


func (this *Payment) GetName() string {
        return this.Name
}


func (this *Payment) GetType() string {
        return this.Type
}


func (this *Payment) GetExpiration() string {
        return this.Expiration
}


func (this *Payment) GetEncryption() string {
        return this.Encryption
}


func (this *Payment) GetNumber() string {
        return this.Number
}


func (this *Payment) GetTenderType() string {
        return this.TenderType
}


func (this *Payment) GetTransactionType() string {
        return this.TransactionType
}


func (this *Payment) GetAba() string {
        return this.Aba
}


func (this *Payment) GetAccount() string {
        return this.Account
}


func (this *Payment) GetAccountType() string {
        return this.AccountType
}


func (this *Payment) GetAuthorizationCode() string {
        return this.AuthorizationCode
}


func (this *Payment) GetCheckNumber() string {
        return this.CheckNumber
}


func (this *Payment) GetCheckType() string {
        return this.CheckType
}


func (this *Payment) GetDriversLicense() string {
        return this.DriversLicense
}


func (this *Payment) GetMagneticInkCheckReader() string {
        return this.MagneticInkCheckReader
}


func (this *Payment) GetLastFour() string {
        if len(this.Number) > PaymentData_MIN && len(this.Number) <= PaymentData_MAX {
            return this.Number[len(this.Number) - 4:]
        } else {
            return "Invalid #"
        }
}


func (this *Payment) IsDefault() bool {
        return this.IsDefault
}


func (this *Payment) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(PaymentData_NAME, this.Name)
        hashMap.Put(PaymentData_TYPE, this.Type)
        hashMap.Put(PaymentData_EXPIRATION, this.Expiration)
        hashMap.Put(PaymentData_NUMBER, this.GetLastFour())
        hashMap.Put(PaymentData_TENDERTYPE, this.TenderType)
        hashMap.Put(PaymentData_TRANSACTIONTYPE, this.TransactionType)
        hashMap.Put(PaymentData_ABA, this.Aba)
        hashMap.Put(PaymentData_ACCOUNT, this.Account)
        hashMap.Put(PaymentData_ACCOUNTTYPE, this.AccountType)
        hashMap.Put(PaymentData_AUTHORIZATIONCODE, this.AuthorizationCode)
        hashMap.Put(PaymentData_CHECKNUMBER, this.CheckNumber)
        hashMap.Put(PaymentData_DRIVERSLICENSE, this.DriversLicense)
        hashMap.Put(PaymentData_MAGNETICINKCHECKREADER, this.MagneticInkCheckReader)
        if this.IsDefault {
            hashMap.Put(EntryData_GetInstance().DEFAULT, EntryData_GetInstance().DEFAULT)
        } else {
            hashMap.Put(EntryData_GetInstance().DEFAULT, StringUtil_GetInstance().EMPTY_STRING)
        }
        return hashMap
}


// throws Exception
func (this *Payment) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        var node *Node = ModDomHelper_CreateNameValueNodes(document, PaymentData_PAYMENT, hashMap); _ = node
        return node
}


func (this *Payment) ToXmlDoc() *Document {
        return nil
}


func (this *Payment) View() string {
        return nil
}


func (this *Payment) ToValidationInfoDoc() *Document {
        return nil
}


func (this *Payment) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *Payment) ValidationInfo() string {
        return nil
}

