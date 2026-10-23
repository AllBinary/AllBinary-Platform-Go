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

type BasicDecimal struct {
    Units int64
    Scaled int
    Factor int
    FactorValue int
}

var BasicDecimal_ZERO_BIGDECIMAL *BasicDecimal = NewBasicDecimal(0)

func (this *BasicDecimal) initFields() {
    
    
    this.Factor = ScaleFactorFactory_GetInstance().DEFAULT_SCALE_FACTOR
    this.FactorValue = ScaleFactorFactory_GetInstance().DEFAULT_SCALE_VALUE
}


func BasicDecimal_CreateBigDecimal(bigDecimal *BasicDecimal) *BasicDecimal {
        return NewBasicDecimal(bigDecimal.GetUnscaled())
}


func NewBasicDecimal(units int64) *BasicDecimal {
    this := &BasicDecimal{}
    this.initFields()
        this.Units = units
        this.UpdateScaled()
    return this
}


func (this *BasicDecimal) Setint(value int) {
        this.Units = int64(value)
        this.UpdateScaled()
}


func (this *BasicDecimal) Setlong(value int64) {
        this.Units = value
        this.UpdateScaled()
}


func (this *BasicDecimal) Set(bigDecimal *BasicDecimal) {
        this.Units = bigDecimal.GetUnscaled()
        this.UpdateScaled()
}


func (this *BasicDecimal) Addint(value int) {
        this.Units = this.Units + int64(value)
        this.UpdateScaled()
}


func (this *BasicDecimal) Addlong(value int64) {
        this.Units = this.Units + value
        this.UpdateScaled()
}


func (this *BasicDecimal) Add(bigDecimal *BasicDecimal) {
        this.Units = this.Units + bigDecimal.GetUnscaled()
        this.UpdateScaled()
}


func (this *BasicDecimal) Subtractint(value int) {
        this.Units = this.Units - int64(value)
        this.UpdateScaled()
}


func (this *BasicDecimal) Subtractlong(value int64) {
        this.Units = this.Units - value
        this.UpdateScaled()
}


func (this *BasicDecimal) Subtract(bigDecimal *BasicDecimal) {
        this.Units = this.Units - bigDecimal.GetUnscaled()
        this.UpdateScaled()
}


func (this *BasicDecimal) Multiplyint(value int) {
        this.Units = this.Units * int64(value)
        this.UpdateScaled()
}


func (this *BasicDecimal) Multiplylong(value int64) {
        this.Units = this.Units * value
        this.UpdateScaled()
}


func (this *BasicDecimal) Multiply(bigDecimal *BasicDecimal) {
        this.Units = this.Units * bigDecimal.GetUnscaled()
        this.UpdateScaled()
}


func (this *BasicDecimal) Divideint(value int) {
        this.Units = this.Units / int64(value)
        this.UpdateScaled()
}


func (this *BasicDecimal) Dividelong(value int64) {
        this.Units = this.Units / value
        this.UpdateScaled()
}


func (this *BasicDecimal) Divide(bigDecimal *BasicDecimal) {
        this.Units = this.Units / bigDecimal.GetUnscaled()
        this.UpdateScaled()
}


func (this *BasicDecimal) GetUnscaled() int64 {
        return this.Units
}


func (this *BasicDecimal) UpdateScaled() {
        this.Scaled = int((this.Units / int64(this.FactorValue)))
}


func (this *BasicDecimal) GetScaled() int {
        return this.Scaled
}


func (this *BasicDecimal) ToString() string {
        return fmt.Sprint(this.Units)
}

func (this *BasicDecimal) String() string {
    return this.ToString()
}


func (this *BasicDecimal) GetScaledFactor() int {
        return this.Factor
}


func (this *BasicDecimal) GetScaledFactorValue() int {
        return this.FactorValue
}

