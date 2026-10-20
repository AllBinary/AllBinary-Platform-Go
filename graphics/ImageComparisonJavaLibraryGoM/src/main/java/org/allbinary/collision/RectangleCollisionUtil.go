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



func NewRectangleCollisionUtil() *RectangleCollisionUtil {
    this := &RectangleCollisionUtil{}
    return this
}


func RectangleCollisionUtil_IsCollision_1(rectangle *Rectangle, point IGPoint) bool {
        if point.GetX() > (int(rectangle.X) + rectangle.GetWidth()) || point.GetY() > (int(rectangle.Y) + rectangle.GetHeight()) || point.GetX() < rectangle.X || point.GetY() < rectangle.Y {
            return false
        } else {
            return true
        }
}

// Java overloads of RectangleCollisionUtil.isCollision
func RectangleCollisionUtil_IsCollision(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*Rectangle](args[0], pass) && Java_IsArg[IGPoint](args[1], pass) {
            return RectangleCollisionUtil_IsCollision_1(Java_Arg[*Rectangle](args[0]), Java_Arg[IGPoint](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*Rectangle](args[0], pass) && Java_IsArg[*Rectangle](args[1], pass) {
            return RectangleCollisionUtil_IsCollision_2(Java_Arg[*Rectangle](args[0]), Java_Arg[*Rectangle](args[1]))
        }
        if len(args) == 8 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) && Java_IsArg[int](args[4], pass) && Java_IsArg[int](args[5], pass) && Java_IsArg[int](args[6], pass) && Java_IsArg[int](args[7], pass) {
            return RectangleCollisionUtil_IsCollision_3(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]), Java_Arg[int](args[4]), Java_Arg[int](args[5]), Java_Arg[int](args[6]), Java_Arg[int](args[7]))
        }
    }
    panic("No Java overload of RectangleCollisionUtil.isCollision for the arguments")
}


func RectangleCollisionUtil_IsCollision_2(rectangle1 *Rectangle, rectangle2 *Rectangle) bool {
        return RectangleCollisionUtil_IsCollision(rectangle1.X, rectangle1.Y, rectangle1.X + rectangle1.Width, rectangle1.Y + rectangle1.Height, rectangle2.X, rectangle2.Y, rectangle2.X + rectangle2.Width, rectangle2.Y + rectangle2.Height)
}


func RectangleCollisionUtil_IsCollision_3(rectX1 int, rectY1 int, rectX2 int, rectY2 int, rect2X1 int, rect2Y1 int, rect2X2 int, rect2Y2 int) bool {
        if rect2X1 >= rectX2 || rect2Y1 >= rectY2 || rect2X2 <= rectX1 || rect2Y2 <= rectY1 {
            return false
        } else {
            return true
        }
}

