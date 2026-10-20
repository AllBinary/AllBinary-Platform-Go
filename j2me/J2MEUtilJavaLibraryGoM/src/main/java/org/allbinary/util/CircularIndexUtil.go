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

type CircularIndexUtil struct {
    Index int
    LastIndex int
    Size int
}

var CircularIndexUtil_NULL_CIRCULAR_INDEX_UTIL *CircularIndexUtil = NewCircularIndexUtil(0, 0)

func (this *CircularIndexUtil) initFields() {
    this.Index = 0
    this.LastIndex = 0
    this.Size = 0
}


func NewCircularIndexUtil(index int, size int) *CircularIndexUtil {
    this := &CircularIndexUtil{}
    this.initFields()
        this.SetSize(size)
        this.SetIndex(index)
    return this
}


func (this *CircularIndexUtil) Next() int {
        if this.Index >= this.LastIndex {
            this.Index = 0
        } else {
            this.Index++
        }
        return this.Index
}


func (this *CircularIndexUtil) Previous() int {
        if this.Index < 1 {
            this.Index = this.LastIndex
        } else {
            this.Index--
        }
        return this.Index
}


func (this *CircularIndexUtil) SetIndex(index int) {
        if index > this.LastIndex {
            return
        }
        this.Index = index
}


func (this *CircularIndexUtil) GetIndex() int {
        return this.Index
}


func CircularIndexUtil_CreateInstance(max_ int) *CircularIndexUtil {
        return NewCircularIndexUtil(0, max_)
}


func CircularIndexUtil_CreateInstanceAt(index int, max_ int) *CircularIndexUtil {
        return NewCircularIndexUtil(index, max_)
}


func (this *CircularIndexUtil) SetSize(size int) {
        this.Size = size
        this.LastIndex = size - 1
        this.Index = 0
}


func (this *CircularIndexUtil) GetSize() int {
        return this.Size
}


func (this *CircularIndexUtil) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().INDEX_LABEL).Appendint(this.Index).Append("lastIndex: ").Appendint(this.LastIndex).Append("size: ").Appendint(this.Size))
}

func (this *CircularIndexUtil) String() string {
    return this.ToString()
}
