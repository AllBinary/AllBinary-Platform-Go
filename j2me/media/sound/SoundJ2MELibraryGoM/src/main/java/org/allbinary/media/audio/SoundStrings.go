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

type SoundStrings struct {
    STOP_ALL string
    CLOSE_ALL string
    SOUND string
    INIT_SOUND string
    STOPPING_SOUND string
    CLOSING_SOUND string
}

var SoundStrings_Instance *SoundStrings = NewSoundStrings()

func (this *SoundStrings) initFields() {
    this.STOP_ALL = "stopAll"
    this.CLOSE_ALL = "closeAll"
    this.SOUND = " Sound: "
    this.INIT_SOUND = "Initializing Sound: "
    this.STOPPING_SOUND = "Stopping Sound: "
    this.CLOSING_SOUND = "Closing Sound: "
}

func NewSoundStrings() *SoundStrings {
    this := &SoundStrings{}
    this.initFields()
    return this
}


func SoundStrings_GetInstance() *SoundStrings {
        return SoundStrings_Instance
}

