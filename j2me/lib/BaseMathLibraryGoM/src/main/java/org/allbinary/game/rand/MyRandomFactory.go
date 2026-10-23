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

//Do not use in client code for money related things.
type MyRandomFactory struct {
    MathUtil *MathUtil
    Rand *Random
}

var MyRandomFactory_Instance *MyRandomFactory = NewMyRandomFactory()

func (this *MyRandomFactory) initFields() {
    this.MathUtil = MathUtil_GetInstance()
    
}


func MyRandomFactory_GetInstance() *MyRandomFactory {
        return MyRandomFactory_Instance
}


func NewMyRandomFactory() *MyRandomFactory {
    this := &MyRandomFactory{}
    this.initFields()
        this.Rand = NewRandom(time.Now().UnixMilli())
    return this
}


func (this *MyRandomFactory) SetSeed(seed int64) {
        this.Rand = NewRandom(seed)
}


func (this *MyRandomFactory) GetNextInt(range_ int) int {
        var div int = (Integer_MAX_VALUE / range_) + 1; _ = div
        return (this.Rand.NextInt() / div)
}


func (this *MyRandomFactory) GetAbsoluteNextInt(range_ int) int {
        return this.MathUtil.Abs(this.GetNextInt(range_))
}


func (this *MyRandomFactory) GetAbsoluteNextIntAllowZero(range_ int) int {
        if range_ == 0 {
            return 0
        } else {
            return this.MathUtil.Abs(this.GetNextInt(range_))
        }
}


func (this *MyRandomFactory) Shuffle(intArray []int) {
        this.ShuffleTotal(intArray, len(intArray) * 7)
}


func (this *MyRandomFactory) ShuffleTotal(intArray []int, shuffleTotal int) {
        var size int = len(intArray); _ = size
        var randomIndex int; _ = randomIndex
        var randomIndex2 int; _ = randomIndex2
        var value int; _ = value
        for index := int(0); index < shuffleTotal; index++ {
            randomIndex = this.GetAbsoluteNextIntAllowZero(size)
            randomIndex2 = this.GetAbsoluteNextIntAllowZero(size)
            value = intArray[randomIndex]
            intArray[randomIndex] = intArray[randomIndex2]
            intArray[randomIndex2] = value
        }
}


func (this *MyRandomFactory) Shuffle2(intArray []int, intArray2 []int) {
        this.Shuffle2Total(intArray, intArray2, len(intArray) * 7)
}


func (this *MyRandomFactory) Shuffle2Total(intArray []int, intArray2 []int, shuffleTotal int) {
        var size int = len(intArray); _ = size
        var randomIndex int; _ = randomIndex
        var randomIndex2 int; _ = randomIndex2
        var value int; _ = value
        var value2 int; _ = value2
        for index := int(0); index < shuffleTotal; index++ {
            randomIndex = this.GetAbsoluteNextIntAllowZero(size)
            randomIndex2 = this.GetAbsoluteNextIntAllowZero(size)
            value = intArray[randomIndex]
            value2 = intArray2[randomIndex]
            intArray[randomIndex] = intArray[randomIndex2]
            intArray[randomIndex2] = value
            intArray2[randomIndex] = intArray2[randomIndex2]
            intArray2[randomIndex2] = value2
        }
}


func MyRandomFactory_Main(args []string) {
        var randomFactory *MyRandomFactory = MyRandomFactory_GetInstance(); _ = randomFactory
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var size2 int = 100; _ = size2
        for index := int(0); index < size2; index++ {
            stringBuilder.Appendint(randomFactory.GetAbsoluteNextInt(3)).Append(commonSeps.COMMA)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
        stringBuilder.Delete(0, stringBuilder.Length())
        var intArray []int = make([]int, 52); _ = intArray
        var intArray2 []int = make([]int, 52); _ = intArray2
        var size int = len(intArray); _ = size
        for index := int(0); index < size; index++ {
            intArray[index] = index
            intArray2[index] = index
        }
        randomFactory.Shuffle2(intArray, intArray2)
        for index := int(0); index < size; index++ {
            stringBuilder.Appendint(intArray[index]).Append(commonSeps.COMMA)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
        stringBuilder.Delete(0, stringBuilder.Length())
        for index := int(0); index < size; index++ {
            stringBuilder.Appendint(intArray2[index]).Append(commonSeps.COMMA)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
        var found bool; _ = found
        for index := int(0); index < size; index++ {
            found = false
            for index2 := int(0); index2 < size; index2++ {
                if intArray[index2] == index {
                    found = true
                }
            }
            if !found {
                panic(errors.New(fmt.Sprint()))
            }
        }
}

