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

type OrientationMotionGestureFactory struct {
    LEFT *MotionGestureInput
    RIGHT *MotionGestureInput
    UP *MotionGestureInput
    DOWN *MotionGestureInput
    ROLL_LEFT *MotionGestureInput
    ROLL_RIGHT *MotionGestureInput
}

var OrientationMotionGestureFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *OrientationMotionGestureFactory) initFields() {
    
    
    
    
    
    
}


func OrientationMotionGestureFactory_GetInstance() *OrientationMotionGestureFactory {
        if OrientationMotionGestureFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            OrientationMotionGestureFactory_Instance = NewOrientationMotionGestureFactory()
        }
        return any(OrientationMotionGestureFactory_Instance).(*OrientationMotionGestureFactory)
}


func NewOrientationMotionGestureFactory() *OrientationMotionGestureFactory {
    this := &OrientationMotionGestureFactory{}
    this.initFields()
        var MAX int = InputFactory_GetInstance().MAX; _ = MAX
        this.LEFT = NewMotionGestureInput(MAX - 20, "Left Orient")
        this.RIGHT = NewMotionGestureInput(MAX - 21, "Right Orient")
        this.UP = NewMotionGestureInput(MAX - 22, "Up Orient")
        this.DOWN = NewMotionGestureInput(MAX - 23, "Down Orient")
        this.ROLL_LEFT = NewMotionGestureInput(MAX - 24, "Roll Left")
        this.ROLL_RIGHT = NewMotionGestureInput(MAX - 25, "Roll Right")
    return this
}

