# [46. 全排列](https://leetcode.cn/problems/permutations)

[English Version](/solution/0000-0099/0046.Permutations/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一个不含重复数字的数组 <code>nums</code> ，返回其 <em>所有可能的全排列</em> 。你可以 <strong>按任意顺序</strong> 返回答案。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3]
<strong>输出：</strong>[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,1]
<strong>输出：</strong>[[0,1],[1,0]]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1]
<strong>输出：</strong>[[1]]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 6</code></li>
	<li><code>-10 &lt;= nums[i] &lt;= 10</code></li>
	<li><code>nums</code> 中的所有整数 <strong>互不相同</strong></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

深度优先搜索。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        res = []
        path = [0] * n
        used = [False] * n

        def dfs(u):
            if u == n:
                res.append(path.copy())
                return
            for i in range(n):
                if not used[i]:
                    path[u] = nums[i]
                    used[i] = True
                    dfs(u + 1)
                    used[i] = False

        dfs(0)
        return res
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> path = new ArrayList<>();
        int n = nums.length;
        boolean[] used = new boolean[n];
        dfs(0, n, nums, used, path, res);
        return res;
    }

    private void dfs(int u, int n, int[] nums, boolean[] used, List<Integer> path, List<List<Integer>> res) {
        if (u == n) {
            res.add(new ArrayList<>(path));
            return;
        }
        for (int i = 0; i < n; ++i) {
            if (!used[i]) {
                path.add(nums[i]);
                used[i] = true;
                dfs(u + 1, n, nums, used, path, res);
                used[i] = false;
                path.remove(path.size() - 1);
            }
        }
    }
}
```

### **JavaScript**

```js
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function (nums) {
    const n = nums.length;
    let res = [];
    let path = [];
    let used = new Array(n).fill(false);
    dfs(0, n, nums, used, path, res);
    return res;
};

function dfs(u, n, nums, used, path, res) {
    if (u == n) {
        res.push(path.slice());
        return;
    }
    for (let i = 0; i < n; ++i) {
        if (!used[i]) {
            path.push(nums[i]);
            used[i] = true;
            dfs(u + 1, n, nums, used, path, res);
            used[i] = false;
            path.pop();
        }
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        int n = nums.size();
        vector<vector<int>> res;
        vector<int> path(n, 0);
        vector<bool> used(n, false);
        dfs(0, n, nums, used, path, res);
        return res;
    }

    void dfs(int u, int n, vector<int>& nums, vector<bool>& used, vector<int>& path, vector<vector<int>>& res) {
        if (u == n)
        {
            res.emplace_back(path);
            return;
        }
        for (int i = 0; i < n; ++i)
        {
            if (!used[i])
            {
                path[u] = nums[i];
                used[i] = true;
                dfs(u + 1, n, nums, used, path, res);
                used[i] = false;
            }
        }
    }
};
```

### **C#**

```cs
using System.Collections.Generic;
using System.Linq;

public class Solution {
    public IList<IList<int>> Permute(int[] nums) {
        var results = new List<IList<int>>();
        var temp = new List<int>();
        var visited = new bool[nums.Length];
        Search(nums, visited, temp, results);
        return results;
    }

    private void Search(int[] nums, bool[] visited, IList<int> temp, IList<IList<int>> results)
    {
        int count = 0;
        for (var i = 0; i < nums.Length; ++i)
        {
            if (visited[i]) continue;
            ++count;
            temp.Add(nums[i]);
            visited[i] = true;
            Search(nums, visited, temp, results);
            temp.RemoveAt(temp.Count - 1);
            visited[i] = false;
        }
        if (count == 0 && temp.Any())
        {
            results.Add(new List<int>(temp));
        }
    }
}
```

### **Go**

```go
func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	path := make([]int, n)
	used := make([]bool, n)
	dfs(0, n, nums, used, path, &res)
	return res
}

func dfs(u, n int, nums []int, used []bool, path []int, res *[][]int) {
	if u == n {
		t := make([]int, n)
		copy(t, path)
		*res = append(*res, t)
		return
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			path[u] = nums[i]
			used[i] = true
			dfs(u+1, n, nums, used, path, res)
			used[i] = false
		}
	}
}
```

### **TypeScript**

```ts
function permute(nums: number[]): number[][] {
    const n = nums.length;
    const res: number[][] = [];
    const dfs = (i: number) => {
        if (i === n) {
            res.push([...nums]);
        }
        for (let j = i; j < n; j++) {
            [nums[i], nums[j]] = [nums[j], nums[i]];
            dfs(i + 1);
            [nums[i], nums[j]] = [nums[j], nums[i]];
        }
    };
    dfs(0);
    return res;
}
```

### **Rust**

```rust
impl Solution {
    fn dfs(i: usize, nums: &mut Vec<i32>, res: &mut Vec<Vec<i32>>) {
        let n = nums.len();
        if i == n {
            res.push(nums.clone());
            return;
        }
        for j in i..n {
            nums.swap(i, j);
            Self::dfs(i + 1, nums, res);
            nums.swap(i, j);
        }
    }

    pub fn permute(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        Self::dfs(0, &mut nums, &mut res);
        res
    }
}
```

### **...**

```

```

<!-- tabs:end -->