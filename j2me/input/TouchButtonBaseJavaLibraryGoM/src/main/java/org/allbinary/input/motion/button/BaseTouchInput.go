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

type BaseTouchInput struct {
    CommonStrings *CommonStrings
    BasicArrayListUtil *BasicArrayListUtil
}

// The Java class BaseTouchInput has subclasses so its variables are this interface
type IBaseTouchInput interface {
    BaseTouchInput_() *BaseTouchInput
    GetList() IBasicArrayList
    ToString() string
}

func (this *BaseTouchInput) BaseTouchInput_() *BaseTouchInput {
    return this
}


func (this *BaseTouchInput) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewBaseTouchInput() *BaseTouchInput {
    this := &BaseTouchInput{}
    this.initFields()
    return this
}


func (this *BaseTouchInput) GetList() IBasicArrayList {
        return this.BasicArrayListUtil.GetImmutableInstance()
}


func (this *BaseTouchInput) ToString() string {
        return reflect.TypeOf(this).String()
}

func (this *BaseTouchInput) String() string {
    return this.ToString()
}

