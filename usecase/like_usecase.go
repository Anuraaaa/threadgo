package usecase

import "github.com/anuraaaa/threadgo/domain"

type LikeUseCase struct {
	repo     domain.LikeRepository
	postRepo domain.PostRepository
}

func NewLikeUseCase(r domain.LikeRepository, pr domain.PostRepository) *LikeUseCase {
	return &LikeUseCase{repo: r, postRepo: pr}
}

func (uc *LikeUseCase) Like(userID, postID uint) error {
	// ensure post exists
	if _, err := uc.postRepo.GetByID(postID); err != nil {
		return err
	}
	l := &domain.Like{UserID: userID, PostID: postID}
	return uc.repo.Create(l)
}

func (uc *LikeUseCase) Unlike(userID, postID uint) error {
	return uc.repo.Delete(userID, postID)
}

func (uc *LikeUseCase) Count(postID uint) (int64, error) { return uc.repo.Count(postID) }
func (uc *LikeUseCase) IsLiked(userID, postID uint) (bool, error) {
	return uc.repo.IsLiked(userID, postID)
}
