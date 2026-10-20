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

type QuoteRequestData struct {
    ID string
    PROJECT_INFO string
    CUSTOMER_COMMENTS string
    BUDGET string
    TIMEFRAME string
    COMMENTS string
}

var QuoteRequestData_Instance *QuoteRequestData = NewQuoteRequestData()

func (this *QuoteRequestData) initFields() {
    this.ID = "QUOTEREQUEST_ID"
    this.PROJECT_INFO = "QUOTEREQUEST_PROJECT_INFO"
    this.CUSTOMER_COMMENTS = "QUOTEREQUEST_CUSTOMER_COMMENTS"
    this.BUDGET = "QUOTEREQUEST_BUDGET"
    this.TIMEFRAME = "QUOTEREQUEST_TIMEFRAME"
    this.COMMENTS = "QUOTEREQUEST_COMMENTS"
}


func QuoteRequestData_GetInstance() *QuoteRequestData {
        return QuoteRequestData_Instance
}


func NewQuoteRequestData() *QuoteRequestData {
    this := &QuoteRequestData{}
    this.initFields()
    return this
}

