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

type NormalJ2MEDefaultGameInputMapping struct {
    PersistentInputMapping
}



func NewNormalJ2MEDefaultGameInputMapping() *NormalJ2MEDefaultGameInputMapping {
    this := &NormalJ2MEDefaultGameInputMapping{}
        this.PersistentInputMapping = *NewPersistentInputMapping(GamePersistanceStrings_GetInstance().DEFAULT_INPUT_MAPPING_RECORD_ID)
        var inputToGameKeyMapping *InputToGameKeyMapping = this.GetInputMapping(); _ = inputToGameKeyMapping
        var gameKeyFactory *GameKeyFactory = GameKeyFactory_GetInstance(); _ = gameKeyFactory
        inputToGameKeyMapping.Add(gameKeyFactory.UP, gameKeyFactory.KEY_NUM2)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, gameKeyFactory.KEY_NUM4)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, gameKeyFactory.KEY_NUM6)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, gameKeyFactory.KEY_NUM8)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, gameKeyFactory.GAME_A)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM3, gameKeyFactory.GAME_B)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, gameKeyFactory.GAME_C)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, gameKeyFactory.GAME_D)
        inputToGameKeyMapping.Add(gameKeyFactory.LEVEL_DOWN, gameKeyFactory.LEVEL_DOWN)
        inputToGameKeyMapping.Add(gameKeyFactory.LEVEL_UP, gameKeyFactory.LEVEL_UP)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, gameKeyFactory.UP)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, gameKeyFactory.DOWN)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, gameKeyFactory.LEFT)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, gameKeyFactory.RIGHT)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_POUND, gameKeyFactory.KEY_POUND)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_STAR, gameKeyFactory.KEY_STAR)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, gameKeyFactory.KEY_NUM1)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM3, gameKeyFactory.KEY_NUM3)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM5, gameKeyFactory.KEY_NUM5)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, gameKeyFactory.KEY_NUM7)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, gameKeyFactory.KEY_NUM9)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM0, gameKeyFactory.KEY_NUM0)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, gameKeyFactory.FIRE)
    return this
}

