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

type BasicPaymentType struct {
    Name string
    Value string
    PaymentGatewayInterfaceFactoryInterface PaymentGatewayInterfaceFactoryInterface
    PaymentGatewayDomNodeFactoryInterface PaymentGatewayDomNodeFactoryInterface
}


func (this *BasicPaymentType) initFields() {
    
    
    
    
}




// Java overloads of new BasicPaymentType
func NewBasicPaymentType(args ...any) *BasicPaymentType {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicPaymentType_1(Java_Arg[string](args[0]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[PaymentGatewayInterfaceFactoryInterface](args[1], pass) && Java_IsArg[PaymentGatewayDomNodeFactoryInterface](args[2], pass) {
            return NewBasicPaymentType_2(Java_Arg[string](args[0]), Java_Arg[PaymentGatewayInterfaceFactoryInterface](args[1]), Java_Arg[PaymentGatewayDomNodeFactoryInterface](args[2]))
        }
    }
    panic("No Java overload of new BasicPaymentType for the arguments")
}
func NewBasicPaymentType_1(paymentMethod string) *BasicPaymentType {
    this := &BasicPaymentType{}
    this.initFields()
        this.Init(paymentMethod)
    return this
}


func NewBasicPaymentType_2(paymentMethod string, paymentGatewayInterfaceFactoryInterface PaymentGatewayInterfaceFactoryInterface, paymentGatewayDomNodeFactoryInterface PaymentGatewayDomNodeFactoryInterface) *BasicPaymentType {
    this := &BasicPaymentType{}
    this.initFields()
        this.Init(paymentMethod)
        this.PaymentGatewayInterfaceFactoryInterface = paymentGatewayInterfaceFactoryInterface
        this.PaymentGatewayDomNodeFactoryInterface = paymentGatewayDomNodeFactoryInterface
    return this
}


func (this *BasicPaymentType) Init(paymentMethod string) {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var UNDERSCORE string = "_"; _ = UNDERSCORE
        hashMap.Put(CommonSeps_GetInstance().SPACE, UNDERSCORE)
        hashMap.Put(AbPathData_GetInstance().EXTENSION_SEP, UNDERSCORE)
        hashMap.Put("-", UNDERSCORE)
        var replace *Replace = NewReplace(hashMap); _ = replace
        this.Name = replace.All(paymentMethod)
        this.Value = paymentMethod
        BasicPaymentTypeUtil_GetInstance().Add(this)
}


func (this *BasicPaymentType) GetName() string {
        return this.Name
}


func (this *BasicPaymentType) GetValue() string {
        return this.Value
}


// throws Exception
func (this *BasicPaymentType) GetPaymentGatewayInterfaceFactoryInterface() PaymentGatewayInterfaceFactoryInterface {
        if this.PaymentGatewayInterfaceFactoryInterface == nil {
            panic(errors.New(fmt.Sprint("PaymentGatewayInterfaceFactoryInterface is Null")))
        }
        return this.PaymentGatewayInterfaceFactoryInterface
}


// throws Exception
func (this *BasicPaymentType) GetPaymentGatewayDomNodeFactoryInterface() PaymentGatewayDomNodeFactoryInterface {
        if this.PaymentGatewayDomNodeFactoryInterface == nil {
            panic(errors.New(fmt.Sprint("PaymentGatewayDomNodeFactoryInterface is Null")))
        }
        return this.PaymentGatewayDomNodeFactoryInterface
}


func (this *BasicPaymentType) GetKey() any {
        return this.GetName()
}


func (this *BasicPaymentType) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.Name)
        vector.Add(this.Value)
        return vector
}


func (this *BasicPaymentType) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(fmt.Sprint(PaymentGatewayData_NAME), this.Name)
        hashMap.Put(fmt.Sprint(PaymentGatewayData_VALUE), this.Value)
        return hashMap
}


func (this *BasicPaymentType) ToString() string {
        return this.Name
}

func (this *BasicPaymentType) String() string {
    return this.ToString()
}

