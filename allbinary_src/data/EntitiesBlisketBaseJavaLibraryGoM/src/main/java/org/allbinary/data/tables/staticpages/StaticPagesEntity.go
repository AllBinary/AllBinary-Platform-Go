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

type StaticPagesEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    TABLENAME string
}

// Implements StaticPagesEntityInterface


func (this *StaticPagesEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    this.TABLENAME = "staticpages"
}


func NewStaticPagesEntity() *StaticPagesEntity {
    this := &StaticPagesEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewStaticPagesDbInitInfo())
        this.initFields()
        this.SetTableName(this.TABLENAME)
    return this
}


func (this *StaticPagesEntity) Insert(values IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, AbSqlBean_INSERT, e)
            }
                }
            }()
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StaticPagesEntity) GetFile(store string, keywords string) string {
        var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
        whereHashMap.Put(StoreFrontData_GetInstance().NAME, store)
        whereHashMap.Put(this.BasicItemData.KEYWORDS, keywords)
        var file string = this.AbSqlBean.GetField(whereHashMap, SearchData_PAGE); _ = file
        return file
}


func (this *StaticPagesEntity) Delete(keywords string) {
        this.AbSqlBean.DeleteWhere(this.BasicItemData.KEYWORDS, keywords)
}


func (this *StaticPagesEntity) DropTable() string {
        var returnStr string = this.AbSqlBean.DropTable(); _ = returnStr
        return returnStr
}


func (this *StaticPagesEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(this.BasicItemData.KEYWORDS, any(updatedValues.Get(this.BasicItemData.KEYWORDS)).(string), updatedValues)
}


func (this *StaticPagesEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TABLENAME).Append(this.SqlStrings.START).Append(StoreFrontData_GetInstance().NAME).Append(this.SqlTypeStrings.SIXTY_CHAR_COLUMN_NOT_NULL).Append(this.BasicItemData.KEYWORDS).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(SearchData_PAGE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(StoreFrontData_GetInstance().NAME).Append(CommonSeps_GetInstance().COMMA_SEP).Append(this.BasicItemData.KEYWORDS).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *StaticPagesEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}

