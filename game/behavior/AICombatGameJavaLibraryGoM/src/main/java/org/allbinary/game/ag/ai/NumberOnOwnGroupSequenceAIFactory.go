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

type NumberOnOwnGroupSequenceAIFactory struct {
    LogUtil *LogUtil
}


func (this *NumberOnOwnGroupSequenceAIFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewNumberOnOwnGroupSequenceAIFactory() *NumberOnOwnGroupSequenceAIFactory {
    this := &NumberOnOwnGroupSequenceAIFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *NumberOnOwnGroupSequenceAIFactory) GetInstance(hashtable *ABHashtable[any, any], artificialIntelligenceInterface []ArtificialIntelligenceInterface, ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) ArtificialIntelligenceInterface {
        var integerArray []int = NumberInSameGroupSequence_GetInstance().NUMBER_ON_SAME_TEAM_SEQUENCE; _ = integerArray
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        for index := int(0); index < len(integerArray); index++ {
            var integerCanBeNull any = hashtable.Get(any(integerArray[index]).(any)); _ = integerCanBeNull
            if integerCanBeNull == nil {
                break
            }
            list.Add(integerCanBeNull)
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Total AI Properties: ")
        stringBuffer.Appendint(list.Size())
        stringBuffer.Append(" == Total AI: ")
        stringBuffer.Appendint(len(artificialIntelligenceInterface))
        stringBuffer.Append(" + 1")
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.GET_INSTANCE)
        if len(artificialIntelligenceInterface) - 1 > list.Size() {
            panic(errors.New(fmt.Sprint("Not enough AI properties.")))
        }
        if len(artificialIntelligenceInterface) - 1 < list.Size() {
            panic(errors.New(fmt.Sprint("Too Many AI properties.")))
        }
        var numberOnSameTeam []int = make([]int, list.Size()); _ = numberOnSameTeam
        for index := int(0); index < len(numberOnSameTeam); index++ {
            numberOnSameTeam[index] = any(list.BasicArrayList_().ObjectArray[index]).(int)
        }
        return NewNumberOnOwnGroupSequenceAI(numberOnSameTeam, artificialIntelligenceInterface, ownerLayerInterface, gameInput)
}

