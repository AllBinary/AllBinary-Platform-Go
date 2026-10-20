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

type NullClip struct {
}

// Implements *Clip

var NullClip_NULL_CONTROL *NullJ2SEControl = NewNullJ2SEControl(BooleanControl_Type.MUTE)

func NewNullClip() *NullClip {
    this := &NullClip{}
    return this
}


// throws LineUnavailableException
func (this *NullClip) Open_1(format *AudioFormat, data []int8, offset int, bufferSize int) {
}

// Java overloads of NullClip.open
func (this *NullClip) Open(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*AudioFormat](args[0], pass) && Java_IsArg[[]int8](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[int](args[3], pass) {
            this.Open_1(Java_Arg[*AudioFormat](args[0]), Java_Arg[[]int8](args[1]), Java_Arg[int](args[2]), Java_Arg[int](args[3]))
            return
        }
        if len(args) == 1 && Java_IsArg[*AudioInputStream](args[0], pass) {
            this.Open_2(Java_Arg[*AudioInputStream](args[0]))
            return
        }
        if len(args) == 0 {
            this.Open_3()
            return
        }
    }
    panic("No Java overload of NullClip.open for the arguments")
}


// throws LineUnavailableException, IOException
func (this *NullClip) Open_2(stream *AudioInputStream) {
}


func (this *NullClip) GetFrameLength() int {
        return 0
}


func (this *NullClip) GetMicrosecondLength() int64 {
        return 0
}


func (this *NullClip) SetFramePosition(frames int) {
}


func (this *NullClip) SetMicrosecondPosition(microseconds int64) {
}


func (this *NullClip) SetLoopPoints(start int, end int) {
}


func (this *NullClip) Loop(count int) {
}


func (this *NullClip) Drain() {
}


func (this *NullClip) Flush() {
}


func (this *NullClip) Start() {
}


func (this *NullClip) Stop() {
}


func (this *NullClip) IsRunning() bool {
        return false
}


func (this *NullClip) IsActive() bool {
        return false
}


func (this *NullClip) GetFormat() *AudioFormat {
        return NullAudioFormat_NULL_AUDIO_FORMAT
}


func (this *NullClip) GetBufferSize() int {
        return 0
}


func (this *NullClip) Available() int {
        return 0
}


func (this *NullClip) GetFramePosition() int {
        return 0
}


func (this *NullClip) GetLongFramePosition() int64 {
        return 0
}


func (this *NullClip) GetMicrosecondPosition() int64 {
        return 0
}


func (this *NullClip) GetLevel() float32 {
        return 0.0
}


func (this *NullClip) GetLineInfo() *Info {
        return NewInfo(reflect.TypeOf(this), NullAudioFormat_NULL_AUDIO_FORMAT)
}


// throws LineUnavailableException
func (this *NullClip) Open_3() {
}


func (this *NullClip) Close() {
}


func (this *NullClip) IsOpen() bool {
        return false
}


func (this *NullClip) GetControls() []Control {
        return make([]Control, 0)
}


func (this *NullClip) IsControlSupported(control *Type) bool {
        return false
}


func (this *NullClip) GetControl(control *Type) Control {
        return NullClip_NULL_CONTROL
}


func (this *NullClip) AddLineListener(listener *LineListener) {
}


func (this *NullClip) RemoveLineListener(listener *LineListener) {
}

