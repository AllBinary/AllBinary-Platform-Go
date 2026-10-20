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

type ComparisonImageInputOutput struct {
    LogUtil *LogUtil
}

// Implements ImageIOInterface

var ComparisonImageInputOutput_ROOT_NAME string = "_Changed_"

func (this *ComparisonImageInputOutput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewComparisonImageInputOutput() *ComparisonImageInputOutput {
    this := &ComparisonImageInputOutput{}
    this.initFields()
    return this
}


// throws Exception
func (this *ComparisonImageInputOutput) Save_1(frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if frame > 0 {
            var imageComparisonResultFrameCacheable *ImageComparisonResultFrameCacheable = any((any(ImageComparisonResultCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(frame)).(*ImageComparisonResultFrameCacheable); _ = imageComparisonResultFrameCacheable
            if imageComparisonResultFrameCacheable != nil {
                var imageComparisonResult *ImageComparisonResult = imageComparisonResultFrameCacheable.GetImageComparisonResult(); _ = imageComparisonResult
                this.Save(imageComparisonResult, imageComparisonResultFrameCacheable.GetFrame())
            } else {
                this.LogUtil.PutF("Comparison Results Not Available for Output: " + fmt.Sprint(frame), this, commonStrings.SAVE)
            }
        } else {
            this.LogUtil.PutF("No Comparison Results: for first frame: " + fmt.Sprint(frame), this, commonStrings.SAVE)
        }
}

// Java overloads of ComparisonImageInputOutput.save
func (this *ComparisonImageInputOutput) Save(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            this.Save_1(Java_Arg[int64](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[*ImageComparisonResult](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            this.Save_2(Java_Arg[*ImageComparisonResult](args[0]), Java_Arg[int64](args[1]))
            return
        }
    }
    panic("No Java overload of ComparisonImageInputOutput.save for the arguments")
}


// throws Exception
func (this *ComparisonImageInputOutput) Save_2(imageComparisonResult *ImageComparisonResult, frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var filePathStringBuffer *StringMaker = NewStringMaker(); _ = filePathStringBuffer
        filePathStringBuffer.Append(ImageOutputData_SAVE_PATH)
        filePathStringBuffer.Append(LongUtil_FillIn(fmt.Sprint(frame)))
        filePathStringBuffer.Append(ComparisonImageInputOutput_ROOT_NAME)
        var filePathStringBuffer1 *StringMaker = NewStringMaker(); _ = filePathStringBuffer1
        filePathStringBuffer1.Append(fmt.Sprint(filePathStringBuffer))
        filePathStringBuffer1.Append("_1")
        filePathStringBuffer1.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        var filePath1 string = fmt.Sprint(filePathStringBuffer1); _ = filePath1
        var filePathStringBuffer2 *StringMaker = NewStringMaker(); _ = filePathStringBuffer2
        filePathStringBuffer2.Append(fmt.Sprint(filePathStringBuffer))
        filePathStringBuffer2.Append("_2")
        filePathStringBuffer2.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        var filePath2 string = fmt.Sprint(filePathStringBuffer2); _ = filePath2
        var bufferedImageCacheables []*BufferedImageCacheable = ChangedPixelsUtil_GenerateBufferedImageChacheables(imageComparisonResult); _ = bufferedImageCacheables
        this.LogUtil.PutF("Comparison Image File Path 1: " + filePath1, this, commonStrings.SAVE)
        this.LogUtil.PutF("Comparison Image File Path 2: " + filePath2, this, commonStrings.SAVE)
        var bufferedImageArray []*BufferedImage = make([]*BufferedImage, 2); _ = bufferedImageArray
        bufferedImageArray[0] = bufferedImageCacheables[0].GetBufferedImage()
        bufferedImageArray[1] = bufferedImageCacheables[1].GetBufferedImage()
        var imagePersistanceUtil *ImagePersistanceUtil = ImagePersistanceUtil_GetInstance(); _ = imagePersistanceUtil
        imagePersistanceUtil.SaveWithImageIO(filePath1, bufferedImageArray[0])
        imagePersistanceUtil.SaveWithImageIO(filePath2, bufferedImageArray[1])
}

