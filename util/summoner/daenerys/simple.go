package daenerys

import (
	"io"
	"text/template"
)

type SimpleHandler struct {
	HandlerName string `json:"handler_name"`
}

func (s *SimpleHandler) Template() string {
	t := `
	s.ANY("/api/commodity/{{.HandlerName}}", {{.HandlerName}}) //

func {{.HandlerName}}(c *httpserver.Context) {
	req := model.{{.HandlerName}}Req{}
	atom := model.Atom{}
	err := c.Bind(c.Request, &req, &atom)
	if err != nil {
		c.JSONAbort(nil, err)
		return
	}
	resp, err := svc.{{.HandlerName}}(c.Ctx, req, atom)
	if err != nil {
		c.JSONAbort(nil, err)
		return
	}
	c.JSON(resp, nil)
}

func (s *Service) {{.HandlerName}}(ctx context.Context, req model.{{.HandlerName}}Req) (model.{{.HandlerName}}Resp, error) {
}

type {{.HandlerName}}Req struct {
}

type {{.HandlerName}}Resp struct {
}
`
	return t
}

func (s *SimpleHandler) Gen(wr io.Writer) error {
	tmpl, err := template.New("").Parse(s.Template())
	if err != nil {
		return err
	}

	if err := tmpl.Execute(wr, s); err != nil {
		return err
	}
	return nil
}
