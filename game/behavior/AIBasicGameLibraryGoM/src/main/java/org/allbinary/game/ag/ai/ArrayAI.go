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

type ArrayAI struct {
    BasicAI
    ArtificialIntelligenceInterfaceArray []ArtificialIntelligenceInterface
}

// The Java class ArrayAI has subclasses so its variables are this interface
type IArrayAI interface {
    IBasicAI
    ArrayAI_() *ArrayAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    GetArtificialIntelligenceInterface() []ArtificialIntelligenceInterface
}

func (this *ArrayAI) ArrayAI_() *ArrayAI {
    return this
}


func (this *ArrayAI) initFields() {
    
}


func NewArrayAI(artificialIntelligenceInterface []ArtificialIntelligenceInterface, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *ArrayAI {
    this := &ArrayAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.ArtificialIntelligenceInterfaceArray = artificialIntelligenceInterface
    return this
}


// throws Exception
func (this *ArrayAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var size int = len(this.ArtificialIntelligenceInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            this.ArtificialIntelligenceInterfaceArray[index].ProcessAI(allBinaryLayerManager)
        }
}


func (this *ArrayAI) GetArtificialIntelligenceInterface() []ArtificialIntelligenceInterface {
        return this.ArtificialIntelligenceInterfaceArray
}

