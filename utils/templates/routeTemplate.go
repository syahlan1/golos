package templates

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/syahlan1/golos/utils"
)

func AddRouteAndImport(controllerName, routePath, packagePath string) error {
	filePath := "routes/routes.go"
	readFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer readFile.Close()

	var lines []string
	scanner := bufio.NewScanner(readFile)
	importAdded := false
	routeAdded := false

	for scanner.Scan() {
		line := scanner.Text()
		if !importAdded && strings.Contains(line, `import (`) {
			lines = append(lines, line)
			importLine := fmt.Sprintf(`	"%s/controllers/%sController"`, packagePath, utils.ToLowerCamelCase(controllerName))
			lines = append(lines, importLine)
			importAdded = true
			continue
		}
		if !routeAdded && strings.Contains(line, `api := app.Group("/api")`) {
			lines = append(lines, line)
			routeLines := []string{
				fmt.Sprintf(``),
				fmt.Sprintf(`	//route %s`, utils.ToKebabCase(routePath)),
				fmt.Sprintf(`	api.Post("/%s/create", %sController.Create%s)`, routePath, utils.ToLowerCamelCase(controllerName), utils.ToCamelCase(controllerName)),
				fmt.Sprintf(`	api.Put("/%s/update/:id", %sController.Update%s)`, routePath, utils.ToLowerCamelCase(controllerName), utils.ToCamelCase(controllerName)),
				fmt.Sprintf(`	api.Put("/%s/delete/:id", %sController.Delete%s)`, routePath, utils.ToLowerCamelCase(controllerName), utils.ToCamelCase(controllerName)),
				fmt.Sprintf(`	api.Get("/%s/show", %sController.Show%s)`, routePath, utils.ToLowerCamelCase(controllerName), utils.ToCamelCase(controllerName)),
				fmt.Sprintf(`	api.Get("/%s/show/:id", %sController.ShowDetail%s)`, routePath, utils.ToLowerCamelCase(controllerName), utils.ToCamelCase(controllerName)),
			}
			lines = append(lines, routeLines...)
			routeAdded = true
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}
