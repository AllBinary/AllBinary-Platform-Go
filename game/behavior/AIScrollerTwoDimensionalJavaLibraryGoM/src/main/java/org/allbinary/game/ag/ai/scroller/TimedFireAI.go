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

type TimedFireAI struct {
    BasicAI
    MaxFireDelayTimeHelper ITimeDelayHelper
    Delay int
}

var TimedFireAI_TIME int = SmallIntegerSingletonFactory_GetInstance().GetAt(1)

func (this *TimedFireAI) initFields() {
    this.MaxFireDelayTimeHelper = NewTimeDelayHelper(0)
    
}


func NewTimedFireAI(delay int, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *TimedFireAI {
    this := &TimedFireAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.Delay = delay
    return this
}


// throws Exception
func (this *TimedFireAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var timeFiredInterface TimeFiredInterface = any(this.GetOwnerLayerInterface()).(TimeFiredInterface); _ = timeFiredInterface
        if this.MaxFireDelayTimeHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) && timeFiredInterface.GetLastFireTime() + int64(this.MaxFireDelayTimeHelper.TimeDelayHelper_().Delay) < this.MaxFireDelayTimeHelper.GetStartTime() {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM1)
            this.MaxFireDelayTimeHelper.TimeDelayHelper_().Delay = this.Delay
        }
}

