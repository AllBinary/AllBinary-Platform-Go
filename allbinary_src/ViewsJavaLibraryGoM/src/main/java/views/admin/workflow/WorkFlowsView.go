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

type WorkFlowsView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    WorkFlowsVector IBasicArrayList
}

// Implements DomNodeInterface

// The Java class WorkFlowsView has subclasses so its variables are this interface
type IWorkFlowsView interface {
    IHttpStoreComponentView
    WorkFlowsView_() *WorkFlowsView
    AddDomNodeInterfaces()
    View() string
    ToXmlNode(document *Document) *Node
}

func (this *WorkFlowsView) WorkFlowsView_() *WorkFlowsView {
    return this
}


func (this *WorkFlowsView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewWorkFlowsView(transformInfoInterface TransformInfoInterface) *WorkFlowsView {
    this := &WorkFlowsView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.WorkFlowsVector = WorkFlowEntityFactory_GetInstance().Create2().Get(this.GetWeblisketSession().GetStoreName()).(IBasicArrayList)
    return this
}


func (this *WorkFlowsView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *WorkFlowsView) View() string {
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


func (this *WorkFlowsView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var node *Node = document.CreateElement(WorkFlowData_GetInstance().WORKFLOWS); _ = node
            var size int = this.WorkFlowsVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var workFlowInterface WorkFlowInterface = any(this.WorkFlowsVector.Get(index)).(WorkFlowInterface); _ = workFlowInterface
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

