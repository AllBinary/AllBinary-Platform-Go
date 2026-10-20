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
    LogUtil *LogUtil
    PreResourceImageUtil *PreResourceImageUtil
    RealFontSize int
    FontImage IOpenGLESImage
}

var TrueTypeFontUtil_Instance *TrueTypeFontUtil = NewTrueTypeFontUtil()

func (this *TrueTypeFontUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PreResourceImageUtil = PreResourceImageUtil_GetInstance()
    this.RealFontSize = 18
    this.FontImage = OpenGLESImage_NULL_OPENGL_IMAGE
}


func TrueTypeFontUtil_GetInstance() *TrueTypeFontUtil {
        return TrueTypeFontUtil_Instance
}


func NewTrueTypeFontUtil() *TrueTypeFontUtil {
    this := &TrueTypeFontUtil{}
        this.TrueTypeFontUtilBase = *NewTrueTypeFontUtilBase(1)
        this.initFields()
    return this
}


func (this *TrueTypeFontUtil) SaveFontAtlasAsFile() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, commonStrings.EXCEPTION, e)
                }
            }()
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            var file *File = ResourceUtil_GetInstance().GetContext().GetFilesDir(); _ = file
            var path string = file.GetAbsolutePath() + CanvasStrings_GetInstance().FONT_ATLAS; _ = path
            this.LogUtil.PutF(path, this, commonStrings.CONSTRUCTOR)
            var fos *FileOutputStream = NewFileOutputStream(path); _ = fos
            var bitmap *Bitmap = (any(this.FontImage.OpenGLESImage_().OpenGLBitmap.GetImage()).(AndroidImageInterface)).GetBitmap(); _ = bitmap
            bitmap.Compress(Bitmap_CompressFormat.PNG, 100, fos)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TrueTypeFontUtil) GetFontBitmapGL(gl GL10, filename string, cellSize int, basicColor *BasicColor) IImage {
        if this.FontImage == OpenGLESImage_NULL_OPENGL_IMAGE {
            var cellsPerRow2 int = this.CELLS_PER_ROW * 2; _ = cellsPerRow2
            var cellsPerRow3 int = this.CELLS_PER_ROW * 3; _ = cellsPerRow3
            var cellsPerRow4 int = this.CELLS_PER_ROW * 4; _ = cellsPerRow4
            var cellsPerRow5 int = this.CELLS_PER_ROW * 5; _ = cellsPerRow5
            var cellsPerRow6 int = this.CELLS_PER_ROW * 6; _ = cellsPerRow6
            var cellsPerRow7 int = this.CELLS_PER_ROW * 7; _ = cellsPerRow7
            var typeface *Typeface = Typeface_DEFAULT; _ = typeface
            var textureSize int = this.GetAsTextureSize(this.CELLS_PER_ROW * cellSize); _ = textureSize
            var bitmap *Bitmap = Bitmap_CreateBitmap(textureSize, textureSize, Bitmap_Config.ARGB_8888); _ = bitmap
            var canvas ICanvas = NewCanvas(bitmap); _ = canvas
            var paint *Paint = NewPaint(); _ = paint
            paint.SetTypeface(typeface)
            paint.SetTextSize(float32(this.RealFontSize))
            paint.SetAlpha(int(basicColor.Alpha))
            paint.SetColor(basicColor.IntValue())
            var biggestHeight int = 0; _ = biggestHeight
            var bounds *Rect = NewRect(); _ = bounds
            var x int; _ = x
            var y int; _ = y
            for index := int(0); index < this.Size; index++ {
                this.CharacterArray[0] = this.Pattern.CharAt(index)
                paint.GetTextBounds(this.CharacterArray, 0, 1, bounds)
                this._characterWidth[index] = bounds.Right
                if bounds.Bottom - bounds.Top > biggestHeight {
                    biggestHeight = bounds.Bottom - bounds.Top
                }
                x = (index % this.CELLS_PER_ROW) * cellSize
                x += (cellSize >> 1)
                x -= (this._characterWidth[index] >> 1)
                y = 0
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
                canvas.DrawText(this.CharacterArray, 0, 1, float32(x) - 3.0, float32(y) - 6.0, paint)
            }
            canvas.Save()
            var image IImage = AndroidImmutableImage_Create(bitmap); _ = image
            this.FontImage = any(this.PreResourceImageUtil.Encapsulate(image)).(IOpenGLESImage)
            return this.FontImage
        } else {
            return this.FontImage
        }
}


func (this *TrueTypeFontUtil) GetFontWidths(filename string, fontSize int) []int {
        var typeface *Typeface = Typeface_DEFAULT; _ = typeface
        var paint *Paint = NewPaint(); _ = paint
        paint.SetTypeface(typeface)
        paint.SetTextSize(float32(fontSize))
        paint.SetARGB(255, 255, 255, 255)
        var bounds *Rect = NewRect(); _ = bounds
        for index := int(0); index < this.Size; index++ {
            this.CharacterArray[0] = this.Pattern.CharAt(index)
            paint.GetTextBounds(this.CharacterArray, 0, 1, bounds)
            if index < this.LastCapIndex {
                if this.CharacterArray[0] == '1' {
                    this._characterWidth[index] = bounds.Right + 3
                } else if this.CharacterArray[0] == 'J' || this.CharacterArray[0] == 'V' || this.CharacterArray[0] == '2' || this.CharacterArray[0] == '9' || this.CharacterArray[0] == 'I' || this.CharacterArray[0] == 'N' || this.CharacterArray[0] == 'U' {
                    this._characterWidth[index] = bounds.Right + 1
                } else if this.CharacterArray[0] == '4' || this.CharacterArray[0] == 'C' || this.CharacterArray[0] == 'E' || this.CharacterArray[0] == 'O' {
                    this._characterWidth[index] = bounds.Right - 2
                } else if this.CharacterArray[0] == 'B' || this.CharacterArray[0] == 'D' || this.CharacterArray[0] == 'G' || this.CharacterArray[0] == 'H' || this.CharacterArray[0] == 'T' || this.CharacterArray[0] == 'W' {
                    this._characterWidth[index] = bounds.Right - 3
                } else if this.CharacterArray[0] == 'A' || this.CharacterArray[0] == 'Q' || this.CharacterArray[0] == 'R' {
                    this._characterWidth[index] = bounds.Right - 5
                } else if this.CharacterArray[0] == 'M' {
                    this._characterWidth[index] = bounds.Right - 6
                } else if this.CharacterArray[0] == 'm' {
                    this._characterWidth[index] = bounds.Right - 8
                } else {
                    this._characterWidth[index] = bounds.Right
                }
            } else {
                if this.CharacterArray[0] == ' ' {
                    this._characterWidth[index] = bounds.Right + 10
                } else if this.CharacterArray[0] == 'l' || this.CharacterArray[0] == 'i' || this.CharacterArray[0] == 'j' || this.CharacterArray[0] == '.' || this.CharacterArray[0] == '!' || this.CharacterArray[0] == '|' {
                    this._characterWidth[index] = bounds.Right + 6
                } else if this.CharacterArray[0] == 'f' || this.CharacterArray[0] == 't' || this.CharacterArray[0] == 'u' || this.CharacterArray[0] == 'v' {
                    this._characterWidth[index] = bounds.Right + 1
                } else if this.CharacterArray[0] == 'r' {
                    this._characterWidth[index] = bounds.Right + 2
                } else if this.CharacterArray[0] == 'a' || this.CharacterArray[0] == 'b' || this.CharacterArray[0] == 'g' || this.CharacterArray[0] == 'u' {
                    this._characterWidth[index] = bounds.Right - 1
                } else if this.CharacterArray[0] == 'o' || this.CharacterArray[0] == 'e' {
                    this._characterWidth[index] = bounds.Right - 2
                } else if this.CharacterArray[0] == 'm' {
                    this._characterWidth[index] = bounds.Right - 7
                } else {
                    this._characterWidth[index] = bounds.Right
                }
            }
        }
        this._characterWidth[0] = (fontSize >> 1) - 2
        return this._characterWidth
}

