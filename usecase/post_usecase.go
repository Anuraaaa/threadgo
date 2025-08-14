package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/anuraaaa/threadgo/domain"
	"github.com/anuraaaa/threadgo/storage"
)

type PostUseCase struct {
	repo domain.PostRepository
	fs   storage.FileStorage
}

func NewPostUseCase(r domain.PostRepository, fs storage.FileStorage) *PostUseCase {
	return &PostUseCase{repo: r, fs: fs}
}

type CreatePostInput struct {
	UserID  uint
	Content string
	File    *multipart.FileHeader // optional
}

type PostListItem struct {
	domain.Post
	LikeCount int64 `json:"like_count"`
}

func (uc *PostUseCase) Create(in CreatePostInput) (*domain.Post, error) {
	if in.Content == "" && in.File == nil {
		return nil, errors.New("content or image required")
	}

	var imageURL *string
	if in.File != nil {
		// open and read
		f, err := in.File.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()
		data := make([]byte, in.File.Size)
		if _, err := f.Read(data); err != nil {
			return nil, err
		}
		name := fmt.Sprintf("%d_%s", time.Now().UnixNano(), in.File.Filename)
		url, err := uc.fs.Save(name, data)
		if err != nil {
			return nil, err
		}
		imageURL = &url
	}
	p := &domain.Post{UserID: in.UserID, Content: in.Content, ImageURL: imageURL}
	if err := uc.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (uc *PostUseCase) List(page, limit int) ([]domain.Post, int64, error) {
	return uc.repo.List(page, limit)
}
