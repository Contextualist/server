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

func newCharacter(db *gorm.DB) character {
	_character := character{}

	_character.characterDo.UseDB(db)
	_character.characterDo.UseModel(&dao.Character{})

	tableName := _character.characterDo.TableName()
	_character.ALL = field.NewField(tableName, "*")
	_character.ID = field.NewUint32(tableName, "crt_id")
	_character.Name = field.NewString(tableName, "crt_name")
	_character.Role = field.NewUint8(tableName, "crt_role")
	_character.Infobox = field.NewString(tableName, "crt_infobox")
	_character.Summary = field.NewString(tableName, "crt_summary")
	_character.Img = field.NewString(tableName, "crt_img")
	_character.Comment = field.NewUint32(tableName, "crt_comment")
	_character.Collects = field.NewUint32(tableName, "crt_collects")
	_character.Dateline = field.NewUint32(tableName, "crt_dateline")
	_character.Lastpost = field.NewUint32(tableName, "crt_lastpost")
	_character.Lock = field.NewInt8(tableName, "crt_lock")
	_character.ImgAnidb = field.NewString(tableName, "crt_img_anidb")
	_character.AnidbID = field.NewUint32(tableName, "crt_anidb_id")
	_character.Ban = field.NewUint8(tableName, "crt_ban")
	_character.Redirect = field.NewUint32(tableName, "crt_redirect")
	_character.Nsfw = field.NewBool(tableName, "crt_nsfw")
	_character.Fields = characterHasOneFields{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fields", "dao.PersonField"),
	}

	_character.fillFieldMap()

	return _character
}

type character struct {
	characterDo characterDo

	ALL      field.Field
	ID       field.Uint32
	Name     field.String
	Role     field.Uint8
	Infobox  field.String
	Summary  field.String
	Img      field.String
	Comment  field.Uint32
	Collects field.Uint32
	Dateline field.Uint32
	Lastpost field.Uint32
	Lock     field.Int8
	ImgAnidb field.String
	AnidbID  field.Uint32
	Ban      field.Uint8
	Redirect field.Uint32
	Nsfw     field.Bool
	Fields   characterHasOneFields

	fieldMap map[string]field.Expr
}

func (c character) Table(newTableName string) *character {
	c.characterDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c character) As(alias string) *character {
	c.characterDo.DO = *(c.characterDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *character) updateTableName(table string) *character {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewUint32(table, "crt_id")
	c.Name = field.NewString(table, "crt_name")
	c.Role = field.NewUint8(table, "crt_role")
	c.Infobox = field.NewString(table, "crt_infobox")
	c.Summary = field.NewString(table, "crt_summary")
	c.Img = field.NewString(table, "crt_img")
	c.Comment = field.NewUint32(table, "crt_comment")
	c.Collects = field.NewUint32(table, "crt_collects")
	c.Dateline = field.NewUint32(table, "crt_dateline")
	c.Lastpost = field.NewUint32(table, "crt_lastpost")
	c.Lock = field.NewInt8(table, "crt_lock")
	c.ImgAnidb = field.NewString(table, "crt_img_anidb")
	c.AnidbID = field.NewUint32(table, "crt_anidb_id")
	c.Ban = field.NewUint8(table, "crt_ban")
	c.Redirect = field.NewUint32(table, "crt_redirect")
	c.Nsfw = field.NewBool(table, "crt_nsfw")

	c.fillFieldMap()

	return c
}

func (c *character) WithContext(ctx context.Context) *characterDo {
	return c.characterDo.WithContext(ctx)
}

func (c character) TableName() string { return c.characterDo.TableName() }

func (c character) Alias() string { return c.characterDo.Alias() }

func (c *character) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *character) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 17)
	c.fieldMap["crt_id"] = c.ID
	c.fieldMap["crt_name"] = c.Name
	c.fieldMap["crt_role"] = c.Role
	c.fieldMap["crt_infobox"] = c.Infobox
	c.fieldMap["crt_summary"] = c.Summary
	c.fieldMap["crt_img"] = c.Img
	c.fieldMap["crt_comment"] = c.Comment
	c.fieldMap["crt_collects"] = c.Collects
	c.fieldMap["crt_dateline"] = c.Dateline
	c.fieldMap["crt_lastpost"] = c.Lastpost
	c.fieldMap["crt_lock"] = c.Lock
	c.fieldMap["crt_img_anidb"] = c.ImgAnidb
	c.fieldMap["crt_anidb_id"] = c.AnidbID
	c.fieldMap["crt_ban"] = c.Ban
	c.fieldMap["crt_redirect"] = c.Redirect
	c.fieldMap["crt_nsfw"] = c.Nsfw

}

func (c character) clone(db *gorm.DB) character {
	c.characterDo.ReplaceDB(db)
	return c
}

type characterHasOneFields struct {
	db *gorm.DB

	field.RelationField
}

func (a characterHasOneFields) Where(conds ...field.Expr) *characterHasOneFields {
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

func (a characterHasOneFields) WithContext(ctx context.Context) *characterHasOneFields {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a characterHasOneFields) Model(m *dao.Character) *characterHasOneFieldsTx {
	return &characterHasOneFieldsTx{a.db.Model(m).Association(a.Name())}
}

type characterHasOneFieldsTx struct{ tx *gorm.Association }

func (a characterHasOneFieldsTx) Find() (result *dao.PersonField, err error) {
	return result, a.tx.Find(&result)
}

func (a characterHasOneFieldsTx) Append(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a characterHasOneFieldsTx) Replace(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a characterHasOneFieldsTx) Delete(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a characterHasOneFieldsTx) Clear() error {
	return a.tx.Clear()
}

func (a characterHasOneFieldsTx) Count() int64 {
	return a.tx.Count()
}

type characterDo struct{ gen.DO }

//Get a Character from database.
//
//where(crt_id=@id)
func (c characterDo) Get(id uint32) (result *dao.Character, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("crt_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (c characterDo) Debug() *characterDo {
	return c.withDO(c.DO.Debug())
}

func (c characterDo) WithContext(ctx context.Context) *characterDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c characterDo) Clauses(conds ...clause.Expression) *characterDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c characterDo) Returning(value interface{}, columns ...string) *characterDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c characterDo) Not(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c characterDo) Or(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c characterDo) Select(conds ...field.Expr) *characterDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c characterDo) Where(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c characterDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *characterDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c characterDo) Order(conds ...field.Expr) *characterDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c characterDo) Distinct(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c characterDo) Omit(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c characterDo) Join(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c characterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c characterDo) RightJoin(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c characterDo) Group(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c characterDo) Having(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c characterDo) Limit(limit int) *characterDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c characterDo) Offset(offset int) *characterDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c characterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *characterDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c characterDo) Unscoped() *characterDo {
	return c.withDO(c.DO.Unscoped())
}

func (c characterDo) Create(values ...*dao.Character) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c characterDo) CreateInBatches(values []*dao.Character, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c characterDo) Save(values ...*dao.Character) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c characterDo) First() (*dao.Character, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Character), nil
	}
}

func (c characterDo) Take() (*dao.Character, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Character), nil
	}
}

func (c characterDo) Last() (*dao.Character, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Character), nil
	}
}

func (c characterDo) Find() ([]*dao.Character, error) {
	result, err := c.DO.Find()
	return result.([]*dao.Character), err
}

func (c characterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Character, err error) {
	buf := make([]*dao.Character, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c characterDo) FindInBatches(result *[]*dao.Character, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c characterDo) Attrs(attrs ...field.AssignExpr) *characterDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c characterDo) Assign(attrs ...field.AssignExpr) *characterDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c characterDo) Joins(fields ...field.RelationField) *characterDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c characterDo) Preload(fields ...field.RelationField) *characterDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c characterDo) FirstOrInit() (*dao.Character, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Character), nil
	}
}

func (c characterDo) FirstOrCreate() (*dao.Character, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Character), nil
	}
}

func (c characterDo) FindByPage(offset int, limit int) (result []*dao.Character, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c characterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *characterDo) withDO(do gen.Dao) *characterDo {
	c.DO = *do.(*gen.DO)
	return c
}
