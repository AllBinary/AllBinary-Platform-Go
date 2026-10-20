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

type PlatformInputMappingFactory struct {
    LogUtil *LogUtil
    InputToGameKeyMapping IPersistentInputMapping
}

var PlatformInputMappingFactory_Instance *PlatformInputMappingFactory = NewPlatformInputMappingFactory()

func (this *PlatformInputMappingFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.InputToGameKeyMapping = PersistentInputMapping_GetNullInstance()
}

func NewPlatformInputMappingFactory() *PlatformInputMappingFactory {
    this := &PlatformInputMappingFactory{}
    this.initFields()
    return this
}


func PlatformInputMappingFactory_GetInstance() *PlatformInputMappingFactory {
        return PlatformInputMappingFactory_Instance
}


func (this *PlatformInputMappingFactory) GetPersistentInputMappingInstance() IPersistentInputMapping {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IPersistentInputMapping) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
                }
            }()
            if this.InputToGameKeyMapping == PersistentInputMapping_GetNullInstance() {
                this.InputToGameKeyMapping = NewNormalJ2MEGameInputMapping()
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.InputToGameKeyMapping
}

