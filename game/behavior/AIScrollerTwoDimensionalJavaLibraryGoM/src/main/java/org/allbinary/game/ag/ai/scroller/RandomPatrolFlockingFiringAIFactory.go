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

type RandomPatrolFlockingFiringAIFactory struct {
}

// Implements ArtificialIntelligenceInterfaceFactoryInterface


func NewRandomPatrolFlockingFiringAIFactory() *RandomPatrolFlockingFiringAIFactory {
    this := &RandomPatrolFlockingFiringAIFactory{}
    return this
}


// throws Exception
func (this *RandomPatrolFlockingFiringAIFactory) GetInstance(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) ArtificialIntelligenceInterface {
        var artificialIntelligenceInterface []ArtificialIntelligenceInterface = make([]ArtificialIntelligenceInterface, 2); _ = artificialIntelligenceInterface
        artificialIntelligenceInterface[0] = NewRandomPatrolAI(hashtable, ownerLayerInterface, gameInput)
        artificialIntelligenceInterface[1] = NewTimedFireAI(1800, ownerLayerInterface, gameInput)
        return NewArrayAI(artificialIntelligenceInterface, ownerLayerInterface, gameInput)
}

