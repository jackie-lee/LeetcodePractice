# -*- coding: UTF-8 -*-
#!/usr/bin/python
'''
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

解题思路2:
该方法的核心是将原问题转变成一个寻找第k小数的问题,假设两个析序列升序排列,这样中位数实际上是第(m+n)/2小的数.
所以只要解决了第k小数的问题，原问题也得以解决。

首先假设数组A和B的元素个数都大于k/2，我们比较A[k/2-1]和B[k/2-1]两个元素，这两个元素分别表示A的第k/2小的元素和B的第k/2小的元素。这两个元素比较共有三种情况：>、<和=。如果A[k/2-1]<B[k/2-1]，这表示A[0]到A[k/2-1]的元素都在A和B合并之后的前k小的元素中。换句话说，A[k/2-1]不可能大于两数组合并之后的第k小值，所以我们可以将其抛弃。

证明也很简单，可以采用反证法。假设A[k/2-1]大于合并之后的第k小值，我们不妨假定其为第（k+1）小值。
由于A[k/2-1]小于B[k/2-1]I，所以B[k/2-1]至少是第（k+2）小值。但实际上，在A中至多存在k/2-1个元素小于A[k/2-1]，B中也至多存在k/2-1个元素小于A[k/2-1]，所以小于A[k/2-1]的元素个数至多有k/2+ k/2-2，小于k，这与A[k/2-1]是第（k+1）的数矛盾。

当A[k/2-1]>B[k/2-1]时存在类似的结论。

当A[k/2-1]=B[k/2-1]时，我们已经找到了第k小的数，也即这个相等的元素，我们将其记为m。由于在A和B中分别有k/2-1个元素小于m，所以m即是第k小的数。(这里可能有人会有疑问，如果k为奇数，则m不是中位数。这里是进行了理想化考虑，在实际代码中略有不同，是先求k/2，然后利用k-k/2获得另一个数。)

通过上面的分析，我们即可以采用递归的方式实现寻找第k小的数。此外我们还需要考虑几个边界条件：

如果A或者B为空，则直接返回B[k-1]或者A[k-1]；
如果k为1，我们只需要返回A[0]和B[0]中的较小值；
如果A[k/2-1]=B[k/2-1]，返回其中一个；
图文详见《两个有序数组中的中位数和Top K问题》
'''

def medianofarrays( a, b ):
	m = len(a)
	n = len(b)
	
	if m > n :
		return medianofarrays(b, a)

	INT_MIN = -1
	INT_MAX = 65535
	lo = 0
	hi = m * 2

	while lo <= hi :
		c1 = (lo + hi) / 2
		c2 = m + n - c1
		L1 = INT_MIN if 0 == c1 else a[(c1 - 1) / 2]
		R1 = INT_MAX if c1 == m * 2 else a[c1 / 2]
		L2 = INI_MIN if 0 == c2 else b[(c2 - 1) / 2]
		R2 = INT_MAX if c2 == n * 2 else b[c2 / 2]

		if L1 > R2 :
			hi = c1 - 1
		elif L2 > R1 :
			lo = c1 + 1
		else :
			break

	return (max(L1, L2) + min(R1, R2)) / 2.0

list1 = (1, 3, 5)
list2 = (2, 4, 7)

print medianofarrays(list1, list2)
