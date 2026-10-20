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

package allbinary

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

type ScoreComparator struct {
    LogUtil *LogUtil
    IsHighestBest bool
}

// Implements RecordComparator


func (this *ScoreComparator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewScoreComparator(isHighestBest bool) *ScoreComparator {
    this := &ScoreComparator{}
    this.initFields()
        this.IsHighestBest = isHighestBest
    return this
}


func (this *ScoreComparator) GetBestScore() int64 {
        return func() any { if this.IsHighestBest { return Long_MAX_VALUE }; return 0 }()
}


func (this *ScoreComparator) Compare(recordOne []int8, recordTwo []int8) int {
        var byteArrayInputStreamOne *ByteArrayInputStream = NewByteArrayInputStream(recordOne); _ = byteArrayInputStreamOne
        var inputStreamOne IDataInputStream = NewDataInputStream(byteArrayInputStreamOne); _ = inputStreamOne
        var byteArrayInputStreamTwo *ByteArrayInputStream = NewByteArrayInputStream(recordTwo); _ = byteArrayInputStreamTwo
        var inputStreamTwo IDataInputStream = NewDataInputStream(byteArrayInputStreamTwo); _ = inputStreamTwo
        var scoreOne int64 = 0; _ = scoreOne
        var scoreTwo int64 = 0; _ = scoreTwo
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "compare", e)
                }
            }()
            inputStreamOne.ReadUTF()
            inputStreamTwo.ReadUTF()
            scoreOne = inputStreamOne.ReadLong()
            scoreTwo = inputStreamTwo.ReadLong()
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        if this.IsHighestBest {
            return this.GetHighTooLow(scoreOne, scoreTwo)
        } else {
            return this.GetLowTooHigh(scoreOne, scoreTwo)
        }
}


func (this *ScoreComparator) GetHighTooLow(scoreOne int64, scoreTwo int64) int {
        if scoreOne > scoreTwo {
            return RecordComparator_PRECEDES
        } else if scoreOne < scoreTwo {
            return RecordComparator_FOLLOWS
        } else {
            return RecordComparator_EQUIVALENT
        }
}


func (this *ScoreComparator) GetLowTooHigh(scoreOne int64, scoreTwo int64) int {
        if scoreOne < scoreTwo {
            return RecordComparator_PRECEDES
        } else if scoreOne > scoreTwo {
            return RecordComparator_FOLLOWS
        } else {
            return RecordComparator_EQUIVALENT
        }
}

