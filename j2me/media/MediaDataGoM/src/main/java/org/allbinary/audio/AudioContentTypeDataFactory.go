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

type AudioContentTypeDataFactory struct {
    MIME_AUDIO_TONE *AudioContentTypeData
    MIME_AUDIO_AMR *AudioContentTypeData
    MIME_AUDIO_AMR_WB *AudioContentTypeData
    MIME_AUDIO_WAV *AudioContentTypeData
    MIME_AUDIO_MIDI *AudioContentTypeData
    MIME_AUDIO_SP_MIDI *AudioContentTypeData
}

var AudioContentTypeDataFactory_Instance *AudioContentTypeDataFactory = NewAudioContentTypeDataFactory()

func (this *AudioContentTypeDataFactory) initFields() {
    this.MIME_AUDIO_TONE = NewAudioContentTypeData("audio/x-tone-seq")
    this.MIME_AUDIO_AMR = NewAudioContentTypeData("audio/amr")
    this.MIME_AUDIO_AMR_WB = NewAudioContentTypeData("audio/amr-wb")
    this.MIME_AUDIO_WAV = NewAudioContentTypeData("audio/x-wav")
    this.MIME_AUDIO_MIDI = NewAudioContentTypeData("audio/midi")
    this.MIME_AUDIO_SP_MIDI = NewAudioContentTypeData("audio/sp-midi")
}


func AudioContentTypeDataFactory_GetInstance() *AudioContentTypeDataFactory {
        return AudioContentTypeDataFactory_Instance
}


func NewAudioContentTypeDataFactory() *AudioContentTypeDataFactory {
    this := &AudioContentTypeDataFactory{}
    this.initFields()
    return this
}

