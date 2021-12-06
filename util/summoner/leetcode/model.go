package leetcode

import (
	"fmt"
	"regexp"
	"strings"
)

type Subject struct {
	Desc string
	Url  string
	Ans  string

	AnsFuncName   string
	AnsParams     []SubjectParam
	AnsReturnType string
}

type SubjectParam struct {
	Name string
	Type string
}

func NewSubject(desc string, url string, ans string) (*Subject, error) {
	subject := &Subject{
		Desc: strings.TrimSpace(desc),
		Url:  strings.TrimSpace(url),
		Ans:  strings.TrimSpace(ans),
	}

	if err := subject.parseAns(); err != nil {
		return nil, err
	}

	return subject, nil
}

func (s *Subject) parseAns() error {
	ansArr := strings.Split(strings.TrimSpace(s.Ans), "\n")

	if len(ansArr) == 0 {
		return fmt.Errorf("ans is empty")
	}

	reg := regexp.MustCompile(`func\s+(\w+)\(([A-Za-z0-9\[\], ]+)\)\s+([A-Za-z0-9\[\]]+)`)
	submatch := reg.FindStringSubmatch(ansArr[0])
	if len(submatch) != 4 {
		return fmt.Errorf("submatch length is not equal to 4: %+v", submatch)
	}

	var paramStr string
	s.AnsFuncName, paramStr, s.AnsReturnType = submatch[1], submatch[2], submatch[3]

	paramReg := regexp.MustCompile(`(\w+)\s+([A-Za-z0-9\[\]]+)`)
	for _, item := range strings.Split(paramStr, ",") {
		paramMatchs := paramReg.FindStringSubmatch(item)
		if len(paramMatchs) != 3 {
			return fmt.Errorf("paramMatchs length is not equal to 3: %+v", paramMatchs)
		}

		s.AnsParams = append(s.AnsParams, SubjectParam{
			Name: paramMatchs[1],
			Type: paramMatchs[2],
		})
	}

	if len(ansArr) == 3 && strings.TrimSpace(ansArr[2]) == "}" {
		var body string
		switch s.AnsReturnType {
		case "int":
			body = "return 0"
		case "string":
			body = "return \"\""
		case "[]int":
			body = "return []int{}"
		}
		if body != "" {
			ansArr[1] = body
			s.Ans = strings.Join(ansArr, "\n")
		}
	}

	return nil
}
