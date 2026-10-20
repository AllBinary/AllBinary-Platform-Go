/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

//OpenGLNone
type SWTJOGLProcessor struct {
}

// Implements Runnable

var SWTJOGLProcessor_Instance *SWTJOGLProcessor = NewSWTJOGLProcessor()

func NewSWTJOGLProcessor() *SWTJOGLProcessor {
    this := &SWTJOGLProcessor{}
    return this
}


func SWTJOGLProcessor_GetInstance() *SWTJOGLProcessor {
        return SWTJOGLProcessor_Instance
}


func (this *SWTJOGLProcessor) Init(display any, comp any) {
}


func (this *SWTJOGLProcessor) IsHolderCreated() bool {
        return true
}


func (this *SWTJOGLProcessor) SetRenderer(renderer any) {
}


func (this *SWTJOGLProcessor) IsJOGL() bool {
        return false
}


func (this *SWTJOGLProcessor) CreateSurface() {
}


func (this *SWTJOGLProcessor) OnSurfaceCreated() {
}


func (this *SWTJOGLProcessor) OnSurfaceChangedForEvent(event any) {
}


func (this *SWTJOGLProcessor) OnSurfaceChanged() {
}


func (this *SWTJOGLProcessor) OnSurfaceChangedWH(width int, height int) {
}


func (this *SWTJOGLProcessor) AddListener() {
}


func (this *SWTJOGLProcessor) AddListeners(midlet IMIDlet) {
}


func (this *SWTJOGLProcessor) Run() {
}


func (this *SWTJOGLProcessor) Clear() {
}


func (this *SWTJOGLProcessor) SetCustom(width int, height int, ratio float32) {
}


func (this *SWTJOGLProcessor) SetRatioProcessor(ratioProcessor IRatioProcessor) {
}


func (this *SWTJOGLProcessor) IsPortraitRatioProcessor() bool {
        return true
}


func (this *SWTJOGLProcessor) CharWidth(font Font, ch rune) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *SWTJOGLProcessor) CharsWidth(font Font, ch []rune, offset int, length int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *SWTJOGLProcessor) StringWidth(font Font, str string) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *SWTJOGLProcessor) SubstringWidth(font Font, str string, offset int, len_ int) int {
        panic(errors.New(fmt.Sprint()))
}

