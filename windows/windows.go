package windows

import (
	"fmt"
	"os"
	"os/exec"

	"cshgitlab.cn-csh.celestica.com/micro-service/truss-patch/patch"
)

func removefile(path string) {
	err := os.Remove(path + "/gengokit/template/NAME-service/svc/transport_http.gotemplate")
	if err != nil {
		fmt.Println("no file name is transport_http.gotemplate")
	}

	err = os.Remove(path + "/gengokit/template/template.go")
	if err != nil {
		fmt.Println("no file name is template.go")
	}
}

func removedir(path string) {
	err := os.RemoveAll(path + "/gengokit/template/NAME-service/svc/client/http")
	if err != nil {
		fmt.Println("no dir name is http")
	}

	err = os.RemoveAll(path + "/gengokit/template/NAME-service/cmd")
	if err != nil {
		fmt.Println("no dir name is cmd")
	}
}

func updatefile(path string) {
	clifile, err := os.OpenFile(path+"/gengokit/template/NAME-service/svc/server/cli/cli.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("cil not found")
	}
	_, err = clifile.WriteString(patch.CliContext)
	if err != nil {
		fmt.Println("cil write fail")
	}
	clifile.Close()
	runfile, err := os.OpenFile(path+"/gengokit/template/NAME-service/svc/server/run.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("run not found")
	}
	_, err = runfile.WriteString(patch.RunContext)
	if err != nil {
		fmt.Println("run write fail")
	}
	runfile.Close()

	winfile, err := os.OpenFile(path+"/wininstall.bat", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("winfile not found")
	}
	_, err = winfile.WriteString(patch.WinBat)
	if err != nil {
		fmt.Println("winfile write fail")
	}
	winfile.Close()
}

func Option(path string) {
	updatefile(path)
	removedir(path)
	removefile(path)
}

func Windows(path string) {
	//windows run
	wininstall := exec.Command(path + "/wininstall.bat")
	err := wininstall.Run()
	if err != nil {
		fmt.Println("wininstall run fail in windows")
	}
}
