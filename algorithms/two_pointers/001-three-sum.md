## Problem

Given an integer array `nums`, find and return all unique triplets `[nums[i], nums[j], nums[k]]`, where the indexes satisfy `i ≠ j`, `i ≠ k`, and `j ≠ k`, and the sum of the elements `nums[i] + nums[j] + nums[k] == 0`.

### Example

Input: `nums = [-1, 0, 1, 2, -1, -4]`

Output: `[[-1, -1, 2], [-1, 0, 1]]`

### Constraints
The solution set must not contain duplicate triplets.

### Solution (Python)
```python
def find_two_sum_pairs(start_index, array_length, target_sum, nums):
    left, right = start_index, array_length - 1
    triplets = []
    
    while left < right:
        current_sum = nums[left] + nums[right]
        
        if current_sum == target_sum:
            triplets.append([-target_sum, nums[left], nums[right]])
            
            while left + 1 < right and nums[left + 1] == nums[left]: left += 1
            while right - 1 > left and nums[right - 1] == nums[right]: right -= 1
            left += 1
            right -= 1

        elif current_sum < target_sum: left += 1
        else: right -= 1
    
    return triplets

def three_sum(nums):
    nums.sort()
    array_length = len(nums)
    result = []
    
    for i in range(array_length):
        if nums[i] > 0: break
        if i > 0 and nums[i] == nums[i - 1]: continue
        
        triplets = find_two_sum_pairs(i + 1, array_length, -nums[i], nums)
        for triplet in triplets: result.append(triplet)
    
    return result
```
