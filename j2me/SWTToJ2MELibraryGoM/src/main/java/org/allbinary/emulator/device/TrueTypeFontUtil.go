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

//TTF
type TrueTypeFontUtil struct {
    TrueTypeFontUtilBase
    PreResourceImageUtil *PreResourceImageUtil
    WidthFloatArray []float32
    FontImage IOpenGLESImage
    WidthScale float32
}

var TrueTypeFontUtil_Instance *TrueTypeFontUtil = NewTrueTypeFontUtil()

func (this *TrueTypeFontUtil) initFields() {
    this.PreResourceImageUtil = PreResourceImageUtil_GetInstance()
    
    this.FontImage = OpenGLESImage_NULL_OPENGL_IMAGE
    
}


func TrueTypeFontUtil_GetInstance() *TrueTypeFontUtil {
        return TrueTypeFontUtil_Instance
}


func NewTrueTypeFontUtil() *TrueTypeFontUtil {
    this := &TrueTypeFontUtil{}
        this.TrueTypeFontUtilBase = *NewTrueTypeFontUtilBase(2)
        this.initFields()
        if AvianUtil_IsAvian() {
            this.WidthFloatArray = []float32{1504.0, 967.0, 527.0, 961.0, 957.0, 1087.0, 920.0, 954.0, 973.0, 961.0, 955.0, 1296.0, 1013.0, 1106.0, 1167.0, 815.0, 815.0, 1216.0, 1110.0, 170.0, 520.0, 1056.0, 815.0, 1448.0, 1142.0, 1345.0, 927.0, 1345.0, 1030.0, 920.0, 1096.0, 1119.0, 1219.0, 1841.0, 1166.0, 1147.0, 1005.0, 879.0, 965.0, 792.0, 964.0, 927.0, 753.0, 1034.0, 916.0, 196.0, 469.0, 877.0, 166.0, 1563.0, 916.0, 1007.0, 965.0, 964.0, 631.0, 777.0, 649.0, 917.0, 1026.0, 1548.0, 993.0, 1028.0, 795.0, 241.0, 798.0, 241.0, 905.0, 1477.0, 393.0, 961.0, 961.0, 961.0, 961.0, 555.0, 1029.0, 1010.0, 1378.0, 952.0, 952.0, 952.0, 952.0, 463.0, 463.0, 926.0, 961.0, 491.0, 930.0, 457.0, 457.0, 644.0, 643.0, 326.0, 186.0, 1219.0, 241.0, 1599.0, 961.0, 302.0, 711.0, 961.0, 961.0, 710.0, 141.0, 1504.0, 1504.0, 1504.0, 1504.0, 1504.0, 1504.0, 1504.0, 1504.0}
        } else {
            this.WidthFloatArray = []float32{1504.0, 967.0, 527.0, 961.0, 957.0, 1087.0, 920.0, 954.0, 973.0, 961.0, 955.0, 1296.0, 1013.0, 1106.0, 1167.0, 815.0, 815.0, 1216.0, 1110.0, 170.0, 520.0, 1056.0, 815.0, 1448.0, 1142.0, 1345.0, 927.0, 1345.0, 1030.0, 920.0, 1096.0, 1119.0, 1219.0, 1841.0, 1166.0, 1147.0, 1005.0, 879.0, 965.0, 792.0, 964.0, 927.0, 753.0, 1034.0, 916.0, 196.0, 469.0, 877.0, 166.0, 1563.0, 916.0, 1007.0, 965.0, 964.0, 631.0, 777.0, 649.0, 917.0, 1026.0, 1548.0, 993.0, 1028.0, 795.0, 241.0, 798.0, 241.0, 905.0, 1477.0, 393.0, 961.0, 555.0, 1029.0, 1010.0, 1378.0, 952.0, 463.0, 463.0, 926.0, 961.0, 491.0, 930.0, 457.0, 457.0, 644.0, 643.0, 326.0, 186.0, 1219.0, 241.0, 1599.0, 961.0, 302.0, 711.0, 961.0, 961.0, 710.0, 141.0, 1504.0, 1504.0}
        }
        this.WidthScale = func() float32 { if (this.Scale == 2) { return 1.2 }; return 1 }()
    return this
}


func (this *TrueTypeFontUtil) GetFontBitmapGL(gl GL10, filename string, cellSize int, basicColor *BasicColor) IImage {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IImage) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, commonStrings.EXCEPTION, e)
            panic(errors.New(fmt.Sprint()))
                }
            }()
            if this.FontImage == OpenGLESImage_NULL_OPENGL_IMAGE {
                var canvasStrings *CanvasStrings = CanvasStrings_GetInstance(); _ = canvasStrings
                var openGLCapabilities *OpenGLCapabilities = OpenGLCapabilities_GetInstance(); _ = openGLCapabilities
                if openGLCapabilities.IsTextureSizeValid(1024) && this.Scale == 2 {
                    var FONT_ATLAS_1024 string = "/font_1024.png"; _ = FONT_ATLAS_1024
                    var image IImage = Image_CreateImage(ResourceUtil_GetInstance().GetResourceAsStream(FONT_ATLAS_1024)); _ = image
                    image.SetName(canvasStrings.FONT_ATLAS)
                    this.FontImage = any(this.PreResourceImageUtil.Encapsulate(image)).(IOpenGLESImage)
                } else if openGLCapabilities.IsTextureSizeValid(512) {
                    var FONT_ATLAS_512 string = "/font_512.png"; _ = FONT_ATLAS_512
                    var image IImage = Image_CreateImage(ResourceUtil_GetInstance().GetResourceAsStream(FONT_ATLAS_512)); _ = image
                    image.SetName(canvasStrings.FONT_ATLAS)
                    this.FontImage = any(this.PreResourceImageUtil.Encapsulate(image)).(IOpenGLESImage)
                } else {
                    var FONT_ATLAS_512 string = "/font_512.png"; _ = FONT_ATLAS_512
                    var image IImage = Image_CreateImage(ResourceUtil_GetInstance().GetResourceAsStream(FONT_ATLAS_512)); _ = image
                    image.SetName(canvasStrings.FONT_ATLAS)
                    this.FontImage = any(this.PreResourceImageUtil.Encapsulate(image)).(IOpenGLESImage)
                }
                tryControl1, tryResult1 = 1, this.FontImage; return
            } else {
                tryControl1, tryResult1 = 1, this.FontImage; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TrueTypeFontUtil) SaveFontAtlasAsFile() {
        var image IImage = this.GetFontBitmap(nil, this.CellSize, BasicColorFactory_GetInstance().WHITE); _ = image
        var imageLoader *ImageLoader = NewImageLoader(); _ = imageLoader
        var imageData *ImageData = (any(image).(*SwtMutableImage)).Image.GetImageData(); _ = imageData
        imageData.TransparentPixel = imageData.GetPixel(0, 0)
        imageLoader.Data = []*ImageData{imageData}
        imageLoader.Save(CommonSeps_GetInstance().PERIOD + CanvasStrings_GetInstance().FONT_ATLAS, SWT_IMAGE_PNG)
}


func (this *TrueTypeFontUtil) GetFontBitmap(filename string, cellSize int, basicColor *BasicColor) IImage {
        var cellsPerRow2 int = this.CELLS_PER_ROW * 2; _ = cellsPerRow2
        var cellsPerRow3 int = this.CELLS_PER_ROW * 3; _ = cellsPerRow3
        var cellsPerRow4 int = this.CELLS_PER_ROW * 4; _ = cellsPerRow4
        var cellsPerRow5 int = this.CELLS_PER_ROW * 5; _ = cellsPerRow5
        var cellsPerRow6 int = this.CELLS_PER_ROW * 6; _ = cellsPerRow6
        var cellsPerRow7 int = this.CELLS_PER_ROW * 7; _ = cellsPerRow7
        var image IImage = Image_CreateImage(this.TextureSize, this.TextureSize); _ = image
        var graphics IGraphics = image.GetGraphics(); _ = graphics
        graphics.SetColor(basicColor.IntValue())
        var biggestHeight int = 0; _ = biggestHeight
        var bounds *Rectangle = NewRectangle(PointFactory_GetInstance().CreateXY(0, 0), cellSize, cellSize); _ = bounds
        var x int; _ = x
        var y int; _ = y
        for index := int(0); index < this.Size; index++ {
            this.CharacterArray[0] = this.Pattern.CharAt(index)
            this._characterWidth[index] = bounds.GetMaxX()
            if bounds.GetMaxY() - bounds.GetPoint().GetX() > biggestHeight {
                biggestHeight = bounds.GetMaxY() - bounds.GetPoint().GetX()
            }
            x = (index % this.CELLS_PER_ROW) * cellSize
            x += (cellSize >> 1)
            x -= (this._characterWidth[index] >> 1)
            y = -(cellSize >> 1)
            if index >= this.CELLS_PER_ROW {
                y += cellSize
            }
            if index >= cellsPerRow2 {
                y += cellSize
            }
            if index >= cellsPerRow3 {
                y += cellSize
            }
            if index >= cellsPerRow4 {
                y += cellSize
            }
            if index >= cellsPerRow5 {
                y += cellSize
            }
            if index >= cellsPerRow6 {
                y += cellSize
            }
            if index >= cellsPerRow7 {
                y += cellSize
            }
            y += cellSize
            y -= (cellSize >> 2)
            graphics.DrawChar(this.CharacterArray[0], x, y, 0)
        }
        return image
}


func (this *TrueTypeFontUtil) GetFontWidths(filename string, fontSize int) []int {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        for index := int(0); index < this.Size; index++ {
            this.CharacterArray[0] = this.Pattern.CharAt(index)
            if this.CharacterArray[0] == ' ' {
                this._characterWidth[index] = (fontSize / 4)
            } else {
                var shortPatternIndex int = this.ShortPattern.IndexOf(this.CharacterArray[0]); _ = shortPatternIndex
                var w int = int((this.WidthFloatArray[shortPatternIndex] / 75.0)) + 19; _ = w
                this._characterWidth[index] = w
                if this.CharacterArray[0] == '.' || this.CharacterArray[0] == '1' {
                    this._characterWidth[index] += 4
                }
                if this.CharacterArray[0] >= 'a' && this.CharacterArray[0] < 'k' {
                    this._characterWidth[index] -= 2
                }
                if this.CharacterArray[0] >= 'k' && this.CharacterArray[0] < 'v' {
                    this._characterWidth[index] -= 4
                }
                if this.CharacterArray[0] >= 'v' && this.CharacterArray[0] < 'z' {
                    this._characterWidth[index] -= 8
                }
                if this.CharacterArray[0] == 'D' {
                    this._characterWidth[index] -= 4
                }
                if this.CharacterArray[0] >= 'A' && this.CharacterArray[0] <= 'Z' {
                    this._characterWidth[index] -= 2
                }
            }
            this._characterWidth[index] *= this.WidthScale
        }
        return this._characterWidth
}

