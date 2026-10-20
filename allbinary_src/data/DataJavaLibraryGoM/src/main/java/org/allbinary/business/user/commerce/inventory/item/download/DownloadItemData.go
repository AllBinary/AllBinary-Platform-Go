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

type DownloadItemData struct {
}

var DownloadItemData_NAME string = "DOWNLOADITEM_NAME"
var DownloadItemData_ID string = "DOWNLOADITEM_ID"
var DownloadItemData_SYSTEM string = "DOWNLOADITEM_SYSTEM"
var DownloadItemData_PLATFORM string = "DOWNLOADITEM_PLATFORM"
var DownloadItemData_SPECIAL_NAME string = "DOWNLOADITEM_SPECIAL_NAME"
var DownloadItemData_VERSION string = "DOWNLOADITEM_VERSION"
var DownloadItemData_CHANGES string = "DOWNLOADITEM_CHANGES"
var DownloadItemData_LICENSE_FILE string = "DOWNLOADITEM_LICENSE_FILE"
var DownloadItemData_FILE string = "DOWNLOADITEM_FILE"
var DownloadItemData_EXTRA_FILES string = "DOWNLOADITEM_EXTRA_FILES"
var DownloadItemData_SIZE string = "DOWNLOADITEM_SIZE"
var DownloadItemData_VALID_TIME string = "DOWNLOADITEM_VALID_TIME"
var DownloadItemData_VALID_TIME_YEARS string = "DOWNLOADITEM_VALID_TIME_YEAR"
var DownloadItemData_VALID_TIME_MONTHS string = "DOWNLOADITEM_VALID_TIME_MONTH"
var DownloadItemData_VALID_TIME_DAYS string = "DOWNLOADITEM_VALID_TIME_DAY"
var DownloadItemData_VALID_TIME_HOURS string = "DOWNLOADITEM_VALID_TIME_HOUR"
var DownloadItemData_VALID_TIME_MINUTES string = "DOWNLOADITEM_VALID_TIME_MINUTES"
var DownloadItemData_VALID_TIME_SECONDS string = "DOWNLOADITEM_VALID_TIME_SECONDS"
var DownloadItemData_RETRIES string = "DOWNLOADITEM_RETRIES"

func NewDownloadItemData() *DownloadItemData {
    this := &DownloadItemData{}
    return this
}

