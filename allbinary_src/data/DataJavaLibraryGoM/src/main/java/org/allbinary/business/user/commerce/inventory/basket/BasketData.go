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

type BasketData struct {
}

var BasketData_BASKET string = "BASKET"
var BasketData_NAME string = "NAME"
var BasketData_BASKETREVIEW string = "BASKET_REVIEW"
var BasketData_TOTALWEIGHT string = "BASKET_TOTALWEIGTH"
var BasketData_ITEMTOTAL string = "BASKET_ITEMTOTAL"
var BasketData_ITEMTOTALINBASKET string = "BASKET_ITEMTOTALINBASKET"
var BasketData_SUBTOTAL string = "BASKET_SUBTOTAL"
var BasketData_ISEMPTY string = "BASKET_ISEMPTY"
var BasketData_INSERT string = "BASKET_INSERT"
var BasketData_DELETE string = "BASKET_DELETE"
var BasketData_ADJUST string = "BASKET_ADJUST"


func NewBasketData() *BasketData {
    this := &BasketData{}
    return this
}

