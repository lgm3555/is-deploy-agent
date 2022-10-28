package deploy

import (
	"fmt"
	"is-deploy-agent/utils"
	"os/exec"
)

func Deploy(node int, worker string) {
	executeShell(node, worker)
}

func executeShell(node int, worker string) {
	shellPath := getShellPath(node, worker)
	cmd := exec.Command(shellPath)
	output, _ := cmd.Output()

	fmt.Println(string(output))
}

func getShellPath(node int, worker string) string {
	models := utils.GetJson()
	podLength := len(models[0].NodeList[node].PodList)

	var shellPath string
	for pods := 0; pods < podLength; pods++ {
		pod := models[0].NodeList[node].PodList[pods]
		name := pod.Name

		if isNameEqual(name, worker) {
			shellPath = pod.ShPath

			break
		}
	}

	return shellPath
}

func isNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}