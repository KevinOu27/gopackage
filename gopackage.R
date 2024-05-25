library(MASS)

generateints = function(n){
  as.integer(runif(n, min = -100, max = 100))
}

generatefloats = function(n){
  runif(n, min = -100, max = 100)
}

computetrimmedmean = function(data){
  mean(data, trim = 0.05)
}

dataint = generateints(100)
datafloats = generatefloats(100)

trimmedints = computetrimmedmean(dataint)
trimmedfloats = computetrimmedmean(datafloats)

trimmedints
trimmedfloats
