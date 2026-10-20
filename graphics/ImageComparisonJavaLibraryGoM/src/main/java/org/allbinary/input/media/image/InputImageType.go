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

type InputImageType struct {
    Name string
    Index int
}

var InputImageType_HashMap *HashMap = StdUtil_GetInstance().CreateHashMap()
var InputImageType_TYPE_VECTOR IBasicArrayList = NewBasicArrayListD()
var InputImageType_CAPTURE *InputImageType = NewInputImageType("Capture", 0)
var InputImageType_COMPARISON *InputImageType = NewInputImageType("Comparison", 1)
var InputImageType_MOTION *InputImageType = NewInputImageType("Motion", 2)

func (this *InputImageType) initFields() {
    
    
}


func NewInputImageType(name string, index int) *InputImageType {
    this := &InputImageType{}
    this.initFields()
        this.Name = name
        this.Index = index
        InputImageType_TYPE_VECTOR.Add(this)
        InputImageType_HashMap.Put(this.GetName(), this)
    return this
}


func (this *InputImageType) ToString() string {
        return "ImageType: " + this.GetName()
}

func (this *InputImageType) String() string {
    return this.ToString()
}


func (this *InputImageType) GetName() string {
        return this.Name
}


func InputImageType_GetInstance(imageTypeString string) *InputImageType {
        return any(InputImageType_HashMap.Get(imageTypeString)).(*InputImageType)
}


func InputImageType_GetAllAsVector() IBasicArrayList {
        return InputImageType_TYPE_VECTOR
}


func (this *InputImageType) GetIndex() int {
        return this.Index
}

