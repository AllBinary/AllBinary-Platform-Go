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

type FileData struct {
    NAME string
    ISFILE string
    ISDIRECTORY string
    ISHIDDEN string
    ISABSOLUTE string
    ROOT_NAME string
    PATH string
    ABSOLUTE_PATH string
    CANONICAL_PATH string
    PARENT string
    MINLEN int
    MAXLEN int
    MAXIMAGEFILESIZE int
    MINIMAGEFILESIZE int
    MAXDOWNLOADABLEFILESIZE int
    MINDOWNLOADABLEFILESIZE int
}

var FileData_Instance *FileData = NewFileData()

func (this *FileData) initFields() {
    this.NAME = "FILE_NAME"
    this.ISFILE = "FILE_ISFILE"
    this.ISDIRECTORY = "FILE_ISDIRECTORY"
    this.ISHIDDEN = "FILE_ISHIDDEN"
    this.ISABSOLUTE = "FILE_ISABSOLUTE"
    this.ROOT_NAME = "FILE_ROOT_NAME"
    this.PATH = "FILE_PATH"
    this.ABSOLUTE_PATH = "FILE_ABSOLUTE_PATH"
    this.CANONICAL_PATH = "FILE_CANONICAL_PATH"
    this.PARENT = "FILE_PARENT"
    this.MINLEN = 1
    this.MAXLEN = 120
    this.MAXIMAGEFILESIZE = 1048576
    this.MINIMAGEFILESIZE = 32
    this.MAXDOWNLOADABLEFILESIZE = 1048576 * 20
    this.MINDOWNLOADABLEFILESIZE = 1
}


func FileData_GetInstance() *FileData {
        return FileData_Instance
}


func NewFileData() *FileData {
    this := &FileData{}
    this.initFields()
    return this
}

