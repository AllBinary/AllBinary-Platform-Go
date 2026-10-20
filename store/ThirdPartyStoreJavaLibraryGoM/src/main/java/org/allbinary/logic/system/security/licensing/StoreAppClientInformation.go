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

type StoreAppClientInformation struct {
    AbeClientInformation
    LongArrayIdentifier []int64
}

// Implements LongArrayIdentifierInterface


func (this *StoreAppClientInformation) initFields() {
    
}


func NewStoreAppClientInformation(name string, version string, specialName string, shortName string, longArrayIdentifier []int64) *StoreAppClientInformation {
    this := &StoreAppClientInformation{}
        this.AbeClientInformation = *NewAbeClientInformation(name, version, specialName, shortName)
        this.initFields()
        this.LongArrayIdentifier = longArrayIdentifier
    return this
}


func (this *StoreAppClientInformation) GetLongArrayIdentifier() []int64 {
        return this.LongArrayIdentifier
}

