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

type MusicManager struct {
    PLAY string
    FOR string
    CommonStrings *CommonStrings
    MusicStrings *MusicStrings
    CommonStateStrings *CommonStateStrings
    BasicArrayListUtil *BasicArrayListUtil
    ResourceUtil *ResourceUtil
    AndroidServicesUtil *AndroidServicesUtil
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    TimeDelayHelper ITimeDelayHelper
    TimeDelayHelper2 ITimeDelayHelper
    SongList IBasicArrayList
    CurrentSongSound ISound
    NextSongSound ISound
    LeftVolume int
    RightVolume int
    MusicServiceClass reflect.Type
    CurrentIntent *Intent
}


func (this *MusicManager) initFields() {
    this.PLAY = "Play "
    this.FOR = " for: "
    this.CommonStrings = CommonStrings_GetInstance()
    this.MusicStrings = MusicStrings_GetInstance()
    this.CommonStateStrings = CommonStateStrings_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.ResourceUtil = ResourceUtil_GetInstance()
    this.AndroidServicesUtil = AndroidServicesUtil_GetInstance()
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(0)
    this.TimeDelayHelper2 = NewTimeDelayHelper(1200)
    
    this.CurrentSongSound = NoSound_GetInstance()
    this.NextSongSound = NoSound_GetInstance()
    this.LeftVolume = 100
    this.RightVolume = 100
    
    
}


func NewMusicManager(musicServiceClass reflect.Type, songList IBasicArrayList) *MusicManager {
    this := &MusicManager{}
    this.initFields()
        PreLogUtil_Put(this.CommonStateStrings.CONTEXT + this.ResourceUtil.GetContext(), this, this.CommonStrings.CONSTRUCTOR)
        this.MusicServiceClass = musicServiceClass
        this.CurrentIntent = NewIntent(this.ResourceUtil.GetContext(), musicServiceClass)
        this.SongList = songList
    return this
}


func (this *MusicManager) NextSong(nextSongSound ISound, leftVolume int, rightVolume int) {
        if nextSongSound == nil {
            nextSongSound = NoSound_GetInstance()
        }
        this.NextSongSound = nextSongSound
        this.LeftVolume = leftVolume
        this.RightVolume = rightVolume
        this.Reset()
}


func (this *MusicManager) Reset() {
        this.TimeDelayHelper.TimeDelayHelper_().Delay = 0
}


func (this *MusicManager) Process() {
        if this.SongList.Size() == 0 {
            return
        }
        if this.TimeDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            this.StartNewSong()
            return
        }
        if this.TimeDelayHelper2.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            if this.AndroidServicesUtil.IsServiceRunning(this.MusicServiceClass.String()) {
            } else {
                this.StartNewSong()
            }
        }
}


func (this *MusicManager) Show() {
        for index := int(this.SongList.Size()); Java_PreDecrement(&index) >= 0;  {
            var sound ISound = any(this.SongList.Get(index)).(ISound); _ = sound
            var duration int64 = int64(sound.GetDuration()); _ = duration
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.PLAY).Append(sound.GetResource()).Append(this.FOR).Appendlong(duration)), this, this.CommonStrings.PROCESS)
        }
}


func (this *MusicManager) StartNewSong() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var resource string = StringUtil_GetInstance().EMPTY_STRING; _ = resource
            if this.CurrentSongSound != nil {
                resource = this.CurrentSongSound.GetResource()
            }
            PreLogUtil_PutOE(fmt.Sprint(this.CommonStrings.EXCEPTION_LABEL) + resource, this, this.CommonStrings.PROCESS, e)
                }
            }()
            this.ResourceUtil.GetContext().StopService(this.CurrentIntent)
            if this.NextSongSound == NoSound_GetInstance() {
                this.CurrentSongSound = any(this.BasicArrayListUtil.GetRandom(this.SongList)).(ISound)
            } else {
                this.CurrentSongSound = this.NextSongSound
                this.NextSongSound = NoSound_GetInstance()
            }
            var duration int64 = int64(this.CurrentSongSound.GetDuration()); _ = duration
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.PLAY).Append(this.CurrentSongSound.GetResource()).Append(this.FOR).Appendlong(duration)), this, this.CommonStrings.PROCESS)
            this.TimeDelayHelper.TimeDelayHelper_().Delay = int(duration)
            this.CurrentIntent.PutExtra(this.MusicStrings.SONG_EXTRA, int(this.ResourceUtil.GetResourceId(this.CurrentSongSound.GetResource())))
            this.CurrentIntent.PutExtra(this.MusicStrings.LEFT_VOLUME, this.LeftVolume)
            this.CurrentIntent.PutExtra(this.MusicStrings.RIGHT_VOLUME, this.RightVolume)
            this.ResourceUtil.GetContext().StartService(this.CurrentIntent)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *MusicManager) Stop() {
        this.ResourceUtil.GetContext().StopService(this.CurrentIntent)
}


func MusicManager_Pause(activity *Activity, musicServiceClass reflect.Type) {
        var commonStateStrings *CommonStateStrings = CommonStateStrings_GetInstance(); _ = commonStateStrings
        var musicPauseIntent *Intent = NewIntent(any(activity).(*Context), musicServiceClass); _ = musicPauseIntent
        musicPauseIntent.PutExtra(commonStateStrings.ON_START_COMMAND, 1)
        activity.StartService(musicPauseIntent)
}


func MusicManager_Resume(activity *Activity, musicServiceClass reflect.Type) {
        if AndroidServicesUtil_GetInstance().IsServiceRunning(musicServiceClass.String()) {
            var commonStateStrings *CommonStateStrings = CommonStateStrings_GetInstance(); _ = commonStateStrings
            var musicResumeIntent *Intent = NewIntent(activity, musicServiceClass); _ = musicResumeIntent
            musicResumeIntent.PutExtra(commonStateStrings.ON_START_COMMAND, 2)
            activity.StartService(musicResumeIntent)
        }
}

