package usecase

import "github.com/anuraaaa/threadgo/domain"

type CommentUseCase struct {
	repo     domain.CommentRepository
	postRepo domain.PostRepository
}

func NewCommentUseCase(r domain.CommentRepository, pr domain.PostRepository) *CommentUseCase {
	return &CommentUseCase{repo: r, postRepo: pr}
}

type CreateCommentInput struct {
	PostID  uint
	UserID  uint
	Content string
}

func (uc *CommentUseCase) Create(in CreateCommentInput) error {
	// optionally: validate post exists
	if _, err := uc.postRepo.GetByID(in.PostID); err != nil {
		return err
	}
	c := &domain.Comment{PostID: in.PostID, UserID: in.UserID, Content: in.Content}
	return uc.repo.Create(c)
}

func (uc *CommentUseCase) List(postID uint, page, limit int) ([]domain.Comment, int64, error) {
	return uc.repo.ListByPost(postID, page, limit)
}
