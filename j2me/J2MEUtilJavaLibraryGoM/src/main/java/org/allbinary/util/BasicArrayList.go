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

//This is a very fast and simple resizable list.
type BasicArrayList struct {
    ArrayUtil *ArrayUtil
    ObjectArray []any
    CurrentIndex int
}

// The Java class BasicArrayList has subclasses so its variables are this interface
type IBasicArrayList interface {
    BasicArrayList_() *BasicArrayList
    AddAt(index int, element any)
    Add(anyType any) bool
    RemoveAt(index int) any
    Remove(anyType any) bool
    RemoveAll2(list IBasicArrayList) bool
    AddAll2(list IBasicArrayList) bool
    AddAllList(list IBasicArrayList) bool
    AddAll(newObjectArray []any) bool
    EnsureCapacity(minSize int)
    TrimToSize()
    IndexOf(anyType any) int
    LastIndexOf(anyType any) int
    Get(index int) any
    Set(index int, element any) any
    Clear()
    Size() int
    IsEmpty() bool
    Contains(anyType any) bool
    ToArray() []any
    ToArrayType(objectArray []any) []any
    Clone() any
    ToString() string
}

func (this *BasicArrayList) BasicArrayList_() *BasicArrayList {
    return this
}

var BasicArrayList_SIZE string = ", Size: "

func (this *BasicArrayList) initFields() {
    this.ArrayUtil = ArrayUtil_GetInstance()
    
    this.CurrentIndex = 0
}


func NewBasicArrayList(objectArray []any) *BasicArrayList {
    this := &BasicArrayList{}
    this.initFields()
        this.ObjectArray = objectArray
    return this
}


func (this *BasicArrayList) AddAt(index int, element any) {
        if index > this.CurrentIndex || index < 0 {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(CommonLabels_GetInstance().INDEX_LABEL)
            stringBuffer.Appendint(index)
            stringBuffer.Append(BasicArrayList_SIZE)
            stringBuffer.Appendint(this.CurrentIndex)
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
        this.EnsureCapacity(this.CurrentIndex + 1)
        copy(this.ObjectArray[index + 1:index + 1 + this.CurrentIndex - index], this.ObjectArray[index:index + this.CurrentIndex - index])
        this.ObjectArray[index] = element
        this.CurrentIndex++
}


func (this *BasicArrayList) Add(anyType any) bool {
        this.EnsureCapacity(this.CurrentIndex + 1)
        this.ObjectArray[Java_PostIncrement(&this.CurrentIndex)] = anyType
        return true
}


func (this *BasicArrayList) RemoveAt(index int) any {
        if index >= this.CurrentIndex {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(CommonLabels_GetInstance().INDEX_LABEL)
            stringBuffer.Appendint(index)
            stringBuffer.Append(BasicArrayList_SIZE)
            stringBuffer.Appendint(this.CurrentIndex)
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
        var oldValue any = this.ObjectArray[index]; _ = oldValue
        var numMoved int = this.CurrentIndex - index - 1; _ = numMoved
        if numMoved > 0 {
            copy(this.ObjectArray[index:index + numMoved], this.ObjectArray[index + 1:index + 1 + numMoved])
        }
        this.ObjectArray[Java_PreDecrement(&this.CurrentIndex)] = nil
        return oldValue
}


func (this *BasicArrayList) Remove(anyType any) bool {
        if anyType == nil {
            for index := int(0); index < this.CurrentIndex; index++ {
                if this.ObjectArray[index] == nil {
                    var numMoved int = this.CurrentIndex - index - 1; _ = numMoved
                    if numMoved > 0 {
                        copy(this.ObjectArray[index:index + numMoved], this.ObjectArray[index + 1:index + 1 + numMoved])
                    }
                    this.ObjectArray[Java_PreDecrement(&this.CurrentIndex)] = nil
                    return true
                }
            }
        } else {
            for index := int(0); index < this.CurrentIndex; index++ {
                if anyType == this.ObjectArray[index] || Java_Equals(anyType, this.ObjectArray[index]) {
                    var numMoved int = this.CurrentIndex - index - 1; _ = numMoved
                    if numMoved > 0 {
                        copy(this.ObjectArray[index:index + numMoved], this.ObjectArray[index + 1:index + 1 + numMoved])
                    }
                    this.ObjectArray[Java_PreDecrement(&this.CurrentIndex)] = nil
                    return true
                }
            }
        }
        return false
}


func (this *BasicArrayList) RemoveAll2(list IBasicArrayList) bool {
        var result bool = true; _ = result
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            if this.Remove(list.Get(index)) {
            } else {
                result = false
            }
        }
        return result
}


func (this *BasicArrayList) AddAll2(list IBasicArrayList) bool {
        this.EnsureCapacity(this.CurrentIndex + list.BasicArrayList_().CurrentIndex)
        var listSize int = list.BasicArrayList_().CurrentIndex; _ = listSize
        for index := int(0); index < listSize; index++ {
            this.ObjectArray[Java_PostIncrement(&this.CurrentIndex)] = list.BasicArrayList_().ObjectArray[index]
        }
        return true
}


func (this *BasicArrayList) AddAllList(list IBasicArrayList) bool {
        var newObjectArray []any = list.ToArray(); _ = newObjectArray
        return this.AddAll(newObjectArray)
}


func (this *BasicArrayList) AddAll(newObjectArray []any) bool {
        var numSize int = len(newObjectArray); _ = numSize
        this.EnsureCapacity(this.CurrentIndex + numSize)
        copy(this.ObjectArray[this.CurrentIndex:this.CurrentIndex + numSize], newObjectArray[0:0 + numSize])
        this.CurrentIndex += numSize
        return numSize != 0
}


func (this *BasicArrayList) EnsureCapacity(minSize int) {
        var oldCapacity int = len(this.ObjectArray); _ = oldCapacity
        if minSize > oldCapacity {
            var newCapacity int = ((oldCapacity * 3) >> 1) + 1; _ = newCapacity
            if newCapacity < minSize {
                newCapacity = minSize
            }
            this.ObjectArray = this.ArrayUtil.CopyOf(this.ObjectArray, newCapacity)
        }
}


func (this *BasicArrayList) TrimToSize() {
        var oldCapacity int = len(this.ObjectArray); _ = oldCapacity
        if this.CurrentIndex < oldCapacity {
            this.ObjectArray = this.ArrayUtil.CopyOf(this.ObjectArray, this.CurrentIndex)
        }
}


func (this *BasicArrayList) IndexOf(anyType any) int {
        if anyType == nil {
            for i := int(0); i < this.CurrentIndex; i++ {
                if this.ObjectArray[i] == nil {
                    return i
                }
            }
        } else {
            for i := int(0); i < this.CurrentIndex; i++ {
                if Java_Equals(anyType, this.ObjectArray[i]) {
                    return i
                }
            }
        }
        return -1
}


func (this *BasicArrayList) LastIndexOf(anyType any) int {
        if anyType == nil {
            for i := int(this.CurrentIndex - 1); i >= 0; i-- {
                if this.ObjectArray[i] == nil {
                    return i
                }
            }
        } else {
            for i := int(this.CurrentIndex - 1); i >= 0; i-- {
                if Java_Equals(anyType, this.ObjectArray[i]) {
                    return i
                }
            }
        }
        return -1
}


func (this *BasicArrayList) Get(index int) any {
        if index >= this.CurrentIndex {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(CommonLabels_GetInstance().INDEX_LABEL)
            stringBuffer.Appendint(index)
            stringBuffer.Append(BasicArrayList_SIZE)
            stringBuffer.Appendint(this.CurrentIndex)
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
        return this.ObjectArray[index]
}


func (this *BasicArrayList) Set(index int, element any) any {
        if index >= this.CurrentIndex {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(CommonLabels_GetInstance().INDEX_LABEL)
            stringBuffer.Appendint(index)
            stringBuffer.Append(BasicArrayList_SIZE)
            stringBuffer.Appendint(this.CurrentIndex)
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
        var oldValue any = this.ObjectArray[index]; _ = oldValue
        this.ObjectArray[index] = element
        return oldValue
}


func (this *BasicArrayList) Clear() {
        for i := int(0); i < this.CurrentIndex; i++ {
            this.ObjectArray[i] = nil
        }
        this.CurrentIndex = 0
}


func (this *BasicArrayList) Size() int {
        return this.CurrentIndex
}


func (this *BasicArrayList) IsEmpty() bool {
        return this.CurrentIndex == 0
}


func (this *BasicArrayList) Contains(anyType any) bool {
        return this.IndexOf(anyType) >= 0
}


func (this *BasicArrayList) ToArray() []any {
        return this.ArrayUtil.CopyOf(this.ObjectArray, this.CurrentIndex)
}


func (this *BasicArrayList) ToArrayType(objectArray []any) []any {
        if len(objectArray) < this.CurrentIndex {
            return this.ArrayUtil.CopyOfType(this.ObjectArray, this.CurrentIndex, reflect.TypeOf(objectArray))
        }
        copy(objectArray[0:0 + this.CurrentIndex], this.ObjectArray[0:0 + this.CurrentIndex])
        if len(objectArray) > this.CurrentIndex {
            objectArray[this.CurrentIndex] = nil
        }
        return objectArray
}


func (this *BasicArrayList) Clone() any {
        var size int = this.Size(); _ = size
        var list IBasicArrayList = NewBasicArrayListS(size); _ = list
        var anyType any; _ = anyType
        for index := int(0); index < size; index++ {
            anyType = this.ObjectArray[index]
            list.Add(anyType)
        }
        return list
}


func (this *BasicArrayList) ToString() string {
        var COMMA_SEP string = CommonSeps_GetInstance().COMMA_SEP; _ = COMMA_SEP
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(0); index < this.CurrentIndex; index++ {
            stringBuffer.Append(fmt.Sprint(this.ObjectArray[index]))
            stringBuffer.Append(COMMA_SEP)
        }
        return fmt.Sprint(stringBuffer)
}

func (this *BasicArrayList) String() string {
    return this.ToString()
}
