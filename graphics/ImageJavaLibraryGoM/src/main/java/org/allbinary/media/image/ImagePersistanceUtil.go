/*
 * 
 * To change this license header, choose License Headers in Project Properties. 
 * To change this template file, choose Tools | Templates 
 * and open the template in the editor.  
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

type ImagePersistanceUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var ImagePersistanceUtil_Instance *ImagePersistanceUtil = NewImagePersistanceUtil()

func (this *ImagePersistanceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewImagePersistanceUtil() *ImagePersistanceUtil {
    this := &ImagePersistanceUtil{}
    this.initFields()
    return this
}


func ImagePersistanceUtil_GetInstance() *ImagePersistanceUtil {
        return ImagePersistanceUtil_Instance
}


// throws Exception
func (this *ImagePersistanceUtil) SaveWithBatik(file *AbFile, bufferedImage *BufferedImage) {
        var batikPNGImageWriter *PNGImageWriter = NewPNGImageWriter(); _ = batikPNGImageWriter
        var fileOutputStream *AbFileOutputStream = AbFileOutputStream_CreateFromAbFile(file); _ = fileOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            fileOutputStream.Flush()
            StreamUtil_GetInstance().Close(fileOutputStream)
            }()
            batikPNGImageWriter.WriteImage(bufferedImage, fileOutputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.LogUtil.PutF("Wrote Image: " + fmt.Sprint(file.GetAbsolutePath()), this, this.CommonStrings.SAVE)
}


func (this *ImagePersistanceUtil) SaveWithImageIO_1(filePath string, bufferedImage *BufferedImage) {
        this.SaveWithImageIO(NewFile(filePath), bufferedImage)
}

// Java overloads of ImagePersistanceUtil.saveWithImageIO
func (this *ImagePersistanceUtil) SaveWithImageIO(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*BufferedImage](args[1], pass) {
            this.SaveWithImageIO_1(Java_Arg[string](args[0]), Java_Arg[*BufferedImage](args[1]))
            return
        }
        if len(args) == 2 && Java_IsArg[*File](args[0], pass) && Java_IsArg[*BufferedImage](args[1], pass) {
            this.SaveWithImageIO_2(Java_Arg[*File](args[0]), Java_Arg[*BufferedImage](args[1]))
            return
        }
    }
    panic("No Java overload of ImagePersistanceUtil.saveWithImageIO for the arguments")
}


func (this *ImagePersistanceUtil) SaveWithImageIO_2(file *File, bufferedImage *BufferedImage) {
        var writer *ImageWriter = nil; _ = writer
        var ios *ImageOutputStream = nil; _ = ios
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.SAVE, e)
                }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.SAVE, e2)
                    }
                }()
                if ios != nil {
                    ios.Flush()
                    if tryControl3 := func() (tryControl3 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        this.LogUtil.Put(this.CommonStrings.EXCEPTION, ios, this.CommonStrings.CLOSE, e)
                            }
                        }()
                        if ios != nil {
                            this.LogUtil.PutF(fmt.Sprint(ios), ios, this.CommonStrings.CLOSE)
                            ios.Close()
                        }
                        return
                    }(); tryControl3 == 1 {
                        tryControl2 = 1; return
                    }
                }
                if writer != nil {
                    writer.Dispose()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            }()
            var iter Iterator = ImageIO_GetImageWritersByFormatName("jpeg"); _ = iter
            if !iter.HasNext() {
                this.LogUtil.PutF("Unable to save image to jpeg file type.", this, this.CommonStrings.SAVE)
                tryControl1 = 1; return
            }
            writer = any(iter.Next()).(*ImageWriter)
            ios = ImageIO_CreateImageOutputStream(file)
            writer.SetOutput(ios)
            var iwp *ImageWriteParam = writer.GetDefaultWriteParam(); _ = iwp
            iwp.SetCompressionMode(ImageWriteParam_MODE_EXPLICIT)
            iwp.SetCompressionQuality(0.95)
            writer.Write(nil, NewIIOImage(bufferedImage, nil, nil), iwp)
            this.LogUtil.PutF("Wrote Image: " + fmt.Sprint(file.GetAbsolutePath()), this, this.CommonStrings.SAVE)
            return
        }(); tryControl1 == 1 {
            return
        }
}

