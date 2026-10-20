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

type MusicStrings struct {
    SONG_EXTRA string
    LEFT_VOLUME string
    RIGHT_VOLUME string
}

var MusicStrings_Instance *MusicStrings = NewMusicStrings()

func (this *MusicStrings) initFields() {
    this.SONG_EXTRA = "SONG"
    this.LEFT_VOLUME = "LEFT_VOLUME"
    this.RIGHT_VOLUME = "RIGHT_VOLUME"
}

func NewMusicStrings() *MusicStrings {
    this := &MusicStrings{}
    this.initFields()
    return this
}


func MusicStrings_GetInstance() *MusicStrings {
        return MusicStrings_Instance
}

