package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Validator struct {
	errors map[string][]string
}

func New() *Validator {
	return &Validator{
		errors: make(map[string][]string),
	}
}

func (v *Validator) Validate(data map[string]string, rules map[string]string) bool {
	v.errors = make(map[string][]string)

	for field, ruleStr := range rules {
		value := data[field]
		ruleList := strings.Split(ruleStr, "|")

		for _, rule := range ruleList {
			if err := v.applyRule(field, value, rule); err != nil {
				v.errors[field] = append(v.errors[field], err.Error())
			}
		}
	}

	return len(v.errors) == 0
}

func (v *Validator) applyRule(field, value, rule string) error {
	parts := strings.SplitN(rule, ":", 2)
	ruleName := parts[0]
	params := ""
	if len(parts) > 1 {
		params = parts[1]
	}

	switch ruleName {
	case "required":
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required", field)
		}
	case "email":
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, value)
		if !matched {
			return fmt.Errorf("%s must be a valid email", field)
		}
	case "min":
		minLen, _ := strconv.Atoi(params)
		if len(value) < minLen {
			return fmt.Errorf("%s must be at least %d characters", field, minLen)
		}
	case "max":
		maxLen, _ := strconv.Atoi(params)
		if len(value) > maxLen {
			return fmt.Errorf("%s must not exceed %d characters", field, maxLen)
		}
	case "numeric":
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return fmt.Errorf("%s must be numeric", field)
		}
	case "integer":
		if _, err := strconv.Atoi(value); err != nil {
			return fmt.Errorf("%s must be an integer", field)
		}
	}

	return nil
}

func (v *Validator) Errors() map[string][]string {
	return v.errors
}

func (v *Validator) FirstError() string {
	for _, errs := range v.errors {
		if len(errs) > 0 {
			return errs[0]
		}
	}
	return ""
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}
