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

type BaseMusicService struct {
    Service
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonStateStrings *CommonStateStrings
    ALREADY_PLAYING string
    WAITING_FOR_MUSIC_TO_END string
    Player *MediaPlayer
    SongId int
    LeftVolume int
    RightVolume int
}

// The Java class BaseMusicService has subclasses so its variables are this interface
type IBaseMusicService interface {
    BaseMusicService_() *BaseMusicService
    OnBind(intent *Intent) *IBinder
    OnCreate()
    OnDestroy()
    Pause()
    Resume()
    Start()
    OnStart(intent *Intent, startid int)
    OnStartCommand(intent *Intent, flags int, startId int) int
    OnStartCommandIntent(intent *Intent)
}

func (this *BaseMusicService) BaseMusicService_() *BaseMusicService {
    return this
}


func (this *BaseMusicService) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonStateStrings = CommonStateStrings_GetInstance()
    this.ALREADY_PLAYING = "This is one song per music service"
    this.WAITING_FOR_MUSIC_TO_END = "Waiting for music to end"
    this.Player = NullAndroidCanvas_NULL_MEDIA_PLAYER
    this.SongId = -1
    this.LeftVolume = -1
    this.RightVolume = -1
}

func NewBaseMusicService() *BaseMusicService {
    this := &BaseMusicService{}
    this.Service = *NewService()
    this.initFields()
    return this
}


func (this *BaseMusicService) OnBind(intent *Intent) *IBinder {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.BIND)
        return nil
}


func (this *BaseMusicService) OnCreate() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.CREATE)
}


func (this *BaseMusicService) OnDestroy() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.DESTROY)
        if this.Player != NullAndroidCanvas_NULL_MEDIA_PLAYER {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.PAUSE)
            this.Player.Stop()
            this.Player.Reset()
            this.Player.Release()
        }
}


func (this *BaseMusicService) Pause() {
        if this.Player != NullAndroidCanvas_NULL_MEDIA_PLAYER {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.PAUSE)
            this.Player.Pause()
        }
}


func (this *BaseMusicService) Resume() {
        if this.Player != NullAndroidCanvas_NULL_MEDIA_PLAYER && !this.Player.IsPlaying() {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.RESUME)
            this.Player.Start()
        }
}


func (this *BaseMusicService) Start() {
        this.Player = MediaPlayer_Create(this, this.SongId)
        this.Player.SetVolume((float32(this.LeftVolume)) / 100.0, (float32(this.RightVolume)) / 100.0)
        this.Player.SetLooping(false)
        this.Player.Start()
}


func (this *BaseMusicService) OnStart(intent *Intent, startid int) {
        this.OnStartCommandIntent(intent)
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.START)
}


func (this *BaseMusicService) OnStartCommand(intent *Intent, flags int, startId int) int {
        this.OnStartCommandIntent(intent)
        return Service_START_STICKY
}


func (this *BaseMusicService) OnStartCommandIntent(intent *Intent) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStateStrings.ON_START_COMMAND)
        var musicStrings *MusicStrings = MusicStrings_GetInstance(); _ = musicStrings
        if intent != nil {
            var command int = intent.GetIntExtra(this.CommonStateStrings.ON_START_COMMAND, -1); _ = command
            this.LogUtil.PutF(CommonLabels_GetInstance().COMMAND_LABEL + command, this, this.CommonStateStrings.ON_START_COMMAND)
            if command == 1 {
                this.Pause()
                return
            } else if command == 2 {
                this.Resume()
                return
            } else {
                this.SongId = intent.GetIntExtra(musicStrings.SONG_EXTRA, -1)
                this.LeftVolume = intent.GetIntExtra(musicStrings.LEFT_VOLUME, -1)
                this.RightVolume = intent.GetIntExtra(musicStrings.RIGHT_VOLUME, -1)
            }
        } else {
            panic(errors.New(fmt.Sprint("Started service without intent")))
        }
        if this.SongId != -1 {
            runtime.GC()
            if this.Player != NullAndroidCanvas_NULL_MEDIA_PLAYER && this.Player.IsPlaying() {
                var player *MediaPlayer = this.Player; _ = player
                this.LogUtil.PutF(this.ALREADY_PLAYING, this, this.CommonStateStrings.ON_START_COMMAND)
                /* Java local class: move it out of the function
type MusicRunnable struct {
    ARunnable
    BaseMusicService IBaseMusicService
}


func (this *MusicRunnable) initFields() {
                    
}
                

                func NewBaseMusicService(baseMusicService IBaseMusicService) *BaseMusicService {
                    this := &BaseMusicService{}
                    this.ARunnable = *NewARunnable()
                    this.initFields()
                        this.BaseMusicService = baseMusicService
                    return this
                }
                

                func (this *BaseMusicService) Run() {
                        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                        var commonStateStrings *CommonStateStrings = CommonStateStrings_GetInstance(); _ = commonStateStrings
                        if tryControl1 := func() (tryControl1 int) {
                            defer func() {
                                if recovered := recover(); recovered != nil {
                                    e := Java_ToError(recovered)
                                    _ = e
                            logUtil.Put(commonStrings.EXCEPTION, this, commonStateStrings.ON_START_COMMAND, e)
                                }
                            }()
                            for this.Player.IsPlaying() {
                                logUtil.PutF(this.BaseMusicService.WAITING_FOR_MUSIC_TO_END, this, commonStateStrings.ON_START_COMMAND)
                                time.Sleep(time.Duration(1200) * time.Millisecond)
                            }
                            this.BaseMusicService.OnStartCommandIntent(this.Intent)
                            return
                        }(); tryControl1 == 1 {
                            return
                        }
                }
*/
                var runnable Runnable = NewMusicRunnable(this); _ = runnable
                var thread IThread = NewThread(runnable); _ = thread
                thread.Start()
                return
            }
            this.Start()
        }
}

