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

type NullPlayerGameInputComposite struct {
}

// Implements PlayerGameInputCompositeInterface

var NullPlayerGameInputComposite_NULL_PLAYER_GAME_INPUT_COMPOSITE *NullPlayerGameInputComposite = NewNullPlayerGameInputComposite()

func NewNullPlayerGameInputComposite() *NullPlayerGameInputComposite {
    this := &NullPlayerGameInputComposite{}
    return this
}


func (this *NullPlayerGameInputComposite) GetName() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


// throws Exception
func (this *NullPlayerGameInputComposite) ProcessInput(layerManager IAllBinaryLayerManager) {
}


func (this *NullPlayerGameInputComposite) InitInputProcessors() {
}


func (this *NullPlayerGameInputComposite) GetPlayerGameInput() IPlayerGameInput {
        return NoPlayerGameInput_GetInstance()
}

