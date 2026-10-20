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

type InApplicationPurchaseFactory struct {
}

var InApplicationPurchaseFactory_Instance *InApplicationPurchaseFactory = NewInApplicationPurchaseFactory()


func InApplicationPurchaseFactory_GetInstance() *InApplicationPurchaseFactory {
        return InApplicationPurchaseFactory_Instance
}


func NewInApplicationPurchaseFactory() *InApplicationPurchaseFactory {
    this := &InApplicationPurchaseFactory{}
    return this
}


func (this *InApplicationPurchaseFactory) Init(anyType any, stateObject any) {
}


func (this *InApplicationPurchaseFactory) OnCreate(anyType any, stateObject any) {
}


func (this *InApplicationPurchaseFactory) OnStart() {
}


func (this *InApplicationPurchaseFactory) OnResult(requestCode int, resultCode int, data any) {
}


func (this *InApplicationPurchaseFactory) OnSaveState(anyType any) {
}


func (this *InApplicationPurchaseFactory) OnStop() {
}


func (this *InApplicationPurchaseFactory) OnDestroy() {
}


func (this *InApplicationPurchaseFactory) Purchase(lockableFeature *LockableFeature) {
}


func (this *InApplicationPurchaseFactory) IsPurchased(lockableFeature *LockableFeature) bool {
        return false
}


func (this *InApplicationPurchaseFactory) Add(lockableFeature *LockableFeature) {
}


func (this *InApplicationPurchaseFactory) GetUserName() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *InApplicationPurchaseFactory) IsEnabled() bool {
        return false
}

