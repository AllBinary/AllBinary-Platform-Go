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

type OrderedIntArrayPermutationUtil struct {
    TOTAL_ITERATIONS_TABLE string
    FINISHED string
}

var OrderedIntArrayPermutationUtil_Instance *OrderedIntArrayPermutationUtil = NewOrderedIntArrayPermutationUtil()

func (this *OrderedIntArrayPermutationUtil) initFields() {
    this.TOTAL_ITERATIONS_TABLE = "Total Iterations: "
    this.FINISHED = "\nFINISHED"
}


func OrderedIntArrayPermutationUtil_GetInstance() *OrderedIntArrayPermutationUtil {
        return OrderedIntArrayPermutationUtil_Instance
}


func NewOrderedIntArrayPermutationUtil() *OrderedIntArrayPermutationUtil {
    this := &OrderedIntArrayPermutationUtil{}
    this.initFields()
    return this
}


func (this *OrderedIntArrayPermutationUtil) GetNext(data []int) {
        var length int = len(data); _ = length
        var i int = length - 1; _ = i
        for data[i - 1] >= data[i] {
            i = i - 1
        }
        var j int = length; _ = j
        for data[j - 1] <= data[i - 1] {
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


func (this *OrderedIntArrayPermutationUtil) Swap(data []int, a int, b int) {
        var temp int = data[a]; _ = temp
        data[a] = data[b]
        data[b] = temp
}


func (this *OrderedIntArrayPermutationUtil) Factorial(dataLength int) int64 {
        var temp int64 = 1; _ = temp
        if dataLength > 1 {
            for i := int(1); i <= dataLength; i++ {
                temp *= int64(i)
            }
        }
        return temp
}


func (this *OrderedIntArrayPermutationUtil) Generate(data []int, visitor IVisitor) {
        var iterations int64 = this.Factorial(len(data)); _ = iterations
        fmt.Print(this.TOTAL_ITERATIONS_TABLE)
        fmt.Println(iterations)
        for count := int64(0); count < iterations - 1; count++ {
            this.GetNext(data)
            visitor.Visit(this)
        }
        fmt.Println(this.FINISHED)
}


func OrderedIntArrayPermutationUtil_Main(args []string) {
        var TOTAL int = 3; _ = TOTAL
        var data []int = make([]int, TOTAL); _ = data
        for i := int(0); i < TOTAL; i++ {
            data[i] = i
        }
        var permutationVisitor IIntPermutationVisitor = NewIntPermutationVisitor(data); _ = permutationVisitor
        var ORIGINAL_DATA_LABEL string = "Original Data: "; _ = ORIGINAL_DATA_LABEL
        fmt.Print(ORIGINAL_DATA_LABEL)
        permutationVisitor.Print(data)
        OrderedIntArrayPermutationUtil_GetInstance().Generate(data, permutationVisitor)
}

