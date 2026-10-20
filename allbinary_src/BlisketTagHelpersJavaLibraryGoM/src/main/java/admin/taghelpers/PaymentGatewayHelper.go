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

type PaymentGatewayHelper struct {
    BasicTable
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    HttpServletRequest *HttpServletRequest
    Portion *Portion
}


func (this *PaymentGatewayHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewPaymentGatewayHelper(hashMap *HashMap, pageContext *PageContext) *PaymentGatewayHelper {
    this := &PaymentGatewayHelper{}
    this.BasicTable = *NewBasicTable()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
            this.HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest)
            this.Portion = NewPortion(hashMap)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *PaymentGatewayHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "add()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br>" + "Exception: " + fmt.Sprint(e) + "<br>"; return
                }
            }()
            var paymentGatewayInterface PaymentGatewayInterface = NewPaymentGatewayInterfaceFactory().GetInstance(this.HttpServletRequest); _ = paymentGatewayInterface
            paymentGatewayInterface.SetStoreName(this.WeblisketSession.GetStoreName())
            PaymentGatewayEntityFactory_GetInstance().Add(paymentGatewayInterface)
            var success string = "Successfully Added New Payment Gateway"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "insert()")
            }
            tryControl1, tryResult1 = 1, success + "<br/>"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update payment gateway information"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br/>" + "Exception: " + fmt.Sprint(e) + "<br/>"; return
                }
            }()
            var paymentGatewayInterface PaymentGatewayInterface = NewPaymentGatewayInterfaceFactory().GetInstance(this.HttpServletRequest); _ = paymentGatewayInterface
            paymentGatewayInterface.SetStoreName(this.WeblisketSession.GetStoreName())
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Gateway Name: ")
                stringBuffer.Append(paymentGatewayInterface.GetName())
                stringBuffer.Append(" HashMap=")
                stringBuffer.Append(fmt.Sprint(paymentGatewayInterface.ToHashMap(true)))
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "update()")
            }
            PaymentGatewayEntityFactory_GetInstance().Update(paymentGatewayInterface)
            var success string = "Successfully Updated Payment Gateway information"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success + "<br/>"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to remove payment gateway from table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "remove()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br/>" + "Exception: " + fmt.Sprint(e) + "<br>"; return
                }
            }()
            var paymentGatewayInterface PaymentGatewayInterface = NewPaymentGatewayInterfaceFactory().GetInstance(this.HttpServletRequest); _ = paymentGatewayInterface
            paymentGatewayInterface.SetStoreName(this.WeblisketSession.GetStoreName())
            var storeName string = paymentGatewayInterface.GetStoreName(); _ = storeName
            var gatewayName string = paymentGatewayInterface.GetName(); _ = gatewayName
            PaymentGatewayEntityFactory_GetInstance().Remove(storeName, BasicPaymentTypeUtil_GetInstance().Get(gatewayName))
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Successfully Removed payment gateway where store name=")
            stringBuffer.Append(storeName)
            stringBuffer.Append(" and gateway name=")
            stringBuffer.Append(gatewayName)
            var success string = fmt.Sprint(stringBuffer); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "delete()")
            }
            tryControl1, tryResult1 = 1, success + "<br/>"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Drop() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to drop payment transaction result table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.DROP, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = PaymentGatewayEntityFactory_GetInstance().DropTable(); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, this.CommonStrings.DROP)
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Create() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to create payment transaction result table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "create()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = PaymentGatewayEntityFactory_GetInstance().CreateTable(); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "create()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Restore() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to restore backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "restore()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Restore Successful"; _ = success
            var result string = AbSqlTableUtil_GetInstance().RestoreTable(PaymentGatewayEntityFactory_GetInstance(), this.Portion); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "restore()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayHelper) Backup() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to make backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Restore Successful"; _ = success
            var result string = AbSqlTableUtil_GetInstance().BackupTable(PaymentGatewayEntityFactory_GetInstance()); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "backup()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

