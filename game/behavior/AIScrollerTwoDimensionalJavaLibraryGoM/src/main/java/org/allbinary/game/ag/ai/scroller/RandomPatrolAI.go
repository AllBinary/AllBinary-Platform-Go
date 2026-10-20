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

type RandomPatrolAI struct {
    BasePatrolAI
    ChangedDirection bool
}


func (this *RandomPatrolAI) initFields() {
    this.ChangedDirection = true
}


func NewRandomPatrolAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *RandomPatrolAI {
    this := &RandomPatrolAI{}
        this.BasePatrolAI = *NewBasePatrolAI(hashtable, ownerLayerInterface, gameInput)
        this.initFields()
        this.UpdateRandomDistance()
    return this
}


func (this *RandomPatrolAI) UpdateRandomDistance() {
        this.CurrentDistance = MyRandomFactory_GetInstance().GetAbsoluteNextInt(int(this.MaxDistance))
}


// throws Exception
func (this *RandomPatrolAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        this.BasePatrolAI.Update()
        this.BasePatrolAI.ProcessKeyAI(this.KeyDirection)
        if this.ChangedDirection {
            this.BasePatrolAI.ProcessKeyAI(Canvas_KEY_NUM1)
            this.UpdateRandomDistance()
            this.ChangedDirection = false
        }
}


func (this *RandomPatrolAI) NextDirection() {
        this.BasePatrolAI.NextDirection()
        this.ChangedDirection = true
}

