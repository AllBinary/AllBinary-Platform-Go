/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type PlatformTextureBaseFactory struct {
    CommonStrings *CommonStrings
}

// The Java class PlatformTextureBaseFactory has subclasses so its variables are this interface
type IPlatformTextureBaseFactory interface {
    PlatformTextureBaseFactory_() *PlatformTextureBaseFactory
    Load(gl GL10, target int, level int, image IImage, border int, flip bool)
}

func (this *PlatformTextureBaseFactory) PlatformTextureBaseFactory_() *PlatformTextureBaseFactory {
    return this
}

var PlatformTextureBaseFactory_NULL_PLATFORM_TEXTURE_BASE_FACTORY IPlatformTextureBaseFactory = NewPlatformTextureBaseFactory()

func (this *PlatformTextureBaseFactory) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewPlatformTextureBaseFactory() *PlatformTextureBaseFactory {
    this := &PlatformTextureBaseFactory{}
    this.initFields()
    return this
}


func (this *PlatformTextureBaseFactory) Load(gl GL10, target int, level int, image IImage, border int, flip bool) {
}

