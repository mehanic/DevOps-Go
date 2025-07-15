#Approach 1: Brute Force
from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[j] == target - nums[i]:
                    return [i, j]
        # Return an empty list if no solution is found
        return []
    


# Two-pass Hash Table

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i in range(len(nums)):
            hashmap[nums[i]] = i
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap and hashmap[complement] != i:
                return [i, hashmap[complement]]
        # If no valid pair is found, return an empty list
        return []
    

#One-pass Hash Table

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap:
                return [i, hashmap[complement]]
            hashmap[nums[i]] = i
        # Return an empty list if no solution is found
        return []
    
#
def twoSum(self, nums, target):
        seen = {}
        for i, v in enumerate(nums):
            remaining = target - v
            if remaining in seen:
                return [seen[remaining], i]
            seen[v] = i
        return []

#

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        h = {}
        for i, num in enumerate(nums):
            n = target - num
            if n not in h:
                h[num] = i
            else:
                return [h[n], i]
            
#

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
        hash_map = dict()
        
        for index, val in enumerate(nums):
            if hash_map.has_key(val):
                return [hash_map.get(val), index]

            hash_map[target - val] = index

#

class Solution:
   def twoSum(self, nums, target):
      """
       :type nums: List[int]
       :type target: int
       :rtype: List[int]
      """
      dic = dict()
      for i, e in enumerate(nums):
          result = target - e
          if result in dic:
            return [dic[result], i]
          dic[e] = i

#

def two_sum(nums, target):
    lookup = {}  # Dictionary to store num: index

    for idx, num in enumerate(nums):
        complement = target - num
        if complement in lookup:
            return lookup[complement], idx
        lookup[num] = idx
#

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        complement_dict = {}
        nums_len = len(nums)
        for i in range(nums_len):
            complement = target - nums[i]
            if complement in complement_dict:
                return [complement_dict[complement], i]
            else:
                if nums[i] in complement_dict:
                    continue
                complement_dict[nums[i]] = i

#
class Solution:
    def twoSum(self, nums, target):
        num_set = {}
        for num_index, num in enumerate(nums):
            if (target-num) in num_set:
                return [num_set[target-num], num_index]
            num_set[num] = num_index

#

def twoSum(self, nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: List[int]
    """
    reidx = []
    if len(nums)<2:
        pass
    for i in range(len(nums)):
        for j in range(i+1,len(nums)):
            sums = nums[i] + nums[j]
            if sums == target:
                return [i,j]
            
class Solution:
    def twoSum(self, nums, target):
        for i,num in enumerate(nums):
            if (target-num) in nums:
                if nums.index(target-num)!=i:
                    return [i,nums.index(target-num)]
                

class Solution:
  def twoSum(self, nums: 'List[int]', target: 'int') -> 'List[int]':
     dict = {}
     for i in range(0, len(nums)):
       if nums[i] in dict:
          return [dict[nums[i]], i]
       else:
          dict[target - nums[i]] = i

class Solution(object):
    def twoSum(self, nums, target):
        a ={}
        for i, num in enumerate(nums):
            if target-num in a:
                return [a[target - num], i]
            else:
                a[num] = i

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hashmap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap:
                return [i, hashmap[complement]]
            hashmap[nums[i]] = i

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_map = {}  # Hash table to store number and its index
        for i, num in enumerate(nums):
            complement = target - num  # Find the complement
            if complement in num_map:
                return [num_map[complement], i]  # Return indices of complement and current number
            num_map[num] = i  # Store the number with its index

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # sort the given array
        sorted_nums = sorted(nums)
        # take two indices i,j
        i = 0 #start
        j = len(sorted_nums)-1 #end
        
        s = sorted_nums[i]+sorted_nums[j] 
        while s != target:
            # move j to left if sum > target (to decrease sum) 
            if s > target:
                j-=1
            else:
                i+=1 # move i to right if sum < target (to increase sum)
            s = sorted_nums[i]+sorted_nums[j]
        # return corresponding indices of elements in original array
        x = nums.index(sorted_nums[i])
        # change value at x to avoid referring to same element twice
        nums[x] = -1
        y = nums.index(sorted_nums[j])
        return [y,x]
    

def twoSum(self, nums: 'List[int]', target: 'int') -> 'List[int]':
    """
    finds two numbers which add to a given target number and returns their indices
    """
    # create dict to store reciprocal digits
    r = {}
    # loop through nums (keep index)
    for x in range(0, len(nums)):
        n = target - nums[x]
        # check dict for reciprocal of current number
        if n in r:
            return [r[n], x]
        # if not in dict, add current number to avoid doubles
        r[nums[x]] = x


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
                      
        mapping = {n: i for i, n in enumerate(nums)}
        
        for i, num in enumerate(nums):
            diff = mapping.get(target - num, False)
            
            if diff and diff != i:
                return [i, diff]
            
    # if not nums or len(nums)<2:
    #     return[-1,-1]
    
    # hashmap = dict()
    
    # for index, value in enumerate(nums):
    #     if target-value in hashmap:
    #         return[hashmap[target-value], index]
    #     hashmap[value] = index
        
    # return [-1,-1]

# class Solution:
#         def twoSum(self, nums: List[int], target: int) -> List[int]:
#         for i, n in enumerate(nums):
#             if target - n in nums and nums.index(target-n) != i:
#                 return sorted([i, nums.index(target - n)])
#         return []


def twoSum(self, nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: List[int]
    """

    return_index = []
    for n in range(0, len(nums)):
        com = target - nums[n]
        if(com in nums and nums.index(com) != n):
            return_index = [n, nums.index(com)]
            break

    if(n == len(nums) - 1 and return_index == []):
        return_index = "No two sums"
    return return_index


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        nums_dict=dict(zip(nums,range(len(nums))))
        for i in range(len(nums)):
            complement=target-nums[i]
            if complement in nums_dict.keys() and nums_dict[complement]!=i:
                return([i,nums_dict[complement]])
            
def twoSum(self, nums, target):        
        temp={}
        for i in range(len(nums)):
            if nums[i] not in temp:
                 temp[nums[i]]=i
            if target-nums[i] in temp.keys() and temp[target-nums[i]]!=i:
                return [temp[target-nums[i]],i]
            else:
                if target-nums[i] in temp.keys() and temp[target-nums[i]]!=i:
                    return [temp[target-nums[i]],i]
                

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        my_dict = {}
        for i in range(len(nums)):
            idx = nums[i]
            diff = target - idx
            if diff in my_dict:
                return [i, my_dict[diff]]
            my_dict[idx] = i


class Solution:
   def twoSum(self, nums, target):
     """
      :type nums: List[int]
      :type target: int
      :rtype: List[int]
     """
     for i in range(len(nums)):
       a = target-nums[i]
     if a in nums and i!=nums.index(a):
       return([i,nums.index(a)])
     
class Solution(object):
    def twoSum(self, nums, target):
        if len(nums) < 2:
            pass
        r = []
        for i in range(len(nums)):
            comp = target - nums[i]
            if comp in r:
                return [i, nums.index(comp)]
            r.append(nums[i])

class Solution:
    def twoSum(self,nums, target):
        if len(nums) < 2:
           pass
    tempDict = {nums[0] : 0}
    for i in range(1, len(nums)):
        checkNum = target-nums[i]
    if(checkNum in tempDict.keys()):
        return [i,tempDict[checkNum]]
    else:
        tempDict[nums[i]] = i


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        nums_hash = {}
        for j, num in enumerate(nums):
            complement = target - num
            if complement in nums_hash:
                return [nums_hash[complement], j]
            else:
                nums_hash[num] = j

class Solution:
    def twoSum(self, nums, target):
        dictionary = {}
        for key, value in enumerate(nums):
            complement = target - value
            if dictionary.__contains__(complement):
                return [dictionary.get(complement), key]
            dictionary[value] = key

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums = nums[0] + nums[1]
        return nums, target
    
class Solution:
    
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        recs = dict((num, i) for i, num in enumerate(nums))
        
        for i, num in enumerate(nums):
            j = recs.get(target-num)
            if j and not j==i:
                return [i,j]
            
class Solution:
    def get_count_map(nums):
        count_map = dict()
        for i in range(0, len(nums)):
            if nums[i] in count_map:
                temp = count_map[nums[i]]
                temp[1].append(i)
                temp = (temp[0]+1, temp[1])
                count_map[nums[i]] = temp
            else:
                count_map[nums[i]] = (1, [i])
        return count_map


    def print_indices(count_map, k):
        for i in count_map:
            if 2*i == k:
                return [count_map[i][1][0], count_map[i][1][1]]
                break
            else:
                if k-i in count_map:
                    return [count_map[i][1][0], count_map[k-i][1][0]]
                    break
    def twoSum(self, nums, target):
        return Solution.print_indices(Solution.get_count_map(nums), target)
    

    def twoSum(self, nums, target):

        numLens = len(nums)
        index = list(range(numLens))
        dict_demo = dict(zip(nums, index))
        for num in nums:
            rem = target - num
            match = dict_demo.get(rem, None)
            if match != None and nums.index(num) != match:
                return [nums.index(num), match]
            
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hash_map = {}
        for idx, num1 in enumerate(nums):
            sub1 = target - num1
            if sub1 in hash_map:
                return [hash_map[sub1],idx]
            hash_map[num1] = idx


    # def twoSum(self, nums, target):
    #     dict = {}
    #     for i in range(0, len(nums)):
    #       complement = target - nums[i]
    #     if dict.get(complement) is not None:
    #          return [i, dict[complement]]

    #     dict[nums[i]] = i

    # raise ValueError("No solution")


class Solution(object):
   def twoSum(self, nums, target):
      """
      :type nums: List[int]
      :type target: int
      :rtype: List[int]
      """
      for ind_i in range(0, len(nums)):
        for ind_j in range(ind_i+1, len(nums)):
      if(nums[ind_i] + nums[ind_j]) == target:
        return ind_i, ind_j
      
    #   def twoSum(self, nums, target):
    # for i in range(len(nums)):
    #     diff = target - nums[i]
    #     if diff in nums and nums.index(diff) != i:
    #         nums[i] = float('nan')
    #         return [i, nums.index(diff)]    
        
def twoSum(self, nums, target):
    dict = {}
    for i in range(len(nums)):
        sol = target-nums[i]
        if sol in dict:
            return dict[sol], i
        dict[nums[i]]=i

class Solution(object):
    def twoSum(self, nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: List[int]
    """
    new_list = {j: i for i,j in enumerate(nums)}
    for n in range(len(nums)):
        if(new_list.has_key(target - nums[n]) and n !=new_list[target - nums[n]]):
            return([n, new_list[target - nums[n]]])
        
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hash_table = {}
        for i in range(0,len(nums)):
            if nums[i] not in hash_table:
                hash_table[nums[i]] = [i]
            else:
                hash_table[nums[i]].append(i)
        
        for k,v in hash_table.iteritems():
            check_key = target - k
            if check_key in hash_table:
                for x in hash_table[check_key]:
                    for v in v:
                        if x == v:
                            continue
                        else:
                            return [v,x]
                    
                
  
        return -1
    
class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        ret=[]
        for i in range(len(nums)):
            if target-nums[i] in nums[i+1::]:
                ret.append(i)
                ret.append(nums[i+1::].index((target-nums[i]))+i+1)
        return ret
    
for a in range(len(nums)):
            for b in range(len(nums)-a):
                if nums[a] + nums[b+a] == target:
                    if a != b+a:
                        return [a,b+a]

class Solution:
    def twoSum(self, nums, target):
        integer_table = {}
        count = 0
        for integer in nums:
            conjugate = target - integer
            if conjugate in integer_table:
                return  [integer_table[conjugate], count]
            integer_table[integer] = count
            count += 1
        return None

