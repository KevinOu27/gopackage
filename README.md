# MSDS 431 Assignment 9: Creating a Go Package

## Overview 

This repository contains a Go package that computes the trimmed mean for slices of floats. 

## Installation

To install this package, run this command:
go get github.com/KevinOu27/gopackage
This will install the package and allow you to run the trimmedmean go package created in this repository.

## Functionality

The Trimmed Mean Calculation is meant to read a slice of floats and compute the trimmed mean. 
To test the package, run 'go test -v' in the directory where the repository is located.

## Testing and Benchmarking

R Script
An R script is provided (gopackage.R) to calculate trimmed means using R's built-in functionalities which allow for a comparison of the outputs between Go and R implementations. It brings a random selected 100 floats and computes the trimmed mean and we can use this trimmed mean in the go unit tests to compare the two. 

Go Testing
The gopackage_test.go includes unit tests that verify the accuracy and efficiency of the trimmed mean calculations. Tests are performed with at least 100 floats and comparing the results against those from the R script.

## Application 

This package is designed with usability in mind where it features a clean, efficiency design suitable for integration into larger projects. 

## Conclusion

The 'go-trimmean' package is supposed to be a reliable method for calculating trimmed means, offering robustness against outliers and data variability.
