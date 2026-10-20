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

type GeographicMapCellType struct {
    LogUtil *LogUtil
    Type int
}

// The Java class GeographicMapCellType has subclasses so its variables are this interface
type IGeographicMapCellType interface {
    GeographicMapCellType_() *GeographicMapCellType
    GetType() int
    ToString() string
}

func (this *GeographicMapCellType) GeographicMapCellType_() *GeographicMapCellType {
    return this
}

var GeographicMapCellType_NULL_GEOGRAPHIC_MAP_CELL_TYPE_ARRAY []IGeographicMapCellType = make([]IGeographicMapCellType, 0)
var GeographicMapCellType_NULL_GEOGRAPHIC_MAP_CELL_TYPE IGeographicMapCellType = NewGeographicMapCellType(Integer_MIN_VALUE, 0)

func (this *GeographicMapCellType) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewGeographicMapCellType(type_ int, reset int) *GeographicMapCellType {
    this := &GeographicMapCellType{}
    this.initFields()
        this.Type = type_
        if type_ != Integer_MIN_VALUE {
            var geographicMapCellTypeFactory IGeographicMapCellTypeFactory = GeographicMapCellTypeFactory_GetInstance(); _ = geographicMapCellTypeFactory
            var geographicMapCellTypeArray []IGeographicMapCellType = geographicMapCellTypeFactory.GetGeographicMapCellTypeArray(); _ = geographicMapCellTypeArray
            if geographicMapCellTypeArray[type_] == nil || reset == 1 {
                geographicMapCellTypeArray[type_] = this
            } else {
                if reset != -1 {
                    panic(errors.New(fmt.Sprint()))
                }
            }
        }
    return this
}


func (this *GeographicMapCellType) GetType() int {
        return this.Type
}


func (this *GeographicMapCellType) ToString() string {
        return fmt.Sprint(SmallIntegerSingletonFactory_GetInstance().GetAt(this.Type))
}

func (this *GeographicMapCellType) String() string {
    return this.ToString()
}

