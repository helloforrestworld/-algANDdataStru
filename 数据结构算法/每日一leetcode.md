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

- https://leetcode-cn.com/problems/implement-strstr/description/ 实现 strStr()--- 题号 28 难度: ⭐

```javascript
/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
  if (needle === '') return 0
  let times = haystack.length - needle.length + 1,
    ret = -1,
    len = needle.length
  if (times <= 0) return -1
  for (let i = 0; i < times; i++) {
    if (haystack.substr(i, len) === needle) {
      ret = i
      break
    }
  }
  return ret
}
```

- https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/ 电话号码组合()--- 题号 17 难度: ⭐⭐⭐

```javascript
/**
 * @param {string} digits
 * @return {string[]}
 */
var ans, tmp

function dfs(str, idx, digits) {
  if (idx === digits.length) {
    ans.push(str)
    return
  }

  var num = Number(digits[idx])
  if (num <= 1) dfs(str, idx + 1, digits)
  else {
    for (var i = 0, len = tmp[num].length; i < len; i++)
      dfs(str + tmp[num][i], idx + 1, digits)
  }
}

var letterCombinations = function(digits) {
  if (!digits.length) return []
  ;(tmp = ['', '', 'abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']),
    (ans = [])

  dfs('', 0, digits)
  return ans
}
```

- https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/ 电话号码组合()--- 题号 17 难度: ⭐⭐⭐

```javascript
/**
 * @param {string}
 * @return {number}
 */
var romanToInt = function(s) {
  var hash = {}
  hash['I'] = 1
  hash['X'] = 10
  hash['C'] = 100
  hash['M'] = 1000
  hash['V'] = 5
  hash['L'] = 50
  hash['D'] = 500

  var sum = 0

  for (var i = 0, len = s.length; i < len; i++) {
    var item = hash[s[i]]

    var nextItem = i + 1 === len ? 0 : hash[s[i + 1]]

    if (nextItem > item) {
      sum += nextItem - item
      i++
    } else sum += item
  }

  return sum
}
```

- https://leetcode-cn.com/problems/generate-parentheses/description/ 括号生成()--- 题号 22 难度: ⭐⭐⭐

```javascript
/**
 * @param {number} n
 * @return {string[]}
 */
var ans

function dfs(s, left, right, n) {
  if (left === n && right === n) {
    ans.push(s)
    return
  }

  if (left + 1 <= n) dfs(s + '(', left + 1, right, n)

  if (right + 1 <= n && right + 1 <= left) dfs(s + ')', left, right + 1, n)
}

var generateParenthesis = function(n) {
  ans = []
  dfs('', 0, 0, n)
  return ans
}
```

- https://leetcode-cn.com/problems/generate-parentheses/description/ 括号生成()--- 题号 22 难度: ⭐⭐⭐

```javascript
/**
 * @param {number} n
 * @return {string[]}
 */
var ans

function dfs(s, left, right, n) {
  if (left === n && right === n) {
    ans.push(s)
    return
  }

  if (left + 1 <= n) dfs(s + '(', left + 1, right, n)

  if (right + 1 <= n && right + 1 <= left) dfs(s + ')', left, right + 1, n)
}

var generateParenthesis = function(n) {
  ans = []
  dfs('', 0, 0, n)
  return ans
}
```

- https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/ 题号 515 难度: ⭐⭐

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var largestValues = function(root) {
  let maxn = []

  let getMax = (a, b = -Number.MAX_VALUE) => Math.max(a, b)

  let dfs = (node, step) => {
    if (!node) return
    maxn[step] = getMax(node.val, maxn[step])
    dfs(node.left, step + 1)
    dfs(node.right, step + 1)
  }

  dfs(root, 0)
  return maxn
}
```

- https://leetcode.com/problems/keyboard-row/description/ 题号 500 难度: ⭐⭐

```javascript
/**
 * @param {string[]} words
 * @return {string[]}
 */
var findWords = function(words) {
  let keys = ['qwertyuiop', 'asdfghjkl', 'zxcvbnm']

  let ans = []

  words.forEach(function(item) {
    let s = new Set()
    let word = item.toLowerCase()

    for (let letter of word) {
      for (let i = 0; i < 3; i++)
        if (keys[i].indexOf(letter) !== -1) {
          s.add(i)
          break
        }
    }

    if (s.size === 1) ans.push(item)
  })

  return ans
}
```

- https://leetcode.com/problems/swap-nodes-in-pairs/description/ 题号 24 难度: ⭐⭐

```javascript
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
  if (!head) return null

  var arr = []

  while (head) {
    var next = head.next
    head.next = null
    arr.push(head)
    head = next
  }

  var len = arr.length

  for (var i = 0; i < len; i += 2) {
    var a = arr[i]
    var b = arr[i + 1]

    if (!b) continue

    arr[i] = b
    arr[i + 1] = a
  }

  for (var i = 0; i < len - 1; i++) arr[i].next = arr[i + 1]

  return arr[0]
}
```

- https://leetcode.com/problems/daily-temperatures/description/ 题号 739 难度: ⭐⭐

```javascript
/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
  let hash = []
  let ans = []

  for (let i = temperatures.length; i--; ) {
    let temperature = temperatures[i]
    ans[i] = Infinity

    for (let j = temperature + 1; j <= 100; j++) {
      if (hash[j] && hash[j] - i < ans[i]) {
        ans[i] = hash[j] - i
      }
    }

    if (ans[i] === Infinity) ans[i] = 0
    hash[temperature] = i
  }

  return ans
}
```

- https://leetcode.com/problems/find-bottom-left-tree-value/description/ 题号 513 难度: ⭐⭐

```javascript
/**
 * @param {TreeNode} root
 * @return {number}
 */
var findBottomLeftValue = function(root) {
  let res = []

  let dfs = (node, step) => {
    if (res[step] === undefined) res[step] = node.val

    node.left && dfs(node.left, step + 1)
    node.right && dfs(node.right, step + 1)
  }

  dfs(root, 0)
  return res.pop()
}
```

- https://leetcode-cn.com/problems/house-robber-ii/description/ 题号 213 难度: ⭐⭐⭐

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
  // 不能同时打劫第一家和最后一家
  function robHouses(nums) {
    var odd = 0
    var even = 0

    for (var i = 0; i < nums.length; i++) {
      var num = nums[i]
      if (i % 2 === 0) {
        even = Math.max(even + num, odd)
      } else {
        odd = Math.max(odd + num, even)
      }
    }

    return Math.max(even, odd)
  }

  if (nums.length <= 1) {
    return robHouses(nums)
  }

  var robHousesExceptLast = robHouses(nums.slice(0, -1))
  var robHousesExceptFirst = robHouses(nums.slice(1))
  return Math.max(robHousesExceptLast, robHousesExceptFirst)
}
```

- https://leetcode.com/problems/daily-temperatures/description/ 题号 739 难度: ⭐⭐⭐

```javascript
/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
  let hash = []
  let ans = []

  for (let i = temperatures.length; i--; ) {
    let temperature = temperatures[i]
    ans[i] = Infinity

    for (let j = temperature + 1; j <= 100; j++) {
      if (hash[j] && hash[j] - i < ans[i]) {
        ans[i] = hash[j] - i
      }
    }

    if (ans[i] === Infinity) ans[i] = 0
    hash[temperature] = i
  }

  return ans
}
```

- https://leetcode.com/problems/degree-of-an-array/description/ 题号 697 难度: ⭐

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var findShortestSubArray = function(nums) {
  let p = new Map()
  let degreeLen = 0

  nums.forEach((item, index) => {
    if (!p.has(item)) p.set(item, {})
    let obj = p.get(item)
    obj.start = obj.start === void 0 ? index : obj.start
    obj.end = index
    obj.count = ~~obj.count + 1
    degreeLen = Math.max(degreeLen, obj.count)
  })

  let ans = Number.POSITIVE_INFINITY
  for (let value of p.values()) {
    if (value.count !== degreeLen) continue
    ans = Math.min(ans, value.end - value.start + 1)
  }

  return ans
}
```

- https://leetcode.com/problems/teemo-attacking/description/ 题号 495 难度: ⭐⭐

```javascript
/**
 * @param {number[]} timeSeries
 * @param {number} duration
 * @return {number}
 */
var findPoisonedDuration = function(timeSeries, duration) {
  let ans = 0

  for (let i = 0, len = timeSeries.length; i < len; i++) {
    if (i === len - 1) {
      ans += duration
      continue
    }

    let curItem = timeSeries[i],
      nextItem = timeSeries[i + 1]

    if (curItem + duration <= nextItem) ans += duration
    else ans += nextItem - curItem
  }

  return ans
}
```

- https://leetcode-cn.com/problems/lru-cache/description/ LRU 缓存机制 题号 146. 难度: ⭐⭐⭐

```javascript
function DoublyLinkListNode(key, value) {
  this.key = key
  this.value = value
  this.prev = this.next = null
}

/**
 * @constructor
 */
var LRUCache = function(capacity) {
  this.head = this.tail = null
  this.maxCapacity = capacity
  this.currSize = 0
  this.hash = {}
}

/**
 * @param {number} key
 * @returns {number}
 */
LRUCache.prototype.get = function(key) {
  if (!this.hash[key]) {
    return -1
  }

  this.moveToHead(key)
  return this.hash[key].value
}

/**
 * @param {number} key
 * @param {number} value
 * @returns {void}
 */
LRUCache.prototype.set = function(key, value) {
  if (this.maxCapacity <= 0) {
    return
  }

  if (!this.hash[key]) {
    if (this.currSize === this.maxCapacity) {
      this.removeLast()
      this.currSize--
    }

    this.hash[key] = new DoublyLinkListNode(key, value)
    this.currSize++
  }

  this.hash[key].value = value
  this.moveToHead(key)
}

LRUCache.prototype.removeLast = function() {
  if (this.tail === null) {
    return
  }

  delete this.hash[this.tail.key]
  var newTail = this.tail.prev

  if (newTail === null) {
    this.head = this.tail = null
    return
  }

  this.tail.prev = null
  newTail.next = null
  this.tail = newTail
}

LRUCache.prototype.moveToHead = function(key) {
  var newHead = this.hash[key]

  if (this.head === null && this.tail === null) {
    this.head = this.tail = newHead
  }

  if (newHead === this.head) {
    return
  }

  if (newHead === this.tail) {
    this.tail = newHead.prev
  }

  if (newHead.prev) {
    newHead.prev.next = newHead.next
  }
  if (newHead.next) {
    newHead.next.prev = newHead.prev
  }

  newHead.prev = null
  newHead.next = this.head
  this.head.prev = newHead
  this.head = newHead
}
```

- https://leetcode.com/problems/daily-temperatures/description/ 题号 739. 难度: ⭐⭐

```javascript
/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
  let hash = []
  let ans = []

  for (let i = temperatures.length; i--; ) {
    let temperature = temperatures[i]
    ans[i] = Infinity

    for (let j = temperature + 1; j <= 100; j++) {
      if (hash[j] && hash[j] - i < ans[i]) {
        ans[i] = hash[j] - i
      }
    }

    if (ans[i] === Infinity) ans[i] = 0
    hash[temperature] = i
  }

  return ans
}
```
