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

type StringMaker struct {
    CharArray []rune
    CurrentLength int
}


func (this *StringMaker) initFields() {
    
    
}


func NewStringMaker() *StringMaker {
    this := &StringMaker{}
    this.initFields()
        this.CharArray = make([]rune, 20)
    return this
}


func (this *StringMaker) Append(string_ string) *StringMaker {
        var stringLength int = len(string_); _ = stringLength
        this.EnsureCapacity(this.CurrentLength + stringLength)
        copy(this.CharArray[this.CurrentLength:], []rune(string_[0:stringLength]))
        this.CurrentLength += stringLength
        return this
}


func (this *StringMaker) Appendchar(c rune) *StringMaker {
        this.EnsureCapacity(this.CurrentLength + 1)
        this.CharArray[Java_PostIncrement(&this.CurrentLength)] = c
        return this
}


func (this *StringMaker) AppendCharArray(charArray []rune, start int, end int) *StringMaker {
        this.EnsureCapacity(this.CurrentLength + (end - start))
        for index := int(start); index < end; index++ {
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = charArray[index]
        }
        return this
}


func (this *StringMaker) Appendbyte(b int8) *StringMaker {
        this.Append(fmt.Sprint(b))
        return this
}


func (this *StringMaker) Appendshort(b int16) *StringMaker {
        this.Append(fmt.Sprint(b))
        return this
}


func (this *StringMaker) Appendint(i int) *StringMaker {
        this.Append(fmt.Sprint(i))
        return this
}


func (this *StringMaker) Appendlong(l int64) *StringMaker {
        this.Append(fmt.Sprint(l))
        return this
}


func (this *StringMaker) Appendfloat(f float32) *StringMaker {
        this.Append(fmt.Sprint(f))
        return this
}


func (this *StringMaker) Appendboolean(bool_ bool) *StringMaker {
        if bool_ {
            this.EnsureCapacity(this.CurrentLength + 4)
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 't'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'r'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'u'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'e'
        } else {
            this.EnsureCapacity(this.CurrentLength + 5)
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'f'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'a'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'l'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 's'
            this.CharArray[Java_PostIncrement(&this.CurrentLength)] = 'e'
        }
        return this
}


func (this *StringMaker) EnsureCapacity(minSize int) {
        var oldCapacity int = len(this.CharArray); _ = oldCapacity
        if minSize > oldCapacity {
            var newCapacity int = oldCapacity * 3; _ = newCapacity
            if newCapacity < minSize {
                newCapacity = minSize
            }
            var copy_ []rune = make([]rune, newCapacity); _ = copy_
            var min_ int = this.Min(len(this.CharArray), newCapacity); _ = min_
            copy(copy_[0:0 + min_], this.CharArray[0:0 + min_])
            this.CharArray = copy_
        }
}


func (this *StringMaker) Min(value int, value2 int) int {
        return func() int { if (value <= value2) { return value }; return value2 }()
}


func (this *StringMaker) Delete(start int, end int) *StringMaker {
        this.CurrentLength -= (end - start)
        return this
}


func (this *StringMaker) Length() int {
        return this.CurrentLength
}


func (this *StringMaker) ToString() string {
        return string(this.CharArray[0:0 + this.CurrentLength])
}

func (this *StringMaker) String() string {
    return this.ToString()
}

