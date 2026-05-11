### Container With Most Water (Two Pointer Approach)

**Input:**
[1, 8, 6, 2, 5, 7]
![alt text](image.png)


---

### Algorithm

``---

### Algorithm

```go
l := 0
r := len(input) - 1
maxArea := 0

for l < r {
    length := min(height[l], height[r]) // e.g. min(8, 7) = 7
    width := r - l
    area := length * width
    maxArea = max(maxArea, area)

    if height[l] < height[r] {
        l++
    } else {
        r--
    }
}