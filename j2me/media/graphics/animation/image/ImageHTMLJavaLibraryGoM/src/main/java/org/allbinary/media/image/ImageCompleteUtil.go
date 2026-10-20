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
    TimeDelayHelper ITimeDelayHelper
    AllTimeDelayHelper ITimeDelayHelper
}

// The Java class ImageCompleteUtil has subclasses so its variables are this interface
type IImageCompleteUtil interface {
    ImageCompleteUtil_() *ImageCompleteUtil
    WaitForImage(image IImage, name string)
    IsReady(image IImage, name string, timeDelayHelper ITimeDelayHelper) bool
    WaitForATime(image IImage, name string, timeDelayHelper ITimeDelayHelper)
    WaitForAll()
}

func (this *ImageCompleteUtil) ImageCompleteUtil_() *ImageCompleteUtil {
    return this
}

var ImageCompleteUtil_Instance IImageCompleteUtil = NewImageCompleteUtil()

func (this *ImageCompleteUtil) initFields() {
    this.TimeDelayHelper = NewTimeDelayHelper(18000)
    this.AllTimeDelayHelper = NewTimeDelayHelper(120000)
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
        this.TimeDelayHelper.SetStartTime()
        this.WaitForATime(image, name, this.TimeDelayHelper)
}


// throws Exception
func (this *ImageCompleteUtil) IsReady(image IImage, name string, timeDelayHelper ITimeDelayHelper) bool {
        var playnImage *PlaynImage = any(image).(*PlaynImage); _ = playnImage
        if !playnImage.GetImage().IsReady() && playnImage.GetImage().Width() + playnImage.GetImage().Height() == 0 {
            if timeDelayHelper.IsTime() {
                panic(errors.New(fmt.Sprint("isReady: Timeout Waiting or GameHtmlHasLoadedResourcesProcessor does not have this Image: " + name)))
            }
            return false
        }
        return true
}


// throws Exception
func (this *ImageCompleteUtil) WaitForATime(image IImage, name string, timeDelayHelper ITimeDelayHelper) {
        var playnImage *PlaynImage = any(image).(*PlaynImage); _ = playnImage
        for !playnImage.GetImage().IsReady() && playnImage.GetImage().Width() + playnImage.GetImage().Height() == 0 {
            if timeDelayHelper.IsTime() {
                panic(errors.New(fmt.Sprint("waitFor: Timeout Waiting or GameHtmlHasLoadedResourcesProcessor does not have this Image: " + name)))
            }
        }
}


// throws Exception
func (this *ImageCompleteUtil) WaitForAll() {
        this.AllTimeDelayHelper.SetStartTime()
        var hashtable IHashtable = GameFeatureImageCacheFactory_GetInstance().GetHashtable(); _ = hashtable
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var size int = len(objectArray); _ = size
        PreLogUtil_Put("Total: " + fmt.Sprint(size), this, "waitForAll")
        for index := int(0); index < size; index++ {
            this.WaitForATime(any(hashtable.Get(objectArray[index])).(IImage), any(objectArray[index]).(string), this.AllTimeDelayHelper)
        }
}

