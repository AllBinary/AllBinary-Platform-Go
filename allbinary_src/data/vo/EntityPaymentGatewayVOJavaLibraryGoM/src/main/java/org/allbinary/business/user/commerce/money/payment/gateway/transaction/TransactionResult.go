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

type TransactionResult struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    KeyValuePairs *HashMap
    Result string
}

// Implements TransactionResultInterface


func (this *TransactionResult) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewTransactionResult(result string) *TransactionResult {
    this := &TransactionResult{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var exceptionMessage string = "TransactionResult Constructor Failed"; _ = exceptionMessage
            this.LogUtil.PutF(exceptionMessage, this, this.CommonStrings.CONSTRUCTOR)
                }
            }()
            this.Result = result
            var tokenizer *Tokenizer = NewTokenizer(CommonSeps_GetInstance().AMPERSAND); _ = tokenizer
            var tokens *BasicArrayList = tokenizer.GetTokensFromString(result, NewBasicArrayListD()); _ = tokens
            var tokenizer2 *Tokenizer = NewTokenizer(CommonSeps_GetInstance().EQUALS); _ = tokenizer2
            var hashtable *ABHashtable[any, any] = tokenizer2.GetTokens(tokens); _ = hashtable
            this.KeyValuePairs.PutAll(hashtable)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransactionResult(keyValuePairs *HashMap) *TransactionResult {
    this := &TransactionResult{}
    this.initFields()
        keyValuePairs = keyValuePairs
        this.Result = "Needs to have the hashmap converted to string for other constructor"
    return this
}


func (this *TransactionResult) ToString() string {
        return this.Result
}

func (this *TransactionResult) String() string {
    return this.ToString()
}


func (this *TransactionResult) ToHashMap() *HashMap {
        return this.KeyValuePairs
}


func (this *TransactionResult) GetAllPossibleKeys() *BasicArrayList {
        var paymentTransactionKeysFactory *PaymentTransactionKeysFactory = PaymentTransactionKeysFactory_GetInstance(); _ = paymentTransactionKeysFactory
        var verisignResponseKeys *BasicArrayList = NewBasicArrayListD(); _ = verisignResponseKeys
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.RESULT))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.PNREF))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.RESPMSG))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.AUTHCODE))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.AVSADDR))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.AVSZIP))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.ORIGRESULT))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.STATUS))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.FRAUDCODE))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.FRAUDMSG))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.ERRCODE))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.SCORE))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.REASON1))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.REASON2))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.REASON3))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION1))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION2))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION3))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION4))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION5))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION6))
        verisignResponseKeys.Add(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION7))
        return verisignResponseKeys
}


func (this *TransactionResult) GetValues() *BasicArrayList {
        var keyVector *BasicArrayList = this.GetAllPossibleKeys(); _ = keyVector
        var valueVector *BasicArrayList = NewBasicArrayListD(); _ = valueVector
        var size int = keyVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var key string = any(keyVector.Get(i)).(string); _ = key
            if this.KeyValuePairs.ContainsKey(key) {
                valueVector.Add(this.KeyValuePairs.Get(key))
            } else {
                var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
                valueVector.Add(stringUtil.EMPTY_STRING)
            }
        }
        return valueVector
}


func (this *TransactionResult) GetPnRef() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().PNREF))).(string)
}


func (this *TransactionResult) GetRespMsg() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().RESPMSG))).(string)
}


func (this *TransactionResult) GetResult() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().RESULT))).(string)
}


func (this *TransactionResult) GetScore() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().SCORE))).(string)
}


func (this *TransactionResult) GetStatus() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().STATUS))).(string)
}


func (this *TransactionResult) GetFraudCode() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().FRAUDCODE))).(string)
}


func (this *TransactionResult) GetFraudMsg() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().FRAUDMSG))).(string)
}


func (this *TransactionResult) GetOrigResult() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().ORIGRESULT))).(string)
}


func (this *TransactionResult) GetAvsAddr() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().AVSADDR))).(string)
}


func (this *TransactionResult) GetAvsZip() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().AVSZIP))).(string)
}


func (this *TransactionResult) GetErrCode() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().ERRCODE))).(string)
}


func (this *TransactionResult) GetException1() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION1))).(string)
}


func (this *TransactionResult) GetException2() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION2))).(string)
}


func (this *TransactionResult) GetException3() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION3))).(string)
}


func (this *TransactionResult) GetException4() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION4))).(string)
}


func (this *TransactionResult) GetException5() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION5))).(string)
}


func (this *TransactionResult) GetException6() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION6))).(string)
}


func (this *TransactionResult) GetException7() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().EXCEPTION7))).(string)
}


func (this *TransactionResult) GetReason1() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().REASON1))).(string)
}


func (this *TransactionResult) GetReason2() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().REASON2))).(string)
}


func (this *TransactionResult) GetReason3() string {
        return any(this.KeyValuePairs.Get(fmt.Sprint(PaymentTransactionKeysFactory_GetInstance().REASON3))).(string)
}

