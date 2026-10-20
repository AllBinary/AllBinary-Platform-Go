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

type BasicRandomAI struct {
    BasicAI
    I_random int
    InputProbability *InputProbability
    Visitor IVisitor
    KeyArray []int
    MyRandomFactory *MyRandomFactory
}


func (this *BasicRandomAI) initFields() {
    this.I_random = 0
    
    
    this.KeyArray = []int{Canvas_UP, Canvas_DOWN, Canvas_LEFT, Canvas_RIGHT, Canvas_KEY_NUM1, Canvas_KEY_NUM7, Canvas_KEY_NUM9}
    this.MyRandomFactory = MyRandomFactory_GetInstance()
}


func NewBasicRandomAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput, inputProbability *InputProbability, visitor IVisitor) *BasicRandomAI {
    this := &BasicRandomAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.InputProbability = inputProbability
        this.Visitor = visitor
    return this
}


// throws Exception
func (this *BasicRandomAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.InputProbability.GetTimeDelayHelper().IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            this.I_random = this.MyRandomFactory.GetAbsoluteNextInt(this.InputProbability.GetMax())
            var repeat bool = false; _ = repeat
            var repeatLikelyhoodIntegerArray []int = this.InputProbability.GetRepeatLikelyhoodIntegerArray(); _ = repeatLikelyhoodIntegerArray
            if this.GetLastKey() != -1 && int(repeatLikelyhoodIntegerArray[this.GetLastKey()]) >= this.I_random {
                repeat = true
            }
            if !repeat {
                var likelyhoodIntegerArray [][]int = this.InputProbability.GetLikelyhoodIntegerArray(); _ = likelyhoodIntegerArray
                this.I_random = this.MyRandomFactory.GetAbsoluteNextInt(this.InputProbability.GetMax())
                var size int = len(this.KeyArray); _ = size
                var likelyhoodIntegerKeyArray []int; _ = likelyhoodIntegerKeyArray
                for index := int(0); index < size; index++ {
                    var key int = this.KeyArray[index]; _ = key
                    likelyhoodIntegerKeyArray = likelyhoodIntegerArray[key]
                    if len(likelyhoodIntegerArray) >= key && likelyhoodIntegerKeyArray[0] != 0 {
                        if this.I_random >= int(likelyhoodIntegerKeyArray[0]) && this.I_random < int(likelyhoodIntegerKeyArray[1]) {
                            this.SetLastKey(key)
                            break
                        }
                    }
                }
            }
            this.Visitor.Visit(this)
        }
}

