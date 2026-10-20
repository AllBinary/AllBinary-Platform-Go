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

type AbeHttpRequestInfoData struct {
    HTTP_USER_AGENT string
    REMOTE_ADDRESS string
    REMOTE_HOST string
    REMOTE_HOST_BY_ADDRESS string
    REMOTE_PORT string
    REQUEST_FILE_PATH string
}

var AbeHttpRequestInfoData_Instance *AbeHttpRequestInfoData = NewAbeHttpRequestInfoData()

func (this *AbeHttpRequestInfoData) initFields() {
    this.HTTP_USER_AGENT = "HTTP_USER_AGENT"
    this.REMOTE_ADDRESS = "REMOTE_ADDRESS"
    this.REMOTE_HOST = "REMOTE_HOST"
    this.REMOTE_HOST_BY_ADDRESS = "REMOTE_HOST_BY_ADDRESS"
    this.REMOTE_PORT = "REMOTE_PORT"
    this.REQUEST_FILE_PATH = "REQUEST_FILE_PATH"
}


func AbeHttpRequestInfoData_GetInstance() *AbeHttpRequestInfoData {
        return AbeHttpRequestInfoData_Instance
}


func NewAbeHttpRequestInfoData() *AbeHttpRequestInfoData {
    this := &AbeHttpRequestInfoData{}
    this.initFields()
    return this
}

