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

type Basket struct {
    Items *HashMap
}

// Implements BasketInterface, *Serializable


func (this *Basket) initFields() {
    this.Items = StdUtil_GetInstance().CreateHashMap()
}


func NewBasket() *Basket {
    this := &Basket{}
    this.initFields()
    return this
}


func (this *Basket) AddItem(id string, num string) {
        if this.Items.ContainsKey(id) {
            var oldNumStr string = any(this.Items.Get(id)).(string); _ = oldNumStr
            var oldNumInt int = fmt.Sprint(oldNumStr); _ = oldNumInt
            var numInt int = fmt.Sprint(num); _ = numInt
            var newNumInt int = int(int(oldNumInt) + int(numInt)); _ = newNumInt
            this.Items.Put(id, fmt.Sprint(newNumInt))
        } else {
            this.Items.Put(id, num)
        }
}


func (this *Basket) RemoveItem(id string) {
        this.Items.Remove(id)
}


func (this *Basket) AdjustItem(id string, num string) {
        if int(fmt.Sprint(num)) > 0 {
            this.Items.Put(id, num)
        } else {
            this.Items.Remove(id)
        }
}


func (this *Basket) GetTotalWeight() string {
        var itemsAndNumberInBasket *HashMap = this.GetItems(); _ = itemsAndNumberInBasket
        var totalWeight *Money = NewMoney(); _ = totalWeight
        if itemsAndNumberInBasket.Size() > 0 {
            var items Set = itemsAndNumberInBasket.KeySet(); _ = items
            var product string; _ = product
            var weightFloat float32; _ = weightFloat
            var itemArray []any = items.ToArray(); _ = itemArray
            var size int = len(itemArray); _ = size
            for index := int(0); index < size; index++ {
                product = any(itemArray[index]).(string)
                weightFloat = float32(NewInventoryEntity().GetWeight(product))
                weightFloat = float32(float32(this.GetNumberOf(product)) * float32(weightFloat))
                totalWeight.Add(fmt.Sprint(weightFloat))
            }
        }
        return fmt.Sprint(totalWeight)
}


func (this *Basket) GetNumberOfItems() int {
        return int(this.Items.Size())
}


func (this *Basket) GetIds() Set {
        var set Set = this.Items.KeySet(); _ = set
        return set
}


func (this *Basket) GetItems() *HashMap {
        return this.Items
}


func (this *Basket) GetNumberOf(product string) int {
        var numStr string = any(this.Items.Get(product)).(string); _ = numStr
        return fmt.Sprint(numStr)
}


// throws MoneyException
func (this *Basket) GetSubTotal() *Money {
        var itemsAndNumberInBasket *HashMap = this.GetItems(); _ = itemsAndNumberInBasket
        var subTotal *Money = NewMoney(); _ = subTotal
        if itemsAndNumberInBasket.Size() > 0 {
            var items Set = itemsAndNumberInBasket.KeySet(); _ = items
            var inventoryEntityFactory *InventoryEntityFactory = InventoryEntityFactory_GetInstance(); _ = inventoryEntityFactory
            var inventoryEntity *InventoryEntity = any(inventoryEntityFactory.GetInventoryEntityInstance()).(*InventoryEntity); _ = inventoryEntity
            var itemArray []any = items.ToArray(); _ = itemArray
            var size int = len(itemArray); _ = size
            for index := int(0); index < size; index++ {
                var product string = any(itemArray[index]).(string); _ = product
                var itemInterface ItemInterface = inventoryEntity.GetItem(product); _ = itemInterface
                var productTotal *Money = itemInterface.GetPrice(); _ = productTotal
                productTotal.Multiply(int(this.GetNumberOf(product)))
                subTotal.Add(fmt.Sprint(productTotal))
            }
        }
        return subTotal
}

