### TwoSum

#### code

```python
class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        tmp = {}
        for x in range(len(nums)):
            m = nums[x]
            n = target - m
            if n in tmp:
                return [x, tmp[n]]
            tmp[m]=x
        return []
```

#### 分析

- 考察： 2数加和 奇偶性

- 先检配对值是否存在 不存在则 原值加入

- 迷思： 只对应一种答案 即不存在 索引覆盖的情况