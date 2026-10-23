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

type SmallIntegerSingletonFactory struct {
    NEGATIVE_MAX int
    POSITIVE_MAX int
    INTEGER_ARRAY []int
    STRING_ARRAY []string
    MIN int
    LastMin int
    LastNegativeMin int
}

var SmallIntegerSingletonFactory_Instance *SmallIntegerSingletonFactory = NewSmallIntegerSingletonFactory()

func (this *SmallIntegerSingletonFactory) initFields() {
    this.NEGATIVE_MAX = 500
    this.POSITIVE_MAX = 0x2D1
    this.INTEGER_ARRAY = make([]int, this.NEGATIVE_MAX + this.POSITIVE_MAX)
    this.STRING_ARRAY = make([]string, this.NEGATIVE_MAX + this.POSITIVE_MAX)
    this.MIN = 0
    this.LastMin = 0
    this.LastNegativeMin = 0
}


func SmallIntegerSingletonFactory_GetInstance() *SmallIntegerSingletonFactory {
        return SmallIntegerSingletonFactory_Instance
}


func (this *SmallIntegerSingletonFactory) GetMin() int {
        var minAllowed int = this.GetMinAllowed(); _ = minAllowed
        if this.MIN <= minAllowed {
            if J2MEUtil_IsJ2ME() {
            } else if AndroidUtil_IsAndroid() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                logUtil.PutF(fmt.Sprint(NewStringMaker().Append("Android InputFactory was initialized before GameMidlet: ").Appendint(this.MIN)), this, "getMin")
                this.InitWithRange(0x291, 6)
            } else {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                logUtil.PutF(fmt.Sprint(NewStringMaker().Append("InputFactory was initialized before GameMidlet or KeyFactoryInitializer - Currently this is occurs on JS build by TouchMotionGestureFactory constructor: ").Appendint(this.MIN)), this, "getMin")
                this.InitWithRange(0x2D0, 6)
            }
        }
        if this.MIN <= minAllowed {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            logUtil.Put("This means you loaded the InputFactory before determining the platform input size requirements.", this, "getMin", errors.New(fmt.Sprint()))
        }
        return this.MIN
}


func (this *SmallIntegerSingletonFactory) GetMinAllowed() int {
        if J2MEUtil_IsJ2ME() {
            return 0
        } else {
            if AndroidUtil_IsAndroid() {
                return 0x101
            } else {
                return 23
            }
        }
}


func (this *SmallIntegerSingletonFactory) InitWithRange(value int, negativeValue int) {
        for index := int(value - 1); index >= this.LastMin; index-- {
            this.INTEGER_ARRAY[index + this.NEGATIVE_MAX] = int(index)
        }
        for index := int(negativeValue - 1); index >= this.LastNegativeMin; index-- {
            this.INTEGER_ARRAY[index] = int(-index)
        }
        if this.LastMin < value {
            this.LastMin = value
            this.MIN = value
        }
        if this.LastNegativeMin < negativeValue {
            this.LastNegativeMin = negativeValue
        }
}


func (this *SmallIntegerSingletonFactory) Init() {
        if this.LastMin < this.POSITIVE_MAX || this.LastNegativeMin < this.NEGATIVE_MAX {
            for index := int(this.POSITIVE_MAX - 1); index >= this.LastMin; index-- {
                this.INTEGER_ARRAY[index + this.NEGATIVE_MAX] = int(index)
            }
            for index := int(this.NEGATIVE_MAX - 1); index >= this.LastNegativeMin; index-- {
                this.INTEGER_ARRAY[index] = int(-index)
            }
            this.LastMin = this.POSITIVE_MAX
            this.MIN = this.LastMin
            this.LastNegativeMin = this.NEGATIVE_MAX
        }
}


func NewSmallIntegerSingletonFactory() *SmallIntegerSingletonFactory {
    this := &SmallIntegerSingletonFactory{}
    this.initFields()
        this.InitWithRange(23, 0)
    return this
}


func (this *SmallIntegerSingletonFactory) GetAt(index int) int {
        return this.INTEGER_ARRAY[index + this.NEGATIVE_MAX]
}


func (this *SmallIntegerSingletonFactory) GetAtNoThrow(index int) int {
        if index + this.NEGATIVE_MAX > len(this.INTEGER_ARRAY) - 1 {
            return this.INTEGER_ARRAY[-1 + this.NEGATIVE_MAX]
        }
        return this.INTEGER_ARRAY[index + this.NEGATIVE_MAX]
}


func (this *SmallIntegerSingletonFactory) CreateInstance(index int) int {
        var integer int = this.GetAt(index); _ = integer
        if integer == 0 {
            integer = int(index)
        }
        return integer
}


func (this *SmallIntegerSingletonFactory) GetString(index int) string {
        var i int = index + this.NEGATIVE_MAX; _ = i
        if this.STRING_ARRAY[i] == "" {
            this.STRING_ARRAY[i] = fmt.Sprint(this.INTEGER_ARRAY[i])
        }
        return this.STRING_ARRAY[i]
}

