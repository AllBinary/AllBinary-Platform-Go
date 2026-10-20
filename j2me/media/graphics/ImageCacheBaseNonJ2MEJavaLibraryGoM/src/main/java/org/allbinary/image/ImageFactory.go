/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

//ImageCacheBaseAndroidJavaLibrary
type ImageFactory struct {
}

var ImageFactory_Instance *ImageFactory = NewImageFactory()

func NewImageFactory() *ImageFactory {
    this := &ImageFactory{}
    return this
}


func ImageFactory_GetInstance() *ImageFactory {
        return ImageFactory_Instance
}


// throws Exception
func (this *ImageFactory) CreateImageUrl(url string) IImage {
        return Image_CreateImage(url)
}


// throws Exception
func (this *ImageFactory) CreateImage(caller string, width int, height int) IImage {
        return Image_CreateImageWH(width, height)
}


// throws Exception
func (this *ImageFactory) CreateImageFromInputStream_1(key any, inputStream IInputStream) IImage {
        var image IImage = Image_CreateImageFromInputStream(inputStream); _ = image
        return image
}

// Java overloads of ImageFactory.createImageFromInputStream
func (this *ImageFactory) CreateImageFromInputStream(args ...any) IImage {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[any](args[0], pass) && Java_IsArg[IInputStream](args[1], pass) {
            return this.CreateImageFromInputStream_1(Java_Arg[any](args[0]), Java_Arg[IInputStream](args[1]))
        }
        if len(args) == 1 && Java_IsArg[IInputStream](args[0], pass) {
            return this.CreateImageFromInputStream_2(Java_Arg[IInputStream](args[0]))
        }
    }
    panic("No Java overload of ImageFactory.createImageFromInputStream for the arguments")
}


// throws Exception
func (this *ImageFactory) CreateImageFromInputStream_2(inputStream IInputStream) IImage {
        return Image_CreateImageFromInputStream(inputStream)
}

