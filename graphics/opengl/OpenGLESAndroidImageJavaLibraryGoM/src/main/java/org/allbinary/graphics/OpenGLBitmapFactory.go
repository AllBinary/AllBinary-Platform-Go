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

type OpenGLBitmapFactory struct {
    PlatformBitmapBaseFactory
    ARGB_8888 *Config
}

var OpenGLBitmapFactory_Instance *OpenGLBitmapFactory = NewOpenGLBitmapFactory()

func (this *OpenGLBitmapFactory) initFields() {
    this.ARGB_8888 = NewConfig(Bitmap_Config.ARGB_8888)
}

func NewOpenGLBitmapFactory() *OpenGLBitmapFactory {
    this := &OpenGLBitmapFactory{}
    this.PlatformBitmapBaseFactory = *NewPlatformBitmapBaseFactory()
    this.initFields()
    return this
}


func OpenGLBitmapFactory_GetInstance() *OpenGLBitmapFactory {
        return OpenGLBitmapFactory_Instance
}


func (this *OpenGLBitmapFactory) CreateBitmap_1(width int, height int, config *Config) *OpenGLBitmap {
        return NewOpenGLBitmap(Bitmap_CreateBitmap(width, height, config.BitmapConfig))
}

// Java overloads of OpenGLBitmapFactory.createBitmap
func (this *OpenGLBitmapFactory) CreateBitmap(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[*Config](args[2], pass) {
            return this.CreateBitmap_1(Java_Arg[int](args[0]), Java_Arg[int](args[1]), Java_Arg[*Config](args[2]))
        }
        if len(args) == 1 && Java_IsArg[IImage](args[0], pass) {
            return this.CreateBitmap_2(Java_Arg[IImage](args[0]))
        }
    }
    panic("No Java overload of OpenGLBitmapFactory.createBitmap for the arguments")
}


// throws IOException
func (this *OpenGLBitmapFactory) DecodeStream(inputStream IInputStream) *OpenGLBitmap {
        return NewOpenGLBitmap(BitmapFactory_DecodeStream(inputStream))
}


func (this *OpenGLBitmapFactory) CreateBitmap_2(image IImage) IPlatformBitmapBase {
        return NewOpenGLBitmap((any(image).(AndroidImageInterface)).GetBitmap())
}

type Config struct {
    BitmapConfig *Config
}


func (this *Config) initFields() {
        
}
    

    func NewOpenGLBitmapFactory(bitmapConfig *Config) *OpenGLBitmapFactory {
        this := &OpenGLBitmapFactory{}
        this.initFields()
            this.BitmapConfig = bitmapConfig
        return this
    }

