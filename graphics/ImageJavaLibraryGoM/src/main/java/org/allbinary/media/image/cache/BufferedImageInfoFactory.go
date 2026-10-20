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

type BufferedImageInfoFactory struct {
}

var BufferedImageInfoFactory_HashMap *HashMap = StdUtil_GetInstance().CreateHashMap()


func NewBufferedImageInfoFactory() *BufferedImageInfoFactory {
    this := &BufferedImageInfoFactory{}
    return this
}


func BufferedImageInfoFactory_GetInstance(width int, height int, type_ int) *BufferedImageInfo {
        var stringbuffer *StringMaker = NewStringMaker(); _ = stringbuffer
        stringbuffer.Appendint(width)
        stringbuffer.Appendint(height)
        stringbuffer.Appendint(type_)
        var key string = fmt.Sprint(stringbuffer); _ = key
        var bufferedImageInfo *BufferedImageInfo = any(BufferedImageInfoFactory_HashMap.Get(key)).(*BufferedImageInfo); _ = bufferedImageInfo
        if bufferedImageInfo == nil {
            bufferedImageInfo = NewBufferedImageInfo(width, height, type_)
            BufferedImageInfoFactory_HashMap.Put(key, bufferedImageInfo)
        }
        return bufferedImageInfo
}

