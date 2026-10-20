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

type InventoryUploadMediaUtil struct {
    LogUtil *LogUtil
    MediaUtil *MediaUtil
    LARGEWIDTH int
    LARGEHEIGHT int
    MEDIUMWIDTH int
    MEDIUMHEIGHT int
    SMALLWIDTH int
    SMALLHEIGHT int
    SMALL string
    MEDIUM string
    LARGE string
    StoreFrontInterface StoreFrontInterface
    ItemInterface ItemInterface
}


func (this *InventoryUploadMediaUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MediaUtil = MediaUtil_GetInstance()
    this.LARGEWIDTH = 1024
    this.LARGEHEIGHT = 1024
    this.MEDIUMWIDTH = 256
    this.MEDIUMHEIGHT = 256
    this.SMALLWIDTH = 128
    this.SMALLHEIGHT = 128
    this.SMALL = "Small"
    this.MEDIUM = "Medium"
    this.LARGE = "Large"
    
    
}


func NewInventoryUploadMediaUtil(storeFrontInterface StoreFrontInterface, itemInterface ItemInterface) *InventoryUploadMediaUtil {
    this := &InventoryUploadMediaUtil{}
    this.initFields()
        this.StoreFrontInterface = storeFrontInterface
        this.ItemInterface = itemInterface
    return this
}


// throws Exception
func (this *InventoryUploadMediaUtil) SaveFiles(byteArray []int8, fileName string, mediaData *MediaData) ItemInterface {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF("Start FileName: " + fileName, this, "saveFiles()")
        }
        if StringValidationUtil_GetInstance().IsEmpty(fileName) {
            panic(errors.New(fmt.Sprint("Image File Name Was Null")))
        }
        var fileUtil *FileUtil = FileUtil_GetInstance(); _ = fileUtil
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(URLGLOBALS.GetWebappPath())
        stringBuffer.Append(this.StoreFrontInterface.GetCurrentHostNamePath())
        stringBuffer.Append(this.ItemInterface.GetCategory())
        var fullPath string = fmt.Sprint(stringBuffer); _ = fullPath
        var imageDirectoryFile *AbFile = AbFile_CreateAbFile(fullPath); _ = imageDirectoryFile
        if !imageDirectoryFile.IsDirectory() {
            panic(errors.New(fmt.Sprint("Unable to save file to non existant directory")))
        }
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(fullPath)
        stringBuffer.Append(fileName)
        stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
        stringBuffer.Append(mediaData.GetName())
        var originalImageFile *AbFile = AbFile_CreateAbFile(fmt.Sprint(stringBuffer)); _ = originalImageFile
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF("Saving Original Image: " + fmt.Sprint(originalImageFile), this, "saveFiles()")
        }
        originalImageFile.CreateNewFile()
        fileUtil.Write(NewByteArrayInputStream(byteArray), originalImageFile)
        var uploadMedia *UploadMediaSingleton = UploadMediaSingleton_GetInstance(); _ = uploadMedia
        var isMediaSupported bool = uploadMedia.IsWriterSupported(mediaData.GetName()) || uploadMedia.IsReaderSupported(mediaData.GetName()); _ = isMediaSupported
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("Is Image Type: ")
            stringBuffer.Append(mediaData.GetName())
            stringBuffer.Append(" supported: ")
            stringBuffer.Appendboolean(isMediaSupported)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "saveFiles()")
        }
        var defaultMediaData *MediaData = MediaData_GetDefault(); _ = defaultMediaData
        this.SetFileNames(fileName)
        if isMediaSupported {
            var isMediaResizable bool = uploadMedia.IsWriterMedia(mediaData.GetName(), MediaTypeData_GetInstance().RESIZABLE_MEDIA); _ = isMediaResizable
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Is ")
                stringBuffer.Append(fmt.Sprint(mediaData))
                stringBuffer.Append(" Image Resize Supported:")
                stringBuffer.Appendboolean(uploadMedia.IsSupported(mediaData.GetName()))
                stringBuffer.Append(" WriterMedia: ")
                stringBuffer.Appendboolean(isMediaResizable)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "saveFiles()")
            }
            if isMediaResizable {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                    this.LogUtil.PutF("Saving Small Image: " + fmt.Sprint(this.ItemInterface.GetSmallImage()), this, "saveFiles()")
                }
                this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetSmallImage(), fullPath, mediaData, this.SMALLWIDTH, this.SMALLHEIGHT)
                this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetMediumImage(), fullPath, mediaData, this.MEDIUMWIDTH, this.MEDIUMHEIGHT)
                this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetLargeImage(), fullPath, mediaData, this.LARGEWIDTH, this.LARGEHEIGHT)
            } else {
                var isConvertable bool = mediaData.IsConvertableTo(defaultMediaData); _ = isConvertable
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("Converting: ")
                    stringBuffer.Append(defaultMediaData.GetName())
                    stringBuffer.Append(" into ")
                    stringBuffer.Append(mediaData.GetName())
                    stringBuffer.Append(" if Convertable is it: ")
                    stringBuffer.Appendboolean(isConvertable)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "saveFiles()")
                }
                if isConvertable {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                        this.LogUtil.PutF("Saving Small Image: " + fmt.Sprint(this.ItemInterface.GetSmallImage()), this, "saveFiles()")
                    }
                    this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetSmallImage(), fullPath, defaultMediaData, this.SMALLWIDTH, this.SMALLHEIGHT)
                    this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetMediumImage(), fullPath, defaultMediaData, this.MEDIUMWIDTH, this.MEDIUMHEIGHT)
                    this.MediaUtil.SaveImageFile(originalImageFile, this.ItemInterface.GetLargeImage(), fullPath, defaultMediaData, this.LARGEWIDTH, this.LARGEHEIGHT)
                } else {
                    var smallImageFile *AbFile = AbFile_CreateAbFile(this.ItemInterface.GetSmallImage()); _ = smallImageFile
                    smallImageFile.CreateNewFile()
                    fileUtil.Write(NewByteArrayInputStream(byteArray), smallImageFile)
                    var mediumImageFile *AbFile = AbFile_CreateAbFile(this.ItemInterface.GetMediumImage()); _ = mediumImageFile
                    smallImageFile.CreateNewFile()
                    fileUtil.Write(NewByteArrayInputStream(byteArray), mediumImageFile)
                    var largeImageFile *AbFile = AbFile_CreateAbFile(this.ItemInterface.GetSmallImage()); _ = largeImageFile
                    largeImageFile.CreateNewFile()
                    fileUtil.Write(NewByteArrayInputStream(byteArray), largeImageFile)
                }
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.END, this, "saveFiles()")
        }
        return this.ItemInterface
}


func (this *InventoryUploadMediaUtil) SetFileNames(fileName string) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.ItemInterface.GetId())
        stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
        stringBuffer.Append(MediaData_GetDefault().GetName())
        var END string = fmt.Sprint(stringBuffer); _ = END
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(fileName)
        stringBuffer.Append(this.SMALL)
        stringBuffer.Append(END)
        var newImageFileName string = fmt.Sprint(stringBuffer); _ = newImageFileName
        this.ItemInterface.SetSmallImage(newImageFileName)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(fileName)
        stringBuffer.Append(this.MEDIUM)
        stringBuffer.Append(END)
        newImageFileName = fmt.Sprint(stringBuffer)
        this.ItemInterface.SetMediumImage(newImageFileName)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(fileName)
        stringBuffer.Append(this.LARGE)
        stringBuffer.Append(END)
        newImageFileName = fmt.Sprint(stringBuffer)
        this.ItemInterface.SetLargeImage(newImageFileName)
}

