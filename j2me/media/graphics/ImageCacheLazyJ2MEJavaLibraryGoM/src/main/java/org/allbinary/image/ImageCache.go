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
    TsUtil *TsUtil
    ConcurrentImageLoadingProcessor IBaseImageLoadingProcessor
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    ResourceUtil *ResourceUtil
    GameGlobalsFactory *GameGlobalsFactory
    GdResources *GDResources
    LoadNowList IBasicArrayList
    LoadSoonList IBasicArrayList
    LoadList IBasicArrayList
    LoadAfterList IBasicArrayList
    Lock *SynchObject
    FirstTime bool
    TotalLoaded int
    ProgressEnded bool
    HasAnyLazyAnimationFactories bool
    Processor IProcessor
    EndProcessor IProcessor
    LOAD_IMAGE_FOR_ANIMATION string
}

// The Java class ImageCache has subclasses so its variables are this interface
type IImageCache interface {
    IImageCacheBase
    ImageCache_() *ImageCache
    FirstProcess()
    AddListener(renderer any)
    WaitForLoadNow()
    LoadImageForAnimation()
    LoadImages()
    LoadImageForAnimations()
    LoadRemainingAnimations()
    LoadImageForLazyAnimation(lazyImageRotationAnimation *LazyImageRotationAnimation) bool
    LoadNextImage()
    LoadImage(image IImage) bool
    Init(image IImage, image2 IImage)
    CreatImage(key string) IImage
    Get(caller string, width int, height int) IImage
    GetWithKey(key any) IImage
    GetIndex(key any) int
    CreateImageFromInputStream(key any, inputStream IInputStream) IImage
    CreateImageLater(key string, width int, height int) IImage
    GetAssociated(lazyImageRotationAnimation *LazyImageRotationAnimation) IBasicArrayList
    Add(lazyImageRotationAnimation *LazyImageRotationAnimation)
    InsertFirst(lazyImageRotationAnimation *LazyImageRotationAnimation)
    ProgressEnded()
    RunTask()
    InitProgress()
    IsLazy() bool
}

func (this *ImageCache) ImageCache_() *ImageCache {
    return this
}

var ImageCache_NULL_IMAGE_CACHE IImageCache = NewImageCache()

func (this *ImageCache) initFields() {
    this.TsUtil = TsUtil_GetInstance()
    this.ConcurrentImageLoadingProcessor = NewConcurrentImageLoadingProcessor(this)
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.ResourceUtil = ResourceUtil_GetInstance()
    this.GameGlobalsFactory = GameGlobalsFactory_GetInstance()
    this.GdResources = GDResources_GetInstance()
    this.LoadNowList = NewBasicArrayListD()
    this.LoadSoonList = NewBasicArrayListD()
    this.LoadList = NewBasicArrayListD()
    this.LoadAfterList = NewBasicArrayListD()
    this.Lock = NewSynchObject()
    this.FirstTime = true
    this.TotalLoaded = 0
    this.ProgressEnded = false
    this.HasAnyLazyAnimationFactories = false
    this.Processor = NewFirstProcessor()
    this.EndProcessor = Processor_GetInstance()
    this.LOAD_IMAGE_FOR_ANIMATION = "Load Image Animation"
}


func NewImageCache() *ImageCache {
    this := &ImageCache{}
    this.ImageCacheBase = *NewImageCacheBase()
    this.initFields()
    return this
}


func (this *ImageCache) FirstProcess() {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var isHTML bool = J2MEUtil_IsHTML(); _ = isHTML
        if isHTML {
            this.Processor = Processor_GetInstance()
            this.EndProcessor = NewHTMLEndProcessor()
        } else {
            this.Processor = NewNotHTMLProcessor()
            this.EndProcessor = NewNotHTMLEndProcessor()
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.PutF(this.CommonStrings.EXCEPTION, this, this.CommonStrings.END_METHOD_NAME)
                    }
                }()
                this.RunTask()
                return
            }(); tryControl1 == 1 {
                return
            }
        }
}


func (this *ImageCache) AddListener(renderer any) {
}


// throws Exception
func (this *ImageCache) WaitForLoadNow() {
        if this.FirstTime {
            var abToGBUtil *ABToGBUtil = ABToGBUtil_GetInstance(); _ = abToGBUtil
            var abCanvas IAllBinaryGameCanvas = any(abToGBUtil.AbCanvas).(IAllBinaryGameCanvas); _ = abCanvas
            for this.LoadNowList.IsEmpty() && (!abCanvas.IsInitialized() || (abCanvas.IsInitialized() && this.HasAnyLazyAnimationFactories)) && !this.ProgressEnded {
                time.Sleep(time.Duration(120) * time.Millisecond)
            }
            this.FirstTime = false
        }
}


// throws Exception
func (this *ImageCache) LoadImageForAnimation() {
        var lazyImageRotationAnimation *LazyImageRotationAnimation = nil; _ = lazyImageRotationAnimation
        { // synchronized (this.Lock)
            if this.LoadNowList.IsEmpty() {
                this.EndProcessor.Process()
                if this.LoadSoonList.IsEmpty() {
                    if this.LoadAfterList.IsEmpty() {
                        if this.FirstTime {
                        } else if this.GameGlobalsFactory.NewCanvas {
                        } else {
                            this.LoadNextImage()
                        }
                        return
                    } else {
                        lazyImageRotationAnimation = any(this.LoadAfterList.Get(0)).(*LazyImageRotationAnimation)
                        if this.LoadImageForLazyAnimation(lazyImageRotationAnimation) {
                            this.LoadAfterList.Remove(lazyImageRotationAnimation)
                        }
                    }
                } else {
                    lazyImageRotationAnimation = any(this.LoadSoonList.Get(0)).(*LazyImageRotationAnimation)
                    if this.LoadImageForLazyAnimation(lazyImageRotationAnimation) {
                        this.LoadSoonList.Remove(lazyImageRotationAnimation)
                    }
                }
                return
            }
            lazyImageRotationAnimation = any(this.LoadNowList.Get(0)).(*LazyImageRotationAnimation)
        }
        if this.LoadImageForLazyAnimation(lazyImageRotationAnimation) {
            { // synchronized (this.Lock)
                this.LoadNowList.Remove(lazyImageRotationAnimation)
            }
            if lazyImageRotationAnimation.LayoutIndex != 0 {
                var list IBasicArrayList = this.GetAssociated(lazyImageRotationAnimation); _ = list
                { // synchronized (this.Lock)
                    var size int = list.Size(); _ = size
                    if size > 0 {
                        this.LoadSoonList.AddAllList(list)
                    }
                }
            }
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            var isHTML bool = J2MEUtil_IsHTML(); _ = isHTML
            if this.LoadNowList.IsEmpty() && (!isHTML || this.FirstTime) {
                progressCanvas.EndFromInitialLazyLoadingComplete()
            } else {
                if this.TotalLoaded % 10 == 0 {
                    progressCanvas.AddNormalPortion(1, this.LOAD_IMAGE_FOR_ANIMATION)
                }
            }
        }
}


// throws Exception
func (this *ImageCache) LoadImages() {
        for !this.LoadList.IsEmpty() || !this.LoadNowList.IsEmpty() {
            this.LoadImageForAnimations()
            this.LoadNextImage()
        }
}


// throws Exception
func (this *ImageCache) LoadImageForAnimations() {
        for !this.LoadNowList.IsEmpty() {
            this.LoadImageForAnimation()
        }
}


// throws Exception
func (this *ImageCache) LoadRemainingAnimations() {
        for !this.LoadAfterList.IsEmpty() || !this.LoadNowList.IsEmpty() {
            for !this.LoadNowList.IsEmpty() {
                this.LoadImageForAnimation()
            }
            var lazyImageRotationAnimation *LazyImageRotationAnimation = nil; _ = lazyImageRotationAnimation
            { // synchronized (this.Lock)
                if !this.LoadAfterList.IsEmpty() {
                    lazyImageRotationAnimation = any(this.LoadAfterList.RemoveAt(0)).(*LazyImageRotationAnimation)
                }
            }
            if lazyImageRotationAnimation != nil {
                this.LoadImageForLazyAnimation(lazyImageRotationAnimation)
            }
        }
}


// throws Exception
func (this *ImageCache) LoadImageForLazyAnimation(lazyImageRotationAnimation *LazyImageRotationAnimation) bool {
        var image IImage = lazyImageRotationAnimation.AnimationInterfaceFactoryInterface.GetImage(); _ = image
        if this.LoadImage(image) {
            lazyImageRotationAnimation.SetRealAnimation()
            return true
        }
        return false
}


// throws Exception
func (this *ImageCache) LoadNextImage() {
        var image IImage = nil; _ = image
        { // synchronized (this.Lock)
            if this.LoadList.Size() == 0 {
                return
            }
            image = any(this.LoadList.RemoveAt(0)).(IImage)
        }
        this.LoadImage(image)
}


// throws Exception
func (this *ImageCache) LoadImage(image IImage) bool {
        if image.IsReady() {
            return true
        } else {
            if image.GetImage() != nil {
                if image.SetReady() {
                    if image.IsReady() {
                        this.TotalLoaded++
                        return true
                    }
                }
            } else {
                var key string = image.GetName(); _ = key
                var image2 IImage = this.CreatImage(key); _ = image2
                if image2.IsReady() {
                    this.Init(image, image2)
                    return true
                } else {
                    image.SetImage(image2.GetImage())
                }
            }
        }
        return false
}


// throws Exception
func (this *ImageCache) Init(image IImage, image2 IImage) {
        image.Init(image2.GetImage())
}


// throws Exception
func (this *ImageCache) CreatImage(key string) IImage {
        var inputStream IInputStream = this.ResourceUtil.GetResourceAsStream(key); _ = inputStream
        var image IImage = ImageFactory_GetInstance().CreateImageFromInputStream(inputStream); _ = image
        image.SetName(key)
        return image
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
            if this.NextIndex > len(this.Widths) - 1 {
                if foundIndex == -1 {
                    foundIndex = this.NextIndex
                    this.Widths[this.NextIndex] = width
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
            var inputStream IInputStream = nil; _ = inputStream
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IImage) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put("Exception: Trying Again After GC", this, this.CommonStrings.GET, e)
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("InputStream: ").Append(StringUtil_GetInstance().ToString(inputStream))), this, this.CommonStrings.GET)
                this.TsUtil.Gc()
                this.TsUtil.Gc()
                this.LogUtil.PutF(Memory_GetInfo(), this, this.CommonStrings.GET)
                time.Sleep(time.Duration(100) * time.Millisecond)
                image = this.CreateImageFromInputStream(key, inputStream)
                    }
                }()
                image = this.CreateImageFromInputStream(key, inputStream)
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            this.Hashtable.Put(key, image)
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
func (this *ImageCache) CreateImageFromInputStream(key any, inputStream IInputStream) IImage {
        var gdLazyResources *GDLazyResources = GDLazyResources_GetInstance(); _ = gdLazyResources
        var resourceStringArray []string = gdLazyResources.RequiredResourcesBeforeLoadingArray; _ = resourceStringArray
        var size int = len(resourceStringArray); _ = size
        for index := int(0); index < size; index++ {
            if key == resourceStringArray[index] {
                return this.CreatImage(any(key).(string))
            }
        }
        this.RunTask()
        var index int = this.GetIndex(key); _ = index
        var width int = gdLazyResources.ImageResourceWidthArray[index]; _ = width
        var height int = gdLazyResources.ImageResourceHeightArray[index]; _ = height
        var image IImage = this.CreateImageLater(any(key).(string), width, height); _ = image
        { // synchronized (this.Lock)
            this.LoadList.Add(image)
        }
        return image
}


// throws Exception
func (this *ImageCache) CreateImageLater(key string, width int, height int) IImage {
        return Image_CreateImageLater(key, width, height)
}


func (this *ImageCache) GetAssociated(lazyImageRotationAnimation *LazyImageRotationAnimation) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        { // synchronized (this.Lock)
            var lazyImageRotationAnimation2 *LazyImageRotationAnimation = nil; _ = lazyImageRotationAnimation2
            var size int = this.LoadAfterList.Size(); _ = size
            for index := int(0); index < size; index++ {
                lazyImageRotationAnimation2 = any(this.LoadAfterList.Get(index)).(*LazyImageRotationAnimation)
                if lazyImageRotationAnimation2.InstanceId == lazyImageRotationAnimation.InstanceId {
                    list.Add(lazyImageRotationAnimation2)
                }
            }
            var size2 int = list.Size(); _ = size2
            for index := int(0); index < size2; index++ {
                this.LoadAfterList.Remove(list.Get(index))
            }
        }
        return list
}


func (this *ImageCache) Add(lazyImageRotationAnimation *LazyImageRotationAnimation) {
        { // synchronized (this.Lock)
            this.LoadAfterList.Add(lazyImageRotationAnimation)
        }
}


// throws Exception
func (this *ImageCache) InsertFirst(lazyImageRotationAnimation *LazyImageRotationAnimation) {
        if this.LoadNowList.Contains(lazyImageRotationAnimation) {
        } else {
            { // synchronized (this.Lock)
                this.LoadNowList.Add(lazyImageRotationAnimation)
                this.LoadAfterList.Remove(lazyImageRotationAnimation)
            }
        }
        this.RunTask()
}


func (this *ImageCache) ProgressEnded() {
        this.ProgressEnded = true
}


// throws Exception
func (this *ImageCache) RunTask() {
        this.Processor.Process()
}


func (this *ImageCache) InitProgress() {
        if this.FirstTime {
            this.FirstTime = false
        }
}


func (this *ImageCache) IsLazy() bool {
        return true
}

type NotHTMLProcessor struct {
    Processor
}


func NewNotHTMLProcessor() *NotHTMLProcessor {
    this := &NotHTMLProcessor{}
    this.Processor = *NewProcessor()
    return this
}
    

    func (this *ImageCache) Process() {
            this.ConcurrentImageLoadingProcessor.RunTask()
    }

type NotHTMLEndProcessor struct {
    Processor
}


func NewNotHTMLEndProcessor() *NotHTMLEndProcessor {
    this := &NotHTMLEndProcessor{}
    this.Processor = *NewProcessor()
    return this
}
    

    func (this *ImageCache) Process() {
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            progressCanvas.EndIfPaintedSinceStart()
    }

type HTMLEndProcessor struct {
    Processor
}


func NewHTMLEndProcessor() *HTMLEndProcessor {
    this := &HTMLEndProcessor{}
    this.Processor = *NewProcessor()
    return this
}
    

    func (this *ImageCache) Process() {
            var size int = this.GdResources.CurrentLayoutRequiredTotal; _ = size
            if size == 0 {
                var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
                progressCanvas.EndIfPaintedSinceStart()
            } else if this.TotalLoaded > size / 12 {
                var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
                progressCanvas.EndIfPaintedSinceStart()
                this.EndProcessor = NewNotHTMLEndProcessor()
            }
    }

type FirstProcessor struct {
    Processor
}


func NewFirstProcessor() *FirstProcessor {
    this := &FirstProcessor{}
    this.Processor = *NewProcessor()
    return this
}
    

    func (this *ImageCache) Process() {
            this.FirstProcess()
    }

