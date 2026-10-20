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

type CommonSeps struct {
    NEW_LINE string
    SEMICOLON string
    EQUALS string
    SPACE string
    COMMA string
    COLON string
    PERIOD string
    UNDERSCORE string
    COLON_SEP string
    COMMA_SEP string
    PARENTHESIS_OPEN string
    PARENTHESIS_CLOSE string
    BRACE_OPEN string
    BRACE_CLOSE string
    BRACKET_OPEN string
    BRACKET_CLOSE string
    DASH string
    PERCENT string
    AMPERSAND string
    FORWARD_SLASH string
    BACK_SLASH string
    QUESTION string
    QUOTE string
    QUOTE_END string
    COMMENT string
    INDENT string
}

var CommonSeps_Instance *CommonSeps = NewCommonSeps()

func (this *CommonSeps) initFields() {
    this.NEW_LINE = "\n"
    this.SEMICOLON = ";"
    this.EQUALS = "="
    this.SPACE = " "
    this.COMMA = ","
    this.COLON = ":"
    this.PERIOD = "."
    this.UNDERSCORE = "_"
    this.COLON_SEP = fmt.Sprint(NewStringMaker().Append(this.COLON).Append(this.SPACE))
    this.COMMA_SEP = fmt.Sprint(NewStringMaker().Append(this.COMMA).Append(this.SPACE))
    this.PARENTHESIS_OPEN = "("
    this.PARENTHESIS_CLOSE = ")"
    this.BRACE_OPEN = "{"
    this.BRACE_CLOSE = "}"
    this.BRACKET_OPEN = "["
    this.BRACKET_CLOSE = "]"
    this.DASH = "-"
    this.PERCENT = "%"
    this.AMPERSAND = "&"
    this.FORWARD_SLASH = "/"
    this.BACK_SLASH = "\\"
    this.QUESTION = "?"
    this.QUOTE = "\""
    this.QUOTE_END = "\","
    this.COMMENT = "//"
    this.INDENT = "    "
}

func NewCommonSeps() *CommonSeps {
    this := &CommonSeps{}
    this.initFields()
    return this
}


func CommonSeps_GetInstance() *CommonSeps {
        return CommonSeps_Instance
}
