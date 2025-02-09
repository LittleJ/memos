package store

import (
	"context"
)

type MemoRelationType string

const (
	MemoRelationReference  MemoRelationType = "REFERENCE"
	MemoRelationAdditional MemoRelationType = "ADDITIONAL"
)

type MemoRelation struct {
	MemoID        int32
	RelatedMemoID int32
	Type          MemoRelationType
}

type FindMemoRelation struct {
	MemoID        *int32
	RelatedMemoID *int32
	Type          *MemoRelationType
}

type DeleteMemoRelation struct {
	MemoID        *int32
	RelatedMemoID *int32
	Type          *MemoRelationType
}

func (s *Store) UpsertMemoRelation(ctx context.Context, create *MemoRelation) (*MemoRelation, error) {
	return s.driver.UpsertMemoRelation(ctx, create)
}

func (s *Store) ListMemoRelations(ctx context.Context, find *FindMemoRelation) ([]*MemoRelation, error) {
	return s.driver.ListMemoRelations(ctx, find)
}

func (s *Store) GetMemoRelation(ctx context.Context, find *FindMemoRelation) (*MemoRelation, error) {
	list, err := s.ListMemoRelations(ctx, find)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

func (s *Store) DeleteMemoRelation(ctx context.Context, delete *DeleteMemoRelation) error {
	return s.driver.DeleteMemoRelation(ctx, delete)
}
