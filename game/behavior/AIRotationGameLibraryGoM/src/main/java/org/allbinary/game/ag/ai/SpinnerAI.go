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

type SpinnerAI struct {
    BasicAI
    IsEven bool
    Direction []int
}


func (this *SpinnerAI) initFields() {
    
    
}


func NewSpinnerAI(direction []int, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *SpinnerAI {
    this := &SpinnerAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.Direction = direction
    return this
}


// throws Exception
func (this *SpinnerAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.IsEven {
            this.BasicAI.ProcessKeyAI(this.Direction[0])
            this.IsEven = false
        } else {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM1)
            this.IsEven = true
        }
}

