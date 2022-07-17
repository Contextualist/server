// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package query

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/bangumi/server/internal/app/internal/cachekey"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
)

func (q Query) GetPerson(ctx context.Context, personID model.PersonID) (model.Person, error) {
	person, err := q.getPerson(ctx, personID)
	if err != nil {
		return model.Person{}, err
	}

	return person, nil
}

func (q Query) getPerson(ctx context.Context, id model.PersonID) (model.Person, error) {
	var key = cachekey.Person(id)

	// try to read from cache
	var r model.Person
	ok, err := q.cache.Get(ctx, key, &r)
	if err != nil {
		return r, errgo.Wrap(err, "cache.Get")
	}

	if ok {
		return r, nil
	}

	r, err = q.person.Get(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return r, domain.ErrPersonNotFound
		}

		return r, errgo.Wrap(err, "personRepo.Get")
	}

	if e := q.cache.Set(ctx, key, r, time.Minute); e != nil {
		q.log.Error("can't set response to cache", zap.Error(e))
	}

	return r, nil
}
