/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type CustomMapGeneratorBase struct {
}

// The Java class CustomMapGeneratorBase has subclasses so its variables are this interface
type ICustomMapGeneratorBase interface {
    CustomMapGeneratorBase_() *CustomMapGeneratorBase
    CopyMapIntoCustomMap()
    GetCustomMapArray() [][]int
    GetCustomType(column int, row int, currentType int) int
}

func (this *CustomMapGeneratorBase) CustomMapGeneratorBase_() *CustomMapGeneratorBase {
    return this
}


func NewCustomMapGeneratorBase() *CustomMapGeneratorBase {
    this := &CustomMapGeneratorBase{}
    return this
}


// throws Exception
func (this *CustomMapGeneratorBase) CopyMapIntoCustomMap() {
}


func (this *CustomMapGeneratorBase) GetCustomMapArray() [][]int {
        return NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY
}


// throws Exception
func (this *CustomMapGeneratorBase) GetCustomType(column int, row int, currentType int) int {
        return -1
}

