// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

func (server *Server) analyzeRootCaluse(name string) error {

	q := metric.NewAnalyzeQuery()
	q.Target = name

	rs, err := server.metricMgr.Query(q)
	if err != nil {
		return err
	}

	//rootContainer := map[string]interface{}{}

	ms := rs.GetFirstMetrics()
	n := 0
	firstResults := []float64{}
	for ms != nil {
		results := []float64{}
		for n, v := range ms.Values {
			switch n {
			case 0:
				results = append(results, v.Value)
			case 1:
				results = append(results, v.Value)
				results = append(results, float64(v.Timestamp.Unix()))
			}
		}
		if n == 0 {
			copy(firstResults, results)
		}
		if 3 <= len(results) {
			if (0.6 < results[0]) && (firstResults[2] < results[2]) {
				logging.Info("%s : %f, %f", ms.Name, results[0], results[1])
			}
		}
		ms = rs.GetNextMetrics()
		n++
	}

	return nil
}

func (server *Server) analyzeFormula(formula kb.Formula) error {
	name := formula.GetVariable().GetName()
	server.analyzeRootCaluse(name)
	return nil
}

func (server *Server) analyzeClause(clause kb.Clause) error {
	for _, formula := range clause.GetFormulas() {
		ok, err := formula.IsSatisfied()
		if err != nil {
			return err
		}
		if ok {
			continue
		}

		err = server.analyzeFormula(formula)
		if err != nil {
			return err
		}
	}

	return nil
}

func (server *Server) analyzeQos(rule kb.Rule) error {
	clauses := rule.GetClauses()

	for _, clause := range clauses {
		ok, err := clause.IsSatisfied()
		if err != nil {
			return err
		}
		if ok {
			continue
		}

		err = server.analyzeClause(clause)
		if err != nil {
			return err
		}
	}

	return nil
}
