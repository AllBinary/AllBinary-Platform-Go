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

type MediaDataFactory struct {
    AVI *MediaData
    BMP *MediaData
    CEL *MediaData
    FITS *MediaData
    FLI *MediaData
    GBR *MediaData
    GIF *MediaData
    GIH *MediaData
    JPG *MediaData
    PIX *MediaData
    PNG *MediaData
    PNM *MediaData
    PSD *MediaData
    PSP *MediaData
    PS *MediaData
    SGI *MediaData
    SWF *MediaData
    TGA *MediaData
    TIFF *MediaData
    WMF *MediaData
}

var MediaDataFactory_Instance *MediaDataFactory = NewMediaDataFactory()

func (this *MediaDataFactory) initFields() {
    this.AVI = NewMediaData("avi")
    this.BMP = NewMediaData("bmp")
    this.CEL = NewMediaData("cel")
    this.FITS = NewMediaData("fits")
    this.FLI = NewMediaData("fli")
    this.GBR = NewMediaData("gbr")
    this.GIF = NewMediaData("gif")
    this.GIH = NewMediaData("gih")
    this.JPG = NewMediaData("jpg")
    this.PIX = NewMediaData("pix")
    this.PNG = NewMediaData("png")
    this.PNM = NewMediaData("pnm")
    this.PSD = NewMediaData("psd")
    this.PSP = NewMediaData("psp")
    this.PS = NewMediaData("ps")
    this.SGI = NewMediaData("SGI")
    this.SWF = NewMediaData("swf")
    this.TGA = NewMediaData("tga")
    this.TIFF = NewMediaData("tiff")
    this.WMF = NewMediaData("wmf")
}

func NewMediaDataFactory() *MediaDataFactory {
    this := &MediaDataFactory{}
    this.initFields()
    return this
}


func MediaDataFactory_GetInstance() *MediaDataFactory {
        return MediaDataFactory_Instance
}

