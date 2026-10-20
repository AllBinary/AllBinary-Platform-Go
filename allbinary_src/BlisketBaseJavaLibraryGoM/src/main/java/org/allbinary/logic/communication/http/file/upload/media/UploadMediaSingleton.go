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

type UploadMediaSingleton struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ReaderFileTypesHashMap *HashMap
    WriterFileTypesHashMap *HashMap
}

var UploadMediaSingleton_UploadMedia *UploadMediaSingleton

func (this *UploadMediaSingleton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewUploadMediaSingleton() *UploadMediaSingleton {
    this := &UploadMediaSingleton{}
    this.initFields()
        this.ReaderFileTypesHashMap = StdUtil_GetInstance().CreateHashMap()
        this.WriterFileTypesHashMap = StdUtil_GetInstance().CreateHashMap()
        var hashMap *HashMap = MediaData_ToHashMap(); _ = hashMap
        var set Set = hashMap.KeySet(); _ = set
        var mediaDataNameArray []any = set.ToArray(); _ = mediaDataNameArray
        var size int = len(mediaDataNameArray); _ = size
        for index := int(0); index < size; index++ {
            var mediaDataName string = any(mediaDataNameArray[index]).(string); _ = mediaDataName
            var mediaData *MediaData = any(hashMap.Get(mediaDataName)).(*MediaData); _ = mediaData
            this.ReaderFileTypesHashMap.Put(mediaData.GetName(), mediaData.GetType())
        }
        var mediaIOUtil *MediaIOUtil = MediaIOUtil_GetInstance(); _ = mediaIOUtil
        var RESIZABLE_MEDIA string = MediaTypeData_GetInstance().RESIZABLE_MEDIA; _ = RESIZABLE_MEDIA
        var readerFileTypes []string = mediaIOUtil.GetReaderFormatNames(); _ = readerFileTypes
        for index := int(0); index < len(readerFileTypes); index++ {
            this.ReaderFileTypesHashMap.Put(readerFileTypes[index], RESIZABLE_MEDIA)
        }
        var writerFileTypes []string = mediaIOUtil.GetWriterFormatNames(); _ = writerFileTypes
        for index := int(0); index < len(writerFileTypes); index++ {
            this.WriterFileTypesHashMap.Put(writerFileTypes[index], RESIZABLE_MEDIA)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Supported Media Readers: ")
            stringBuffer.Append(fmt.Sprint(this.ReaderFileTypesHashMap))
            stringBuffer.Append(" Writers: ")
            stringBuffer.Append(fmt.Sprint(this.WriterFileTypesHashMap))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// synchronized
func UploadMediaSingleton_GetInstance() *UploadMediaSingleton {
        if UploadMediaSingleton_UploadMedia == nil {
            UploadMediaSingleton_UploadMedia = NewUploadMediaSingleton()
        }
        return UploadMediaSingleton_UploadMedia
}


func (this *UploadMediaSingleton) IsReaderSupported(mediaFileType string) bool {
        return this.ReaderFileTypesHashMap.ContainsKey(mediaFileType)
}


func (this *UploadMediaSingleton) IsWriterSupported(mediaFileType string) bool {
        return this.ReaderFileTypesHashMap.ContainsKey(mediaFileType)
}


func (this *UploadMediaSingleton) IsSupported(mediaFileType string) bool {
        return this.IsReaderSupported(mediaFileType) && this.IsWriterSupported(mediaFileType)
}


func (this *UploadMediaSingleton) IsMedia_1(mediaFileType string, aMediaType string) bool {
        return this.IsReaderMedia(mediaFileType, aMediaType) && this.IsWriterMedia(mediaFileType, aMediaType)
}

// Java overloads of UploadMediaSingleton.isMedia
func (this *UploadMediaSingleton) IsMedia(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.IsMedia_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return this.IsMedia_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[*HashMap](args[2]))
        }
    }
    panic("No Java overload of UploadMediaSingleton.isMedia for the arguments")
}


func (this *UploadMediaSingleton) IsReaderMedia(mediaFileType string, aMediaType string) bool {
        return this.IsMedia(mediaFileType, aMediaType, this.ReaderFileTypesHashMap)
}


func (this *UploadMediaSingleton) IsWriterMedia(mediaFileType string, aMediaType string) bool {
        return this.IsMedia(mediaFileType, aMediaType, this.WriterFileTypesHashMap)
}


func (this *UploadMediaSingleton) IsMedia_2(mediaFileType string, aMediaType string, hashMap *HashMap) bool {
        if hashMap.ContainsKey(mediaFileType) {
            var mediaType string = any(hashMap.Get(mediaFileType)).(string); _ = mediaType
            if strings.Compare(mediaType, aMediaType) == 0 {
                return true
            } else {
                return false
            }
        } else {
            return false
        }
}

