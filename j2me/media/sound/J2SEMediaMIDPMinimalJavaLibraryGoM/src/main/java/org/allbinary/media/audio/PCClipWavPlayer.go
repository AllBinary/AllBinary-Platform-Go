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

type PCClipWavPlayer struct {
    BasicPlayer
    LogUtil *LogUtil
    AudioInputStream *AudioInputStream
    Clip *Clip
}

// Implements *LineListener


func (this *PCClipWavPlayer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewPCClipWavPlayer(inputStream IInputStream) *PCClipWavPlayer {
    this := &PCClipWavPlayer{}
    this.BasicPlayer = *NewBasicPlayer()
    this.initFields()
        var audioInputStream *AudioInputStream = NewAudioInputStream(NewByteArrayInputStream(NullUtil_GetInstance().NULL_BYTE_ARRAY), NullAudioFormat_NULL_AUDIO_FORMAT, 0); _ = audioInputStream
        var clip *Clip = NewNullClip(); _ = clip
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            audioInputStream = AudioSystem_GetAudioInputStream(inputStream)
            clip = this.Create(audioInputStream)
            if clip == nil {
                this.LogUtil.Put("Clip was null", this, this.CommonStrings.CONSTRUCTOR, errors.New(fmt.Sprint()))
                panic(errors.New(fmt.Sprint()))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        this.AudioInputStream = audioInputStream
        this.Clip = clip
    return this
}


func (this *PCClipWavPlayer) Close() {
        MusicThreadPool_GetInstance().RunTask(/* Java anonymous class body: implement it as a named type
            

        func (this *PCClipWavPlayer) Run() {
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                        }
                    }()
                    this.Close2()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}


// throws Exception
func (this *PCClipWavPlayer) Close2() {
        this.Clip.Drain()
        this.Clip.Flush()
        this.Clip.Close()
}


func (this *PCClipWavPlayer) GetContentType() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


// throws MediaException
// synchronized
func (this *PCClipWavPlayer) Start() {
        MusicThreadPool_GetInstance().RunTask(/* Java anonymous class body: implement it as a named type
            

        func (this *PCClipWavPlayer) Run() {
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                        }
                    }()
                    this.Start2()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}


// throws Exception
func (this *PCClipWavPlayer) Start2() {
        this.Clip.SetFramePosition(0)
        this.Clip.Loop(this.GetLoopCount())
        this.Clip.Start()
        this.BasicPlayer.Start()
}


// throws Exception
func (this *PCClipWavPlayer) Create(audioInputStream *AudioInputStream) *Clip {
        var clip *Clip = AudioSystem_GetClip(); _ = clip
        clip.AddLineListener(this)
        clip.Open(audioInputStream)
        return clip
}


// throws MediaException
// synchronized
func (this *PCClipWavPlayer) Stop() {
        MusicThreadPool_GetInstance().RunTask(/* Java anonymous class body: implement it as a named type
            

        func (this *PCClipWavPlayer) Run() {
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                        }
                    }()
                    this.Stop2()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}


// throws Exception
func (this *PCClipWavPlayer) Stop2() {
        this.Clip.Stop()
        this.BasicPlayer.Stop()
}


func (this *PCClipWavPlayer) GetControl(controlType string) Control {
        return NewNullControl()
}


func (this *PCClipWavPlayer) GetControls() []Control {
        return make([]Control, 0)
}


func (this *PCClipWavPlayer) SetVolume(leftVolume int, rightVolume int) {
        this.SetVolumeF((float32(leftVolume)) / 100.0)
}


func (this *PCClipWavPlayer) GetVolume() float32 {
        var masterGainFloatControl *FloatControl = any(this.Clip.GetControl(FloatControl_Type.MASTER_GAIN)).(*FloatControl); _ = masterGainFloatControl
        return float32(math.Pow(10, masterGainFloatControl.GetValue() / 20))
}


func (this *PCClipWavPlayer) SetVolumeF(volume float32) {
        if volume < 0 || volume > 1 {
            panic(errors.New(fmt.Sprint("Volume: " + fmt.Sprint(volume))))
        }
        var masterGainFloatControl *FloatControl = any(this.Clip.GetControl(FloatControl_Type.MASTER_GAIN)).(*FloatControl); _ = masterGainFloatControl
        masterGainFloatControl.SetValue(20 * float32(math.Log10(volume)))
}


func (this *PCClipWavPlayer) GetDuration() int64 {
        return this.Clip.GetMicrosecondLength() / 1000
}


func (this *PCClipWavPlayer) Update(event *LineEvent) {
        if Java_Equals(event.GetType(), LineEvent_Type.STOP) {
            var size int = this.ListenersList.Size(); _ = size
            for index := int(0); index < size; index++ {
                var listener PlayerListener = any(this.ListenersList.Get(size)).(PlayerListener); _ = listener
                listener.PlayerUpdate(this, PlayerListener_END_OF_MEDIA, nil)
            }
        }
}

