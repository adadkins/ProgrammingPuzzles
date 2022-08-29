Question from Amazon Inteveriew:

Given a set of boxes/weights in an array, return the smallest set that is larger than the rest of the boxes.

Solution: The smallest set will be the largest X boxes so sort the array and keep adding until the Largest X boxes are heavier than the rest. A possible optimiaztion is to get sums once and then subtract/add as we move the box to the smallest set so we dont have to continually loop to sum up the boxes. 