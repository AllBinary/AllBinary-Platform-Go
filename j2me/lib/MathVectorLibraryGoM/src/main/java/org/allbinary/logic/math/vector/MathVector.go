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

type MathVector struct {
    MagnitudeBasicDecimal *BasicDecimal
    Direction int
    ResultBasicDecimal *BasicDecimal
}

// Implements MathVectorInterface

// The Java class MathVector has subclasses so its variables are this interface
type IMathVector interface {
    MathVector_() *MathVector
    GetMagnitude() *BasicDecimal
    GetDirection() int
    GetResult() *BasicDecimal
}

func (this *MathVector) MathVector_() *MathVector {
    return this
}


func (this *MathVector) initFields() {
    
    
    
}


func NewMathVector(magnitudeBasicDecimal *BasicDecimal, direction int) *MathVector {
    this := &MathVector{}
    this.initFields()
        this.MagnitudeBasicDecimal = magnitudeBasicDecimal
        this.Direction = direction
    return this
}


func (this *MathVector) GetMagnitude() *BasicDecimal {
        return this.MagnitudeBasicDecimal
}


func (this *MathVector) GetDirection() int {
        return this.Direction
}


func (this *MathVector) GetResult() *BasicDecimal {
        return this.ResultBasicDecimal
}

