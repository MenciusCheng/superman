package gendbinfo

import (
	"fmt"
	"github.com/MenciusCheng/superman/util/gendbinfo/config"
	"github.com/MenciusCheng/superman/util/gendbinfo/genmysql"
	"github.com/MenciusCheng/superman/util/gendbinfo/model"
	"github.com/MenciusCheng/superman/util/gendbinfo/tools"
	"os/exec"
)

func Gen(orm genmysql.MySqlDB) {
	modeldb := genmysql.New(orm)

	pkg := modeldb.GenModel()

	list, _ := model.Generate(pkg)

	for _, v := range list {
		path := config.GetOutDir() + "/" + v.FileName
		tools.WriteFile(path, []string{v.FileCtx}, true)

		cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
		fmt.Println(string(cmd))

		cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
		fmt.Println(string(cmd))
	}
}
