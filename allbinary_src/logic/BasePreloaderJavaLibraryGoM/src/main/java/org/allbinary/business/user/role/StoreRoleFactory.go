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

type StoreRoleFactory struct {
    Roles IBasicArrayList
    SubscriberRoles IBasicArrayList
    WholesaleRoles IBasicArrayList
}

var StoreRoleFactory_Instance *StoreRoleFactory = NewStoreRoleFactory()

func (this *StoreRoleFactory) initFields() {
    this.Roles = NewBasicArrayListD()
    this.SubscriberRoles = NewBasicArrayListD()
    this.WholesaleRoles = NewBasicArrayListD()
}


func StoreRoleFactory_GetInstance() *StoreRoleFactory {
        return StoreRoleFactory_Instance
}


func NewStoreRoleFactory() *StoreRoleFactory {
    this := &StoreRoleFactory{}
    this.initFields()
        var basicUserRoleFactory *BasicUserRoleFactory = BasicUserRoleFactory_GetInstance(); _ = basicUserRoleFactory
        this.Roles.Add(basicUserRoleFactory.CUSTOMER)
        this.Roles.Add(basicUserRoleFactory.SUBSCRIBERCUSTOMER)
        this.Roles.Add(basicUserRoleFactory.WHOLESALECUSTOMER)
        this.SubscriberRoles.Add(basicUserRoleFactory.SUBSCRIBERCUSTOMER)
        this.WholesaleRoles.Add(basicUserRoleFactory.WHOLESALECUSTOMER)
    return this
}


func (this *StoreRoleFactory) GetRoles() IBasicArrayList {
        return this.Roles
}


func (this *StoreRoleFactory) GetSubscriberRoles() IBasicArrayList {
        return this.SubscriberRoles
}


func (this *StoreRoleFactory) GetWholesaleRoles() IBasicArrayList {
        return this.WholesaleRoles
}

