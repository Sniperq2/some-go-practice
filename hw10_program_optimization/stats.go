package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)
	var user User
	var lowerEmail string
	var splittedEmail []string
	var err error
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for scanner.Scan() {
		if err = json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, err
		}
		if strings.HasSuffix(user.Email, domain) {
			lowerEmail = strings.ToLower(user.Email)
			splittedEmail = strings.SplitN(lowerEmail, "@", 2)
			result[splittedEmail[1]]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
