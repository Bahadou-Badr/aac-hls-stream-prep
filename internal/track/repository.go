package track

import "sync"

type Repository interface {
	Save(track *Track) error
	GetByID(id string) (*Track, error)
}

type InMemoryRepository struct {
	mu     sync.RWMutex
	tracks map[string]*Track
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tracks: make(map[string]*Track),
	}
}

func (r *InMemoryRepository) Save(track *Track) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tracks[track.ID] = track
	return nil
}

func (r *InMemoryRepository) GetByID(id string) (*Track, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	track, ok := r.tracks[id]
	if !ok {
		return nil, ErrNotFound
	}

	return track, nil
}
