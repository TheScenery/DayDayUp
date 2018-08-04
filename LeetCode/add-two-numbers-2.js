/*
    Title:2. Add Two Numbers
    Description:
    You are given two non-empty linked lists representing two non-negative integers.
    The digits are stored in reverse order and each of their nodes contain a single digit. 
    Add the two numbers and return it as a linked list.
    You may assume the two numbers do not contain any leading zero, except the number 0 itself. 
    Example:
    Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    Output: 7 -> 0 -> 8
    Explanation: 342 + 465 = 807.

    Language: Javascript
*/

class ListNode {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

function addTwoNumbers(l1, l2) {
    let isOver = false;
    let bitSum = 0;
    const result = new ListNode(bitSum);
    let resultTail = result;
    let l1Tail = l1;
    let l2Tail = l2;
    while (l1Tail || l2Tail) {
        const val1 = l1Tail && l1Tail.val || 0;
        const val2 = l2Tail && l2Tail.val || 0;
        bitSum = val1 + val2 + (isOver ? 1 : 0);
        isOver = false;
        if (bitSum >= 10) {
            bitSum -= 10;
            isOver = true;
        }
        resultTail.val = bitSum;
        l1Tail = l1Tail && l1Tail.next;
        l2Tail = l2Tail && l2Tail.next;
        if (l1Tail || l2Tail) {
            resultTail = resultTail.next = new ListNode(0);
        }
    }
    if (isOver) {
        resultTail.next = new ListNode(1);
    }
    return result;
}

function outPutList(l) {
    const valueArray = [];
    valueArray.push(l.val);
    let lTail = l;
    while (lTail.next) {
        lTail = lTail.next;
        valueArray.push(lTail.val);
    }
    console.log(valueArray.join(' -> '));
}

function testMain() {
    const l1 = new ListNode(2);
    let newNode = new ListNode(4);
    l1.next = newNode;
    let newNode1 = new ListNode(3);
    newNode.next = newNode1;

    const l2 = new ListNode(5);
    newNode = new ListNode(6);
    l2.next = newNode;
    newNode1 = new ListNode(4);
    newNode.next = newNode1;

    const l3 = addTwoNumbers(l1, l2);
    outPutList(l3);
}

testMain()