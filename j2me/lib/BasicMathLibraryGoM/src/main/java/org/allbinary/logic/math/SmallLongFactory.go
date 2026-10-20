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

type SmallLongFactory struct {
    NEGATIVE_MAX int
    POSITIVE_MAX int
    INTEGER_ARRAY []int64
    STRING_ARRAY []string
    MIN int
    LastMin int
    LastNegativeMin int
}

var SmallLongFactory_Instance *SmallLongFactory = NewSmallLongFactory()

func (this *SmallLongFactory) initFields() {
    this.NEGATIVE_MAX = 500
    this.POSITIVE_MAX = 0x2D1
    this.INTEGER_ARRAY = make([]int64, this.NEGATIVE_MAX + this.POSITIVE_MAX)
    this.STRING_ARRAY = make([]string, this.NEGATIVE_MAX + this.POSITIVE_MAX)
    this.MIN = 0
    this.LastMin = 0
    this.LastNegativeMin = 0
}


func SmallLongFactory_GetInstance() *SmallLongFactory {
        return SmallLongFactory_Instance
}


func (this *SmallLongFactory) InitWithRange(value int, negativeValue int) {
        for index := int(value - 1); index >= this.LastMin; index-- {
            this.INTEGER_ARRAY[index + this.NEGATIVE_MAX] = int64(index)
        }
        for index := int(negativeValue - 1); index >= this.LastNegativeMin; index-- {
            this.INTEGER_ARRAY[index] = int64(-index)
        }
        if this.LastMin < value {
            this.LastMin = value
            this.MIN = value
        }
        if this.LastNegativeMin < negativeValue {
            this.LastNegativeMin = negativeValue
        }
}


func (this *SmallLongFactory) Init() {
        if this.INTEGER_ARRAY[0] == 0 {
            if this.LastMin < this.POSITIVE_MAX || this.LastNegativeMin < this.NEGATIVE_MAX {
                for index := int(this.POSITIVE_MAX - 1); index >= this.LastMin; index-- {
                    this.INTEGER_ARRAY[index + this.NEGATIVE_MAX] = int64(index)
                }
                for index := int(this.NEGATIVE_MAX - 1); index >= this.LastNegativeMin; index-- {
                    this.INTEGER_ARRAY[index] = int64(-index)
                }
                this.LastMin = this.POSITIVE_MAX
                this.LastNegativeMin = this.NEGATIVE_MAX
            }
        }
}


func NewSmallLongFactory() *SmallLongFactory {
    this := &SmallLongFactory{}
    this.initFields()
    return this
}


func (this *SmallLongFactory) GetInstanceNoThrow(index int64) int64 {
        if index + int64(this.NEGATIVE_MAX) > int64(len(this.INTEGER_ARRAY) - 1) {
            return this.INTEGER_ARRAY[-1 + this.NEGATIVE_MAX]
        }
        return this.INTEGER_ARRAY[(int(index)) + this.NEGATIVE_MAX]
}


func (this *SmallLongFactory) GetString(index int) string {
        var i int = index + this.NEGATIVE_MAX; _ = i
        if this.STRING_ARRAY[i] == "" {
            this.STRING_ARRAY[i] = fmt.Sprint(this.INTEGER_ARRAY[i])
        }
        return this.STRING_ARRAY[i]
}
