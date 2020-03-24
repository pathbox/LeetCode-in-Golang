将第一个出现的字母映射成 1，第二个出现的字母映射成 2

对于 egg
e -> 1
g -> 2
也就是将 egg 的 e 换成 1, g 换成 2, 就变成了 122

对于 add
a -> 1
d -> 2
也就是将 add 的 a 换成 1, d 换成 2, 就变成了 122

egg -> 122, add -> 122
都变成了 122，所以两个字符串异构。

为了方便，我们也可以将当前字母直接映射为当前字母的下标加 1。因为 0 是我们的默认值，所以不能直接赋值为下标，而是「下标加 1」

```java
public boolean isIsomorphic(String s, String t) {
    return isIsomorphicHelper(s).equals(isIsomorphicHelper(t));
}

private String isIsomorphicHelper(String s) {
    int[] map = new int[128];
    int n = s.length();
    StringBuilder sb = new StringBuilder();
    for (int i = 0; i < n; i++) {
        char c = s.charAt(i);
        //当前字母第一次出现,赋值
        if (map[c] == 0) {
            map[c] = i + 1;
        }
        sb.append(map[c]);
    }
    return sb.toString();
}
```