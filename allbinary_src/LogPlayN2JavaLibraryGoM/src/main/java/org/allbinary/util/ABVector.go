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


func (this *ABVector[E]) ElementAt(index int) E {
        return this.Vector.ElementAt(index)
}


func (this *ABVector[E]) FirstElement() E {
        return this.Vector.FirstElement()
}


func (this *ABVector[E]) LastElement() E {
        return this.Vector.LastElement()
}


func (this *ABVector[E]) SetElementAt(obj E, index int) {
        this.Vector.SetElementAt(obj, index)
}


func (this *ABVector[E]) RemoveElementAt(index int) {
        this.Vector.RemoveElementAt(index)
}


func (this *ABVector[E]) AddElement(obj E) {
        this.Vector.AddElement(obj)
}


func (this *ABVector[E]) InsertElementAt(obj E, index int) {
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


func (this *ABVector[E]) Set(index int, element E) E {
        return this.Vector.Set(index, element)
}


func (this *ABVector[E]) Add(element E) bool {
        return this.Vector.Add(element)
}


func (this *ABVector[E]) Remove(obj any) bool {
        return this.Vector.Remove(obj)
}


// Java overloaded method: Go needs a different name
func (this *ABVector[E]) Add(index int, element E) {
        this.Vector.Add(index, element)
}


// Java overloaded method: Go needs a different name
func (this *ABVector[E]) Remove(index int) E {
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

