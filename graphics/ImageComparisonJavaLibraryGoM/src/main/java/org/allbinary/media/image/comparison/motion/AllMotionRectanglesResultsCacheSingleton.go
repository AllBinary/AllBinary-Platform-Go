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

type AllMotionRectanglesResultsCacheSingleton struct {
    LogUtil *LogUtil
}

var AllMotionRectanglesResultsCacheSingleton_CacheInterface CacheInterface = nil

func (this *AllMotionRectanglesResultsCacheSingleton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func init() {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var STATIC_BLOCK string = "Static Block"; _ = STATIC_BLOCK
        var instance string = "MotionRectanglesResultsCacheSingleton"; _ = instance
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            logUtil.Put(commonStrings.EXCEPTION, instance, STATIC_BLOCK, e)
                }
            }()
            logUtil.PutF(commonStrings.START, instance, STATIC_BLOCK)
            AllMotionRectanglesResultsCacheSingleton_CacheInterface = CacheInterfaceFactory_GetInstance(CacheTypeFactory_GetInstance().CACHE, CachePolicyFactory_GetInstance().ONE_MINUTE_FIVE_MAX)
            logUtil.PutF(commonStrings.END, instance, STATIC_BLOCK)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func NewAllMotionRectanglesResultsCacheSingleton() *AllMotionRectanglesResultsCacheSingleton {
    this := &AllMotionRectanglesResultsCacheSingleton{}
    this.initFields()
    return this
}


func AllMotionRectanglesResultsCacheSingleton_GetInstance() CacheInterface {
        return AllMotionRectanglesResultsCacheSingleton_CacheInterface
}

