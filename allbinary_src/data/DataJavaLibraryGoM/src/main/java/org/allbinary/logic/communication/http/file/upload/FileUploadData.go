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

type FileUploadData struct {
    FILE string
    UPLOAD_TO_FILE_PATH string
    FILE_DATA string
}

var FileUploadData_Instance *FileUploadData = NewFileUploadData()

func (this *FileUploadData) initFields() {
    this.FILE = "download_file_permissions.xml"
    this.UPLOAD_TO_FILE_PATH = "FILE_UPLOAD_UPLOAD_TO_FILE_PATH"
    this.FILE_DATA = "FILE_UPLOAD_FILE_DATA"
}

func NewFileUploadData() *FileUploadData {
    this := &FileUploadData{}
    this.initFields()
    return this
}


func FileUploadData_GetInstance() *FileUploadData {
        return FileUploadData_Instance
}

