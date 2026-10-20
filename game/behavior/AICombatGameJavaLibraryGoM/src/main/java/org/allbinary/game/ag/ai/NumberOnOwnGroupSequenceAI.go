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

type NumberOnOwnGroupSequenceAI struct {
    SequenceAI
    NumberOnSameTeam []int
}


func (this *NumberOnOwnGroupSequenceAI) initFields() {
    
}


func NewNumberOnOwnGroupSequenceAI(numberOnSameTeam []int, artificialIntelligenceInterface []ArtificialIntelligenceInterface, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *NumberOnOwnGroupSequenceAI {
    this := &NumberOnOwnGroupSequenceAI{}
        this.SequenceAI = *NewSequenceAI(artificialIntelligenceInterface, ownerLayerInterface, gameInput)
        this.initFields()
        this.NumberOnSameTeam = numberOnSameTeam
    return this
}


// throws Exception
func (this *NumberOnOwnGroupSequenceAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var layerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = layerInterface
        var index int = this.GetIndex(); _ = index
        if len(this.NumberOnSameTeam) > index {
            var size int = GroupLayerManagerListener_GetInstance().GetGroupSizeFromInterface(layerInterface); _ = size
            if int(this.NumberOnSameTeam[index]) > size {
                this.Next()
            }
        }
        this.SequenceAI.ProcessAI(allBinaryLayerManager)
}

