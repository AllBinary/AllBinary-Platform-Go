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

type CachePolicy struct {
    Name string
    MaxTime int64
    MaxSize int64
}


func (this *CachePolicy) initFields() {
    
    
    
}


func NewCachePolicy(name string, maxTime int64, maxSize int64) *CachePolicy {
    this := &CachePolicy{}
    this.initFields()
        this.Name = name
        this.SetMaxTime(maxTime)
        this.SetMaxSize(maxSize)
    return this
}


func (this *CachePolicy) GetMaxTime() int64 {
        return this.MaxTime
}


func (this *CachePolicy) SetMaxTime(maxTime int64) {
        this.MaxTime = maxTime
}


func (this *CachePolicy) GetMaxSize() int64 {
        return this.MaxSize
}


func (this *CachePolicy) SetMaxSize(maxSize int64) {
        this.MaxSize = maxSize
}


func (this *CachePolicy) GetName() string {
        return this.Name
}

