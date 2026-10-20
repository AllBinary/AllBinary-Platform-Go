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

type BasicItemData struct {
    ITEM string
    TOTAL string
    ID string
    NUMBER string
    INBASKETS string
    WEIGHT string
    NEWORUSED string
    SUMMARY string
    DISTRIBUTOR string
    IDUSEDBYDISTRIBUTOR string
    PRODUCEDBY string
    PRODUCTIONDATE string
    STARTPRODUCTIONDATE string
    DESCRIPTION string
    KEYWORDS string
    CATEGORY string
    TYPE string
    IMAGE string
    SMALLIMAGE string
    MEDIUMIMAGE string
    LARGEIMAGE string
    PRICE string
    COMMENT string
    CUSTOMS string
    DOWNLOADS string
    GROUPS string
    OPTIONS string
    PERMISSIONS string
    SPECIALS string
}

var BasicItemData_Instance *BasicItemData = NewBasicItemData()

func (this *BasicItemData) initFields() {
    this.ITEM = "BASICITEM"
    this.TOTAL = "BASICITEM_TOTAL"
    this.ID = "BASICITEM_ID"
    this.NUMBER = "BASICITEM_NUMBER"
    this.INBASKETS = "BASICITEM_IN_BASKETS"
    this.WEIGHT = "BASICITEM_WEIGHT"
    this.NEWORUSED = "BASICITEM_NEW_OR_USED"
    this.SUMMARY = "BASICITEM_SUMMARY"
    this.DISTRIBUTOR = "BASICITEM_DISTRIBUTOR"
    this.IDUSEDBYDISTRIBUTOR = "BASICITEM_IDUSEDBYDISTRIBUTOR"
    this.PRODUCEDBY = "BASICITEM_PRODUCEDBY"
    this.PRODUCTIONDATE = "BASICITEM_PRODUCTIONDATE"
    this.STARTPRODUCTIONDATE = "BASICITEM_STARTPRODUCTIONDATE"
    this.DESCRIPTION = "BASICITEM_DESCRIPTION"
    this.KEYWORDS = "BASICITEM_KEYWORDS"
    this.CATEGORY = "BASICITEM_CATEGORY"
    this.TYPE = "BASICITEM_TYPE"
    this.IMAGE = "BASICITEM_IMG"
    this.SMALLIMAGE = "BASICITEM_SMALL_IMG"
    this.MEDIUMIMAGE = "BASICITEM_MEDIUM_IMG"
    this.LARGEIMAGE = "BASICITEM_LARGE_IMG"
    this.PRICE = "BASICITEM_PRICE"
    this.COMMENT = "BASICITEM_COMMENT"
    this.CUSTOMS = "BASICITEM_CUSTOMS"
    this.DOWNLOADS = "BASICITEM_DOWNLOADS"
    this.GROUPS = "BASICITEM_GROUPS"
    this.OPTIONS = "BASICITEM_OPTIONS"
    this.PERMISSIONS = "BASICITEM_PERMISSIONS"
    this.SPECIALS = "BASICITEM_SPECIALS"
}

func NewBasicItemData() *BasicItemData {
    this := &BasicItemData{}
    this.initFields()
    return this
}


func BasicItemData_GetInstance() *BasicItemData {
        return BasicItemData_Instance
}

