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

type ImageCache struct {
    ImageCacheBase
    CommonStrings *CommonStrings
    TsUtil *TsUtil
}

// The Java class ImageCache has subclasses so its variables are this interface
type IImageCache interface {
    IImageCacheBase
    ImageCache_() *ImageCache
    AddListener(renderer any)
    GetIndex(key any) int
    Get(caller string, width int, height int) IImage
    GetWithKey(key any) IImage
    LoadImageForAnimation()
    ProgressEnded()
    RunTask()
    IsLazy() bool
}

func (this *ImageCache) ImageCache_() *ImageCache {
    return this
}

var ImageCache_NULL_IMAGE_CACHE IImageCache = NewImageCache()

func (this *ImageCache) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
}


func NewImageCache() *ImageCache {
    this := &ImageCache{}
    this.ImageCacheBase = *NewImageCacheBase()
    this.initFields()
    return this
}


func (this *ImageCache) AddListener(renderer any) {
}


func (this *ImageCache) GetIndex(key any) int {
        var gdResources *GDResources = GDResources_GetInstance(); _ = gdResources
        var resourceStringArray []string = gdResources.ResourceStringArray; _ = resourceStringArray
        var size int = len(resourceStringArray); _ = size
        for index := int(0); index < size; index++ {
            if resourceStringArray[index] == key {
                return index
            }
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("unable to find key: ").Append(StringUtil_GetInstance().ToString(key))), this, this.CommonStrings.RUN)
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *ImageCache) Get(caller string, width int, height int) IImage {
        var foundIndex int = this.GetIndexWH(width, height); _ = foundIndex
        var image IImage = this.GetFromAvailable(foundIndex, width, height); _ = image
        if image == NullImage_NULL_IMAGE {
            this.Volume += width * height
            if this.Volume > 32000 {
                this.TsUtil.Gc()
                this.Volume = 0
            }
            image = this.CreateImage(caller, width, height)
            var widths []int = this.Widths; _ = widths
            if this.NextIndex > len(widths) - 1 {
                if foundIndex == -1 {
                    foundIndex = this.NextIndex
                    widths[this.NextIndex] = width
                    this.Heights[this.NextIndex] = height
                    this.NextIndex++
                }
                this.ListOfList[foundIndex].Add(image)
            }
        }
        return image
}


// throws Exception
func (this *ImageCache) GetWithKey(key any) IImage {
        var image IImage = this.GetImage(key); _ = image
        if image == NullImage_NULL_IMAGE {
            var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
            var resourceInputStream IInputStream = resourceUtil.GetResourceAsStream(any(key).(string)); _ = resourceInputStream
            if resourceInputStream == nil {
                panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Image resource is not available for key: ").Append(StringUtil_GetInstance().ToString(key))))))
            }
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IImage) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put("Exception: Trying Again After GC", this, this.CommonStrings.GET, e)
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("InputStream: ").Append(StringUtil_GetInstance().ToString(resourceInputStream))), this, this.CommonStrings.GET)
                this.TsUtil.Gc()
                this.TsUtil.Gc()
                this.LogUtil.PutF(Memory_GetInfo(), this, this.CommonStrings.GET)
                time.Sleep(time.Duration(100) * time.Millisecond)
                image = this.CreateImageFromInputStream(key, resourceInputStream)
                    }
                }()
                image = this.CreateImageFromInputStream(key, resourceInputStream)
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            if J2MEUtil_IsHTML() {
            } else {
                resourceInputStream.Close()
            }
            this.Hashtable.Put(key, image)
        }
        return image
}


// throws Exception
func (this *ImageCache) LoadImageForAnimation() {
}


func (this *ImageCache) ProgressEnded() {
}


func (this *ImageCache) RunTask() {
}


func (this *ImageCache) IsLazy() bool {
        return false
}

