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

type PointsUtil struct {
}

var PointsUtil_Instance *PointsUtil = NewPointsUtil()

func NewPointsUtil() *PointsUtil {
    this := &PointsUtil{}
    return this
}


func PointsUtil_GetInstance() *PointsUtil {
        return PointsUtil_Instance
}


func (this *PointsUtil) AdjustXY(points [][][]int, x int, y int) [][][]int {
        var size int = len(points); _ = size
        var newPoints [][][]int = func() [][][]int { array := make([][][]int, size); for index := range array { array[index] = make([][]int, 0); for index2 := range array[index] { array[index][index2] = make([]int, 0) } }; return array }(); _ = newPoints
        for index := int(0); index < size; index++ {
            var size2 int = len(points[index]); _ = size2
            newPoints[index] = func() [][]int { array := make([][]int, size2); for index := range array { array[index] = make([]int, 2) }; return array }()
            for index2 := int(0); index2 < size2; index2++ {
                if points[index][index2][0] != 1000 {
                    newPoints[index][index2][0] = points[index][index2][0] + x
                    newPoints[index][index2][1] = points[index][index2][1] + y
                } else {
                    newPoints[index][index2][0] = 1000
                    newPoints[index][index2][1] = 1000
                }
            }
        }
        return newPoints
}


func (this *PointsUtil) Adjust_1(points [][]int, both int) [][]int {
        return this.Adjust(points, both, both)
}

// Java overloads of PointsUtil.adjust
func (this *PointsUtil) Adjust(args ...any) [][]int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[[][]int](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.Adjust_1(Java_Arg[[][]int](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 3 && Java_IsArg[[][]int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.Adjust_2(Java_Arg[[][]int](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]))
        }
    }
    panic("No Java overload of PointsUtil.adjust for the arguments")
}


func (this *PointsUtil) Adjust_2(points [][]int, x int, y int) [][]int {
        var newPoints [][]int = func() [][]int { array := make([][]int, len(points)); for index := range array { array[index] = make([]int, 2) }; return array }(); _ = newPoints
        for index2 := int(0); index2 < len(points); index2++ {
            if points[index2][0] != 1000 {
                newPoints[index2][0] = points[index2][0] + x
                newPoints[index2][1] = points[index2][1] + y
            } else {
                newPoints[index2][0] = 1000
                newPoints[index2][1] = 1000
            }
        }
        return newPoints
}

