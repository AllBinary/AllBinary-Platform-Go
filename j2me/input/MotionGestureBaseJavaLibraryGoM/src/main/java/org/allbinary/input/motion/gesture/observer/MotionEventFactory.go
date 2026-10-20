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

type MotionEventFactory struct {
    Source any
    Id int
}

// Implements AllBinaryEventObjectFactoryInterface

var MotionEventFactory_Index int = 0

func (this *MotionEventFactory) initFields() {
    
    
}


func NewMotionEventFactory(source any, id int) *MotionEventFactory {
    this := &MotionEventFactory{}
    this.initFields()
        this.Source = source
        this.Id = id
        MotionEventFactory_Index = TouchMotionGestureFactory_GetInstance().LAST_MOTION.GetId()
    return this
}


func (this *MotionEventFactory) GetInstance() IAllBinaryEventObject {
        var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
        var input IInput = inputFactory.GetInstanceById(Java_PostIncrement(&MotionEventFactory_Index)); _ = input
        if input == inputFactory.NO_INPUT {
            return NewMotionGestureEvent(this.Source, this.Id, any(TouchMotionGestureFactory_GetInstance().NO_MOTION).(*MotionGestureInput))
        } else {
            return NewMotionGestureEvent(this.Source, this.Id, any(input).(*MotionGestureInput))
        }
}

