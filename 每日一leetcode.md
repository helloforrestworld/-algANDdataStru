- https://leetcode-cn.com/problems/reverse-integer/description/ 反转有范围整数
```javascript
/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let str = String(x)
    let isMinus = str[0] === '-'
    let ret
    if (isMinus) {
        str = str.slice(1)
    }
    ret = Number(str.replace(/()(?=\d)/g, ',').split(',').reverse().join(''))
    ret = isMinus ? ret * -1 : ret
    if (ret < -(2**31) || ret > (2**31) -1) {return 0}
    return ret
};
```

- https://leetcode-cn.com/problems/two-sum/description/ 两数之和
```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let result = []
    nums.forEach((item, index) => {
        let index1 = index, index2 = null
        for (let i = index + 1; i < nums.length; i++) {
            const sum = item + nums[i]
            if (sum === target) {
                index2 = i
                break;
            }
        }
        if (index2) {
            result = [index1, index2]
            return
        }
    })
    return result
};
```

