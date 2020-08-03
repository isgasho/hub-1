// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"context"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/tektoncd/hub/api/gen/auth"
	"github.com/tektoncd/hub/api/pkg/app"
)

type service struct {
	api           app.Config
	logger        *zap.SugaredLogger
	db            *gorm.DB
	oauth         *oauth2.Config
	jwtSigningKey string
}

// New returns the auth service implementation.
func New(api app.Config) auth.Service {
	return &service{api, api.Logger(), api.DB(), api.OAuthConfig(), api.JWTSigningKey()}
}

// Authenticates users against GitHub OAuth
func (s *service) Authenticate(ctx context.Context, p *auth.AuthenticatePayload) (res *auth.AuthenticateResult, err error) {
	s.logger.Info("auth.Authenticate")
	return
}