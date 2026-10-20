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

type AndroidKeyFactory struct {
    ActivityFractureUtil
    LogUtil *LogUtil
    UNKNOWN *AndroidGameKey
    KEYCODE_BUTTON_L1 *AndroidGameKey
    KEYCODE_BUTTON_R1 *AndroidGameKey
    KEYCODE_BUTTON_THUMBR *AndroidGameKey
    KEYCODE_BUTTON_THUMBL *AndroidGameKey
    KEYCODE_BUTTON_START *AndroidGameKey
    KEYCODE_BUTTON_MODE *AndroidGameKey
    KEYCODE_BUTTON_B *AndroidGameKey
    KEYCODE_BUTTON_A *AndroidGameKey
    KEYCODE_BUTTON_X *AndroidGameKey
    KEYCODE_BUTTON_Y *AndroidGameKey
    KEYCODE_BUTTON_L2 *AndroidGameKey
    KEYCODE_BUTTON_R2 *AndroidGameKey
    KEYCODE_BUTTON_SELECT *AndroidGameKey
    KEYCODE_BUTTON_C *AndroidGameKey
    KEYCODE_BUTTON_Z *AndroidGameKey
    KEYCODE_BUTTON_1 *AndroidGameKey
    KEYCODE_BUTTON_10 *AndroidGameKey
    KEYCODE_BUTTON_11 *AndroidGameKey
    KEYCODE_BUTTON_12 *AndroidGameKey
    KEYCODE_BUTTON_13 *AndroidGameKey
    KEYCODE_BUTTON_14 *AndroidGameKey
    KEYCODE_BUTTON_15 *AndroidGameKey
    KEYCODE_BUTTON_16 *AndroidGameKey
    KEYCODE_BUTTON_2 *AndroidGameKey
    KEYCODE_BUTTON_3 *AndroidGameKey
    KEYCODE_BUTTON_4 *AndroidGameKey
    KEYCODE_BUTTON_5 *AndroidGameKey
    KEYCODE_BUTTON_6 *AndroidGameKey
    KEYCODE_BUTTON_7 *AndroidGameKey
    KEYCODE_BUTTON_8 *AndroidGameKey
    KEYCODE_BUTTON_9 *AndroidGameKey
    SOFT_LEFT *AndroidGameKey
    SOFT_RIGHT *AndroidGameKey
    HOME *AndroidGameKey
    BACK *AndroidGameKey
    CALL *AndroidGameKey
    ENDCALL *AndroidGameKey
    ZERO *AndroidGameKey
    ONE *AndroidGameKey
    TWO *AndroidGameKey
    THREE *AndroidGameKey
    FOUR *AndroidGameKey
    FIVE *AndroidGameKey
    SIX *AndroidGameKey
    SEVEN *AndroidGameKey
    EIGHT *AndroidGameKey
    NINE *AndroidGameKey
    STAR *AndroidGameKey
    POUND *AndroidGameKey
    DPAD_UP *AndroidGameKey
    DPAD_DOWN *AndroidGameKey
    DPAD_LEFT *AndroidGameKey
    DPAD_RIGHT *AndroidGameKey
    DPAD_CENTER *AndroidGameKey
    VOLUME_UP *AndroidGameKey
    VOLUME_DOWN *AndroidGameKey
    POWER *AndroidGameKey
    CAMERA *AndroidGameKey
    CLEAR *AndroidGameKey
    A *AndroidGameKey
    B *AndroidGameKey
    C *AndroidGameKey
    D *AndroidGameKey
    E *AndroidGameKey
    F *AndroidGameKey
    G *AndroidGameKey
    H *AndroidGameKey
    I *AndroidGameKey
    J *AndroidGameKey
    K *AndroidGameKey
    L *AndroidGameKey
    M *AndroidGameKey
    N *AndroidGameKey
    O *AndroidGameKey
    P *AndroidGameKey
    Q *AndroidGameKey
    R *AndroidGameKey
    S *AndroidGameKey
    T *AndroidGameKey
    U *AndroidGameKey
    V *AndroidGameKey
    W *AndroidGameKey
    X *AndroidGameKey
    Y *AndroidGameKey
    Z *AndroidGameKey
    COMMA *AndroidGameKey
    PERIOD *AndroidGameKey
    ALT_LEFT *AndroidGameKey
    ALT_RIGHT *AndroidGameKey
    SHIFT_LEFT *AndroidGameKey
    SHIFT_RIGHT *AndroidGameKey
    TAB *AndroidGameKey
    SPACE *AndroidGameKey
    SYM *AndroidGameKey
    EXPLORER *AndroidGameKey
    ENVELOPE *AndroidGameKey
    ENTER *AndroidGameKey
    DEL *AndroidGameKey
    GRAVE *AndroidGameKey
    MINUS *AndroidGameKey
    EQUALS *AndroidGameKey
    LEFT_BRACKET *AndroidGameKey
    RIGHT_BRACKET *AndroidGameKey
    BACKSLASH *AndroidGameKey
    SEMICOLON *AndroidGameKey
    APOSTROPHE *AndroidGameKey
    SLASH *AndroidGameKey
    AT *AndroidGameKey
    NUM *AndroidGameKey
    HEADSETHOOK *AndroidGameKey
    FOCUS *AndroidGameKey
    PLUS *AndroidGameKey
    MENU *AndroidGameKey
    NOTIFICATION *AndroidGameKey
    SEARCH *AndroidGameKey
}

// The Java class AndroidKeyFactory has subclasses so its variables are this interface
type IAndroidKeyFactory interface {
    IActivityFractureUtil
    AndroidKeyFactory_() *AndroidKeyFactory
    Init()
    IsSubmission(input IInput) bool
    IsDelete(input IInput) bool
    IsBackSpace(input IInput) bool
    IsLeft(input IInput) bool
    IsRight(input IInput) bool
    IsUp(input IInput) bool
    IsDown(input IInput) bool
    IsEnter(input IInput) bool
}

func (this *AndroidKeyFactory) AndroidKeyFactory_() *AndroidKeyFactory {
    return this
}


func (this *AndroidKeyFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.UNKNOWN = NewAndroidGameKey(KeyEvent_KEYCODE_UNKNOWN, "UNKNOWN")
    this.KEYCODE_BUTTON_L1 = this.UNKNOWN
    this.KEYCODE_BUTTON_R1 = this.UNKNOWN
    this.KEYCODE_BUTTON_THUMBR = this.UNKNOWN
    this.KEYCODE_BUTTON_THUMBL = this.UNKNOWN
    this.KEYCODE_BUTTON_START = this.UNKNOWN
    this.KEYCODE_BUTTON_MODE = this.UNKNOWN
    this.KEYCODE_BUTTON_B = this.UNKNOWN
    this.KEYCODE_BUTTON_A = this.UNKNOWN
    this.KEYCODE_BUTTON_X = this.UNKNOWN
    this.KEYCODE_BUTTON_Y = this.UNKNOWN
    this.KEYCODE_BUTTON_L2 = this.UNKNOWN
    this.KEYCODE_BUTTON_R2 = this.UNKNOWN
    this.KEYCODE_BUTTON_SELECT = this.UNKNOWN
    this.KEYCODE_BUTTON_C = this.UNKNOWN
    this.KEYCODE_BUTTON_Z = this.UNKNOWN
    this.KEYCODE_BUTTON_1 = this.UNKNOWN
    this.KEYCODE_BUTTON_10 = this.UNKNOWN
    this.KEYCODE_BUTTON_11 = this.UNKNOWN
    this.KEYCODE_BUTTON_12 = this.UNKNOWN
    this.KEYCODE_BUTTON_13 = this.UNKNOWN
    this.KEYCODE_BUTTON_14 = this.UNKNOWN
    this.KEYCODE_BUTTON_15 = this.UNKNOWN
    this.KEYCODE_BUTTON_16 = this.UNKNOWN
    this.KEYCODE_BUTTON_2 = this.UNKNOWN
    this.KEYCODE_BUTTON_3 = this.UNKNOWN
    this.KEYCODE_BUTTON_4 = this.UNKNOWN
    this.KEYCODE_BUTTON_5 = this.UNKNOWN
    this.KEYCODE_BUTTON_6 = this.UNKNOWN
    this.KEYCODE_BUTTON_7 = this.UNKNOWN
    this.KEYCODE_BUTTON_8 = this.UNKNOWN
    this.KEYCODE_BUTTON_9 = this.UNKNOWN
    this.SOFT_LEFT = NewAndroidGameKey(KeyEvent_KEYCODE_SOFT_LEFT, "SOFT_LEFT")
    this.SOFT_RIGHT = NewAndroidGameKey(KeyEvent_KEYCODE_SOFT_RIGHT, "SOFT_RIGHT")
    this.HOME = NewAndroidGameKey(KeyEvent_KEYCODE_HOME, "HOME")
    this.BACK = NewAndroidGameKey(KeyEvent_KEYCODE_BACK, "BACK")
    this.CALL = NewAndroidGameKey(KeyEvent_KEYCODE_CALL, "CALL")
    this.ENDCALL = NewAndroidGameKey(KeyEvent_KEYCODE_ENDCALL, "ENDCALL")
    this.ZERO = NewAndroidGameKey(KeyEvent_KEYCODE_0, CommonPhoneStrings_GetInstance().ZERO)
    this.ONE = NewAndroidGameKey(KeyEvent_KEYCODE_1, CommonPhoneStrings_GetInstance().ONE)
    this.TWO = NewAndroidGameKey(KeyEvent_KEYCODE_2, CommonPhoneStrings_GetInstance().TWO)
    this.THREE = NewAndroidGameKey(KeyEvent_KEYCODE_3, CommonPhoneStrings_GetInstance().THREE)
    this.FOUR = NewAndroidGameKey(KeyEvent_KEYCODE_4, CommonPhoneStrings_GetInstance().FOUR)
    this.FIVE = NewAndroidGameKey(KeyEvent_KEYCODE_5, CommonPhoneStrings_GetInstance().FIVE)
    this.SIX = NewAndroidGameKey(KeyEvent_KEYCODE_6, CommonPhoneStrings_GetInstance().SIX)
    this.SEVEN = NewAndroidGameKey(KeyEvent_KEYCODE_7, CommonPhoneStrings_GetInstance().SEVEN)
    this.EIGHT = NewAndroidGameKey(KeyEvent_KEYCODE_8, CommonPhoneStrings_GetInstance().EIGHT)
    this.NINE = NewAndroidGameKey(KeyEvent_KEYCODE_9, CommonPhoneStrings_GetInstance().NINE)
    this.STAR = NewAndroidGameKey(KeyEvent_KEYCODE_STAR, CommonPhoneStrings_GetInstance().STAR)
    this.POUND = NewAndroidGameKey(KeyEvent_KEYCODE_POUND, CommonPhoneStrings_GetInstance().POUND)
    this.DPAD_UP = NewAndroidGameKey(KeyEvent_KEYCODE_DPAD_UP, "UP")
    this.DPAD_DOWN = NewAndroidGameKey(KeyEvent_KEYCODE_DPAD_DOWN, "DOWN")
    this.DPAD_LEFT = NewAndroidGameKey(KeyEvent_KEYCODE_DPAD_LEFT, "LEFT")
    this.DPAD_RIGHT = NewAndroidGameKey(KeyEvent_KEYCODE_DPAD_RIGHT, "RIGHT")
    this.DPAD_CENTER = NewAndroidGameKey(KeyEvent_KEYCODE_DPAD_CENTER, "CENTER")
    this.VOLUME_UP = NewAndroidGameKey(KeyEvent_KEYCODE_VOLUME_UP, "VOLUME_UP")
    this.VOLUME_DOWN = NewAndroidGameKey(KeyEvent_KEYCODE_VOLUME_DOWN, "VOLUME_DOWN")
    this.POWER = NewAndroidGameKey(KeyEvent_KEYCODE_POWER, "POWER")
    this.CAMERA = NewAndroidGameKey(KeyEvent_KEYCODE_CAMERA, "CAMERA")
    this.CLEAR = NewAndroidGameKey(KeyEvent_KEYCODE_CLEAR, "CLEAR")
    this.A = NewAndroidGameKey(KeyEvent_KEYCODE_A, "A")
    this.B = NewAndroidGameKey(KeyEvent_KEYCODE_B, "B")
    this.C = NewAndroidGameKey(KeyEvent_KEYCODE_C, "C")
    this.D = NewAndroidGameKey(KeyEvent_KEYCODE_D, "D")
    this.E = NewAndroidGameKey(KeyEvent_KEYCODE_E, "E")
    this.F = NewAndroidGameKey(KeyEvent_KEYCODE_F, "F")
    this.G = NewAndroidGameKey(KeyEvent_KEYCODE_G, "G")
    this.H = NewAndroidGameKey(KeyEvent_KEYCODE_H, "H")
    this.I = NewAndroidGameKey(KeyEvent_KEYCODE_I, "I")
    this.J = NewAndroidGameKey(KeyEvent_KEYCODE_J, "J")
    this.K = NewAndroidGameKey(KeyEvent_KEYCODE_K, "K")
    this.L = NewAndroidGameKey(KeyEvent_KEYCODE_L, "L")
    this.M = NewAndroidGameKey(KeyEvent_KEYCODE_M, "M")
    this.N = NewAndroidGameKey(KeyEvent_KEYCODE_N, "N")
    this.O = NewAndroidGameKey(KeyEvent_KEYCODE_O, "O")
    this.P = NewAndroidGameKey(KeyEvent_KEYCODE_P, "P")
    this.Q = NewAndroidGameKey(KeyEvent_KEYCODE_Q, "Q")
    this.R = NewAndroidGameKey(KeyEvent_KEYCODE_R, "R")
    this.S = NewAndroidGameKey(KeyEvent_KEYCODE_S, "S")
    this.T = NewAndroidGameKey(KeyEvent_KEYCODE_T, "T")
    this.U = NewAndroidGameKey(KeyEvent_KEYCODE_U, "U")
    this.V = NewAndroidGameKey(KeyEvent_KEYCODE_V, "V")
    this.W = NewAndroidGameKey(KeyEvent_KEYCODE_W, "W")
    this.X = NewAndroidGameKey(KeyEvent_KEYCODE_X, "X")
    this.Y = NewAndroidGameKey(KeyEvent_KEYCODE_Y, "Y")
    this.Z = NewAndroidGameKey(KeyEvent_KEYCODE_Z, "Z")
    this.COMMA = NewAndroidGameKey(KeyEvent_KEYCODE_COMMA, "COMMA")
    this.PERIOD = NewAndroidGameKey(KeyEvent_KEYCODE_PERIOD, AbPathData_GetInstance().EXTENSION_SEP)
    this.ALT_LEFT = NewAndroidGameKey(KeyEvent_KEYCODE_ALT_LEFT, "ALT_LEFT")
    this.ALT_RIGHT = NewAndroidGameKey(KeyEvent_KEYCODE_ALT_RIGHT, "ALT_RIGHT")
    this.SHIFT_LEFT = NewAndroidGameKey(KeyEvent_KEYCODE_SHIFT_LEFT, "SHIFT_LEFT")
    this.SHIFT_RIGHT = NewAndroidGameKey(KeyEvent_KEYCODE_SHIFT_RIGHT, "SHIFT_RIGHT")
    this.TAB = NewAndroidGameKey(KeyEvent_KEYCODE_TAB, "TAB")
    this.SPACE = NewAndroidGameKey(KeyEvent_KEYCODE_SPACE, "SPACE")
    this.SYM = NewAndroidGameKey(KeyEvent_KEYCODE_SYM, "SYM")
    this.EXPLORER = NewAndroidGameKey(KeyEvent_KEYCODE_EXPLORER, "EXPLORER")
    this.ENVELOPE = NewAndroidGameKey(KeyEvent_KEYCODE_ENVELOPE, "ENVELOPE")
    this.ENTER = NewAndroidGameKey(KeyEvent_KEYCODE_ENTER, "ENTER")
    this.DEL = NewAndroidGameKey(KeyEvent_KEYCODE_DEL, "DEL")
    this.GRAVE = NewAndroidGameKey(KeyEvent_KEYCODE_GRAVE, "GRAVE")
    this.MINUS = NewAndroidGameKey(KeyEvent_KEYCODE_MINUS, "MINUS")
    this.EQUALS = NewAndroidGameKey(KeyEvent_KEYCODE_EQUALS, "EQUALS")
    this.LEFT_BRACKET = NewAndroidGameKey(KeyEvent_KEYCODE_LEFT_BRACKET, "LEFT_BRACKET")
    this.RIGHT_BRACKET = NewAndroidGameKey(KeyEvent_KEYCODE_RIGHT_BRACKET, "RIGHT_BRACKET")
    this.BACKSLASH = NewAndroidGameKey(KeyEvent_KEYCODE_BACKSLASH, "BACKSLASH")
    this.SEMICOLON = NewAndroidGameKey(KeyEvent_KEYCODE_SEMICOLON, "SEMICOLON")
    this.APOSTROPHE = NewAndroidGameKey(KeyEvent_KEYCODE_APOSTROPHE, "APOSTROPHE")
    this.SLASH = NewAndroidGameKey(KeyEvent_KEYCODE_SLASH, "SLASH")
    this.AT = NewAndroidGameKey(KeyEvent_KEYCODE_AT, "AT")
    this.NUM = NewAndroidGameKey(KeyEvent_KEYCODE_NUM, "NUM")
    this.HEADSETHOOK = NewAndroidGameKey(KeyEvent_KEYCODE_HEADSETHOOK, "HEADSETHOOK")
    this.FOCUS = NewAndroidGameKey(KeyEvent_KEYCODE_FOCUS, "FOCUS")
    this.PLUS = NewAndroidGameKey(KeyEvent_KEYCODE_PLUS, "PLUS")
    this.MENU = NewAndroidGameKey(KeyEvent_KEYCODE_MENU, "MENU")
    this.NOTIFICATION = NewAndroidGameKey(KeyEvent_KEYCODE_NOTIFICATION, "NOTIFICATION")
    this.SEARCH = NewAndroidGameKey(KeyEvent_KEYCODE_SEARCH, "SEARCH")
}

func NewAndroidKeyFactory() *AndroidKeyFactory {
    this := &AndroidKeyFactory{}
    this.ActivityFractureUtil = *NewActivityFractureUtil()
    this.initFields()
    return this
}


func (this *AndroidKeyFactory) Init() {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF("Start/End", this, commonStrings.INIT)
}


func (this *AndroidKeyFactory) IsSubmission(input IInput) bool {
        if input == this.ENTER {
            return true
        } else if input == this.DPAD_CENTER {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsDelete(input IInput) bool {
        return false
}


func (this *AndroidKeyFactory) IsBackSpace(input IInput) bool {
        if input == this.DEL {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsLeft(input IInput) bool {
        if input == this.DPAD_LEFT {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsRight(input IInput) bool {
        if input == this.DPAD_RIGHT {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsUp(input IInput) bool {
        if input == this.DPAD_UP {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsDown(input IInput) bool {
        if input == this.DPAD_DOWN {
            return true
        }
        return false
}


func (this *AndroidKeyFactory) IsEnter(input IInput) bool {
        if input == this.DPAD_CENTER {
            return true
        } else if input == this.ENTER {
            return true
        } else if input == this.CALL {
            return true
        }
        return false
}

