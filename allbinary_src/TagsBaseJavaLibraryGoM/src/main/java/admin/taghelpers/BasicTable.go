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

type BasicTable struct {
    CommonStrings *CommonStrings
}

// Implements BasicTableInterface

// The Java class BasicTable has subclasses so its variables are this interface
type IBasicTable interface {
    BasicTable_() *BasicTable
    Drop() string
    Create() string
    Restore() string
    Backup() string
}

func (this *BasicTable) BasicTable_() *BasicTable {
    return this
}


func (this *BasicTable) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewBasicTable() *BasicTable {
    this := &BasicTable{}
    this.initFields()
    return this
}


// throws Exception
func (this *BasicTable) Drop() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *BasicTable) Create() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *BasicTable) Restore() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *BasicTable) Backup() string {
        panic(errors.New(fmt.Sprint()))
}

