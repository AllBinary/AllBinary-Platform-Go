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

type ArrayUtil struct {
    MathUtil *MathUtil
}

var ArrayUtil_Instance *ArrayUtil = NewArrayUtil()

func (this *ArrayUtil) initFields() {
    this.MathUtil = MathUtil_GetInstance()
}

func NewArrayUtil() *ArrayUtil {
    this := &ArrayUtil{}
    this.initFields()
    return this
}


func ArrayUtil_GetInstance() *ArrayUtil {
        return ArrayUtil_Instance
}


func (this *ArrayUtil) CopyOf(original []any, newLength int) []any {
        return this.CopyOfType(original, newLength, reflect.TypeOf(original))
}


func (this *ArrayUtil) CopyOfType(original []any, newLength int, newType reflect.Type) []any {
        var copy_ []any = make([]any, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfbyte(original []int8, newLength int) []int8 {
        var copy_ []int8 = make([]int8, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfshort(original []int16, newLength int) []int16 {
        var copy_ []int16 = make([]int16, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfint(original []int, newLength int) []int {
        var copy_ []int = make([]int, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOflong(original []int64, newLength int) []int64 {
        var copy_ []int64 = make([]int64, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfchar(original []rune, newLength int) []rune {
        var copy_ []rune = make([]rune, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOffloat(original []float32, newLength int) []float32 {
        var copy_ []float32 = make([]float32, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfdouble(original []float64, newLength int) []float64 {
        var copy_ []float64 = make([]float64, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) CopyOfboolean(original []bool, newLength int) []bool {
        var copy_ []bool = make([]bool, newLength); _ = copy_
        copy(copy_[0:0 + this.MathUtil.Min(len(original), newLength)], original[0:0 + this.MathUtil.Min(len(original), newLength)])
        return copy_
}


func (this *ArrayUtil) Reverse(intArray []int) {
        var length int = len(intArray) - 1; _ = length
        var temp int; _ = temp
        var size int = length / 2; _ = size
        for index := int(0); index <= size; index++ {
            temp = intArray[index]
            intArray[index] = intArray[length - index]
            intArray[length - index] = temp
        }
}


func (this *ArrayUtil) Flip(intArray []int, width int, height int) {
        var length int = height - 1; _ = length
        var temp int; _ = temp
        var index int; _ = index
        var rowIndex int; _ = rowIndex
        for y := int(0); y <= length / 2; y++ {
            for x := int(0); x < width; x++ {
                rowIndex = ((length - y) * width)
                index = x + (y * width)
                temp = intArray[index]
                intArray[index] = intArray[rowIndex + x]
                intArray[rowIndex + x] = temp
            }
        }
}


func (this *ArrayUtil) ToStringFromTwoDimensionaArray(twoDimensionalIntArray [][]int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var endIndex int = len(twoDimensionalIntArray); _ = endIndex
        var endIndex2 int = len(twoDimensionalIntArray[0]); _ = endIndex2
        for index := int(0); index < endIndex; index++ {
            stringBuffer.Append(commonSeps.BRACE_OPEN)
            for index2 := int(0); index2 < endIndex2; index2++ {
                stringBuffer.Appendint(int(twoDimensionalIntArray[index][index2]))
                stringBuffer.Append(commonSeps.COMMA_SEP)
            }
            stringBuffer.Append(commonSeps.BRACE_CLOSE)
            stringBuffer.Append(commonSeps.NEW_LINE)
        }
        return fmt.Sprint(stringBuffer)
}


func (this *ArrayUtil) ToStringFromObjectArray(objectArray []any) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        for index := int(0); index < len(objectArray); index++ {
            stringBuffer.Append(StringUtil_GetInstance().ToString(objectArray[index]))
            stringBuffer.Append(commonSeps.COMMA_SEP)
        }
        return fmt.Sprint(stringBuffer)
}


func (this *ArrayUtil) ToStringFromByteArray(byteArray []int8) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        for index := int(0); index < len(byteArray); index++ {
            stringBuffer.Appendbyte(byteArray[index])
            stringBuffer.Append(commonSeps.COMMA_SEP)
        }
        return fmt.Sprint(stringBuffer)
}

