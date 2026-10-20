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

type Sound struct {
    Resource string
    Player *Player
}

// Implements SoundInterface, Runnable

// The Java class Sound has subclasses so its variables are this interface
type ISound interface {
    Sound_() *Sound
    GetResource() string
    GetPlayerP() *Player
    SetPlayerP(player *Player)
    Init()
    GetDuration() int
    Run()
}

func (this *Sound) Sound_() *Sound {
    return this
}


func (this *Sound) initFields() {
    
    this.Player = NoPlayer_NO_PLAYER
}


func NewSound(resource string) *Sound {
    this := &Sound{}
    this.initFields()
        this.Resource = resource
    return this
}


func (this *Sound) GetResource() string {
        return this.Resource
}


func (this *Sound) GetPlayerP() *Player {
        return this.Player
}


func (this *Sound) SetPlayerP(player *Player) {
        this.Player = player
}


// throws Exception
func (this *Sound) Init() {
}


func (this *Sound) GetDuration() int {
        return -1
}


func (this *Sound) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_Put(commonStrings.EXCEPTION, this, commonStrings.RUN)
                }
            }()
            this.Player.Start()
            return
        }(); tryControl1 == 1 {
            return
        }
}

