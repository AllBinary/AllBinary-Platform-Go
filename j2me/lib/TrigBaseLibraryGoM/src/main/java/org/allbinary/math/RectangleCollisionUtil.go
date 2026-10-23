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

type RectangleCollisionUtil struct {
}

var RectangleCollisionUtil_Instance *RectangleCollisionUtil = NewRectangleCollisionUtil()


func RectangleCollisionUtil_GetInstance() *RectangleCollisionUtil {
        return RectangleCollisionUtil_Instance
}


func NewRectangleCollisionUtil() *RectangleCollisionUtil {
    this := &RectangleCollisionUtil{}
    return this
}


func (this *RectangleCollisionUtil) IsCollision(rectX1 int, rectY1 int, rectX2 int, rectY2 int, rect2X1 int, rect2Y1 int, rect2X2 int, rect2Y2 int) bool {
        if rect2X1 >= rectX2 || rect2Y1 >= rectY2 || rect2X2 <= rectX1 || rect2Y2 <= rectY1 {
            return false
        } else {
            return true
        }
}


func (this *RectangleCollisionUtil) IsInside(rectX1 int, rectY1 int, rectX2 int, rectY2 int, x int, y int) bool {
        if x >= rectX2 || y >= rectY2 || x <= rectX1 || y <= rectY1 {
            return false
        } else {
            return true
        }
}

