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

//Note: if request is destroyed or messed up this will not work. Use fileRequestParams
type RequestParams struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Map Map
    KEY string
    VALUE string
}

// The Java class RequestParams has subclasses so its variables are this interface
type IRequestParams interface {
    RequestParams_() *RequestParams
    SetMap(map_ Map)
    GetMap() Map
    ToXmlNode(document *Document) *Node
    ToHashMap() *HashMap
}

func (this *RequestParams) RequestParams_() *RequestParams {
    return this
}


func (this *RequestParams) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.KEY = "Key: "
    this.VALUE = " Value: "
}




// Java overloads of new RequestParams
func NewRequestParams(args ...any) *RequestParams {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewRequestParams_1()
        }
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return NewRequestParams_2(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*PageContext](args[0], pass) {
            return NewRequestParams_3(Java_Arg[*PageContext](args[0]))
        }
    }
    panic("No Java overload of new RequestParams for the arguments")
}
func NewRequestParams_1() *RequestParams {
    this := &RequestParams{}
    this.initFields()
        this.LogUtil.PutF("New RequestParams Size: 0", this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func NewRequestParams_2(request *HttpServletRequest) *RequestParams {
    this := &RequestParams{}
    this.initFields()
        this.Map = request.GetParameterMap()
        this.LogUtil.PutF("RequestParams Size: " + fmt.Sprint(this.GetMap().KeySet().Size()), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func NewRequestParams_3(pageContext *PageContext) *RequestParams {
    this := &RequestParams{}
    this.initFields()
        this.Map = pageContext.GetRequest().GetParameterMap()
        this.LogUtil.PutF("Request Params Size: " + fmt.Sprint(this.GetMap().KeySet().Size()), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *RequestParams) SetMap(map_ Map) {
        this.Map = map_
}


func (this *RequestParams) GetMap() Map {
        return this.Map
}


func (this *RequestParams) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode(document)", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var node *Node = document.CreateElement(RequestData_REQUEST); _ = node
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var keys Set = this.Map.KeySet(); _ = keys
            var keyArray []any = keys.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for i := int(0); i < size; i++ {
                var key string = any(keyArray[i]).(string); _ = key
                var values []string = any(this.Map.Get(key)).([]string); _ = values
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.KEY)
                stringBuffer.Append(key)
                stringBuffer.Append(this.VALUE)
                stringBuffer.Append(values[0])
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "toXmlNode(document)")
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, RequestData_PARAMETER, key, values[0]))
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *RequestParams) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var keys Set = this.Map.KeySet(); _ = keys
        var keyArray []any = keys.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for i := int(0); i < size; i++ {
            var key string = any(keyArray[i]).(string); _ = key
            var values []string = any(this.Map.Get(key)).([]string); _ = values
            hashMap.Put(key, values[0])
        }
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append("RequestParams: ")
        stringBuffer.Append(fmt.Sprint(this.Map))
        stringBuffer.Append("\ntoHashMap: ")
        stringBuffer.Append(fmt.Sprint(hashMap))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "toHashMap()")
        return hashMap
}

