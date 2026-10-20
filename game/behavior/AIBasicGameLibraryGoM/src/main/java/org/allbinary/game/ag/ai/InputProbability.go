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

type InputProbability struct {
    LikelyhoodIntegerArray [][]int
    RepeatLikelyhoodIntegerArray []int
    Max int
    TimeDelayHelper ITimeDelayHelper
}

var InputProbability_INPUT_PROBABILITY int = SmallIntegerSingletonFactory_GetInstance().GetAt(0)

func (this *InputProbability) initFields() {
    
    
    
    
}


func NewInputProbability(likelyhoodIntegerArray [][]int, repeatLikelyhoodIntegerArray []int, max_ int, timeDelayHelper ITimeDelayHelper) *InputProbability {
    this := &InputProbability{}
    this.initFields()
        this.LikelyhoodIntegerArray = likelyhoodIntegerArray
        this.RepeatLikelyhoodIntegerArray = repeatLikelyhoodIntegerArray
        this.Max = max_
        this.TimeDelayHelper = timeDelayHelper
    return this
}


func (this *InputProbability) GetLikelyhoodIntegerArray() [][]int {
        return this.LikelyhoodIntegerArray
}


func (this *InputProbability) GetRepeatLikelyhoodIntegerArray() []int {
        return this.RepeatLikelyhoodIntegerArray
}


func (this *InputProbability) GetMax() int {
        return this.Max
}


func (this *InputProbability) GetTimeDelayHelper() ITimeDelayHelper {
        return this.TimeDelayHelper
}


func (this *InputProbability) ToHashtable() *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        hashtable.Put(InputProbability_INPUT_PROBABILITY, this)
        return hashtable
}

