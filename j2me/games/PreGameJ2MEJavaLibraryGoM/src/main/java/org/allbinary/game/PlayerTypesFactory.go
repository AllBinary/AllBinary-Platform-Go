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

type PlayerTypesFactory struct {
    PLAYER_TYPE_ONE *PlayerType
    PLAYER_TYPE_TWO *PlayerType
    PLAYER_TYPE_THREE *PlayerType
}

var PlayerTypesFactory_SINGLETON *PlayerTypesFactory = NewPlayerTypesFactory()

func (this *PlayerTypesFactory) initFields() {
    this.PLAYER_TYPE_ONE = NewPlayerType("Player Type One", 0)
    this.PLAYER_TYPE_TWO = NewPlayerType("Player Type Two", 1)
    this.PLAYER_TYPE_THREE = NewPlayerType("Player Type Three", 2)
}

func NewPlayerTypesFactory() *PlayerTypesFactory {
    this := &PlayerTypesFactory{}
    this.initFields()
    return this
}


func PlayerTypesFactory_GetInstance() *PlayerTypesFactory {
        return PlayerTypesFactory_SINGLETON
}

