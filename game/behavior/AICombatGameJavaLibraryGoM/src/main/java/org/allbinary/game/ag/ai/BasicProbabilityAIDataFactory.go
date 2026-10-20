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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type BasicProbabilityAIDataFactory struct {
}


func NewBasicProbabilityAIDataFactory() *BasicProbabilityAIDataFactory {
    this := &BasicProbabilityAIDataFactory{}
    return this
}


func (this *BasicProbabilityAIDataFactory) GetInstance() *ABHashtable[any, any] {
        var likelyhoodIntegerArray [][]int = func() [][]int { array := make([][]int, Canvas_KEY_NUM9 + 1); for index := range array { array[index] = make([]int, 2) }; return array }(); _ = likelyhoodIntegerArray
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        likelyhoodIntegerArray[Canvas_UP][0] = smallIntegerSingletonFactory.GetAt(0)
        likelyhoodIntegerArray[Canvas_UP][1] = smallIntegerSingletonFactory.GetAt(10)
        likelyhoodIntegerArray[Canvas_DOWN][0] = smallIntegerSingletonFactory.GetAt(10)
        likelyhoodIntegerArray[Canvas_DOWN][1] = smallIntegerSingletonFactory.GetAt(14)
        likelyhoodIntegerArray[Canvas_LEFT][0] = smallIntegerSingletonFactory.GetAt(14)
        likelyhoodIntegerArray[Canvas_LEFT][1] = smallIntegerSingletonFactory.GetAt(18)
        likelyhoodIntegerArray[Canvas_RIGHT][0] = smallIntegerSingletonFactory.GetAt(18)
        likelyhoodIntegerArray[Canvas_RIGHT][1] = smallIntegerSingletonFactory.GetAt(22)
        likelyhoodIntegerArray[Canvas_KEY_NUM7][0] = smallIntegerSingletonFactory.GetAt(22)
        likelyhoodIntegerArray[Canvas_KEY_NUM7][1] = smallIntegerSingletonFactory.GetAt(26)
        likelyhoodIntegerArray[Canvas_KEY_NUM9][0] = smallIntegerSingletonFactory.GetAt(26)
        likelyhoodIntegerArray[Canvas_KEY_NUM9][1] = smallIntegerSingletonFactory.GetAt(30)
        likelyhoodIntegerArray[Canvas_KEY_NUM1][0] = smallIntegerSingletonFactory.GetAt(30)
        likelyhoodIntegerArray[Canvas_KEY_NUM1][1] = smallIntegerSingletonFactory.GetAt(40)
        var repeatLikelyhoodIntegerArray []int = make([]int, Canvas_KEY_NUM9 + 1); _ = repeatLikelyhoodIntegerArray
        repeatLikelyhoodIntegerArray[Canvas_UP] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_DOWN] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_LEFT] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_RIGHT] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_KEY_NUM7] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_KEY_NUM9] = smallIntegerSingletonFactory.GetAt(0)
        repeatLikelyhoodIntegerArray[Canvas_KEY_NUM1] = smallIntegerSingletonFactory.GetAt(10)
        var inputProbability *InputProbability = NewInputProbability(likelyhoodIntegerArray, repeatLikelyhoodIntegerArray, 42, NoTimeDelayHelper_SINGLETON); _ = inputProbability
        return inputProbability.ToHashtable()
}

