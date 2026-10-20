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

type CategoryData struct {
    PARENT string
    NAME string
    PATH string
    LABEL string
    VIEW string
    LEVEL string
    REQUEST string
    ROOTCATEGORY string
    UNCRYPTED_EXTENSION string
    ENCRYPTED_EXTENSION string
}

var CategoryData_Instance *CategoryData = NewCategoryData()

func (this *CategoryData) initFields() {
    this.PARENT = "CATEGORY_PARENT"
    this.NAME = "CATEGORY_NAME"
    this.PATH = "CATEGORY_PATH"
    this.LABEL = "CATEGORY_LABEL"
    this.VIEW = "CATEGORY_VIEW_CATEGORY"
    this.LEVEL = "CATEGORY_LEVEL"
    this.REQUEST = "request"
    this.ROOTCATEGORY = "Category"
    this.UNCRYPTED_EXTENSION = CommonDataFileStrings_GetInstance().UNCRYPTED_EXTENSION
    this.ENCRYPTED_EXTENSION = CommonDataFileStrings_GetInstance().ENCRYPTED_EXTENSION
}


func CategoryData_GetInstance() *CategoryData {
        return CategoryData_Instance
}


func NewCategoryData() *CategoryData {
    this := &CategoryData{}
    this.initFields()
    return this
}

