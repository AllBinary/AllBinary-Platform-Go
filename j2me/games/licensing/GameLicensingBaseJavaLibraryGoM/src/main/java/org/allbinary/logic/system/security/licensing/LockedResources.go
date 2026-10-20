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

type LockedResources struct {
    LOCKED_DEMO_GAME_FEATURE_RESOURCE string
}

var LockedResources_Instance *LockedResources = NewLockedResources()

func (this *LockedResources) initFields() {
    this.LOCKED_DEMO_GAME_FEATURE_RESOURCE = "/locked_demo_game_feature_64_by_64.png"
}

func NewLockedResources() *LockedResources {
    this := &LockedResources{}
    this.initFields()
    return this
}


func LockedResources_GetInstance() *LockedResources {
        return LockedResources_Instance
}

