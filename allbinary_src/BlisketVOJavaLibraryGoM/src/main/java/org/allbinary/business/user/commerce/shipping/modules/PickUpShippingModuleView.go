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

type PickUpShippingModuleView struct {
    PickUpShippingModule
    LogUtil *LogUtil
}

// Implements DomNodeInterface


func (this *PickUpShippingModuleView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPickUpShippingModuleView() *PickUpShippingModuleView {
    this := &PickUpShippingModuleView{}
    this.PickUpShippingModule = *NewPickUpShippingModule()
    this.initFields()
    return this
}


func (this *PickUpShippingModuleView) IsValid() bool {
        return nil
}


func (this *PickUpShippingModuleView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *PickUpShippingModuleView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *PickUpShippingModuleView) ToXmlDoc() *Document {
        return nil
}


func (this *PickUpShippingModuleView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var shippingMethodNode *Node = document.CreateElement(ShippingMethodData_NAME); _ = shippingMethodNode
            shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodData_NAME, this.GetName()))
            shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodData_DESCRIPTION, this.GetDescription()))
            tryControl1, tryResult1 = 1, shippingMethodNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PickUpShippingModuleView) ValidationInfo() string {
        return nil
}


func (this *PickUpShippingModuleView) View() string {
        return nil
}

