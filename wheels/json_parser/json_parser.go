package jsonparser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type JSONTokenType int

const (
	JSONTokenTypeString JSONTokenType = iota
	JSONTokenTypeNumber
	JSONTokenTypeSyntax
	JSONTokenTypeBool
	JSONTokenTypeNull
)

type JSONToken struct {
	Value     string
	Loc       int
	TokenType JSONTokenType
}

func lex(jsonRaw string) ([]*JSONToken, error) {
	var tokens []*JSONToken
	var token *JSONToken
	for idx := 0; idx < len(jsonRaw); {
		// fmt.Printf("tokens: %+v", spew.Sdump(tokens))
		// space
		for unicode.IsSpace(rune(jsonRaw[idx])) {
			idx++
		}
		// syntax
		char := jsonRaw[idx]
		if char == '[' || char == ']' || char == ',' || char == '{' || char == '}' || char == ':' {
			token := &JSONToken{
				Value:     string(char),
				Loc:       idx,
				TokenType: JSONTokenTypeSyntax,
			}
			tokens = append(tokens, token)
			idx++
			continue
		}
		// string
		if jsonRaw[idx] == '"' {
			idx++
			var val string
			for jsonRaw[idx] != '"' {
				val += string(jsonRaw[idx])
				idx++
			}
			token := &JSONToken{
				Value:     val,
				Loc:       idx,
				TokenType: JSONTokenTypeString,
			}
			tokens = append(tokens, token)
			idx++
			continue
		}
		// number
		char = jsonRaw[idx]
		if isNumeric(string(char)) {
			num := string(char)
			idx++
			for isNumeric(string(jsonRaw[idx])) {
				num += string(jsonRaw[idx])
				idx++
			}
			token := &JSONToken{
				Value:     num,
				Loc:       idx,
				TokenType: JSONTokenTypeNumber,
			}
			tokens = append(tokens, token)
			continue
		}
		// bool
		token, idx = lexKeyword(jsonRaw, idx, "true", JSONTokenTypeBool)
		if token != nil {
			tokens = append(tokens, token)
			continue
		}
		token, idx = lexKeyword(jsonRaw, idx, "false", JSONTokenTypeBool)
		if token != nil {
			tokens = append(tokens, token)
			continue
		}
		// null
		token, idx = lexKeyword(jsonRaw, idx, "null", JSONTokenTypeNull)
		if token != nil {
			tokens = append(tokens, token)
			continue
		}
	}
	return tokens, nil
}

func lexKeyword(jsonRaw string, idx int, keyword string, tokenType JSONTokenType) (token *JSONToken, newIdx int) {
	origIdx := idx
	keywordIdx := 0
	for keywordIdx < len(keyword) && jsonRaw[idx] == keyword[keywordIdx] {
		idx++
		keywordIdx++
	}
	if keywordIdx == len(keyword) {
		token = &JSONToken{
			Value:     keyword,
			Loc:       idx,
			TokenType: tokenType,
		}
		newIdx = idx
		return
	} else {
		return nil, origIdx
	}
}

func isNumeric(s string) bool {
	match, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, s)
	return match
}

type JSONValueType int

const (
	JSONValueTypeString JSONValueType = iota
	JSONValueTypeNumber
	JSONValueTypeBoolean
	JSONValueTypeNull
	JSONValueTypeArray
	JSONValueTypeObject
)

type JSONValue struct {
	Str     string
	Number  int
	Boolean bool
	Arr     []*JSONValue
	Obj     map[string]*JSONValue
	Type    JSONValueType
}

func parse(tokens []*JSONToken, idx int) (newIdx int, jsonVal JSONValue, err error) {
	switch tokens[idx].TokenType {
	case JSONTokenTypeString:
		return idx + 1, JSONValue{
			Str:  tokens[idx].Value,
			Type: JSONValueTypeString,
		}, nil
	case JSONTokenTypeNumber:
		num, _ := strconv.Atoi(tokens[idx].Value)
		return idx + 1, JSONValue{
			Number: num,
			Type:   JSONValueTypeNumber,
		}, nil
	case JSONTokenTypeBool:
		boolVal := tokens[idx].Value == "true"
		return idx + 1, JSONValue{
			Boolean: boolVal,
			Type:    JSONValueTypeBoolean,
		}, nil
	case JSONTokenTypeNull:
		return idx + 1, JSONValue{
			Type: JSONValueTypeNull,
		}, nil
	case JSONTokenTypeSyntax:
		// object
		if tokens[idx].Value == "{" {
			jsonObj := make(map[string]*JSONValue)
			for idx = idx + 1; idx < len(tokens); {
				if tokens[idx].Value == "}" {
					return idx + 1, JSONValue{
						Obj:  jsonObj,
						Type: JSONValueTypeObject,
					}, nil
				}
				if tokens[idx].Value == "," {
					idx++
					continue
				}
				keyToken := tokens[idx]
				if keyToken.TokenType != JSONTokenTypeString {
					return 0, JSONValue{}, fmt.Errorf("expected `string` type, got `%v`", keyToken.TokenType)
				}
				newIdx, objKey, err := parse(tokens, idx)
				if err != nil {
					return 0, JSONValue{}, err
				}
				idx = newIdx
				if tokens[idx].TokenType != JSONTokenTypeSyntax && tokens[idx].Value != ":" {
					return 0, JSONValue{}, fmt.Errorf("expected `:`, got `%v`", tokens[idx].TokenType)
				}
				idx++
				newIdx, objVal, err := parse(tokens, idx)
				if err != nil {
					return 0, JSONValue{}, err
				}
				idx = newIdx
				jsonObj[objKey.Str] = &objVal
			}
		}
		// array
		if tokens[idx].Value == "[" {
			var jsonArr []*JSONValue
			for idx = idx + 1; idx < len(tokens); {
				if tokens[idx].Value == "]" {
					return idx + 1, JSONValue{
						Arr:  jsonArr,
						Type: JSONValueTypeArray,
					}, nil
				}
				if tokens[idx].Value == "," {
					idx++
					continue
				}
				newIdx, arrEle, err := parse(tokens, idx)
				if err != nil {
					return 0, JSONValue{}, err
				}
				idx = newIdx
				jsonArr = append(jsonArr, &arrEle)
			}
		}
	}

	return 0, JSONValue{}, errors.New("unreachable path!")
}

func prettyPrint(jsonVal JSONValue, indentLevel int) string {
	switch jsonVal.Type {
	case JSONValueTypeString:
		return fmt.Sprintf(`"%v"`, jsonVal.Str)
	case JSONValueTypeBoolean:
		return fmt.Sprintf("%v", jsonVal.Boolean)
	case JSONValueTypeNumber:
		return fmt.Sprintf("%v", jsonVal.Number)
	case JSONValueTypeNull:
		return "null"
	case JSONValueTypeObject:
		var kvs []string
		var indent string
		for ; indentLevel > 0; indentLevel-- {
			indent += "  "
		}
		for k, v := range jsonVal.Obj {
			kvs = append(kvs, fmt.Sprintf("%v\"%v\": %v", indent, k, prettyPrint(*v, indentLevel+1)))
		}
		objStr := strings.Join(kvs, ",\n")
		return fmt.Sprintf("{\n%v}", objStr)
	case JSONValueTypeArray:
		var eles []string
		for _, ele := range jsonVal.Arr {
			eles = append(eles, prettyPrint(*ele, indentLevel+1))
		}
		arrStr := strings.Join(eles, ", ")
		return fmt.Sprintf("[ %v ]", arrStr)
	}

	return ""
}
