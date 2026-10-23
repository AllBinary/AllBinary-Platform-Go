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

type Table struct {
    CommonStrings *CommonStrings
}

// Implements TableInterface

// The Java class Table has subclasses so its variables are this interface
type ITable interface {
    Table_() *Table
    Drop() string
    Create() string
    Restore() string
    Backup() string
    Insert() string
    Delete() string
    Update() string
}

func (this *Table) Table_() *Table {
    return this
}


func (this *Table) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewTable() *Table {
    this := &Table{}
    this.initFields()
    return this
}


// throws Exception
func (this *Table) Drop() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Create() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Restore() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Backup() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Insert() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Delete() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *Table) Update() string {
        panic(errors.New(fmt.Sprint()))
}

