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

type GameMode struct {
    Name string
}

var GameMode_NONE *GameMode = NewGameMode(StringUtil_GetInstance().NULL_STRING)
var GameMode_SERVER *GameMode = NewGameMode("Server")
var GameMode_CLIENT *GameMode = NewGameMode("Client")

func (this *GameMode) initFields() {
    this.Name = StringUtil_GetInstance().EMPTY_STRING
}


func NewGameMode(name string) *GameMode {
    this := &GameMode{}
    this.initFields()
        this.SetName(name)
    return this
}


func (this *GameMode) SetName(name string) {
        this.Name = name
}


func (this *GameMode) GetName() string {
        return this.Name
}


func (this *GameMode) ToString() string {
        return this.GetName()
}

func (this *GameMode) String() string {
    return this.ToString()
}

