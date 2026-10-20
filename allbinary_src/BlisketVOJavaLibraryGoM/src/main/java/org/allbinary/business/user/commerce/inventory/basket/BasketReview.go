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

type BasketReview struct {
    Items IBasicArrayList
}


func (this *BasketReview) initFields() {
    this.Items = NewBasicArrayListD()
}


func NewBasketReview() *BasketReview {
    this := &BasketReview{}
    this.initFields()
    return this
}


func (this *BasketReview) IsValid() bool {
        return true
}


func (this *BasketReview) AddItem(item IItem) {
        this.Items.Add(item)
}


func (this *BasketReview) GetItems() IBasicArrayList {
        return this.Items
}


func (this *BasketReview) RemoveItem(id string) {
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            if strings.Compare(item.GetId(), id) == 0 {
                this.Items.Remove(item)
            }
        }
}


func (this *BasketReview) AdjustItem(id string, num string) {
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            if strings.Compare(item.GetId(), id) == 0 {
                item.SetNumber(num)
            }
        }
}


func (this *BasketReview) GetTotalWeight() string {
        var weightFloat float32 = 0.0; _ = weightFloat
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            weightFloat += float32(float32(item.GetWeight()))
        }
        return fmt.Sprint(float32(weightFloat))
}


func (this *BasketReview) GetNumberOfItems() int {
        return int(this.Items.Size())
}


func (this *BasketReview) GetIds() Set {
        var idSet *HashSet = NewHashSet(); _ = idSet
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            idSet.Add(item.GetId())
        }
        return any(idSet).(Set)
}


func (this *BasketReview) GetNumberOf(id string) int {
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            if strings.Compare(item.GetId(), id) == 0 {
                return int(item.GetNumber())
            }
        }
        return int(0)
}


// throws Exception
func (this *BasketReview) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(BasketData_BASKET); _ = node
        var itemArray []any = this.Items.ToArray(); _ = itemArray
        var size int = len(itemArray); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = any(itemArray[index]).(IItem); _ = item
            node.AppendChild(NewItemView(item, NewBasicArrayListD()).ToXmlNode(document))
        }
        return node
}

