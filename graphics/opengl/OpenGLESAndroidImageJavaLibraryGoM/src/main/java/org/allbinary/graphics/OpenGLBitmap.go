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

type OpenGLBitmap struct {
    PlatformBitmapBase
    Bitmap *Bitmap
}


func (this *OpenGLBitmap) initFields() {
    
}


func NewOpenGLBitmap(bitmap any) *OpenGLBitmap {
    this := &OpenGLBitmap{}
    this.PlatformBitmapBase = *NewPlatformBitmapBase()
    this.initFields()
        this.Bitmap = any(bitmap).(*Bitmap)
    return this
}


func (this *OpenGLBitmap) GetPixels(pixels []int, offset int, stride int, x int, y int, width int, height int) {
        this.Bitmap.GetPixels(pixels, offset, stride, x, y, width, height)
}


func (this *OpenGLBitmap) SetPixels(pixels []int, offset int, stride int, x int, y int, width int, height int) {
        this.Bitmap.SetPixels(pixels, offset, stride, x, y, width, height)
}


func (this *OpenGLBitmap) GetWidth() int {
        return this.Bitmap.GetWidth()
}


func (this *OpenGLBitmap) GetHeight() int {
        return this.Bitmap.GetHeight()
}


func (this *OpenGLBitmap) Recycle() {
        this.Bitmap.Recycle()
}

