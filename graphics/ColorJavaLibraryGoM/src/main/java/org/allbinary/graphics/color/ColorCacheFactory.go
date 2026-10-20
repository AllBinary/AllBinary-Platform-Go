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

type ColorCacheFactory struct {
}

var ColorCacheFactory_CacheInterface AutomaticCacheInterface = nil


func init() {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var STATIC_BLOCK string = "Static Block"; _ = STATIC_BLOCK
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            logUtil.Put(commonStrings.EXCEPTION, reflect.TypeOf((*ColorCacheFactory)(nil)).Elem(), STATIC_BLOCK, e)
                }
            }()
            logUtil.PutF(commonStrings.START, reflect.TypeOf((*ColorCacheFactory)(nil)).Elem(), STATIC_BLOCK)
            ColorCacheFactory_CacheInterface = AutomaticCacheInterfaceFactory_GetInstance(NewColorCacheableFactory(), CacheTypeFactory_GetInstance().CACHE, CachePolicyFactory_GetInstance().THIRTY_MINUTES_TEN_THOUSAND_MAX)
            logUtil.PutF(commonStrings.END, reflect.TypeOf((*ColorCacheFactory)(nil)).Elem(), STATIC_BLOCK)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func NewColorCacheFactory() *ColorCacheFactory {
    this := &ColorCacheFactory{}
    return this
}


func ColorCacheFactory_GetInstance() AutomaticCacheInterface {
        return ColorCacheFactory_CacheInterface
}

