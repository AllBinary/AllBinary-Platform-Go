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

type PaymentGateway struct {
    StoreName string
    Name string
    Enable string
    PaymentMethod string
    Mode string
    TestProtocol string
    TestServer string
    TestPort string
    TestPath string
    ServerProtocol string
    Server string
    ServerPort string
    ServerPath string
    Timeout string
    UserName string
    Password string
    ProxyProtocol string
    ProxyServer string
    ProxyPort string
    ProxyPath string
    ProxyUserName string
    ProxyPassword string
    ProxyTimeout string
    Special1 string
    Special2 string
    Special3 string
    Special4 string
    Special5 string
    Special6 string
    Special7 string
    Special8 string
    Special9 string
    LastModified string
    TimeEntered string
}

// Implements PaymentGatewayInterface


func (this *PaymentGateway) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewPaymentGateway(hashMap *HashMap) *PaymentGateway {
    this := &PaymentGateway{}
    this.initFields()
        this.GetFormData(hashMap)
    return this
}


func (this *PaymentGateway) GetFormData(hashMap *HashMap) {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var empty string = stringUtil.EMPTY_STRING; _ = empty
        this.StoreName = any(hashMap.Get(fmt.Sprint(StoreFrontData_GetInstance().NAME))).(string)
        this.Name = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_NAME))).(string)
        this.Enable = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(EntryData_GetInstance().ENABLE))).(string))
        this.PaymentMethod = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(PaymentData_METHOD))).(string))
        this.Mode = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(PaymentGatewayData_MODE))).(string))
        this.TestProtocol = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(PaymentGatewayData_TESTPROTOCOL))).(string))
        this.TestServer = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(PaymentGatewayData_TESTSERVER))).(string))
        this.TestPort = stringUtil.GetNonNull(any(hashMap.Get(fmt.Sprint(PaymentGatewayData_TESTPORT))).(string))
        this.TestPath = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_TESTPATH))).(string)
        if this.TestPath == "" {
            this.TestPath = empty
        }
        this.ServerProtocol = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SERVERPROTOCOL))).(string)
        if this.ServerProtocol == "" {
            this.ServerProtocol = empty
        }
        this.Server = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SERVER))).(string)
        if this.Server == "" {
            this.Server = empty
        }
        this.ServerPort = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SERVERPORT))).(string)
        if this.ServerPort == "" {
            this.ServerPort = empty
        }
        this.ServerPath = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SERVERPATH))).(string)
        if this.ServerPath == "" {
            this.ServerPath = empty
        }
        this.UserName = any(hashMap.Get(fmt.Sprint(UserData_USERNAME))).(string)
        if this.UserName == "" {
            this.UserName = empty
        }
        this.Password = any(hashMap.Get(fmt.Sprint(UserData_PASSWORD))).(string)
        if this.Password == "" {
            this.Password = empty
        }
        this.Timeout = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_TIMEOUT))).(string)
        if this.Timeout == "" {
            this.Timeout = empty
        }
        this.ProxyProtocol = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYPROTOCOL))).(string)
        if this.ProxyProtocol == "" {
            this.ProxyProtocol = empty
        }
        this.ProxyServer = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYSERVER))).(string)
        if this.ProxyServer == "" {
            this.ProxyServer = empty
        }
        this.ProxyPort = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYPORT))).(string)
        if this.ProxyPort == "" {
            this.ProxyPort = empty
        }
        this.ProxyPath = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYPATH))).(string)
        if this.ProxyPath == "" {
            this.ProxyPath = empty
        }
        this.ProxyUserName = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYUSERNAME))).(string)
        if this.ProxyUserName == "" {
            this.ProxyUserName = empty
        }
        this.ProxyPassword = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYPASSWORD))).(string)
        if this.ProxyPassword == "" {
            this.ProxyPassword = empty
        }
        this.ProxyTimeout = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_PROXYTIMEOUT))).(string)
        if this.ProxyTimeout == "" {
            this.ProxyTimeout = empty
        }
        this.Special1 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL1))).(string)
        if this.Special1 == "" {
            this.Special1 = empty
        }
        this.Special2 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL2))).(string)
        if this.Special2 == "" {
            this.Special2 = empty
        }
        this.Special3 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL3))).(string)
        if this.Special3 == "" {
            this.Special3 = empty
        }
        this.Special4 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL4))).(string)
        if this.Special4 == "" {
            this.Special4 = empty
        }
        this.Special5 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL5))).(string)
        if this.Special5 == "" {
            this.Special5 = empty
        }
        this.Special6 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL6))).(string)
        if this.Special6 == "" {
            this.Special6 = empty
        }
        this.Special7 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL7))).(string)
        if this.Special7 == "" {
            this.Special7 = empty
        }
        this.Special8 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL8))).(string)
        if this.Special8 == "" {
            this.Special8 = empty
        }
        this.Special9 = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL9))).(string)
        if this.Special9 == "" {
            this.Special9 = empty
        }
        this.LastModified = any(hashMap.Get(fmt.Sprint(EntryData_GetInstance().LASTMODIFIED))).(string)
        if this.LastModified == "" {
            this.LastModified = empty
        }
        this.TimeEntered = any(hashMap.Get(fmt.Sprint(EntryData_GetInstance().TIMECREATED))).(string)
        if this.TimeEntered == "" {
            this.TimeEntered = empty
        }
}


func (this *PaymentGateway) GetStoreName() string {
        return this.StoreName
}


func (this *PaymentGateway) GetName() string {
        return this.Name
}


func (this *PaymentGateway) GetEnable() string {
        return this.Enable
}


func (this *PaymentGateway) GetPaymentMethod() string {
        return this.PaymentMethod
}


func (this *PaymentGateway) GetMode() string {
        return this.Mode
}


func (this *PaymentGateway) GetTestProtocol() string {
        return this.TestProtocol
}


func (this *PaymentGateway) GetTestServer() string {
        return this.TestServer
}


func (this *PaymentGateway) GetTestPort() string {
        return this.TestPort
}


func (this *PaymentGateway) GetTestPath() string {
        return this.TestPath
}


func (this *PaymentGateway) GetServerProtocol() string {
        return this.ServerProtocol
}


func (this *PaymentGateway) GetServer() string {
        return this.Server
}


func (this *PaymentGateway) GetServerPort() string {
        return this.ServerPort
}


func (this *PaymentGateway) GetServerPath() string {
        return this.ServerPath
}


func (this *PaymentGateway) GetUserName() string {
        return this.UserName
}


func (this *PaymentGateway) GetPassword() string {
        return this.Password
}


func (this *PaymentGateway) GetProxyProtocol() string {
        return this.ProxyProtocol
}


func (this *PaymentGateway) GetProxyServer() string {
        return this.ProxyServer
}


func (this *PaymentGateway) GetProxyPath() string {
        return this.ProxyPath
}


func (this *PaymentGateway) GetProxyPort() string {
        return this.ProxyPort
}


func (this *PaymentGateway) GetProxyUserName() string {
        return this.ProxyUserName
}


func (this *PaymentGateway) GetProxyPassword() string {
        return this.ProxyPassword
}


func (this *PaymentGateway) GetTimeout() string {
        return this.Timeout
}


func (this *PaymentGateway) GetProxyTimeout() string {
        return this.ProxyTimeout
}


func (this *PaymentGateway) GetSpecial1() string {
        return this.Special1
}


func (this *PaymentGateway) GetSpecial2() string {
        return this.Special2
}


func (this *PaymentGateway) GetSpecial3() string {
        return this.Special3
}


func (this *PaymentGateway) GetSpecial4() string {
        return this.Special4
}


func (this *PaymentGateway) GetSpecial5() string {
        return this.Special5
}


func (this *PaymentGateway) GetSpecial6() string {
        return this.Special6
}


func (this *PaymentGateway) GetSpecial7() string {
        return this.Special7
}


func (this *PaymentGateway) GetSpecial8() string {
        return this.Special8
}


func (this *PaymentGateway) GetSpecial9() string {
        return this.Special9
}


func (this *PaymentGateway) GetLastModified() string {
        return this.LastModified
}


func (this *PaymentGateway) GetTimeEntered() string {
        return this.TimeEntered
}


func (this *PaymentGateway) SetEnable(value string) {
        this.Enable = value
}


func (this *PaymentGateway) SetStoreName(value string) {
        this.StoreName = value
}


func (this *PaymentGateway) SetName(value string) {
        this.Name = value
}


func (this *PaymentGateway) SetPaymentMethod(value string) {
        this.PaymentMethod = value
}


func (this *PaymentGateway) SetMode(value string) {
        this.Mode = value
}


func (this *PaymentGateway) SetTestProtocol(value string) {
        this.TestProtocol = value
}


func (this *PaymentGateway) SetTestServer(value string) {
        this.TestServer = value
}


func (this *PaymentGateway) SetTestPort(value string) {
        this.TestPort = value
}


func (this *PaymentGateway) SetTestPath(value string) {
        this.TestPath = value
}


func (this *PaymentGateway) SetServerProtocol(value string) {
        this.ServerProtocol = value
}


func (this *PaymentGateway) SetServer(value string) {
        this.Server = value
}


func (this *PaymentGateway) SetServerPort(value string) {
        this.ServerPort = value
}


func (this *PaymentGateway) SetServerPath(value string) {
        this.ServerPath = value
}


func (this *PaymentGateway) SetTimeout(value string) {
        this.Timeout = value
}


func (this *PaymentGateway) SetUserName(value string) {
        this.UserName = value
}


func (this *PaymentGateway) SetPassword(value string) {
        this.Password = value
}


func (this *PaymentGateway) SetProxyProtocol(value string) {
        this.ProxyProtocol = value
}


func (this *PaymentGateway) SetProxyServer(value string) {
        this.ProxyServer = value
}


func (this *PaymentGateway) SetProxyPort(value string) {
        this.ProxyPort = value
}


func (this *PaymentGateway) SetProxyPath(value string) {
        this.ProxyPath = value
}


func (this *PaymentGateway) SetProxyUserName(value string) {
        this.ProxyUserName = value
}


func (this *PaymentGateway) SetProxyPassword(value string) {
        this.ProxyPassword = value
}


func (this *PaymentGateway) SetProxyTimeout(value string) {
        this.ProxyTimeout = value
}


func (this *PaymentGateway) SetSpecial1(value string) {
        this.Special1 = value
}


func (this *PaymentGateway) SetSpecial2(value string) {
        this.Special2 = value
}


func (this *PaymentGateway) SetSpecial3(value string) {
        this.Special3 = value
}


func (this *PaymentGateway) SetSpecial4(value string) {
        this.Special4 = value
}


func (this *PaymentGateway) SetSpecial5(value string) {
        this.Special5 = value
}


func (this *PaymentGateway) SetSpecial6(value string) {
        this.Special6 = value
}


func (this *PaymentGateway) SetSpecial7(value string) {
        this.Special7 = value
}


func (this *PaymentGateway) SetSpecial8(value string) {
        this.Special8 = value
}


func (this *PaymentGateway) SetSpecial9(value string) {
        this.Special9 = value
}


func (this *PaymentGateway) SetLastModified(value string) {
        this.LastModified = value
}


func (this *PaymentGateway) SetTimeEntered(value string) {
        this.TimeEntered = value
}


// throws Exception
func (this *PaymentGateway) ToHashMap() *HashMap {
        return this.ToHashMap(true)
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *PaymentGateway) ToHashMap(isEncrypted bool) *HashMap {
        var paymentGatewayHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = paymentGatewayHashMap
        paymentGatewayHashMap.Put(EntryData_GetInstance().ENABLE, this.Enable)
        paymentGatewayHashMap.Put(fmt.Sprint(StoreFrontData_GetInstance().NAME), this.StoreName)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_NAME), this.GetName())
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentData_METHOD), this.PaymentMethod)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_MODE), this.Mode)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_TESTPROTOCOL), this.TestProtocol)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_TESTSERVER), this.TestServer)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_TESTPORT), this.TestPort)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_TESTPATH), this.TestPath)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SERVERPROTOCOL), this.ServerProtocol)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SERVER), this.Server)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SERVERPORT), this.ServerPort)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SERVERPATH), this.ServerPath)
        if isEncrypted {
            var superCrypt *SuperCrypt = NewSuperCrypt(PaymentGatewayData_CRYPTNUM); _ = superCrypt
            paymentGatewayHashMap.Put(fmt.Sprint(UserData_USERNAME), superCrypt.Encrypt(this.UserName))
            paymentGatewayHashMap.Put(fmt.Sprint(UserData_PASSWORD), superCrypt.Encrypt(this.Password))
        } else {
            paymentGatewayHashMap.Put(fmt.Sprint(UserData_USERNAME), this.UserName)
            paymentGatewayHashMap.Put(fmt.Sprint(UserData_PASSWORD), this.Password)
        }
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_TIMEOUT), this.Timeout)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYPROTOCOL), this.ProxyProtocol)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYSERVER), this.ProxyServer)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYPORT), this.ProxyPort)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYPATH), this.ProxyPath)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYUSERNAME), this.ProxyUserName)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYPASSWORD), this.ProxyPassword)
        paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_PROXYTIMEOUT), this.ProxyTimeout)
        if isEncrypted {
            var superCrypt *SuperCrypt = NewSuperCrypt(PaymentGatewayData_CRYPTNUM); _ = superCrypt
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL1), superCrypt.Encrypt(this.Special1))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL2), superCrypt.Encrypt(this.Special2))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL3), superCrypt.Encrypt(this.Special3))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL4), superCrypt.Encrypt(this.Special4))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL5), superCrypt.Encrypt(this.Special5))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL6), superCrypt.Encrypt(this.Special6))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL7), superCrypt.Encrypt(this.Special7))
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL8), superCrypt.Encrypt(this.Special8))
        } else {
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL1), this.Special1)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL2), this.Special2)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL3), this.Special3)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL4), this.Special4)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL5), this.Special5)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL6), this.Special6)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL7), this.Special7)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL8), this.Special8)
            paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL9), this.Special9)
        }
        paymentGatewayHashMap.Put(fmt.Sprint(EntryData_GetInstance().LASTMODIFIED), this.LastModified)
        return paymentGatewayHashMap
}


func (this *PaymentGateway) ToVector() *BasicArrayList {
        return this.ToVector(true)
}


// Java overloaded method: Go needs a different name
func (this *PaymentGateway) ToVector(isEncrypted bool) *BasicArrayList {
        var updateVector *BasicArrayList = NewBasicArrayListD(); _ = updateVector
        updateVector.Add(this.Enable)
        updateVector.Add(this.StoreName)
        updateVector.Add(this.GetName())
        updateVector.Add(this.PaymentMethod)
        updateVector.Add(this.Mode)
        updateVector.Add(this.TestProtocol)
        updateVector.Add(this.TestServer)
        updateVector.Add(this.TestPort)
        updateVector.Add(this.TestPath)
        updateVector.Add(this.ServerProtocol)
        updateVector.Add(this.Server)
        updateVector.Add(this.ServerPort)
        updateVector.Add(this.ServerPath)
        if isEncrypted {
            var superCrypt *SuperCrypt = NewSuperCrypt(PaymentGatewayData_CRYPTNUM); _ = superCrypt
            updateVector.Add(superCrypt.Encrypt(this.UserName))
            updateVector.Add(superCrypt.Encrypt(this.Password))
        } else {
            updateVector.Add(this.UserName)
            updateVector.Add(this.Password)
        }
        updateVector.Add(this.Timeout)
        updateVector.Add(this.ProxyProtocol)
        updateVector.Add(this.ProxyServer)
        updateVector.Add(this.ProxyPort)
        updateVector.Add(this.ProxyPath)
        updateVector.Add(this.ProxyUserName)
        updateVector.Add(this.ProxyPassword)
        updateVector.Add(this.ProxyTimeout)
        if isEncrypted {
            var superCrypt *SuperCrypt = NewSuperCrypt(PaymentGatewayData_CRYPTNUM); _ = superCrypt
            updateVector.Add(superCrypt.Encrypt(this.Special1))
            updateVector.Add(superCrypt.Encrypt(this.Special2))
            updateVector.Add(superCrypt.Encrypt(this.Special3))
            updateVector.Add(superCrypt.Encrypt(this.Special4))
            updateVector.Add(superCrypt.Encrypt(this.Special5))
            updateVector.Add(superCrypt.Encrypt(this.Special6))
            updateVector.Add(superCrypt.Encrypt(this.Special7))
            updateVector.Add(superCrypt.Encrypt(this.Special8))
            updateVector.Add(superCrypt.Encrypt(this.Special9))
        } else {
            updateVector.Add(this.Special1)
            updateVector.Add(this.Special2)
            updateVector.Add(this.Special3)
            updateVector.Add(this.Special4)
            updateVector.Add(this.Special5)
            updateVector.Add(this.Special6)
            updateVector.Add(this.Special7)
            updateVector.Add(this.Special8)
            updateVector.Add(this.Special9)
        }
        return updateVector
}


// throws Exception
func (this *PaymentGateway) GetKey() any {
        return any(NewPaymentGatewayPrimaryKey(this.GetStoreName(), this.GetName())).(any)
}

