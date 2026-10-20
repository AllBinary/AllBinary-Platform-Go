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

type PCKeyFactory struct {
    UNKNOWN *PCGameKey
    ENTER *PCGameKey
    _A *PCGameKey
    _B *PCGameKey
    _C *PCGameKey
    _D *PCGameKey
    _E *PCGameKey
    _F *PCGameKey
    _G *PCGameKey
    _H *PCGameKey
    _I *PCGameKey
    _J *PCGameKey
    _K *PCGameKey
    _L *PCGameKey
    _M *PCGameKey
    _N *PCGameKey
    _O *PCGameKey
    _P *PCGameKey
    _Q *PCGameKey
    _R *PCGameKey
    _S *PCGameKey
    _T *PCGameKey
    _U *PCGameKey
    _V *PCGameKey
    _W *PCGameKey
    _X *PCGameKey
    _Y *PCGameKey
    _Z *PCGameKey
    A *PCGameKey
    B *PCGameKey
    C *PCGameKey
    D *PCGameKey
    E *PCGameKey
    F *PCGameKey
    G *PCGameKey
    H *PCGameKey
    I *PCGameKey
    J *PCGameKey
    K *PCGameKey
    L *PCGameKey
    M *PCGameKey
    N *PCGameKey
    O *PCGameKey
    P *PCGameKey
    Q *PCGameKey
    R *PCGameKey
    S *PCGameKey
    T *PCGameKey
    U *PCGameKey
    V *PCGameKey
    W *PCGameKey
    X *PCGameKey
    Y *PCGameKey
    Z *PCGameKey
    ZERO *PCGameKey
    ONE *PCGameKey
    TWO *PCGameKey
    THREE *PCGameKey
    FOUR *PCGameKey
    FIVE *PCGameKey
    SIX *PCGameKey
    SEVEN *PCGameKey
    EIGHT *PCGameKey
    NINE *PCGameKey
    DPAD_UP *PCGameKey
    DPAD_DOWN *PCGameKey
    DPAD_LEFT *PCGameKey
    DPAD_RIGHT *PCGameKey
    DPAD_UP2 *PCGameKey
    DPAD_DOWN2 *PCGameKey
    DPAD_LEFT2 *PCGameKey
    DPAD_RIGHT2 *PCGameKey
    COMMA *PCGameKey
    SPACE *PCGameKey
    DEL *PCGameKey
    STAR *PCGameKey
    QUESTION *PCGameKey
    PERIOD *PCGameKey
    ESCAPE *PCGameKey
    BACK_SPACE *PCGameKey
}

var PCKeyFactory_SINGLETON *PCKeyFactory = NewPCKeyFactory()

func (this *PCKeyFactory) initFields() {
    this.UNKNOWN = NewPCGameKey(254, "UNKNOWN")
    this.ENTER = NewPCGameKey(13, "Enter")
    this._A = NewPCGameKey(65, "A")
    this._B = NewPCGameKey(66, "B")
    this._C = NewPCGameKey(67, "C")
    this._D = NewPCGameKey(68, "D")
    this._E = NewPCGameKey(69, "E")
    this._F = NewPCGameKey(70, "F")
    this._G = NewPCGameKey(71, "G")
    this._H = NewPCGameKey(72, "H")
    this._I = NewPCGameKey(73, "I")
    this._J = NewPCGameKey(74, "J")
    this._K = NewPCGameKey(75, "K")
    this._L = NewPCGameKey(76, "L")
    this._M = NewPCGameKey(77, "M")
    this._N = NewPCGameKey(78, "N")
    this._O = NewPCGameKey(79, "O")
    this._P = NewPCGameKey(80, "P")
    this._Q = NewPCGameKey(81, "Q")
    this._R = NewPCGameKey(82, "R")
    this._S = NewPCGameKey(83, "S")
    this._T = NewPCGameKey(84, "T")
    this._U = NewPCGameKey(85, "U")
    this._V = NewPCGameKey(86, "V")
    this._W = NewPCGameKey(87, "W")
    this._X = NewPCGameKey(88, "X")
    this._Y = NewPCGameKey(89, "Y")
    this._Z = NewPCGameKey(90, "Z")
    this.A = NewPCGameKey(97, "a")
    this.B = NewPCGameKey(98, "b")
    this.C = NewPCGameKey(99, "c")
    this.D = NewPCGameKey(100, "d")
    this.E = NewPCGameKey(101, "e")
    this.F = NewPCGameKey(102, "f")
    this.G = NewPCGameKey(103, "g")
    this.H = NewPCGameKey(104, "h")
    this.I = NewPCGameKey(105, "i")
    this.J = NewPCGameKey(106, "j")
    this.K = NewPCGameKey(107, "k")
    this.L = NewPCGameKey(108, "l")
    this.M = NewPCGameKey(109, "m")
    this.N = NewPCGameKey(110, "n")
    this.O = NewPCGameKey(111, "o")
    this.P = NewPCGameKey(112, "p")
    this.Q = NewPCGameKey(113, "q")
    this.R = NewPCGameKey(114, "r")
    this.S = NewPCGameKey(115, "s")
    this.T = NewPCGameKey(116, "t")
    this.U = NewPCGameKey(117, "u")
    this.V = NewPCGameKey(118, "v")
    this.W = NewPCGameKey(119, "w")
    this.X = NewPCGameKey(120, PositionStrings_GetInstance().X)
    this.Y = NewPCGameKey(121, PositionStrings_GetInstance().Y)
    this.Z = NewPCGameKey(122, "z")
    this.ZERO = NewPCGameKey(48, CommonPhoneStrings_GetInstance().ZERO)
    this.ONE = NewPCGameKey(49, CommonPhoneStrings_GetInstance().ONE)
    this.TWO = NewPCGameKey(50, CommonPhoneStrings_GetInstance().TWO)
    this.THREE = NewPCGameKey(51, CommonPhoneStrings_GetInstance().THREE)
    this.FOUR = NewPCGameKey(52, CommonPhoneStrings_GetInstance().FOUR)
    this.FIVE = NewPCGameKey(53, CommonPhoneStrings_GetInstance().FIVE)
    this.SIX = NewPCGameKey(54, CommonPhoneStrings_GetInstance().SIX)
    this.SEVEN = NewPCGameKey(55, CommonPhoneStrings_GetInstance().SEVEN)
    this.EIGHT = NewPCGameKey(56, CommonPhoneStrings_GetInstance().EIGHT)
    this.NINE = NewPCGameKey(57, CommonPhoneStrings_GetInstance().NINE)
    this.DPAD_UP = NewPCGameKey(38, CommonPhoneStrings_GetInstance().UP)
    this.DPAD_DOWN = NewPCGameKey(40, CommonPhoneStrings_GetInstance().DOWN)
    this.DPAD_LEFT = NewPCGameKey(37, CommonPhoneStrings_GetInstance().LEFT)
    this.DPAD_RIGHT = NewPCGameKey(39, CommonPhoneStrings_GetInstance().RIGHT)
    this.DPAD_UP2 = this.DPAD_UP
    this.DPAD_DOWN2 = this.DPAD_DOWN
    this.DPAD_LEFT2 = this.DPAD_LEFT
    this.DPAD_RIGHT2 = this.DPAD_RIGHT
    this.COMMA = NewPCGameKey(44, "COMMA")
    this.SPACE = NewPCGameKey(32, "SPACE")
    this.DEL = NewPCGameKey(46, "DEL")
    this.STAR = NewPCGameKey(42, CommonPhoneStrings_GetInstance().STAR)
    this.QUESTION = NewPCGameKey(47, "?")
    this.PERIOD = NewPCGameKey(46, CommonSeps_GetInstance().PERIOD)
    this.ESCAPE = NewPCGameKey(27, "Esc")
    this.BACK_SPACE = NewPCGameKey(8, "Backspace")
}

func NewPCKeyFactory() *PCKeyFactory {
    this := &PCKeyFactory{}
    this.initFields()
    return this
}


func PCKeyFactory_GetInstance() *PCKeyFactory {
        return PCKeyFactory_SINGLETON
}


func (this *PCKeyFactory) IsSubmission(input IInput) bool {
        if input == this.ENTER {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsDelete(input IInput) bool {
        if input == this.DEL {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsBackSpace(input IInput) bool {
        if input == this.BACK_SPACE {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsUp(input IInput) bool {
        if input == this.DPAD_UP || input == this.DPAD_UP2 {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsDown(input IInput) bool {
        if input == this.DPAD_DOWN || input == this.DPAD_DOWN2 {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsLeft(input IInput) bool {
        if input == this.DPAD_LEFT || input == this.DPAD_LEFT2 {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsRight(input IInput) bool {
        if input == this.DPAD_RIGHT || input == this.DPAD_RIGHT2 {
            return true
        }
        return false
}


func (this *PCKeyFactory) IsEnter(input IInput) bool {
        if input == this.ENTER {
            return true
        }
        return false
}

