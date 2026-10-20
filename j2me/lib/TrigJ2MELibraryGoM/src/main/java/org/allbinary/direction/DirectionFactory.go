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

type DirectionFactory struct {
    NAME string
    LEFT *Direction
    RIGHT *Direction
    DOWN *Direction
    UP *Direction
    UP_LEFT *Direction
    UP_RIGHT *Direction
    DOWN_LEFT *Direction
    DOWN_RIGHT *Direction
    NOT_BORDERED_WITH *Direction
    NO_DIRECTION *Direction
    TOP *Direction
    BOTTOM *Direction
    ALL_DIRECTION *Direction
}

var DirectionFactory_Instance *DirectionFactory = NewDirectionFactory()

func (this *DirectionFactory) initFields() {
    this.NAME = "DIRECTION"
    this.LEFT = NewDirection(CommonPhoneStrings_GetInstance().LEFT, "WEST", 0, 1)
    this.RIGHT = NewDirection(CommonPhoneStrings_GetInstance().RIGHT, "EAST", 1, 0)
    this.DOWN = NewDirection(CommonPhoneStrings_GetInstance().DOWN, "SOUTH", 2, 0)
    this.UP = NewDirection(CommonPhoneStrings_GetInstance().UP, "NORTH", 3, 0)
    
    
    
    
    this.NOT_BORDERED_WITH = NewDirection("NOT_BORDERED_WITH", "NOT_BORDERED_WITH", 8, 0)
    this.NO_DIRECTION = NewDirection("NO_DIRECTION", "NO_DIRECTION", 9, 0)
    this.TOP = NewDirection("TOP", "TOP", 10, 0)
    this.BOTTOM = NewDirection("BOTTOM", "BOTTOM", 11, 0)
    this.ALL_DIRECTION = NewDirection("ALL_DIRECTION", "ALL_DIRECTION", 12, 0)
}


func DirectionFactory_GetInstance() *DirectionFactory {
        return DirectionFactory_Instance
}


func NewDirectionFactory() *DirectionFactory {
    this := &DirectionFactory{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        this.UP_LEFT = NewDirection("UP LEFT", fmt.Sprint(stringBuffer.Append(this.UP.GetOtherName()).Append(commonSeps.SPACE).Append(this.LEFT.GetOtherName())), 4, 0)
        stringBuffer.Delete(0, stringBuffer.Length())
        this.UP_RIGHT = NewDirection("UP RIGHT", fmt.Sprint(stringBuffer.Append(this.UP.GetOtherName()).Append(commonSeps.SPACE).Append(this.RIGHT.GetOtherName())), 5, 0)
        stringBuffer.Delete(0, stringBuffer.Length())
        this.DOWN_LEFT = NewDirection("DOWN LEFT", fmt.Sprint(stringBuffer.Append(this.DOWN.GetOtherName()).Append(commonSeps.SPACE).Append(this.LEFT.GetOtherName())), 6, 0)
        stringBuffer.Delete(0, stringBuffer.Length())
        this.DOWN_RIGHT = NewDirection("DOWN RIGHT", fmt.Sprint(stringBuffer.Append(this.DOWN.GetOtherName()).Append(commonSeps.SPACE).Append(this.RIGHT.GetOtherName())), 7, 0)
    return this
}

