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

type ImageArrayAnimation struct {
    IndexedAnimation
    ImageArray []IImage
    TotalFrames int
    CircularIndexUtil *CircularIndexUtil
    Anchor int
}

// The Java class ImageArrayAnimation has subclasses so its variables are this interface
type IImageArrayAnimation interface {
    IIndexedAnimation
    ImageArrayAnimation_() *ImageArrayAnimation
    GetAnimationSize() int
    NextFrame()
    PreviousFrame()
    SetFrame(index int)
    GetFrame() int
    GetSize() int
    SetSequence(sequence []int)
    GetSequence() []int
    GetImage(index int) IImage
    GetImageArray() []IImage
    SetImageArray(imageArray []IImage)
    PaintXY(graphics IGraphics, x int, y int)
}

func (this *ImageArrayAnimation) ImageArrayAnimation_() *ImageArrayAnimation {
    return this
}


func (this *ImageArrayAnimation) initFields() {
    this.ImageArray = NullImage_NULL_IMAGE_ARRAY
    
    this.CircularIndexUtil = CircularIndexUtil_NULL_CIRCULAR_INDEX_UTIL
    this.Anchor = Anchor_TOP_LEFT
}


func NewImageArrayAnimation(imageArray []IImage, animationBehavior IAnimationBehavior) *ImageArrayAnimation {
    this := &ImageArrayAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.SetImageArray(imageArray)
    return this
}


// throws Exception
func (this *ImageArrayAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *ImageArrayAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
}


func (this *ImageArrayAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
}


func (this *ImageArrayAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
}


func (this *ImageArrayAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *ImageArrayAnimation) GetSize() int {
        return this.TotalFrames
}


func (this *ImageArrayAnimation) SetSequence(sequence []int) {
}


func (this *ImageArrayAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *ImageArrayAnimation) GetImage(index int) IImage {
        return this.ImageArray[index]
}


func (this *ImageArrayAnimation) GetImageArray() []IImage {
        return this.ImageArray
}


func (this *ImageArrayAnimation) SetImageArray(imageArray []IImage) {
        this.ImageArray = imageArray
        this.TotalFrames = len(imageArray)
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(this.TotalFrames)
}


func (this *ImageArrayAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.ImageArray[this.CircularIndexUtil.GetIndex()], x, y, this.Anchor)
}

