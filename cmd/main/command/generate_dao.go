package command

import (
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/istrings"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/los"
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func NewGenerateDaoCommand() *cobra.Command {
	var save bool
	var takeBy []string
	var deleteBy []string
	var batchTakeBy []string

	var command = &cobra.Command{
		Use:  "dao",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var pkgName = os.Getenv("GOPACKAGE")

			// dao只能从model生成
			if "model" != pkgName {
				return
			}

			var path = os.Getenv("GOFILE")

			if "" == path {
				cmd.PrintErrf("GOFILE environment variable not set\n")

				os.Exit(1)
			}

			var fset = token.NewFileSet()
			astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if nil != err {
				cmd.PrintErrf("error parsing file: %v\n", err)

				return
			}

			ast.Inspect(astFile, func(node ast.Node) bool {
				if tp, ok := node.(*ast.TypeSpec); ok {
					if _, ok := tp.Type.(*ast.StructType); !ok {
						return true
					}

					// do generate interface file
					if err := GenerateDaoDoGenerateInterface(tp.Name.Name, save, takeBy, deleteBy, batchTakeBy); nil != err {
						cmd.PrintErrf("generate interface file failed: %v\n", err)

						return false
					}

					// do generate implementation file
					if err := GenerateDaoFiles(tp.Name.Name, save, takeBy, deleteBy, batchTakeBy); nil != err {
						cmd.PrintErrf("generate dao files failed: %v\n", err)

						return false
					}
				}

				return true
			})
		},
	}

	command.Flags().BoolVar(&save, "save", false, "Generate Save method code to file")
	command.Flags().StringSliceVar(&takeBy, "take-by", nil, "Generate TakeBy[FIELD] methods code to file")
	command.Flags().StringSliceVar(&deleteBy, "delete-by", nil, "Generate DeleteBy[FIELD] methods code to file")
	command.Flags().StringSliceVar(&batchTakeBy, "batch-take-by", nil, "Generate BatchTakeBy[FIELD] methods code to file")

	return command
}

func GenerateDaoDoGenerateInterface(modelName string, save bool, takeBy []string, deleteBy []string, batchTakeBy []string) error {
	var template = `package dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
    "github.com/Luna-CY/Golang-Project-Template/model"
)

// {{.ModelName}} Data access object for {{.ModelName}} definition
type {{.ModelName}} interface {
	Transactional
	
	{{.Methods}}
}

`

	var methods = new(strings.Builder)

	// save method
	if save {
		var saveMethodCode = `// Save{{.ModelName}} save {{.ModelName}} to db
	// if {{.ModelName}} id is 0, it will create a new record, otherwise, it will update the record
	Save{{.ModelName}}(ctx context.Context, {{.LowerModelName}} *model.{{.ModelName}}) errors.Error`

		methods.WriteString(strings.NewReplacer("{{.ModelName}}", modelName, "{{.LowerModelName}}", strings.ToLower(string(modelName[0]))+modelName[1:]).Replace(saveMethodCode))
	}

	if 0 != len(takeBy) {
		// take by methods
		var takeByList = make([]string, len(takeBy))
		for i, item := range takeBy {
			var tokens = strings.Split(item, "=")

			if 3 > len(tokens) {
				return fmt.Errorf("无效的take-by参数配置: %s，每个选项必须以=号分割三个值：字段名=字段类型=字段零值", item)
			}

			var takeByMethodCode = `

	// Take{{.ModelName}}By{{.FieldName}} get {{.ModelName}} by {{.FieldName}} from db
	// if {{.ModelName}} not found, return error type with errors.ErrorTypeRecordNotFound
	Take{{.ModelName}}By{{.FieldName}}(ctx context.Context, {{.LowerFieldName}} {{.FieldType}}, lock bool) (*model.{{.ModelName}}, errors.Error)`

			takeByList[i] = strings.NewReplacer("{{.ModelName}}", modelName, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1]).Replace(takeByMethodCode)
		}

		methods.WriteString(strings.Join(takeByList, ""))
	}

	if 0 != len(deleteBy) {
		// delete by methods
		var deleteByList = make([]string, len(deleteBy))
		for i, item := range deleteBy {
			var tokens = strings.Split(item, "=")

			if 3 > len(tokens) {
				return fmt.Errorf("无效的delete-by参数配置: %s，每个选项必须以=号分割三个值：字段名=字段类型=字段零值", item)
			}

			var deleteByMethodCode = `

	// Delete{{.ModelName}}By{{.FieldName}} delete {{.ModelName}} by {{.FieldName}} from db
	Delete{{.ModelName}}By{{.FieldName}}(ctx context.Context, {{.LowerFieldName}} {{.FieldType}}) errors.Error`

			deleteByList[i] = strings.NewReplacer("{{.ModelName}}", modelName, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1]).Replace(deleteByMethodCode)
		}

		methods.WriteString(strings.Join(deleteByList, ""))
	}

	if 0 != len(batchTakeBy) {
		// batch take by methods
		var batchTakeByList = make([]string, len(batchTakeBy))
		for i, item := range batchTakeBy {
			var tokens = strings.Split(item, "=")

			if 2 > len(tokens) {
				return fmt.Errorf("无效的batch-take-by参数配置: %s，每个选项必须以=号分割两个值：字段名=字段类型", item)
			}

			var batchTakeByMethodCode = `

	// BatchTake{{.ModelName}}By{{.FieldName}} get {{.ModelName}} by {{.FieldName}} from db
	BatchTake{{.ModelName}}By{{.FieldName}}(ctx context.Context, {{.LowerFieldName}} []{{.FieldType}}, lock bool) ([]*model.{{.ModelName}}, errors.Error)`

			batchTakeByList[i] = strings.NewReplacer("{{.ModelName}}", modelName, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1]).Replace(batchTakeByMethodCode)
		}

		methods.WriteString(strings.Join(batchTakeByList, ""))
	}

	var content = strings.NewReplacer("{{.ModelName}}", modelName, "{{.Methods}}", methods.String()).Replace(template)

	path, err := filepath.Abs(filepath.Join("..", "internal", "interface", "dao", fmt.Sprintf("%s.go", istrings.CamelCaseToUnderscore(modelName))))
	if nil != err {
		return fmt.Errorf("获取绝对路径失败: %s, err: %s", path, err)
	}

	exists, err := los.CheckPathExists(path)
	if nil != err {
		return fmt.Errorf("检查文件失败: %s, err: %s", path, err)
	}

	if !exists {
		if err := los.WriteToFile(path, content); nil != err {
			return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
		}
	}

	return nil
}

var implementCode = `package {{.LowerModelName}}

import "github.com/Luna-CY/Golang-Project-Template/internal/dao"

func New() *{{.ModelName}} {
	return &{{.ModelName}}{
		BaseDao: dao.New(),
	}
}

type {{.ModelName}} struct {
	*dao.BaseDao
}
`

var saveImplementCode = `package {{.LowerModelName}}

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
	"time"
)

func (cls *{{.ModelName}}) Save{{.ModelName}}(ctx context.Context, {{.LowerModelName}} *model.{{.ModelName}}) errors.Error {
	if nil == {{.LowerModelName}} {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.Save{{.ModelName}} {{.LowerModelName}} is nil stack %s", string(debug.Stack()))

		return errors.ErrorServerInternalError("{{.Code}}.17")
	}

	{{.LowerModelName}}.UpdateTime = pointer.New(time.Now().Unix())
	if 0 == {{.LowerModelName}}.Id {
		{{.LowerModelName}}.CreateTime = pointer.New(time.Now().Unix())

		if err := cls.GetDb(ctx).Model(new(model.{{.ModelName}})).Create(&{{.LowerModelName}}).Error; nil != err {
			logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.Save{{.ModelName}} create {{.LowerModelName}} failed, err %v, stack %s", err, string(debug.Stack()))

			return errors.ErrorServerInternalError("{{.Code}}.27")
		}

		return nil
	}

	if err := cls.GetDb(ctx).Model(new(model.{{.ModelName}})).Where("id = ?", {{.LowerModelName}}.Id).Updates(&{{.LowerModelName}}).Error; nil != err {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.Save{{.ModelName}} save {{.LowerModelName}} failed, err %v, stack %s", err, string(debug.Stack()))

		return errors.ErrorServerInternalError("{{.Code}}.36")
	}

	return nil
}
`

var takeByCode = `package {{.LowerModelName}}

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"gorm.io/gorm"
	"runtime/debug"
)

func (cls *{{.ModelName}}) Take{{.ModelName}}By{{.FieldName}}(ctx context.Context, {{.LowerFieldName}} {{.FieldType}}, lock bool) (*model.{{.ModelName}}, errors.Error) {
	if {{.DefaultValue}} == {{.LowerFieldName}} {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.TakeBy{{.FieldName}} {{.LowerFieldName}} is %v stack %s", {{.DefaultValue}}, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("{{.Code}}.17")
	}

	var session = cls.GetDb(ctx).Model(new(model.{{.ModelName}}))
	session = dao.Lock(session, lock)

	var {{.LowerModelName}} *model.{{.ModelName}}
	if err := session.Where("{{.LowerFieldName}} = ?", {{.LowerFieldName}}).Take(&{{.LowerModelName}}).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrorRecordNotFound("{{.Code}}.26")
		}

		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.Take{{.ModelName}}By{{.FieldName}} take {{.LowerModelName}} by {{.LowerFieldName}} failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("{{.Code}}.31")
	}

	return {{.LowerModelName}}, nil
}
`

var deleteByCode = `package {{.LowerModelName}}

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *{{.ModelName}}) Delete{{.ModelName}}By{{.FieldName}}(ctx context.Context, {{.LowerFieldName}} {{.FieldType}}) errors.Error {
	if {{.DefaultValue}} == {{.LowerFieldName}} {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.DeleteBy{{.FieldName}}: {{.LowerFieldName}} is %v stack %s", {{.DefaultValue}}, string(debug.Stack()))

		return errors.ErrorServerInternalError("{{.Code}}.15")
	}

	if err := cls.GetDb(ctx).Model(new(model.{{.ModelName}})).Where("{{.LowerFieldName}} = ?", {{.LowerFieldName}}).Delete(nil).Error; nil != err {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.DeleteBy{{.FieldName}}: delete {{.ModelName}} failed, err %v, stack %s", err, string(debug.Stack()))

		return errors.ErrorServerInternalError("{{.Code}}.21")
	}

	return nil
}
`

var batchTakeByCode = `package {{.LowerModelName}}

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *{{.ModelName}}) BatchTake{{.ModelName}}By{{.FieldName}}(ctx context.Context, values []{{.FieldType}}, lock bool) ([]*model.{{.ModelName}}, errors.Error) {
	if 0 == len(values) {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.BatchTakeById: values is empty stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("{{.Code}}.16")
	}

	var session = cls.GetDb(ctx).Model(new(model.{{.ModelName}}))
	session = dao.Lock(session, lock)

	var data []*model.{{.ModelName}}
	if err := session.Where("{{.LowerFieldName}} in (?)", values).Find(&data).Error; nil != err {
		logger.SugarLogger(ctx).Errorf("I.D.{{.ModelName}}.BatchTakeBy{{.FieldName}}: batch take by {{.LowerFieldName}} failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("{{.Code}}.26")
	}

	return data, nil
}
`

func GenerateDaoFiles(modelName string, save bool, takeBy []string, deleteBy []string, batchTakeBy []string) error {
	root, err := filepath.Abs(filepath.Join("..", "internal", "dao", istrings.CamelCaseToUnderscore(modelName)))
	if nil != err {
		return fmt.Errorf("获取绝对路径失败: %s, err: %s", root, err)
	}

	if err := os.MkdirAll(root, 0755); nil != err {
		return fmt.Errorf("创建文件夹失败: %s, err: %s", root, err)
	}

	var kvs = []string{"{{.ModelName}}", modelName, "{{.LowerModelName}}", strings.ToLower(string(modelName[0])) + modelName[1:]}

	{
		// generate impl.go
		var path = filepath.Join(root, "impl.go")

		if err := los.WriteToFile(path, strings.NewReplacer(kvs...).Replace(implementCode)); nil != err {
			return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
		}

	}

	if save {
		var code = fmt.Sprintf("%s.%s.%s", "ID"+istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), "S"+istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]))
		var kvs = append(kvs, "{{.Code}}", code)

		// generate impl_{{ModelName}}_save_{{ModelName}}.go
		var path = filepath.Join(root, fmt.Sprintf("impl_%s_save_%s.go", istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(modelName)))

		exists, err := los.CheckPathExists(path)
		if nil != err {
			return fmt.Errorf("检查文件失败: %s, err: %s", path, err)
		}

		if !exists {
			if err := los.WriteToFile(path, strings.NewReplacer(kvs...).Replace(saveImplementCode)); nil != err {
				return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
			}
		}
	}

	for _, item := range takeBy {
		var tokens = strings.Split(item, "=")
		if 3 > len(tokens) {
			return fmt.Errorf("无效的take-by参数配置: %s，每个选项必须以=号分割三个值：字段名=字段类型=字段零值", item)
		}

		if "" == tokens[2] {
			tokens[2] = "\"\""
		}

		var code = fmt.Sprintf("%s.%s.%s", "ID"+istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), "T"+istrings.GetUpperChars(modelName)+"B"+istrings.GetUpperChars(tokens[0])+"_"+strings.ToUpper(tokens[0][len(tokens[0])-2:]))
		var kvs = append(kvs, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1], "{{.DefaultValue}}", tokens[2], "{{.Code}}", code)

		// generate impl_{{ModelName}}_take_by_{{FieldName}}.go
		var path = filepath.Join(root, fmt.Sprintf("impl_%s_take_%s_by_%s.go", istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(tokens[0])))

		exists, err := los.CheckPathExists(path)
		if nil != err {
			return fmt.Errorf("检查文件失败: %s, err: %s", path, err)
		}

		if !exists {
			if err := los.WriteToFile(path, strings.NewReplacer(kvs...).Replace(takeByCode)); nil != err {
				return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
			}
		}
	}

	for _, item := range deleteBy {
		var tokens = strings.Split(item, "=")
		if 3 > len(tokens) {
			return fmt.Errorf("无效的delete-by参数配置: %s，每个选项必须以=号分割三个值：字段名=字段类型=字段零值", item)
		}

		if "" == tokens[2] {
			tokens[2] = "\"\""
		}

		var code = fmt.Sprintf("%s.%s.%s", "ID"+istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), "D"+istrings.GetUpperChars(modelName)+"B"+istrings.GetUpperChars(tokens[0])+"_"+strings.ToUpper(tokens[0][len(tokens[0])-2:]))
		var kvs = append(kvs, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1], "{{.DefaultValue}}", tokens[2], "{{.Code}}", code)

		// generate impl_{{ModelName}}_delete_by_{{FieldName}}.go
		var path = filepath.Join(root, fmt.Sprintf("impl_%s_delete_%s_by_%s.go", istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(tokens[0])))

		exists, err := los.CheckPathExists(path)
		if nil != err {
			return fmt.Errorf("检查文件失败: %s, err: %s", path, err)
		}

		if !exists {
			if err := los.WriteToFile(path, strings.NewReplacer(kvs...).Replace(deleteByCode)); nil != err {
				return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
			}
		}
	}

	for _, item := range batchTakeBy {
		var tokens = strings.Split(item, "=")
		if 2 > len(tokens) {
			return fmt.Errorf("无效的batch-take-by参数配置: %s，每个选项必须以=号分割两个值：字段名=字段类型", item)
		}

		var code = fmt.Sprintf("%s.%s.%s", "ID"+istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), istrings.GetUpperChars(modelName)+"_"+strings.ToUpper(modelName[len(modelName)-2:]), "BT"+istrings.GetUpperChars(modelName)+"B"+istrings.GetUpperChars(tokens[0])+"_"+strings.ToUpper(tokens[0][len(tokens[0])-2:]))
		var kvs = append(kvs, "{{.FieldName}}", tokens[0], "{{.LowerFieldName}}", strings.ToLower(tokens[0]), "{{.FieldType}}", tokens[1], "{{.Code}}", code)

		// generate impl_{{ModelName}}_batch_take_by_{{FieldName}}.go
		var path = filepath.Join(root, fmt.Sprintf("impl_%s_batch_take_%s_by_%s.go", istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(modelName), istrings.CamelCaseToUnderscore(tokens[0])))

		exists, err := los.CheckPathExists(path)
		if nil != err {
			return fmt.Errorf("检查文件失败: %s, err: %s", path, err)
		}

		if !exists {
			if err := los.WriteToFile(path, strings.NewReplacer(kvs...).Replace(batchTakeByCode)); nil != err {
				return fmt.Errorf("写入文件失败: %s, err: %s", path, err)
			}
		}
	}

	return nil
}
