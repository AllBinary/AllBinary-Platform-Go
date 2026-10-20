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

type OpenGLCapabilities struct {
    LogUtil *LogUtil
    CUSTOM_GL_SURFACE_VIEW bool
    StringUtil *StringUtil
    Initialized bool
    GlVersionString string
    GlShaderVersionString string
    ShaderVersion int
    GlRenderer string
    GlVendor string
    GlExtensions string
    PossiblyAccelerated bool
    AcceleratedString string
    MaxTextureSize int
    VERSION_1_0 string
    VERSION_1_1 string
    VERSION_2_0 string
    VERSION_3_0 string
    VERSION_3_1 string
    VERSION_3_2 string
    VERSION_HIGHER_THAN_EXISTS string
    VERSION_UNK string
    GL_EXT_GPU_SHADER_100 string
    GL_EXT_GPU_SHADER4 string
    GlVersion string
    GlInstanceVersion string
    GlExtensionDrawTexture bool
    GlThreedDrawTexture bool
    GlExtensionGPUShader4 bool
    GlExtensionGPUShader100 bool
    VertexBufferObjectSupport bool
}

var OpenGLCapabilities_Instance *OpenGLCapabilities = NewOpenGLCapabilities()

func (this *OpenGLCapabilities) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CUSTOM_GL_SURFACE_VIEW = true
    this.StringUtil = StringUtil_GetInstance()
    this.Initialized = false
    this.GlVersionString = this.StringUtil.EMPTY_STRING
    this.GlShaderVersionString = this.StringUtil.EMPTY_STRING
    this.ShaderVersion = 0
    this.GlRenderer = this.StringUtil.EMPTY_STRING
    this.GlVendor = this.StringUtil.EMPTY_STRING
    this.GlExtensions = this.StringUtil.EMPTY_STRING
    
    this.AcceleratedString = this.StringUtil.EMPTY_STRING
    this.MaxTextureSize = 64
    this.VERSION_1_0 = "1.0"
    this.VERSION_1_1 = "1.1"
    this.VERSION_2_0 = "2.0"
    this.VERSION_3_0 = "3.0"
    this.VERSION_3_1 = "3.1"
    this.VERSION_3_2 = "3.2"
    this.VERSION_HIGHER_THAN_EXISTS = "999.999"
    this.VERSION_UNK = "Unk"
    this.GL_EXT_GPU_SHADER_100 = "GL_ARB_shading_language_100"
    this.GL_EXT_GPU_SHADER4 = "GL_EXT_gpu_shader4"
    this.GlVersion = this.GlVersionString
    this.GlInstanceVersion = this.VERSION_UNK
    
    
    
    
    
}


func OpenGLCapabilities_GetInstance() *OpenGLCapabilities {
        return OpenGLCapabilities_Instance
}


func NewOpenGLCapabilities() *OpenGLCapabilities {
    this := &OpenGLCapabilities{}
    this.initFields()
    return this
}


func (this *OpenGLCapabilities) IsInitialized() bool {
        return this.Initialized
}


func (this *OpenGLCapabilities) RequireInitialization() {
}


func (this *OpenGLCapabilities) InitCapabilities(gl GL10) {
        var METHOD_NAME string = "initGLCapabilities"; _ = METHOD_NAME
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, METHOD_NAME, e)
                }
            }()
            var features *Features = Features_GetInstance(); _ = features
            var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
            var openGLImageSpecificFactory *OpenGLImageSpecificFactory = OpenGLImageSpecificFactory_GetInstance(); _ = openGLImageSpecificFactory
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            this.GlVersionString = gl.GlGetString(GL10.GL_VERSION)
            var GL_SHADING_LANGUAGE_VERSION int = 0x8b8c; _ = GL_SHADING_LANGUAGE_VERSION
            this.GlShaderVersionString = gl.GlGetString(GL_SHADING_LANGUAGE_VERSION)
            if this.GlShaderVersionString == "" {
                this.GlShaderVersionString = this.StringUtil.EMPTY_STRING
            }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, METHOD_NAME, e)
                    }
                }()
                if this.GlShaderVersionString != "" && strings.Index(this.GlShaderVersionString, string('.')) >= 0 {
                    var startIndex int = strings.LastIndex(this.GlShaderVersionString, string(' ')); _ = startIndex
                    if startIndex >= 0 {
                        this.GlShaderVersionString = this.GlShaderVersionString[startIndex + 1:]
                    }
                    this.ShaderVersion = func() int { value, _ := strconv.Atoi(strings.ReplaceAll(this.GlShaderVersionString, string(CommonSeps_GetInstance().PERIOD), string(StringUtil_GetInstance().EMPTY_STRING))); return value }()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            this.GlRenderer = gl.GlGetString(GL10.GL_RENDERER)
            this.GlVendor = gl.GlGetString(GL10.GL_VENDOR)
            this.GlExtensions = gl.GlGetString(GL10.GL_EXTENSIONS)
            if strings.Index(this.GlExtensions, string(this.GL_EXT_GPU_SHADER_100)) >= 0 {
                this.GlExtensionGPUShader100 = true
            } else {
                this.GlExtensionGPUShader100 = false
            }
            if strings.Index(this.GlExtensions, string(this.GL_EXT_GPU_SHADER4)) >= 0 {
                this.GlExtensionGPUShader4 = true
            } else {
                this.GlExtensionGPUShader4 = false
            }
            if this.GlRenderer == "" {
                this.GlRenderer = this.StringUtil.EMPTY_STRING
            }
            if strings.Index(strings.ToLower(this.GlRenderer), string("pixelflinger")) >= 0 {
                this.AcceleratedString = "Probably Not for " + this.GlRenderer
                this.PossiblyAccelerated = false
            } else {
                this.AcceleratedString = "Probably for " + this.GlRenderer
                this.PossiblyAccelerated = true
            }
            this.GlExtensionDrawTexture = false
            if this.GlVersionString == "" {
                this.GlVersionString = this.StringUtil.EMPTY_STRING
                this.GlVersion = this.VERSION_UNK
            } else if strings.Index(this.GlVersionString, string(" 1.0")) >= 0 {
                this.GlVersion = this.VERSION_1_0
            } else if strings.Index(this.GlVersionString, string(" 1.1")) >= 0 {
                this.GlVersion = this.VERSION_1_1
            } else {
                this.GlVersion = this.VERSION_UNK
            }
            if func() bool { _, ok := any(gl).(GL11); return ok }() {
                this.GlInstanceVersion = this.VERSION_1_1
            } else if func() bool { _, ok := any(gl).(GL10); return ok }() {
                this.GlInstanceVersion = this.VERSION_1_0
            }
            if this.GlVendor == "" {
                this.GlVendor = this.StringUtil.EMPTY_STRING
            }
            if this.GlExtensions == "" {
                this.GlExtensions = this.StringUtil.EMPTY_STRING
            }
            if this.PossiblyAccelerated {
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("VBO:?").Appendboolean((this.GlInstanceVersion == this.VERSION_1_1)).Append("||").Appendboolean(this.IsExtension(openGLFeatureFactory.OPENGL_VERTEX_BUFFER_OBJECT))), this, METHOD_NAME)
                if (this.GlInstanceVersion == this.VERSION_1_1 || this.IsExtension(openGLFeatureFactory.OPENGL_VERTEX_BUFFER_OBJECT)) {
                    if AndroidUtil_IsAndroid() {
                        PreLogUtil_Put("VBO implementation was detected, but disabled by default for Android", this, METHOD_NAME)
                    } else {
                        PreLogUtil_Put("VBO implementation was detected, but disabled for now", this, METHOD_NAME)
                    }
                }
            }
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            if features.IsDefault(openGLFeatureFactory.OPENGL_AUTO_SELECT) {
                if this.VertexBufferObjectSupport {
                    openGLImageSpecificFactory.SetImageFactory(NewOpenGLESGL11VBOImageFactory())
                } else {
                    stringBuffer.Append("OpenGL is on but ")
                    stringBuffer.Append(stringUtil.ToString(openGLFeatureFactory.OPENGL_DRAW_TEXTURE))
                    stringBuffer.Append(" was not available")
                    PreLogUtil_Put(fmt.Sprint(stringBuffer), this, METHOD_NAME)
                    openGLImageSpecificFactory.SetImageFactory(NewOpenGLESGL10ImageFactory())
                }
            } else {
                stringBuffer.Append(stringUtil.ToString(openGLFeatureFactory.OPENGL_AUTO_SELECT))
                stringBuffer.Append(" is not on")
                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, METHOD_NAME)
                openGLImageSpecificFactory.SetImageFactory(NewOpenGLESGL10ImageFactory())
            }
            var maxTextureSizeArray []int = make([]int, 1); _ = maxTextureSizeArray
            gl.GlGetIntegerv(GL10.GL_MAX_TEXTURE_SIZE, maxTextureSizeArray, 0)
            this.MaxTextureSize = maxTextureSizeArray[0]
            this.Initialized = true
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OpenGLCapabilities) IsExtension(gameFeature *OpenGLFeature) bool {
        var index int = strings.Index(this.GlExtensions, string(gameFeature.GetName())); _ = index
        if index >= 0 {
            return true
        } else {
            return false
        }
}


func (this *OpenGLCapabilities) IsGlExtensionDrawTexture() bool {
        this.RequireInitialization()
        return this.GlExtensionDrawTexture
}


func (this *OpenGLCapabilities) IsGlExtensionGPUShader100() bool {
        return this.GlExtensionGPUShader100
}


func (this *OpenGLCapabilities) IsGlExtensionGPUShader4() bool {
        this.RequireInitialization()
        return this.GlExtensionGPUShader4
}


func (this *OpenGLCapabilities) GetGlVersion() string {
        this.RequireInitialization()
        return this.GlVersion
}


func (this *OpenGLCapabilities) GetGlVersionString() string {
        this.RequireInitialization()
        return this.GlVersionString
}


func (this *OpenGLCapabilities) GetGlShaderVersion() string {
        this.RequireInitialization()
        return this.GlShaderVersionString
}


func (this *OpenGLCapabilities) IsGlThreedDrawTexture() bool {
        this.RequireInitialization()
        return this.GlThreedDrawTexture
}


func (this *OpenGLCapabilities) GetGlRenderer() string {
        this.RequireInitialization()
        return this.GlRenderer
}


func (this *OpenGLCapabilities) IsVertexBufferObjectSupport() bool {
        this.RequireInitialization()
        return this.VertexBufferObjectSupport
}


func (this *OpenGLCapabilities) IsTextureSizeValid(widthAndHeight int) bool {
        this.RequireInitialization()
        if this.MaxTextureSize >= widthAndHeight {
            return true
        }
        return false
}


func (this *OpenGLCapabilities) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("GL_VERSION: ")
        stringBuffer.Append(this.GlVersionString)
        stringBuffer.Append(" GL_SHADING_LANGUAGE_VERSION: ")
        stringBuffer.Append(this.GlShaderVersionString)
        stringBuffer.Append(" GL_RENDERER: ")
        stringBuffer.Append(this.GlRenderer)
        stringBuffer.Append(" GL_VENDOR: ")
        stringBuffer.Append(this.GlVendor)
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(" Is Accelerated: ")
        stringBuffer.Append(this.AcceleratedString)
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(" VBO Support: ")
        stringBuffer.Appendboolean(this.IsVertexBufferObjectSupport())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(" Max Texture Size: ")
        stringBuffer.Appendint(this.MaxTextureSize)
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(" GL_EXTENSIONS: ")
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, commonStrings.TOSTRING, e)
                }
            }()
            var tokenizer *Tokenizer = NewTokenizer(commonSeps.SPACE); _ = tokenizer
            var list IBasicArrayList = tokenizer.GetTokensFromString(this.GlExtensions, NewBasicArrayListD()); _ = list
            var size int = list.Size(); _ = size
            for index := int(0); index < size; index++ {
                stringBuffer.Append(commonSeps.NEW_LINE)
                stringBuffer.Append(this.StringUtil.ToString(list.BasicArrayList_().ObjectArray[index]))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return fmt.Sprint(stringBuffer)
}

func (this *OpenGLCapabilities) String() string {
    return this.ToString()
}

