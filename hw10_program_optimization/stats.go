package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"-"`
	Email    string
	Phone    string `json:"-"`
	Password string `json:"-"`
	Address  string `json:"-"`
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	scanner := bufio.NewScanner(r)
	counter := 0
	for scanner.Scan() {
		var user User
		if err = json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return
		}
		result[counter] = user
		counter++
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	domainName := fmt.Sprintf(".%s", domain)
	for _, user := range u {
		if strings.HasSuffix(user.Email, domainName) {
			lowerEmail := strings.ToLower(user.Email)
			splittedEmail := strings.SplitN(lowerEmail, "@", 2)
			result[splittedEmail[1]]++
		}
	}
	return result, nil
}
