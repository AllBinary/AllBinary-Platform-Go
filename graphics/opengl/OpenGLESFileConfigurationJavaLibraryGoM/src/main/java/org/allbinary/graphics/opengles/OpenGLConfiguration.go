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

type OpenGLConfiguration struct {
    LogUtil *LogUtil
    FILE string
    CommonStrings *CommonStrings
    TsUtil *TsUtil
    Opengl bool
    Type *OpenGLFeature
    ImageColor *OpenGLFeature
    Color *OpenGLFeature
    VersionSelector *OpenGLFeature
}

var OpenGLConfiguration_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *OpenGLConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.FILE = "OpenGLConfiguration.dat"
    this.CommonStrings = CommonStrings_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.Opengl = false
    this.Type = OpenGLFeatureFactory_GetInstance().OPENGL_AS_GAME_THREAD
    this.ImageColor = OpenGLFeatureFactory_GetInstance().IMAGE_COLOR_DEPTH_4444
    this.Color = OpenGLFeatureFactory_GetInstance().IMAGE_COLOR_DEPTH_4444
    this.VersionSelector = OpenGLFeatureFactory_GetInstance().OPENGL_AUTO_SELECT
}


func OpenGLConfiguration_GetInstance() *OpenGLConfiguration {
        if OpenGLConfiguration_Instance == NullUtil_GetInstance().NULL_OBJECT {
            OpenGLConfiguration_Instance = NewOpenGLConfiguration()
        }
        return any(OpenGLConfiguration_Instance).(*OpenGLConfiguration)
}


func NewOpenGLConfiguration() *OpenGLConfiguration {
    this := &OpenGLConfiguration{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            if FileFactory_GetInstance().IsFile(this.FILE) {
                this.Read()
            } else {
                this.Write()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *OpenGLConfiguration) Read() {
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
        var fileInputStream IInputStream = fileInputStreamFactory.GetFileInputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileInputStream
        var dataInputStream IDataInputStream = NewDataInputStream(fileInputStream); _ = dataInputStream
        var openGLValue int = dataInputStream.ReadInt(); _ = openGLValue
        if openGLValue == 0 {
            this.SetOpenGL(false)
        } else if openGLValue == 1 {
            this.SetOpenGL(true)
        } else {
            panic(errors.New(fmt.Sprint("Invalid OpenGL Setting")))
        }
        var version string = dataInputStream.ReadUTF(); _ = version
        if this.TsUtil.CompareTo(version, openGLFeatureFactory.OPENGL_AUTO_SELECT.GetName()) == 0 {
            this.SetVersionSelector(openGLFeatureFactory.OPENGL_AUTO_SELECT)
        } else if this.TsUtil.CompareTo(version, openGLFeatureFactory.OPENGL_MINIMUM.GetName()) == 0 {
            this.SetVersionSelector(openGLFeatureFactory.OPENGL_MINIMUM)
        } else {
            panic(errors.New(fmt.Sprint("OpenGLConfiguration: Error reading version selector: " + version)))
        }
        var type_ string = dataInputStream.ReadUTF(); _ = type_
        if this.TsUtil.CompareTo(type_, openGLFeatureFactory.OPENGL_AS_GAME_THREAD.GetName()) == 0 {
            this.SetType(openGLFeatureFactory.OPENGL_AS_GAME_THREAD)
        } else if this.TsUtil.CompareTo(type_, openGLFeatureFactory.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS.GetName()) == 0 {
            this.SetType(openGLFeatureFactory.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS)
        } else {
            panic(errors.New(fmt.Sprint("OpenGLConfiguration: Error reading image color")))
        }
        var imageColor string = dataInputStream.ReadUTF(); _ = imageColor
        if this.TsUtil.CompareTo(imageColor, openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444.GetName()) == 0 {
            this.SetImageColor(openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444)
        } else if this.TsUtil.CompareTo(imageColor, openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444.GetName()) == 0 {
            this.SetImageColor(openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444)
        } else if this.TsUtil.CompareTo(imageColor, openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444.GetName()) == 0 {
            this.SetImageColor(openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444)
        } else {
            panic(errors.New(fmt.Sprint("OpenGLConfiguration: Error reading image color: " + fmt.Sprint(imageColor))))
        }
        var color string = dataInputStream.ReadUTF(); _ = color
        if this.TsUtil.CompareTo(color, openGLFeatureFactory.OPENGL_COLOR_DEPTH_4444.GetName()) == 0 {
            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_4444)
        } else if this.TsUtil.CompareTo(color, openGLFeatureFactory.OPENGL_COLOR_DEPTH_565.GetName()) == 0 {
            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_565)
        } else if this.TsUtil.CompareTo(color, openGLFeatureFactory.OPENGL_COLOR_DEPTH_8888.GetName()) == 0 {
            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_8888)
        } else {
            panic(errors.New(fmt.Sprint("OpenGLConfiguration: Error reading color: " + fmt.Sprint(color))))
        }
        PreLogUtil_Put("Read Configuration: " + this.ToString(), this, "read")
}


// throws Exception
func (this *OpenGLConfiguration) Write() {
        var closeable *Closeable = NullCloseable_NULL_CLOSEABLE; _ = closeable
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            StreamUtil_GetInstance().Close(closeable)
            }()
            PreLogUtil_Put("Write Configuration: " + this.ToString(), this, "write")
            var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
            var fileOutputStream IOutputStream = fileInputStreamFactory.GetFileOutputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileOutputStream
            var dataOutputStream *AbDataOutputStream = NewAbDataOutputStream(fileOutputStream); _ = dataOutputStream
            closeable = dataOutputStream
            if this.IsOpenGL() {
                dataOutputStream.WriteInt(1)
            } else {
                dataOutputStream.WriteInt(0)
            }
            dataOutputStream.WriteUTF(this.GetVersionSelector().GetName())
            dataOutputStream.WriteUTF(this.GetType().GetName())
            dataOutputStream.WriteUTF(this.GetImageColor().GetName())
            dataOutputStream.WriteUTF(this.GetColor().GetName())
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *OpenGLConfiguration) Init() {
        var features *Features = Features_GetInstance(); _ = features
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(MainFeatureFactory_GetInstance().STATIC) {
            var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
            if this.IsOpenGL() {
                if features.IsFeature(openGLFeatureFactory.OPENGL_2D) || features.IsFeature(openGLFeatureFactory.OPENGL_3D) || features.IsFeature(openGLFeatureFactory.OPENGL_2D_AND_3D) {
                    if !features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
                        this.LogUtil.PutF("Turning on OpenGL as Default", this, this.CommonStrings.INIT)
                        features.AddDefault(OpenGLFeatureFactory_GetInstance().OPENGL)
                        this.LogUtil.PutF("Using OpenGL Type Feature: " + fmt.Sprint(this.GetType()), this, this.CommonStrings.INIT)
                        features.AddDefault(this.GetType())
                        this.LogUtil.PutF("Using OpenGL ImageColor Feature: " + fmt.Sprint(this.GetImageColor()), this, this.CommonStrings.INIT)
                        features.AddDefault(this.GetImageColor())
                        this.LogUtil.PutF("Using OpenGL Color Feature: " + fmt.Sprint(this.GetColor()), this, this.CommonStrings.INIT)
                        features.AddDefault(this.GetColor())
                        this.LogUtil.PutF("Using OpenGL Version Selector Feature: " + fmt.Sprint(this.GetVersionSelector()), this, this.CommonStrings.INIT)
                        features.AddDefault(this.GetVersionSelector())
                    }
                } else {
                    this.LogUtil.PutF("OpenGL is enabled, but this game does not support OpenGL (Probably from installing with the same package name from a game that had OpenGL available)", this, this.CommonStrings.INIT)
                }
            } else {
                this.LogUtil.PutF("OpenGL is Off", this, this.CommonStrings.INIT)
            }
        } else {
            if this.IsOpenGL() && !features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
                PreLogUtil_Put("OpenGL is set but not Enabled since statics where not cleared (Reboot or Complete exit is Required)", this, this.CommonStrings.INIT)
            }
        }
        PreLogUtil_Put(this.ToString(), this, this.CommonStrings.INIT)
}


// throws Exception
func (this *OpenGLConfiguration) Update(gameFeature IFeature, colorLocked bool) {
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        var modified bool = false; _ = modified
        if gameFeature == openGLFeatureFactory.OPENGL {
            if features.IsFeature(gameFeature) {
                if !this.IsOpenGL() {
                    this.SetOpenGL(true)
                    modified = true
                }
            } else {
                if this.IsOpenGL() {
                    this.SetOpenGL(false)
                    modified = true
                }
            }
        }
        if gameFeature == openGLFeatureFactory.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS || gameFeature == openGLFeatureFactory.OPENGL_AS_GAME_THREAD {
            if features.IsFeature(gameFeature) {
                if gameFeature != this.GetType() {
                    this.SetType(any(gameFeature).(*OpenGLFeature))
                    modified = true
                }
            }
        }
        if gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444 || gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_565 || gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_8888 {
            if features.IsFeature(gameFeature) {
                if this.GetImageColor() != gameFeature {
                    this.SetImageColor(any(gameFeature).(*OpenGLFeature))
                    if colorLocked {
                        if gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_4444 {
                            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_4444)
                        } else if gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_565 {
                            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_565)
                        } else if gameFeature == openGLFeatureFactory.IMAGE_COLOR_DEPTH_8888 {
                            this.SetColor(openGLFeatureFactory.OPENGL_COLOR_DEPTH_8888)
                        }
                    }
                    modified = true
                }
            }
        }
        if gameFeature == openGLFeatureFactory.OPENGL_COLOR_DEPTH_4444 || gameFeature == openGLFeatureFactory.OPENGL_COLOR_DEPTH_565 || gameFeature == openGLFeatureFactory.OPENGL_COLOR_DEPTH_8888 {
            if features.IsFeature(gameFeature) {
                if this.GetColor() != gameFeature {
                    this.SetColor(any(gameFeature).(*OpenGLFeature))
                    modified = true
                }
            }
        }
        if gameFeature == openGLFeatureFactory.OPENGL_AUTO_SELECT || gameFeature == openGLFeatureFactory.OPENGL_MINIMUM {
            if features.IsFeature(gameFeature) {
                this.SetVersionSelector(any(gameFeature).(*OpenGLFeature))
            }
            modified = true
        }
        if modified {
            this.Write()
        }
}


func (this *OpenGLConfiguration) SetOpenGL(opengl bool) {
        this.Opengl = opengl
}


func (this *OpenGLConfiguration) IsOpenGL() bool {
        return this.Opengl
}


func (this *OpenGLConfiguration) SetImageColor(imageColor *OpenGLFeature) {
        this.ImageColor = imageColor
}


func (this *OpenGLConfiguration) GetImageColor() *OpenGLFeature {
        return this.ImageColor
}


func (this *OpenGLConfiguration) SetColor(color *OpenGLFeature) {
        this.Color = color
}


func (this *OpenGLConfiguration) GetColor() *OpenGLFeature {
        return this.Color
}


func (this *OpenGLConfiguration) SetVersionSelector(versionSelector *OpenGLFeature) {
        this.VersionSelector = versionSelector
}


func (this *OpenGLConfiguration) GetVersionSelector() *OpenGLFeature {
        return this.VersionSelector
}


func (this *OpenGLConfiguration) SetType(type_ *OpenGLFeature) {
        this.Type = type_
}


func (this *OpenGLConfiguration) GetType() *OpenGLFeature {
        return this.Type
}


func (this *OpenGLConfiguration) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append(" isOpenGL: ")
        stringBuffer.Appendboolean(this.IsOpenGL())
        stringBuffer.Append(" VersionSelector: ")
        stringBuffer.Append(stringUtil.ToString(this.GetVersionSelector()))
        stringBuffer.Append(" Type: ")
        stringBuffer.Append(stringUtil.ToString(this.GetType()))
        stringBuffer.Append(" Image Color: ")
        stringBuffer.Append(stringUtil.ToString(this.GetImageColor()))
        stringBuffer.Append(" Color: ")
        stringBuffer.Append(stringUtil.ToString(this.GetColor()))
        return fmt.Sprint(stringBuffer)
}

func (this *OpenGLConfiguration) String() string {
    return this.ToString()
}

