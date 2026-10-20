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

type ItemView struct {
    ItemInterface OrderItemInterface
    Vector IBasicArrayList
}


func (this *ItemView) initFields() {
    
    
}


func NewItemView(itemInterface OrderItemInterface, vector IBasicArrayList) *ItemView {
    this := &ItemView{}
    this.initFields()
        this.ItemInterface = itemInterface
        this.Vector = vector
    return this
}


// throws Exception
func (this *ItemView) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ItemInterface.ToHashMap(); _ = hashMap
        var keySet Set = hashMap.KeySet(); _ = keySet
        var node *Node = document.CreateElement(ItemData_ITEM); _ = node
        var basicItemView *BasicItemView = NewBasicItemView(this.ItemInterface, this.Vector); _ = basicItemView
        node.AppendChild(basicItemView.ToXmlNode(document))
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var nameArray []any = keySet.ToArray(); _ = nameArray
        var size int = len(nameArray); _ = size
        for index := int(0); index < size; index++ {
            var name string = any(nameArray[index]).(string); _ = name
            var value string = any(hashMap.Get(name)).(string); _ = value
            value = stringUtil.GetNonNull(value)
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        var shippingAddressNode *Node = document.CreateElement(StreetAddressData_NAME); _ = shippingAddressNode
        shippingAddressNode.AppendChild(this.ItemInterface.GetShippingAddress().ToXmlNode(document))
        return node
}

