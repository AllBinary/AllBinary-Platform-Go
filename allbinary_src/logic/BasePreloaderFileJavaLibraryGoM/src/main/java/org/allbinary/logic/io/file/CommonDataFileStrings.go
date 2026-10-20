/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type CommonDataFileStrings struct {
    XML string
    _XML string
    JSON string
    _JSON string
    XSL string
    _XSL string
    JAVA string
    _JAVA string
    KT string
    _KT string
    TS string
    _TS string
    JS string
    _JS string
    CPP string
    _CPP string
    H string
    _H string
    PY string
    _PY string
    GO string
    _GO string
    CS string
    _CS string
    ADB string
    _ADB string
    ADS string
    _ADS string
    ADA string
    _ADA string
    RUST string
    _RUST string
    C string
    _C string
    CSHARP string
    _CSHARP string
    ERL string
    _ERL string
    HS string
    _HS string
    HX string
    _HX string
    LUA string
    _LUA string
    PAS string
    _PAS string
    ML string
    _ML string
    PL string
    _PL string
    RB string
    _RB string
    ZIG string
    _ZIG string
    UNCRYPTED_EXTENSION string
    ENCRYPTED_EXTENSION string
}

var CommonDataFileStrings_Instance *CommonDataFileStrings = NewCommonDataFileStrings()

func (this *CommonDataFileStrings) initFields() {
    this.XML = "xml"
    this._XML = ".xml"
    this.JSON = "json"
    this._JSON = ".json"
    this.XSL = "xsl"
    this._XSL = ".xsl"
    this.JAVA = "java"
    this._JAVA = ".java"
    this.KT = "kt"
    this._KT = ".kt"
    this.TS = "ts"
    this._TS = ".ts"
    this.JS = "js"
    this._JS = ".js"
    this.CPP = "cpp"
    this._CPP = ".cpp"
    this.H = "h"
    this._H = ".h"
    this.PY = "py"
    this._PY = ".py"
    this.GO = "go"
    this._GO = ".go"
    this.CS = "cs"
    this._CS = ".cs"
    this.ADB = "adb"
    this._ADB = ".adb"
    this.ADS = "ads"
    this._ADS = ".ads"
    this.ADA = "ada"
    this._ADA = ".ada"
    this.RUST = "rs"
    this._RUST = ".rs"
    this.C = "c"
    this._C = ".c"
    this.CSHARP = "cs"
    this._CSHARP = ".cs"
    this.ERL = "erl"
    this._ERL = ".erl"
    this.HS = "hs"
    this._HS = ".hs"
    this.HX = "hx"
    this._HX = ".hx"
    this.LUA = "lua"
    this._LUA = ".lua"
    this.PAS = "pas"
    this._PAS = ".pas"
    this.ML = "ml"
    this._ML = ".ml"
    this.PL = "pl"
    this._PL = ".pl"
    this.RB = "rb"
    this._RB = ".rb"
    this.ZIG = "zig"
    this._ZIG = ".zig"
    this.UNCRYPTED_EXTENSION = this.XML
    this.ENCRYPTED_EXTENSION = "abd"
}

func NewCommonDataFileStrings() *CommonDataFileStrings {
    this := &CommonDataFileStrings{}
    this.initFields()
    return this
}


func CommonDataFileStrings_GetInstance() *CommonDataFileStrings {
        return CommonDataFileStrings_Instance
}

