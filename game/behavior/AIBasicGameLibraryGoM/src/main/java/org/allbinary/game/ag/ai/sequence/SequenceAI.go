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

type SequenceAI struct {
    ArrayAI
    Index int
}

// The Java class SequenceAI has subclasses so its variables are this interface
type ISequenceAI interface {
    IArrayAI
    SequenceAI_() *SequenceAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    Next()
    GetIndex() int
    SetIndex(index int)
    GetSelectedArtificialIntelligenceInterface() ArtificialIntelligenceInterface
    GetName() string
    ToString() string
}

func (this *SequenceAI) SequenceAI_() *SequenceAI {
    return this
}

var SequenceAI_SEQUENCE_AI string = "Sequence AI"

func (this *SequenceAI) initFields() {
    
}


func NewSequenceAI(artificialIntelligenceInterfaceArray []ArtificialIntelligenceInterface, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *SequenceAI {
    this := &SequenceAI{}
        this.ArrayAI = *NewArrayAI(artificialIntelligenceInterfaceArray, ownerLayerInterface, gameInput)
        this.initFields()
    return this
}


// throws Exception
func (this *SequenceAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        this.GetArtificialIntelligenceInterface()[this.Index].ProcessAI(allBinaryLayerManager)
}


func (this *SequenceAI) Next() {
        this.Index++
        var artificialIntelligenceInterface ArtificialIntelligenceInterface = this.GetSelectedArtificialIntelligenceInterface(); _ = artificialIntelligenceInterface
        if artificialIntelligenceInterface.GetId() == ArtificialIntelligence_AI_ID {
            var artificialIntelligenceTransitionInterface ArtificialIntelligenceTransitionInterface = any(artificialIntelligenceInterface).(ArtificialIntelligenceTransitionInterface); _ = artificialIntelligenceTransitionInterface
            artificialIntelligenceTransitionInterface.Transition()
        }
}


func (this *SequenceAI) GetIndex() int {
        return this.Index
}


func (this *SequenceAI) SetIndex(index int) {
        this.Index = index
}


func (this *SequenceAI) GetSelectedArtificialIntelligenceInterface() ArtificialIntelligenceInterface {
        return this.GetArtificialIntelligenceInterface()[this.Index]
}


func (this *SequenceAI) GetName() string {
        return SequenceAI_SEQUENCE_AI
}


func (this *SequenceAI) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.ArrayAI.ToString())
        stringBuffer.Append(" Selected AI: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetSelectedArtificialIntelligenceInterface()))
        return fmt.Sprint(stringBuffer)
}

func (this *SequenceAI) String() string {
    return this.ToString()
}

