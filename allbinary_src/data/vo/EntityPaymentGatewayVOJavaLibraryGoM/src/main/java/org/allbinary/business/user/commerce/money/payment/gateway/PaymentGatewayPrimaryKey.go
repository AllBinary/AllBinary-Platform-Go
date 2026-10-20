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

type PaymentGatewayPrimaryKey struct {
    StoreName string
    Name string
}

// Implements *Serializable


func (this *PaymentGatewayPrimaryKey) initFields() {
    
    
}


func NewPaymentGatewayPrimaryKey(storeName string, gatewayName string) *PaymentGatewayPrimaryKey {
    this := &PaymentGatewayPrimaryKey{}
    this.initFields()
        this.SetStoreName(storeName)
        this.SetName(gatewayName)
    return this
}


func (this *PaymentGatewayPrimaryKey) GetStoreName() string {
        return this.StoreName
}


func (this *PaymentGatewayPrimaryKey) SetStoreName(storeName string) {
        this.StoreName = storeName
}


func (this *PaymentGatewayPrimaryKey) GetName() string {
        return this.Name
}


func (this *PaymentGatewayPrimaryKey) SetName(name string) {
        this.Name = name
}


func (this *PaymentGatewayPrimaryKey) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf((*PaymentGatewayPrimaryKey)(nil)).Elem().GetName())
        stringBuffer.Append(reflect.TypeOf((*StoreFront)(nil)).Elem().GetName())
        stringBuffer.Append(" =")
        stringBuffer.Append(this.GetStoreName())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(reflect.TypeOf((*PaymentGateway)(nil)).Elem().GetName())
        stringBuffer.Append("=")
        stringBuffer.Append(this.GetName())
        return fmt.Sprint(stringBuffer)
}

func (this *PaymentGatewayPrimaryKey) String() string {
    return this.ToString()
}

