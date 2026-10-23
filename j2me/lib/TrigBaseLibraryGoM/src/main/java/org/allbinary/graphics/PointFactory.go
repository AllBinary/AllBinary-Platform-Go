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
    ZERO_ZERO IGPoint
}

var PointFactory_Instance *PointFactory = NewPointFactory()

func (this *PointFactory) initFields() {
    this.ZERO_ZERO = this.CreateXYZ(0, 0, 0)
}


func PointFactory_GetInstance() *PointFactory {
        return PointFactory_Instance
}


func (this *PointFactory) Init() {
}


func NewPointFactory() *PointFactory {
    this := &PointFactory{}
    this.initFields()
    return this
}


func (this *PointFactory) CreateXY(x int, y int) IGPoint {
        return NewGPoint(x, y, 0)
}


func (this *PointFactory) CreateXYZ(x int, y int, z int) IGPoint {
        return NewGPoint(x, y, z)
}

