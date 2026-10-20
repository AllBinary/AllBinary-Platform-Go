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

type PrimitiveLongSingleton struct {
    UNK []rune
    NUMBER_ARRAY []int8
    NUMBER_CHAR_ARRAY []rune
    NUMBER_CHAR_ARRAYS [][]rune
    ZERO []rune
}

var PrimitiveLongSingleton_Instance *PrimitiveLongSingleton = NewPrimitiveLongSingleton()

func (this *PrimitiveLongSingleton) initFields() {
    this.UNK = []rune{'U', 'n', 'k'}
    this.NUMBER_ARRAY = []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    this.NUMBER_CHAR_ARRAY = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
    this.NUMBER_CHAR_ARRAYS = [][]rune{[]rune{'0'}, []rune{'1'}, []rune{'2'}, []rune{'3'}, []rune{'4'}, []rune{'5'}, []rune{'6'}, []rune{'7'}, []rune{'8'}, []rune{'9'}}
    this.ZERO = []rune{this.NUMBER_CHAR_ARRAY[0]}
}

func NewPrimitiveLongSingleton() *PrimitiveLongSingleton {
    this := &PrimitiveLongSingleton{}
    this.initFields()
    return this
}


func PrimitiveLongSingleton_GetInstance() *PrimitiveLongSingleton {
        return PrimitiveLongSingleton_Instance
}


// throws Exception
func (this *PrimitiveLongSingleton) GetNumberFromChar(numberAsChar rune) int8 {
        var size int = len(this.NUMBER_CHAR_ARRAY); _ = size
        for index := int(0); index < size; index++ {
            if this.NUMBER_CHAR_ARRAY[index] == numberAsChar {
                return this.NUMBER_ARRAY[index]
            }
        }
        panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("numberAsChar was not in array: ").Appendchar(numberAsChar)))))
}
