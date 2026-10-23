/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type ABVector[E any] struct {
    Vector
}


func NewABVector[E any]() *ABVector[E] {
    this := &ABVector[E]{}
    this.Vector = *NewVector()
    return this
}


func (this *ABVector[E]) EnsureCapacity(minCapacity int) {
        this.Vector.EnsureCapacity(minCapacity)
}


func (this *ABVector[E]) Capacity() int {
        return this.Vector.Capacity()
}


func (this *ABVector[E]) Size() int {
        return this.Vector.Size()
}


func (this *ABVector[E]) IsEmpty() bool {
        return this.Vector.IsEmpty()
}


func (this *ABVector[E]) Contains(obj any) bool {
        return this.Vector.Contains(obj)
}


func (this *ABVector[E]) IndexOf(obj any) int {
        return this.Vector.IndexOf(obj)
}


func (this *ABVector[E]) LastIndexOf(obj any) int {
        return this.Vector.LastIndexOf(obj)
}


func (this *ABVector[E]) ElementAt(index int) IE {
        return this.Vector.ElementAt(index)
}


func (this *ABVector[E]) FirstElement() IE {
        return this.Vector.FirstElement()
}


func (this *ABVector[E]) LastElement() IE {
        return this.Vector.LastElement()
}


func (this *ABVector[E]) SetElementAt(obj IE, index int) {
        this.Vector.SetElementAt(obj, index)
}


func (this *ABVector[E]) RemoveElementAt(index int) {
        this.Vector.RemoveElementAt(index)
}


func (this *ABVector[E]) AddElement(obj IE) {
        this.Vector.AddElement(obj)
}


func (this *ABVector[E]) InsertElementAt(obj IE, index int) {
        this.Vector.InsertElementAt(obj, index)
}


func (this *ABVector[E]) RemoveElement(obj any) bool {
        return this.Vector.RemoveElement(obj)
}


func (this *ABVector[E]) RemoveAllElements() {
        this.Vector.RemoveAllElements()
}


func (this *ABVector[E]) Clone() any {
        return this.Vector.Clone()
}


func (this *ABVector[E]) ToArray() []any {
        return this.Vector.ToArray()
}


func (this *ABVector[E]) Set(index int, element IE) IE {
        return this.Vector.Set(index, element)
}


func (this *ABVector[E]) Add_1(element IE) bool {
        return this.Vector.Add(element)
}

// Java overloads of ABVector.add
func (this *ABVector[E]) Add(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IE](args[0], pass) {
            return this.Add_1(Java_Arg[IE](args[0]))
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[IE](args[1], pass) {
            return this.Add_2(Java_Arg[int](args[0]), Java_Arg[IE](args[1]))
        }
    }
    panic("No Java overload of ABVector.add for the arguments")
}


func (this *ABVector[E]) Remove_1(obj any) bool {
        return this.Vector.Remove(obj)
}

// Java overloads of ABVector.remove
func (this *ABVector[E]) Remove(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[any](args[0], pass) {
            return this.Remove_1(Java_Arg[any](args[0]))
        }
        if len(args) == 1 && Java_IsArg[int](args[0], pass) {
            return this.Remove_2(Java_Arg[int](args[0]))
        }
    }
    panic("No Java overload of ABVector.remove for the arguments")
}


func (this *ABVector[E]) Add_2(index int, element IE) {
        this.Vector.Add(index, element)
}


func (this *ABVector[E]) Remove_2(index int) IE {
        return this.Vector.Remove(index)
}


func (this *ABVector[E]) Clear() {
        this.Vector.Clear()
}


func (this *ABVector[E]) Equals(obj any) bool {
        return this.Vector.Equals(obj)
}


func (this *ABVector[E]) HashCode() int {
        return this.Vector.HashCode()
}


func (this *ABVector[E]) ToString() string {
        return this.Vector.ToString()
}

func (this *ABVector[E]) String() string {
    return this.ToString()
}


func (this *ABVector[E]) RemoveRange(fromIndex int, toIndex int) {
        this.Vector.RemoveRange(fromIndex, toIndex)
}

