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

type GameTickDisplayInfoSingleton struct {
    DisplayInfoSingleton *DisplayInfoSingleton
    LastWidth int
    LastHeight int
    LastHalfWidth int
    LastHalfHeight int
}

var GameTickDisplayInfoSingleton_Instance *GameTickDisplayInfoSingleton = NewGameTickDisplayInfoSingleton()

func (this *GameTickDisplayInfoSingleton) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.LastWidth = this.DisplayInfoSingleton.GetLastWidth()
    this.LastHeight = this.DisplayInfoSingleton.GetLastHeight()
    this.LastHalfWidth = this.DisplayInfoSingleton.GetLastHalfWidth()
    this.LastHalfHeight = this.DisplayInfoSingleton.GetLastHalfHeight()
}

func NewGameTickDisplayInfoSingleton() *GameTickDisplayInfoSingleton {
    this := &GameTickDisplayInfoSingleton{}
    this.initFields()
    return this
}


func GameTickDisplayInfoSingleton_GetInstance() *GameTickDisplayInfoSingleton {
        return GameTickDisplayInfoSingleton_Instance
}


func (this *GameTickDisplayInfoSingleton) Update() {
        this.LastWidth = this.DisplayInfoSingleton.GetLastWidth()
        this.LastHeight = this.DisplayInfoSingleton.GetLastHeight()
        this.LastHalfWidth = this.DisplayInfoSingleton.GetLastHalfWidth()
        this.LastHalfHeight = this.DisplayInfoSingleton.GetLastHalfHeight()
}


func (this *GameTickDisplayInfoSingleton) GetLastWidth() int {
        return this.LastWidth
}


func (this *GameTickDisplayInfoSingleton) GetLastHeight() int {
        return this.LastHeight
}


func (this *GameTickDisplayInfoSingleton) GetLastHalfWidth() int {
        return this.LastHalfWidth
}


func (this *GameTickDisplayInfoSingleton) GetLastHalfHeight() int {
        return this.LastHalfHeight
}

