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

type DiveAndDirectionalTrackingAIFactory struct {
}

// Implements ArtificialIntelligenceInterfaceFactoryInterface


func NewDiveAndDirectionalTrackingAIFactory() *DiveAndDirectionalTrackingAIFactory {
    this := &DiveAndDirectionalTrackingAIFactory{}
    return this
}


// throws Exception
func (this *DiveAndDirectionalTrackingAIFactory) GetInstance(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) ArtificialIntelligenceInterface {
        var visitorCanBeNull any = hashtable.Get(any(BasicAI_AI_VISITOR).(any)); _ = visitorCanBeNull
        if visitorCanBeNull == nil {
            visitorCanBeNull = ThrustAIVisitorFactory_GetInstance()
        }
        var hashtable2 *ABHashtable[any, any] = NewBasicProbabilityAIDataFactory().GetInstance(); _ = hashtable2
        hashtable2.Put(BasicAI_AI_VISITOR, LastKeyAIVisitorFactory_GetInstance())
        var artificialIntelligenceInterface ArtificialIntelligenceInterface = NewBasicRandomAIFactory().GetInstance(hashtable2, ownerLayerInterface, gameInput); _ = artificialIntelligenceInterface
        return NewDiveAndDirectionalTrackingAI(ownerLayerInterface, artificialIntelligenceInterface, gameInput, any(visitorCanBeNull).(IVisitor))
}

