# Hyperloglog
* Reference: https://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf

# Problem
* How many unique visitors in a day
* Solution: use HashMap. 
* Problem: With scale, space complexity.

# Question to ask 
* Do we actually need a fully accurate count?
* What is the probability that 
    * any binary number ends with 1 or 0? P = 1/2
    * For 2 binary number, P = 1/4
    * For n, P = 1/2^n
* Max would 2^n where n = max zeroes at the end.

```
1001
0000
1100
1101
0001

Max zeroes at end = 4
unique entries = 2^4 = 16
but here with this dataset, we got only 5.
```

# Solution to above (using bucket) => LogLog algorithm

1. Bucket by leading 2 zeroes for the above example
```
0 -- 00 -> 00(00), 0001
1 -- 01 -> 
2 -- 10 -> 1001
3 -- 11 -> 11(00), 1101
```
2. Calculate the mean of number of trailing zeroes per bucket
```
mean = 2 + 0 + 0 + 2 / 4 = 1
```
3. Calculate the unique entries as 
```
// 0.79 was the constant after experimentation as per the paper
= 0.79 * 4 * 2^mean = 6.32
```

# Closer to the target use harmonic mean => hyperloglog algorithm.

```
harmonic mean = 1/2^2 + 1/2^0 + 1/2^0 + 1/2^2
New formula = 0.79 * 4 * 4 / harmonic mean = 5.056
```

