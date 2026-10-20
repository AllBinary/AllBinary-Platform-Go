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

type CustomizersView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CustomizersVector *BasicArrayList
}

// Implements DomNodeInterface


func (this *CustomizersView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewCustomizersView(transformInfoInterface TransformInfoInterface) *CustomizersView {
    this := &CustomizersView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.CustomizersVector = TransformInfoEntityBuilder_GetInstance().GetNames(this.GetWeblisketSession().GetStoreName())
    return this
}


func (this *CustomizersView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var node *Node = document.CreateElement(TransformTemplateCustomizersData_NAME); _ = node
            var unsortedCustomizerViewVector *BasicArrayList = NewBasicArrayListD(); _ = unsortedCustomizerViewVector
            var size int = this.CustomizersVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var viewName string = any(this.CustomizersVector.Get(index)).(string); _ = viewName
                if strings.Index(viewName, string(CustomizerTransformInfoData_NAME)) > 0 && strings.Index(viewName, string(HttpStoreComponentView_GLOBALS2.EDIT)) > 0 && strings.Index(viewName, string(BodyData_GetInstance().VIEWNAMEKEY)) > 0 && strings.Index(viewName, string(TitleData_GetInstance().VIEWNAMEKEY)) > 0 {
                    unsortedCustomizerViewVector.Add(viewName)
                }
            }
            var objectArray []any = any(unsortedCustomizerViewVector.ToArray()).([]any); _ = objectArray
            Arrays_Sort(objectArray, NewStringComparator())
            for index := int(0); index < len(objectArray); index++ {
                var viewName string = any(objectArray[index]).(string); _ = viewName
                var viewNameNode *Node = ModDomHelper_CreateNameValueNodes(document, TransformTemplateCustomizerData_NAME, viewName, TransformTemplateCustomizerUtil_GetInstance().GetPageNameHack(viewName, this.GetWeblisketSession().GetStoreName())); _ = viewNameNode
                node.AppendChild(viewNameNode)
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CustomizersView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *CustomizersView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            this.AddDomNodeInterfaces()
            tryControl1, tryResult1 = 1, this.HttpStoreComponentView.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

