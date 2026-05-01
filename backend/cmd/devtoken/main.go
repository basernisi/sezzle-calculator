package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jnsilvag/sezzle-calculator/backend/internal/adapters/auth"
)

func main() {
	subject := flag.String("sub", "local-dev-user", "token subject")
	expiresIn := flag.Duration("expires-in", time.Hour, "token lifetime")
	flag.Parse()

	secret := os.Getenv("JWT_SECRET")
	if strings.TrimSpace(secret) == "" {
		log.Fatal("JWT_SECRET is required")
	}

	issuer := auth.NewJWTIssuer(secret)
	token, err := issuer.IssueToken(nil, *subject)
	if err != nil {
		log.Fatalf("issue token: %v", err)
	}

	log.Printf("expires in approximately %s", expiresIn.String())
	os.Stdout.WriteString(token + "\n")
}
