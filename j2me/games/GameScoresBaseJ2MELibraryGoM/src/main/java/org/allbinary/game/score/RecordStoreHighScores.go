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

type RecordStoreHighScores struct {
    HighScores
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PlatformRecordIdUtil *PlatformRecordIdUtil
    TsUtil *TsUtil
    RECORD_ID string
    MAXHIGHSCORES int
    GameInfo *GameInfo
    AbeClientInformation AbeClientInformationInterface
    RecordComparatorInterface RecordComparator
}

var RecordStoreHighScores_HashTable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable()

func (this *RecordStoreHighScores) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.PlatformRecordIdUtil = PlatformRecordIdUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.RECORD_ID = "_HS"
    this.MAXHIGHSCORES = 100
    
    
    
}


func NewRecordStoreHighScores(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, name string, heading string, columnTwoHeading string, recordComparatorInterface RecordComparator) *RecordStoreHighScores {
    this := &RecordStoreHighScores{}
        this.HighScores = *NewHighScores(name, heading, columnTwoHeading)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.GameInfo = gameInfo
        this.RecordComparatorInterface = recordComparatorInterface
        this.Load()
    return this
}


// synchronized
func RecordStoreHighScores_GetInstance(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, highScoreName string, heading string, columnTwoHeading string, recordComparatorInterface RecordComparator) IHighScores {
        var highScoresCanBeNull any = RecordStoreHighScores_HashTable.Get(any(highScoreName).(any)); _ = highScoresCanBeNull
        if highScoresCanBeNull == nil {
            var highScores IHighScores = NewRecordStoreHighScores(abeClientInformation, gameInfo, highScoreName, heading, columnTwoHeading, recordComparatorInterface); _ = highScores
            RecordStoreHighScores_HashTable.Put(highScores.GetName(), highScores)
            return highScores
        }
        return any(highScoresCanBeNull).(IHighScores)
}


func (this *RecordStoreHighScores) GetRecordId(abeClientInformation AbeClientInformationInterface) string {
        return this.PlatformRecordIdUtil.GetRecordId(abeClientInformation, fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().UNDERSCORE).Append(this.GetName()).Append(this.RECORD_ID)))
}


// synchronized
func (this *RecordStoreHighScores) AddHighScore(newHighScore *HighScore) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.ADD, e)
                }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.ADD, e)
                    }
                }()
                if recordStore != nil {
                    PreLogUtil_Put("Closing RecordStore", this, this.CommonStrings.ADD)
                    recordStore.CloseRecordStore()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Adding HighScore: ").Appendlong(newHighScore.GetScore())), this, this.CommonStrings.ADD)
            if this.IsTooManyHighScores() {
                this.LogUtil.PutF("Removing Lowest Score", this, this.CommonStrings.ADD)
                this.RemoveLowestHighScore()
            }
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(this.AbeClientInformation), true)
            var highScoreBytes []int8 = newHighScore.GetAsBytes(); _ = highScoreBytes
            var recordId int = recordStore.AddRecord(highScoreBytes, 0, len(highScoreBytes)); _ = recordId
            this.Load()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *RecordStoreHighScores) RemoveLowestHighScore() {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "removeLowestHighScore", e)
                }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "removeLowestHighScore", e)
                    }
                }()
                if recordStore != nil {
                    PreLogUtil_Put("Closing RecordStore", this, "removeLowestHighScore")
                    recordStore.CloseRecordStore()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            }()
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(this.AbeClientInformation), true)
            var recordEnum RecordEnumeration = recordStore.EnumerateRecords(NullRecordFilter_NULL_RECORD_FILTER, NullRecordComparator_NULL_RECORD_COMPARATOR, true); _ = recordEnum
            var scoreComparator *ScoreComparator = (any(this.RecordComparatorInterface).(*ScoreComparator)); _ = scoreComparator
            var bestHighScore *HighScore = NewHighScore(-1, "none", GameInfo_NONE, scoreComparator.GetBestScore()); _ = bestHighScore
            var recordAsBytes []int8; _ = recordAsBytes
            var byteArrayInputStream *ByteArrayInputStream; _ = byteArrayInputStream
            var inputStream IDataInputStream; _ = inputStream
            for recordEnum.HasNextElement() {
                var id int = recordEnum.NextRecordId(); _ = id
                recordAsBytes = this.TsUtil.GetRecord(recordStore, id)
                if recordAsBytes != nil {
                    byteArrayInputStream = NewByteArrayInputStream(recordAsBytes)
                    inputStream = NewDataInputStream(byteArrayInputStream)
                    var name string = inputStream.ReadUTF(); _ = name
                    var nextScore int64 = inputStream.ReadLong(); _ = nextScore
                    var nextCurrentHighScore *HighScore = NewHighScore(id, name, GameInfo_NONE, nextScore); _ = nextCurrentHighScore
                    if this.RecordComparatorInterface.Compare(nextCurrentHighScore.GetAsBytes(), bestHighScore.GetAsBytes()) == RecordComparator_FOLLOWS {
                        bestHighScore = nextCurrentHighScore
                    }
                }
            }
            if bestHighScore.GetId() != -1 {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Removing Lowest HighScore: ").Appendlong(bestHighScore.GetScore())), this, this.CommonStrings.LOAD)
                recordStore.DeleteRecord(bestHighScore.GetId())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *RecordStoreHighScores) Load() {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("No High Scores", this, this.CommonStrings.LOAD, e)
                }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.LOAD, e)
                    }
                }()
                if recordStore != nil {
                    PreLogUtil_Put("Closing RecordStore", this, this.CommonStrings.LOAD)
                    recordStore.CloseRecordStore()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            }()
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(this.AbeClientInformation), true)
            this.SetList(NewBasicArrayListD())
            var recordEnum RecordEnumeration = recordStore.EnumerateRecords(NullRecordFilter_NULL_RECORD_FILTER, NullRecordComparator_NULL_RECORD_COMPARATOR, true); _ = recordEnum
            var recordAsBytes []int8; _ = recordAsBytes
            var byteArrayInputStream *ByteArrayInputStream; _ = byteArrayInputStream
            var inputStream IDataInputStream; _ = inputStream
            for recordEnum.HasNextElement() {
                var id int = recordEnum.NextRecordId(); _ = id
                recordAsBytes = this.TsUtil.GetRecord(recordStore, id)
                if recordAsBytes != nil {
                    byteArrayInputStream = NewByteArrayInputStream(recordAsBytes)
                    inputStream = NewDataInputStream(byteArrayInputStream)
                    if tryControl2 := func() (tryControl2 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        this.LogUtil.Put("EOF", this, this.CommonStrings.LOAD, e)
                        panic(e)
                            }
                        }()
                        var name string = inputStream.ReadUTF(); _ = name
                        var score int64 = inputStream.ReadLong(); _ = score
                        var newHighScore *HighScore = NewHighScore(id, name, GameInfo_NONE, score); _ = newHighScore
                        var list IBasicArrayList = this.GetList(); _ = list
                        var size int = list.Size(); _ = size
                        var lastIndex int = size; _ = lastIndex
                        for index := int(0); index < size; index++ {
                            var highScore *HighScore = any(list.BasicArrayList_().ObjectArray[index]).(*HighScore); _ = highScore
                            if this.RecordComparatorInterface.Compare(newHighScore.GetAsBytes(), highScore.GetAsBytes()) == RecordComparator_PRECEDES {
                                lastIndex = index
                                break
                            }
                        }
                        list.AddAt(lastIndex, newHighScore)
                        return
                    }(); tryControl2 == 1 {
                        tryControl1 = 1; return
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *RecordStoreHighScores) IsTooManyHighScores() bool {
        if this.GetList() != nil && this.GetList().Size() < this.MAXHIGHSCORES {
            return false
        } else {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("HighScores RecordStore Max Reached: ").Appendint(this.MAXHIGHSCORES)), this, "isTooManyHighScores")
            return true
        }
}


// throws Exception
// synchronized
func (this *RecordStoreHighScores) IsBestScore(newHighScore *HighScore) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.ADD, e)
            panic(e)
                }
            }()
            if !this.IsTooManyHighScores() {
                this.LogUtil.PutF("Slot Available for a High Score", this, "isBestScore")
                tryControl1, tryResult1 = 1, true; return
            } else {
                var list IBasicArrayList = this.GetList(); _ = list
                var size int = list.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var highScore *HighScore = any(list.BasicArrayList_().ObjectArray[index]).(*HighScore); _ = highScore
                    if this.RecordComparatorInterface.Compare(newHighScore.GetAsBytes(), highScore.GetAsBytes()) == RecordComparator_FOLLOWS {
                        this.LogUtil.PutF("Obtained a High Score", this, "isBestScore")
                        tryControl1, tryResult1 = 1, true; return
                    }
                }
            }
            this.LogUtil.PutF("Not a High Score", this, "isBestScore")
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RecordStoreHighScores) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.HighScores.ToString())
        var list IBasicArrayList = this.GetList(); _ = list
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var highScore *HighScore = any(list.BasicArrayList_().ObjectArray[index]).(*HighScore); _ = highScore
            stringBuffer.Append(highScore.GetScoreString())
            stringBuffer.Append(CommonSeps_GetInstance().COMMA_SEP)
        }
        return fmt.Sprint(stringBuffer)
}

func (this *RecordStoreHighScores) String() string {
    return this.ToString()
}

