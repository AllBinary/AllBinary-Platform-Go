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

type AbSqlTableUtil struct {
    LogUtil *LogUtil
    EXTENSION string
    END string
    SAVING_BACKUP_PATH string
    FILE_LABEL string
    METHOD_RESTORE_TABLE string
    METHOD_BACKUP_TABLE string
    METHOD_BACKUP_FILE string
    METHOD_GET_OUTPUT_STREAM string
    TABLE_CREATION_SUCCESS string
    DROPPED_SUCCESS string
    SAVING string
    APPENDING string
    TABLE_LABEL string
    BACKUP_SUCCESS string
    ERROR_CREATING string
    TOTAL_LABEL string
    SECTION_LABEL string
    DASH string
    PORTION_RESTORED string
    SpecialCharArray []rune
    NEW_LINE string
    StringUtil *StringUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    SqlStrings *SqlStrings
}

var AbSqlTableUtil_Instance *AbSqlTableUtil = NewAbSqlTableUtil()

func (this *AbSqlTableUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.EXTENSION = ".adb"
    this.END = "')\n"
    this.SAVING_BACKUP_PATH = "Saving Backup: Path: "
    this.FILE_LABEL = " File: "
    this.METHOD_RESTORE_TABLE = "restoreTable()"
    this.METHOD_BACKUP_TABLE = "backupTable()"
    this.METHOD_BACKUP_FILE = "backupFile()"
    this.METHOD_GET_OUTPUT_STREAM = "getOutputStream()"
    this.TABLE_CREATION_SUCCESS = "Table Creation Successful: "
    this.DROPPED_SUCCESS = " Dropped Successfully"
    this.SAVING = "Saving: "
    this.APPENDING = "Appending: "
    this.TABLE_LABEL = "Table: "
    this.BACKUP_SUCCESS = " Backup Success"
    this.ERROR_CREATING = "Error Creating: "
    this.TOTAL_LABEL = " Total: "
    this.SECTION_LABEL = " Section: "
    this.DASH = " - "
    this.PORTION_RESTORED = " Portion Restored"
    this.SpecialCharArray = []rune{'\n', '\f', '\r'}
    this.NEW_LINE = "\\n"
    this.StringUtil = StringUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.SqlStrings = SqlStrings_GetInstance()
}

func NewAbSqlTableUtil() *AbSqlTableUtil {
    this := &AbSqlTableUtil{}
    this.initFields()
    return this
}


func AbSqlTableUtil_GetInstance() *AbSqlTableUtil {
        return AbSqlTableUtil_Instance
}


// synchronized
func (this *AbSqlTableUtil) GetOutputStream(backupPath string, tableName string) *OutputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *OutputStream) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put("Create File", this, this.METHOD_GET_OUTPUT_STREAM, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var fileName string = tableName + this.EXTENSION; _ = fileName
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.SAVING + tableName, this, this.METHOD_GET_OUTPUT_STREAM)
            }
            var backupFilePath *AbPath = NewAbPath(backupPath, fileName); _ = backupFilePath
            var backupFile *AbFile = AbFile_CreateAbFile(backupFilePath); _ = backupFile
            if backupFile.Exists() {
                this.BackupFile(backupFilePath, backupPath, tableName)
                backupFile.Delete()
            }
            backupFile.CreateNewFile()
            var outputStream *OutputStream = NewAbFileOutputStream(backupFile); _ = outputStream
            tryControl1, tryResult1 = 1, outputStream; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *AbSqlTableUtil) BackupFile(path *AbPath, backupPath string, tableName string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.METHOD_BACKUP_FILE, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var timeLong int64 = int64(calendar.GetTimeInMillis()); _ = timeLong
            var time string = fmt.Sprint(timeLong); _ = time
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var fileName string = tableName + this.EXTENSION; _ = fileName
            stringBuffer.Append(backupPath)
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            stringBuffer.Append(time)
            var backupAbPath *AbPath = NewAbPath(fmt.Sprint(stringBuffer)); _ = backupAbPath
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.SAVING_BACKUP_PATH)
                stringBuffer.Append(backupAbPath.ToFileSystemString())
                stringBuffer.Append(this.FILE_LABEL)
                stringBuffer.Append(fileName)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_BACKUP_FILE)
            }
            Directory_Create(backupAbPath)
            var backupFileBak *AbFile = AbFile_CreateAbFile(backupAbPath.ToFileSystemString()); _ = backupFileBak
            backupFileBak.CreateNewFile()
            FileUtil_GetInstance().Copy(path, backupAbPath)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlTableUtil) ConvertNewLines(value string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var index int = 0; _ = index
        var lastIndex int = 0; _ = lastIndex
        for index < len(value) {
            index = value.IndexOf(this.SpecialCharArray[0], lastIndex)
            if index != -1 {
                var nextLine string = value[lastIndex:index - 1]; _ = nextLine
                stringBuffer.Append(nextLine)
                stringBuffer.Append(this.NEW_LINE)
                lastIndex = index + 1
            } else {
                break
            }
        }
        if lastIndex < len(value) {
            stringBuffer.Append(value[lastIndex:len(value)])
        }
        return fmt.Sprint(stringBuffer)
}


// synchronized
func (this *AbSqlTableUtil) BackupTable(abSqlTable *AbSqlTable) string {
        var tableName string = abSqlTable.GetTableName(); _ = tableName
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put("Backup Table Failed\nSQL Statement", this, this.METHOD_BACKUP_TABLE, e)
            }
            tryControl1, tryResult1 = 1, this.TABLE_LABEL + tableName + " Backup Failed"; return
                }
            }()
            var count int = 0; _ = count
            var sqlStatement string = fmt.Sprint(this.SqlStrings.SELECT_ALL_FROM) + tableName; _ = sqlStatement
            var path string = org_Allbinary.Globals.URLGLOBALS.GetMainPath() + PATH_GLOBALS.GetInstance().BACKUP_PATH; _ = path
            if !Directory_Create(NewAbPath(path)) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.ERROR_CREATING + fmt.Sprint(path), this, this.METHOD_BACKUP_TABLE)
                }
            }
            var rset *ResultSet = abSqlTable.ExecuteSQLStatement(sqlStatement); _ = rset
            var rsmd *ResultSetMetaData = rset.GetMetaData(); _ = rsmd
            var colNum int = rsmd.GetColumnCount(); _ = colNum
            var QUERY_START string = fmt.Sprint(&strings.Builder{}.Append(this.SqlStrings.INSERT_INTO).Append(tableName).Append(this.SqlStrings.VALUES)); _ = QUERY_START
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var outputStream *OutputStream = this.GetOutputStream(path, tableName); _ = outputStream
            for rset.Next() {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(QUERY_START)
                for i := int(1); i < colNum; i++ {
                    var value string = rset.GetString(i); _ = value
                    stringBuffer.Append(this.ConvertNewLines(value))
                    stringBuffer.Append(this.SqlStrings.SINGLE_QUOTE_COMMA_SEP)
                }
                stringBuffer.Append(rset.GetString(colNum))
                stringBuffer.Append(this.END)
                var sqlStatementLine string = fmt.Sprint(stringBuffer); _ = sqlStatementLine
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.APPENDING + sqlStatementLine, this, this.METHOD_BACKUP_TABLE)
                }
                outputStream.Write(Java_GetBytes(sqlStatementLine))
            }
            StreamUtil_GetInstance().Close(outputStream)
            tryControl1, tryResult1 = 1, this.TABLE_LABEL + tableName + this.BACKUP_SUCCESS; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *AbSqlTableUtil) RestoreTable(abSqlTable *AbSqlTable, portion *Portion) string {
        var tableName string = abSqlTable.GetTableName(); _ = tableName
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put("Restore Table Failed\nSQL Statement", this, this.METHOD_RESTORE_TABLE, e)
            }
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.TABLE_LABEL)
            stringBuffer.Append(tableName)
            stringBuffer.Append(" Restoration Failed")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
                }
            }()
            var path string = org_Allbinary.Globals.URLGLOBALS.GetMainPath() + PATH_GLOBALS.GetInstance().BACKUP_PATH; _ = path
            var current int = int(portion.GetCurrent()); _ = current
            if current == 0 {
                if Directory_Create(NewAbPath(path)) {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                        this.LogUtil.PutF(this.ERROR_CREATING + fmt.Sprint(path), this, this.METHOD_RESTORE_TABLE)
                    }
                }
            }
            var backupFile *AbFile = AbFile_CreateAbFile(path, tableName + this.EXTENSION); _ = backupFile
            var bufferedLineReader *BufferedLineReader = NewBufferedLineReader(backupFile); _ = bufferedLineReader
            var size int64 = bufferedLineReader.GetSize(); _ = size
            var section int64 = size / int64(int(portion.GetTotal())) + 1; _ = section
            var start int64 = section * int64(current); _ = start
            var end int64 = start + section; _ = end
            if end > size {
                end = size
            }
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.TOTAL_LABEL)
            stringBuffer.Append(size)
            stringBuffer.Append(this.SECTION_LABEL)
            stringBuffer.Append(start)
            stringBuffer.Append(this.DASH)
            stringBuffer.Append(end)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_RESTORE_TABLE)
            }
            bufferedLineReader.ReadUpToLines(start)
            var line string = this.StringUtil.EMPTY_STRING; _ = line
            for bufferedLineReader.GetCurrent() < end && (func() any { line = bufferedLineReader.ReadLine(); return line }()) != nil {
                if len(line) > 1 {
                    abSqlTable.ExecuteSQLStatement(line)
                }
            }
            stringBuffer.Append(this.CommonSeps.SPACE)
            stringBuffer.Append(this.TABLE_LABEL)
            stringBuffer.Append(tableName)
            stringBuffer.Append(this.PORTION_RESTORED)
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

