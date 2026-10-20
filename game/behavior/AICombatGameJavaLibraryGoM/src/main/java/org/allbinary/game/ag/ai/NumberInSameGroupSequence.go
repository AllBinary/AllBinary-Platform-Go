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

type NumberInSameGroupSequence struct {
    NUMBER_ON_SAME_TEAM_SEQUENCE_ZERO int
    NUMBER_ON_SAME_TEAM_SEQUENCE_ONE int
    NUMBER_ON_SAME_TEAM_SEQUENCE_TWO int
    NUMBER_ON_SAME_TEAM_SEQUENCE_THREE int
    NUMBER_ON_SAME_TEAM_SEQUENCE_FOUR int
    NUMBER_ON_SAME_TEAM_SEQUENCE_FIVE int
    NUMBER_ON_SAME_TEAM_SEQUENCE_SIX int
    NUMBER_ON_SAME_TEAM_SEQUENCE_SEVEN int
    NUMBER_ON_SAME_TEAM_SEQUENCE_EIGHT int
    NUMBER_ON_SAME_TEAM_SEQUENCE_NINE int
    NUMBER_ON_SAME_TEAM_SEQUENCE []int
}


func (this *NumberInSameGroupSequence) initFields() {
    
    
    
    
    
    
    
    
    
    
    
}


func NumberInSameGroupSequence_GetInstance() *NumberInSameGroupSequence {
        return NewNumberInSameGroupSequence()
}


func NewNumberInSameGroupSequence() *NumberInSameGroupSequence {
    this := &NumberInSameGroupSequence{}
    this.initFields()
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_ZERO = smallIntegerSingletonFactory.GetAt(100)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_ONE = smallIntegerSingletonFactory.GetAt(101)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_TWO = smallIntegerSingletonFactory.GetAt(102)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_THREE = smallIntegerSingletonFactory.GetAt(103)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_FOUR = smallIntegerSingletonFactory.GetAt(104)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_FIVE = smallIntegerSingletonFactory.GetAt(105)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_SIX = smallIntegerSingletonFactory.GetAt(106)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_SEVEN = smallIntegerSingletonFactory.GetAt(107)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_EIGHT = smallIntegerSingletonFactory.GetAt(108)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE_NINE = smallIntegerSingletonFactory.GetAt(109)
        this.NUMBER_ON_SAME_TEAM_SEQUENCE = []int{this.NUMBER_ON_SAME_TEAM_SEQUENCE_ZERO, this.NUMBER_ON_SAME_TEAM_SEQUENCE_ONE, this.NUMBER_ON_SAME_TEAM_SEQUENCE_THREE, this.NUMBER_ON_SAME_TEAM_SEQUENCE_FOUR, this.NUMBER_ON_SAME_TEAM_SEQUENCE_FIVE, this.NUMBER_ON_SAME_TEAM_SEQUENCE_SIX, this.NUMBER_ON_SAME_TEAM_SEQUENCE_SEVEN, this.NUMBER_ON_SAME_TEAM_SEQUENCE_EIGHT, this.NUMBER_ON_SAME_TEAM_SEQUENCE_NINE}
    return this
}

