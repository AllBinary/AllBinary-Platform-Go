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

type TouchButtonTurnRightResource struct {
    TouchButtonResource
}

var TouchButtonTurnRightResource_Instance ITouchButtonResource = NewTouchButtonTurnRightResource()


func TouchButtonTurnRightResource_GetInstance() ITouchButtonResource {
        return TouchButtonTurnRightResource_Instance
}


func NewTouchButtonTurnRightResource() *TouchButtonTurnRightResource {
    this := &TouchButtonTurnRightResource{}
        this.TouchButtonResource = *NewTouchButtonResource("/touch_button_turn_right_arrow_64_by_64.png", "/touch_button_turn_right_arrow_hint_64_by_64.png")
    return this
}

