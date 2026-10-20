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

type MotionRectanglesImageInputOutput struct {
    LogUtil *LogUtil
    NUMBER_OF_IMAGES int
}

// Implements ImageIOInterface


func (this *MotionRectanglesImageInputOutput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NUMBER_OF_IMAGES = 1
}


func NewMotionRectanglesImageInputOutput() *MotionRectanglesImageInputOutput {
    this := &MotionRectanglesImageInputOutput{}
    this.initFields()
    return this
}


// throws Exception
func (this *MotionRectanglesImageInputOutput) Save_1(frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var motionRectanglesResultsFrameCacheable *MotionRectanglesResultsFrameCacheable = any((any(AllMotionRectanglesResultsCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(frame)).(*MotionRectanglesResultsFrameCacheable); _ = motionRectanglesResultsFrameCacheable
        if motionRectanglesResultsFrameCacheable != nil {
            var motionRectangles IMotionRectangles = motionRectanglesResultsFrameCacheable.GetMotionRectangles(); _ = motionRectangles
            this.Save(motionRectangles, frame)
            motionRectanglesResultsFrameCacheable = any((any(ConsolidatedMotionRectanglesResultsCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(frame)).(*MotionRectanglesResultsFrameCacheable)
            motionRectangles = motionRectanglesResultsFrameCacheable.GetMotionRectangles()
            this.Save(motionRectangles, frame)
            motionRectanglesResultsFrameCacheable = any((any(ConstrainedMotionRectanglesResultsCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(frame)).(*MotionRectanglesResultsFrameCacheable)
            motionRectangles = motionRectanglesResultsFrameCacheable.GetMotionRectangles()
            this.Save(motionRectangles, frame)
        } else {
            this.LogUtil.PutF("No Motion Rectangle Results for Frame: " + fmt.Sprint(frame), this, commonStrings.SAVE)
        }
}

// Java overloads of MotionRectanglesImageInputOutput.save
func (this *MotionRectanglesImageInputOutput) Save(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            this.Save_1(Java_Arg[int64](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[IMotionRectangles](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            this.Save_2(Java_Arg[IMotionRectangles](args[0]), Java_Arg[int64](args[1]))
            return
        }
    }
    panic("No Java overload of MotionRectanglesImageInputOutput.save for the arguments")
}


// throws Exception
func (this *MotionRectanglesImageInputOutput) Save_2(motionRectangles IMotionRectangles, frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var filePathStringBuffer *StringMaker = NewStringMaker(); _ = filePathStringBuffer
        filePathStringBuffer.Append(ImageOutputData_SAVE_PATH)
        filePathStringBuffer.Append(LongUtil_FillIn(fmt.Sprint(frame)))
        filePathStringBuffer.Append("_")
        filePathStringBuffer.Append(motionRectangles.GetName())
        var filePathStringBuffer1 *StringMaker = NewStringMaker(); _ = filePathStringBuffer1
        filePathStringBuffer1.Append(fmt.Sprint(filePathStringBuffer))
        filePathStringBuffer1.Append("_1")
        filePathStringBuffer1.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        var filePathStringBuffer2 *StringMaker = NewStringMaker(); _ = filePathStringBuffer2
        filePathStringBuffer2.Append(fmt.Sprint(filePathStringBuffer))
        filePathStringBuffer2.Append("_2")
        filePathStringBuffer2.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        var filePathStringBuffer3 *StringMaker = NewStringMaker(); _ = filePathStringBuffer3
        filePathStringBuffer3.Append(fmt.Sprint(filePathStringBuffer))
        filePathStringBuffer3.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        this.LogUtil.PutF("Motion Image File Path 1: " + fmt.Sprint(filePathStringBuffer1), this, commonStrings.SAVE)
        this.LogUtil.PutF("Motion Image File Path 2: " + fmt.Sprint(filePathStringBuffer2), this, commonStrings.SAVE)
        this.LogUtil.PutF("Motion Image File Path 2: " + fmt.Sprint(filePathStringBuffer3), this, commonStrings.SAVE)
        var bufferedImageCacheables []*BufferedImageCacheable = make([]*BufferedImageCacheable, this.NUMBER_OF_IMAGES); _ = bufferedImageCacheables
        var imageComparisonInfo *ImageComparisonResult = motionRectangles.GetImageComparisonResult(); _ = imageComparisonInfo
        var originalBufferedImage *BufferedImage = imageComparisonInfo.GetBufferedImages()[0]; _ = originalBufferedImage
        var bufferedImageInfo *BufferedImageInfo = BufferedImageInfoFactory_GetInstance(imageComparisonInfo.ImageWidth, imageComparisonInfo.ImageHeight, originalBufferedImage.GetType()); _ = bufferedImageInfo
        var poolInterface PoolInterface = BufferedImagePoolSingleton_GetInstance(); _ = poolInterface
        for index := int(0); index < this.NUMBER_OF_IMAGES; index++ {
            bufferedImageCacheables[index] = any(poolInterface.Remove(bufferedImageInfo)).(*BufferedImageCacheable)
        }
        bufferedImageCacheables[0].GetBufferedImage().SetData(imageComparisonInfo.GetBufferedImages()[1].GetData())
        if this.NUMBER_OF_IMAGES > 1 {
            bufferedImageCacheables[1].GetBufferedImage().SetData(imageComparisonInfo.GetBufferedImages()[1].GetData())
        }
        var bufferedImageGraphicsArray []IGraphics; _ = bufferedImageGraphicsArray
        bufferedImageGraphicsArray = make([]IGraphics, this.NUMBER_OF_IMAGES)
        for index := int(0); index < this.NUMBER_OF_IMAGES; index++ {
            bufferedImageGraphicsArray[index] = bufferedImageCacheables[index].GetBufferedImage().GetGraphics()
        }
        var rectangelList IBasicArrayList = motionRectangles.GetVector(); _ = rectangelList
        var size int = rectangelList.Size(); _ = size
        for index2 := int(0); index2 < size; index2++ {
            var motionRectangle *MotionRectangle = any(rectangelList.Get(index2)).(*MotionRectangle); _ = motionRectangle
            var rectangle *Rectangle = motionRectangle.GetRectangle(); _ = rectangle
            for index := int(0); index < this.NUMBER_OF_IMAGES; index++ {
                bufferedImageGraphicsArray[index].SetColor(Color_ORANGE)
                bufferedImageGraphicsArray[index].DrawRect(rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height)
            }
        }
        var imagePersistanceUtil *ImagePersistanceUtil = ImagePersistanceUtil_GetInstance(); _ = imagePersistanceUtil
        if this.NUMBER_OF_IMAGES > 1 {
            imagePersistanceUtil.SaveWithImageIO(fmt.Sprint(filePathStringBuffer1), bufferedImageCacheables[1].GetBufferedImage())
        }
        imagePersistanceUtil.SaveWithImageIO(fmt.Sprint(filePathStringBuffer2), bufferedImageCacheables[0].GetBufferedImage())
        if this.NUMBER_OF_IMAGES > 2 {
            imagePersistanceUtil.SaveWithImageIO(fmt.Sprint(filePathStringBuffer3), bufferedImageCacheables[2].GetBufferedImage())
        }
}

