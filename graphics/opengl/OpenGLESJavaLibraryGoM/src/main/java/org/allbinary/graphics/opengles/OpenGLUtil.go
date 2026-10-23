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

type OpenGLUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RenderStrings *RendererStrings
    CanvasStrings *CanvasStrings
    DisplayInfoSingleton *DisplayInfoSingleton
    PreResourceImageUtil *PreResourceImageUtil
    LockObject *SynchObject
    AnyType *SynchObject
    List IBasicArrayList
    RunnableList IBasicArrayList
    Created bool
    SurfaceCreatedAndInitialized bool
}

var OpenGLUtil_Instance *OpenGLUtil = NewOpenGLUtil()

func (this *OpenGLUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.RenderStrings = RendererStrings_GetInstance()
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.PreResourceImageUtil = PreResourceImageUtil_GetInstance()
    this.LockObject = NewSynchObject()
    this.AnyType = NewSynchObject()
    this.List = NewBasicArrayListD()
    this.RunnableList = NewBasicArrayListD()
    this.Created = false
    this.SurfaceCreatedAndInitialized = false
}

func NewOpenGLUtil() *OpenGLUtil {
    this := &OpenGLUtil{}
    this.initFields()
    return this
}


func OpenGLUtil_GetInstance() *OpenGLUtil {
        return OpenGLUtil_Instance
}


func (this *OpenGLUtil) OnSurfaceCreated_1(gl GL10) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.RenderStrings.ON_SURFACE_CREATED, e)
                }
            }()
            if !this.Created {
                this.Created = true
                PreLogUtil_Put(fmt.Sprint(CommonLabels_GetInstance().START_LABEL) + fmt.Sprint(OpenGLCapabilities_GetInstance()), this, this.RenderStrings.ON_SURFACE_CREATED)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of OpenGLUtil.onSurfaceCreated
func (this *OpenGLUtil) OnSurfaceCreated(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[GL10](args[0], pass) {
            this.OnSurfaceCreated_1(Java_Arg[GL10](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[GL10](args[0], pass) && Java_IsArg[ILoadTextures](args[1], pass) {
            this.OnSurfaceCreated_2(Java_Arg[GL10](args[0]), Java_Arg[ILoadTextures](args[1]))
            return
        }
    }
    panic("No Java overload of OpenGLUtil.onSurfaceCreated for the arguments")
}


func (this *OpenGLUtil) OnSurfaceCreated_2(gl GL10, loadTextures ILoadTextures) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.RenderStrings.ON_SURFACE_CREATED, e)
                }
            }()
            loadTextures.Load(gl)
            OpenGLLogUtil_GetInstance().LogError(gl)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *OpenGLUtil) OnSurfaceChanged(gl GL10, graphics IOpenGLESGraphics) {
        if !this.SurfaceCreatedAndInitialized {
            graphics.Init()
            this.SurfaceCreatedAndInitialized = true
        }
        graphics.Update()
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.Update(graphics)
        (any(OpenGLImageCacheFactory_GetInstance()).(*OpenGLImageCache)).Update(gl)
        this.ProcessRunnables()
        this.Set(gl)
}


func (this *OpenGLUtil) AddImage(image IImage) IImage {
        var encapsulateImage IImage = this.PreResourceImageUtil.Encapsulate(image); _ = encapsulateImage
        if encapsulateImage != image {
            { // synchronized (this.AnyType)
                this.List.Add(encapsulateImage)
            }
            return encapsulateImage
        }
        return image
}


func (this *OpenGLUtil) Clear() {
        { // synchronized (this.AnyType)
            this.List.Clear()
        }
}


func (this *OpenGLUtil) AddRunnable(runnable Runnable) {
        { // synchronized (this.LockObject)
            if !this.RunnableList.Contains(runnable) {
                this.RunnableList.Add(runnable)
                this.DisplayInfoSingleton.Add(this.CanvasStrings.SCALED_IMAGES)
            }
        }
}


func (this *OpenGLUtil) ProcessRunnables() {
        { // synchronized (this.LockObject)
            var runnable Runnable; _ = runnable
            var size int = this.RunnableList.Size(); _ = size
            for index := int(0); index < size; index++ {
                runnable = any(this.RunnableList.Get(index)).(Runnable)
                runnable.Run()
            }
            this.RunnableList.Clear()
        }
}


// throws Exception
func (this *OpenGLUtil) Set(gl GL10) {
        var size int = this.List.Size(); _ = size
        var image IOpenGLESImage; _ = image
        for index2 := int(0); index2 < size; index2++ {
            image = any(this.List.Get(index2)).(IOpenGLESImage)
            image.Set(gl)
        }
}

