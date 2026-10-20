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

type OrderStoreAdminUser struct {
    User
}



func NewOrderStoreAdminUser() *OrderStoreAdminUser {
    this := &OrderStoreAdminUser{}
        this.User = *NewUser()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewOrderStoreAdminUser(userHashMap *HashMap) *OrderStoreAdminUser {
    this := &OrderStoreAdminUser{}
        this.User = *NewUser(userHashMap)
    return this
}


func (this *OrderStoreAdminUser) ValidateSession(weblisketSession WeblisketSessionInterface) {
        this.User.ValidateSession(weblisketSession)
        this.UpdateSession(weblisketSession)
}


func (this *OrderStoreAdminUser) UpdateSession(weblisketSession WeblisketSessionInterface) {
        weblisketSession.SetStoreName(this.GetPermissions())
}

