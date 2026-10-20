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

type DirectionUtil struct {
    DirectionArray []IAngle
    DirectionArray2 []IAngle
}

var DirectionUtil_Instance *DirectionUtil = NewDirectionUtil()

func (this *DirectionUtil) initFields() {
    this.DirectionArray = make([]IAngle, 10)
    this.DirectionArray2 = make([]IAngle, 10)
}


func DirectionUtil_GetInstance() *DirectionUtil {
        return DirectionUtil_Instance
}


func NewDirectionUtil() *DirectionUtil {
    this := &DirectionUtil{}
    this.initFields()
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        this.DirectionArray[directionFactory.LEFT.GetValue()] = angleFactory.LEFT
        this.DirectionArray[directionFactory.RIGHT.GetValue()] = angleFactory.RIGHT
        this.DirectionArray[directionFactory.UP.GetValue()] = angleFactory.UP
        this.DirectionArray[directionFactory.DOWN.GetValue()] = angleFactory.DOWN
        this.DirectionArray2[directionFactory.LEFT.GetValue()] = angleFactory.LEFT
        this.DirectionArray2[directionFactory.RIGHT.GetValue()] = angleFactory.RIGHT
        this.DirectionArray2[directionFactory.UP.GetValue()] = angleFactory.UP
        this.DirectionArray2[directionFactory.DOWN.GetValue()] = angleFactory.DOWN
    return this
}


func (this *DirectionUtil) GetAngle(direction *Direction) IAngle {
        return this.DirectionArray2[direction.GetValue()]
}


func (this *DirectionUtil) GetFrameAngle(direction *Direction) IAngle {
        return this.DirectionArray[direction.GetValue()]
}

