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

// The Java class IntPermutationVisitor has subclasses so its variables are this interface
type IIntPermutationVisitor interface {
    IVisitor
    IntPermutationVisitor_() *IntPermutationVisitor
    Visit(anyType any) any
    Print(args ...any)
}

func (this *IntPermutationVisitor) IntPermutationVisitor_() *IntPermutationVisitor {
    return this
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


func (this *IntPermutationVisitor) Print_1(data []int) {
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        this.Print(data, stringBuilder)
}

// Java overloads of IntPermutationVisitor.print
func (this *IntPermutationVisitor) Print(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[[]int](args[0], pass) {
            this.Print_1(Java_Arg[[]int](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[[]int](args[0], pass) && Java_IsArg[*StringMaker](args[1], pass) {
            this.Print_2(Java_Arg[[]int](args[0]), Java_Arg[*StringMaker](args[1]))
            return
        }
    }
    panic("No Java overload of IntPermutationVisitor.print for the arguments")
}


func (this *IntPermutationVisitor) Print_2(data []int, stringBuilder *StringMaker) {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var size int = len(data); _ = size
        for i := int(0); i < size; i++ {
            stringBuilder.Appendint(data[i])
            stringBuilder.Append(commonSeps.SPACE)
        }
        fmt.Println(fmt.Sprint(stringBuilder))
}

