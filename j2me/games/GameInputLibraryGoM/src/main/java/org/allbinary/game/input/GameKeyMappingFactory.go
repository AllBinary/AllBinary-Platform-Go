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

type GameKeyMappingFactory struct {
    IntKeyToGameKey []*GameKey
}

var GameKeyMappingFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameKeyMappingFactory) initFields() {
    this.IntKeyToGameKey = make([]*GameKey, InputFactory_GetInstance().MAX)
}


func GameKeyMappingFactory_GetInstance() *GameKeyMappingFactory {
        if GameKeyMappingFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameKeyMappingFactory_Instance = NewGameKeyMappingFactory()
        }
        return any(GameKeyMappingFactory_Instance).(*GameKeyMappingFactory)
}


func NewGameKeyMappingFactory() *GameKeyMappingFactory {
    this := &GameKeyMappingFactory{}
    this.initFields()
        var gameKeyFactory *GameKeyFactory = GameKeyFactory_GetInstance(); _ = gameKeyFactory
        this.SetGameKey(gameKeyFactory.NONE)
        this.SetGameKey(gameKeyFactory.UP)
        this.SetGameKey(gameKeyFactory.DOWN)
        this.SetGameKey(gameKeyFactory.LEFT)
        this.SetGameKey(gameKeyFactory.RIGHT)
        this.SetGameKey(gameKeyFactory.KEY_POUND)
        this.SetGameKey(gameKeyFactory.KEY_STAR)
        this.SetGameKey(gameKeyFactory.KEY_NUM0)
        this.SetGameKey(gameKeyFactory.KEY_NUM1)
        this.SetGameKey(gameKeyFactory.KEY_NUM2)
        this.SetGameKey(gameKeyFactory.KEY_NUM3)
        this.SetGameKey(gameKeyFactory.KEY_NUM4)
        this.SetGameKey(gameKeyFactory.KEY_NUM5)
        this.SetGameKey(gameKeyFactory.KEY_NUM6)
        this.SetGameKey(gameKeyFactory.KEY_NUM7)
        this.SetGameKey(gameKeyFactory.KEY_NUM8)
        this.SetGameKey(gameKeyFactory.KEY_NUM9)
        this.SetGameKey(gameKeyFactory.GAME_A)
        this.SetGameKey(gameKeyFactory.GAME_B)
        this.SetGameKey(gameKeyFactory.GAME_C)
        this.SetGameKey(gameKeyFactory.GAME_D)
        this.SetGameKey(gameKeyFactory.FIRE)
        this.SetGameKey(gameKeyFactory.LEVEL_DOWN)
        this.SetGameKey(gameKeyFactory.LEVEL_UP)
    return this
}


func (this *GameKeyMappingFactory) GetGameKey(key int) *GameKey {
        return this.IntKeyToGameKey[key]
}


func (this *GameKeyMappingFactory) SetGameKey(gameKey *GameKey) {
        this.IntKeyToGameKey[gameKey.GetId()] = gameKey
}

