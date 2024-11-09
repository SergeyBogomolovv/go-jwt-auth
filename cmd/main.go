package main

import (
	"fmt"

	"github.com/SergeyBogomolovv/go-jwt-auth/internal/config"
)

func main() {
	cfg := config.New()
	fmt.Println(cfg)
}
