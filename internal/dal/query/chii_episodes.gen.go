// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/bangumi/server/internal/dal/dao"
)

func newEpisode(db *gorm.DB) episode {
	_episode := episode{}

	_episode.episodeDo.UseDB(db)
	_episode.episodeDo.UseModel(&dao.Episode{})

	tableName := _episode.episodeDo.TableName()
	_episode.ALL = field.NewField(tableName, "*")
	_episode.ID = field.NewUint32(tableName, "ep_id")
	_episode.SubjectID = field.NewUint32(tableName, "ep_subject_id")
	_episode.Sort = field.NewFloat32(tableName, "ep_sort")
	_episode.Type = field.NewInt8(tableName, "ep_type")
	_episode.Disc = field.NewUint8(tableName, "ep_disc")
	_episode.Name = field.NewString(tableName, "ep_name")
	_episode.NameCn = field.NewString(tableName, "ep_name_cn")
	_episode.Rate = field.NewInt8(tableName, "ep_rate")
	_episode.Duration = field.NewString(tableName, "ep_duration")
	_episode.Airdate = field.NewString(tableName, "ep_airdate")
	_episode.Online = field.NewString(tableName, "ep_online")
	_episode.Comment = field.NewUint32(tableName, "ep_comment")
	_episode.Resources = field.NewUint32(tableName, "ep_resources")
	_episode.Desc = field.NewString(tableName, "ep_desc")
	_episode.Dateline = field.NewUint32(tableName, "ep_dateline")
	_episode.Lastpost = field.NewUint32(tableName, "ep_lastpost")
	_episode.Lock = field.NewUint8(tableName, "ep_lock")
	_episode.Ban = field.NewUint8(tableName, "ep_ban")
	_episode.Subject = episodeBelongsToSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_episode.fillFieldMap()

	return _episode
}

type episode struct {
	episodeDo episodeDo

	ALL       field.Field
	ID        field.Uint32
	SubjectID field.Uint32
	Sort      field.Float32
	Type      field.Int8
	Disc      field.Uint8
	Name      field.String
	NameCn    field.String
	Rate      field.Int8
	Duration  field.String
	Airdate   field.String
	Online    field.String
	Comment   field.Uint32
	Resources field.Uint32
	Desc      field.String
	Dateline  field.Uint32
	Lastpost  field.Uint32
	Lock      field.Uint8
	Ban       field.Uint8
	Subject   episodeBelongsToSubject

	fieldMap map[string]field.Expr
}

func (e episode) Table(newTableName string) *episode {
	e.episodeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e episode) As(alias string) *episode {
	e.episodeDo.DO = *(e.episodeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *episode) updateTableName(table string) *episode {
	e.ALL = field.NewField(table, "*")
	e.ID = field.NewUint32(table, "ep_id")
	e.SubjectID = field.NewUint32(table, "ep_subject_id")
	e.Sort = field.NewFloat32(table, "ep_sort")
	e.Type = field.NewInt8(table, "ep_type")
	e.Disc = field.NewUint8(table, "ep_disc")
	e.Name = field.NewString(table, "ep_name")
	e.NameCn = field.NewString(table, "ep_name_cn")
	e.Rate = field.NewInt8(table, "ep_rate")
	e.Duration = field.NewString(table, "ep_duration")
	e.Airdate = field.NewString(table, "ep_airdate")
	e.Online = field.NewString(table, "ep_online")
	e.Comment = field.NewUint32(table, "ep_comment")
	e.Resources = field.NewUint32(table, "ep_resources")
	e.Desc = field.NewString(table, "ep_desc")
	e.Dateline = field.NewUint32(table, "ep_dateline")
	e.Lastpost = field.NewUint32(table, "ep_lastpost")
	e.Lock = field.NewUint8(table, "ep_lock")
	e.Ban = field.NewUint8(table, "ep_ban")

	e.fillFieldMap()

	return e
}

func (e *episode) WithContext(ctx context.Context) *episodeDo { return e.episodeDo.WithContext(ctx) }

func (e episode) TableName() string { return e.episodeDo.TableName() }

func (e episode) Alias() string { return e.episodeDo.Alias() }

func (e *episode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *episode) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 19)
	e.fieldMap["ep_id"] = e.ID
	e.fieldMap["ep_subject_id"] = e.SubjectID
	e.fieldMap["ep_sort"] = e.Sort
	e.fieldMap["ep_type"] = e.Type
	e.fieldMap["ep_disc"] = e.Disc
	e.fieldMap["ep_name"] = e.Name
	e.fieldMap["ep_name_cn"] = e.NameCn
	e.fieldMap["ep_rate"] = e.Rate
	e.fieldMap["ep_duration"] = e.Duration
	e.fieldMap["ep_airdate"] = e.Airdate
	e.fieldMap["ep_online"] = e.Online
	e.fieldMap["ep_comment"] = e.Comment
	e.fieldMap["ep_resources"] = e.Resources
	e.fieldMap["ep_desc"] = e.Desc
	e.fieldMap["ep_dateline"] = e.Dateline
	e.fieldMap["ep_lastpost"] = e.Lastpost
	e.fieldMap["ep_lock"] = e.Lock
	e.fieldMap["ep_ban"] = e.Ban

}

func (e episode) clone(db *gorm.DB) episode {
	e.episodeDo.ReplaceDB(db)
	return e
}

type episodeBelongsToSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a episodeBelongsToSubject) Where(conds ...field.Expr) *episodeBelongsToSubject {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a episodeBelongsToSubject) WithContext(ctx context.Context) *episodeBelongsToSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a episodeBelongsToSubject) Model(m *dao.Episode) *episodeBelongsToSubjectTx {
	return &episodeBelongsToSubjectTx{a.db.Model(m).Association(a.Name())}
}

type episodeBelongsToSubjectTx struct{ tx *gorm.Association }

func (a episodeBelongsToSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a episodeBelongsToSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a episodeBelongsToSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a episodeBelongsToSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a episodeBelongsToSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a episodeBelongsToSubjectTx) Count() int64 {
	return a.tx.Count()
}

type episodeDo struct{ gen.DO }

//GetByID ...
//
//where(ep_ep_id=@id)
func (e episodeDo) GetByID(id uint32) (result *dao.Episode, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("ep_ep_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = e.UnderlyingDB().Where(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = e.UnderlyingDB().Where(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

//GetBySubjectID ...
//
//where(ep_subject_id=@id ORDER BY `ep_disc`,`ep_type`,`ep_sort`)
func (e episodeDo) GetBySubjectID(id uint32) (result []*dao.Episode, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("ep_subject_id=@id ORDER BY `ep_disc`,`ep_type`,`ep_sort` ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = e.UnderlyingDB().Where(generateSQL.String(), params).Find(&result)
	} else {
		executeSQL = e.UnderlyingDB().Where(generateSQL.String()).Find(&result)
	}
	err = executeSQL.Error
	return
}

//GetFirst get first episode of a repository.
//where(ep_subject_id=@subjectID ORDER BY `ep_disc`,`ep_type`,`ep_sort` LIMIT 1)
func (e episodeDo) GetFirst(subjectID uint32) (result *dao.Episode, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["subjectID"] = subjectID
	generateSQL.WriteString("get first episode of a repository. where(ep_subject_id=@subjectID ORDER BY `ep_disc`,`ep_type`,`ep_sort` LIMIT 1) ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

//CountBySubjectID ...
//
//sql(select count(ep_id) from chii_episodes where ep_subject_id=@id)
func (e episodeDo) CountBySubjectID(id uint32) (result int64, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("select count(ep_id) from chii_episodes where ep_subject_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (e episodeDo) Debug() *episodeDo {
	return e.withDO(e.DO.Debug())
}

func (e episodeDo) WithContext(ctx context.Context) *episodeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e episodeDo) Clauses(conds ...clause.Expression) *episodeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e episodeDo) Returning(value interface{}, columns ...string) *episodeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e episodeDo) Not(conds ...gen.Condition) *episodeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e episodeDo) Or(conds ...gen.Condition) *episodeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e episodeDo) Select(conds ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e episodeDo) Where(conds ...gen.Condition) *episodeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e episodeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *episodeDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e episodeDo) Order(conds ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e episodeDo) Distinct(cols ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e episodeDo) Omit(cols ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e episodeDo) Join(table schema.Tabler, on ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e episodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *episodeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e episodeDo) RightJoin(table schema.Tabler, on ...field.Expr) *episodeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e episodeDo) Group(cols ...field.Expr) *episodeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e episodeDo) Having(conds ...gen.Condition) *episodeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e episodeDo) Limit(limit int) *episodeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e episodeDo) Offset(offset int) *episodeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e episodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *episodeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e episodeDo) Unscoped() *episodeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e episodeDo) Create(values ...*dao.Episode) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e episodeDo) CreateInBatches(values []*dao.Episode, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e episodeDo) Save(values ...*dao.Episode) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e episodeDo) First() (*dao.Episode, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Episode), nil
	}
}

func (e episodeDo) Take() (*dao.Episode, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Episode), nil
	}
}

func (e episodeDo) Last() (*dao.Episode, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Episode), nil
	}
}

func (e episodeDo) Find() ([]*dao.Episode, error) {
	result, err := e.DO.Find()
	return result.([]*dao.Episode), err
}

func (e episodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Episode, err error) {
	buf := make([]*dao.Episode, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e episodeDo) FindInBatches(result *[]*dao.Episode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e episodeDo) Attrs(attrs ...field.AssignExpr) *episodeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e episodeDo) Assign(attrs ...field.AssignExpr) *episodeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e episodeDo) Joins(fields ...field.RelationField) *episodeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e episodeDo) Preload(fields ...field.RelationField) *episodeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e episodeDo) FirstOrInit() (*dao.Episode, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Episode), nil
	}
}

func (e episodeDo) FirstOrCreate() (*dao.Episode, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Episode), nil
	}
}

func (e episodeDo) FindByPage(offset int, limit int) (result []*dao.Episode, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e episodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e *episodeDo) withDO(do gen.Dao) *episodeDo {
	e.DO = *do.(*gen.DO)
	return e
}
