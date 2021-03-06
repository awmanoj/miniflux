// Copyright 2018 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package scheduler // import "miniflux.app/service/scheduler"

import (
	"time"

	"miniflux.app/config"
	"miniflux.app/logger"
	"miniflux.app/model"
	"miniflux.app/storage"
	"miniflux.app/worker"
)

// Serve starts the internal scheduler.
func Serve(store *storage.Storage, pool *worker.Pool) {
	logger.Info(`Starting scheduler...`)

	go feedScheduler(
		store,
		pool,
		config.Opts.PollingFrequency(),
		config.Opts.BatchSize(),
	)

	go cleanupScheduler(
		store,
		config.Opts.CleanupFrequencyHours(),
		config.Opts.CleanupArchiveReadDays(),
		config.Opts.CleanupArchiveUnreadDays(),
		config.Opts.CleanupRemoveSessionsDays(),
	)
}

func feedScheduler(store *storage.Storage, pool *worker.Pool, frequency, batchSize int) {
	c := time.Tick(time.Duration(frequency) * time.Minute)
	for range c {
		jobs, err := store.NewBatch(batchSize)
		if err != nil {
			logger.Error("[Scheduler:Feed] %v", err)
		} else {
			logger.Debug("[Scheduler:Feed] Pushing %d jobs", len(jobs))
			pool.Push(jobs)
		}
	}
}

func cleanupScheduler(store *storage.Storage, frequency, archiveReadDays, archiveUnreadDays, sessionsDays int) {
	c := time.Tick(time.Duration(frequency) * time.Hour)
	for range c {
		nbSessions := store.CleanOldSessions(sessionsDays)
		nbUserSessions := store.CleanOldUserSessions(sessionsDays)
		logger.Info("[Scheduler:Cleanup] Cleaned %d sessions and %d user sessions", nbSessions, nbUserSessions)

		if rowsAffected, err := store.ArchiveEntries(model.EntryStatusRead, archiveReadDays); err != nil {
			logger.Error("[Scheduler:ArchiveReadEntries] %v", err)
		} else {
			logger.Info("[Scheduler:ArchiveReadEntries] %d entries changed", rowsAffected)
		}

		if rowsAffected, err := store.ArchiveEntries(model.EntryStatusUnread, archiveUnreadDays); err != nil {
			logger.Error("[Scheduler:ArchiveUnreadEntries] %v", err)
		} else {
			logger.Info("[Scheduler:ArchiveUnreadEntries] %d entries changed", rowsAffected)
		}
	}
}
