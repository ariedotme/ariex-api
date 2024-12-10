package database

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func InitSupabase() {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	Client = supabase.CreateClient(url, key)
}
