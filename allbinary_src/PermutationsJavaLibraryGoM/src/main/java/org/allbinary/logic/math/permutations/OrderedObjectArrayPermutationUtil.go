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

type OrderedObjectArrayPermutationUtil struct {
    TOTAL_ITERATIONS_TABLE string
    FINISHED string
}

var OrderedObjectArrayPermutationUtil_Instance *OrderedObjectArrayPermutationUtil = NewOrderedObjectArrayPermutationUtil()

func (this *OrderedObjectArrayPermutationUtil) initFields() {
    this.TOTAL_ITERATIONS_TABLE = "Total Iterations: "
    this.FINISHED = "\nFINISHED"
}


func OrderedObjectArrayPermutationUtil_GetInstance() *OrderedObjectArrayPermutationUtil {
        return OrderedObjectArrayPermutationUtil_Instance
}


func NewOrderedObjectArrayPermutationUtil() *OrderedObjectArrayPermutationUtil {
    this := &OrderedObjectArrayPermutationUtil{}
    this.initFields()
    return this
}


func (this *OrderedObjectArrayPermutationUtil) GetNext(data []any) {
        var length int = len(data); _ = length
        var i int = length - 1; _ = i
        var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
        for tsUtil.HashCode(data[i - 1]) >= tsUtil.HashCode(data[i]) {
            i = i - 1
        }
        var j int = length; _ = j
        for tsUtil.HashCode(data[j - 1]) <= tsUtil.HashCode(data[i - 1]) {
            j = j - 1
        }
        this.Swap(data, i - 1, j - 1)
        i++
        j = length
        for i < j {
            this.Swap(data, i - 1, j - 1)
            i++
            j--
        }
}


func (this *OrderedObjectArrayPermutationUtil) Swap(data []any, a int, b int) {
        var temp any = data[a]; _ = temp
        data[a] = data[b]
        data[b] = temp
}


func (this *OrderedObjectArrayPermutationUtil) Factorial(dataLength int) int64 {
        var temp int64 = 1; _ = temp
        if dataLength > 1 {
            for i := int(1); i <= dataLength; i++ {
                temp *= int64(i)
            }
        }
        return temp
}


func (this *OrderedObjectArrayPermutationUtil) Generate(data []any, visitor *Visitor) {
        var iterations int64 = this.Factorial(len(data)); _ = iterations
        fmt.Print(this.TOTAL_ITERATIONS_TABLE)
        fmt.Println(iterations)
        for count := int64(0); count < iterations - 1; count++ {
            this.GetNext(data)
            visitor.Visit(this)
        }
        fmt.Println(this.FINISHED)
}


func OrderedObjectArrayPermutationUtil_Main(args []string) {
        var TOTAL int = 7; _ = TOTAL
        var data []any = make([]any, TOTAL); _ = data
        for i := int(0); i < TOTAL; i++ {
            data[i] = NewComparableObject()
        }
        Arrays_Sort(data)
        var permutationVisitor *ObjectPermutationVisitor = NewObjectPermutationVisitor(data); _ = permutationVisitor
        var ORIGINAL_DATA_LABEL string = "Original Data: "; _ = ORIGINAL_DATA_LABEL
        fmt.Print(ORIGINAL_DATA_LABEL)
        permutationVisitor.Visit(data)
        OrderedObjectArrayPermutationUtil_GetInstance().Generate(data, permutationVisitor)
}

