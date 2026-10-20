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

//ImageCacheOpenGLESJavaLibrary
type OpenGLImageCache struct {
    ImageCache
    ImageCache IImageCache
    PreResourceImageUtil *PreResourceImageUtil
    Gl GL10
    Lock *SynchObject
    List IBasicArrayList
    Renderer IAllBinaryRendererBase3
}


func (this *OpenGLImageCache) initFields() {
    this.ImageCache = ImageCacheFactory_GetInstance()
    this.PreResourceImageUtil = PreResourceImageUtil_GetInstance()
    this.Gl = NullGL10_NULL_GL10
    this.Lock = NewSynchObject()
    this.List = NewBasicArrayListD()
    this.Renderer = NewAllBinaryRendererBase3()
}


func NewOpenGLImageCache() *OpenGLImageCache {
    this := &OpenGLImageCache{}
    this.ImageCache = *NewImageCache()
    this.initFields()
    return this
}


func (this *OpenGLImageCache) AddListener(renderer any) {
        this.Renderer = any(renderer).(IAllBinaryRendererBase3)
}


// throws Exception
func (this *OpenGLImageCache) Update(gl GL10) {
        this.Gl = gl
        { // synchronized (this.Lock)
            for index := int(this.List.Size() - 1); index >= 0; index-- {
                var openGLESImage IOpenGLESImage = (any(this.List.BasicArrayList_().ObjectArray[index]).(IOpenGLESImage)); _ = openGLESImage
                if openGLESImage != OpenGLESImage_NULL_OPENGL_IMAGE {
                    openGLESImage.Set(gl)
                }
            }
        }
}


// throws Exception
func (this *OpenGLImageCache) CreateImage(caller string, width int, height int) IImage {
        var textureSize int = width; _ = textureSize
        if height > width {
            textureSize = height
        }
        for (textureSize % 4) != 0 {
            textureSize++
        }
        width = textureSize
        height = textureSize
        var image2 IImage = this.ImageCache.Get(caller, width, height); _ = image2
        var image IImage = this.PreResourceImageUtil.Encapsulate(image2); _ = image
        { // synchronized (this.Lock)
            if image != NullImage_NULL_IMAGE {
                this.List.Add(image)
            }
        }
        return image
}


// throws Exception
func (this *OpenGLImageCache) CreateImageFromInputStream(key any, inputStream IInputStream) IImage {
        var cachedImage IImage = this.ImageCache.GetWithKey(key); _ = cachedImage
        var image IImage = this.PreResourceImageUtil.Encapsulate(cachedImage); _ = image
        { // synchronized (this.Lock)
            if image != NullImage_NULL_IMAGE {
                this.List.Add(image)
            }
        }
        return image
}


func (this *OpenGLImageCache) GetGlP() GL10 {
        return this.Gl
}


func (this *OpenGLImageCache) Init(image IImage) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            { // synchronized (this.Lock)
                if this.List.Contains(image) {
                    panic(errors.New(fmt.Sprint()))
                }
                this.List.Add(image)
            }
            this.Renderer.Add(image)
            return
        }(); tryControl1 == 1 {
            return
        }
}

