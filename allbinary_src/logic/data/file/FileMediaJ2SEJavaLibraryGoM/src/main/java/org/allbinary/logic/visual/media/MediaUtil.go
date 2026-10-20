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

//import com.sun.imageio.plugins.common.ImageUtil;
type MediaUtil struct {
    LogUtil *LogUtil
}

var MediaUtil_Instance *MediaUtil = NewMediaUtil()

func (this *MediaUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func MediaUtil_GetInstance() *MediaUtil {
        return MediaUtil_Instance
}


func NewMediaUtil() *MediaUtil {
    this := &MediaUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *MediaUtil) SaveImageFile(originalImageFile *AbFile, newImageFileName string, category string, mediaData *MediaData, newWidth int, newHeight int) {
        if originalImageFile == nil || !originalImageFile.IsFile() {
            panic(errors.New(fmt.Sprint("Original Image File Does Not Exist.")))
        }
        var bufferedImage *BufferedImage = ImageIOUtil_Read(originalImageFile); _ = bufferedImage
        if bufferedImage == nil {
            panic(errors.New(fmt.Sprint("Unable to find ImageReader for this file.")))
        }
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().VIEW) {
            var hashMap *HashMap = this.GetImageBufferPropertyHashMap(bufferedImage); _ = hashMap
            this.LogUtil.PutF("Image Properties: " + fmt.Sprint(hashMap), this, "saveImageFile()")
        }
        var imageFile *AbFile = AbFile_CreateAbFile(category + newImageFileName); _ = imageFile
        imageFile.CreateNewFile()
        var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
        var newBufferedImage *BufferedImage = imageUtil.CreateBufferedImageForSave(bufferedImage, newWidth, newHeight); _ = newBufferedImage
        var isWritten bool = ImageIOUtil_Write(any(newBufferedImage).(*RenderedImage), mediaData.GetName(), imageFile); _ = isWritten
        if !isWritten {
            panic(errors.New(fmt.Sprint("Unable to write.")))
        }
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().VIEW) {
            var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Get Path: ")
            stringBuffer.Append(originalImageFile.GetPath())
            stringBuffer.Append("\nNewImageFileName: ")
            stringBuffer.Append(newImageFileName)
            stringBuffer.Append("\nCategory: ")
            stringBuffer.Append(category)
            stringBuffer.Append("\nSave File Type: ")
            stringBuffer.Append(mediaData.GetName())
            stringBuffer.Append("\nNew").Append(commonLabels.WIDTH_LABEL)
            stringBuffer.Appendint(newWidth)
            stringBuffer.Append("\nNew").Append(commonLabels.HEIGHT_LABEL)
            stringBuffer.Appendint(newHeight)
            stringBuffer.Append("\nFile Length: ")
            stringBuffer.Appendlong(int64(originalImageFile.Length()))
            stringBuffer.Append("\nNew File Length: ")
            stringBuffer.Appendlong(int64(imageFile.Length()))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "saveImageFile()")
        }
}


func MediaUtil_GetImageBufferPropertyHashMap(bufferedImage *BufferedImage) *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var propertyStringArray []string = bufferedImage.GetPropertyNames(); _ = propertyStringArray
        if propertyStringArray != nil {
            for index := int(0); index < len(propertyStringArray); index++ {
                var propertyObject any = bufferedImage.GetProperty(propertyStringArray[index]); _ = propertyObject
                hashMap.Put(propertyStringArray[index], fmt.Sprint(propertyObject))
            }
        }
        return hashMap
}

