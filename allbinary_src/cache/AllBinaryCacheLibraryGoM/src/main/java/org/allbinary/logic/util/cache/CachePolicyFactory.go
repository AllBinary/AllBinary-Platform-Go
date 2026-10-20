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

type CachePolicyFactory struct {
    NAME string
    MAX_TIME_THOUSAND_MAX *CachePolicy
    THIRTY_MINUTES_TEN_THOUSAND_MAX *CachePolicy
    THIRTY_MINUTES_FIFTY_THOUSAND_MAX *CachePolicy
    ONE_MINUTE_FIVE_MAX *CachePolicy
    ONE_MINUTE_TEN_MAX *CachePolicy
    ONE_MINUTE_ONE_HUNDRED_MAX *CachePolicy
}

var CachePolicyFactory_Instance *CachePolicyFactory = NewCachePolicyFactory()

func (this *CachePolicyFactory) initFields() {
    this.NAME = "lru"
    this.MAX_TIME_THOUSAND_MAX = NewCachePolicy(this.NAME, Long_MAX_VALUE, 1000)
    this.THIRTY_MINUTES_TEN_THOUSAND_MAX = NewCachePolicy(this.NAME, 1000 * 60 * 30, 10000)
    this.THIRTY_MINUTES_FIFTY_THOUSAND_MAX = NewCachePolicy(this.NAME, 1000 * 60 * 30, 100000)
    this.ONE_MINUTE_FIVE_MAX = NewCachePolicy(this.NAME, 1000 * 60 * 1, 5)
    this.ONE_MINUTE_TEN_MAX = NewCachePolicy(this.NAME, 1000 * 60 * 1, 10)
    this.ONE_MINUTE_ONE_HUNDRED_MAX = NewCachePolicy(this.NAME, 1000 * 60 * 1, 100)
}

func NewCachePolicyFactory() *CachePolicyFactory {
    this := &CachePolicyFactory{}
    this.initFields()
    return this
}


func CachePolicyFactory_GetInstance() *CachePolicyFactory {
        return CachePolicyFactory_Instance
}

