# About
This repository includes simulations for the Theremin (Total Hippocampal ERror MINimization) model described in [Correcting the Hebbian Mistake: Toward a Fully Error-Driven Hippocampus (Zheng et al., 2021)](https://doi.org/10.1101/2021.10.29.466546).

# Brief summary
Theremin is the latest implementation of a standard hippocampus model under the [Leabra](https://github.com/emer/leabra) framework. The basic idea for this name comes from the increasing (almost covering every essential pathways) error-driven learning dynamics in the model, as more and more neurophysiological data comes to support the biologically plausible temporal difference error-driven learning principles. 

Theremin learns to do the AB-AC task (Barnes & Underwood, 1959) but can be adjusted to perform other experiments as needed. We provide a GUI Theremin, which could be used to observe the model running in real time and do some arbitrary simple tweaks, and a non-GUI Theremin used for getting simulation results. 

# Instructions for running the models
## Installation
1. Install [Golang](https://go.dev/doc/install) 1.18
2. Install [Emergent](https://github.com/emer/emergent/wiki/Install)

## Main Theremin
After successfully installing Go environments and needed libraries, under the `sims` directory, run `go build` to generate the main exec file `sims`.

### GUI
To run GUI version Theremin, simply type `./sims`. Plenty functionalities are included in the GUI for easy experimentations. 

### Non-GUI
To get a quick look of what the model produces, we can run 1 run locally using the fastest configuration (small hippocampus learning a short AB-AC list). First, search for the keyword "OuterLoopParams" in `Theremin.go` and replace the codes there with: 

```go
// OuterLoopParams are the parameters to run for outer crossed factor testing
var OuterLoopParams = []string{"SmallHip"}

// InnerLoopParams are the parameters to run for inner crossed factor testing
var InnerLoopParams = []string{"List020"}
```

After building the model, run `./sims --nogui --runs 1` to run 1 run locally. Note that we provide options to save different files under the `CmdArgs()` function.

To replicate the results shown in the paper, keep `Theremin.go` the same as in this repository, run `./sims --nogui --run X --runs 1` on the server to create one job, and repeat this process 30 times with X ranging from 0 to 29. We used [Grunt](https://github.com/emer/grunt) for creating this parallel batch run workflow on the server to reduce running time. 

## Testing Effect Theremin
Results for testing effect shown in the paper could be replicated here (GUI not used). Under the `sims/testing_effect` directory:
1. `go build`
2. repeat `./testing_effect --nogui --run X --runs 1` 30 times with X ranging from 0 to 29

# Run different variants of Theremin
For model comparison, different variants could be run by altering several lines of code in `Theremin.go`. To change the model to a specific variant, search for the variant's name in `Theremin.go` (for ThetaPhase, NoEDL, NoDynMF, and NoPretrain) or `def_params.go` (for NoDGLearn) and follow instructions there. The running procedure remains the same as in [main Theremin](#main-theremin).