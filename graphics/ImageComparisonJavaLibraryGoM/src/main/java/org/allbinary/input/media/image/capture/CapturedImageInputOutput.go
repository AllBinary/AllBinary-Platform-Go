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

type CapturedImageInputOutput struct {
    LogUtil *LogUtil
}

// Implements ImageIOInterface


func (this *CapturedImageInputOutput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewCapturedImageInputOutput() *CapturedImageInputOutput {
    this := &CapturedImageInputOutput{}
    this.initFields()
    return this
}


// throws Exception
func (this *CapturedImageInputOutput) Save_1(frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var capturedBufferedImageCacheable *BufferedImageFrameCacheable = (any((any(CapturedBufferedImagesCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(frame)).(*BufferedImageFrameCacheable)); _ = capturedBufferedImageCacheable
        this.LogUtil.PutF(("Saving: " + fmt.Sprint(capturedBufferedImageCacheable)), this, commonStrings.SAVE)
        this.Save(capturedBufferedImageCacheable.GetBufferedImage(), frame)
}

// Java overloads of CapturedImageInputOutput.save
func (this *CapturedImageInputOutput) Save(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[int64](args[0], pass) {
            this.Save_1(Java_Arg[int64](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[*BufferedImage](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            this.Save_2(Java_Arg[*BufferedImage](args[0]), Java_Arg[int64](args[1]))
            return
        }
    }
    panic("No Java overload of CapturedImageInputOutput.save for the arguments")
}


func (this *CapturedImageInputOutput) Save_2(bufferedImage *BufferedImage, frame int64) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
        var filePathStringBuffer *StringMaker = NewStringMaker(); _ = filePathStringBuffer
        filePathStringBuffer.Append(ImageOutputData_SAVE_PATH)
        filePathStringBuffer.Append(LongUtil_FillIn(fmt.Sprint(frame)))
        filePathStringBuffer.Append(MediaDataFactory_GetInstance().JPG.GetExtension())
        var filePath string = fmt.Sprint(filePathStringBuffer); _ = filePath
        this.LogUtil.PutF(("Image File Path: " + filePath + imageUtil.ToString(bufferedImage)), this, commonStrings.SAVE)
        var imagePersistanceUtil *ImagePersistanceUtil = ImagePersistanceUtil_GetInstance(); _ = imagePersistanceUtil
        imagePersistanceUtil.SaveWithImageIO(filePath, bufferedImage)
}

