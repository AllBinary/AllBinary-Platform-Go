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

type NullTickable struct {
    NAME string
}

// Implements TickableInterface

var NullTickable_SINGLETON *NullTickable = NewNullTickable()

func (this *NullTickable) initFields() {
    this.NAME = "NulLTickable"
}

func NewNullTickable() *NullTickable {
    this := &NullTickable{}
    this.initFields()
    return this
}


func NullTickable_GetInstance() *NullTickable {
        return NullTickable_SINGLETON
}


// throws Exception
func (this *NullTickable) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
}


func (this *NullTickable) GetName() string {
        return this.NAME
}

