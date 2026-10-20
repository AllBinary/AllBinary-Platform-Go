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

type BasicTopViewGeographicMapCellType struct {
    LogUtil *LogUtil
    Name string
    Cost int
    Types []int
}


func (this *BasicTopViewGeographicMapCellType) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func BasicTopViewGeographicMapCellType_CreateType(name string, type_ int, cost int) *BasicTopViewGeographicMapCellType {
        var types IBasicArrayList = NewBasicArrayListS(1); _ = types
        types.Add(type_)
        return NewBasicTopViewGeographicMapCellType(name, types, cost)
}


func NewBasicTopViewGeographicMapCellType(name string, types IBasicArrayList, cost int) *BasicTopViewGeographicMapCellType {
    this := &BasicTopViewGeographicMapCellType{}
    this.initFields()
        var size int = types.Size(); _ = size
        var typeArray []int = make([]int, size); _ = typeArray
        var typeAsInteger int; _ = typeAsInteger
        var type_ int; _ = type_
        for index := int(0); index < size; index++ {
            typeAsInteger = (any(types.Get(index)).(int))
            type_ = int(typeAsInteger)
            if GeographicMapCellTypeFactory_GetInstance().GetGeographicMapCellTypeArray()[type_] == nil {
                NewRaceTrackGeographicMapCellType(name, type_, cost, 0)
            } else {
            }
            typeArray[index] = type_
        }
        this.Name = name
        this.Cost = cost
        this.Types = typeArray
    return this
}


func (this *BasicTopViewGeographicMapCellType) IsType(type_ IGeographicMapCellType) bool {
        return this.HasType(type_.GetType())
}


func (this *BasicTopViewGeographicMapCellType) HasType(type_ int) bool {
        var size int = len(this.Types); _ = size
        for index := int(0); index < size; index++ {
            if this.Types[index] == type_ {
                return true
            }
        }
        return false
}


func (this *BasicTopViewGeographicMapCellType) GetTypes() []int {
        return this.Types
}


func (this *BasicTopViewGeographicMapCellType) ToString() string {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var size int = len(this.Types); _ = size
        for index := int(0); index < size; index++ {
            stringMaker.Appendint(this.Types[index]).Append(CommonSeps_GetInstance().COMMA)
        }
        return fmt.Sprint(stringMaker)
}

func (this *BasicTopViewGeographicMapCellType) String() string {
    return this.ToString()
}

