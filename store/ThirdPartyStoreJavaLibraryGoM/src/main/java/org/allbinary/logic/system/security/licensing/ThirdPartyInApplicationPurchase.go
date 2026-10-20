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

type ThirdPartyInApplicationPurchase struct {
}

// Implements ThirdPartyInApplicationPurchaseInterface

// The Java class ThirdPartyInApplicationPurchase has subclasses so its variables are this interface
type IThirdPartyInApplicationPurchase interface {
    ThirdPartyInApplicationPurchase_() *ThirdPartyInApplicationPurchase
    IsPurchased(lockableFeature *LockableFeature) bool
    OnCreate(anyType any, stateObject any)
    OnDestroy()
    OnStart()
    OnActivityResult(requestCode int, resultCode int, dataObject any)
    OnSaveInstanceState(anyType any)
    OnStop()
    RequestPurchase(lockableFeature *LockableFeature)
    RequestReceipts()
    Add(lockableFeature *LockableFeature)
    GetUserName() string
    IsEnabled() bool
    RequestProducts()
}

func (this *ThirdPartyInApplicationPurchase) ThirdPartyInApplicationPurchase_() *ThirdPartyInApplicationPurchase {
    return this
}


func NewThirdPartyInApplicationPurchase() *ThirdPartyInApplicationPurchase {
    this := &ThirdPartyInApplicationPurchase{}
    return this
}


func (this *ThirdPartyInApplicationPurchase) IsPurchased(lockableFeature *LockableFeature) bool {
        return false
}


func (this *ThirdPartyInApplicationPurchase) OnCreate(anyType any, stateObject any) {
}


func (this *ThirdPartyInApplicationPurchase) OnDestroy() {
}


func (this *ThirdPartyInApplicationPurchase) OnStart() {
}


func (this *ThirdPartyInApplicationPurchase) OnActivityResult(requestCode int, resultCode int, dataObject any) {
}


func (this *ThirdPartyInApplicationPurchase) OnSaveInstanceState(anyType any) {
}


func (this *ThirdPartyInApplicationPurchase) OnStop() {
}


func (this *ThirdPartyInApplicationPurchase) RequestPurchase(lockableFeature *LockableFeature) {
}


func (this *ThirdPartyInApplicationPurchase) RequestReceipts() {
}


func (this *ThirdPartyInApplicationPurchase) Add(lockableFeature *LockableFeature) {
}


func (this *ThirdPartyInApplicationPurchase) GetUserName() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *ThirdPartyInApplicationPurchase) IsEnabled() bool {
        return false
}


func (this *ThirdPartyInApplicationPurchase) RequestProducts() {
}

