package model

import (
	"github.com/MenciusCheng/superman/util/gendbinfo/config"
	"github.com/MenciusCheng/superman/util/gendbinfo/genstruct"
	"github.com/MenciusCheng/superman/util/gendbinfo/mybigcamel"
	"strings"
)

type _Model struct {
	info DBInfo
	pkg  *genstruct.GenPackage
}

// Generate build code string.生成代码
func Generate(info DBInfo) (out []GenOutInfo, m _Model) {
	m = _Model{
		info: info,
	}

	// struct
	if config.GetIsOutFileByTableName() {
		outByTable := m.GenerateByTableName()
		out = append(out, outByTable...)
	} else {
		var stt GenOutInfo
		stt.FileCtx = m.generate()
		stt.FileName = info.DbName + ".go"

		if name := config.GetOutFileName(); len(name) > 0 {
			stt.FileName = name + ".go"
		}
		out = append(out, stt)
	}

	// ------end

	// gen function
	//if config.GetIsOutFunc() {
	//	out = append(out, m.generateFunc()...)
	//}
	// -------------- end
	return
}

// GetPackage gen struct on table
func (m *_Model) GetPackage() genstruct.GenPackage {
	if m.pkg == nil {
		var pkg genstruct.GenPackage
		pkg.SetPackage(m.info.PackageName) //package name

		tablePrefix := config.GetTablePrefix()

		for _, tab := range m.info.TabList {
			var sct genstruct.GenStruct

			sct.SetTableName(tablePrefix + tab.Name)

			//如果设置了表前缀
			// if tablePrefix != "" {
			// 	tab.Name = strings.TrimLeft(tab.Name, tablePrefix)
			// }

			sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
			sct.SetNotes(tab.Notes)
			sct.AddElement(m.genTableElement(tab.Em)...) // build element.构造元素
			sct.SetCreatTableStr(tab.SQLBuildStr)
			pkg.AddStruct(sct)
		}
		m.pkg = &pkg
	}

	return *m.pkg
}

// GetPackageByTableName Generate multiple model files based on the table name. 根据表名生成多个model文件
func (m *_Model) GenerateByTableName() (out []GenOutInfo) {
	if m.pkg == nil {
		for _, tab := range m.info.TabList {
			var pkg genstruct.GenPackage
			pkg.SetPackage(m.info.PackageName) //package name
			var sct genstruct.GenStruct
			sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
			sct.SetNotes(tab.Notes)
			sct.AddElement(m.genTableElement(tab.Em)...) // build element.构造元素
			sct.SetCreatTableStr(tab.SQLBuildStr)
			sct.SetTableName(tab.Name)
			pkg.AddStruct(sct)
			var stt GenOutInfo
			stt.FileCtx = pkg.Generate()
			stt.FileName = tab.Name + ".go"
			out = append(out, stt)
		}
	}
	return
}

func (m *_Model) generate() string {
	m.pkg = nil
	m.GetPackage()
	return m.pkg.Generate()
}

// genTableElement Get table columns and comments.获取表列及注释
func (m *_Model) genTableElement(cols []ColumnsInfo) (el []genstruct.GenElement) {
	_tagGorm := config.GetDBTag()
	_tagJSON := config.GetURLTag()

	for _, v := range cols {
		var tmp genstruct.GenElement
		var isPK bool
		if strings.EqualFold(v.Type, "gorm.Model") { // gorm model
			tmp.SetType(v.Type) //
		} else {
			tmp.SetName(getCamelName(v.Name))
			tmp.SetNotes(v.Notes)
			tmp.SetType(getTypeName(v.Type, v.IsNull))
			// 是否输出gorm标签
			if len(_tagGorm) > 0 {
				// not simple output. 默认只输出gorm主键和字段标签
				if !config.GetSimple() {
					for _, v1 := range v.Index {
						switch v1.Key {
						// case ColumnsKeyDefault:
						case ColumnsKeyPrimary: // primary key.主键
							tmp.AddTag(_tagGorm, "primaryKey")
							isPK = true
						case ColumnsKeyUnique: // unique key.唯一索引
							tmp.AddTag(_tagGorm, "unique")
						case ColumnsKeyIndex: // index key.复合索引
							uninStr := getUninStr("index", ":", v1.KeyName)
							// 兼容 gorm 本身 sort 标签
							if v1.KeyName == "sort" {
								uninStr = "index"
							}
							if v1.KeyType == "FULLTEXT" {
								uninStr += ",class:FULLTEXT"
							}
							tmp.AddTag(_tagGorm, uninStr)
						case ColumnsKeyUniqueIndex: // unique index key.唯一复合索引
							tmp.AddTag(_tagGorm, getUninStr("uniqueIndex", ":", v1.KeyName))
						}
					}
				} else {
					for _, v1 := range v.Index {
						switch v1.Key {
						// case ColumnsKeyDefault:
						case ColumnsKeyPrimary: // primary key.主键
							tmp.AddTag(_tagGorm, "primaryKey")
							isPK = true
						}
					}
				}
			}
		}

		if len(v.Name) > 0 {
			// 是否输出gorm标签
			if len(_tagGorm) > 0 {
				// not simple output
				if !config.GetSimple() {
					tmp.AddTag(_tagGorm, "column:"+v.Name)
					tmp.AddTag(_tagGorm, "type:"+v.Type)
					if !v.IsNull {
						tmp.AddTag(_tagGorm, "not null")
					}
					// default tag
					if len(v.Gormt) > 0 {
						tmp.AddTag(_tagGorm, v.Gormt)
					}
				} else {
					tmp.AddTag(_tagGorm, "column:"+v.Name)
				}
			}

			// json tag
			if config.GetIsWEBTag() {
				if isPK && config.GetIsWebTagPkHidden() {
					tmp.AddTag(_tagJSON, "-")
				} else {
					if config.GetWebTagType() == 0 {
						tmp.AddTag(_tagJSON, mybigcamel.UnSmallMarshal(mybigcamel.Marshal(v.Name)))
					} else {
						tmp.AddTag(_tagJSON, mybigcamel.UnMarshal(v.Name))
					}
				}
			}

		}

		tmp.ColumnName = v.Name // 列名
		el = append(el, tmp)

		// ForeignKey
		//if config.GetIsForeignKey() && len(v.ForeignKeyList) > 0 {
		//	fklist := m.genForeignKey(v)
		//	el = append(el, fklist...)
		//}
		// -----------end
	}

	return
}
