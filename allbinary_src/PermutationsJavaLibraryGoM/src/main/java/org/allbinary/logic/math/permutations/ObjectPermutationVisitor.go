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


func (this *ObjectPermutationVisitor) Print_1(data []any) {
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        this.Print(data, stringBuilder)
}

// Java overloads of ObjectPermutationVisitor.print
func (this *ObjectPermutationVisitor) Print(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[[]any](args[0], pass) {
            this.Print_1(Java_Arg[[]any](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[[]any](args[0], pass) && Java_IsArg[*StringMaker](args[1], pass) {
            this.Print_2(Java_Arg[[]any](args[0]), Java_Arg[*StringMaker](args[1]))
            return
        }
    }
    panic("No Java overload of ObjectPermutationVisitor.print for the arguments")
}


func (this *ObjectPermutationVisitor) Print_2(data []any, stringBuilder *StringMaker) {
        var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
        var size int = len(data); _ = size
        for i := int(0); i < size; i++ {
            stringBuilder.Appendint(tsUtil.HashCode(data[i]))
            stringBuilder.Append(CommonSeps_GetInstance().SPACE)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
}

