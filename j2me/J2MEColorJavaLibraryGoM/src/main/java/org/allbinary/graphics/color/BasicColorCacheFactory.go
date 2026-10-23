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

type BasicColorCacheFactory struct {
    Hashtable *ABHashtable[any, any]
}

var BasicColorCacheFactory_Instance *BasicColorCacheFactory = NewBasicColorCacheFactory()

func (this *BasicColorCacheFactory) initFields() {
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
}


func BasicColorCacheFactory_GetInstance() *BasicColorCacheFactory {
        return BasicColorCacheFactory_Instance
}


func NewBasicColorCacheFactory() *BasicColorCacheFactory {
    this := &BasicColorCacheFactory{}
    this.initFields()
    return this
}


func (this *BasicColorCacheFactory) Add(basicDefaultColor *BasicColor) {
        this.Hashtable.Put(int(basicDefaultColor.IntValue()), basicDefaultColor)
}


// synchronized
func (this *BasicColorCacheFactory) GetAndOrCreate(integer int) *BasicColor {
        var basicColor *BasicColor = any(this.Hashtable.Get(integer)).(*BasicColor); _ = basicColor
        if basicColor == nil {
            basicColor = BasicColorFactory_GetInstance().CreateInstance(int(integer))
            this.Hashtable.Put(integer, basicColor)
        }
        return basicColor
}

