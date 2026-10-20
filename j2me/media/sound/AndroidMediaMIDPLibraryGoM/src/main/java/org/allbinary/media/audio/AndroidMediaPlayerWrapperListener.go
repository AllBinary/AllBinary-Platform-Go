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

type AndroidMediaPlayerWrapperListener struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AndroidMediaPlayerWrapper *AndroidMediaPlayerWrapper
    MOnBufferingUpdateListener *OnBufferingUpdateListener
    MOnPreparedListener *OnPreparedListener
    MOnErrorListener *OnErrorListener
    MOnCompletionListener *OnCompletionListener
}


func (this *AndroidMediaPlayerWrapperListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AndroidMediaPlayerWrapper = AndroidMediaPlayerWrapper_NULL_ANDROID_MEDIA_PLAYER_WRAPPER
    this.MOnBufferingUpdateListener = NewAndroidMediaPlayerOnBufferingUpdateListener(this)
    this.MOnPreparedListener = NewAndroidMediaPlayerOnPreparedListener(this)
    this.MOnErrorListener = NewAndroidMediaPlayerOnErrorListener(this)
    this.MOnCompletionListener = NewAndroidMediaPlayerOnCompletionListener(this)
}


func NewAndroidMediaPlayerWrapperListener(androidMediaPlayerWrapper *AndroidMediaPlayerWrapper) *AndroidMediaPlayerWrapperListener {
    this := &AndroidMediaPlayerWrapperListener{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
            this.AndroidMediaPlayerWrapper = androidMediaPlayerWrapper
            var mediaPlayer *MediaPlayer = androidMediaPlayerWrapper.GetMediaPlayer(); _ = mediaPlayer
            mediaPlayer.SetOnCompletionListener(this.MOnCompletionListener)
            mediaPlayer.SetOnBufferingUpdateListener(this.MOnBufferingUpdateListener)
            mediaPlayer.SetOnPreparedListener(this.MOnPreparedListener)
            mediaPlayer.SetOnErrorListener(this.MOnErrorListener)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}

type MediaPlayerOnBufferingUpdateListener struct {
}

// Implements *OnBufferingUpdateListener

// The Java class MediaPlayerOnBufferingUpdateListener has subclasses so its variables are this interface
type IMediaPlayerOnBufferingUpdateListener interface {
    MediaPlayerOnBufferingUpdateListener_() *MediaPlayerOnBufferingUpdateListener
    OnBufferingUpdate(mediaPlayer *MediaPlayer, i int)
}

func (this *MediaPlayerOnBufferingUpdateListener) MediaPlayerOnBufferingUpdateListener_() *MediaPlayerOnBufferingUpdateListener {
    return this
}


func NewMediaPlayerOnBufferingUpdateListener() *MediaPlayerOnBufferingUpdateListener {
    this := &MediaPlayerOnBufferingUpdateListener{}
    return this
}
    

    func (this *AndroidMediaPlayerWrapperListener) OnBufferingUpdate(mediaPlayer *MediaPlayer, i int) {
    }

type MediaPlayerOnPreparedListener struct {
}

// Implements *OnPreparedListener

// The Java class MediaPlayerOnPreparedListener has subclasses so its variables are this interface
type IMediaPlayerOnPreparedListener interface {
    MediaPlayerOnPreparedListener_() *MediaPlayerOnPreparedListener
    OnPrepared(mp *MediaPlayer)
}

func (this *MediaPlayerOnPreparedListener) MediaPlayerOnPreparedListener_() *MediaPlayerOnPreparedListener {
    return this
}


func NewMediaPlayerOnPreparedListener() *MediaPlayerOnPreparedListener {
    this := &MediaPlayerOnPreparedListener{}
    return this
}
    

    func (this *AndroidMediaPlayerWrapperListener) OnPrepared(mp *MediaPlayer) {
    }

type MediaPlayerOnErrorListener struct {
}

// Implements *OnErrorListener

// The Java class MediaPlayerOnErrorListener has subclasses so its variables are this interface
type IMediaPlayerOnErrorListener interface {
    MediaPlayerOnErrorListener_() *MediaPlayerOnErrorListener
    OnError(mp *MediaPlayer, what int, extra int) bool
}

func (this *MediaPlayerOnErrorListener) MediaPlayerOnErrorListener_() *MediaPlayerOnErrorListener {
    return this
}


func NewMediaPlayerOnErrorListener() *MediaPlayerOnErrorListener {
    this := &MediaPlayerOnErrorListener{}
    return this
}
    

    func (this *AndroidMediaPlayerWrapperListener) OnError(mp *MediaPlayer, what int, extra int) bool {
            return false
    }

type MediaPlayerOnCompletionListener struct {
}

// Implements *OnCompletionListener

// The Java class MediaPlayerOnCompletionListener has subclasses so its variables are this interface
type IMediaPlayerOnCompletionListener interface {
    MediaPlayerOnCompletionListener_() *MediaPlayerOnCompletionListener
    OnCompletion(mp *MediaPlayer)
}

func (this *MediaPlayerOnCompletionListener) MediaPlayerOnCompletionListener_() *MediaPlayerOnCompletionListener {
    return this
}


func NewMediaPlayerOnCompletionListener() *MediaPlayerOnCompletionListener {
    this := &MediaPlayerOnCompletionListener{}
    return this
}
    

    func (this *AndroidMediaPlayerWrapperListener) OnCompletion(mp *MediaPlayer) {
    }

type AndroidMediaPlayerOnBufferingUpdateListener struct {
    MediaPlayerOnBufferingUpdateListener
    ON_BUFFERING_UPDATE string
    AndroidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener
}


func (this *AndroidMediaPlayerOnBufferingUpdateListener) initFields() {
        this.ON_BUFFERING_UPDATE = "onBufferingUpdate()"
        
}
    

    func NewAndroidMediaPlayerWrapperListener(androidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener) *AndroidMediaPlayerWrapperListener {
        this := &AndroidMediaPlayerWrapperListener{}
        this.MediaPlayerOnBufferingUpdateListener = *NewMediaPlayerOnBufferingUpdateListener()
        this.initFields()
            this.AndroidMediaPlayerWrapperListener = androidMediaPlayerWrapperListener
        return this
    }
    

    func (this *AndroidMediaPlayerWrapperListener) OnBufferingUpdate(mediaPlayer *MediaPlayer, i int) {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            logUtil.PutF(fmt.Sprint(NewStringMaker().Append("Update buffer: ").Appendint(i).Append("%")), this, this.ON_BUFFERING_UPDATE)
            this.AndroidMediaPlayerWrapperListener.AndroidMediaPlayerWrapper.Update(PlayerListener_DEVICE_UNAVAILABLE)
    }

type AndroidMediaPlayerOnPreparedListener struct {
    MediaPlayerOnPreparedListener
    ON_PREPARE string
    AndroidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener
}


func (this *AndroidMediaPlayerOnPreparedListener) initFields() {
        this.ON_PREPARE = "onPrepare()"
        
}
    

    func NewAndroidMediaPlayerWrapperListener(androidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener) *AndroidMediaPlayerWrapperListener {
        this := &AndroidMediaPlayerWrapperListener{}
        this.MediaPlayerOnPreparedListener = *NewMediaPlayerOnPreparedListener()
        this.initFields()
            this.AndroidMediaPlayerWrapperListener = androidMediaPlayerWrapperListener
        return this
    }
    

    func (this *AndroidMediaPlayerWrapperListener) OnPrepared(mp *MediaPlayer) {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.PutF(commonStrings.START, this, this.ON_PREPARE)
            this.AndroidMediaPlayerWrapperListener.AndroidMediaPlayerWrapper.Update(PlayerListener_DEVICE_AVAILABLE)
    }

type AndroidMediaPlayerOnErrorListener struct {
    MediaPlayerOnErrorListener
    ON_ERROR string
    AndroidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener
}


func (this *AndroidMediaPlayerOnErrorListener) initFields() {
        this.ON_ERROR = "onError()"
        
}
    

    func NewAndroidMediaPlayerWrapperListener(androidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener) *AndroidMediaPlayerWrapperListener {
        this := &AndroidMediaPlayerWrapperListener{}
        this.MediaPlayerOnErrorListener = *NewMediaPlayerOnErrorListener()
        this.initFields()
            this.AndroidMediaPlayerWrapperListener = androidMediaPlayerWrapperListener
        return this
    }
    

    func (this *AndroidMediaPlayerWrapperListener) OnError(mp *MediaPlayer, what int, extra int) bool {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            logUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append("What: ").Appendint(what).Append(" Extra: ").Appendint(extra)), this, this.ON_ERROR)
            this.AndroidMediaPlayerWrapperListener.AndroidMediaPlayerWrapper.Update(PlayerListener_ERROR)
            return true
    }

type AndroidMediaPlayerOnCompletionListener struct {
    MediaPlayerOnCompletionListener
    ON_COMPLETE string
    AndroidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener
}


func (this *AndroidMediaPlayerOnCompletionListener) initFields() {
        this.ON_COMPLETE = "onComplete()"
        
}
    

    func NewAndroidMediaPlayerWrapperListener(androidMediaPlayerWrapperListener *AndroidMediaPlayerWrapperListener) *AndroidMediaPlayerWrapperListener {
        this := &AndroidMediaPlayerWrapperListener{}
        this.MediaPlayerOnCompletionListener = *NewMediaPlayerOnCompletionListener()
        this.initFields()
            this.AndroidMediaPlayerWrapperListener = androidMediaPlayerWrapperListener
        return this
    }
    

    func (this *AndroidMediaPlayerWrapperListener) OnCompletion(mp *MediaPlayer) {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.PutF(commonStrings.START, this, this.ON_COMPLETE)
            this.AndroidMediaPlayerWrapperListener.AndroidMediaPlayerWrapper.Update(PlayerListener_END_OF_MEDIA)
    }

