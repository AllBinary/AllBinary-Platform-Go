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

type ImageCompleteUtil struct {
    ISREADY string
    COMPLETE string
}

// The Java class ImageCompleteUtil has subclasses so its variables are this interface
type IImageCompleteUtil interface {
    ImageCompleteUtil_() *ImageCompleteUtil
    WaitForImage(image IImage, name string)
    IsReady(image IImage, name string, timeDelayHelper ITimeDelayHelper) bool
    HandleTimeout(name string)
    WaitForAll()
}

func (this *ImageCompleteUtil) ImageCompleteUtil_() *ImageCompleteUtil {
    return this
}

var ImageCompleteUtil_Instance IImageCompleteUtil = NewImageCompleteUtil()

func (this *ImageCompleteUtil) initFields() {
    this.ISREADY = "isReady"
    this.COMPLETE = "Image loading complete for: "
}

func NewImageCompleteUtil() *ImageCompleteUtil {
    this := &ImageCompleteUtil{}
    this.initFields()
    return this
}


func ImageCompleteUtil_GetInstance() IImageCompleteUtil {
        return ImageCompleteUtil_Instance
}


// throws Exception
func (this *ImageCompleteUtil) WaitForImage(image IImage, name string) {
}


// throws Exception
func (this *ImageCompleteUtil) IsReady(image IImage, name string, timeDelayHelper ITimeDelayHelper) bool {
        var playnImage *PlaynImage = any(image).(*PlaynImage); _ = playnImage
        var playnCoreImage IImage = any(playnImage.GetImage()).(IImage); _ = playnCoreImage
        if !playnCoreImage.IsReady() || playnCoreImage.Width() + playnCoreImage.Height() <= 0 {
            this.HandleTimeout(name)
            return false
        }
        if !image.IsReady() {
            image.Init(image.GetImage())
            if !image.IsReady() {
                PreLogUtil_Put("Waiting for Not Ready Image: " + image.GetName(), this, this.ISREADY)
            }
        }
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        PreLogUtil_Put(this.COMPLETE + image.GetName(), this, commonStrings.SUCCESS)
        return true
}


// throws Exception
func (this *ImageCompleteUtil) HandleTimeout(name string) {
}


// throws Exception
func (this *ImageCompleteUtil) WaitForAll() {
}

