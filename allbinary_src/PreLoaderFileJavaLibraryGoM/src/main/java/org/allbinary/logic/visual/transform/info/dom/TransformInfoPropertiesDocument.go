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

type TransformInfoPropertiesDocument struct {
    LogUtil *LogUtil
    Document *Document
}


func (this *TransformInfoPropertiesDocument) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewTransformInfoPropertiesDocument(filePath string) *TransformInfoPropertiesDocument {
    this := &TransformInfoPropertiesDocument{}
    this.initFields()
        var xmlFile *AbFile = AbFile_CreateAbFile(filePath); _ = xmlFile
        this.Document = DomDocumentFileHelper_CreateDocument(xmlFile)
    return this
}


func (this *TransformInfoPropertiesDocument) ToTransformInfoPropertiesHashMap() *HashMap {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HashMap) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "toTransformInfoPropertiesHashMap()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var transformInfoPropertiesHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = transformInfoPropertiesHashMap
            var transformInfosNode *Node = this.Document.GetElementsByTagName(TransformInfosData_GetInstance().NAME).Item(0); _ = transformInfosNode
            var transformInfosChildNodeList *NodeList = transformInfosNode.GetChildNodes(); _ = transformInfosChildNodeList
            var transformInfoNodeVector *BasicArrayList = DomSearchHelper_GetAllNodes(TransformInfoData_GetInstance().NAME, transformInfosChildNodeList); _ = transformInfoNodeVector
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Size: " + fmt.Sprint(transformInfoNodeVector.Size()), this, "toTransformInfoPropertiesHashMap()")
            }
            var size int = transformInfoNodeVector.Size(); _ = size
            var node *Node; _ = node
            var transformInfoProperties *TransformInfoProperties; _ = transformInfoProperties
            for index := int(0); index < size; index++ {
                node = any(transformInfoNodeVector.Get(index)).(*Node)
                transformInfoProperties = TransformInfoPropertiesFactory_GetInstance().GetInstance(node)
                transformInfoPropertiesHashMap.Put(transformInfoProperties.GetName(), transformInfoProperties)
            }
            tryControl1, tryResult1 = 1, transformInfoPropertiesHashMap; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

