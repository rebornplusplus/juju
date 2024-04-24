// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package base

import (
	"sort"
	"time"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
)

// ControllerBases returns the supported workload bases available to it at the
// execution time.
func ControllerBases(now time.Time, requestedBase Base, imageStream string) ([]Base, error) {
	return seriesToBase(requestedBase, func(requestedSeries string) (set.Strings, error) {
		// The DistroInfo is currently in series, so we need to convert it to
		// base to get the correct series.
		return ControllerSeries(now, requestedSeries, imageStream)
	})
}

// WorkloadBases returns the supported workload bases available to it at the
// execution time.
func WorkloadBases(now time.Time, requestedBase Base, imageStream string) ([]Base, error) {
	return seriesToBase(requestedBase, func(requestedSeries string) (set.Strings, error) {
		// The DistroInfo is currently in series, so we need to convert it to
		// base to get the correct series.
		return WorkloadSeries(now, requestedSeries, imageStream)
	})
}

func seriesToBase(requestedBase Base, fn func(string) (set.Strings, error)) ([]Base, error) {
	var requestedSeries string
	if !requestedBase.Empty() {
		var err error
		requestedSeries, err = GetSeriesFromBase(requestedBase)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}

	series, err := fn(requestedSeries)
	if err != nil {
		return nil, errors.Trace(err)
	}

	bases := map[Base]struct{}{}
	for _, s := range series.Values() {
		b, err := GetBaseFromSeries(s)
		if err != nil {
			return nil, errors.Trace(err)
		}
		bases[b] = struct{}{}
	}

	results := make([]Base, 0, len(bases))
	for b := range bases {
		results = append(results, b)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].String() < results[j].String()
	})
	return results, nil
}
