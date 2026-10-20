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

type ABStack[E any] struct {
    Stack
}


func NewABStack[E any]() *ABStack[E] {
    this := &ABStack[E]{}
    this.Stack = *NewStack()
    return this
}


func (this *ABStack[E]) Push(item E) E {
        return this.Stack.Push(item)
}


func (this *ABStack[E]) Pop() E {
        return this.Stack.Pop()
}


func (this *ABStack[E]) Peek() E {
        return this.Stack.Peek()
}


func (this *ABStack[E]) Empty() bool {
        return this.Stack.Empty()
}


func (this *ABStack[E]) Search(obj any) int {
        return this.Stack.Search(obj)
}


func (this *ABStack[E]) RemoveElement(obj any) bool {
        return this.Stack.RemoveElement(obj)
}

