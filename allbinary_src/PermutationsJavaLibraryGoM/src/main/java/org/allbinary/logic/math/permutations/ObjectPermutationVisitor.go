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

type ObjectPermutationVisitor struct {
    Visitor
    Data []any
    StringBuilder *StringMaker
}


func (this *ObjectPermutationVisitor) initFields() {
    
    this.StringBuilder = NewStringMaker()
}


func NewObjectPermutationVisitor(data []any) *ObjectPermutationVisitor {
    this := &ObjectPermutationVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
        this.Data = data
    return this
}


func (this *ObjectPermutationVisitor) Visit(anyType any) any {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        this.Print(this.Data, this.StringBuilder)
        return NullUtil_GetInstance().NULL_OBJECT
}


func (this *ObjectPermutationVisitor) Print(data []any) {
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        this.Print(data, stringBuilder)
}


// Java overloaded method: Go needs a different name
func (this *ObjectPermutationVisitor) Print(data []any, stringBuilder *StringMaker) {
        var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
        var size int = len(data); _ = size
        for i := int(0); i < size; i++ {
            stringBuilder.Appendint(int(tsUtil.HashCode(data[i])))
            stringBuilder.Append(CommonSeps_GetInstance().SPACE)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
}

