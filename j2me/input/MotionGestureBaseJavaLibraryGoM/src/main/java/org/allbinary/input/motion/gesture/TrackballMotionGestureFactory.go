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

type TrackballMotionGestureFactory struct {
    LEFT *MotionGestureInput
    RIGHT *MotionGestureInput
    UP *MotionGestureInput
    DOWN *MotionGestureInput
}

var TrackballMotionGestureFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *TrackballMotionGestureFactory) initFields() {
    
    
    
    
}


func TrackballMotionGestureFactory_GetInstance() *TrackballMotionGestureFactory {
        if TrackballMotionGestureFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            TrackballMotionGestureFactory_Instance = NewTrackballMotionGestureFactory()
        }
        return any(TrackballMotionGestureFactory_Instance).(*TrackballMotionGestureFactory)
}


func NewTrackballMotionGestureFactory() *TrackballMotionGestureFactory {
    this := &TrackballMotionGestureFactory{}
    this.initFields()
        var MAX int = InputFactory_GetInstance().MAX; _ = MAX
        this.LEFT = NewMotionGestureInput(MAX - 26, "L Trackball")
        this.RIGHT = NewMotionGestureInput(MAX - 27, "R Trackball")
        this.UP = NewMotionGestureInput(MAX - 28, "Up Trackball")
        this.DOWN = NewMotionGestureInput(MAX - 29, "D Trackball")
    return this
}

