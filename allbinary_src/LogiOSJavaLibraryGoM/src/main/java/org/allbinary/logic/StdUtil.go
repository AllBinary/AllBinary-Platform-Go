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

//ActualPlatform
type StdUtil struct {
    EMPTY_VECTOR *ABVector[any]
    NULL_TABLE *ABHashtable[any, any]
    NULL_MAP *ABHashMap[any, any]
}

var StdUtil_Instance *StdUtil = NewStdUtil()

func (this *StdUtil) initFields() {
    this.EMPTY_VECTOR = this.CreateVector()
    this.NULL_TABLE = this.CreateHashtable()
    this.NULL_MAP = this.CreateHashMap()
}

func NewStdUtil() *StdUtil {
    this := &StdUtil{}
    this.initFields()
    return this
}


func StdUtil_GetInstance() *StdUtil {
        return StdUtil_Instance
}


func (this *StdUtil) CreateStack() *ABStack[any] {
        return NewABStack[any]()
}


func (this *StdUtil) CreateVector() *ABVector[any] {
        return NewABVector[any]()
}


func (this *StdUtil) CreateHashtable() *ABHashtable[any, any] {
        return NewABHashtable[any, any]()
}


func (this *StdUtil) CreateHashMap() *ABHashMap[any, any] {
        return NewABHashMap[any, any]()
}

