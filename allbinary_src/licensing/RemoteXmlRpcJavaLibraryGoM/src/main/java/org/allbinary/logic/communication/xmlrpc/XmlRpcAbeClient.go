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

type XmlRpcAbeClient struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RemoteMethod string
    ClientInfo AbeClientInformationInterface
    Client XmlRpcHandler
    Server int
    Start int
    MaxServers int
    IsDone bool
    START_SERVER string
    TRYING string
    SEP string
    CLIENT_INFO string
    RESULT string
    INVALID string
    EXCEPTION_IN_CLIENT string
    SERVER_REPORTED_ERROR string
    UNKNOWN_ERROR string
    TRYING_OTHER_SERVERS string
    HOST_NOT_RESOLVED_MSG string
    HOST_NOT_RESOLVED string
    MyRandomFactory *MyRandomFactory
}

// The Java class XmlRpcAbeClient has subclasses so its variables are this interface
type IXmlRpcAbeClient interface {
    XmlRpcAbeClient_() *XmlRpcAbeClient
    Get(anyType any) any
    TryAnother(anyType any) any
    GetClientInfo() AbeClientInformationInterface
    GetClient() XmlRpcHandler
    SetClient(client IXmlRpcClient)
    GetServer() int
    SetServer(server int)
    GetStart() int
    SetStart(start int)
    GetMaxServers() int
    SetMaxServers(maxServers int)
    IsIsDone() bool
    SetIsDone(isDone bool)
    GetRemoteMethod() string
}

func (this *XmlRpcAbeClient) XmlRpcAbeClient_() *XmlRpcAbeClient {
    return this
}

var XmlRpcAbeClient_IsOnline bool = true

func (this *XmlRpcAbeClient) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.Client = NullXmlRpcHandler_NULL_XML_RPC_HANDLER
    
    
    
    
    this.START_SERVER = "Start With Server #"
    this.TRYING = "Trying Server #"
    this.SEP = CommonLabels_GetInstance().COLON_SEP
    this.CLIENT_INFO = "Client Info: \n"
    this.RESULT = CommonLabels_GetInstance().RESULT_ + CommonSeps_GetInstance().NEW_LINE
    this.INVALID = "License data is Invalid Trying Other Servers"
    this.EXCEPTION_IN_CLIENT = "Exception in client"
    this.SERVER_REPORTED_ERROR = "Server reported error"
    this.UNKNOWN_ERROR = "Unknown License Retrieval Failure"
    this.TRYING_OTHER_SERVERS = "IOException Trying Other Servers"
    this.HOST_NOT_RESOLVED_MSG = "Not Trying Again Since Host Unresolved"
    this.HOST_NOT_RESOLVED = "Host is unresolved"
    this.MyRandomFactory = MyRandomFactory_GetInstance()
}


func NewXmlRpcAbeClient(clientInfo AbeClientInformationInterface, remoteMethod string) *XmlRpcAbeClient {
    this := &XmlRpcAbeClient{}
    this.initFields()
        this.RemoteMethod = remoteMethod
        this.ClientInfo = clientInfo
        if clientInfo.GetNumberOfLicenseServers() > 1 {
            this.MaxServers = clientInfo.GetNumberOfLicenseServers() - 2
            this.Start = this.MyRandomFactory.GetAbsoluteNextInt(this.MaxServers) + 1
        } else if clientInfo.GetNumberOfLicenseServers() == 1 {
            this.MaxServers = 0
            this.Start = 0
        }
        this.IsDone = false
        this.Server = this.Start
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.START_SERVER)
        stringBuffer.Appendint(this.GetServer())
        stringBuffer.Append(this.SEP)
        stringBuffer.Append(clientInfo.GetLicenseServer(this.GetServer()))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


// throws Exception
func (this *XmlRpcAbeClient) Get(anyType any) any {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
        return NullUtil_GetInstance().NULL_OBJECT
}


// throws Exception
func (this *XmlRpcAbeClient) TryAnother(anyType any) any {
        if this.GetServer() < this.GetMaxServers() {
            this.SetServer(this.GetServer() + 1)
        } else {
            this.SetServer(0)
        }
        if this.GetServer() != this.GetStart() && !this.IsIsDone() {
            return this.Get(anyType)
        } else {
            this.SetIsDone(true)
        }
        panic(errors.New(fmt.Sprint("Tried All Servers But Still Failed")))
}


func (this *XmlRpcAbeClient) GetClientInfo() AbeClientInformationInterface {
        return this.ClientInfo
}


func (this *XmlRpcAbeClient) GetClient() XmlRpcHandler {
        return this.Client
}


func (this *XmlRpcAbeClient) SetClient(client IXmlRpcClient) {
        this.Client = client
}


func (this *XmlRpcAbeClient) GetServer() int {
        return this.Server
}


func (this *XmlRpcAbeClient) SetServer(server int) {
        this.Server = server
}


func (this *XmlRpcAbeClient) GetStart() int {
        return this.Start
}


func (this *XmlRpcAbeClient) SetStart(start int) {
        this.Start = start
}


func (this *XmlRpcAbeClient) GetMaxServers() int {
        return this.MaxServers
}


func (this *XmlRpcAbeClient) SetMaxServers(maxServers int) {
        this.MaxServers = maxServers
}


func (this *XmlRpcAbeClient) IsIsDone() bool {
        return this.IsDone
}


func (this *XmlRpcAbeClient) SetIsDone(isDone bool) {
        this.IsDone = isDone
}


func (this *XmlRpcAbeClient) GetRemoteMethod() string {
        return this.RemoteMethod
}

