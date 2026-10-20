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

type BasicArrayListUtil struct {
    NULL_ARRAY_OF_ARRAY [][]IBasicArrayList
    MyRandomFactory *MyRandomFactory
    ImmutableInstance IBasicArrayList
}

var BasicArrayListUtil_Instance *BasicArrayListUtil = NewBasicArrayListUtil()

func (this *BasicArrayListUtil) initFields() {
    this.NULL_ARRAY_OF_ARRAY = func() [][]IBasicArrayList { array := make([][]IBasicArrayList, 0); for index := range array { array[index] = make([]IBasicArrayList, 0) }; return array }()
    this.MyRandomFactory = MyRandomFactory_GetInstance()
    this.ImmutableInstance = NewImmutableBasicArrayList("Immutable", 0)
}

func NewBasicArrayListUtil() *BasicArrayListUtil {
    this := &BasicArrayListUtil{}
    this.initFields()
    return this
}


func BasicArrayListUtil_GetInstance() *BasicArrayListUtil {
        return BasicArrayListUtil_Instance
}


func (this *BasicArrayListUtil) GetRandom(list IBasicArrayList) any {
        var i_random int = this.GetRandomIndex(list); _ = i_random
        return list.BasicArrayList_().ObjectArray[i_random]
}


func (this *BasicArrayListUtil) GetRandomIndex(list IBasicArrayList) int {
        return this.MyRandomFactory.GetAbsoluteNextInt(list.Size())
}


func (this *BasicArrayListUtil) Reverse(list IBasicArrayList) {
        var lastMinusOneIndex int = list.Size() - 2; _ = lastMinusOneIndex
        var temp any; _ = temp
        var size int = lastMinusOneIndex / 2; _ = size
        var secondSwapIndex int; _ = secondSwapIndex
        for index := int(0); index <= size; index++ {
            secondSwapIndex = lastMinusOneIndex - index
            temp = list.RemoveAt(index)
            list.AddAt(index, list.RemoveAt(secondSwapIndex))
            list.AddAt(secondSwapIndex + 1, temp)
        }
}


func (this *BasicArrayListUtil) Log(list IBasicArrayList) string {
        var S_LABEL string = " s: "; _ = S_LABEL
        var size int = list.Size(); _ = size
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(CommonLabels_GetInstance().TOTAL_LABEL)
        stringBuffer.Appendint(size)
        for index := int(size - 1); index >= 0; index-- {
            var usedList IBasicArrayList = any(list.BasicArrayList_().ObjectArray[index]).(IBasicArrayList); _ = usedList
            stringBuffer.Append(S_LABEL)
            stringBuffer.Appendint(usedList.Size())
        }
        return fmt.Sprint(stringBuffer)
}


func (this *BasicArrayListUtil) GetImmutableInstance() IBasicArrayList {
        return this.ImmutableInstance
}


func BasicArrayListUtil_Main(args []string) {
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        list.Add(commonPhoneStrings.ONE)
        list.Add(commonPhoneStrings.TWO)
        list.Add(commonPhoneStrings.THREE)
        list.Add(commonPhoneStrings.FOUR)
        list.Add(commonPhoneStrings.FIVE)
        list.Add(commonPhoneStrings.SIX)
        list.Add(commonPhoneStrings.SEVEN)
        list.Add(commonPhoneStrings.EIGHT)
        list.Add(commonPhoneStrings.NINE)
        BasicArrayListUtil_GetInstance().Reverse(list)
        fmt.Println(fmt.Sprint(list))
}
