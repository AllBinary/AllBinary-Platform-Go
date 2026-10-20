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

type XmlRpcRemoteLogClient struct {
    XmlRpcAbeClient
    NoCrypt *NoCrypt
}


func (this *XmlRpcRemoteLogClient) initFields() {
    this.NoCrypt = NewNoCrypt()
}


func NewXmlRpcRemoteLogClient(clientInfo AbeClientInformationInterface) *XmlRpcRemoteLogClient {
    this := &XmlRpcRemoteLogClient{}
        this.XmlRpcAbeClient = *NewXmlRpcAbeClient(clientInfo, "SSLLogServ.logUtil")
        this.initFields()
    return this
}


// throws Exception
func (this *XmlRpcRemoteLogClient) Get(anyType any) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println("IOException Trying Other Servers")
            if !strings.HasPrefix(e.Error(), XmlRpcAbeClient_HOST_NOT_RESOLVED) {
                tryControl1, tryResult1 = 1, this.TryAnother(anyType); return
            } else {
                panic(errors.New(fmt.Sprint(XmlRpcAbeClient_HOST_NOT_RESOLVED_MSG)))
            }
                }
            }()
            var param IBasicArrayList = NewBasicArrayListD(); _ = param
            var serverUrl string = this.GetClientInfo().GetLicenseServer(this.GetServer()); _ = serverUrl
            var index int = strings.LastIndex(serverUrl, string("/")); _ = index
            serverUrl = serverUrl[0:index + 1] + "logservssl.php"
            var xmlRpcClient IXmlRpcClient = NewXmlRpcClient(serverUrl); _ = xmlRpcClient
            this.SetClient(xmlRpcClient)
            xmlRpcClient.SetBasicAuthentication(nil, nil)
            var hashtable *ABHashtable[any, any] = any(anyType).(*ABHashtable[any, any]); _ = hashtable
            param.Add(hashtable)
            var result any = xmlRpcClient.Execute(this.GetRemoteMethod(), param, this.NoCrypt); _ = result
            fmt.Println(fmt.Sprint(CommonLabels_GetInstance().RESULT_ + CommonSeps_GetInstance().NEW_LINE) + fmt.Sprint(result))
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

