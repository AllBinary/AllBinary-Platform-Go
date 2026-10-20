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
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    BasicArrayListUtil *BasicArrayListUtil
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    PlayerStateUtil *PlayerStateUtil
    TimeDelayHelper ITimeDelayHelper
    PlayerListener PlayerListener
    PLAY string
    SONG string
    NEXT_SONG string
    STOPPING string
    ENDING string
    WAITING_FOR_MEDIA_TO_END string
    ALREADY_PLAYING string
    ALREADY_ENDED string
    SongList IBasicArrayList
    CurrentSongSound ISound
    NextSongSound ISound
    Reset bool
    NoDuration bool
    Stopped bool
}


func (this *MusicManager) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.PlayerStateUtil = PlayerStateUtil_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(0)
    this.PlayerListener = NewMusicManagerPlayerListener()
    this.PLAY = "Play "
    this.SONG = " for: "
    this.NEXT_SONG = fmt.Sprint(CommonLabels_GetInstance().NEXT) + " Song: "
    this.STOPPING = "Stopping Current Song: "
    this.ENDING = "Ending Current Song: "
    this.WAITING_FOR_MEDIA_TO_END = "Waiting for media to end"
    this.ALREADY_PLAYING = "Already Playing: "
    this.ALREADY_ENDED = "Last Song already ended: "
    
    this.CurrentSongSound = NoSound_GetInstance()
    this.NextSongSound = NoSound_GetInstance()
    
    
    this.Stopped = true
}


func NewMusicManager(musicServiceClass reflect.Type, songList IBasicArrayList) *MusicManager {
    this := &MusicManager{}
    this.initFields()
        this.SongList = songList
    return this
}


func (this *MusicManager) StartNewSong() {
        if this.NextSongSound == NoSound_GetInstance() {
            var randomSongSound ISound = any(this.BasicArrayListUtil.GetRandom(this.SongList)).(ISound); _ = randomSongSound
            this.NextSong(randomSongSound, 0, 0)
        }
        this.Process()
}


func (this *MusicManager) NextSong(nextSongSound ISound, leftVolume int, rightVolume int) {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.NEXT_SONG).Append(nextSongSound.GetResource())), this, this.CommonStrings.PROCESS)
        this.NextSongSound = nextSongSound
        this.Reset()
        this.Stopped = false
}


func (this *MusicManager) Reset() {
        this.Reset = true
}


func (this *MusicManager) Process() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var resource string = StringUtil_GetInstance().EMPTY_STRING; _ = resource
            if this.CurrentSongSound != NoSound_GetInstance() {
                resource = this.CurrentSongSound.GetResource()
            }
            PreLogUtil_PutOE(fmt.Sprint(this.CommonStrings.EXCEPTION_LABEL) + resource, this, this.CommonStrings.PROCESS, e)
                }
            }()
            if this.Stopped {
                tryControl1 = 1; return
            }
            if this.SongList.Size() == 0 {
                tryControl1 = 1; return
            }
            if (this.TimeDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) && !this.NoDuration) || this.Reset {
                this.Reset = false
                this.NoDuration = false
                var endingCurrentSongSound ISound = this.CurrentSongSound; _ = endingCurrentSongSound
                var nextSongSound ISound = this.NextSongSound; _ = nextSongSound
                this.NextSongSound = NoSound_GetInstance()
                if nextSongSound == NoSound_GetInstance() {
                    this.CurrentSongSound = any(BasicArrayListUtil_GetInstance().GetRandom(this.SongList)).(ISound)
                } else {
                    this.CurrentSongSound = nextSongSound
                }
                var startingCurrentSongSound ISound = this.CurrentSongSound; _ = startingCurrentSongSound
                var duration int64 = int64(this.CurrentSongSound.GetDuration()); _ = duration
                this.TimeDelayHelper.TimeDelayHelper_().Delay = int(duration)
                if duration <= 0 {
                    var NO_DURATION_FOR string = "No Duration for: "; _ = NO_DURATION_FOR
                    PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(NO_DURATION_FOR).Append(this.CurrentSongSound.GetResource())), this, this.CommonStrings.PROCESS)
                    this.CurrentSongSound.GetPlayerP().AddPlayerListener(this.PlayerListener)
                    this.NoDuration = true
                }
                if endingCurrentSongSound != NoSound_GetInstance() {
                    if endingCurrentSongSound == startingCurrentSongSound && endingCurrentSongSound.GetPlayerP().GetState() == Player_STARTED {
                        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.ALREADY_PLAYING).Append(endingCurrentSongSound.GetResource())), this, this.CommonStrings.PROCESS)
                        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.STOPPING).Append(endingCurrentSongSound.GetResource()).Append(this.SONG).Appendlong(duration)), this, this.CommonStrings.PROCESS)
                        endingCurrentSongSound.GetPlayerP().Stop()
                        this.WaitForStateChange(endingCurrentSongSound, startingCurrentSongSound)
                    } else {
                        if endingCurrentSongSound.GetPlayerP().GetState() == Player_STARTED {
                            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.STOPPING).Append(endingCurrentSongSound.GetResource()).Append(this.SONG).Appendlong(duration)), this, this.CommonStrings.PROCESS)
                            endingCurrentSongSound.GetPlayerP().Stop()
                            this.WaitForStateChange(endingCurrentSongSound, startingCurrentSongSound)
                        } else {
                            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.ALREADY_ENDED).Append(this.PLAY).Append(startingCurrentSongSound.GetResource())), this, this.CommonStrings.PROCESS)
                            if AvianUtil_IsAvian() {
                                endingCurrentSongSound.GetPlayerP().Stop()
                                startingCurrentSongSound.GetPlayerP().Stop()
                            }
                            startingCurrentSongSound.GetPlayerP().Start()
                        }
                    }
                } else {
                    PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.PLAY).Append(this.CurrentSongSound.GetResource()).Append(this.SONG).Appendlong(duration)), this, this.CommonStrings.PROCESS)
                    this.CurrentSongSound.GetPlayerP().Start()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *MusicManager) WaitForStateChange(endingCurrentSongSound ISound, startingCurrentSongSound ISound) {
        for endingCurrentSongSound.GetPlayerP().GetState() == Player_STARTED {
            PreLogUtil_Put(this.WAITING_FOR_MEDIA_TO_END, this, this.CommonStrings.PROCESS)
            time.Sleep(time.Duration(100) * time.Millisecond)
        }
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.PlayerStateUtil.Convert(endingCurrentSongSound.GetPlayerP().GetState())).Append(this.CommonSeps.SPACE).Append(this.PLAY).Append(startingCurrentSongSound.GetResource())), this, this.CommonStrings.PROCESS)
        startingCurrentSongSound.GetPlayerP().Start()
}


// throws Exception
func (this *MusicManager) Stop() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var resource string = StringUtil_GetInstance().EMPTY_STRING; _ = resource
            if currentSongSound != NoSound_GetInstance() {
                resource = currentSongSound.GetResource()
            }
            PreLogUtil_PutOE(fmt.Sprint(this.CommonStrings.EXCEPTION_LABEL) + resource, this, this.CommonStrings.END, e)
                }
            }()
            var currentSongSound ISound = this.CurrentSongSound; _ = currentSongSound
            if currentSongSound != NoSound_GetInstance() {
                this.Stopped = true
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.ENDING).Append(currentSongSound.GetResource())), this, this.CommonStrings.PROCESS)
                currentSongSound.GetPlayerP().Stop()
            }
            this.TimeDelayHelper.SetStartTime(0)
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.END).Append(StringUtil_GetInstance().ToString(currentSongSound))), this, this.CommonStrings.END)
            return
        }(); tryControl1 == 1 {
            return
        }
}

//Handle Media ending for Avian
type MusicManagerPlayerListener struct {
}

// Implements PlayerListener


func NewMusicManagerPlayerListener() *MusicManagerPlayerListener {
    this := &MusicManagerPlayerListener{}
    return this
}
    

    func (this *MusicManager) PlayerUpdate(player *Player, event string, eventData any) {
            PreLogUtil_Put(event, this, this.CommonStrings.PROCESS)
            if event == PlayerListener_END_OF_MEDIA || event == PlayerListener_STOPPED || event == PlayerListener_CLOSED {
                this.Reset()
            }
    }

