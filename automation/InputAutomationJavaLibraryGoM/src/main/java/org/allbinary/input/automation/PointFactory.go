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

type PointFactory struct {
}

var PointFactory_PointArray [][]*Point = func() [][]*Point { array := make([][]*Point, 3000); for index := range array { array[index] = make([]*Point, 3000) }; return array }()


func NewPointFactory() *PointFactory {
    this := &PointFactory{}
    return this
}


func PointFactory_GetInstance(x int, y int) *Point {
        var point *Point = PointFactory_PointArray[x][y]; _ = point
        if point == nil {
            PointFactory_PointArray[x][y] = NewPoint(x, y)
            point = PointFactory_PointArray[x][y]
        }
        return point
}

