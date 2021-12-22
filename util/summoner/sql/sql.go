package sql

import (
	"bytes"
	"fmt"
	"text/template"
)

type Data struct {
	Name int `json:"name"`
}

func genPostSql() {
	res := bytes.Buffer{}

	for i := 0; i < 64; i++ {
		t := "ALTER TABLE `post_{{.Name}}`\n" + "ADD COLUMN `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '帖子类型 1动态 2视频 3文章' AFTER `campid`;\n"
		tmpl, err := template.New("").Parse(t)
		if err != nil {
			panic(err)
		}

		wr := &bytes.Buffer{}
		d := Data{Name: i}
		err = tmpl.Execute(wr, d)
		if err != nil {
			panic(err)
		}
		res.Write(wr.Bytes())
	}

	fmt.Println(res.String())
}
