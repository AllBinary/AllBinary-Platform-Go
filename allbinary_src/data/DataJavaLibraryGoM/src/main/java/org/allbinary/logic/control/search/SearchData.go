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

type SearchData struct {
}

var SearchData_TOTAL_NUMBER_PAGES string = "SEARCHDATA_TOTALNUMBEROFPAGES"
var SearchData_TOTAL_NUMBER_ITEMS_ON_THIS_PAGE string = "SEARCHDATA_TOTALNUMBEROFITEMSONTHISPAGE"
var SearchData_TOTAL_NUMBER_ITEMS string = "SEARCHDATA_TOTALNUMBEROFITEMS"
var SearchData_DEFAULT_PAGE_LENGTH int = 10
var SearchData_GENERATESTATICPAGES string = "Generate Static Pages"
var SearchData_SEARCHBYSUMMARY string = "Search By Summary"
var SearchData_SEARCHBYDESCRIPTION string = "Search By Description"
var SearchData_SEARCHBYCATEGORY string = "Search By Category"
var SearchData_SEARCHCATEGORIES string = "Search Categories"
var SearchData_FILENAME string = "SEARCHDATA_BASEFILENAME"
var SearchData_COLUMNNAME string = "COLUMN_NAME"
var SearchData_COLUMNVALUE string = "COLUMN_VALUE"
var SearchData_PARAMS string = "SEARCHDATA_PARAMS"
var SearchData_LENGTH string = "SEARCHDATA_LENGTH"
var SearchData_ORDER string = "SEARCHDATA_ORDER"
var SearchData_PAGE string = "SEARCHDATA_PAGE"
var SearchData_FIELDS string = "SEARCHDATA_FIELDS"
var SearchData_FIELD string = "SEARCHDATA_FIELD"
var SearchData_SORTBY string = "SEARCHDATA_SORTBY"
var SearchData_TYPE string = "SEARCHDATA_TYPE"
var SearchData_PAGE_INFO string = "pageInfo"
var SearchData_SEARCH string = "search"
var SearchData_MAKEPUBLIC string = "Make Public"


func NewSearchData() *SearchData {
    this := &SearchData{}
    return this
}

