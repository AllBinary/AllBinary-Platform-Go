/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type GamePersistanceStrings struct {
    SAVED_GAME_RECORD_ID string
    SAVED_GAME_CONFIGURATION_RECORD_ID string
    DEFAULT_INPUT_MAPPING_RECORD_ID string
    SAVED_INPUT_CONFIGURATION_RECORD_ID string
}

var GamePersistanceStrings_Instance *GamePersistanceStrings = NewGamePersistanceStrings()

func (this *GamePersistanceStrings) initFields() {
    this.SAVED_GAME_RECORD_ID = "_SG"
    this.SAVED_GAME_CONFIGURATION_RECORD_ID = "_SGC"
    this.DEFAULT_INPUT_MAPPING_RECORD_ID = "_DIM"
    this.SAVED_INPUT_CONFIGURATION_RECORD_ID = "_SIC"
}

func NewGamePersistanceStrings() *GamePersistanceStrings {
    this := &GamePersistanceStrings{}
    this.initFields()
    return this
}


func GamePersistanceStrings_GetInstance() *GamePersistanceStrings {
        return GamePersistanceStrings_Instance
}

