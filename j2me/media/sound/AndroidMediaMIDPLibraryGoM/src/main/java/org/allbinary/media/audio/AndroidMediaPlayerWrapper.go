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

type AndroidMediaPlayerWrapper struct {
    BasicPlayer
    LogUtil *LogUtil
    MediaPlayer *MediaPlayer
}

var AndroidMediaPlayerWrapper_NULL_ANDROID_MEDIA_PLAYER_WRAPPER *AndroidMediaPlayerWrapper = AndroidMediaPlayerWrapper_CreatePlayerWrapper()

func (this *AndroidMediaPlayerWrapper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MediaPlayer = NullAndroidCanvas_NULL_MEDIA_PLAYER
}


func AndroidMediaPlayerWrapper_CreatePlayerWrapper() *AndroidMediaPlayerWrapper {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *AndroidMediaPlayerWrapper) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(errors.New(fmt.Sprint()))
                }
            }()
            tryControl1, tryResult1 = 1, NewAndroidMediaPlayerWrapper(StringUtil_GetInstance().EMPTY_STRING); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func NewAndroidMediaPlayerWrapper(resource string) *AndroidMediaPlayerWrapper {
    this := &AndroidMediaPlayerWrapper{}
    this.BasicPlayer = *NewBasicPlayer()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(fmt.Sprint(this.CommonStrings.EXCEPTION_LABEL) + resource, this, this.CommonStrings.CONSTRUCTOR, e)
            panic(e)
                }
            }()
            if resource == StringUtil_GetInstance().EMPTY_STRING {
            } else {
                var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
                var mediaPlayer *MediaPlayer = MediaPlayer_Create(resourceUtil.GetContext(), int(resourceUtil.GetResourceId(resource))); _ = mediaPlayer
                if mediaPlayer == nil {
                    panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Failed to create media player for: ").Append(resource).Append(" with id: ").Append(fmt.Sprint(resourceUtil.GetResourceId(resource)))))))
                }
                this.SetMediaPlayer(mediaPlayer)
                this.MediaPlayer.SetLooping(false)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *AndroidMediaPlayerWrapper) SetLoopCount(count int) {
        this.BasicPlayer.SetLoopCount(count)
        if this.MediaPlayer != NullAndroidCanvas_NULL_MEDIA_PLAYER && this.MediaPlayer != nil {
            if count == 0 {
                this.MediaPlayer.SetLooping(false)
            } else {
            }
        }
}


// synchronized
func (this *AndroidMediaPlayerWrapper) AddPlayerListener(playerListener PlayerListener) {
        this.BasicPlayer.AddPlayerListener(playerListener)
}


func (this *AndroidMediaPlayerWrapper) RemovePlayerListener(playerListener PlayerListener) {
        this.BasicPlayer.RemovePlayerListener(playerListener)
}


func (this *AndroidMediaPlayerWrapper) GetState() int {
        if this.MediaPlayer.IsPlaying() {
            return Player_STARTED
        } else {
            return Player_PREFETCHED
        }
}


func (this *AndroidMediaPlayerWrapper) Close() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CLOSE, e)
                }
            }()
            this.MediaPlayer.Release()
            this.MediaPlayer = NullAndroidCanvas_NULL_MEDIA_PLAYER
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws MediaException
// synchronized
func (this *AndroidMediaPlayerWrapper) Start() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.START_METHOD_NAME, e)
                }
            }()
            if this.MediaPlayer.IsPlaying() {
                this.MediaPlayer.Pause()
                this.MediaPlayer.SeekTo(0)
            }
            this.MediaPlayer.Start()
            this.BasicPlayer.Start()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws MediaException
// synchronized
func (this *AndroidMediaPlayerWrapper) Stop() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "stop", e)
                }
            }()
            this.MediaPlayer.Stop()
            this.MediaPlayer.Prepare()
            this.BasicPlayer.Stop()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidMediaPlayerWrapper) Update(event string) {
        this.LogUtil.PutF("LineEvent: " + event, this, this.CommonStrings.UPDATE)
        var size int = this.ListenersList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var listener PlayerListener = any(this.ListenersList.BasicArrayList_().ObjectArray[index]).(PlayerListener); _ = listener
            listener.PlayerUpdate(this, event, NullUtil_GetInstance().NULL_OBJECT)
        }
}


func (this *AndroidMediaPlayerWrapper) SetVolume(leftVolume int, rightVolume int) {
        this.MediaPlayer.SetVolume((float32(leftVolume)) / 100.0, (float32(rightVolume)) / 100.0)
}


func (this *AndroidMediaPlayerWrapper) GetDuration() int64 {
        return int64(this.MediaPlayer.GetDuration())
}


func (this *AndroidMediaPlayerWrapper) SetMediaPlayer(mediaPlayer *MediaPlayer) {
        this.MediaPlayer = mediaPlayer
}


func (this *AndroidMediaPlayerWrapper) GetMediaPlayer() *MediaPlayer {
        return this.MediaPlayer
}

