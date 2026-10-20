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

type BasicPlatormGeographicMapCellType struct {
    Types []int
}


func (this *BasicPlatormGeographicMapCellType) initFields() {
    
}


func BasicPlatormGeographicMapCellType_CreateType(type_ int) *BasicPlatormGeographicMapCellType {
        var types IBasicArrayList = NewBasicArrayListS(1); _ = types
        types.Add(SmallIntegerSingletonFactory_GetInstance().GetAt(type_))
        var basicPlatormGeographicMapCellType *BasicPlatormGeographicMapCellType = NewBasicPlatormGeographicMapCellType(types); _ = basicPlatormGeographicMapCellType
        return basicPlatormGeographicMapCellType
}


func NewBasicPlatormGeographicMapCellType(types IBasicArrayList) *BasicPlatormGeographicMapCellType {
    this := &BasicPlatormGeographicMapCellType{}
    this.initFields()
        var size int = types.Size(); _ = size
        var typeArray []int = make([]int, size); _ = typeArray
        var typeAsInteger int; _ = typeAsInteger
        var type_ int; _ = type_
        for index := int(0); index < size; index++ {
            typeAsInteger = (any(types.Get(index)).(int))
            type_ = int(typeAsInteger)
            NewGeographicMapCellType(type_, 0)
            typeArray[index] = type_
        }
        this.Types = typeArray
    return this
}


func (this *BasicPlatormGeographicMapCellType) IsType(type_ IGeographicMapCellType) bool {
        return this.HasType(type_.GetType())
}


func (this *BasicPlatormGeographicMapCellType) HasType(type_ int) bool {
        var size int = len(this.Types); _ = size
        for index := int(0); index < size; index++ {
            if this.Types[index] == type_ {
                return true
            }
        }
        return false
}

