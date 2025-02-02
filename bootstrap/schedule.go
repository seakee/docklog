// Copyright 2024 Seakee.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package bootstrap

import (
	"context"

	"github.com/seakee/dockmon/app/job"
	"github.com/seakee/dockmon/app/pkg/schedule"
)

func (a *App) startSchedule(ctx context.Context) {
	s := schedule.New(a.Logger, a.Redis["dockmon"], a.TraceID)

	job.Register(a.Logger, a.Redis, a.MysqlDB, a.Feishu, s)

	s.Start()

	a.Logger.Info(ctx, "Schedule loaded successfully")
}
