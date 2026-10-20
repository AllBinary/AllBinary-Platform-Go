/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type LayerUtil struct {
}

var LayerUtil_HASHTABLE string = "Hashtable: "

func NewLayerUtil() *LayerUtil {
    this := &LayerUtil{}
    return this
}


func LayerUtil_ToString(hashtable *ABHashtable[any, any], x int, y int, z int) string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(LayerUtil_HASHTABLE)
        stringBuffer.Append(StringUtil_GetInstance().ToString(hashtable))
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Appendint(x)
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Appendint(y)
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Appendint(z)
        return fmt.Sprint(stringBuffer)
}

