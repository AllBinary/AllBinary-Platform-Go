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

type AbeHttpRequestInfo struct {
    NAME string
    AbeHttpRequestInfoData *AbeHttpRequestInfoData
    CommonSeps *CommonSeps
    HttpUserAgent string
    RemoteAddress string
    RemoteHost string
    RemoteHostByAddr string
    RemotePort string
    RequestedFilePath string
}


func (this *AbeHttpRequestInfo) initFields() {
    this.NAME = "AbeHttpRequestInfo"
    this.AbeHttpRequestInfoData = AbeHttpRequestInfoData_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    
    
    
    
    
    
}


func NewAbeHttpRequestInfo(hashMap *HashMap) *AbeHttpRequestInfo {
    this := &AbeHttpRequestInfo{}
    this.initFields()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.HttpUserAgent = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.HTTP_USER_AGENT)).(string))
        this.RemoteAddress = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.REMOTE_ADDRESS)).(string))
        this.RemoteHost = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.REMOTE_HOST)).(string))
        this.RemoteHostByAddr = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.REMOTE_HOST_BY_ADDRESS)).(string))
        this.RemotePort = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.REMOTE_PORT)).(string))
        this.RequestedFilePath = stringUtil.GetNonNull(any(hashMap.Get(this.AbeHttpRequestInfoData.REQUEST_FILE_PATH)).(string))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewAbeHttpRequestInfo(httpServletRequest *HttpServletRequest) *AbeHttpRequestInfo {
    this := &AbeHttpRequestInfo{}
    this.initFields()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.HttpUserAgent = stringUtil.EMPTY_STRING
        var enumuration Enumeration = httpServletRequest.GetHeaderNames(); _ = enumuration
        for enumuration.HasMoreElements() {
            var key string = any(enumuration.NextElement()).(string); _ = key
            var value string = httpServletRequest.GetHeader(key); _ = value
            if strings.Index(key, string("user")) >= 0 {
                this.HttpUserAgent = "key: " + key + " value: " + value
            }
        }
        this.RemoteHostByAddr = stringUtil.EMPTY_STRING
        this.RemoteAddress = httpServletRequest.GetRemoteAddr()
        this.RemoteHost = httpServletRequest.GetRemoteHost()
        this.RequestedFilePath = httpServletRequest.GetServletPath()
        this.RemotePort = fmt.Sprint(httpServletRequest.GetRemotePort())
    return this
}


func (this *AbeHttpRequestInfo) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(this.AbeHttpRequestInfoData.HTTP_USER_AGENT, this.HttpUserAgent)
        hashMap.Put(this.AbeHttpRequestInfoData.REMOTE_ADDRESS, this.RemoteAddress)
        hashMap.Put(this.AbeHttpRequestInfoData.REMOTE_HOST, this.RemoteHost)
        hashMap.Put(this.AbeHttpRequestInfoData.REMOTE_HOST_BY_ADDRESS, this.RemoteHostByAddr)
        hashMap.Put(this.AbeHttpRequestInfoData.REMOTE_PORT, this.RemotePort)
        hashMap.Put(this.AbeHttpRequestInfoData.REQUEST_FILE_PATH, this.RequestedFilePath)
        return hashMap
}


func (this *AbeHttpRequestInfo) ToVector() *BasicArrayList {
        var vector *BasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.HttpUserAgent)
        vector.Add(this.RemoteAddress)
        vector.Add(this.RemoteHost)
        vector.Add(this.RemoteHostByAddr)
        vector.Add(this.RemotePort)
        vector.Add(this.RequestedFilePath)
        return vector
}


func (this *AbeHttpRequestInfo) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.NAME).Append(this.CommonSeps.COLON_SEP).Append(this.AbeHttpRequestInfoData.HTTP_USER_AGENT).Append(this.CommonSeps.COLON_SEP).Append(this.HttpUserAgent).Append(this.CommonSeps.SPACE).Append(this.AbeHttpRequestInfoData.REMOTE_ADDRESS).Append(this.CommonSeps.COLON_SEP).Append(this.RemoteAddress).Append(this.CommonSeps.SPACE).Append(this.AbeHttpRequestInfoData.REMOTE_HOST).Append(this.CommonSeps.COLON_SEP).Append(this.RemoteHost).Append(this.CommonSeps.SPACE).Append(this.AbeHttpRequestInfoData.REMOTE_HOST_BY_ADDRESS).Append(this.CommonSeps.COLON_SEP).Append(this.RemoteHostByAddr).Append(this.CommonSeps.SPACE).Append(this.AbeHttpRequestInfoData.REMOTE_PORT).Append(this.CommonSeps.COLON_SEP).Append(this.RemotePort).Append(this.CommonSeps.SPACE).Append(this.AbeHttpRequestInfoData.REQUEST_FILE_PATH).Append(this.CommonSeps.COLON_SEP).Append(this.RequestedFilePath))
}

func (this *AbeHttpRequestInfo) String() string {
    return this.ToString()
}

