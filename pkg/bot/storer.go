package bot

import (
	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

func NewBotStore(db *bolt.DB) (*BotStore, error) {
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("bot"))
		return nil
	})
	return &BotStore{
		db: db,
	}, nil
}

type BotStore struct {
	db *bolt.DB
}

func (b *BotStore) SaveFilterID(userID id.UserID, filterID string) {
	b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bot"))
		return bucket.Put([]byte("filter"), []byte(filterID))
	})
}

func (b *BotStore) LoadFilterID(userID id.UserID) string {
	log.Debug().Str("method", "LoadFilterID").Str("userID", userID.String()).Send()
	// TODO implement me
	return ""
}

func (b *BotStore) SaveNextBatch(userID id.UserID, nextBatchToken string) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bot"))
		return bucket.Put([]byte("batch"), []byte(nextBatchToken))
	})
	if err != nil {
		log.Error().Err(err).Msg("SaveNextBatch")
	}
}

func (b *BotStore) LoadNextBatch(userID id.UserID) string {
	result := ""
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bot"))
		result = string(bucket.Get([]byte("batch")))
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("LoadNextBatch")
	}

	return result
}

func (b *BotStore) SaveRoom(room *mautrix.Room) {
	panic("not implemented") // TODO: Implement
}

func (b *BotStore) LoadRoom(roomID id.RoomID) *mautrix.Room {
	panic("not implemented") // TODO: Implement
}