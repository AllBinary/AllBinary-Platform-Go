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

type PaymentTransactionKeysFactory struct {
    TRXTYPE *PaymentTransactionKeys
    TENDER *PaymentTransactionKeys
    ACCT *PaymentTransactionKeys
    EXPDATE *PaymentTransactionKeys
    AMT *PaymentTransactionKeys
    AUTHCODE *PaymentTransactionKeys
    MICR *PaymentTransactionKeys
    CHECKNUM *PaymentTransactionKeys
    NAME *PaymentTransactionKeys
    STREET *PaymentTransactionKeys
    CITY *PaymentTransactionKeys
    STATE *PaymentTransactionKeys
    ZIP *PaymentTransactionKeys
    DL *PaymentTransactionKeys
    EMAIL *PaymentTransactionKeys
    COMMENT1 *PaymentTransactionKeys
    COMMENT2 *PaymentTransactionKeys
    ORIGID *PaymentTransactionKeys
    PONUM *PaymentTransactionKeys
    DESC *PaymentTransactionKeys
    DESC1TO4 *PaymentTransactionKeys
    INVNUM *PaymentTransactionKeys
    SHIPTOZIP *PaymentTransactionKeys
    SWIPE *PaymentTransactionKeys
    TAXAMT *PaymentTransactionKeys
    COMMCARD *PaymentTransactionKeys
    DUTYAMT *PaymentTransactionKeys
    FRIEGHTAMT *PaymentTransactionKeys
    ORDERDATE *PaymentTransactionKeys
    TAXEXEMPT *PaymentTransactionKeys
    COUNTRYCODE *PaymentTransactionKeys
    CUSTCODE *PaymentTransactionKeys
    CVV2 *PaymentTransactionKeys
    ABA *PaymentTransactionKeys
    ACCTTYPE *PaymentTransactionKeys
    DISCOUNT *PaymentTransactionKeys
    FIRSTNAME *PaymentTransactionKeys
    LASTNAME *PaymentTransactionKeys
    SHIPFROMZIP *PaymentTransactionKeys
    PRENOTE *PaymentTransactionKeys
    CHKTYPE *PaymentTransactionKeys
    DOB *PaymentTransactionKeys
    PHONENUM *PaymentTransactionKeys
    SS *PaymentTransactionKeys
    COMPANYNAME *PaymentTransactionKeys
    COUNTRY *PaymentTransactionKeys
    FREIGHTAMT *PaymentTransactionKeys
    SHIPTOCITY *PaymentTransactionKeys
    SHIPTOFIRSTNAME *PaymentTransactionKeys
    SHIPTOLASTNAME *PaymentTransactionKeys
    SHIPTOSTATE *PaymentTransactionKeys
    SHIPTOSTREET *PaymentTransactionKeys
    SPECIAL1 *PaymentTransactionKeys
    SPECIAL2 *PaymentTransactionKeys
    SPECIAL3 *PaymentTransactionKeys
    PARTNER *PaymentTransactionKeys
    VENDOR *PaymentTransactionKeys
    USER *PaymentTransactionKeys
    PWD *PaymentTransactionKeys
    CHKNUM *PaymentTransactionKeys
    AVSADDR *PaymentTransactionKeys
    AVSZIP *PaymentTransactionKeys
    PNREF *PaymentTransactionKeys
    RESPMSG *PaymentTransactionKeys
    RESULT *PaymentTransactionKeys
    ORIGRESULT *PaymentTransactionKeys
    STATUS *PaymentTransactionKeys
    FRAUDCODE *PaymentTransactionKeys
    FRAUDMSG *PaymentTransactionKeys
    ERRCODE *PaymentTransactionKeys
    SCORE *PaymentTransactionKeys
    REASON1 *PaymentTransactionKeys
    REASON2 *PaymentTransactionKeys
    REASON3 *PaymentTransactionKeys
    EXCEPTION1 *PaymentTransactionKeys
    EXCEPTION2 *PaymentTransactionKeys
    EXCEPTION3 *PaymentTransactionKeys
    EXCEPTION4 *PaymentTransactionKeys
    EXCEPTION5 *PaymentTransactionKeys
    EXCEPTION6 *PaymentTransactionKeys
    EXCEPTION7 *PaymentTransactionKeys
}

var PaymentTransactionKeysFactory_Instance *PaymentTransactionKeysFactory = NewPaymentTransactionKeysFactory()

func (this *PaymentTransactionKeysFactory) initFields() {
    this.TRXTYPE = NewPaymentTransactionKeys("TRXTYPE")
    this.TENDER = NewPaymentTransactionKeys("TENDER")
    this.ACCT = NewPaymentTransactionKeys("ACCT")
    this.EXPDATE = NewPaymentTransactionKeys("EXPDATE")
    this.AMT = NewPaymentTransactionKeys("AMT")
    this.AUTHCODE = NewPaymentTransactionKeys("AUTHCODE")
    this.MICR = NewPaymentTransactionKeys("MICR")
    this.CHECKNUM = NewPaymentTransactionKeys("CHECKNUM")
    this.NAME = NewPaymentTransactionKeys("NAME")
    this.STREET = NewPaymentTransactionKeys("STREET")
    this.CITY = NewPaymentTransactionKeys("CITY")
    this.STATE = NewPaymentTransactionKeys("STATE")
    this.ZIP = NewPaymentTransactionKeys("ZIP")
    this.DL = NewPaymentTransactionKeys("DL")
    this.EMAIL = NewPaymentTransactionKeys("EMAIL")
    this.COMMENT1 = NewPaymentTransactionKeys("COMMENT1")
    this.COMMENT2 = NewPaymentTransactionKeys("COMMENT2")
    this.ORIGID = NewPaymentTransactionKeys("ORIGID")
    this.PONUM = NewPaymentTransactionKeys("PONUM")
    this.DESC = NewPaymentTransactionKeys("PaymentTransactionDESC")
    this.DESC1TO4 = NewPaymentTransactionKeys("DESC1TO4")
    this.INVNUM = NewPaymentTransactionKeys("INVNUM")
    this.SHIPTOZIP = NewPaymentTransactionKeys("SHIPTOZIP")
    this.SWIPE = NewPaymentTransactionKeys("SWIPE")
    this.TAXAMT = NewPaymentTransactionKeys("TAXAMT")
    this.COMMCARD = NewPaymentTransactionKeys("COMMCARD")
    this.DUTYAMT = NewPaymentTransactionKeys("DUTYAMT")
    this.FRIEGHTAMT = NewPaymentTransactionKeys("FRIEGHTAMT")
    this.ORDERDATE = NewPaymentTransactionKeys("ORDERDATE")
    this.TAXEXEMPT = NewPaymentTransactionKeys("TAXEXEMPT")
    this.COUNTRYCODE = NewPaymentTransactionKeys("COUNTRYCODE")
    this.CUSTCODE = NewPaymentTransactionKeys("CUSTCODE")
    this.CVV2 = NewPaymentTransactionKeys("CVV2")
    this.ABA = NewPaymentTransactionKeys("ABA")
    this.ACCTTYPE = NewPaymentTransactionKeys("ACCTTYPE")
    this.DISCOUNT = NewPaymentTransactionKeys("DISCOUNT")
    this.FIRSTNAME = NewPaymentTransactionKeys("FIRSTNAME")
    this.LASTNAME = NewPaymentTransactionKeys("LASNAME")
    this.SHIPFROMZIP = NewPaymentTransactionKeys("SHIPFROMZIP")
    this.PRENOTE = NewPaymentTransactionKeys("PRENOTE")
    this.CHKTYPE = NewPaymentTransactionKeys("CHKTYPE")
    this.DOB = NewPaymentTransactionKeys("DOB")
    this.PHONENUM = NewPaymentTransactionKeys("PHONENUM")
    this.SS = NewPaymentTransactionKeys("SS")
    this.COMPANYNAME = NewPaymentTransactionKeys("COMPANYNAME")
    this.COUNTRY = NewPaymentTransactionKeys("COUNTRY")
    this.FREIGHTAMT = NewPaymentTransactionKeys("FREIGHTAMT")
    this.SHIPTOCITY = NewPaymentTransactionKeys("SHIPTOCITY")
    this.SHIPTOFIRSTNAME = NewPaymentTransactionKeys("SHIPTOFIRSTNAME")
    this.SHIPTOLASTNAME = NewPaymentTransactionKeys("SHIPTOLASTNAME")
    this.SHIPTOSTATE = NewPaymentTransactionKeys("SHIPTOSTATE")
    this.SHIPTOSTREET = NewPaymentTransactionKeys("SHIPTOSTREET")
    this.SPECIAL1 = NewPaymentTransactionKeys("SPECIAL1")
    this.SPECIAL2 = NewPaymentTransactionKeys("SPECIAL2")
    this.SPECIAL3 = NewPaymentTransactionKeys("SPECIAL3")
    this.PARTNER = NewPaymentTransactionKeys("PARTNER")
    this.VENDOR = NewPaymentTransactionKeys("VENDOR")
    this.USER = NewPaymentTransactionKeys("USER")
    this.PWD = NewPaymentTransactionKeys("PWD")
    this.CHKNUM = NewPaymentTransactionKeys("CHKNUM")
    this.AVSADDR = NewPaymentTransactionKeys("AVSADDR")
    this.AVSZIP = NewPaymentTransactionKeys("AVSZIP")
    this.PNREF = NewPaymentTransactionKeys("PNREF")
    this.RESPMSG = NewPaymentTransactionKeys("RESPMSG")
    this.RESULT = NewPaymentTransactionKeys("RESULT")
    this.ORIGRESULT = NewPaymentTransactionKeys("ORIGRESULT")
    this.STATUS = NewPaymentTransactionKeys("STATUS")
    this.FRAUDCODE = NewPaymentTransactionKeys("FRAUDCODE")
    this.FRAUDMSG = NewPaymentTransactionKeys("FRAUDMSG")
    this.ERRCODE = NewPaymentTransactionKeys("ERRCODE")
    this.SCORE = NewPaymentTransactionKeys("SCORE")
    this.REASON1 = NewPaymentTransactionKeys("REASON1")
    this.REASON2 = NewPaymentTransactionKeys("REASON2")
    this.REASON3 = NewPaymentTransactionKeys("REASON3")
    this.EXCEPTION1 = NewPaymentTransactionKeys("EXCEPTION1")
    this.EXCEPTION2 = NewPaymentTransactionKeys("EXCEPTION2")
    this.EXCEPTION3 = NewPaymentTransactionKeys("EXCEPTION3")
    this.EXCEPTION4 = NewPaymentTransactionKeys("EXCEPTION4")
    this.EXCEPTION5 = NewPaymentTransactionKeys("EXCEPTION5")
    this.EXCEPTION6 = NewPaymentTransactionKeys("EXCEPTION6")
    this.EXCEPTION7 = NewPaymentTransactionKeys("EXCEPTION7")
}

func NewPaymentTransactionKeysFactory() *PaymentTransactionKeysFactory {
    this := &PaymentTransactionKeysFactory{}
    this.initFields()
    return this
}


func PaymentTransactionKeysFactory_GetInstance() *PaymentTransactionKeysFactory {
        return PaymentTransactionKeysFactory_Instance
}

