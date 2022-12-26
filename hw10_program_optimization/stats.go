package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domainName := fmt.Sprintf(".%s", domain)
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)
	for scanner.Scan() {
		var user User
		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, err
		}
		if strings.HasSuffix(user.Email, domainName) {
			lowerEmail := strings.ToLower(user.Email)
			splittedEmail := strings.SplitN(lowerEmail, "@", 2)
			result[splittedEmail[1]]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
