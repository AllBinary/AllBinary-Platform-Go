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

type InventoryView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}

// Implements DomNodeInterface

// The Java class InventoryView has subclasses so its variables are this interface
type IInventoryView interface {
    IHttpStoreComponentView
    InventoryView_() *InventoryView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *InventoryView) InventoryView_() *InventoryView {
    return this
}


func (this *InventoryView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewInventoryView(transformInfoInterface TransformInfoInterface) *InventoryView {
    this := &InventoryView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *InventoryView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Starting", this, "toXmlNode")
            }
            var inventoryEntityInterface *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntityInterface
            var inventoryNode *Node = document.CreateElement(InventoryData_INVENTORY); _ = inventoryNode
            var itemVector IBasicArrayList = inventoryEntityInterface.GetItems(StoreFrontFactory_GetInstance(this.GetTransformInfoInterface().GetStoreName())); _ = itemVector
            inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, SearchData_TOTAL_NUMBER_ITEMS_ON_THIS_PAGE, fmt.Sprint(int(itemVector.Size()))))
            var size int = itemVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var itemInterface ItemInterface = any(itemVector.Get(index)).(ItemInterface); _ = itemInterface
                if itemInterface != nil {
                    var node *Node = NewBasicItemView(itemInterface, NewBasicArrayListD()).ToXmlNode(document); _ = node
                    inventoryNode.AppendChild(node)
                } else {
                    this.LogUtil.PutF("Inventory", this, "toXmlNode")
                }
            }
            tryControl1, tryResult1 = 1, inventoryNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InventoryView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


func (this *InventoryView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view Inventory"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
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

