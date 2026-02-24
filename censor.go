package main

import (
	"strings"
)

func censor(body string) string {
	words := strings.Split(body, " ")
	for i := range words {
		lower := strings.ToLower(words[i])
		if lower == "kerfuffle" || lower == "sharbert" || lower == "fornax" {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
