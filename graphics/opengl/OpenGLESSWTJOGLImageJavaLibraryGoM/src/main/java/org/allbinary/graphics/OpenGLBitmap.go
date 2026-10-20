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
    Image IImage
}


func (this *OpenGLBitmap) initFields() {
    
}


func NewOpenGLBitmap(bitmap any) *OpenGLBitmap {
    this := &OpenGLBitmap{}
    this.PlatformBitmapBase = *NewPlatformBitmapBase()
    this.initFields()
        this.Image = any(bitmap).(IImage)
    return this
}


func (this *OpenGLBitmap) GetImage() IImage {
        return this.Image
}


func (this *OpenGLBitmap) GetDepth() int {
        if this.Image.IsMutable() {
            var mutableImage *SwtMutableImage = any(this.Image).(*SwtMutableImage); _ = mutableImage
            return mutableImage.GetDepth()
        } else {
            var immutableImage *SwtImmutableImage = any(this.Image).(*SwtImmutableImage); _ = immutableImage
            return immutableImage.GetDepth()
        }
}


func (this *OpenGLBitmap) GetPixels(pixels []int, offset int, stride int, x int, y int, width int, height int) {
        if this.Image.IsMutable() {
            var mutableImage *SwtMutableImage = any(this.Image).(*SwtMutableImage); _ = mutableImage
            mutableImage.GetRGB(pixels, offset, stride, x, y, width, height)
        } else {
            var immutableImage *SwtImmutableImage = any(this.Image).(*SwtImmutableImage); _ = immutableImage
            immutableImage.GetRGB(pixels, offset, stride, x, y, width, height)
        }
}


func (this *OpenGLBitmap) SetPixels(pixels []int, offset int, stride int, x int, y int, width int, height int) {
        this.Image.SetRGB2(pixels, offset, height, x, y, width, height)
}


func (this *OpenGLBitmap) GetWidth() int {
        return this.Image.GetWidth()
}


func (this *OpenGLBitmap) GetHeight() int {
        return this.Image.GetHeight()
}


func (this *OpenGLBitmap) Recycle() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}

