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

type GameGlobalsFactory struct {
    NewCanvas bool
    NewDisplaybleTime int64
    Point IGPoint
}

var GameGlobalsFactory_Instance *GameGlobalsFactory = NewGameGlobalsFactory()

func (this *GameGlobalsFactory) initFields() {
    this.NewCanvas = true
    this.NewDisplaybleTime = int64(Integer_MIN_VALUE)
    this.Point = PointFactory_GetInstance().ZERO_ZERO
}

func NewGameGlobalsFactory() *GameGlobalsFactory {
    this := &GameGlobalsFactory{}
    this.initFields()
    return this
}


func GameGlobalsFactory_GetInstance() *GameGlobalsFactory {
        return GameGlobalsFactory_Instance
}

