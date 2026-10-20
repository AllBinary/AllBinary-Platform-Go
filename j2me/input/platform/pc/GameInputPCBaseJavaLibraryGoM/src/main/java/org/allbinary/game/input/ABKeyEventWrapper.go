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

type ABKeyEventWrapper struct {
    VK_ENTER int
    VK_BACK_SPACE int
    VK_TAB int
    VK_CANCEL int
    VK_CLEAR int
    VK_SHIFT int
    VK_CONTROL int
    VK_ALT int
    VK_PAUSE int
    VK_CAPS_LOCK int
    VK_ESCAPE int
    VK_SPACE int
    VK_PAGE_UP int
    VK_PAGE_DOWN int
    VK_END int
    VK_HOME int
    VK_LEFT int
    VK_UP int
    VK_RIGHT int
    VK_DOWN int
    VK_COMMA int
    VK_MINUS int
    VK_PERIOD int
    VK_SLASH int
    VK_0 int
    VK_1 int
    VK_2 int
    VK_3 int
    VK_4 int
    VK_5 int
    VK_6 int
    VK_7 int
    VK_8 int
    VK_9 int
    VK_SEMICOLON int
    VK_EQUALS int
    VK_A int
    VK_B int
    VK_C int
    VK_D int
    VK_E int
    VK_F int
    VK_G int
    VK_H int
    VK_I int
    VK_J int
    VK_K int
    VK_L int
    VK_M int
    VK_N int
    VK_O int
    VK_P int
    VK_Q int
    VK_R int
    VK_S int
    VK_T int
    VK_U int
    VK_V int
    VK_W int
    VK_X int
    VK_Y int
    VK_Z int
    VK_OPEN_BRACKET int
    VK_BACK_SLASH int
    VK_CLOSE_BRACKET int
    VK_NUMPAD0 int
    VK_NUMPAD1 int
    VK_NUMPAD2 int
    VK_NUMPAD3 int
    VK_NUMPAD4 int
    VK_NUMPAD5 int
    VK_NUMPAD6 int
    VK_NUMPAD7 int
    VK_NUMPAD8 int
    VK_NUMPAD9 int
    VK_MULTIPLY int
    VK_ADD int
    VK_SEPARATER int
    VK_SEPARATOR int
    VK_SUBTRACT int
    VK_DECIMAL int
    VK_DIVIDE int
    VK_DELETE int
    VK_NUM_LOCK int
    VK_SCROLL_LOCK int
    VK_F1 int
    VK_F2 int
    VK_F3 int
    VK_F4 int
    VK_F5 int
    VK_F6 int
    VK_F7 int
    VK_F8 int
    VK_F9 int
    VK_F10 int
    VK_F11 int
    VK_F12 int
    VK_F13 int
    VK_F14 int
    VK_F15 int
    VK_F16 int
    VK_F17 int
    VK_F18 int
    VK_F19 int
    VK_F20 int
    VK_F21 int
    VK_F22 int
    VK_F23 int
    VK_F24 int
    VK_PRINTSCREEN int
    VK_INSERT int
    VK_HELP int
    VK_META int
    VK_BACK_QUOTE int
    VK_QUOTE int
    VK_KP_UP int
    VK_KP_DOWN int
    VK_KP_LEFT int
    VK_KP_RIGHT int
    VK_DEAD_GRAVE int
    VK_DEAD_ACUTE int
    VK_DEAD_CIRCUMFLEX int
    VK_DEAD_TILDE int
    VK_DEAD_MACRON int
    VK_DEAD_BREVE int
    VK_DEAD_ABOVEDOT int
    VK_DEAD_DIAERESIS int
    VK_DEAD_ABOVERING int
    VK_DEAD_DOUBLEACUTE int
    VK_DEAD_CARON int
    VK_DEAD_CEDILLA int
    VK_DEAD_OGONEK int
    VK_DEAD_IOTA int
    VK_DEAD_VOICED_SOUND int
    VK_DEAD_SEMIVOICED_SOUND int
    VK_AMPERSAND int
    VK_ASTERISK int
    VK_QUOTEDBL int
    VK_LESS int
    VK_GREATER int
    VK_BRACELEFT int
    VK_BRACERIGHT int
    VK_AT int
    VK_COLON int
    VK_CIRCUMFLEX int
    VK_DOLLAR int
    VK_EURO_SIGN int
    VK_EXCLAMATION_MARK int
    VK_INVERTED_EXCLAMATION_MARK int
    VK_LEFT_PARENTHESIS int
    VK_NUMBER_SIGN int
    VK_PLUS int
    VK_RIGHT_PARENTHESIS int
    VK_UNDERSCORE int
    VK_WINDOWS int
    VK_CONTEXT_MENU int
    VK_FINAL int
    VK_CONVERT int
    VK_NONCONVERT int
    VK_ACCEPT int
    VK_MODECHANGE int
    VK_KANA int
    VK_KANJI int
    VK_ALPHANUMERIC int
    VK_KATAKANA int
    VK_HIRAGANA int
    VK_FULL_WIDTH int
    VK_HALF_WIDTH int
    VK_ROMAN_CHARACTERS int
    VK_ALL_CANDIDATES int
    VK_PREVIOUS_CANDIDATE int
    VK_CODE_INPUT int
    VK_JAPANESE_KATAKANA int
    VK_JAPANESE_HIRAGANA int
    VK_JAPANESE_ROMAN int
    VK_KANA_LOCK int
    VK_INPUT_METHOD_ON_OFF int
    VK_CUT int
    VK_COPY int
    VK_PASTE int
    VK_UNDO int
    VK_AGAIN int
    VK_FIND int
    VK_PROPS int
    VK_STOP int
    VK_COMPOSE int
    VK_ALT_GRAPH int
    VK_BEGIN int
    VK_UNDEFINED int
}

var ABKeyEventWrapper_Instance *ABKeyEventWrapper = NewABKeyEventWrapper()

func (this *ABKeyEventWrapper) initFields() {
    this.VK_ENTER = KeyEvent_VK_ENTER
    this.VK_BACK_SPACE = KeyEvent_VK_BACK_SPACE
    this.VK_TAB = KeyEvent_VK_TAB
    this.VK_CANCEL = KeyEvent_VK_CANCEL
    this.VK_CLEAR = KeyEvent_VK_CLEAR
    this.VK_SHIFT = KeyEvent_VK_SHIFT
    this.VK_CONTROL = KeyEvent_VK_CONTROL
    this.VK_ALT = KeyEvent_VK_ALT
    this.VK_PAUSE = KeyEvent_VK_PAUSE
    this.VK_CAPS_LOCK = KeyEvent_VK_CAPS_LOCK
    this.VK_ESCAPE = KeyEvent_VK_ESCAPE
    this.VK_SPACE = KeyEvent_VK_SPACE
    this.VK_PAGE_UP = KeyEvent_VK_PAGE_UP
    this.VK_PAGE_DOWN = KeyEvent_VK_PAGE_DOWN
    this.VK_END = KeyEvent_VK_END
    this.VK_HOME = KeyEvent_VK_HOME
    this.VK_LEFT = KeyEvent_VK_LEFT
    this.VK_UP = KeyEvent_VK_UP
    this.VK_RIGHT = KeyEvent_VK_RIGHT
    this.VK_DOWN = KeyEvent_VK_DOWN
    this.VK_COMMA = KeyEvent_VK_COMMA
    this.VK_MINUS = KeyEvent_VK_MINUS
    this.VK_PERIOD = KeyEvent_VK_PERIOD
    this.VK_SLASH = KeyEvent_VK_SLASH
    this.VK_0 = KeyEvent_VK_0
    this.VK_1 = KeyEvent_VK_1
    this.VK_2 = KeyEvent_VK_2
    this.VK_3 = KeyEvent_VK_3
    this.VK_4 = KeyEvent_VK_4
    this.VK_5 = KeyEvent_VK_5
    this.VK_6 = KeyEvent_VK_6
    this.VK_7 = KeyEvent_VK_7
    this.VK_8 = KeyEvent_VK_8
    this.VK_9 = KeyEvent_VK_9
    this.VK_SEMICOLON = KeyEvent_VK_SEMICOLON
    this.VK_EQUALS = KeyEvent_VK_EQUALS
    this.VK_A = KeyEvent_VK_A
    this.VK_B = KeyEvent_VK_B
    this.VK_C = KeyEvent_VK_C
    this.VK_D = KeyEvent_VK_D
    this.VK_E = KeyEvent_VK_E
    this.VK_F = KeyEvent_VK_F
    this.VK_G = KeyEvent_VK_G
    this.VK_H = KeyEvent_VK_H
    this.VK_I = KeyEvent_VK_I
    this.VK_J = KeyEvent_VK_J
    this.VK_K = KeyEvent_VK_K
    this.VK_L = KeyEvent_VK_L
    this.VK_M = KeyEvent_VK_M
    this.VK_N = KeyEvent_VK_N
    this.VK_O = KeyEvent_VK_O
    this.VK_P = KeyEvent_VK_P
    this.VK_Q = KeyEvent_VK_Q
    this.VK_R = KeyEvent_VK_R
    this.VK_S = KeyEvent_VK_S
    this.VK_T = KeyEvent_VK_T
    this.VK_U = KeyEvent_VK_U
    this.VK_V = KeyEvent_VK_V
    this.VK_W = KeyEvent_VK_W
    this.VK_X = KeyEvent_VK_X
    this.VK_Y = KeyEvent_VK_Y
    this.VK_Z = KeyEvent_VK_Z
    this.VK_OPEN_BRACKET = KeyEvent_VK_OPEN_BRACKET
    this.VK_BACK_SLASH = KeyEvent_VK_BACK_SLASH
    this.VK_CLOSE_BRACKET = KeyEvent_VK_CLOSE_BRACKET
    this.VK_NUMPAD0 = KeyEvent_VK_NUMPAD0
    this.VK_NUMPAD1 = KeyEvent_VK_NUMPAD1
    this.VK_NUMPAD2 = KeyEvent_VK_NUMPAD2
    this.VK_NUMPAD3 = KeyEvent_VK_NUMPAD3
    this.VK_NUMPAD4 = KeyEvent_VK_NUMPAD4
    this.VK_NUMPAD5 = KeyEvent_VK_NUMPAD5
    this.VK_NUMPAD6 = KeyEvent_VK_NUMPAD6
    this.VK_NUMPAD7 = KeyEvent_VK_NUMPAD7
    this.VK_NUMPAD8 = KeyEvent_VK_NUMPAD8
    this.VK_NUMPAD9 = KeyEvent_VK_NUMPAD9
    this.VK_MULTIPLY = KeyEvent_VK_MULTIPLY
    this.VK_ADD = KeyEvent_VK_ADD
    this.VK_SEPARATER = KeyEvent_VK_SEPARATER
    this.VK_SEPARATOR = KeyEvent_VK_SEPARATOR
    this.VK_SUBTRACT = KeyEvent_VK_SUBTRACT
    this.VK_DECIMAL = KeyEvent_VK_DECIMAL
    this.VK_DIVIDE = KeyEvent_VK_DIVIDE
    this.VK_DELETE = KeyEvent_VK_DELETE
    this.VK_NUM_LOCK = KeyEvent_VK_NUM_LOCK
    this.VK_SCROLL_LOCK = KeyEvent_VK_SCROLL_LOCK
    this.VK_F1 = KeyEvent_VK_F1
    this.VK_F2 = KeyEvent_VK_F2
    this.VK_F3 = KeyEvent_VK_F3
    this.VK_F4 = KeyEvent_VK_F4
    this.VK_F5 = KeyEvent_VK_F5
    this.VK_F6 = KeyEvent_VK_F6
    this.VK_F7 = KeyEvent_VK_F7
    this.VK_F8 = KeyEvent_VK_F8
    this.VK_F9 = KeyEvent_VK_F9
    this.VK_F10 = KeyEvent_VK_F10
    this.VK_F11 = KeyEvent_VK_F11
    this.VK_F12 = KeyEvent_VK_F12
    this.VK_F13 = KeyEvent_VK_F13
    this.VK_F14 = KeyEvent_VK_F14
    this.VK_F15 = KeyEvent_VK_F15
    this.VK_F16 = KeyEvent_VK_F16
    this.VK_F17 = KeyEvent_VK_F17
    this.VK_F18 = KeyEvent_VK_F18
    this.VK_F19 = KeyEvent_VK_F19
    this.VK_F20 = KeyEvent_VK_F20
    this.VK_F21 = KeyEvent_VK_F21
    this.VK_F22 = KeyEvent_VK_F22
    this.VK_F23 = KeyEvent_VK_F23
    this.VK_F24 = KeyEvent_VK_F24
    this.VK_PRINTSCREEN = KeyEvent_VK_PRINTSCREEN
    this.VK_INSERT = KeyEvent_VK_INSERT
    this.VK_HELP = KeyEvent_VK_HELP
    this.VK_META = KeyEvent_VK_META
    this.VK_BACK_QUOTE = KeyEvent_VK_BACK_QUOTE
    this.VK_QUOTE = KeyEvent_VK_QUOTE
    this.VK_KP_UP = KeyEvent_VK_KP_UP
    this.VK_KP_DOWN = KeyEvent_VK_KP_DOWN
    this.VK_KP_LEFT = KeyEvent_VK_KP_LEFT
    this.VK_KP_RIGHT = KeyEvent_VK_KP_RIGHT
    this.VK_DEAD_GRAVE = KeyEvent_VK_DEAD_GRAVE
    this.VK_DEAD_ACUTE = KeyEvent_VK_DEAD_ACUTE
    this.VK_DEAD_CIRCUMFLEX = KeyEvent_VK_DEAD_CIRCUMFLEX
    this.VK_DEAD_TILDE = KeyEvent_VK_DEAD_TILDE
    this.VK_DEAD_MACRON = KeyEvent_VK_DEAD_MACRON
    this.VK_DEAD_BREVE = KeyEvent_VK_DEAD_BREVE
    this.VK_DEAD_ABOVEDOT = KeyEvent_VK_DEAD_ABOVEDOT
    this.VK_DEAD_DIAERESIS = KeyEvent_VK_DEAD_DIAERESIS
    this.VK_DEAD_ABOVERING = KeyEvent_VK_DEAD_ABOVERING
    this.VK_DEAD_DOUBLEACUTE = KeyEvent_VK_DEAD_DOUBLEACUTE
    this.VK_DEAD_CARON = KeyEvent_VK_DEAD_CARON
    this.VK_DEAD_CEDILLA = KeyEvent_VK_DEAD_CEDILLA
    this.VK_DEAD_OGONEK = KeyEvent_VK_DEAD_OGONEK
    this.VK_DEAD_IOTA = KeyEvent_VK_DEAD_IOTA
    this.VK_DEAD_VOICED_SOUND = KeyEvent_VK_DEAD_VOICED_SOUND
    this.VK_DEAD_SEMIVOICED_SOUND = KeyEvent_VK_DEAD_SEMIVOICED_SOUND
    this.VK_AMPERSAND = KeyEvent_VK_AMPERSAND
    this.VK_ASTERISK = KeyEvent_VK_ASTERISK
    this.VK_QUOTEDBL = KeyEvent_VK_QUOTEDBL
    this.VK_LESS = KeyEvent_VK_LESS
    this.VK_GREATER = KeyEvent_VK_GREATER
    this.VK_BRACELEFT = KeyEvent_VK_BRACELEFT
    this.VK_BRACERIGHT = KeyEvent_VK_BRACERIGHT
    this.VK_AT = KeyEvent_VK_AT
    this.VK_COLON = KeyEvent_VK_COLON
    this.VK_CIRCUMFLEX = KeyEvent_VK_CIRCUMFLEX
    this.VK_DOLLAR = KeyEvent_VK_DOLLAR
    this.VK_EURO_SIGN = KeyEvent_VK_EURO_SIGN
    this.VK_EXCLAMATION_MARK = KeyEvent_VK_EXCLAMATION_MARK
    this.VK_INVERTED_EXCLAMATION_MARK = KeyEvent_VK_INVERTED_EXCLAMATION_MARK
    this.VK_LEFT_PARENTHESIS = KeyEvent_VK_LEFT_PARENTHESIS
    this.VK_NUMBER_SIGN = KeyEvent_VK_NUMBER_SIGN
    this.VK_PLUS = KeyEvent_VK_PLUS
    this.VK_RIGHT_PARENTHESIS = KeyEvent_VK_RIGHT_PARENTHESIS
    this.VK_UNDERSCORE = KeyEvent_VK_UNDERSCORE
    this.VK_WINDOWS = KeyEvent_VK_WINDOWS
    this.VK_CONTEXT_MENU = KeyEvent_VK_CONTEXT_MENU
    this.VK_FINAL = KeyEvent_VK_FINAL
    this.VK_CONVERT = KeyEvent_VK_CONVERT
    this.VK_NONCONVERT = KeyEvent_VK_NONCONVERT
    this.VK_ACCEPT = KeyEvent_VK_ACCEPT
    this.VK_MODECHANGE = KeyEvent_VK_MODECHANGE
    this.VK_KANA = KeyEvent_VK_KANA
    this.VK_KANJI = KeyEvent_VK_KANJI
    this.VK_ALPHANUMERIC = KeyEvent_VK_ALPHANUMERIC
    this.VK_KATAKANA = KeyEvent_VK_KATAKANA
    this.VK_HIRAGANA = KeyEvent_VK_HIRAGANA
    this.VK_FULL_WIDTH = KeyEvent_VK_FULL_WIDTH
    this.VK_HALF_WIDTH = KeyEvent_VK_HALF_WIDTH
    this.VK_ROMAN_CHARACTERS = KeyEvent_VK_ROMAN_CHARACTERS
    this.VK_ALL_CANDIDATES = KeyEvent_VK_ALL_CANDIDATES
    this.VK_PREVIOUS_CANDIDATE = KeyEvent_VK_PREVIOUS_CANDIDATE
    this.VK_CODE_INPUT = KeyEvent_VK_CODE_INPUT
    this.VK_JAPANESE_KATAKANA = KeyEvent_VK_JAPANESE_KATAKANA
    this.VK_JAPANESE_HIRAGANA = KeyEvent_VK_JAPANESE_HIRAGANA
    this.VK_JAPANESE_ROMAN = KeyEvent_VK_JAPANESE_ROMAN
    this.VK_KANA_LOCK = KeyEvent_VK_KANA_LOCK
    this.VK_INPUT_METHOD_ON_OFF = KeyEvent_VK_INPUT_METHOD_ON_OFF
    this.VK_CUT = KeyEvent_VK_CUT
    this.VK_COPY = KeyEvent_VK_COPY
    this.VK_PASTE = KeyEvent_VK_PASTE
    this.VK_UNDO = KeyEvent_VK_UNDO
    this.VK_AGAIN = KeyEvent_VK_AGAIN
    this.VK_FIND = KeyEvent_VK_FIND
    this.VK_PROPS = KeyEvent_VK_PROPS
    this.VK_STOP = KeyEvent_VK_STOP
    this.VK_COMPOSE = KeyEvent_VK_COMPOSE
    this.VK_ALT_GRAPH = KeyEvent_VK_ALT_GRAPH
    this.VK_BEGIN = KeyEvent_VK_BEGIN
    this.VK_UNDEFINED = KeyEvent_VK_UNDEFINED
}

func NewABKeyEventWrapper() *ABKeyEventWrapper {
    this := &ABKeyEventWrapper{}
    this.initFields()
    return this
}


func ABKeyEventWrapper_GetInstance() *ABKeyEventWrapper {
        return ABKeyEventWrapper_Instance
}

