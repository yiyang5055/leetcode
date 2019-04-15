# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        result = None
        carrybit = 0
        current_node = None
        addend_node = l1
        summand_node = l2

        while addend_node or summand_node or carrybit:

            addend = addend_node.val if addend_node else 0
            summand = summand_node.val if summand_node else 0
            (carrybit, node_val) = divmod(addend + summand + carrybit, 10)

            if result:
                current_node.next = ListNode(node_val)
                current_node = current_node.next
            else:
                result = ListNode(node_val)
                current_node = result

            if addend_node:
                addend_node = addend_node.next
            if summand_node:
                summand_node = summand_node.next

        return result
        