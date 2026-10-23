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

type CharacterSetData struct {
    US_ASCII string
    ISO_8859_1 string
    UTF_8 string
    UTF_16 string
    UTF_16BE string
    UTF_16LE string
}

var CharacterSetData_Instance *CharacterSetData = NewCharacterSetData()

func (this *CharacterSetData) initFields() {
    this.US_ASCII = "US-ASCII"
    this.ISO_8859_1 = "ISO-8859-1"
    this.UTF_8 = "UTF-8"
    this.UTF_16 = "UTF-16"
    this.UTF_16BE = "UTF-16BE"
    this.UTF_16LE = "UTF-16LE"
}


func CharacterSetData_GetInstance() *CharacterSetData {
        return CharacterSetData_Instance
}


func NewCharacterSetData() *CharacterSetData {
    this := &CharacterSetData{}
    this.initFields()
    return this
}

