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
    FilterStringBasicArrayList IBasicArrayList
}

// Implements VisitorInterface

// The Java class BooleanFileVisitor has subclasses so its variables are this interface
type IBooleanFileVisitor interface {
    BooleanFileVisitor_() *BooleanFileVisitor
    Visit(args ...any) any
    GetFilterStringBasicArrayList() IBasicArrayList
    SetFilterStringBasicArrayList(filterStringBasicArrayList IBasicArrayList)
}

func (this *BooleanFileVisitor) BooleanFileVisitor_() *BooleanFileVisitor {
    return this
}


func (this *BooleanFileVisitor) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewBooleanFileVisitor(filterStringBasicArrayList IBasicArrayList) *BooleanFileVisitor {
    this := &BooleanFileVisitor{}
    this.initFields()
        this.SetFilterStringBasicArrayList(filterStringBasicArrayList)
        PreLogUtil_Put("Filter BasicArrayList: " + fmt.Sprint(this.GetFilterStringBasicArrayList()), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *BooleanFileVisitor) Visit_1(anyType any) any {
        return any(this.Visit(any(anyType).(*AbFile)).(any)).(any)
}

// Java overloads of BooleanFileVisitor.visit
func (this *BooleanFileVisitor) Visit(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[any](args[0], pass) {
            return this.Visit_1(Java_Arg[any](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return this.Visit_2(Java_Arg[*AbFile](args[0]))
        }
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.Visit_3(Java_Arg[*AbFile](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of BooleanFileVisitor.visit for the arguments")
}


func (this *BooleanFileVisitor) Visit_2(file *AbFile) bool {
        var list IBasicArrayList = this.GetFilterStringBasicArrayList(); _ = list
        var size int = list.Size(); _ = size
        var nextFileFilterString string; _ = nextFileFilterString
        for index := int(0); index < size; index++ {
            nextFileFilterString = any(list.Get(index)).(string)
            if this.Visit(file, nextFileFilterString).(bool) {
                return true
            }
        }
        return false
}


func (this *BooleanFileVisitor) GetFilterStringBasicArrayList() IBasicArrayList {
        return this.FilterStringBasicArrayList
}


func (this *BooleanFileVisitor) SetFilterStringBasicArrayList(filterStringBasicArrayList IBasicArrayList) {
        this.FilterStringBasicArrayList = filterStringBasicArrayList
}


func (this *BooleanFileVisitor) Visit_3(file *AbFile, fileNameString string) bool {
        return false
}

