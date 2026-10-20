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

type CanvasUtil struct {
}



func NewCanvasUtil() *CanvasUtil {
    this := &CanvasUtil{}
    return this
}


func CanvasUtil_GetKeyName(key int) string {
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        if key == Canvas_UP {
            return commonPhoneStrings.UP
        } else if key == Canvas_DOWN {
            return commonPhoneStrings.DOWN
        } else if key == Canvas_RIGHT {
            return commonPhoneStrings.RIGHT
        } else if key == Canvas_LEFT {
            return commonPhoneStrings.LEFT
        } else if key == Canvas_FIRE {
            return commonPhoneStrings.FIRE
        }
        return "UNKNOWN KEY"
}

