/*There are two sorted arrays nums1 and nums2 of size m and n respectively.

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
*/
/*
思路2：
两个有序数组长度之和为偶数：
两个a和b有序数组的长度分别是m和n。对a数组的一个划分i和对b数组的一个划分j，使得a数组的左半部分同b数组的左半部分看作合并后有序数组c的左半部分；而a数组的右半部分同b数组的右半部分看作合并后有序数组c的右半部分。假设有i、j存在，则需满足如下条件：
1、length(a[0, i)) + length(b[0, j]) = length(a[i, m)) + length(b[j, n])
(i-0)+(j-0)=(m-i)+(n-j)
i=(m+n)/2-j
2、max(a[i-1], b[j-1]) <= min(a[i], b[j])
则中位数是 (max(a[i-1], b[j-1]) + min(a[i], b[j])) / 2
两个有序数组长度之和为奇数：
两个a和b有序数组的长度分别是m和n。对a数组的一个划分i和对b数组的一个划分j，使得a数组的左半部分同b数组的左半部分看作合并后有序数组c的左半部分；而a数组的右半部分同b数组的右半部分看作合并后有序数组c的右半部分。假设有i、j存在，则需满足如下条件：
1、length(a[0, i)) + length(b[0, j]) + 1 = length(a[i, m)) + length(b[j, n])
(i-0)+(j-0)+1=(m-i)+(n-j)
i=(m+n-1)/2-j
2、max(a[i-1], b[j-1]) <= min(a[i], b[j])
则中位数是 min(a[i], b[j])
分析求解：
两个有序数组长度不管是基数还是偶数，都需要满足如上两个条件。其中第二个条件相同，而第一个条件我们可以归纳简化为i=(m+n)/2-j，因为当m+n为基数时 (m+n)/2-j = (m+n-1)/2-j。所以，只要找出数组a划分i，则可以通过公式i=(m+n)/2-j，得到数组b划分j，并且满足如下两个条件：
1、i=(m+n)/2-j
2、max(a[i-1], b[j-1]) <= min(a[i], b[j])
我们可以在数组a中，通过折半查找来取得划分i。其中IMIN为开始索引0，IMAX为数组a长度m
1）将数组a折半 i = (IMAX - IMIN) / 2
2）其中a[i-1]<=a[i]，b[j-1]<=b[j]一定是成立的。
若b[j-1]>a[i]不符合条件，i只有向后移动一位变大，相应的j向前移动一位变小，才有可能符合条件b[j-1]<=a[i]。所以，此时的i并不是有效划分，有效划分在i+1~IMAX范围内。IMIN赋值为i+1，继续执行1）。
若a[i-1]>b[j]不符合条件，i只有向前移动一位变小，相应的j向后移动一位变大，才有可能符合条件a[i-1]<=b[j]。所以，此时的i并不是有效划分，有效划分在IMIN~i-1范围内。IMAX为i-1，继续执行1）。
3）最后找到合适的i和j，就可以直接计算得到中位数。
注意：
1、数组a和数组b长度的关系
必须保证0<=i<=m, 0<=j<=n；假设0<=i<=m成立，如果满足0<=j<=n，m和n之间的关系。
1）
i=(m+n)/2-j =>
j=(m+n)/2-i =>
因为i<=m，所以j=(m+n)/2-i>=(m+n)/2-m=(n-m)/2 =>
若(n-m)/2>=0，则j>=0 =>
m<=n
2）
i=(m+n)/2-j =>
j=(m+n)/2-i =>
因为i>=0，j=(m+n)/2-i<=(m+n)/2 =>
若(m+n)/2<=n，则j<=n =>
m<=n
3）
总结：当m<=n时，0<=i<=m，可以保证0<=j<=n。所以，在算法中，始终要保持数组a的长度要不大于数组b的长度。
数组a可能为空数组
2、i，j移动过程中的越界问题
1）若b[j-1]>a[i]不符合条件，i只有向后移动一位变大，相应的j向前移动一位变小，才有可能符合条件b[j-1]<=a[i]。
为保证数组不越界，j不能等于0，i不能等于m。
这两种情况要单独判断：
当j==0时，左半部分的最大值是a[i-1]，右半部分的最小值是min(a[i], b[0])；
当i==m时，左半部分的最大值是max(a[m-1], b[j-1])，右半部分的最小值是b[j]。
2）若a[i-1]>b[j]不符合条件，i只有向前移动一位变小，相应的j向后移动一位变大，才有可能符合条件a[i-1]<=b[j]。所以，此时的i并不是有效划分，有效划分在0~i范围内。IMIN赋值为0，IMAX为i，继续执行1）。
为保证数组不越界，i不能等于0，j不能等于n。
这两种情况要单独判断：
当i==0时，左半部分的最大值是b[j-1]，右半部分的最小值是min(a[0], b[j])；
当j==n时，左半部分的最大值是max(a[i-1], b[n-1])，右半部分的最小值是a[i]。

*/

package main
import (
	"fmt"
	"math"
)

func medianofarrays (a []int, lena int, b []int, lenb int) float32 {
	fmt.Println(a[0], lena, b[0], lenb)

	if lena > lenb {
		return medianofarrays(b, lenb, a, lena)
	}

	var iMin, iMax int = 0, lena

	for iMin <= iMax {
		var i int = (iMax - iMin) / 2 + iMin
		var j int = (lena + lenb) / 2 - i

		if 0 != j && i != lena && b[j - 1] > a[i] {
			i++
			iMin = i
		} else if 0 != i && j != lenb && a[i - 1] > b[j] {
			i--
			iMax = i
		} else {
			var rMin int
			if 0 == lena {
				rMin = b[lenb / 2]
			} else if i == lena {
				rMin = b[j]
			} else if j == lenb {
				rMin = a[i]
			} else if 0 == j {
				rMin = int(math.Min(float64(a[i]), float64(b[0])))
			} else if 0 == i {
				rMin = int(math.Min(float64(a[0]), float64(b[j])))
			} else {
				rMin = int(math.Min(float64(a[i]), float64(b[j])))
			}

			if 0 != (lena + lenb) % 2 {
				return float32(rMin)
			}


			var lMax int
			if 0 == lena {
				lMax = b[lenb / 2 - 1]
			} else if i == lena {
				lMax = int(math.Max(float64(a[lena - 1]), float64(b[j - 1])))
			} else if j == lenb {
				lMax = int(math.Max(float64(a[i - 1]), float64(b[lenb - 1])))
			} else if 0 == j {
				lMax = a[i - 1]
			} else if 0 == i {
				lMax = b[j - 1]
			} else {
				lMax = int(math.Max(float64(a[i - 1]), float64(b[j - 1])))
			}

			return float32(rMin + lMax) / 2.0
		}
	}

	return 0.0
}

func main () {
	var arr1 = []int {1, 3, 7}
	var arr2 = []int {2, 4}

	fmt.Println(medianofarrays(arr1, 3, arr2, 2))
}
