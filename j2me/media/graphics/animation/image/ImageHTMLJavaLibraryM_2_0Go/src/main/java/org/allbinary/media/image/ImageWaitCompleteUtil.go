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

type ImageWaitCompleteUtil struct {
    ImageCompleteUtil
    LogUtil *LogUtil
    TimeDelayHelper ITimeDelayHelper
    AllTimeDelayHelper ITimeDelayHelper
}


func (this *ImageWaitCompleteUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(18000)
    this.AllTimeDelayHelper = NewTimeDelayHelper(120000)
}


func NewImageWaitCompleteUtil() *ImageWaitCompleteUtil {
    this := &ImageWaitCompleteUtil{}
    this.ImageCompleteUtil = *NewImageCompleteUtil()
    this.initFields()
    return this
}


// throws Exception
func (this *ImageWaitCompleteUtil) WaitForImage(image IImage, name string) {
        this.TimeDelayHelper.SetStartTimeTNT()
        this.WaitForATime(image, name, this.TimeDelayHelper)
}


// throws Exception
func (this *ImageWaitCompleteUtil) HandleTimeout(name string) {
        if this.TimeDelayHelper.IsTimeTNT() {
            panic(errors.New(fmt.Sprint("isReady: Timeout Waiting or GameHtmlHasLoadedResourcesProcessor does not have this Image: " + name)))
        }
}


// throws Exception
func (this *ImageWaitCompleteUtil) WaitForATime(image IImage, name string, timeDelayHelper ITimeDelayHelper) {
        var playnImage *PlaynImage = any(image).(*PlaynImage); _ = playnImage
        var playnCoreImage *CoreImage = any(playnImage.GetImage()).(*CoreImage); _ = playnCoreImage
        for !playnCoreImage.IsReady() || playnCoreImage.Width() + playnCoreImage.Height() <= 0 {
            if timeDelayHelper.IsTimeTNT() {
                panic(errors.New(fmt.Sprint("waitFor: Timeout Waiting or GameHtmlHasLoadedResourcesProcessor does not have this Image: " + name)))
            }
        }
        if !image.IsReady() {
            image.Init(image.GetImage())
        }
}


// throws Exception
func (this *ImageWaitCompleteUtil) WaitForAll() {
        this.AllTimeDelayHelper.SetStartTimeTNT()
        var hashtable *ABHashtable[any, any] = GameFeatureImageCacheFactory_GetInstance().GetHashtableP(); _ = hashtable
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var size int = len(objectArray); _ = size
        PreLogUtil_Put("Image Total: " + fmt.Sprint(size), this, "waitForAll")
        for index := int(0); index < size; index++ {
            this.WaitForATime(any(hashtable.Get(objectArray[index])).(IImage), any(objectArray[index]).(string), this.AllTimeDelayHelper)
        }
}

