/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
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

type ColorCacheable struct {
    Key any
    Color *Color
}

// Implements CacheableInterface


func (this *ColorCacheable) initFields() {
    
    
}


func NewColorCacheable(key any) *ColorCacheable {
    this := &ColorCacheable{}
    this.initFields()
        this.Key = key
        this.Color = NewColor(int((any(key).(int))))
    return this
}


func (this *ColorCacheable) GetKey() any {
        return this.Key
}


func (this *ColorCacheable) GetColor() *Color {
        return this.Color
}


func (this *ColorCacheable) SetColor(color *Color) {
        this.Color = color
}

