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

type AndroidInfoFactory struct {
    Version int
}

var AndroidInfoFactory_Instance *AndroidInfoFactory = NewAndroidInfoFactory()

func (this *AndroidInfoFactory) initFields() {
    this.Version = func() int { value, _ := strconv.Atoi(Build_VERSION.SDK); return value }()
}

func NewAndroidInfoFactory() *AndroidInfoFactory {
    this := &AndroidInfoFactory{}
    this.initFields()
    return this
}


func AndroidInfoFactory_GetInstance() *AndroidInfoFactory {
        return AndroidInfoFactory_Instance
}


func (this *AndroidInfoFactory) GetVersion() int {
        return this.Version
}

