class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        merge_array = list()
        i, j = 0, 0
        n_len = len(nums1)
        m_len = len(nums2)
        while i < n_len and j < m_len:
            if nums1[i] < nums2[j]:
                merge_array.append(nums1[i])
                i += 1
            else:
                merge_array.append(nums2[j])
                j += 1
        if i >= n_len:
            merge_array.extend(nums2[j:])
        else:
            merge_array.extend(nums1[i:])

        merge_len = n_len + m_len
        media = (merge_len - 1) / 2
        if merge_len % 2 == 0:
            return (merge_array[media] + merge_array[media + 1]) / 2.0
        else:
            return float(merge_array[media])
        