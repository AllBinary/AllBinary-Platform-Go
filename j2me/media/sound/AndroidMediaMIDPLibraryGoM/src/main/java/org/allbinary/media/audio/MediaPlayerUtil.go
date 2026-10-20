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

type MediaPlayerUtil struct {
    LogUtil *LogUtil
}

var MediaPlayerUtil_Instance *MediaPlayerUtil = NewMediaPlayerUtil()

func (this *MediaPlayerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewMediaPlayerUtil() *MediaPlayerUtil {
    this := &MediaPlayerUtil{}
    this.initFields()
    return this
}


func MediaPlayerUtil_GetInstance() *MediaPlayerUtil {
        return MediaPlayerUtil_Instance
}


// throws Exception
func (this *MediaPlayerUtil) WaitForMediaPlayer(mediaPlayer *MediaPlayer) {
        var MESSAGE string = "Not Stopped Waiting"; _ = MESSAGE
        var METHOD_NAME string = "wait"; _ = METHOD_NAME
        var index int = 0; _ = index
        for mediaPlayer.IsPlaying() && index < 50 {
            this.LogUtil.PutF(MESSAGE, this, METHOD_NAME)
            time.Sleep(time.Duration(100) * time.Millisecond)
            index++
        }
}

