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

type ABHashtable[K any, V any] struct {
    Hashtable
}





// Java overloads of new ABHashtable
func NewABHashtable[K any, V any](args ...any) *ABHashtable[K, V] {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[int](args[0], pass) {
            return NewABHashtable_1[K, V](Java_Arg[int](args[0]))
        }
        if len(args) == 0 {
            return NewABHashtable_2[K, V]()
        }
    }
    panic("No Java overload of new ABHashtable for the arguments")
}
func NewABHashtable_1[K any, V any](initialCapacity int) *ABHashtable[K, V] {
    this := &ABHashtable[K, V]{}
        this.Hashtable = *NewHashtable(initialCapacity)
    return this
}


func NewABHashtable_2[K any, V any]() *ABHashtable[K, V] {
    this := &ABHashtable[K, V]{}
        this.Hashtable = *NewHashtable()
    return this
}

