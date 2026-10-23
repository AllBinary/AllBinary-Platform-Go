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

type BasicWorkFlow struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    WorkFlowDoc *Document
    WorkFlowName string
    StoreName string
}

// The Java class BasicWorkFlow has subclasses so its variables are this interface
type IBasicWorkFlow interface {
    BasicWorkFlow_() *BasicWorkFlow
    GetName() string
    GetStoreName() string
    GetKey() any
    ToVector() IBasicArrayList
    ToHashMap() *HashMap
    ToXmlNode(document *Document) *Node
    ToXmlDoc() *Document
    IsValid() bool
    ValidationInfo() string
    ToValidationInfoDoc() *Document
    ToValidationInfoNode(document *Document) *Node
}

func (this *BasicWorkFlow) BasicWorkFlow_() *BasicWorkFlow {
    return this
}


func (this *BasicWorkFlow) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}




// Java overloads of new BasicWorkFlow
func NewBasicWorkFlow(args ...any) *BasicWorkFlow {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewBasicWorkFlow_1()
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewBasicWorkFlow_2(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewBasicWorkFlow_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new BasicWorkFlow for the arguments")
}
func NewBasicWorkFlow_1() *BasicWorkFlow {
    this := &BasicWorkFlow{}
    this.initFields()
        this.WorkFlowName = StringUtil_GetInstance().EMPTY_STRING
        this.WorkFlowDoc = DomDocumentHelper_Create()
        var workFlowNode *Node = this.WorkFlowDoc.CreateElement(WorkFlowData_GetInstance().WORKFLOW); _ = workFlowNode
        this.WorkFlowDoc.AppendChild(workFlowNode)
    return this
}


func NewBasicWorkFlow_2(hashMap *HashMap, pageContext *PageContext) *BasicWorkFlow {
    this := &BasicWorkFlow{}
    this.initFields()
        var weblisketSession *WeblisketSession = NewWeblisketSession(hashMap, pageContext); _ = weblisketSession
        this.StoreName = weblisketSession.GetStoreName()
        var requestHashMap *HashMap = NewRequestParams(pageContext).ToHashMap(); _ = requestHashMap
        this.WorkFlowName = any(requestHashMap.Get(WorkFlowData_GetInstance().NAME)).(string)
        this.WorkFlowDoc = DomDocumentHelper_Create(any(requestHashMap.Get(WorkFlowData_GetInstance().DATA)).(string))
    return this
}


func NewBasicWorkFlow_3(hashMap *HashMap) *BasicWorkFlow {
    this := &BasicWorkFlow{}
    this.initFields()
        this.WorkFlowName = any(hashMap.Get(WorkFlowData_GetInstance().NAME)).(string)
        this.StoreName = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string)
        this.WorkFlowDoc = DomDocumentHelper_Create(any(hashMap.Get(WorkFlowData_GetInstance().DATA)).(string))
    return this
}


func (this *BasicWorkFlow) GetName() string {
        return this.WorkFlowName
}


func (this *BasicWorkFlow) GetStoreName() string {
        return this.StoreName
}


func (this *BasicWorkFlow) GetKey() any {
        return any(this.GetName()).(any)
}


// throws Exception
func (this *BasicWorkFlow) ToVector() IBasicArrayList {
        var values IBasicArrayList = NewBasicArrayListD(); _ = values
        values.Add(this.WorkFlowName)
        values.Add(this.StoreName)
        values.Add(XmlDocumentHelper_ToString(this.WorkFlowDoc))
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Add(time)
        values.Add(time)
        return values
}


// throws Exception
func (this *BasicWorkFlow) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(WorkFlowData_GetInstance().NAME, this.WorkFlowName)
        hashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
        hashMap.Put(WorkFlowData_GetInstance().DATA, XmlDocumentHelper_ToString(this.WorkFlowDoc))
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        hashMap.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return hashMap
}


func (this *BasicWorkFlow) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var hashMap *HashMap = this.ToHashMap(); _ = hashMap
            var keySet Set = hashMap.KeySet(); _ = keySet
            var node *Node = document.CreateElement(WorkFlowData_GetInstance().WORKFLOW); _ = node
            var nameArray []any = keySet.ToArray(); _ = nameArray
            var size int = len(nameArray); _ = size
            for index := int(0); index < size; index++ {
                var name string = any(nameArray[index]).(string); _ = name
                var value string = any(hashMap.Get(name)).(string); _ = value
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicWorkFlow) ToXmlDoc() *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlDoc()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, this.WorkFlowDoc; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicWorkFlow) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var returnBoolean bool = true; _ = returnBoolean
            if this.WorkFlowName == "" {
                returnBoolean = false
            }
            tryControl1, tryResult1 = 1, returnBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicWorkFlow) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Getting Validation Info"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.WorkFlowName == "" {
                stringBuffer.Append("No WorkFlow name specified.")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicWorkFlow) ToValidationInfoDoc() *Document {
        return nil
}


func (this *BasicWorkFlow) ToValidationInfoNode(document *Document) *Node {
        return nil
}

