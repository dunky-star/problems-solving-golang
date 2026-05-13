![Add binary algorithm](image.png)

---

### Algorithm

```go
i := len(a) - 1
j := len(b) - 1
carry := 0
result := ""

for i >= 0 || j >= 0 || carry == 1 {
    sum := carry

    if i >= 0 {
        sum += (a[i] - '0')
        i--
    }

    if j >= 0 {
        sum += (b[j] - '0')
        j--
    }

    bit := sum % 2
    carry = sum / 2

    result = bit + result   // add to left
}