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

type BaseCircularPool struct {
    CircularIndexUtil *CircularIndexUtil
    OBJECT_ARRAY []any
}

// The Java class BaseCircularPool has subclasses so its variables are this interface
type IBaseCircularPool interface {
    BaseCircularPool_() *BaseCircularPool
    GetNextInstance() any
    Init(allBinaryObjectFactoryInterface AllBinaryObjectFactoryInterface)
    GetInstance(index int) any
}

func (this *BaseCircularPool) BaseCircularPool_() *BaseCircularPool {
    return this
}


func (this *BaseCircularPool) initFields() {
    this.CircularIndexUtil = CircularIndexUtil_NULL_CIRCULAR_INDEX_UTIL
    this.OBJECT_ARRAY = NullUtil_GetInstance().NULL_OBJECT_ARRAY
}

func NewBaseCircularPool() *BaseCircularPool {
    this := &BaseCircularPool{}
    this.initFields()
    return this
}


// throws Exception
// synchronized
func (this *BaseCircularPool) GetNextInstance() any {
        var anyType any = this.OBJECT_ARRAY[this.CircularIndexUtil.GetIndex()]; _ = anyType
        this.CircularIndexUtil.Next()
        return anyType
}


func (this *BaseCircularPool) Init(allBinaryObjectFactoryInterface AllBinaryObjectFactoryInterface) {
        this.CircularIndexUtil.SetIndex(0)
        var size int = this.CircularIndexUtil.GetSize(); _ = size
        for index := int(0); index < size; index++ {
            this.OBJECT_ARRAY[index] = allBinaryObjectFactoryInterface.GetInstance()
        }
}


// throws Exception
func (this *BaseCircularPool) GetInstance(index int) any {
        return this.OBJECT_ARRAY[index]
}

