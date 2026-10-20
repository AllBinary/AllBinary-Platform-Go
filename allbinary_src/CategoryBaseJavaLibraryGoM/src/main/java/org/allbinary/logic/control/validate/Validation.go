/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type Validation struct {
    CommonStrings *CommonStrings
}

// Implements ValidationInterface

// The Java class Validation has subclasses so its variables are this interface
type IValidation interface {
    Validation_() *Validation
    IsValid() bool
    ToValidationInfoNode(document *Document) *Node
    ToValidationInfoDoc() *Document
    ValidationInfo() string
}

func (this *Validation) Validation_() *Validation {
    return this
}


func (this *Validation) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewValidation() *Validation {
    this := &Validation{}
    this.initFields()
    return this
}


// throws Exception
func (this *Validation) IsValid() bool {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Validation) ToValidationInfoNode(document *Document) *Node {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Validation) ToValidationInfoDoc() *Document {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Validation) ValidationInfo() string {
        panic(errors.New(fmt.Sprint()))
}

