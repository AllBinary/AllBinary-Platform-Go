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

type IntPermutationVisitor struct {
    Visitor
    Data []int
    StringBuilder *StringMaker
}


func (this *IntPermutationVisitor) initFields() {
    
    this.StringBuilder = NewStringMaker()
}


func NewIntPermutationVisitor(data []int) *IntPermutationVisitor {
    this := &IntPermutationVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
        this.Data = data
    return this
}


func (this *IntPermutationVisitor) Visit(anyType any) any {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        this.Print(this.Data, this.StringBuilder)
        return NullUtil_GetInstance().NULL_OBJECT
}


func (this *IntPermutationVisitor) Print(data []int) {
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        this.Print(data, stringBuilder)
}


// Java overloaded method: Go needs a different name
func (this *IntPermutationVisitor) Print(data []int, stringBuilder *StringMaker) {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var size int = len(data); _ = size
        for i := int(0); i < size; i++ {
            stringBuilder.Appendint(data[i])
            stringBuilder.Append(commonSeps.SPACE)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
}

