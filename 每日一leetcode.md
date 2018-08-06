- https://leetcode-cn.com/problems/reverse-integer/description/ 反转有范围整数 --- 题号 7 难度: ⭐

```javascript
/**
 * @param {number} x
 * @return {number}
 */

// 方法一
var reverse = function(x) {
  let str = String(x)
  let isMinus = str[0] === '-'
  let ret
  if (isMinus) {
    str = str.slice(1)
  }
  ret = Number(
    str
      .replace(/()(?=\d)/g, ',')
      .split(',')
      .reverse()
      .join('')
  )
  ret = isMinus ? ret * -1 : ret
  if (ret < -(2 ** 31) || ret > 2 ** 31 - 1) {
    return 0
  }
  return ret
}

// 方法二
var reverse = function(x) {
  let ret = 0,
    pop = 0
  let slice = x < 0 ? Math.ceil : Math.floor
  while (x !== 0) {
    pop = x % 10
    ret = ret * 10 + pop
    x = slice(x / 10)
  }
  if (ret > 2 ** 31 - 1 || ret < -(2 ** 31)) {
    return 0
  } else {
    return ret
  }
}
```

- https://leetcode-cn.com/problems/two-sum/description/ 两数之和 --- 题号 1 难度: ⭐

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  let result = []
  nums.forEach((item, index) => {
    let index1 = index,
      index2 = null
    for (let i = index + 1; i < nums.length; i++) {
      const sum = item + nums[i]
      if (sum === target) {
        index2 = i
        break
      }
    }
    if (index2) {
      result = [index1, index2]
      return
    }
  })
  return result
}
```

- https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/ 删除排序数组中的重复项 --- 题号 26 难度: ⭐

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  let cur = NaN
  for (var i = 0; i < nums.length; i++) {
    if (nums[i] !== cur) {
      cur = nums[i]
    } else {
      nums.splice(i, 1)
      i--
    }
  }
  return nums.length
}
```

- https://leetcode-cn.com/problems/trapping-rain-water/description/ 42. 接雨水 难度: ⭐ ⭐ ⭐

```javascript
/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function(height) {
  let len = height.length
  let leftMaxn = []
  let rightMaxn = []

  let maxn = 0
  for (let i = 0; i < len; i++) {
    leftMaxn[i] = maxn
    maxn = Math.max(maxn, height[i])
  }

  maxn = 0
  for (let i = height.length; i--; ) {
    rightMaxn[i] = maxn
    maxn = Math.max(maxn, height[i])
  }

  let sum = 0
  for (let i = 0; i < len; i++) {
    let left = leftMaxn[i]
    let right = rightMaxn[i]
    let minn = Math.min(left, right)

    if (minn > height[i]) {
      sum += minn - height[i]
    }
  }

  return sum
}
```

- https://leetcode-cn.com/problems/container-with-most-water/description/ 11. 盛最多水的容器 ⭐ ⭐

```javascript
/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
  let max = -Infinity
  height.forEach((item, index) => {
    let h, w, totalWater
    for (let i = index + 1; i < height.length; i++) {
      h = height[i] > item ? item : height[i]
      w = i - index
      totalWater = w * h
      if (totalWater > max) max = totalWater
    }
  })

  return max
}
```

- https://leetcode-cn.com/problems/palindrome-number/description/ 回文数 --- 题号 9 难度: ⭐

```javascript
/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
  let toStr = String(x),
    ret = true
  for (let i = 0; i < toStr.length; i++) {
    if (i >= ~~(toStr.length / 2))
      if (toStr[i] !== toStr[toStr.length - 1 - i]) {
        ret = false
      }
  }
  return ret
}
```

- https://leetcode-cn.com/problems/4sum/description/ 4SUM --- 题号 18 难度: ⭐⭐

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function(nums, target) {
  var len = nums.length
  if (len < 4) return []
  nums.sort((a, b) => (a < b ? -1 : a > b ? 1 : 0))
  // console.log(nums);
  var ans = []
  for (var i = 0; i < len - 3; i++) {
    for (var j = i + 1; j < len - 2; j++) {
      var k = j + 1
      var l = len - 1

      while (k < l) {
        var sum = nums[i] + nums[j] + nums[k] + nums[l]
        if (sum === target) {
          ans.push([nums[i], nums[j], nums[k], nums[l]])
          while (nums[l--] === nums[l] && nums[k++] === nums[k] && k < l);
        } else if (sum < target) while (nums[k++] === nums[k] && k < l);
        else while (nums[l--] === nums[l] && k < l);
      }
      while (nums[j] === nums[j + 1]) j++
    }
    while (nums[i] === nums[i + 1]) i++
  }
  return ans
}
```

- https://leetcode-cn.com/problems/3sum-closest/description/ 3SUM-closest --- 题号 16 难度: ⭐⭐

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function(nums, target) {
  var len = nums.length,
    minDiff = Number.MAX_VALUE,
    diff,
    left,
    right,
    i,
    j

  nums.sort(function(a, b) {
    return a - b
  })

  for (i = 0; i < len; i++) {
    left = i + 1
    right = len - 1

    while (left < right) {
      diff = target - nums[i] - nums[left] - nums[right]

      if (diff === 0) {
        return target
      } else if (diff > 0) {
        left++
      } else {
        right--
      }

      if (Math.abs(diff) < Math.abs(minDiff)) {
        minDiff = diff
      }
    }
  }

  return target - minDiff
}
```
