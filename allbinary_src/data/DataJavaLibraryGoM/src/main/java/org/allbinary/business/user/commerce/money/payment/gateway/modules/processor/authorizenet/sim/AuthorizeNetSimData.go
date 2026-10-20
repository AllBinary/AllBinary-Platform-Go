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

type AuthorizeNetSimData struct {
}

var AuthorizeNetSimData_X_FP_HASH string = "x_fp_hash"
var AuthorizeNetSimData_X_FP_SEQUENCE string = "x_fp_sequence"
var AuthorizeNetSimData_X_FP_TIMESTAMP string = "x_fp_timestamp"
var AuthorizeNetSimData_X_SHOW_FORM string = "x_show_form"
var AuthorizeNetSimData_X_HEADER_HTML_PAYMENT_FORM string = "x_header_html_payment_form"
var AuthorizeNetSimData_X_FOOTER_HTML_PAYMENT_FORM string = "x_footer_html_payment_form"
var AuthorizeNetSimData_X_HEADER_HTML_RECEIPT string = "x_header_html_receipt"
var AuthorizeNetSimData_X_FOOTER_HTML_RECEIPT string = "x_footer_html_receipt"
var AuthorizeNetSimData_X_RECEIPT_LINK_TEXT string = "x_receipt_link_text"
var AuthorizeNetSimData_X_RECEIPT_LINK_URL string = "x_receipt_link_url"
var AuthorizeNetSimData_X_LOGO_URL string = "x_logo_url"
var AuthorizeNetSimData_X_BACKGROUND_URL string = "x_background_url"
var AuthorizeNetSimData_X_COLOR_BACKGROUND string = "x_color_background"
var AuthorizeNetSimData_X_COLOR_LINK string = "x_color_link"
var AuthorizeNetSimData_X_COLOR_TEXT string = "x_color_text"
var AuthorizeNetSimData_X_RELAY_URL string = "x_relay_url"


func NewAuthorizeNetSimData() *AuthorizeNetSimData {
    this := &AuthorizeNetSimData{}
    return this
}

