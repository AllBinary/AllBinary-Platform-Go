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

type ItemColorFactory struct {
    INVERT_PAINT int
    PAINT int
    TEXT_FIELD_NO_FOCUS int
}

var ItemColorFactory_Instance *ItemColorFactory = NewItemColorFactory()

func (this *ItemColorFactory) initFields() {
    this.INVERT_PAINT = 0xFF000000
    this.PAINT = 0xFFe07718
    this.TEXT_FIELD_NO_FOCUS = 0x7F7F7F
}

func NewItemColorFactory() *ItemColorFactory {
    this := &ItemColorFactory{}
    this.initFields()
    return this
}


func ItemColorFactory_GetInstance() *ItemColorFactory {
        return ItemColorFactory_Instance
}

