This problem was asked by Pivotal.

Write an algorithm that finds the total number of set bits in all integers between 1 and N.

NOTES: I wanted to try something similar to a prime number seive. I expected storing the results in a map to be faster than constantly looping over a string, but the string method is 2+ orders of magnitude faster. 