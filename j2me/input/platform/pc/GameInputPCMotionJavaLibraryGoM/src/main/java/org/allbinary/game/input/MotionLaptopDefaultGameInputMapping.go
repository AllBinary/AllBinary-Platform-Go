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

type MotionLaptopDefaultGameInputMapping struct {
    PersistentInputMapping
}



func NewMotionLaptopDefaultGameInputMapping() *MotionLaptopDefaultGameInputMapping {
    this := &MotionLaptopDefaultGameInputMapping{}
        this.PersistentInputMapping = *NewPersistentInputMapping(GamePersistanceStrings_GetInstance().DEFAULT_INPUT_MAPPING_RECORD_ID)
        var pcKeyFactory *PCKeyFactory = PCKeyFactory_GetInstance(); _ = pcKeyFactory
        var inputToGameKeyMapping *InputToGameKeyMapping = this.GetInputMapping(); _ = inputToGameKeyMapping
        var gameKeyFactory *GameKeyFactory = GameKeyFactory_GetInstance(); _ = gameKeyFactory
        var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
        var basicTouchInputFactory *BasicTouchInputFactory = BasicTouchInputFactory_GetInstance(); _ = basicTouchInputFactory
        inputToGameKeyMapping.Add(gameKeyFactory.FIRE, pcKeyFactory.ENTER)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, pcKeyFactory.FOUR)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, pcKeyFactory._P)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, pcKeyFactory.P)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, pcKeyFactory.DPAD_UP)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, pcKeyFactory.DPAD_DOWN)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, pcKeyFactory.DPAD_LEFT)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, pcKeyFactory.DPAD_RIGHT)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, pcKeyFactory._Y)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, pcKeyFactory._I)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, pcKeyFactory._U)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, pcKeyFactory._H)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, pcKeyFactory._J)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, pcKeyFactory._N)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM5, pcKeyFactory._M)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, pcKeyFactory.Y)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, pcKeyFactory.I)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, pcKeyFactory.U)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, pcKeyFactory.H)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, pcKeyFactory.J)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, pcKeyFactory.N)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM5, pcKeyFactory.M)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, pcKeyFactory._Q)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, pcKeyFactory._E)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, pcKeyFactory._W)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, pcKeyFactory._A)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, pcKeyFactory._D)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, pcKeyFactory._S)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM5, pcKeyFactory._X)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, pcKeyFactory.Q)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, pcKeyFactory.E)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, pcKeyFactory.W)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, pcKeyFactory.A)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, pcKeyFactory.D)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, pcKeyFactory.S)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM5, pcKeyFactory.X)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, touchMotionGestureFactory.UP)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, touchMotionGestureFactory.DOWN)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, touchMotionGestureFactory.LEFT)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, touchMotionGestureFactory.RIGHT)
        inputToGameKeyMapping.Add(gameKeyFactory.UP, basicTouchInputFactory.UP)
        inputToGameKeyMapping.Add(gameKeyFactory.LEFT, basicTouchInputFactory.LEFT)
        inputToGameKeyMapping.Add(gameKeyFactory.RIGHT, basicTouchInputFactory.RIGHT)
        inputToGameKeyMapping.Add(gameKeyFactory.DOWN, basicTouchInputFactory.DOWN)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM0, basicTouchInputFactory.SPECIAL_BUTTON_FIVE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM7, basicTouchInputFactory.SPECIAL_BUTTON_FOUR)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM9, basicTouchInputFactory.SPECIAL_BUTTON_THREE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM1, basicTouchInputFactory.SPECIAL_BUTTON_TWO)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM3, basicTouchInputFactory.SPECIAL_BUTTON_ONE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM2, basicTouchInputFactory.SPECIAL_BUTTON_SIX)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM3, pcKeyFactory.THREE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM0, pcKeyFactory.ZERO)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM0, pcKeyFactory.ONE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_NUM0, pcKeyFactory.SPACE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_POUND, pcKeyFactory.DEL)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_POUND, pcKeyFactory.TWO)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_STAR, pcKeyFactory.ESCAPE)
        inputToGameKeyMapping.Add(gameKeyFactory.KEY_STAR, pcKeyFactory.STAR)
        inputToGameKeyMapping.Add(gameKeyFactory.LEVEL_DOWN, pcKeyFactory.COMMA)
        inputToGameKeyMapping.Add(gameKeyFactory.LEVEL_UP, pcKeyFactory.PERIOD)
    return this
}

