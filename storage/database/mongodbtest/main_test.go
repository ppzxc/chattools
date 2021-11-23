package mongodbtest

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/mongodb"
	"testing"
	"time"
)

var (
	service database.Service
)

// ----------------------------
// 		TEST MAIN FUNCTION
// ----------------------------
func TestMain(m *testing.M) {
	s, err := mongodb.NewMongoDbInstance(context.Background(), "SCRAM-SHA-256", "admin", "localhost", 27017, "dev", "link", "dev", 10*time.Second, 5*time.Second)
	if err != nil {
		panic(err)
	}
	service = s
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = service.InitializeTable(ctx, true, true, false)
	cancel()
	if err != nil {
		panic(err)
	}
	m.Run()
}
