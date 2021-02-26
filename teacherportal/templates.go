package teacherportal

import (
	"html/template"
)

var rootTemplate *template.Template

// ImportTemplates imports templates and parse them
func ImportTemplates() error {
	var err error

	rootTemplate, err = template.ParseFiles(
		"github.com/zechtz/distributed/teacherportal/public/students.gohtml",
		"github.com/zechtz/distributed/teacherportal/public/student.gohtml",
	)

	if err != nil {
		return err
	}
	return nil
}
