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

type GeographicMapCellTypeFactory struct {
    LogUtil *LogUtil
    GeographicMapCellTypeArray []IGeographicMapCellType
}

// The Java class GeographicMapCellTypeFactory has subclasses so its variables are this interface
type IGeographicMapCellTypeFactory interface {
    GeographicMapCellTypeFactory_() *GeographicMapCellTypeFactory
    Get(type_ int) IGeographicMapCellType
    GetGeographicMapCellTypeArray() []IGeographicMapCellType
    GetStartType() int
    GetEndType() int
    GetEmptyType() int
    IsPath(cellType IGeographicMapCellType) bool
}

func (this *GeographicMapCellTypeFactory) GeographicMapCellTypeFactory_() *GeographicMapCellTypeFactory {
    return this
}

var GeographicMapCellTypeFactory_Instance IGeographicMapCellTypeFactory = NewGeographicMapCellTypeFactory()

func (this *GeographicMapCellTypeFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GeographicMapCellTypeArray = make([]IGeographicMapCellType, 512)
}

func NewGeographicMapCellTypeFactory() *GeographicMapCellTypeFactory {
    this := &GeographicMapCellTypeFactory{}
    this.initFields()
    return this
}


func GeographicMapCellTypeFactory_GetInstance() IGeographicMapCellTypeFactory {
        return GeographicMapCellTypeFactory_Instance
}


func (this *GeographicMapCellTypeFactory) Get(type_ int) IGeographicMapCellType {
        return this.GeographicMapCellTypeArray[type_]
}


func (this *GeographicMapCellTypeFactory) GetGeographicMapCellTypeArray() []IGeographicMapCellType {
        return this.GeographicMapCellTypeArray
}


func (this *GeographicMapCellTypeFactory) GetStartType() int {
        panic(errors.New(fmt.Sprint()))
}


func (this *GeographicMapCellTypeFactory) GetEndType() int {
        panic(errors.New(fmt.Sprint()))
}


func (this *GeographicMapCellTypeFactory) GetEmptyType() int {
        panic(errors.New(fmt.Sprint()))
}


func (this *GeographicMapCellTypeFactory) IsPath(cellType IGeographicMapCellType) bool {
        panic(errors.New(fmt.Sprint()))
}

