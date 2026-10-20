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

type BooleanFactory struct {
    TRUE bool
    FALSE bool
    TRUE_STRING string
    FALSE_STRING string
    YES string
    NO string
}

var BooleanFactory_Instance *BooleanFactory = NewBooleanFactory()

func (this *BooleanFactory) initFields() {
    this.TRUE = true
    this.FALSE = false
    this.TRUE_STRING = fmt.Sprint(this.TRUE)
    this.FALSE_STRING = fmt.Sprint(this.FALSE)
    this.YES = "yes"
    this.NO = "no"
}

func NewBooleanFactory() *BooleanFactory {
    this := &BooleanFactory{}
    this.initFields()
    return this
}


func BooleanFactory_GetInstance() *BooleanFactory {
        return BooleanFactory_Instance
}


func (this *BooleanFactory) ToStringb(bool_ bool) string {
        if bool_ {
            return this.TRUE_STRING
        } else {
            return this.FALSE_STRING
        }
}
