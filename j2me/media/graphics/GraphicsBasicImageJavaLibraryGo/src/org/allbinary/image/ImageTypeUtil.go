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

type ImageTypeUtil struct {
    PNG string
    JPG string
    GIF string
    List IBasicArrayList
}

var ImageTypeUtil_SINGLETON *ImageTypeUtil = NewImageTypeUtil()

func (this *ImageTypeUtil) initFields() {
    this.PNG = "png"
    this.JPG = "jpg"
    this.GIF = "gif"
    this.List = NewBasicArrayListD()
}


func NewImageTypeUtil() *ImageTypeUtil {
    this := &ImageTypeUtil{}
    this.initFields()
        this.List.Add(this.JPG)
        this.List.Add(this.GIF)
        this.List.Add(this.PNG)
    return this
}


func (this *ImageTypeUtil) IsSupported(mediaFileType string) bool {
        mediaFileType = strings.ToLower(mediaFileType)
        for index := int(this.List.Size() - 1); index >= 0; index-- {
            var name string = any(this.List.BasicArrayList_().ObjectArray[index]).(string); _ = name
            if strings.Compare(mediaFileType, name) == 0 {
                return true
            }
        }
        return false
}


func ImageTypeUtil_GetInstance() *ImageTypeUtil {
        return ImageTypeUtil_SINGLETON
}

