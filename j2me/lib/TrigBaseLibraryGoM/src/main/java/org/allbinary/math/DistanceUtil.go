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

type DistanceUtil struct {
    MathUtil *MathUtil
}

var DistanceUtil_Instance *DistanceUtil = NewDistanceUtil()

func (this *DistanceUtil) initFields() {
    this.MathUtil = MathUtil_GetInstance()
}

func NewDistanceUtil() *DistanceUtil {
    this := &DistanceUtil{}
    this.initFields()
    return this
}


func DistanceUtil_GetInstance() *DistanceUtil {
        return DistanceUtil_Instance
}


func (this *DistanceUtil) GetDistance(x1 int, y1 int, x2 int, y2 int) int {
        var dx int = x1 - x2; _ = dx
        var dy int = y1 - y2; _ = dy
        return this.MathUtil.Sqrt((dx * dx) + (dy * dy))
}

