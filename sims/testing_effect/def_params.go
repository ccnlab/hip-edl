// Copyright (c) 2020, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/emer/emergent/params"

// ParamSets is the default set of parameters -- Base is always applied, and others can be optionally
// selected to apply on top of that
var ParamSets = params.Sets{
	{Name: "Base", Desc: "these are the best params", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: "Prjn", Desc: "keeping default params for generic prjns",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "true",
					"Prjn.Learn.Norm.On":     "true",
					"Prjn.Learn.WtBal.On":    "false",
				}},
			{Sel: ".EcCa1Prjn", Desc: "encoder projections -- no norm, moment",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.04",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
				}},
			{Sel: ".HippoCHL", Desc: "hippo CHL projections -- no norm, moment, but YES wtbal = sig better",
				Params: params.Params{
					"Prjn.CHL.Hebb":          "0.01",
					"Prjn.Learn.Lrate":       "0.2",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
				}},
			{Sel: ".PPath", Desc: "performant path, new Dg error-driven EcCa1Prjn prjns",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.15",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
				}},
			{Sel: "#CA1ToECout", Desc: "extra strong from CA1 to ECout",
				Params: params.Params{
					"Prjn.WtScale.Abs": "4.0",
				}},
			{Sel: "#InputToECin", Desc: "one-to-one input to EC",
				Params: params.Params{
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.8",
					"Prjn.WtInit.Var":  "0.0",
				}},
			{Sel: "#ECoutToECin", Desc: "one-to-one out to in",
				Params: params.Params{
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.9",
					"Prjn.WtInit.Var":  "0.01",
					"Prjn.WtScale.Rel": "0.5",
				}},
			{Sel: "#DGToCA3", Desc: "Mossy fibers: strong, non-learning",
				Params: params.Params{
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.9",
					"Prjn.WtInit.Var":  "0.01",
					"Prjn.WtScale.Rel": "4",
				}},
			{Sel: "#CA3ToCA3", Desc: "CA3 recurrent cons: rel=2 still the best",
				Params: params.Params{
					"Prjn.WtScale.Rel": "2",
					"Prjn.Learn.Lrate": "0.1",
				}},
			{Sel: "#ECinToDG", Desc: "DG learning is surprisingly critical: maxed out fast, hebbian works best",
				Params: params.Params{
					"Prjn.Learn.Learn":       "true", // NoDGLearn key. Set to false for NoDGLearn
					"Prjn.CHL.Hebb":          "0.2",
					"Prjn.CHL.SAvgCor":       "0.1",
					"Prjn.CHL.MinusQ1":       "true",
					"Prjn.Learn.Lrate":       "0.05",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
				}},
			{Sel: "#CA3ToCA1", Desc: "Schaffer collaterals -- slower, less hebb",
				Params: params.Params{
					"Prjn.CHL.Hebb":          "0.01",
					"Prjn.CHL.SAvgCor":       "0.4",
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
				}},
			{Sel: ".EC", Desc: "all EC layers: only pools, no layer-level",
				Params: params.Params{
					"Layer.Act.Gbar.L":        "0.1",
					"Layer.Inhib.ActAvg.Init": "0.2",
					"Layer.Inhib.Layer.On":    "false",
					"Layer.Inhib.Pool.Gi":     "2.0",
					"Layer.Inhib.Pool.On":     "true",
				}},
			{Sel: "#DG", Desc: "very sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.01",
					"Layer.Inhib.Layer.Gi":    "3.8",
				}},
			{Sel: "#CA3", Desc: "sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.02",
					"Layer.Inhib.Layer.Gi":    "2.8",
					"Layer.Learn.AvgL.Gain":   "2.5",
				}},
			{Sel: "#CA1", Desc: "CA1 only Pools",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.1",
					"Layer.Inhib.Layer.On":    "false",
					"Layer.Inhib.Pool.On":     "true",
					"Layer.Inhib.Pool.Gi":     "2.4",
					"Layer.Learn.AvgL.Gain":   "2.5",
				}},
		},
		// NOTE: it is essential not to put Pat / Hip params here, as we have to use Base
		// to initialize the network every time, even if it is a different size..
	}},
	{Name: "List010", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "10",
				}},
		},
	}},
	{Name: "List020", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "20",
				}},
		},
	}},
	{Name: "List030", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "30",
				}},
		},
	}},
	{Name: "List040", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "40",
				}},
		},
	}},
	{Name: "List050", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "50",
				}},
		},
	}},
	{Name: "List060", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "60",
				}},
		},
	}},
	{Name: "List070", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "70",
				}},
		},
	}},
	{Name: "List080", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "80",
				}},
		},
	}},
	{Name: "List090", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "90",
				}},
		},
	}},
	{Name: "List100", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "100",
				}},
		},
	}},
	{Name: "List125", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "125",
				}},
		},
	}},
	{Name: "List150", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "150",
				}},
		},
	}},
	{Name: "List200", Desc: "list size", Sheets: params.Sheets{
		"Pat": &params.Sheet{
			{Sel: "PatParams", Desc: "pattern params",
				Params: params.Params{
					"PatParams.ListSize": "200",
				}},
		},
	}},
	{Name: "SmallHip", Desc: "hippo size", Sheets: params.Sheets{
		"Hip": &params.Sheet{
			{Sel: "HipParams", Desc: "hip sizes",
				Params: params.Params{
					"HipParams.ECPool.Y":  "7",
					"HipParams.ECPool.X":  "7",
					"HipParams.CA1Pool.Y": "10",
					"HipParams.CA1Pool.X": "10",
					"HipParams.CA3Size.Y": "20",
					"HipParams.CA3Size.X": "20",
					"HipParams.DGRatio":   "2.236", // 1.5 before, sqrt(5) aligns with Ketz et al. 2013
				}},
		},
	}},
	{Name: "MedHip", Desc: "hippo size", Sheets: params.Sheets{
		"Hip": &params.Sheet{
			{Sel: "HipParams", Desc: "hip sizes",
				Params: params.Params{
					"HipParams.ECPool.Y":  "7",
					"HipParams.ECPool.X":  "7",
					"HipParams.CA1Pool.Y": "15",
					"HipParams.CA1Pool.X": "15",
					"HipParams.CA3Size.Y": "30",
					"HipParams.CA3Size.X": "30",
					"HipParams.DGRatio":   "2.236", // 1.5 before
				}},
		},
	}},
	{Name: "BigHip", Desc: "hippo size", Sheets: params.Sheets{
		"Hip": &params.Sheet{
			{Sel: "HipParams", Desc: "hip sizes",
				Params: params.Params{
					"HipParams.ECPool.Y":  "7",
					"HipParams.ECPool.X":  "7",
					"HipParams.CA1Pool.Y": "20",
					"HipParams.CA1Pool.X": "20",
					"HipParams.CA3Size.Y": "40",
					"HipParams.CA3Size.X": "40",
					"HipParams.DGRatio":   "2.236", // 1.5 before
				}},
		},
	}},
	{Name: "EDL", Desc: "EDL or NoEDL in testing effect", Sheets: params.Sheets{
		"TE": &params.Sheet{
			{Sel: "TEParams", Desc: "EDL or NoEDL for testing effect",
				Params: params.Params{
					"TEParams.EDL": "true",
				}},
		},
	}},
	{Name: "NoEDL", Desc: "EDL or NoEDL in testing effect", Sheets: params.Sheets{
		"TE": &params.Sheet{
			{Sel: "TEParams", Desc: "EDL or NoEDL for testing effect",
				Params: params.Params{
					"TEParams.EDL": "false",
				}},
		},
	}},
	{Name: "RP", Desc: "Retrieval Practice or Restudy in testing effect", Sheets: params.Sheets{
		"TE": &params.Sheet{
			{Sel: "TEParams", Desc: "Retrieval Practice or Restudy in testing effect",
				Params: params.Params{
					"TEParams.IsRP": "true",
				}},
		},
	}},
	{Name: "RS", Desc: "Retrieval Practice or Restudy in testing effect", Sheets: params.Sheets{
		"TE": &params.Sheet{
			{Sel: "TEParams", Desc: "Retrieval Practice or Restudy in testing effect",
				Params: params.Params{
					"TEParams.IsRP": "false",
				}},
		},
	}},
}
