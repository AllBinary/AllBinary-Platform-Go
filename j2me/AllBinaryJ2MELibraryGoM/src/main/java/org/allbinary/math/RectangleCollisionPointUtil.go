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

type RectangleCollisionPointUtil struct {
}



func NewRectangleCollisionPointUtil() *RectangleCollisionPointUtil {
    this := &RectangleCollisionPointUtil{}
    return this
}


func RectangleCollisionPointUtil_AllPointsInside(rectX1 int, rectY1 int, rectX2 int, rectY2 int, list IBasicArrayList) bool {
        var rectangleCollisionUtil *RectangleCollisionUtil = RectangleCollisionUtil_GetInstance(); _ = rectangleCollisionUtil
        var point IGPoint; _ = point
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            point = any(list.BasicArrayList_().ObjectArray[index]).(IGPoint)
            if !rectangleCollisionUtil.IsInside(rectX1, rectY1, rectX2, rectY2, point.GetX(), point.GetY()) {
                return false
            }
        }
        return true
}


func RectangleCollisionPointUtil_AllPointsInsideCell(rectX1 int, rectY1 int, rectX2 int, rectY2 int, list IBasicArrayList, xCellSize int, yCellSize int) bool {
        var rectangleCollisionUtil *RectangleCollisionUtil = RectangleCollisionUtil_GetInstance(); _ = rectangleCollisionUtil
        var point IGPoint; _ = point
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            point = any(list.BasicArrayList_().ObjectArray[index]).(IGPoint)
            if !rectangleCollisionUtil.IsInside(rectX1, rectY1, rectX2, rectY2, point.GetX() * xCellSize, point.GetY() * yCellSize) {
                return false
            }
        }
        return true
}

