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

type ImmutableBasicArrayList struct {
    BasicArrayListS
    Name string
}


func (this *ImmutableBasicArrayList) initFields() {
    
}


func NewImmutableBasicArrayList(name string, size int) *ImmutableBasicArrayList {
    this := &ImmutableBasicArrayList{}
        this.BasicArrayListS = *NewBasicArrayListS(size)
        this.initFields()
        this.Name = name
    return this
}


func (this *ImmutableBasicArrayList) Add(e any) bool {
        ForcedLogUtil_Log(this.Name, this)
        return true
}


func (this *ImmutableBasicArrayList) AddAt(index int, element any) {
        ForcedLogUtil_Log(this.Name, this)
}


func (this *ImmutableBasicArrayList) ToString() string {
        return this.Name
}

func (this *ImmutableBasicArrayList) String() string {
    return this.ToString()
}
