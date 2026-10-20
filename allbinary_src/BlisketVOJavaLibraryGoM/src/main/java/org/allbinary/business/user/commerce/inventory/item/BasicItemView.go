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

type BasicItemView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BasicItemData *BasicItemData
    ItemInterface ItemInterface
    Vector IBasicArrayList
}

// Implements DomNodeInterface


func (this *BasicItemView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
}


func NewBasicItemView(itemInterface ItemInterface, vector IBasicArrayList) *BasicItemView {
    this := &BasicItemView{}
    this.initFields()
        this.ItemInterface = itemInterface
        this.Vector = vector
    return this
}


// throws Exception
func (this *BasicItemView) ToXmlNode(document *Document) *Node {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "toXmlNode")
        }
        var hashMap *HashMap = this.ItemInterface.ToHashMap(); _ = hashMap
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        hashMap.Put(this.BasicItemData.IMAGE, EMPTY_STRING)
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var keySet Set = hashMap.KeySet(); _ = keySet
        var node *Node = document.CreateElement(this.BasicItemData.ITEM); _ = node
        var nameArray []any = keySet.ToArray(); _ = nameArray
        var size2 int = len(nameArray); _ = size2
        for index := int(0); index < size2; index++ {
            var name string = any(nameArray[index]).(string); _ = name
            var value string = any(hashMap.Get(name)).(string); _ = value
            value = stringUtil.GetNonNull(value)
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        var totalNode *Node = ModDomHelper_CreateNameValueNodes(document, this.BasicItemData.TOTAL, fmt.Sprint(this.ItemInterface.GetTotal())); _ = totalNode
        node.AppendChild(totalNode)
        var size int = this.Vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var domNodeInterface DomNodeInterface = any(this.Vector.Get(index)).(DomNodeInterface); _ = domNodeInterface
            var customNode *Node = domNodeInterface.ToXmlNode(document); _ = customNode
            node.AppendChild(customNode)
        }
        return node
}


func (this *BasicItemView) ToXmlDoc() *Document {
        return nil
}


func (this *BasicItemView) View() string {
        return nil
}

