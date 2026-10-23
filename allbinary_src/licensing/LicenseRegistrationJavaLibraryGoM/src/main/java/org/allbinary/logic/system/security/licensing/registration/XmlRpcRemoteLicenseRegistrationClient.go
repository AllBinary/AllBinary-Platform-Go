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

type XmlRpcRemoteLicenseRegistrationClient struct {
    XmlRpcAbeClient
    LogUtil *LogUtil
    PAGE string
}


func (this *XmlRpcRemoteLicenseRegistrationClient) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PAGE = "licenseregistrationserverssl.php"
}


func NewXmlRpcRemoteLicenseRegistrationClient(clientInfo AbeClientInformationInterface) *XmlRpcRemoteLicenseRegistrationClient {
    this := &XmlRpcRemoteLicenseRegistrationClient{}
        this.XmlRpcAbeClient = *NewXmlRpcAbeClient(clientInfo, "LicenseRegistrationServer.process")
        this.initFields()
        this.SetServer(0)
    return this
}


// throws Exception
func (this *XmlRpcRemoteLicenseRegistrationClient) Get(anyType any, cryptInterface CryptInterface) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(XmlRpcAbeClient_TRYING_OTHER_SERVERS + ExceptionUtil_GetInstance().GetStackTrace(e), this, this.CommonStrings.GET)
            if !strings.HasPrefix(e.Error(), XmlRpcAbeClient_HOST_NOT_RESOLVED) {
                tryControl1, tryResult1 = 1, this.TryAnother(anyType); return
            } else {
                panic(errors.New(fmt.Sprint(XmlRpcAbeClient_HOST_NOT_RESOLVED_MSG)))
            }
                }
            }()
            var server string = this.GetClientInfo().GetLicenseServer(this.GetServer()); _ = server
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(XmlRpcAbeClient_TRYING)
            stringBuffer.Appendint(this.GetServer())
            stringBuffer.Append(SEP)
            stringBuffer.Append(server)
            this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START_LABEL) + fmt.Sprint(stringBuffer), this, this.CommonStrings.GET)
            var param IBasicArrayList = NewBasicArrayListD(); _ = param
            var serverUrl string = this.GetClientInfo().GetLicenseServer(this.GetServer()); _ = serverUrl
            var index int = strings.LastIndex(serverUrl, string("/")); _ = index
            serverUrl = serverUrl[0:index + 1] + this.PAGE
            var xmlRpcClient IXmlRpcClient = NewXmlRpcClient(serverUrl); _ = xmlRpcClient
            this.SetClient(xmlRpcClient)
            xmlRpcClient.SetBasicAuthentication(nil, nil)
            var hashtable *ABHashtable[any, any] = any(anyType).(*ABHashtable[any, any]); _ = hashtable
            this.LogUtil.PutF(fmt.Sprint(XmlRpcAbeClient_CLIENT_INFO) + fmt.Sprint(hashtable), this, this.CommonStrings.GET)
            param.Add(hashtable)
            var result any = xmlRpcClient.Execute(this.GetRemoteMethod(), param, cryptInterface); _ = result
            this.LogUtil.PutF(fmt.Sprint(XmlRpcAbeClient_RESULT) + fmt.Sprint(result), this, this.CommonStrings.GET)
            this.IsOnline = true
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

