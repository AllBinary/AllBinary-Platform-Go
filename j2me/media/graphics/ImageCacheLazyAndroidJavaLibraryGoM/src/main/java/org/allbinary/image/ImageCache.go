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
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SystemWrapper *ABSystemWrapper
    TsUtil *TsUtil
}

// The Java class ImageCache has subclasses so its variables are this interface
type IImageCache interface {
    IImageCacheBase
    ImageCache_() *ImageCache
    Get(args ...any) IImage
    GetIndex(key any) int
    CreateImage(key any, inputStream IInputStream) IImage
    IsLazy() bool
}

func (this *ImageCache) ImageCache_() *ImageCache {
    return this
}

var ImageCache_NULL_IMAGE_CACHE IImageCache = NewImageCache()

func (this *ImageCache) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SystemWrapper = ABSystemWrapper_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
}


func NewImageCache() *ImageCache {
    this := &ImageCache{}
    this.ImageCacheBase = *NewImageCacheBase()
    this.initFields()
    return this
}


// throws Exception
func (this *ImageCache) Get_1(caller string, width int, height int) IImage {
        var foundIndex int = this.GetIndexWH(width, height); _ = foundIndex
        var image IImage = this.GetFromAvailable(foundIndex, width, height); _ = image
        if image == NullImage_NULL_IMAGE {
            this.Volume += width * height
            if this.Volume > 32000 {
                this.TsUtil.Gc()
                this.Volume = 0
            }
            image = this.CreateImage(caller, width, height)
            if foundIndex == -1 {
                foundIndex = this.NextIndex
                this.Widths[this.NextIndex] = width
                this.Heights[this.NextIndex] = height
                this.NextIndex++
            }
            this.ListOfList[foundIndex].Add(image)
        }
        return image
}

// Java overloads of ImageCache.get
func (this *ImageCache) Get(args ...any) IImage {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.Get_1(Java_Arg[string](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]))
        }
        if len(args) == 1 && Java_IsArg[any](args[0], pass) {
            return this.Get_2(Java_Arg[any](args[0]))
        }
    }
    panic("No Java overload of ImageCache.get for the arguments")
}


// throws Exception
func (this *ImageCache) Get_2(key any) IImage {
        var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
        var resourceId int = resourceUtil.GetResourceId(any(key).(string)); _ = resourceId
        var image IImage = this.GetImage(resourceId); _ = image
        if image == NullImage_NULL_IMAGE {
            var inputStream IInputStream = nil; _ = inputStream
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IImage) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put("Exception: Trying Again After GC", this, this.CommonStrings.GET, e)
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("InputStream: ").Append(fmt.Sprint(inputStream))), this, this.CommonStrings.GET)
                this.TsUtil.Gc()
                this.TsUtil.Gc()
                this.LogUtil.PutF(Memory_GetInfo(), this, this.CommonStrings.GET)
                time.Sleep(time.Duration(100) * time.Millisecond)
                image = this.CreateImage(key, inputStream)
                    }
                }()
                this.LogUtil.PutF(Memory_GetInfo(), this, this.CommonStrings.GET)
                image = this.CreateImage(key, inputStream)
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            this.Hashtable.Put(resourceId, image)
        }
        return image
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
func (this *ImageCache) CreateImage(key any, inputStream IInputStream) IImage {
        return Image_CreateImageLater(any(key).(string), -1, -1)
}


func (this *ImageCache) IsLazy() bool {
        panic(errors.New(fmt.Sprint()))
}

