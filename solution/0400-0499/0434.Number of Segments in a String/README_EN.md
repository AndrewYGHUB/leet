# [434. Number of Segments in a String](https://leetcode.com/problems/number-of-segments-in-a-string)

[中文文档](/solution/0400-0499/0434.Number%20of%20Segments%20in%20a%20String/README.md)

## Description

<p>Given a string <code>s</code>, return <em>the number of segments in the string</em>.</p>

<p>A <strong>segment</strong> is defined to be a contiguous sequence of <strong>non-space characters</strong>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;Hello, my name is John&quot;
<strong>Output:</strong> 5
<strong>Explanation:</strong> The five segments are [&quot;Hello,&quot;, &quot;my&quot;, &quot;name&quot;, &quot;is&quot;, &quot;John&quot;]
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;Hello&quot;
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= s.length &lt;= 300</code></li>
	<li><code>s</code> consists of lowercase and uppercase English letters, digits, or one of the following characters <code>&quot;!@#$%^&amp;*()_+-=&#39;,.:&quot;</code>.</li>
	<li>The only space character in <code>s</code> is <code>&#39; &#39;</code>.</li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def countSegments(self, s: str) -> int:
        return len(s.split())
```

```python
class Solution:
    def countSegments(self, s: str) -> int:
        res = 0
        for i in range(len(s)):
            if s[i] != ' ' and (i == 0 or s[i - 1] == ' '):
                res += 1
        return res
```

### **Java**

```java
class Solution {
    public int countSegments(String s) {
        int res = 0;
        for (String t : s.split(" ")) {
            if (!"".equals(t)) {
                ++res;
            }
        }
        return res;
    }
}
```

```java
class Solution {
    public int countSegments(String s) {
        int res = 0;
        for (int i = 0; i < s.length(); ++i) {
            if (s.charAt(i) != ' ' && (i == 0 || s.charAt(i - 1) == ' ')) {
                ++res;
            }
        }
        return res;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int countSegments(string s) {
        int res = 0;
        for (int i = 0; i < s.size(); ++i)
        {
            if (s[i] != ' ' && (i == 0 || s[i - 1] == ' '))
                ++res;
        }
        return res;
    }
};
```

### **Go**

```go
func countSegments(s string) int {
	res := 0
	for i, c := range s {
		if c != ' ' && (i == 0 || s[i-1] == ' ') {
			res++
		}
	}
	return res
}
```

### **...**

```

```

<!-- tabs:end -->