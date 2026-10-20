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

type EchoPaymentGatewayData struct {
}

var EchoPaymentGatewayData_TRANSACTION_TYPE string = "transaction_type"
var EchoPaymentGatewayData_MERCHANT_ECHO_ID string = "merchant_echo_id"
var EchoPaymentGatewayData_ORDER_TYPE string = "order_type"
var EchoPaymentGatewayData_MERCHANT_PIN string = "merchant_pin"
var EchoPaymentGatewayData_ISP_ECHO_ID string = "isp_echo_id"
var EchoPaymentGatewayData_ISP_PIN string = "isp_pin"
var EchoPaymentGatewayData_BILLING_IP_ADDRESS string = "billing_ip_address"
var EchoPaymentGatewayData_MERCHANT_EMAIL string = "merchant_email"
var EchoPaymentGatewayData_AUTHORIZATION string = "authorization"
var EchoPaymentGatewayData_BILLING_PREFIX string = "billing_prefix"
var EchoPaymentGatewayData_BILLING_NAME string = "billing_name"
var EchoPaymentGatewayData_BILLING_FIRST_NAME string = "billing_first_name"
var EchoPaymentGatewayData_BILLING_LAST_NAME string = "billing_last_name"
var EchoPaymentGatewayData_BILLING_COMPANY_NAME string = "billing_company_name"
var EchoPaymentGatewayData_BILLING_ADDRESS1 string = "billing_address1"
var EchoPaymentGatewayData_BILLING_ADDRESS2 string = "billing_address2"
var EchoPaymentGatewayData_BILLING_CITY string = "billing_city"
var EchoPaymentGatewayData_BILLING_STATE string = "billing_state"
var EchoPaymentGatewayData_BILLING_ZIP string = "billing_zip"
var EchoPaymentGatewayData_BILLING_COUNTRY string = "billing_country"
var EchoPaymentGatewayData_BILLING_PHONE string = "billing_phone"
var EchoPaymentGatewayData_BILLING_FAX string = "billing_fax"
var EchoPaymentGatewayData_BILLING_EMAIL string = "billing_email"
var EchoPaymentGatewayData_CC_NUMBER string = "cc_number"
var EchoPaymentGatewayData_CCEXP_MONTH string = "ccexp_month"
var EchoPaymentGatewayData_CCEXP_YEAR string = "ccexp_year"
var EchoPaymentGatewayData_CNP_RECURRING string = "cnp_recurring"
var EchoPaymentGatewayData_Cnp_security string = "cnp_security"
var EchoPaymentGatewayData_COUNTER string = "counter"
var EchoPaymentGatewayData_DEBUG string = "debug"
var EchoPaymentGatewayData_EC_ACCOUNT string = "ec_account"
var EchoPaymentGatewayData_EC_ACCOUNT_TYPE string = "ec_account_type"
var EchoPaymentGatewayData_EC_ADDRESS1 string = "ec_address1"
var EchoPaymentGatewayData_EC_ADDRESS2 string = "ec_address2"
var EchoPaymentGatewayData_EC_BANK_NAME string = "ec_bank_name"
var EchoPaymentGatewayData_EC_CITY string = "ec_city"
var EchoPaymentGatewayData_EC_EMAIL string = "ec_email"
var EchoPaymentGatewayData_EC_FIRST_NAME string = "ec_first_name"
var EchoPaymentGatewayData_EC_ID_COUNTRY string = "ec_id_country"
var EchoPaymentGatewayData_EC_ID_EXP_MM string = "ec_id_exp_mm"
var EchoPaymentGatewayData_EC_ID_EXP_DD string = "ec_id_exp_dd"
var EchoPaymentGatewayData_EC_ID_EXP_YY string = "ec_id_exp_yy"
var EchoPaymentGatewayData_EC_ID_NUMBER string = "ec_id_number"
var EchoPaymentGatewayData_EC_ID_STATE string = "ec_id_state"
var EchoPaymentGatewayData_EC_ID_TYPE string = "ec_id_type"
var EchoPaymentGatewayData_EC_LAST_NAME string = "ec_last_name"
var EchoPaymentGatewayData_EC_BUSINESS_NAME string = "ec_BUSINESS_NAME"
var EchoPaymentGatewayData_EC_LICENSE_NUMBER string = "ec_LICENSE_NUMBER"
var EchoPaymentGatewayData_EC_LICENSE_STATE string = "ec_LICENSE_STATE"
var EchoPaymentGatewayData_EC_MERCHANT_REF string = "ec_MERCHANT_REF"
var EchoPaymentGatewayData_EC_NBDS_CODE string = "ec_NBDS_CODE"
var EchoPaymentGatewayData_EC_OTHER_NAME string = "ec_other_name"
var EchoPaymentGatewayData_EC_PAYEE string = "ec_payee"
var EchoPaymentGatewayData_EC_PAYMENT_TYPE string = "ec_payment_type"
var EchoPaymentGatewayData_EC_RT string = "ec_rt"
var EchoPaymentGatewayData_EC_SERIAL_NUMBER string = "ec_serial_number"
var EchoPaymentGatewayData_EC_STATE string = "ec_state"
var EchoPaymentGatewayData_EC_TRANSACTION_DT string = "ec_transaction_dt"
var EchoPaymentGatewayData_EC_ZIP string = "ec_zip"
var EchoPaymentGatewayData_GRAND_TOTAL string = "grand_total"
var EchoPaymentGatewayData_MERCHANT_TRACE_NBR string = "merchant_trace_nbr"
var EchoPaymentGatewayData_ORDER_NUMBER string = "order_number"
var EchoPaymentGatewayData_ORIGINAL_AMOUNT string = "original_amount"
var EchoPaymentGatewayData_ORIGINAL_TRANDATE_MM string = "original_trandate_mm"
var EchoPaymentGatewayData_ORIGINAL_TRANDATE_DD string = "original_trandate_dd"
var EchoPaymentGatewayData_ORIGINAL_TRANDATE_YYYY string = "original_trandate_yyyy"
var EchoPaymentGatewayData_ORIGINAL_REFERENCE string = "original_reference"
var EchoPaymentGatewayData_PRODUCT_DESCRIPTION string = "product_description"
var EchoPaymentGatewayData_PURCHASE_ORDER_NUMBER string = "purchase_order_number"
var EchoPaymentGatewayData_SALES_TAX string = "sales_tax"


func NewEchoPaymentGatewayData() *EchoPaymentGatewayData {
    this := &EchoPaymentGatewayData{}
    return this
}

