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

type BooleanFileVisitor struct {
    CommonStrings *CommonStrings
    FilterStringBasicArrayList *BasicArrayList
}

// Implements VisitorInterface


func (this *BooleanFileVisitor) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewBooleanFileVisitor(filterStringBasicArrayList *BasicArrayList) *BooleanFileVisitor {
    this := &BooleanFileVisitor{}
    this.initFields()
        this.SetFilterStringBasicArrayList(filterStringBasicArrayList)
        PreLogUtil_Put("Filter BasicArrayList: " + fmt.Sprint(this.GetFilterStringBasicArrayList()), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *BooleanFileVisitor) Visit(anyType any) any {
        return any(this.Visit(any(anyType).(*AbFile))).(any)
}


// Java overloaded method: Go needs a different name
func (this *BooleanFileVisitor) Visit(file *AbFile) bool {
        var list *BasicArrayList = this.GetFilterStringBasicArrayList(); _ = list
        var size int = list.Size(); _ = size
        var nextFileFilterString string; _ = nextFileFilterString
        for index := int(0); index < size; index++ {
            nextFileFilterString = any(list.Get(index)).(string)
            if this.Visit(file, nextFileFilterString) {
                return true
            }
        }
        return false
}


func (this *BooleanFileVisitor) GetFilterStringBasicArrayList() *BasicArrayList {
        return this.FilterStringBasicArrayList
}


func (this *BooleanFileVisitor) SetFilterStringBasicArrayList(filterStringBasicArrayList *BasicArrayList) {
        this.FilterStringBasicArrayList = filterStringBasicArrayList
}


// Java overloaded method: Go needs a different name
func (this *BooleanFileVisitor) Visit(file *AbFile, fileNameString string) bool {
        return false
}

