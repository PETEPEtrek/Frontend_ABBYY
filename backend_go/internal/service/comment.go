package service

import (
	"backend_go/internal/entity"
)

type CommentService struct {
	commentRepo entity.CommentRepository
}

func NewCommentService(commentRepo entity.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) Get(id uint) (*entity.Comment, error) {
	commentDB, err := s.commentRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return commentDB, nil
	}
}

func (s *CommentService) Update(comment *entity.Comment) error {
	_, err := s.commentRepo.Get(comment.ID)
	if err != nil {
		return err
	}
	err = s.commentRepo.Update(comment)
	return err
}

func (s *CommentService) Delete(comment *entity.Comment) error {
	err := s.commentRepo.Delete(comment.ID)
	return err
}
