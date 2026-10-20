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

type J2MEKeyFactory struct {
    GameKeyFactory *GameKeyFactory
}

var J2MEKeyFactory_Instance *J2MEKeyFactory = NewJ2MEKeyFactory()

func (this *J2MEKeyFactory) initFields() {
    this.GameKeyFactory = GameKeyFactory_GetInstance()
}


func J2MEKeyFactory_GetInstance() *J2MEKeyFactory {
        return J2MEKeyFactory_Instance
}


func NewJ2MEKeyFactory() *J2MEKeyFactory {
    this := &J2MEKeyFactory{}
    this.initFields()
    return this
}


func (this *J2MEKeyFactory) Init() {
        var gameKeyArray []*GameKey = []*GameKey{this.GameKeyFactory.DOWN, this.GameKeyFactory.FIRE, this.GameKeyFactory.GAME_A, this.GameKeyFactory.GAME_B, this.GameKeyFactory.GAME_C, this.GameKeyFactory.GAME_D, this.GameKeyFactory.KEY_NUM0, this.GameKeyFactory.KEY_NUM1, this.GameKeyFactory.KEY_NUM2, this.GameKeyFactory.KEY_NUM3, this.GameKeyFactory.KEY_NUM4, this.GameKeyFactory.KEY_NUM5, this.GameKeyFactory.KEY_NUM6, this.GameKeyFactory.KEY_NUM7, this.GameKeyFactory.KEY_NUM8, this.GameKeyFactory.KEY_NUM9, this.GameKeyFactory.KEY_POUND, this.GameKeyFactory.KEY_STAR, this.GameKeyFactory.LEFT, this.GameKeyFactory.LEVEL_DOWN, this.GameKeyFactory.LEVEL_UP, this.GameKeyFactory.NONE, this.GameKeyFactory.RIGHT, this.GameKeyFactory.UP}; _ = gameKeyArray
        var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
        for index := int(len(gameKeyArray)); Java_PreDecrement(&index) >= 0;  {
            inputFactory.Add(gameKeyArray[index].GetId(), gameKeyArray[index])
        }
}


func (this *J2MEKeyFactory) IsSubmission(input IInput) bool {
        if input == this.GameKeyFactory.FIRE {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsDelete(input IInput) bool {
        if input == this.GameKeyFactory.GAME_D {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsBackSpace(input IInput) bool {
        return false
}


func (this *J2MEKeyFactory) IsUp(input IInput) bool {
        if input == this.GameKeyFactory.UP {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsDown(input IInput) bool {
        if input == this.GameKeyFactory.DOWN {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsLeft(input IInput) bool {
        if input == this.GameKeyFactory.LEFT {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsRight(input IInput) bool {
        if input == this.GameKeyFactory.RIGHT {
            return true
        }
        return false
}


func (this *J2MEKeyFactory) IsEnter(input IInput) bool {
        if input == this.GameKeyFactory.FIRE {
            return true
        }
        return false
}

