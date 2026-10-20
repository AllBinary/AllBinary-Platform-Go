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

type TouchMotionGestureFactory struct {
    LEFT *MotionGestureInput
    RIGHT *MotionGestureInput
    UP *MotionGestureInput
    DOWN *MotionGestureInput
    PRESSED *MotionGestureInput
    RELEASED *MotionGestureInput
    DIAGONAL_DOWN_RIGHT *MotionGestureInput
    DIAGONAL_DOWN_LEFT *MotionGestureInput
    DIAGONAL_UP_RIGHT *MotionGestureInput
    DIAGONAL_UP_LEFT *MotionGestureInput
    SCROLL_UP *MotionGestureInput
    SCROLL_DOWN *MotionGestureInput
    NO_MOTION *MotionGestureInput
    LAST_MOTION *MotionGestureInput
}

var TouchMotionGestureFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *TouchMotionGestureFactory) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func TouchMotionGestureFactory_GetInstance() *TouchMotionGestureFactory {
        if TouchMotionGestureFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            TouchMotionGestureFactory_Instance = NewTouchMotionGestureFactory()
        }
        return any(TouchMotionGestureFactory_Instance).(*TouchMotionGestureFactory)
}


func NewTouchMotionGestureFactory() *TouchMotionGestureFactory {
    this := &TouchMotionGestureFactory{}
    this.initFields()
        var MAX int = InputFactory_GetInstance().MAX; _ = MAX
        this.LEFT = NewMotionGestureInput(MAX - 3, "Left Touch")
        this.RIGHT = NewMotionGestureInput(MAX - 4, "Right Touch")
        this.UP = NewMotionGestureInput(MAX - 5, "Up Touch")
        this.DOWN = NewMotionGestureInput(MAX - 6, "Down Touch")
        this.PRESSED = NewMotionGestureInput(MAX - 7, "Screen Press")
        this.RELEASED = NewMotionGestureInput(MAX - 8, "Released")
        this.DIAGONAL_DOWN_RIGHT = NewMotionGestureInput(MAX - 9, "Diagonal Dn R")
        this.DIAGONAL_DOWN_LEFT = NewMotionGestureInput(MAX - 10, "Diagonal Dn L")
        this.DIAGONAL_UP_RIGHT = NewMotionGestureInput(MAX - 11, "Diagonal Up R")
        this.DIAGONAL_UP_LEFT = NewMotionGestureInput(MAX - 12, "Diagonal Up L")
        this.SCROLL_UP = NewMotionGestureInput(MAX - 13, "Scroll Up")
        this.SCROLL_DOWN = NewMotionGestureInput(MAX - 14, "Scroll Down")
        this.NO_MOTION = NewMotionGestureInput(MAX - 2, "No Motion")
        this.LAST_MOTION = NewMotionGestureInput(MAX - 19, "Last Motion")
    return this
}

