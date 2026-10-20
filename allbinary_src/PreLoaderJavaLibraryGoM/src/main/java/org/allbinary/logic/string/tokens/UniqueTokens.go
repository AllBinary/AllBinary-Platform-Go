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

type UniqueTokens struct {
    SpecialCharacters *BasicArrayList
}


func (this *UniqueTokens) initFields() {
    this.SpecialCharacters = NewBasicArrayListD()
}


func NewUniqueTokens() *UniqueTokens {
    this := &UniqueTokens{}
    this.initFields()
        this.SpecialCharacters.Add("!")
        this.SpecialCharacters.Add("@")
        this.SpecialCharacters.Add(CommonPhoneStrings_GetInstance().POUND)
        this.SpecialCharacters.Add("$")
        this.SpecialCharacters.Add("%")
        this.SpecialCharacters.Add("^")
        this.SpecialCharacters.Add(CommonSeps_GetInstance().AMPERSAND)
        this.SpecialCharacters.Add(CommonPhoneStrings_GetInstance().STAR)
        this.SpecialCharacters.Add("(")
        this.SpecialCharacters.Add(")")
        this.SpecialCharacters.Add("-")
        this.SpecialCharacters.Add("_")
        this.SpecialCharacters.Add("+")
        this.SpecialCharacters.Add("=")
        this.SpecialCharacters.Add("\\")
        this.SpecialCharacters.Add("|")
        this.SpecialCharacters.Add(AbPathData_GetInstance().EXTENSION_SEP)
        this.SpecialCharacters.Add(",")
        this.SpecialCharacters.Add("<")
        this.SpecialCharacters.Add(">")
        this.SpecialCharacters.Add("?")
        this.SpecialCharacters.Add("/")
        this.SpecialCharacters.Add("~")
        this.SpecialCharacters.Add("`")
    return this
}


// throws Exception
func (this *UniqueTokens) GetWhithoutDashesAndSkipNumberOnlyTokens(stringVector *BasicArrayList) *HashSet {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HashSet) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var hashSet *HashSet = NewHashSet(); _ = hashSet
            var index int = 0; _ = index
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            for index < stringVector.Size() {
                var keywords string = any(stringVector.Get(index)).(string); _ = keywords
                var tokenizer *Tokenizer = NewTokenizer(commonSeps.COMMA); _ = tokenizer
                var keywordVector *BasicArrayList = tokenizer.GetTokensFromString(keywords, NewBasicArrayListD()); _ = keywordVector
                for forIndex := int(0); forIndex < keywordVector.Size(); forIndex++ {
                    var cleanString string = any(keywordVector.Get(forIndex)).(string); _ = cleanString
                    cleanString = strings.TrimSpace(cleanString)
                    if strings.Index(cleanString, string('-')) != -1 {
                        cleanString = NewReplace("-", commonSeps.SPACE).All(cleanString)
                    }
                    hashSet.Add(cleanString)
                    if strings.Index(cleanString, string(' ')) != -1 {
                        var spaceTokenizer *Tokenizer = NewTokenizer(commonSeps.SPACE); _ = spaceTokenizer
                        var subKeywordVector *BasicArrayList = spaceTokenizer.GetTokensFromString(cleanString, NewBasicArrayListD()); _ = subKeywordVector
                        for spaceIndex := int(0); spaceIndex < subKeywordVector.Size(); spaceIndex++ {
                            var subCleanString string = any(subKeywordVector.Get(spaceIndex)).(string); _ = subCleanString
                            if !this.IsSpecialCharacter(subCleanString) && !this.NumberOnly(subCleanString) {
                                hashSet.Add(subCleanString)
                            }
                        }
                    }
                }
                index++
            }
            tryControl1, tryResult1 = 1, hashSet; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *UniqueTokens) NumberOnly(subCleaningString string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            func() int { value, _ := strconv.Atoi(subCleaningString); return value }()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *UniqueTokens) IsSpecialCharacter(subCleaningString string) bool {
        for index := int(0); index < this.SpecialCharacters.Size(); index++ {
            if strings.Compare(subCleaningString, any(this.SpecialCharacters.Get(index)).(string)) == 0 {
                return true
            }
        }
        return false
}

