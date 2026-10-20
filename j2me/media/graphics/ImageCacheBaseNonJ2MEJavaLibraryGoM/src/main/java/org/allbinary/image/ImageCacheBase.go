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

type ImageCacheBase struct {
    LogUtil *LogUtil
    ImageFactory *ImageFactory
    Hashtable *ABHashtable[any, any]
    SIZE int
    Widths []int
    Heights []int
    ListOfList []IBasicArrayList
    AvailableListOfList []IBasicArrayList
    Volume int
    NextIndex int
}

// The Java class ImageCacheBase has subclasses so its variables are this interface
type IImageCacheBase interface {
    ImageCacheBase_() *ImageCacheBase
    GetImage(resourceId any) IImage
    ReleaseAll()
    GetIndexWH(width int, height int) int
    GetFromAvailable(foundIndex int, width int, height int) IImage
    Get(caller string, width int, height int) IImage
    GetWithKey(key any) IImage
    CreateImage(caller string, width int, height int) IImage
    CreateImageFromInputStream(key any, inputStream IInputStream) IImage
    ToString() string
    GetHashtableP() *ABHashtable[any, any]
    Init(image IImage)
    InitProgress()
}

func (this *ImageCacheBase) ImageCacheBase_() *ImageCacheBase {
    return this
}


func (this *ImageCacheBase) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageFactory = ImageFactory_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    this.SIZE = 128
    this.Widths = make([]int, this.SIZE)
    this.Heights = make([]int, this.SIZE)
    this.ListOfList = make([]IBasicArrayList, this.SIZE)
    this.AvailableListOfList = make([]IBasicArrayList, this.SIZE)
    this.Volume = 0
    this.NextIndex = 0
}


func NewImageCacheBase() *ImageCacheBase {
    this := &ImageCacheBase{}
    this.initFields()
        for index := int(len(this.ListOfList) - 1); index >= 0; index-- {
            this.ListOfList[index] = NewBasicArrayListD()
        }
        for index := int(len(this.AvailableListOfList) - 1); index >= 0; index-- {
            this.AvailableListOfList[index] = NewBasicArrayListD()
        }
    return this
}


func (this *ImageCacheBase) GetImage(resourceId any) IImage {
        var imageCanBeNull any = this.Hashtable.Get(resourceId); _ = imageCanBeNull
        if imageCanBeNull == nil {
            return NullImage_NULL_IMAGE
        }
        return any(imageCanBeNull).(IImage)
}


func (this *ImageCacheBase) ReleaseAll() {
        for index := int(len(this.ListOfList) - 1); index >= 0; index-- {
            this.AvailableListOfList[index].Clear()
            this.AvailableListOfList[index].AddAllList(this.ListOfList[index])
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("ImageCache: ").Append(this.ToString())), this, "releaseAll")
}


func (this *ImageCacheBase) GetIndexWH(width int, height int) int {
        var foundIndex int = -1; _ = foundIndex
        var size int = len(this.Widths); _ = size
        for index := int(0); index < size; index++ {
            if this.Widths[index] == width && this.Heights[index] == height {
                foundIndex = index
                return foundIndex
            }
        }
        return foundIndex
}


func (this *ImageCacheBase) GetFromAvailable(foundIndex int, width int, height int) IImage {
        if foundIndex != -1 {
            if this.AvailableListOfList[foundIndex].Size() > 0 {
                var list IBasicArrayList = this.AvailableListOfList[foundIndex]; _ = list
                return any(list.RemoveAt(list.Size() - 1)).(IImage)
            }
        }
        return NullImage_NULL_IMAGE
}


// throws Exception
func (this *ImageCacheBase) Get(caller string, width int, height int) IImage {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *ImageCacheBase) GetWithKey(key any) IImage {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *ImageCacheBase) CreateImage(caller string, width int, height int) IImage {
        return this.ImageFactory.CreateImage(caller, width, height)
}


// throws Exception
func (this *ImageCacheBase) CreateImageFromInputStream(key any, inputStream IInputStream) IImage {
        return this.ImageFactory.CreateImageFromInputStream(key, inputStream)
}


func (this *ImageCacheBase) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(this.NextIndex - 1); index >= 0; index-- {
            var width int = this.Widths[index]; _ = width
            var height int = this.Heights[index]; _ = height
            var total int = this.ListOfList[index].Size(); _ = total
            var totalAvailable int = this.AvailableListOfList[index].Size(); _ = totalAvailable
            stringBuffer.Append(" w: ")
            stringBuffer.Appendint(width)
            stringBuffer.Append(" h: ")
            stringBuffer.Appendint(height)
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            stringBuffer.Append(CommonLabels_GetInstance().TOTAL_LABEL)
            stringBuffer.Appendint(total)
            stringBuffer.Append(" available: ")
            stringBuffer.Appendint(totalAvailable)
        }
        return fmt.Sprint(stringBuffer)
}

func (this *ImageCacheBase) String() string {
    return this.ToString()
}


func (this *ImageCacheBase) GetHashtableP() *ABHashtable[any, any] {
        return this.Hashtable
}


func (this *ImageCacheBase) Init(image IImage) {
}


func (this *ImageCacheBase) InitProgress() {
}

