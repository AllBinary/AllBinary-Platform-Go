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

type OpenGLFeatureFactory struct {
    OPENGL_OPTIONS *OpenGLFeature
    OPENGL *OpenGLFeature
    OPENGL_SURFACE_VIEW *OpenGLFeature
    OPENGL_2D *OpenGLFeature
    OPENGL_3D *OpenGLFeature
    OPENGL_2D_AND_3D *OpenGLFeature
    OPENGL_AUTO_SELECT *OpenGLFeature
    OPENGL_MINIMUM *OpenGLFeature
    OPENGL_AS_GAME_THREAD *OpenGLFeature
    OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS *OpenGLFeature
    IMAGE_COLOR_DEPTH_4444 *OpenGLFeature
    IMAGE_COLOR_DEPTH_8888 *OpenGLFeature
    IMAGE_COLOR_DEPTH_565 *OpenGLFeature
    OPENGL_COLOR_DEPTH_4444 *OpenGLFeature
    OPENGL_COLOR_DEPTH_8888 *OpenGLFeature
    OPENGL_COLOR_DEPTH_565 *OpenGLFeature
    OPENGL_DRAW_TEXTURE *OpenGLFeature
    OPENGL_VERTEX_BUFFER_OBJECT *OpenGLFeature
    OPENGL_SIMPLE_OBJECT3D_PROCESSOR *OpenGLFeature
    OPENGL_COMPLEX_OBJECT3D_PROCESSOR *OpenGLFeature
    OPENGL_SIMPLE_TEXTURE_PROCESSOR *OpenGLFeature
    OPENGL_COMPLEX_TEXTURE_PROCESSOR *OpenGLFeature
}

var OpenGLFeatureFactory_Instance *OpenGLFeatureFactory = NewOpenGLFeatureFactory()

func (this *OpenGLFeatureFactory) initFields() {
    this.OPENGL_OPTIONS = NewOpenGLFeature("OpenGL Options")
    this.OPENGL = NewOpenGLFeature("OpenGL")
    this.OPENGL_SURFACE_VIEW = NewOpenGLFeature("GLSurfaceView")
    this.OPENGL_2D = NewOpenGLFeature("OpenGL 2D")
    this.OPENGL_3D = NewOpenGLFeature("OpenGL 3D")
    this.OPENGL_2D_AND_3D = NewOpenGLFeature("OpenGL 2D/3D")
    this.OPENGL_AUTO_SELECT = NewOpenGLFeature("Auto Select")
    this.OPENGL_MINIMUM = NewOpenGLFeature("Minimum")
    this.OPENGL_AS_GAME_THREAD = NewOpenGLFeature("As game thread")
    this.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS = NewOpenGLFeature("As different threads")
    this.IMAGE_COLOR_DEPTH_4444 = NewOpenGLFeature("4444")
    this.IMAGE_COLOR_DEPTH_8888 = NewOpenGLFeature("8888")
    this.IMAGE_COLOR_DEPTH_565 = NewOpenGLFeature("565")
    this.OPENGL_COLOR_DEPTH_4444 = NewOpenGLFeature("4444")
    this.OPENGL_COLOR_DEPTH_8888 = NewOpenGLFeature("8888")
    this.OPENGL_COLOR_DEPTH_565 = NewOpenGLFeature("565")
    this.OPENGL_DRAW_TEXTURE = NewOpenGLFeature("GL_OES_draw_texture")
    this.OPENGL_VERTEX_BUFFER_OBJECT = NewOpenGLFeature("vertex_buffer_object")
    this.OPENGL_SIMPLE_OBJECT3D_PROCESSOR = NewOpenGLFeature("Simple Object3d Processor")
    this.OPENGL_COMPLEX_OBJECT3D_PROCESSOR = NewOpenGLFeature("Complex Object3d Processor")
    this.OPENGL_SIMPLE_TEXTURE_PROCESSOR = NewOpenGLFeature("Simple Texture Processor")
    this.OPENGL_COMPLEX_TEXTURE_PROCESSOR = NewOpenGLFeature("Complex Texture Processor")
}


func OpenGLFeatureFactory_GetInstance() *OpenGLFeatureFactory {
        return OpenGLFeatureFactory_Instance
}


func NewOpenGLFeatureFactory() *OpenGLFeatureFactory {
    this := &OpenGLFeatureFactory{}
    this.initFields()
    return this
}

