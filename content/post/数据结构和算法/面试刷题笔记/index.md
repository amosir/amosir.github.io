---
title: "面试刷题笔记"
date: 2021-10-28T13:37:07+08:00
draft: false
image: "img/Nowd8MgUK8A.jpg"
categories: 
tag:
---


# 面试刷题笔记

## 地下城游戏

https://leetcode-cn.com/problems/dungeon-game/

```java
class Solution {

    public int calculateMinimumHP(int[][] dungeon) {
        int rows = dungeon.length;
        int cols = dungeon[0].length;
        int[][] dp = new int[rows][cols];
        dp[rows - 1][cols - 1] = dungeon[rows - 1][cols - 1] > 0 ? 0 : -dungeon[rows - 1][cols - 1];
        for(int i = rows - 2;i >= 0;i--){
            dp[i][cols - 1] = Math.max(dp[i + 1][cols - 1] - dungeon[i][cols - 1], 0); 
        }
        for(int j = cols - 2;j >= 0;j--){
            dp[rows - 1][j] = Math.max(dp[rows - 1][j + 1] - dungeon[rows - 1][j], 0); 
        }

        for(int i = rows - 2;i >= 0;i--){
            for(int j = cols - 2;j >= 0;j--){
                dp[i][j] = Math.max(0, Math.min(dp[i][j + 1], dp[i+1][j]) - dungeon[i][j]);
            }
        }
        return dp[0][0] + 1;
    }
}
```

## 最大正方形

https://www.nowcoder.com/practice/0058c4092cec44c2975e38223f10470e?tpId=117&tqId=37832&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int solve (char[][] matrix) {
    if(matrix == null || matrix.length == 0 || matrix[0].length == 0) return 0;
    // dp[i][j]表示矩阵中以坐标(i,j)为右下角的最大正方形边长
    int[][] dp = new int[matrix.length][matrix[0].length];
    // 最大边长
    int maxLen = 0;
    for(int i = 0;i < matrix.length;i++){
        for(int j = 0;j < matrix[0].length;j++){
            if(matrix[i][j] == '1'){
                if(i == 0 || j == 0){
                    dp[i][j] = 1;
                }else{
                    dp[i][j] = Math.min(Math.min(dp[i - 1][j],dp[i - 1][j - 1]),dp[i][j - 1]) + 1;
                }
            }
            if(maxLen < dp[i][j]) maxLen = dp[i][j];
        }
    }
    return maxLen * maxLen;
}
```

## 大数加法

https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475?tpId=117&tqId=37842&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public String solve (String s, String t) {
    StringBuilder sb = new StringBuilder();
    int i = s.length() - 1, j = t.length() - 1;
    int sum = 0;
    int shift = 0;
    while(i >= 0 && j >= 0){
        sum = (s.charAt(i) - '0') + (t.charAt(j) - '0') + shift;
        if(sum >= 10){
            shift = 1;
            sb.append(sum - 10);
        }else{
            sb.append(sum);
            shift = 0;
        }
        i--;
        j--;
    }
    while(i >= 0){
        sum = (s.charAt(i) - '0') + shift;
        if(sum >= 10){
            shift = 1;
            sb.append(sum - 10);
        }else{
            sb.append(sum);
            shift = 0;
        }
        i--;
    }
    while(j >= 0){
        sum = (t.charAt(j) - '0') + shift;
        if(sum >= 10){
            shift = 1;
            sb.append(sum - 10);
        }else{
            sb.append(sum);
            shift = 0;
        }
        j--;
    }
    if(shift > 0){
        sb.append(shift);
    }
    return sb.reverse().toString();
}
```

## 字符串乘法

https://leetcode-cn.com/problems/multiply-strings/

```java
class Solution {
    public String multiply(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0")) {
            return "0";
        }
        int m = num1.length(), n = num2.length();
        int[] ansArr = new int[m + n];
        for (int i = m - 1; i >= 0; i--) {
            int x = num1.charAt(i) - '0';
            for (int j = n - 1; j >= 0; j--) {
                int y = num2.charAt(j) - '0';
                ansArr[i + j + 1] += x * y;
            }
        }
        for (int i = m + n - 1; i > 0; i--) {
            ansArr[i - 1] += ansArr[i] / 10;
            ansArr[i] %= 10;
        }
        int index = ansArr[0] == 0 ? 1 : 0;
        StringBuffer ans = new StringBuffer();
        while (index < m + n) {
            ans.append(ansArr[index]);
            index++;
        }
        return ans.toString();
    }
}
```



## 链表相加

https://leetcode-cn.com/problems/add-two-numbers

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode cur1 = l1, cur2 = l2;
        ListNode dummy = new ListNode();
        ListNode cur, end = dummy;
        int carry = 0;
        while(cur1 != null && cur2 != null){
            int sum = cur1.val + cur2.val + carry;
            if(sum > 9){
                sum = sum - 10;
                carry = 1;
            }else{
                carry = 0;
            }
            cur = new ListNode(sum);
            end.next = cur;
            end = cur;
            cur1 = cur1.next;
            cur2 = cur2.next;
        }

        while(cur1 != null){
            int sum = cur1.val + carry;
            if(sum > 9){
                sum = sum - 10;
                carry = 1;
            }else{
                carry = 0;
            }
            cur = new ListNode(sum);
            end.next = cur;
            end = cur;
            cur1 = cur1.next;
        }

        while(cur2 != null){
            int sum = cur2.val + carry;
            if(sum > 9){
                sum = sum - 10;
                carry = 1;
            }else{
                carry = 0;
            }
            cur = new ListNode(sum);
            end.next = cur;
            end = cur;
            cur2 = cur2.next;
        }
        if(carry != 0){
            cur = new ListNode(carry);
            end.next = cur;
            end = cur;
        }
        return dummy.next;
    }
}
```



## 重排链表

https://leetcode-cn.com/problems/reorder-list/

```java
public void reorderList(ListNode head) {
    if(head == null || head.next == null) return;
    // 快慢指针找中点
    ListNode fast = head, slow = head;
    // 后面一条链表的头结点
    ListNode otherHead = null;
    while(fast.next != null && fast.next.next != null){
        fast = fast.next.next;
        slow = slow.next;
    }
    // 分割成两条链表
    otherHead = slow.next;
    slow.next = null;

    // 翻转第二条链表
    ListNode cur = otherHead.next;
    otherHead.next = null;
    while(cur != null){
        ListNode next = cur.next;
        cur.next = otherHead;
        otherHead = cur;
        cur = next;
    }

    // 合并链表
    // 第一条链表的当前节点
    cur = head;
    while(otherHead != null){
        // 第二条链表当前节点的后一个节点
        ListNode otherTemp = otherHead.next;
        otherHead.next = cur.next;
        cur.next = otherHead;
        cur = otherHead.next;
        otherHead = otherTemp;
    }
}
```

## 奇偶链表

https://leetcode-cn.com/problems/odd-even-linked-list/

```java
class Solution {
    public ListNode oddEvenList(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode oddPos = head, evenPos = head.next, evenHead = evenPos;
        while(oddPos.next != null && evenPos.next != null){
            oddPos.next = evenPos.next;
            oddPos = oddPos.next;
            evenPos.next = oddPos.next;
            evenPos = evenPos.next;
        }
        oddPos.next = evenHead;
        return head;
    }
}
```



## 复杂链表的复制

https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    public Node copyRandomList(Node head) {
        if(head == null) return head;
        Node cur = head;
        while(cur != null){
            Node temp = cur.next;
            Node newNode = new Node(cur.val);
            newNode.next = temp;
            cur.next = newNode;
            cur = temp;
        }
        cur = head;
        while(cur != null){
            if(cur.random != null){
                cur.next.random = cur.random.next;
            }
            cur = cur.next.next;
        }
        Node newHead = head.next, pos1 = head, pos2 = newHead;
        while(pos2.next != null){
            pos1.next = pos2.next;
            pos1 = pos1.next;
            pos2.next = pos1.next;
            pos2 = pos2.next;
        }
        pos1.next = null;
        return newHead;
    }
}
```

## 二叉树所有路径之和

https://www.nowcoder.com/practice/185a87cd29eb42049132aed873273e83?tpId=117&tqId=37715&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
static class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;
    public TreeNode(int val){
        this.val = val;
    }
}

public int sumNumbers (TreeNode root) {
    if(root == null) return 0;
    return helper(root,0);
}

private int helper(TreeNode root, int sum){
    if(root.left == null && root.right == null){
        return sum * 10 + root.val;
    }else{
        int ret = 0;
        if(root.left != null){
            ret += helper(root.left,sum * 10 + root.val);
        }
        if(root.right != null){
            ret += helper(root.right, sum * 10 + root.val);
        }
        return ret;
    }

}
```

## 根节点到叶子节点的数字之和

https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

```java
class Solution {
    public int sumNumbers(TreeNode root) {
        return preOrderTraverse(root, 0);
    }
    private int preOrderTraverse(TreeNode root, int temp){
        int sum = 0;
        if(root == null){
            return 0;
        }else if(root.left == null && root.right == null){
            sum = temp * 10 + root.val;
        }else{
            temp = temp * 10 + root.val;
            sum += preOrderTraverse(root.left, temp);
            sum += preOrderTraverse(root.right, temp);
        }
        return sum;
    }
}
```

## 买股票的最佳时机(一次交易)

https://www.nowcoder.com/practice/64b4262d4e6d4f6181cd45446a5821ec?tpId=117&tqId=37717&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int maxProfit (int[] prices) {
    if(prices == null || prices.length == 0) return 0;
    int minPrice = prices[0], maxProfit = 0;
    for(int i = 1;i < prices.length;i++){
        minPrice = Math.min(prices[i], minPrice);
        maxProfit = Math.max(prices[i] - minPrice, maxProfit);
    }
    return maxProfit;
}
```

## **股票交易的最大收益（二次交易）**

https://www.nowcoder.com/practice/4892d3ff304a4880b7a89ba01f48daf9?tpId=117&tqId=37847&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int maxProfit (int[] prices) {
    if(prices == null || prices.length == 0) return 0;
    // buy1表示一天结束时只进行了一次买入
    // sell1表示一天结束时只进行了一次买入和一次卖出
    // buy2表示一天结束时完成了一笔交易，进行了第二次买入
    // sell2表示一天结束时完成了两笔交易
    int buy1 = -prices[0], sell1 = 0, buy2 = -prices[0], sell2 = 0;
    for(int i = 1;i < prices.length;i++){
        // 在前一天的基础上什么都不做或者在没有任何操作的情况下进行一次买入
        buy1 = Math.max(buy1,-prices[i]);
        // 在前一天的基础上什么都不做或者在有一次买入的情况下进行一次卖出
        sell1 = Math.max(buy1 + prices[i], sell1);
        // 在前一天的基础上什么都不做或者在完成一次交易的情况下再进行一次买入
        buy2 = Math.max(buy2, sell1 - prices[i]);
        // 在前一天的基础上什么都不做或者在进入第二次买入的情况下进行一次卖出
        sell2 = Math.max(sell2, buy2 + prices[i]);
    }
    return sell2;
}
```

## **股票(无限次交易)**

https://www.nowcoder.com/practice/9e5e3c2603064829b0a0bbfca10594e9?tpId=117&tqId=37846&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
class Solution {
    public int maxProfit(int[] prices) {
        int[][] dp = new int[prices.length][2];
        dp[0][0] = 0;
        dp[0][1] = -prices[0];
        for(int i = 1;i < prices.length;i++){
            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
            dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][0] - prices[i]);
        }
        return dp[prices.length - 1][0];
    }
}
```

## 打家劫舍1

https://leetcode-cn.com/problems/house-robber/

```java
class Solution {

    
    public int rob(int[] nums) {
        if(nums.length == 1) return nums[0];
        int dp_0 = 0;
        int dp_1 = nums[0];
        for(int i = 2;i <= nums.length;i++){
            int dp_i = Math.max(dp_0 + nums[i - 1], dp_1);
            dp_0 = dp_1;
            dp_1 = dp_i;
        }
        return dp_1;
    }
}
```

## 打家劫舍2

https://leetcode-cn.com/problems/house-robber-ii/

```java
class Solution {
    public int rob(int[] nums) {
        if(nums.length == 1) return nums[0];
        return Math.max(rob(nums,0,nums.length - 2), rob(nums,1,nums.length - 1));
    }

    public int rob(int[] nums, int start, int end) {
        if(nums.length == 1) return nums[0];
        int dp_0 = 0;
        int dp_1 = nums[start];
        for(int i = start + 1;i <= end;i++){
            int dp_i = Math.max(dp_0 + nums[i], dp_1);
            dp_0 = dp_1;
            dp_1 = dp_i;
        }
        return dp_1;
    }
}
```

## 打家劫舍3

https://leetcode-cn.com/problems/house-robber-iii/

```java
class Solution {
    private HashMap<TreeNode, Integer> memo = new HashMap();
    public int rob(TreeNode root) {
        if(root == null) return 0;
        if(memo.containsKey(root)) return memo.get(root);
        int do_it = root.val + (root.left == null ? 0 : rob(root.left.left) + rob(root.left.right)) + (root.right == null ? 0 : rob(root.right.left) + rob(root.right.right));
        int not_do = rob(root.left) + rob(root.right);
        int res = Math.max(do_it, not_do);
        memo.put(root, res);
        return res;
    }
}
```

## 路径总和

https://leetcode-cn.com/problems/path-sum/

```java
class Solution {
    public boolean hasPathSum(TreeNode root, int targetSum) {
        if(root == null){
            return false;
        }else if(root.left == null && root.right == null){
            return targetSum == root.val;
        }else{
            return hasPathSum(root.left, targetSum - root.val) || hasPathSum(root.right, targetSum - root.val);
        }
    }
}
```



## 路径总和2

https://leetcode-cn.com/problems/path-sum-ii/submissions/

```java
class Solution {
    List<List<Integer>> ret;
    ArrayList<Integer> path;
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        ret = new ArrayList();
        path = new ArrayList();
        helper(root, targetSum);
        return ret;
    }

    private void helper(TreeNode root, int sum){
        if(root == null) return;
        sum -= root.val;
        path.add(root.val);
        if(root.left == null && root.right == null && sum == 0){
            ret.add(new ArrayList(path));
        }
        helper(root.left, sum);
        helper(root.right, sum);
        sum += root.val;
        path.remove(path.size() - 1);
    }
}
```

## 二叉树的最大深度

https://www.nowcoder.com/practice/8a2b2bf6c19b4f23a9bdb9b233eefa73?tpId=117&tqId=37721&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int maxDepth (TreeNode root) {
    if(root == null) return 0;
    int depthLeft = maxDepth(root.left);
    int depthRight = maxDepth(root.right);
    return depthLeft > depthRight ? depthLeft + 1 : depthRight + 1;
}
```

## 二叉树的最大宽度

https://leetcode-cn.com/problems/maximum-width-of-binary-tree/

```java
class Solution {
   public int widthOfBinaryTree(TreeNode root) {
        if(root == null) return 0;
        int maxWidth = 1;
        root.val = 0;
        Deque<TreeNode> queue = new LinkedList<>();
        queue.addFirst(root);
        while(!queue.isEmpty()){
            maxWidth = Math.max(maxWidth, queue.getFirst().val - queue.getLast().val + 1);
            for(int i = queue.size();i > 0;i--){
                TreeNode node = queue.pollLast();
                if(node.left != null){
                    node.left.val = 2 * node.val + 1;
                    queue.addFirst(node.left);
                }
                if(node.right != null){
                    node.right.val = 2 * node.val + 2;
                    queue.addFirst(node.right);
                }
            }
        }
        return maxWidth;
    }
}
```

## 平衡二叉树

https://leetcode-cn.com/problems/balanced-binary-tree/

```java
class Solution {
    public boolean isBalanced(TreeNode root) {
        if(root == null) return true;
        if(root.left == null && root.right == null) return true;
        int minus = getHeight(root.left) - getHeight(root.right);
        return isBalanced(root.left) && isBalanced(root.right) && minus <= 1 && minus >= -1;
    }

    private int getHeight(TreeNode root){
        if(root == null) return 0;
        if(root.left == null && root.right == null) return 1;
        int left = getHeight(root.left);
        int right = getHeight(root.right);
        return Math.max(left, right) + 1;
    }
}
```



## **输出二叉树的右视图**

https://www.nowcoder.com/practice/c9480213597e45f4807880c763ddd5f0?tpId=117&tqId=37848&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
private ArrayList<Integer> levelOrder(TreeNode root){
    if(root == null) return new ArrayList<>(0);
    // 返回
    ArrayList<Integer> ret = new ArrayList<>();

    // 辅助队列
    Queue<TreeNode> queue = new LinkedList<>();
    queue.add(root);

    // 记录每一层的节点数
    int layerCnt = 1;
    while(!queue.isEmpty()){
        TreeNode temp = queue.poll();
        layerCnt --;

        // 子节点入队列
        if(temp.left != null){
            queue.add(temp.left);
        }

        if(temp.right != null){
            queue.add(temp.right);
        }

        // 当前处理的节点是该层最后一个节点
        if(layerCnt == 0) {
            ret.add(temp.val);
            layerCnt = queue.size();
        }

    }
    return ret;
}
```

## 之字形遍历二叉树

https://www.nowcoder.com/practice/47e1687126fa461e8a3aff8632aa5559?tpId=117&tqId=37722&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ArrayList<ArrayList<Integer>> zigzagLevelOrder(TreeNode root) {
    if(root == null) return new ArrayList<>(0);
    ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
    // 初始保存第一层
    LinkedList<Integer> layer = new LinkedList<>();
    // 每一层的节点数
    int nodeCnt = 1;
    // 奇数层还是偶数层
    boolean odd = true;
    Queue<TreeNode> queue = new LinkedList<>();
    queue.offer(root);
    while (!queue.isEmpty()){
        TreeNode temp = queue.poll();
        nodeCnt --;
        // 当前奇数层, 尾插法
        if(odd){
            layer.addLast(temp.val);
        }
        // 当前偶数层, 头插法
        else{
            layer.addFirst(temp.val);
        }
        if(temp.left!=null){
            queue.offer(temp.left);
        }
        if(temp.right!=null){
            queue.offer(temp.right);
        }
        if(nodeCnt == 0){
            ret.add(new ArrayList<>(layer));
            nodeCnt = queue.size();
            layer = new LinkedList<>();
            odd = !odd;
        }
    }
    return ret;
}

```

## 树的非递归遍历

```java
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        if(root == null) return new ArrayList(0);
        ArrayList<Integer> ret = new ArrayList();
        Stack<TreeNode> stack = new Stack();
        stack.push(root);
        while(!stack.isEmpty()){
            TreeNode cur = stack.pop();
            ret.add(cur.val);
            if(cur.right != null){
                stack.add(cur.right);
            }
            if(cur.left != null){
                stack.add(cur.left);
            }
        }
        return ret;
    }

		public List<Integer> inorderTraversal(TreeNode root) {
		        if(root == null) return new ArrayList(0);
		        TreeNode cur = root;
		        Stack<TreeNode> stack = new Stack();
		        List<Integer> ret = new ArrayList();
		        while(!stack.isEmpty() || cur != null){
		            while(cur != null){
		                stack.push(cur);
		                cur = cur.left;
		            }
		            cur = stack.pop();
		            ret.add(cur.val);
		            cur = cur.right;
		        }
		        return ret;
    }

		public List<Integer> postorderTraversal(TreeNode root) {
        if(root == null) return new ArrayList(0);
        ArrayList<Integer> ret = new ArrayList();
        TreeNode cur = root;
        TreeNode lastVisited = null;
        Stack<TreeNode> stack = new Stack();
        while(!stack.isEmpty() || cur != null){
            while(cur != null){
                stack.push(cur);
                cur = cur.left;
            }
            cur = stack.peek();
            if(cur.right != null && cur.right != lastVisited){
                cur = cur.right;
            }else{
                lastVisited = cur;
                ret.add(cur.val);
                stack.pop();
                cur = null;
            }
        }
        return ret;
    }
  
    public List<List<Integer>> levelOrder(TreeNode root) {
        if(root == null) return new ArrayList(0);
        List<List<Integer>> ret = new ArrayList();
        List<Integer> layer = new ArrayList();
        int cnt = 1;
        Queue<TreeNode> queue = new LinkedList();
        queue.add(root);
        while(!queue.isEmpty()){
            TreeNode node = queue.poll();
            layer.add(node.val);
            cnt--;
            if(node.left != null){
                queue.add(node.left);
            }
            if(node.right != null){
                queue.add(node.right);
            }
            if(cnt == 0){
                ret.add(layer);
                cnt = queue.size();
                layer = new ArrayList(queue.size());
            }            
        }
        return ret;
    }
}
```



## 二叉树转化为链表

https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/

```java
class Solution {
    public Node treeToDoublyList(Node root) {
        if(root == null) return root;
        Stack<Node> stack = new Stack();
        Node cur = root;
        Node head = null, tail = null;
        while(!stack.isEmpty() || cur != null){
            if(cur != null){
                stack.push(cur);
                cur = cur.left;
            }else{
                cur = stack.pop();
                if(head == null){
                    head = cur;
                    tail = cur;
                }else{
                    tail.right = cur;
                    cur.left = tail;
                    tail = cur;
                }
                cur = cur.right;
            }
        }
        head.left = tail;
        tail.right = head;
        return head;
    }
}
```

## 判断完全二叉树

https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/

```java
class Solution {
    public boolean isCompleteTree(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList();
        queue.add(root);
        while(!queue.isEmpty()){
            TreeNode cur = queue.poll();
            if(cur != null){
                queue.add(cur.left);
                queue.add(cur.right);
            }else{
                while(!queue.isEmpty() && queue.peek() == null){
                    queue.poll();
                }
                if(!queue.isEmpty()){
                    return false;
                }
            }
        }
        return true;
    }
}
```



## 最长回文子串

https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af?tpId=117&tqId=37789&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int getLongestPalindrome(String A, int n) {
    if(A== null || n == 0) return 0;
    boolean[][] dp = new boolean[n][n];
    // 记录出现的最大长度
    int maxLen = 1;
    // len + 1表示字串的长度
    for(int len = 0;len < n;len++){
        // 子串的开始位置
        for(int i = 0;i < n - len;i++){
            // 子串的结束位置
            int j = i + len;
            if((A.charAt(i) == A.charAt(j)) && (j - i < 2 || dp[i+1][j-1])) {
                dp[i][j] = true;
                if(len + 1 > maxLen) maxLen = len + 1;
            }
        }
    }
    return maxLen;
}
```

## 顺时针旋转矩阵

https://www.nowcoder.com/practice/2e95333fbdd4451395066957e24909cc?tpId=117&tqId=37790&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
// 两次翻转
public int[][] rotateMatrix_v1(int[][] mat, int n) {
    // 右对角线翻转
    // 注意只需要翻转对称轴左边的
    for(int i = 0;i < n;i++){
        for(int j = 0;j < n - i - 1;j ++){
            mat[i][j] =  mat[i][j] ^  mat[n - j - 1][n - i - 1];
            mat[n - j - 1][n - i - 1] = mat[i][j] ^  mat[n - j - 1][n - i - 1];
            mat[i][j] =  mat[i][j] ^  mat[n - j - 1][n - i - 1];
        }
    }

    // 横向对称轴翻转
    for(int i = 0;i < n / 2;i++){
        for(int j = 0;j < n;j++){
            mat[i][j] =  mat[i][j] ^  mat[n - 1 - i][j];
            mat[n - 1- i][j] =  mat[i][j] ^  mat[n - 1 - i][j];
            mat[i][j] =  mat[i][j] ^  mat[n - 1 - i][j];
        }
    }
    return mat;
}

// 直接翻转
public int[][] rotateMatrix(int[][] mat, int n) {
    int[][] ret = new int[n][n];
    for(int i = 0;i < n;i++){
        for(int j = 0;j < n;j ++){
            // 注意方向
            ret[i][j] = mat[j][n - i - 1];
        }
    }
    return ret;
}
```

## 旋转数组

```java
class Solution {
    public void rotate(int[] nums, int k) {
        k = k % nums.length;
        rotate(nums, 0, nums.length - 1);
        rotate(nums, 0, k - 1);
        rotate(nums, k, nums.length - 1);
    }

    private void rotate(int[] nums, int left, int right) {
        int temp;
        while(left < right){
            temp = nums[left];
            nums[left] = nums[right];
            nums[right] = temp;
            left++;
            right--;
        }
    }
}
```

## 旋转数组的最小值

https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

```java
class Solution {
    public int findMin(int[] nums) {
        int left = 0, right = nums.length - 1;
        while(left < right){
            int mid = left + (right - left) / 2;
            if(nums[right] < nums[mid]){
                left = mid + 1;
            }else if(nums[right] > nums[mid]){
                right = mid;
            }else{
                right--;
            }
        }
        return nums[left];
    }
}
```

## 顺时针打印矩阵

https://leetcode-cn.com/problems/spiral-matrix-ii/

```java
class Solution {
    public int[][] generateMatrix(int n) {
        int top = 0, bottom = n - 1, left = 0, right = n - 1;
        int[][] matrix = new int[n][n];
        int num = 1;
        while(top <= bottom && left <= right){
            for(int i = left;i <= right;i++){
                matrix[top][i] = num++;
            }
            top++;
            for(int i = top;i <= bottom;i++){
                matrix[i][right] = num++;
            }
            right--;
            if(left <= right && top <= bottom){
                for(int i = right;i >= left;i--){
                    matrix[bottom][i] = num++;
                }
                bottom--;
                for(int i = bottom;i >= top;i--){
                    matrix[i][left] = num++;
                }
                left++;
            }
        }
        return matrix;
    }
}
```

## 最大子数组和

https://www.nowcoder.com/practice/554aa508dd5d4fefbf0f86e5fe953abd?tpId=117&tqId=37797&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int maxsumofSubarray (int[] arr) {
    // dp[i]表示以第i个元素结尾的子数组的最大和
    int[] dp = new int[arr.length];
    dp[0] = arr[0];
    int max = dp[0];
    for(int i = 1;i < arr.length;i++){
        if(dp[i - 1] + arr[i] >= arr[i]){
            dp[i] = dp[i - 1] + arr[i];
        }else{
            dp[i] = arr[i];
        }
        if(dp[i] > max){
            max = dp[i];
        }
    }
    return max;
}
```

## 数字字符串转化为IP地址

https://www.nowcoder.com/practice/ce73540d47374dbe85b3125f57727e1e?tpId=117&tqId=37725&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ArrayList<String> restoreIpAddresses (String s) {
    // 注意不能返回null
    if(s == null || s.length() < 4) return new ArrayList<>(0);
    ArrayList<String> ret = new ArrayList<>();
    helper(ret,"",s,0);
    return ret;
}

public void helper(ArrayList<String> ret, String sb, String s, int dotCnt){
    // 三个点时前三部分已经验证通过，只需要验证最后一部分符合
    if(dotCnt == 3){
        if(Integer.valueOf(s) > 255 || (s.length() > 1 && s.charAt(0) == '0')) return;
        sb += s;
        ret.add(sb);
        return;
    }
    // 从当前剩余的字符中选择
    for(int i = 1;i < 4 && i < s.length();i++){
        String temp = s.substring(0, i);
        // 验证符合要求
        if(Integer.valueOf(temp) > 255 || (temp.length() > 1 && temp.charAt(0) == '0')) return;
        helper(ret,sb + temp + ".",s.substring(i),dotCnt + 1);
    }
}
```

## 翻转链表

https://leetcode-cn.com/problems/reverse-linked-list/

```java
class Solution {
    public ListNode reverseList(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode cur = head;
        head = null;
        while(cur != null){
            ListNode temp = cur.next;
            cur.next = head;
            head = cur;
            cur = temp;
        }
        return head;
    }
}
```

## 翻转链表2

https://leetcode-cn.com/problems/reverse-linked-list-ii/

```java
class Solution {
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if(head == null || head.next == null) return head;
        ListNode dummy = new ListNode(-1), pre = dummy;
        dummy.next = head;
        for(int i = 1;i < left ;i++){
            pre = pre.next;
        }
        ListNode cur = pre.next;
        for(int i = left;i < right;i++){
            ListNode next = cur.next;
            cur.next = next.next;
            next.next = pre.next;
            pre.next = next;
        }
        return dummy.next;
    }
}
```

## 旋转链表

https://leetcode-cn.com/problems/rotate-list/

```java
public class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if(head == null || head.next == null) return head;
        int len = 1;
        ListNode cur;
        for(cur = head;cur != null && cur.next != null;cur = cur.next){
            len++;
        }
        // 闭合成环
        cur.next = head;
        k = k % len;
        cur = head;
        for(int i = 1;i < len - k;i++){
            cur = cur.next;
        }
        head = cur.next;
        cur.next = null;
        return head;
    }
}
```

## 链表内指定区间翻转

https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c?tpId=117&tqId=37726&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ListNode reverseBetween(ListNode head, int m, int n) {
    ListNode dummy = new ListNode();
    dummy.next = head;
    ListNode pre = dummy;
    // 找到翻转开始位置的前一个节点
    for(int i = 1;i < m;i++){
        pre = pre.next;
    }
    ListNode cur = pre.next, temp;
    // 翻转m, n之间的节点
    for(int j = m;j < n;j++){
        temp = cur.next;
        cur.next = temp.next;
        temp.next = pre.next;
        pre.next = temp;
    }
    return dummy.next;
}
```

## 删除有序链表中重复出现的元素

https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024?tpId=117&tqId=37729&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ListNode deleteDuplicates (ListNode head) {
    if(head == null || head.next == null) return head;

    // 伪头结点
    ListNode dummy = new ListNode(-1);
    dummy.next = head;
    // 前面的节点和当前节点
    ListNode pre = dummy, cur = dummy.next;
    while(cur != null && cur.next != null){
        // 当前节点和后面的节点不相等时开始处理下一个节点
        if(cur.val != cur.next.val){
            cur = cur.next;
            pre = pre.next;
        }else{
            ListNode temp = cur.next;
            // 找到第一个和当前节点不等的节点
            while(temp != null && temp.val == cur.val){
                temp = temp.next;
            }
            // 剪除之间的节点
            pre.next = temp;
            cur = temp;
        }
    }
    return dummy.next;
}
```

## 删除有序链表中重复出现的元素2

https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/

```java
class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode dummy = new ListNode(-1);
        dummy.next = head;
        ListNode cur = dummy;
        while(cur.next != null && cur.next.next != null){
            if(cur.next.val != cur.next.next.val){
                cur = cur.next;
            }else{
                ListNode temp = cur.next;
                while(temp != null && temp.val == cur.next.val){
                    temp = temp.next;
                }
                cur.next = temp;
            }
        }
        return dummy.next;
    }
}
```



## 翻转字符串里的单词

https://leetcode-cn.com/problems/reverse-words-in-a-string/

```java
class Solution {
    public String reverseWords(String s) {
        StringBuilder sb = new StringBuilder();
        int left = 0, right = s.length() - 1;
        while(left <= right && s.charAt(left) == ' '){
           left++;
        }
        while(right >= 0 && s.charAt(right) == ' '){
           right--;
        }
        for(int i = left;i <= right;i++){
            char ch = s.charAt(i);
            if(ch != ' '){
                sb.append(ch);
            }else if(s.charAt(i - 1) != ' '){
                sb.append(ch);
            }
        }
        reverseString(sb, 0, sb.length() - 1);
        reverseAllWords(sb, 0, sb.length() - 1);
        return sb.toString();

    }

    private void reverseString(StringBuilder sb, int start, int end){
        while(start < end){
            char temp = sb.charAt(start);
            sb.setCharAt(start, sb.charAt(end));
            sb.setCharAt(end, temp);
            start++;
            end--;
        }
    }

    private void reverseAllWords(StringBuilder sb, int start, int end){
        int left = start, right = start;
        while(right <= end){
            while(right <= end && sb.charAt(right) != ' '){
                right++;
            }
            reverseString(sb, left, right - 1);
            left = right + 1;
            right = left;
        }
    }
}
```

## 单词搜索

https://leetcode-cn.com/problems/word-search/

````java
class Solution {
    public boolean exist(char[][] board, String word) {
        boolean find = false;
        for(int i = 0;i < board.length;i++){
            for(int j = 0;j < board[0].length;j++){
                if(find){
                    return true;
                }else{
                    find = find(board, i,j, word, 0);
                }
            }
        }
        return find;
    }

    private boolean find(char[][] board, int i, int j, String word, int pos) {
        if (pos == word.length()) {
            return true;
        }
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] != word.charAt(pos))
            return false;
        char ch = board[i][j];
        board[i][j] = '.';
        boolean res = find(board, i + 1, j, word, pos + 1) || find(board, i, j + 1, word, pos + 1) ||
                find(board, i - 1, j, word, pos + 1) || find(board, i, j - 1, word, pos + 1);
        board[i][j] = ch;
        return res;
    }
}
````

## 单词拆分

https://leetcode-cn.com/problems/word-break/

```java
class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        // dp[i]表示前0位是否可以表示
        boolean[] dp = new boolean[s.length() + 1];
        // 空串可以表示
        dp[0] = true;
        for(int i = 0;i < s.length();i++){
            for(int j = i + 1;j <= s.length();j++){
                if(dp[i] && wordDict.contains(s.substring(i, j))){
                    dp[j] = true;
                }
            }
        }
        return dp[s.length()];
    }
}
```

## K个一组翻转链表

https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e?tpId=117&tqId=37746&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC50 {

    class ListNode {
        int val;
        ListNode next = null;

        public ListNode(int val){
            this.val = val;
        }
    }

    public ListNode reverseKGroup(ListNode head, int k) {
        if(head == null || head.next == null || k == 1) return head;
        // 长度
        int len = 0;
        ListNode temp = head;
        while (temp != null){
            len ++;
            temp = temp.next;
        }
        if(len < k) return head;
        ListNode dummy = new ListNode(0), pre = dummy, cur = head, next;
        dummy.next = head;
        // 长度为K的子链的数量
        int groups = len / k;
        for(int i = 0;i < groups;i++){
            // 翻转子链
            for(int j = 1;j < k;j++){
                next = cur.next;
                cur.next = next.next;
                next.next = pre.next;
                pre.next = next;
            }
            pre = cur;
            cur = cur.next;
        }
        return dummy.next;
    }
}
```



## 括号生成

https://www.nowcoder.com/practice/c9addb265cdf4cdd92c092c655d164ca?tpId=117&tqId=37748&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC26 {
    public ArrayList<String> generateParenthesis (int n) {
        ArrayList<String> ret = new ArrayList<>();
        helper(ret, new StringBuilder(), 0 , 0, n);
        return ret;
    }

    private void helper(ArrayList<String> ret, StringBuilder sb, int openBrace, int closeBrace, int n){
        if(sb.length() == n + n){
            ret.add(sb.toString());
            return;
        }
        if(openBrace < n){
            sb.append('(');
            helper(ret, sb, openBrace + 1, closeBrace, n);
            sb.deleteCharAt(sb.length() - 1);
        }
        if(closeBrace < openBrace){
            sb.append(')');
            helper(ret, sb, openBrace, closeBrace + 1, n);
            sb.deleteCharAt(sb.length() - 1);
        }
    }
}
```

## 最长有效括号

https://leetcode-cn.com/problems/longest-valid-parentheses/

```java
class Solution {
    public int longestValidParentheses(String s) {
        if(s == null || s.length() == 0) return 0;
        Stack<Integer> stack = new Stack();
        stack.push(-1);
        int maxLen = 0;
        for(int i = 0;i < s.length();i++){
            char ch = s.charAt(i);
            if(ch == '('){
                stack.push(i);
            }else{
                stack.pop();
                if(!stack.isEmpty()){
                    maxLen = Math.max(maxLen, i - stack.peek());
                }else{
                    stack.push(i);
                }
            }
        }
        return maxLen;
    }
}
```

## 合并有序链表

https://www.nowcoder.com/practice/a479a3f0c4554867b35356e0d57cf03d?tpId=117&tqId=37735&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    if(l1 == null){
        return l2;
    }
    if(l2 == null){
        return l1;
    }
    ListNode head = null, tail = null;
    while(l1 != null && l2 != null){
        if(l1.val <= l2.val){
            if(head == null){
                head = l1;
                tail = l1;
            }else{
                tail.next = l1;
                tail = tail.next;
            }
            l1 = l1.next;
        }else{
            if(head == null){
                head = l2;
                tail = l2;
            }else{
                tail.next = l2;
                tail = tail.next;
            }
            l2 = l2.next;
        }
    }

    if(l1 != null){
        tail.next = l1;
    }
    if(l2 != null){
        tail.next = l2;
    }
    return head;
}
```

## 合并K个有序链表

https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=117&tqId=37747&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC51 {
    /**
     * 基于优先级队列实现
     *
     * @param lists
     * @return
     */
    public ListNode mergeKLists_v1(ArrayList<ListNode> lists) {
        // 空数组的情况
        if (lists == null || lists.size() == 0) return null;
        if (lists.size() == 1) return lists.get(0);

        // 小根堆
        PriorityQueue<ListNode> queue = new PriorityQueue<>(new Comparator<ListNode>() {
            @Override
            public int compare(ListNode o1, ListNode o2) {
                return o1.val - o2.val;
            }
        });
        // 头节点加入小根堆，注意不是所有节点
        for (ListNode node : lists) {
            if (node != null) {
                queue.add(node);
            }
        }
        ListNode dummy = new ListNode(-1), end = dummy;
        // 小根堆的根节点是最小值
        while (!queue.isEmpty()) {
            ListNode temp = queue.poll();
            end.next = temp;
            end = end.next;
            // 当前节点所在链表的下一个节点入小根堆
            if(temp.next!=null){
                queue.add(temp.next);
            }
        }
        return dummy.next;
    }

    /**
     * 分治法
     *
     * @param lists
     * @return
     */
    public ListNode mergeKLists(ArrayList<ListNode> lists) {
        // 空数组的情况
        if (lists == null || lists.size() == 0) return null;
        if (lists.size() == 1) return lists.get(0);
        return mergeKList(lists,0, lists.size() - 1);
    }

    private ListNode mergeKList(ArrayList<ListNode> lists,int left,int right){
        if(left == right) {
            return lists.get(left);
        }else if(left > right){
            return null;
        }else{
            int mid = left + (right - left) / 2;
            ListNode list1 = mergeKList(lists,left,mid);
            ListNode list2 = mergeKList(lists,mid + 1, right);
            return mergeTwoList(list1,list2);
        }
    }

    // 合并两个链表
    private ListNode mergeTwoList(ListNode list1, ListNode list2){
        ListNode node1 = list1, node2 = list2;
        ListNode dummy = new ListNode(-1), end = dummy;
        while (node1 != null && node2 != null){
            if(node1.val <= node2.val){
                end.next = node1;
                node1 = node1.next;
            }else{
                end.next = node2;
                node2 = node2.next;
            }
            end = end.next;
        }

        if(node1 != null){
            end.next = node1;
        }

        if(node2 != null){
            end.next = node2;
        }
        return dummy.next;
    }

}
```

## **链表的奇偶重排**

https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3?tpId=117&tqId=37845&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ListNode oddEvenList (ListNode head) {
    if(head == null || head.next == null) return head;
    // 分别表示偶数链表的头结点,奇数链表的当前节点和偶数链表的当前节点
    ListNode evenHead = head.next, oddCur = head, evenCur = evenHead;
    while(evenCur != null && evenCur.next != null){
        oddCur.next = evenCur.next;
        oddCur = oddCur.next;
        evenCur.next = oddCur.next;
        evenCur = evenCur.next;
    }
    oddCur.next = evenHead;
    return head;
}
```

## 排序链表

https://leetcode-cn.com/problems/sort-list/

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode sortList(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode fast = head.next, slow = head;
        while(fast != null && fast.next != null){
            fast = fast.next.next;
            slow = slow.next;
        }
        ListNode head2 = sortList(slow.next);
        slow.next = null;
        head = sortList(head);
        ListNode dummy = new ListNode(-1);
        ListNode cur1 = head, cur2 = head2, cur = dummy;
        while(cur1 != null && cur2 != null){
            if(cur1.val <= cur2.val){
                cur.next = cur1;
                cur1 = cur1.next;
            }else{
                cur.next = cur2;
                cur2 = cur2.next;
            }
            cur = cur.next;
        }
        if(cur1 != null){
            cur.next = cur1;
        }
        if(cur2 != null){
            cur.next = cur2;
        }
        return dummy.next;
    }
}
```

## 链表相加

https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b?tpId=117&tqId=37814&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC40 {

    static class ListNode {
        int val;
        ListNode next = null;

        public ListNode(int val){
            this.val = val;
        }
    }

    public ListNode addInList (ListNode head1, ListNode head2) {
        Stack<Integer> stack1 = new Stack<>();
        Stack<Integer> stack2 = new Stack<>();
        ListNode temp1 = head1, temp2 = head2;
        // 返回的头结点
        ListNode head = null;

        while(temp1 != null || temp2 != null){
            if(temp1 != null){
                stack1.push(temp1.val);
                temp1 = temp1.next;
            }

            if(temp2 != null){
                stack2.push(temp2.val);
                temp2 = temp2.next;
            }
        }
        int carry = 0;
        while(!stack1.isEmpty() || !stack2.isEmpty() || carry != 0){
            int num1 = 0, num2 = 0;
            if(!stack1.isEmpty()){
                num1 = stack1.pop();
            }
            if(!stack2.isEmpty()){
                num2 = stack2.pop();
            }
            int sum = num1 + num2 + carry;
            carry = sum > 9 ? 1 : 0;
            ListNode node = new ListNode(sum > 9 ? sum - 10 : sum);
            if(head == null){
                head = node;
            }else{
                node.next = head;
                head = node;
            }
        }
        return head;
    }
}
```

## 删除链表的倒数第K个节点

https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6?tpId=117&tqId=37750&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ListNode removeNthFromEnd(ListNode head, int n) {
        // 头结点为空
        if(head == null) return head;

        // 添加头结点防止删除head的情况
        ListNode dummy = new ListNode();
        dummy.next = head;
        head = dummy;

        ListNode fast = head;
        // 慢指针指向待删除节点的前一个
        ListNode slow = head;
        // 快指针移动n位
        while(n > 0){
            fast = fast.next;
            if(fast == null){
                return head;
            }
            n--;
        }
        // 快指针到达尾部时慢指针指向待删除节点的前一个节点
        while(fast.next != null){
            fast = fast.next;
            slow = slow.next;
        }
        slow.next = slow.next.next;
        return dummy.next;
}
```

## 求路径

https://leetcode-cn.com/problems/unique-paths/

```java
public int uniquePaths (int m, int n) {
    if(m < 1 || n < 1) return 0;
    // dp[i][j]表示从(0,0)位置到(i,j)的走法
    int[][] dp = new int[m][n];
    dp[0][0] = 1;

    for(int i = 0;i < m;i++){
        for(int j = 0;j < n;j++){
            if(i == 0 || j == 0){
                dp[i][j] = 1;
            }else{
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
    }
    return dp[m - 1][n - 1];
}
```

## 最小编辑距离

https://www.nowcoder.com/practice/05fed41805ae4394ab6607d0d745c8e4?tpId=117&tqId=37801&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int minEditCost (String str1, String str2, int ic, int dc, int rc) {
    if(str1 == null && str2 == null) return 0;
    if(str1 == null) return str2.length() * ic;
    if(str2 == null) return str1.length() * rc;
    int len1 = str1.length(), len2 = str2.length();
    // dp[i][j]表示str1的前i个字符转化为str2的前j个字符的最小代价
    int[][] dp = new int[len1 + 1][len2 + 1];
    for(int i = 1;i <= len1;i++){
        dp[i][0] = dc * i;
    }

    for(int j = 1;j <= len2;j++){
        dp[0][j] = ic * j;
    }
    for(int i = 1;i <= len1;i++){
        char ch1 = str1.charAt(i - 1);
        for(int j = 1;j <= len2;j++){
            char ch2 = str2.charAt(j - 1);
            if(ch1 == ch2){
                dp[i][j] = dp[i - 1][j - 1];
            }else{
                int insertCost = dp[i][j - 1] + ic;
                int replaceCost = dp[i - 1][j - 1] + rc;
                int deleteCost = dp[i - 1][j] + dc;
                dp[i][j] = minCost(insertCost, minCost(replaceCost, deleteCost));
            }
        }
    }
    return dp[len1][len2];
}
private int minCost(int a,int b){
    return a < b ? a : b;
}
```

##  在排序数组中查找元素的第一个和最后一个位置

https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

```java
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] ret = {-1,-1};
        int left = 0, right = nums.length - 1, mid;
        int leftIdx = -1, rightIdx = -1;
        // 左边
        while(left <= right){
            mid = (right - left) / 2 + left;
            if(nums[mid] < target){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        leftIdx = left;
        left = 0;
        right = nums.length - 1;
        // 右边
        while(left <= right){
            mid = (right - left) / 2 + left;
            if(nums[mid] <= target){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        rightIdx = left - 1;
        if(leftIdx <= rightIdx && nums[leftIdx] == nums[rightIdx] && nums[leftIdx] == target){
            ret[0] = leftIdx;
            ret[1] = rightIdx;
        }
        return ret;
    }
}
```

## 两个数组找两个差值最小的数

```java
class Solution {
    public int smallestDifference(int[] a, int[] b) {
        Arrays.sort(a);
        Arrays.sort(b);
        long minDistance = Long.MAX_VALUE;
        for(int i = 0, j = 0;i < a.length && j < b.length;){
            long distance = a[i] < b[j] ? (long)b[j] - (long)a[i] : (long)a[i] - (long)b[j];
            if(distance < minDistance){
                minDistance = distance;
            }
            if(a[i] < b[j]){
                i++;
            }else if(a[i] == b[j]){
                return 0;
            }else{
                j++;
            }
        }
        return (int)minDistance;
    }
}
```

## 搜索旋转排序数组

https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

```java
class Solution {
    public int search(int[] nums, int target) {
        if(nums == null || nums.length == 0) return -1;
        int left = 0, right = nums.length - 1, mid;
        while(left <= right){
            mid = left + (right - left) / 2;
            if(nums[mid] == target){
                return mid;
            }
            if(nums[mid] >= nums[left]){
                if(target >= nums[left] && target < nums[mid]){
                    right = mid - 1;
                }else{
                    left = mid + 1;
                }
            }else{
                if(target > nums[mid] && target <= nums[right]){
                    left = mid + 1;
                }else{
                    right = mid - 1;
                }
            }

        }
        return -1;
    }
}
```



## 在两个长度相同的排序数组中找中位数

https://www.nowcoder.com/practice/6fbe70f3a51d44fa9395cfc49694404f?tpId=117&tqId=37808&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int findMedianinTwoSortedAray (int[] arr1, int[] arr2) {
    int left1 = 0, right1 = arr1.length - 1;
    int left2 = 0, right2 = arr2.length - 1;
    while(left1 < right1 && left2 < right2){
        int mid1 = left1 + (right1 - left1) / 2;
        int mid2 = left2 + (right2 - left2) / 2;
        // 等于1表示奇数个
        boolean flag = (((right1 - left1 + 1) & 1) == 1);
        if(arr1[mid1] == arr2[mid2]){
            return arr1[mid1];
        }else if(arr1[mid1] < arr2[mid2]){
            if(flag){
                left1 = mid1;
                right2 = mid2;
            }else{
                // 要保持两个数组去掉的部分长度相等
                left1 = mid1 + 1;
                right2 = mid2;
            }
        }else{
            if(flag){
                right1 = mid1;
                left2 = mid2;
            }else{
                left2 = mid2 + 1;
                right1 = mid1;
            }
        }
    }
    return Math.min(arr1[left1], arr2[left2]);
}
```

## 在两个长度不同的排序数组中找中位数

https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

```java
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len = nums1.length + nums2.length;
        boolean odd = true;
        if((len & 1) == 0){
            odd = false;
        }
        if(odd){
            return findMedianSortedArrays(nums1, nums2, len / 2 + 1) * 1.0;
        }else{
            return (findMedianSortedArrays(nums1, nums2, len / 2) + findMedianSortedArrays(nums1, nums2, len / 2 + 1)) / 2.0; 
        }
    }

    private int findMedianSortedArrays(int[] nums1, int[] nums2, int k){
        int idx1 = 0, idx2 = 0, cnt = 0, temp = 0;
        while(idx1 < nums1.length && idx2 < nums2.length && cnt < k){
            if(nums1[idx1] < nums2[idx2]){
                temp = nums1[idx1];
                idx1++;
                cnt++;
            }else{
                temp = nums2[idx2];
                idx2++;
                cnt++;
            }
        }
        if(idx1 < nums1.length && cnt < k){
             temp = nums1[idx1 + k - cnt - 1];
        }
        if(idx2 < nums2.length && cnt < k){
            temp = nums2[idx2 + k - cnt - 1];
        }
        return temp;
    }
}
```

## 区间合并

https://www.nowcoder.com/practice/69f4e5b7ad284a478777cb2a17fb5e6a?tpId=117&tqId=37737&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public ArrayList<Interval> merge(ArrayList<Interval> intervals) {
    if(intervals == null || intervals.size() == 0) return new ArrayList<>(0);

    ArrayList<Interval> ret = new ArrayList<>();
    int len = intervals.size();
    // 按区间左边的位置从小到大排序
    Collections.sort(intervals, new Comparator<Interval>() {
        @Override
        public int compare(Interval o1, Interval o2) {
            return o1.start - o2.start;
        }
    });
    for(int i = 1;i < len;i++){
        // 前一个区间
        Interval front = intervals.get(i - 1);
        // 后一个区间
        Interval back = intervals.get(i);
        // 后一个区间的左边界小于等于前一个区间的右边界，即有重叠
        if(back.start <= front.end){
            // 更新后一个区间的左右边界
            back.start = front.start;
            back.end = back.end > front.end ? back.end : front.end;
        }
        // 没有重叠直接将前一个区间加入到结果集
        else{
            ret.add(front);
        }
    }
    // 加入最后一个区间
    ret.add(intervals.get(intervals.size() - 1));
    return ret;
}
```

## 区间列表的交集

https://leetcode-cn.com/problems/interval-list-intersections/

```java
class Solution {
    public int[][] intervalIntersection(int[][] firstList, int[][] secondList) {
        if(firstList == null || firstList.length == 0 || secondList == null || secondList.length == 0) return new int[0][0];
        List<int[]> ret = new ArrayList();
        for(int i = 0,j = 0;i < firstList.length && j < secondList.length;){
            int left = Math.max(firstList[i][0], secondList[j][0]);
            int right = Math.min(firstList[i][1], secondList[j][1]);
            if(left <= right){
                ret.add(new int[]{left, right});
            }
            if(firstList[i][1] < secondList[j][1]){
                i++;
            }else{
                j++;
            }
        }
        return ret.toArray(new int[ret.size()][2]);
    }
}
```



## N皇后问题

https://www.nowcoder.com/practice/c76408782512486d91eea181107293b6?tpId=117&tqId=37811&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC39 {

    private int cnt = 0;

    public int Nqueen (int n) {
        if(n <= 1) return n;
        boolean[][] plate = new boolean[n][n];
        place(plate, 0);
        return cnt;
    }

    private void place(boolean[][] plate, int rowIndex){
        if(rowIndex >= plate.length){
            cnt++;
        }else{
            for(int j = 0;j < plate.length;j++){
                if(canPlace(plate, rowIndex, j)){
                    plate[rowIndex][j] = true;
                    place(plate, rowIndex + 1);
                    plate[rowIndex][j] = false;
                }
            }
        }
    }

    // 当前位置是否可以放置皇后
    private boolean canPlace(boolean[][] plate, int rowIndex, int colIndex){
        // 同一列
        for(int i = 0;i < rowIndex;i++){
            if(plate[i][colIndex]){
                return false;
            }
        }

        // 同一对角线
        for(int i = 0;i < rowIndex;i++){
            for(int j = 0;j < colIndex;j ++){
                if(plate[i][j] && rowIndex - i == colIndex - j){
                    return false;
                }
            }
        }

        // 同一斜对角线
        for(int i = 0;i < rowIndex;i++){
            for(int j = plate.length - 1;j > colIndex;j --){
                if(plate[i][j] && rowIndex - i == j - colIndex){
                    return false;
                }
            }
        }
        return true;
    }

    public static void main(String[] args) {
        NC39 nc39 = new NC39();
        System.out.println(nc39.Nqueen(8));
    }
}
```

## 子集

https://leetcode-cn.com/problems/subsets/

```java
class Solution {

    private List<List<Integer>> sets = new ArrayList<List<Integer>>();
    public List<List<Integer>> subsets(int[] nums) {
        find(nums, 0, new ArrayList());
        return sets;
    }

    private void find(int[] nums, int idx, List<Integer> oneSet){
        if(idx == nums.length){
            sets.add(new ArrayList(oneSet));
            return;
        }
        oneSet.add(nums[idx]);
        find(nums, idx + 1, oneSet);
        oneSet.remove(oneSet.size() - 1);

        find(nums, idx + 1, oneSet);
    }
}
```

## 最长无重复子数组

https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4?tpId=117&tqId=37816&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC41 {
    public int maxLength (int[] arr) {
        int maxLen = 0;
        HashMap<Integer,Integer> map = new HashMap<>();
        for(int start = 0, end = 0;end < arr.length;end++){
            // 存在
            if(map.containsKey(arr[end])){
                start = max(start,map.get(arr[end]) + 1);
            }
            map.put(arr[end], end);
            maxLen = max(maxLen, end - start + 1);
        }
        return maxLen;
    }

    private int max(int a,int b){
        return a > b ? a : b;
    }
}
```

## 最长无重复子串

https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

```java
class Solution {
    public int lengthOfLongestSubstring(String s) {
        HashMap<Character, Integer> window = new HashMap();
        int left = 0, right = 0, maxLen = 0;
        while(right < s.length()){
            char ch = s.charAt(right);
            right++;
            window.put(ch, window.getOrDefault(ch, 0) + 1);
            while(window.get(ch) > 1){
                char temp = s.charAt(left);
                left++;
                window.put(temp, window.get(temp) - 1);
            }
            if(right - left > maxLen){
                    maxLen = right - left;
            }
        }
        return maxLen;
    }
}
```

## 最小覆盖子串

https://leetcode-cn.com/problems/minimum-window-substring/

```java
class Solution {
    public String minWindow(String s, String t) {
        HashMap<Character, Integer> need = new HashMap();
        HashMap<Character, Integer> window = new HashMap();
        int left = 0, right = 0, minLen = Integer.MAX_VALUE, start = 0, valid = 0;
        for(int i = 0;i < t.length();i++){
            need.put(t.charAt(i), need.getOrDefault(t.charAt(i), 0) + 1);
        }
        while(right < s.length()){
            char ch = s.charAt(right);
            right++;
            window.put(ch, window.getOrDefault(ch, 0) + 1);
            if(window.get(ch).equals(need.get(ch))){
                valid++;
            }
            while(valid == need.size()){
                if(right - left < minLen){
                    minLen = right - left;
                    start = left;
                }
                char leftCh = s.charAt(left);
                left++;
                if(window.get(leftCh).equals(need.get(leftCh))){
                    valid--;
                }
                window.put(leftCh, window.get(leftCh) - 1);
            }
        }
        return minLen == Integer.MAX_VALUE ? "" : s.substring(start, start + minLen);
    }
}
```

## 字母异位词

https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

```java
class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        ArrayList<Integer> indexs = new ArrayList();
        HashMap<Character, Integer> window = new HashMap();
        HashMap<Character, Integer> need = new HashMap();
        for(int i = 0;i < p.length();i++){
            char ch = p.charAt(i);
            need.put(ch, need.getOrDefault(ch, 0) + 1);
        }
        int left = 0, right = 0, valid = 0;
        while(right < s.length()){
            char ch = s.charAt(right);
            right++;
            window.put(ch, window.getOrDefault(ch , 0) + 1);
            if(window.get(ch).equals(need.get(ch))){
                valid++;
            }
            while(right - left >= p.length()){
                if(valid == need.size()){
                    indexs.add(left);
                }
                ch = s.charAt(left);
                left++;
                if(window.get(ch).equals(need.get(ch))){
                    valid--;
                }
                window.put(ch, window.get(ch) - 1);
            }
        }
        return indexs;
    }
}
```





## 无重复数字的全排列

https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863?tpId=117&tqId=37739&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC43 {
    public ArrayList<ArrayList<Integer>> permute(int[] num) {
        if(num == null || num.length == 0) return new ArrayList<>(0);
        // 升序排序
        Arrays.sort(num);
        // 记录num对应位置的数是否已经使用
        boolean[] used = new boolean[num.length];
        ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
        helper(num,used,ret,new ArrayList<>());
        return ret;
    }

    private void helper(int[] num, boolean[] used, ArrayList<ArrayList<Integer>> ret, ArrayList<Integer> sequence){
        // 形成一个序列
        if(sequence.size() == num.length){
            ret.add(new ArrayList<>(sequence));
        }else{
            for(int i = 0;i < num.length;i++){
                if(used[i] == false){
                    sequence.add(num[i]);
                    // 标记为已使用
                    used[i] = true;
                    // 处理序列的下一个数
                    helper(num,used,ret,sequence);
                    // 还原
                    used[i] = false;
                    sequence.remove(sequence.size() - 1);
                }
            }
        }
    }
}
```

## 有重复数字的全排列

```java
class Solution {

    public List<List<Integer>> permuteUnique(int[] nums) {
        boolean visited[] = new boolean[nums.length];
        List<List<Integer>> ret = new ArrayList();
        ArrayList<Integer> temp = new ArrayList();
        Arrays.sort(nums);
        helper(ret, temp, nums, 0, visited);
        return ret;
    }
    private void helper(List<List<Integer>> all, List<Integer> one, int[] nums, int idx, boolean visited[]){
        if(idx == nums.length){
            all.add(new ArrayList(one));
            return;
        }
        for(int i = 0;i < nums.length;i++){
						// 去重
            if(visited[i] || (i > 0 && nums[i] == nums[i - 1] && !visited[i - 1])){
                continue;
            }
            visited[i] = true;
            one.add(nums[i]);
            helper(all, one, nums, idx + 1, visited);
            visited[i] = false;
            one.remove(one.size() - 1);
        }
    }
}
```

## 解数独

https://leetcode-cn.com/problems/sudoku-solver/

```java
class Solution {
    // 是否已经有了有效地匹配
    private boolean valid = false;
    // column[i][j] 表示第i列是否出现过数字j + 1
    private boolean[][] column = new boolean[9][9];
    // block[i][j][k] 表示第(i,j)个小正方形内是否出现过数字k + 1
    private boolean[][][] block = new boolean[3][3][9];
    // row[i][j] 表示第i行是否出现过数字j + 1
    private boolean[][] row = new boolean[9][9];
    // 所有待填充的位置
    private ArrayList<int[]> todo = new ArrayList<>();
    public void solveSudoku(char[][] board) {
        for(int i = 0;i < 9;i++){
            for(int j = 0;j < 9;j++){
                if(board[i][j] == '.'){
                    todo.add(new int[]{i, j});
                }else{
                    // 记录已经有数的情况
                    int digit = board[i][j] - '0';
                    row[i][digit - 1] = true;
                    column[j][digit - 1] = true;
                    block[i / 3][j / 3][digit - 1] = true;
                }
            }
        }
        dfs(board, 0);
    }

    private void dfs(char[][] board, int pos){
        // 所有位置都已经填充完毕
        if(pos == todo.size()){
            valid = true;
            return;
        }

        int i = todo.get(pos)[0];
        int j = todo.get(pos)[1];
        for(int digit = 1;digit <= 9;digit ++){
            // 该位置可以填入该数
            if(!row[i][digit - 1] && !column[j][digit - 1] && !block[i / 3][j / 3][digit - 1] && !valid){
                row[i][digit - 1] = column[j][digit - 1] = block[i / 3][j / 3][digit - 1] = true;
                board[i][j] = (char)(digit + '0');
                dfs(board, pos + 1);
                row[i][digit - 1] = column[j][digit - 1] = block[i / 3][j / 3][digit - 1] = false;
            }
        }
    }

}
```

## 下一个排列

https://leetcode-cn.com/problems/next-permutation/

```java
class Solution {
    public void nextPermutation(int[] nums) {
        int i;
        for(i = nums.length - 2;i >= 0 && nums[i] >= nums[i + 1];i--){}
        int j;
        if(i >= 0) {
            for(j = nums.length - 1;j > i && nums[j] <= nums[i];j--){}
            nums[i] = nums[i] ^ nums[j];
            nums[j] = nums[i] ^ nums[j];
            nums[i] = nums[i] ^ nums[j];
        }
        for(i = i + 1,j = nums.length - 1;i < j;i++,j--){
            nums[i] = nums[i] ^ nums[j];
            nums[j] = nums[i] ^ nums[j];
            nums[i] = nums[i] ^ nums[j];
        }
    }
}
```

## 最大交换

https://leetcode-cn.com/problems/maximum-swap/

```java
class Solution {
    public int maximumSwap(int num) {
        char[] nums = String.valueOf(num).toCharArray();
        int[] lastPos = new int[10];
        // 记录每个数字最后一次出现的位置
        for(int i = 0;i < nums.length;i++){
            lastPos[nums[i] - '0'] = i;
        }
        // 从左往右遍历
        for(int i = 0;i < nums.length;i++){
            // 找到比位置i的数大的数中最大的数与该位置交换
            for(int j = 9;j > nums[i] - '0';j--){
                if(lastPos[j] > i){
                    swap(nums, i ,lastPos[j]);
                    return Integer.valueOf(new String(nums));
                }
            }
        }
        return num;
    }
}
```



## 复原IP地址

https://leetcode-cn.com/problems/restore-ip-addresses/

```java
class Solution {
    private List<String> res = new ArrayList();

    public List<String> restoreIpAddresses(String s) {
        helper(s,-1,0, new ArrayList<String>());
        return res;
    }

    private void helper(String str, int lastDotIdx, int dotCnt, ArrayList<String> segments){
        if(dotCnt > 4) return;
        if(lastDotIdx == str.length() - 1){
            if(dotCnt == 4){
                res.add(String.join(".", segments));
            }
            return;
        }
        for(int i = lastDotIdx + 1;i < str.length() && i <= lastDotIdx + 3;i++){
            if(judge(str, lastDotIdx + 1, i)){
                String seg = str.substring(lastDotIdx + 1, i + 1);
                segments.add(seg);
                helper(str, i, dotCnt + 1, segments);
                segments.remove(segments.size() - 1);
            }
        }
    }

    private boolean judge(String str, int start, int end){
        // 每一个部分长度为1~3
        if(end - start + 1 >= 4){
            return false;
        }
        // 每一段可以只有一个0但是长度不为1时不能以0开头
        if(end - start > 0 && str.charAt(start) == '0') return false;
        // 计算对应的数
        int num = 0;
        for(int i = start;i <= end;i++){
            num = num * 10 + (str.charAt(i) - '0');
        }
        // 最大是255
        if(num > 255) return false;
        return true;
    }
}
```



## 正则表达式匹配

https://leetcode-cn.com/problems/regular-expression-matching/

```java
class Solution {
    public boolean isMatch(String s, String p) {
        HashMap<Integer, Boolean> memo = new HashMap<>();
        return match(s,0,p,0, memo);
    }

    private boolean match(String s, int i, String p, int j, HashMap<Integer, Boolean> memo){
        if(j == p.length()){
            return i == s.length();
        }
        if(i > s.length()) return false;
        Integer key = i << p.length() ^ j;
        if(memo.containsKey(key)){
            return memo.get(key);
        }
        boolean firstMatch = i < s.length() && (s.charAt(i) == p.charAt(j) || p.charAt(j) == '.');
        if(j + 1 < p.length() && p.charAt(j + 1) == '*'){
            memo.put(key, match(s, i, p, j + 2, memo) || (firstMatch && match(s, i + 1, p, j, memo)));
        }else{
            memo.put(key, firstMatch && match(s, i + 1, p, j + 1, memo));
        }
        return memo.get(key);
    }
    
}
```

## 通配符匹配

https://www.nowcoder.com/practice/e96f1a44d4e44d9ab6289ee080099322?tpId=117&tqId=37741&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public boolean isMatch(String s, String p) {
    int sLen = s.length(), pLen = p.length();
    // dp[i][j]表示s的前i个字符和p的前j个字符是否匹配
    boolean[][] dp = new boolean[sLen + 1][pLen + 1];
    dp[0][0] = true;
    // 匹配串为空的情况
    for(int i = 1;i <= pLen;i++){
        if(p.charAt(i - 1) == '*'){
            dp[0][i] = dp[0][i - 1];
        }
    }
    for(int j = 1; j <= pLen;j++){
        for(int i = 1;i <= sLen;i++){
            // 当前位置模式串和匹配串字符相等或者模式串当前字符是'?', 是否匹配取决于前一个位置
            if(s.charAt(i - 1) == p.charAt(j - 1) || p.charAt(j - 1) == '?'){
                dp[i][j] = dp[i - 1][j - 1];
            }
            // 当前位置不相等且模式串当前字符是'*',取决于匹配串是够存在子串能与模式串*之前的子串匹配
            else if(p.charAt(j - 1) == '*'){
                // dp[i][j - 1]表示*匹配空串
                //  dp[i - 1][j - 1]表示*匹配一个字符
                // dp[i - 1][j]表示*匹配多个字符
                dp[i][j] = dp[i - 1][j] || dp[i][j - 1] || dp[i - 1][j - 1];
            }
        }
    }
    return dp[sLen][pLen];
}

public class Solution {
    public boolean isMatch(String s, String p) {
        HashMap<Integer, Boolean> memo = new HashMap<>();
        return match(s,0,p,0, memo);
    }

    private boolean match(String s, int i, String p, int j, HashMap<Integer, Boolean> memo){
        if(j == p.length()){
            return i == s.length();
        }
        if(i > s.length()) return false;
        Integer key = i << p.length() ^ j;
        if(memo.containsKey(key)){
            return memo.get(key);
        }
        boolean firstMatch = i < s.length() && (s.charAt(i) == p.charAt(j) || p.charAt(j) == '?');
        if(firstMatch){
            return match(s, i + 1, p, j + 1, memo);
        }
        if(j < p.length() && p.charAt(j) == '*'){
            memo.put(key, match(s, i, p, j + 1, memo) || match(s, i + 1, p, j, memo));
        }
        return memo.getOrDefault(key, false);
    }
}
```

## 加起来和为目标值的组合

https://www.nowcoder.com/practice/75e6cd5b85ab41c6a7c43359a74e869a?tpId=117&tqId=37742&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC46 {

    public ArrayList<ArrayList<Integer>> combinationSum2(int[] num, int target) {
        if(num == null || num.length == 0) return new ArrayList<>(0);
        Arrays.sort(num);
        boolean[] used = new boolean[num.length];
        ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
        helper(ret, new ArrayList<Integer>(),num, 0, target);
        return ret;
    }

    private void helper(ArrayList<ArrayList<Integer>> ret, ArrayList<Integer> sequence,int[] num,int startIndex, int target){
        // 找到一组符合的组合
        if(target == 0){
            ret.add(new ArrayList<>(sequence));
            return;
        }
        // 这里使用一个索引值startIndex表示前面的已访问，不能使用标记数组，因为组合内部是有序的
        for(int i = startIndex;i < num.length;i++){
            // 去重
            if(i > startIndex && num[i] == num[i - 1]) continue;
            // 剪枝
            if(target >= num[i]){
                sequence.add(num[i]);
                helper(ret,sequence,num,i + 1,target - num[i]);
                sequence.remove(sequence.size() - 1);
            }
        }
    }

    public static void main(String[] args) {
        NC46 nc46 = new NC46();
        int[] arr = {100,10,20,70,60,10,50};
        ArrayList<ArrayList<Integer>> ret = nc46.combinationSum2(arr,80);
        System.out.println();
    }
}
```

## 在转动过的数组找目标值

https://www.nowcoder.com/practice/7cd13986c79d4d3a8d928d490db5d707?tpId=117&tqId=37744&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int search (int[] A, int target) {
    if(A == null || A.length == 0) return -1;

    // 二分查找
    for(int left = 0, right = A.length - 1;left <= right;){
        int mid = left + (right - left) / 2;
        // 找到目标元素
        if(A[mid] == target) return mid;
        // 左侧有序
        if(A[mid] >= A[left]){
            // 在范围内
            if(A[left] <= target && A[mid] > target){
                right = mid - 1;
            }
            // 不在范围内
            else{
                left = mid + 1;
            }
        }
        // 右侧有序
        else{
            if(A[mid] < target && A[right] >= target){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
    }
    return -1;
}
```

## 矩阵元素查找

https://www.nowcoder.com/practice/3afe6fabdb2c46ed98f06cfd9a20f2ce?tpId=117&tqId=37788&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int[] findElement(int[][] mat, int n, int m, int x) {
    if(mat == null || mat.length == 0 || mat[0].length == 0) return null;
    for(int i = mat.length - 1;i >= 0;){
        for(int j = 0;j < mat[0].length;){
            if(mat[i][j] == x){
                return new int[]{i, j};
            }else if(mat[i][j] < x){
                j++;
            }else{
                i--;
            }
        }
    }
    return null;
}
```

## **寻找第K大**

https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf?tpId=117&tqId=37791&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
int find(int[] a,int i, int j,int k){
    int temp = a[i];
    int left = i,right = j;
    while(left < right){
        while(left < right && a[right] <= temp) right--;
        a[left] = a[right];
        while(left < right && a[left] >= temp) left++;
        a[right] = a[left];
    }
    a[left] = temp;

    if(left == k - 1){
        return a[left];
    }else if(left < k - 1){
        return find(a,left + 1,j, k);
    }else{
        return find(a,i,left - 1, k);
    }
}
```

## 二叉搜索树的第K小元素

https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

```java
public class Solution {
    public int kthSmallest(TreeNode root, int k) {
        if(root == null) return 0;
        int cnt = 0;
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        while(!stack.isEmpty() || cur != null){
            if(cur != null){
                stack.push(cur);
                cur = cur.left;
                continue;
            }
            cur = stack.pop();
            k--;
            if(k == 0){
                return cur.val;
            }
            cur = cur.right;
        }
        return -1;
    }
}
```

## **最大数**

https://www.nowcoder.com/practice/fc897457408f4bbe9d3f87588f497729?tpId=117&tqId=37835&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public String solve (int[] nums) {
    // 转化为字符串
    String[] strs = new String[nums.length];
    for(int i = 0;i < nums.length;i++){
        strs[i] = String.valueOf(nums[i]);
    }
    // 排序
    // 顺序是: 相邻两个数组合结果更大时的顺序
    Arrays.sort(strs, new Comparator<String>() {
        @Override
        public int compare(String o1, String o2) {
            String temp1 = o1.concat(o2);
            String temp2 = o2.concat(o1);
            if(temp1.compareTo(temp2) > 0){
                return -1;
            }else{
                return 1;
            }
        }
    });
    // 全为0的情况
    if(strs[0].charAt(0) == '0') return "0";
    StringBuilder sb = new StringBuilder();
    for (String str : strs) {
        sb.append(str);
    }
    return sb.toString();
}
```

## 缺失的第一个正数

https://leetcode-cn.com/problems/first-missing-positive/

```java
class Solution {
    public int firstMissingPositive(int[] nums) {
        for(int i = 0;i < nums.length;i++){
            while(nums[i] > 0 && nums[i] <= nums.length && nums[i] != nums[nums[i] - 1]){
                swap(nums, i, nums[i] - 1);
            }
        }
        for(int i = 0;i < nums.length;i++){
            if(nums[i] != i + 1){
                return i + 1;
            }
        }
        return nums.length + 1;
    }

    private void swap(int[] nums, int i, int j){
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
```

## **二分查找-II**

https://www.nowcoder.com/practice/4f470d1d3b734f8aaf2afb014185b395?tpId=117&tqId=37829&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey  

```java
public int search (int[] nums, int target) {
    if(nums == null || nums.length == 0) return -1;
    int left = 0, right = nums.length - 1;
    while(left < right){
        int mid = left + (right - left) / 2;
        if(nums[mid] == target){
            right = mid;
        }else if(nums[mid] < target){
            left = mid + 1;
        }else{
            right = mid - 1;
        }
    }
    return nums[left] == target ? left : -1;
}
```

## 寻找峰值

https://leetcode-cn.com/problems/find-peak-element/

```java
class Solution {
    public int findPeakElement(int[] nums) {
        return binaryFind(nums, 0, nums.length - 1);
    }

    private int binaryFind(int[] nums, int left, int right){
        if(left == right){
            return nums[left];
        }
        int mid = left + (right - left) / 2;
        if(nums[mid] <= nums[mid + 1]){
            return binaryFind(nums, mid + 1, right);
        }else{
            return binaryFind(nums, left, mid);
        }
    }
}
```

## 找出一个未排序数组中第K大的数

https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

```java
class Solution {
    public int findKthLargest(int[] nums, int k) {
        return helper(nums, 0, nums.length - 1, k);
    }

    private int helper(int[] nums, int start, int end, int k){
        int left = start, right = end;
        int temp = nums[start];
        while(left < right){
            while(left < right && nums[right] <= temp) right--;
            nums[left] = nums[right];
            while(left < right && nums[left] >= temp) left++;
            nums[right] = nums[left];
        }
        nums[left] = temp;
        if(left == k - 1){
            return nums[left];
        }else if(left < k - 1){
            return helper(nums, left + 1, end, k);
        }else{
            return helper(nums, start, left - 1, k);
        }
    }
}
```

## 数组中的逆序对

https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

```java
class Solution {
    public int reversePairs(int[] nums) {
        if(nums.length <= 1) return 0;
        return getReversePairs(nums, 0, nums.length - 1, new int[nums.length]);
    }

    private int getReversePairs(int[] nums, int left, int right, int[] temp){
        // 只有一个元素时必不可能是逆序对
        if(left == right) return 0;
        // 求mid左右两个子数组的逆序数
        int mid = left + (right - left) / 2;
        int leftPairs = getReversePairs(nums, left, mid, temp);
        int rightPairs = getReversePairs(nums, mid + 1, right, temp);
        // 如果两个子数组是递增的，那必不可能有跨两个子数组的逆序对
        if(nums[mid] <= nums[mid + 1]) return leftPairs + rightPairs;
        // 求跨两个子数组的逆序对
        int spanPairs = getSpanPairs(nums, left ,mid, right, temp);
        return leftPairs + rightPairs + spanPairs;
    }

    private int getSpanPairs(int[] nums, int left, int mid, int right, int[] temp) {
        for(int i = left;i <= right;i++){
            temp[i] = nums[i];
        }
        int i = left, j = mid + 1, cnt = 0, k = left;
        // 两种情况:
        // 1. 左边数组的当前元素比右边数组的当前元素小或等于, 不构成逆序
        // 2. 左边数组的当前元素比右边数组的当前元素大，那么左边数组当前位置及之后所有元素与右边数组的当前元素构成逆序对, 更新逆序对数
        while(i <= mid && j <= right){
            if(temp[i] <= temp[j]){
                nums[k++] = temp[i];
                i++;
            }else{
                nums[k++] = temp[j];
                cnt += (mid - i + 1);
                j++;
            }
        }
        while(i <= mid){
            nums[k++] = temp[i];
            i++;
        }
        while(j <= right){
            nums[k++] = temp[j];
            j++;
        }
        return cnt;
    }
}
```

## **字符串出现次数的TopK问题**

https://www.nowcoder.com/practice/fd711bdfa0e840b381d7e1b82183b3ee?tpId=117&tqId=37809&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
class Pair{
    String key;
    Integer cnt;

    public Pair(String key, Integer cnt) {
        this.key = key;
        this.cnt = cnt;
    }
}
public String[][] topKstrings (String[] strings, int k) {
    if(strings == null || strings.length == 0 || k == 0) return new String[0][0];

    // 记录每个key出现的频次
    HashMap<String,Integer> map = new HashMap<>();
    for(String str : strings){
        map.put(str, map.getOrDefault(str, 0 ) + 1);
    }

    // 构建大顶堆
    PriorityQueue<Pair> bigHeap = new PriorityQueue<>(new Comparator<Pair>() {
        @Override
        public int compare(Pair o1, Pair o2) {
            if(o1.cnt < o2.cnt){
                return 1;
            }else if(o1.cnt.compareTo(o2.cnt) == 0 && o1.key.compareTo(o2.key) > 0){
                return 1;
            }else{
                return -1;
            }
        }
    });

    // 入堆
    for(String key : map.keySet()){
        bigHeap.add(new Pair(key,map.get(key)));
    }

    // 找出前k个
    int cnt = 0;
    String[][] ret = new String[k][2];
    while(cnt < k && !bigHeap.isEmpty()){
        Pair pair = bigHeap.poll();
        ret[cnt][0] = pair.key;
        ret[cnt++][1] = String.valueOf(pair.cnt);
    }
    return ret;
}
```



## 数组中和为0的三元组

https://www.nowcoder.com/practice/345e2ed5f81d4017bbb8cc6055b0b711?tpId=117&tqId=37751&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC54 {
    public ArrayList<ArrayList<Integer>> threeSum(int[] num) {
        // 空数组的情况
        if(num == null || num.length < 3) return new ArrayList<>(0);
        // 排序
        Arrays.sort(num);
        // 第一个数大于0则三个数的何必不可能等于0
        if(num[0] > 0) return new ArrayList<>(0);
        // 最后的返回结果
        ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
        // 第一个数只能不能取最后两个
        for(int i = 0;i < num.length - 2;i++){
            // 第二个数,第三个数,后两个数之和
            int j = i + 1, k = num.length - 1,target = -num[i];
            // 第二个数在第三个数前面
            while(j < k){
                // 三数之和等于0
                if(num[j] + num[k] == -num[i]){
                    ArrayList<Integer> temp = new ArrayList<>(3);
                    temp.add(num[i]);
                    temp.add(num[j]);
                    temp.add(num[k]);
                    ret.add(temp);
                    // 第二个数重复，跳过,j + 1 < k而不是j < k是因为只有两个数还有数时才需要跳过
                    while(j + 1 < k && num[j] == num[j + 1]) j++;
                    // 第三个数重复，跳过
                    while(j + 1 < k && num[k] == num[k - 1]) k--;
                    j++;
                    k--;
                }else if(num[j] + num[k] < -num[i]){
                    j++;
                }else{
                    k--;
                }
            }
            // 第一个数重复，跳过
            while (i + 1 < num.length - 2 && num[i] == num[i+1]) i++;
        }
        return ret;
    }
}
```

## 完全平方数

https://leetcode-cn.com/problems/perfect-squares/

```java
class Solution {
    public int numSquares(int n) {
        int[] dp = new int[n + 1];
        for(int i = 1;i <= n;i++){
            dp[i] = i;
            for(int j = 1;j * j <= i;j++){
                dp[i] = Math.min(dp[i], dp[i - j * j] + 1);
            }
        }
        return dp[n];
    }
}
```



## 矩阵的最小路径和

https://www.nowcoder.com/practice/7d21b6be4c6b429bb92d219341c4f8bb?tpId=117&tqId=37823&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int minPathSum (int[][] matrix) {
    if(matrix == null || matrix.length == 0 || matrix[0].length == 0) return 0;

    int m = matrix.length;
    int n = matrix[0].length;
    // dp[i][j]表示从matrix[0][0]走到matrix[i][j]位置的最小路径和
    int[][] dp = new int[m][n];
    dp[0][0] = matrix[0][0];
    // 第一行元素只能从左往右
    for(int j = 1;j < n;j ++){
        dp[0][j] = dp[0][j - 1] + matrix[0][j];
    }

    // 第一列
    for(int i = 1;i < m;i ++){
        dp[i][0] = dp[i - 1][0] + matrix[i][0];
    }

    for(int i = 1;i < m;i ++){
        for(int j = 1;j < n;j++){
            dp[i][j] = min(dp[i - 1][j],dp[i][j - 1]) + matrix[i][j];
        }
    }
    return dp[m - 1][n - 1];
}
```

## **矩阵最长递增路径**

https://www.nowcoder.com/practice/7a71a88cdf294ce6bdf54c899be967a2?tpId=117&tqId=37850&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
// memo[i][j]表示从矩阵(i,j)位置开始的最长长度
private int[][] memo;

public int solve (int[][] matrix) {
    if(matrix == null || matrix.length == 0 || matrix[0].length == 0) return 0;
    // 行数和列数
    int rows = matrix.length, cols = matrix[0].length;
    // 最长长度
    int maxValue = 0;
    // 初始化备忘录
    memo = new int[rows][cols];
    // 遍历每一个节点,计算从该位置开始的最长长度
    for(int i = 0;i < rows;i++){
        for(int j = 0;j < cols;j++){
            maxValue = Math.max(walk(matrix, i, j, -1), maxValue);
        }
    }
    return maxValue;
}

private int walk(int[][] matrix,int i, int j,int lastValue){
    int maxLen = 0;
    if(lastValue >= matrix[i][j]) return maxLen;
    // 当前位置开始的最长长度已知
    if(memo[i][j] != 0){
        return memo[i][j];
    }
    // 上
    if(i > 0){
        maxLen = Math.max(maxLen, walk(matrix,i - 1,j, matrix[i][j]));
    }
    // 左
    if(j > 0){
        maxLen = Math.max(maxLen, walk(matrix,i,j - 1, matrix[i][j]));
    }
    // 下
    if(i < matrix.length - 1){
        maxLen = Math.max(maxLen, walk(matrix,i + 1, j, matrix[i][j]));
    }
    // 右
    if(j < matrix[0].length - 1){
        maxLen = Math.max(maxLen, walk(matrix,i, j + 1, matrix[i][j]));
    }
    // 更新当前位置开始的最长长度
    memo[i][j] = maxLen + 1;
    return maxLen + 1;
}
```

## **判断一棵二叉树是否为搜索二叉树和完全二叉树**

https://www.nowcoder.com/practice/f31fc6d3caf24e7f8b4deb5cd9b5fa97?tpId=117&tqId=37822&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC60 {

    public boolean[] judgeIt(TreeNode root) {
        if (root == null || (root.left == null && root.right == null)) return new boolean[]{true, true};
        boolean bst = isBST(root);
        boolean completed = isCompletedTree(root);
        return new boolean[]{bst, completed};
    }

    // 判断是否是搜索二叉树
    private boolean isBST(TreeNode root) {
        ArrayList<Integer> sequence = new ArrayList<>();
        Stack<TreeNode> stack = new Stack<>();
        TreeNode temp = root;
        while (temp != null || !stack.isEmpty()) {
            if (temp != null) {
                stack.push(temp);
                temp = temp.left;
            } else {
                temp = stack.pop();
                sequence.add(temp.val);
                temp = temp.right;
            }
        }
        // 中序遍历序列是否为升序序列
        int ssize = sequence.size();
        for (int i = 1; i < ssize; i++) {
            if (sequence.get(i) < sequence.get(i - 1)) {
                return false;
            }
        }
        return true;
    }

    // 判断是否是完全二叉树
    private boolean isCompletedTree(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        // 当前节点不空就将其全部子节点入队列
        while (queue.peek() != null) {
            TreeNode temp = queue.poll();
            queue.add(temp.left);
            queue.add(temp.right);
        }
        // 如果空节点后面还有非空节点则不是完全二叉树
        while(!queue.isEmpty() && queue.peek() == null) {
            queue.poll();
        }
        return queue.isEmpty();
    }
}
```

## 根据前序和中序构建二叉树

https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return helper(preorder, 0, inorder, 0, inorder.length - 1);
    }

    public TreeNode helper(int[] preorder, int preStart, int[] inorder, int inStart,int inEnd){
        if(inStart <= inEnd){
            TreeNode root = new TreeNode(preorder[preStart]);
            int i;
            for(i = inStart;i <= inEnd;i++){
                if(inorder[i] == preorder[preStart]){
                    break;
                }
            }
            root.left = helper(preorder, preStart + 1,inorder, inStart, i - 1);
            root.right = helper(preorder,preStart + i - inStart + 1 , inorder, i + 1, inEnd);
            return root;
        }else{
            return null;
        }
    }
}
```



## **树的直径**

```java
public class NC99 {

    class Interval {
        int start;
        int end;
    }

    // 图的边节点
    class Edge{
        int end;
        int value;

    }

    public int solve (int n, Interval[] Tree_edge, int[] Edge_value) {
        if(n == 0 || Tree_edge == null || Edge_value == null || Tree_edge.length != Edge_value.length) return 0;
        // 邻接表法表示图
        Map<Integer, List<Edge>> graph = new HashMap<>();
        for(int i = 0;i < Tree_edge.length;i++){
            Interval interval = Tree_edge[i];
            // 起点、终点和权重
            int start  = interval.start;
            int end = interval.end;
            int value = Edge_value[i];
            // 构建边节点
            Edge edge1 = new Edge();
            edge1.end = end;
            edge1.value = value;
            if(!graph.containsKey(start)){
                List<Edge> edges = new ArrayList<>();
                graph.put(start,edges);
            }
            graph.get(start).add(edge1);

            // 双向的边节点
            Edge edge2 = new Edge();
            edge2.end = start;
            edge2.value = value;
            if(!graph.containsKey(end)){
                List<Edge> edges = new ArrayList<>();
                graph.put(end,edges);
            }
            graph.get(end).add(edge2);
        }

        // remote[0]表示从节点0开始的最长路径长度，remote[1]表示最长路径的端点
        int[] remote = new int[]{0,0};

        dfs(graph,0,-1,0,remote);

        int[] res = new int[]{0,0};

        dfs(graph, remote[1], -1, 0, res);

        return res[0];
    }

    private void dfs(Map<Integer, List<Edge>> graph,int from, int prev,int path, int[] remote){
        List<Edge> edges = graph.get(from);
        for(Edge edge : edges){
            if(edge.end != prev){
                path += edge.value;
                if(path > remote[0]){
                    remote[0] = path;
                    remote[1] = edge.end;
                }
                dfs(graph,edge.end, from,path, remote);
                path -= edge.value;
            }
        }
    }
}
```

## 二叉树的最大路径和

https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int maxSum = Integer.MIN_VALUE;
    public int maxPathSum(TreeNode root) {
        if(root == null) return 0;
        helper(root);
        return maxSum;
    }
    private int helper(TreeNode root){
        if(root == null) return 0;
        int leftMaxSum = Math.max(helper(root.left), 0);
        int rightMaxSum = Math.max(helper(root.right), 0);
        int tempMaxSum = leftMaxSum + rightMaxSum + root.val;
        maxSum = Math.max(maxSum, tempMaxSum);
        return root.val + Math.max(leftMaxSum, rightMaxSum);
    }
}
```

## 二叉树的直径

https://leetcode-cn.com/problems/diameter-of-binary-tree/

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int maxNodeCnt = 1;
    public int diameterOfBinaryTree(TreeNode root) {
        helper(root);
        return maxNodeCnt - 1;
    }
    private int helper(TreeNode root){
        if(root == null) return 0;
        int left = helper(root.left);
        int right = helper(root.right);
        maxNodeCnt = Math.max(maxNodeCnt, left + right + 1);
        return Math.max(left, right) + 1;
    }
}
```



## **在二叉树中找到两个节点的最近公共祖先**

https://www.nowcoder.com/practice/e0cc33a83afe4530bcec46eba3325116?tpId=117&tqId=37826&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int lowestCommonAncestor(TreeNode root, int o1, int o2) {
    TreeNode commonAncestor = commonAncestor(root, o1, o2);
    return commonAncestor != null ? commonAncestor.val : -1;
}

private TreeNode commonAncestor(TreeNode root, int o1, int o2){
    if(root == null || root.val == o1 || root.val == o2) return root;
    TreeNode left = commonAncestor(root.left,o1,o2);
    TreeNode right = commonAncestor(root.right,o1,o2);
    if(left != null && right != null) return root;
    return left != null ? left : right;
}

static class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;
}
```

## 满二叉树的公共祖先

```java
public int binarySearch(int one,int two,int three,int left,int right)

    {

        int middle=(left+right)/2;

        if((one-middle)*(two-middle)<=0||

                (one-middle)*(three-middle)<=0||

                (two-middle)*(three-middle)<=0)

        {

            //说明至少有一个结点在不同的一侧

            //此时的middle结点就是最小公共祖先

            return middle;

        }else if(one>middle){

            //确定二分方向，用two>middle或者three>middle也可以，

            //反正这三个结点都在middle的某一侧

            //向右侧二分

            return binarySearch(one,two,three,middle+1,right);

        }else{

            //向左侧二分

            return binarySearch(one,two,three,left,middle-1);

        }
    }
```

## 最长重复子数组

https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

```java
class Solution {
    // 动态规划方式
    public int findLength_v1(int[] nums1, int[] nums2) {
        int[][] dp = new int[nums1.length][nums2.length];
        int maxLength = 0;
        for(int i = nums1.length - 1;i >= 0;i--){
            for(int j = nums2.length - 1;j >= 0;j--){
                if(nums1[i] == nums2[j]){
                    if(i == nums1.length - 1 || j == nums2.length - 1){
                        dp[i][j] = 1;
                    }else{
                        dp[i][j] = nums1[i] == nums2[j] ? dp[i + 1][j + 1] + 1 : 0;
                    }
                }
                maxLength = Math.max(maxLength, dp[i][j]);
            }
        }
        return maxLength;
    }

    // 滑动窗口方式
    public int findLength(int[] nums1, int[] nums2) {
        int maxLength = 0;
        for(int i = 0;i < nums1.length;i++){
            maxLength = Math.max(getMaxCommonLength(nums1, i, nums2, 0), maxLength);
        }
        for(int j = 0;j < nums2.length;j++){
            maxLength = Math.max(getMaxCommonLength(nums1, 0, nums2, j), maxLength);
        }

        return maxLength;
    }
    private int getMaxCommonLength(int[] nums1, int start1, int[] nums2, int start2){
        int cnt = 0, max = 0;
        while(start1 < nums1.length && start2 < nums2.length){
            if(nums1[start1] == nums2[start2]){
                cnt++;
            }else{
                cnt = 0;
            }
            max = Math.max(cnt, max);
            start1++;
            start2++;
        }
        return max;
    }
}
```

## 滑动窗口最大值

https://leetcode-cn.com/problems/sliding-window-maximum/

```java
class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int[] ret = new int[nums.length - k + 1];
        int pos = 0;
        Deque<Integer> deque = new LinkedList();
        for(int i = 0;i < nums.length;i++){
            while(!deque.isEmpty() && nums[deque.peekFirst()] <= nums[i]){
                deque.pollFirst();
            }
            while(!deque.isEmpty() && deque.peekLast() + k - 1 < i){
                deque.pollLast();
            }
            deque.offerFirst(i);
            if(i + 1 >= k){
                ret[pos++] = nums[deque.peekLast()];
            }
        }
        return ret;
    }
}
```

## 每日温度

https://leetcode-cn.com/problems/daily-temperatures/

```java
class Solution {
    public int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> stack = new Stack<>();
        int[] res = new int[temperatures.length];
        for(int i = 0;i < temperatures.length;i++) {
            // 当前温度比之前的温度高，计算那一天后需要多少天才会遇到更高的温度
            while(!stack.isEmpty() && temperatures[stack.peek()] < temperatures[i]){
                int index = stack.peek();
                res[index] = i - index;
                stack.pop();
            }
            stack.push(i);
        }
        return res;
    }
}
```

## 长度最小的子数组

https://leetcode-cn.com/problems/minimum-size-subarray-sum/

```java
class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int minLen = Integer.MAX_VALUE;
        int left = 0, right = 0, sum = 0;
        while(right < nums.length){
            sum += nums[right++];
            while(sum >= target){
                minLen = Math.min(minLen, right - left);
                sum -= nums[left];
                left++;
            }
        }
        return minLen == Integer.MAX_VALUE ? 0 : minLen;
    }
}
```

## 和为K的子数组

https://leetcode-cn.com/problems/subarray-sum-equals-k/

```java
class Solution {
    public int subarraySum(int[] nums, int k) {
        int sum = 0;
        int cnt = 0;
        HashMap<Integer, Integer> map = new HashMap<>();
        map.put(0, 1);
        for(int i = 0;i < nums.length;i++){
            sum += nums[i];
            if(map.containsKey(sum - k)){
                cnt += map.get(sum - k);
            }
            map.put(sum, map.getOrDefault(sum, 0) + 1);
        }
        return cnt;
    }
}
```

## 132模式

https://leetcode-cn.com/problems/132-pattern/

```java
class Solution {
    public boolean find132pattern(int[] nums) {
        Deque<Integer> queue = new LinkedList();
        int k = Integer.MIN_VALUE;
        for(int i = nums.length - 1; i >= 0;i--){
            if(nums[i] < k) return true;
            while(!queue.isEmpty() && nums[i] > queue.peekLast()){
                k = Math.max(k, queue.pollLast());
            }
            queue.addLast(nums[i]);
        }
        return false;
    }
}
```



## 最大连续1的个数

https://leetcode-cn.com/problems/max-consecutive-ones-iii/

```java
class Solution {
    public int longestOnes(int[] nums, int k) {
        int left = 0, right = 0, cnt = 0;
        int maxLen = -1;
        while(right < nums.length){
            if(nums[right] == 0){
                cnt++;
            }
            while(cnt > k){
                int temp = nums[left];
                if(temp == 0){
                    cnt--;
                }
                left++;
            }
            maxLen = Math.max(maxLen, right - left + 1);
            right++;
        }
        return maxLen;
    }
}
```



## 子数组最大乘积

https://www.nowcoder.com/practice/9c158345c867466293fc413cff570356?tpId=117&tqId=37785&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public double maxProduct(double[] arr) {
    if(arr == null || arr.length == 0) return 0;
    // maxProduct记录上一个位置之前的最大值, minProduct表示上一个位置之前的最小值
    // res表示全局最大值
    double maxProduct = arr[0],minProduct = arr[0], res = arr[0];
    for(int i = 1;i < arr.length;i++){
        double temp_max = maxProduct;
        // 更新当前位置的最大值和最小值
        maxProduct = Math.max(arr[i], Math.max(maxProduct * arr[i], minProduct * arr[i]));
        minProduct = Math.min(arr[i], Math.min(temp_max * arr[i], minProduct * arr[i]));
        res = Math.max(res, maxProduct);
    }
    return res;
}
```

## **数组中的最长连续子序列**

https://www.nowcoder.com/practice/eac1c953170243338f941959146ac4bf?tpId=117&tqId=37810&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int MLS (int[] arr) {
    if(arr == null || arr.length == 0) return 0;
    Arrays.sort(arr);
    // 当前最长连续序列长度
    int maxLS = 1;
    // 当前长度
    int len = 1;
    for(int i = 0;i < arr.length - 1;i++){
        if(arr[i + 1] - arr[i] == 0){
            continue;
        }else if(arr[i + 1] - arr[i] == 1){
            len++;
        }else{
            len = 1;
        }
        maxLS = maxLS > len ? maxLS : len;
    }
    return maxLS;
}
```

## 最长递增子序列

https://www.nowcoder.com/practice/9cf027bf54714ad889d4f30ff0ae5481?tpId=117&tqId=37796&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

**动态规划方法:**

```java
public class NC91 {

    // 动态规划方法
    // 复杂度为N^2,运行超时
    public int[] LIS_v1 (int[] arr) {
        if(arr == null || arr.length == 0) return new int[0];
        // 初始值
        int[] dp = new int[arr.length];
        Arrays.fill(dp,1);

        // 最长子序列长度
        int maxLen = 1;

        for(int i = 1;i < arr.length;i++){
            for(int j = 0;j < i;j++){
                if(arr[i] > arr[j]){
                    dp[i] = Math.max(dp[i],dp[j] + 1);
                    maxLen = Math.max(maxLen, dp[i]);
                }
            }
        }

        int[] ret = new int[maxLen];
        for(int k = arr.length - 1;k >= 0;k--){
            if(dp[k] == maxLen){
                ret[--maxLen] = arr[k];
            }
        }
        return ret;
    }
}
```

**贪心 + 二分**:

```java
public class NC91 {
    // 贪心 + 二分查找方式
    public int[] LIS (int[] arr) {
        if(arr == null || arr.length == 0) return new int[0];
        // dp[i]表示以arr[i]结尾的元素的最长子序列长度
        int[] dp = new int[arr.length];
        Arrays.fill(dp,1);

        // 临时记录最长子序列
        int[] res = new int[arr.length];
        // res数组实际的长度
        int resLen = 1;
        res[0] = arr[0];

        for(int i = 1;i < arr.length;i++){
            // 当前值比序列最后一个值大，直接加入
            if(arr[i] > res[resLen - 1]){
                res[resLen++] = arr[i];
                dp[i] = resLen;
            }
            else{
                // 二分查找找到第一个大于等于arr[i]的元素位置
                int left = 0, right = resLen - 1;
                while(left <= right){
                    int mid = left + (right - left) / 2;
                    // 注意等于的位置
                    if(res[mid] < arr[i]){
                        left = mid + 1;
                    }else{
                        right = mid - 1;
                    }
                }
                // left就是第一个大于等于arr[i]的位置，替换该位置不会影响目前得到的最长子序列的长度，但是是另一个最长子序列
                res[left] = arr[i];
                dp[i] = left + 1;
            }
        }

        int[] ret = new int[resLen];
        for(int k = arr.length - 1;k >= 0;k--){
            if(dp[k] == resLen){
                ret[--resLen] = arr[k];
            }
        }
        return ret;

    }
}
```

## 俄罗斯套娃信封问题

https://leetcode-cn.com/problems/russian-doll-envelopes/

```java
class Solution {
    public int maxEnvelopes(int[][] envelopes) {
        if(envelopes.length == 1) return 1;

        Arrays.sort(envelopes, new Comparator<int[]>(){
            public int compare(int[] a,int[] b){
                return a[0] != b[0] ? a[0] - b[0] : b[1] - a[1];
            }
        });

        int[] weight = new int[envelopes.length];
        for(int i = 0;i < envelopes.length;i++){
            weight[i] = envelopes[i][1];
        }

        int[] top = new int[envelopes.length];
        int cnt = 0;
        for(int i = 0;i < weight.length;i++){
            int left = 0, right = cnt - 1;
            int lv = weight[i];
            while(left <= right){
                int mid = left + (right - left) / 2;
                if(top[mid] == lv){
                    right = mid - 1;
                }else if(top[mid] < lv){
                    left = mid + 1;
                }else{
                    right = mid - 1;
                }
            }
            if(left == cnt) cnt++;
            top[left] = lv;
        }
        return cnt;
    }
}
```



## **最长公共子序列-II**

https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11?tpId=117&tqId=37798&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public String LCS (String s1, String s2) {
    if (s1 == null || s2 == null || s1.length() == 0 || s2.length() == 0) return "-1";
    int rows = s1.length();
    int cols = s2.length();
    // dp[i][j]表示s1的前i个字符和s2的前j个字符的最长公共子序列
    int[][] dp = new int[rows + 1][cols + 1];
    // 填充DP数组，寻找最大长度
    for (int i = 0; i <= rows; i++) {
        for (int j = 0; j <= cols; j++) {
            if (i == 0 || j == 0) {
                dp[i][j] = 0;
            } else {
                if (s1.charAt(i - 1) == s2.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                }
            }
        }
    }
    StringBuilder sb = new StringBuilder();
    int i = rows, j = cols;
    while (i > 0 && j > 0) {
        if (s1.charAt(i - 1) == s2.charAt(j - 1)) {
            sb.insert(0, s1.charAt(i - 1));
            i--;
            j--;
        } else {
            if (dp[i - 1][j] > dp[i][j - 1]) {
                i--;
            } else {
                j--;
            }
        }
    }
    return sb.length() == 0 ? "-1" : sb.toString();
}
```

## **最长公共子串**

https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac?tpId=117&tqId=37799&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public String LCS (String str1, String str2) {
    int maxLen = 0, endI = -1;
    int str1Len  = str1.length(), str2Len = str2.length();
    // dp[i][j]表示str1中第i个字符之前的部分和str2第j个字符之前的部分的最长公共子串长度
    int[][] dp = new int[str1Len][str2Len];
    int i = 0, j = 0;
    for(i = 0;i < str1Len;i++){
        for(j = 0;j < str2Len;j++){
            if(str1.charAt(i) == str2.charAt(j)){
                if(i == 0 || j == 0){
                    dp[i][j] = 1;
                }else{
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                }
            }

            if(maxLen < dp[i][j]){
                maxLen = dp[i][j];
                endI = i;
            }
        }
    }
    return str1.substring(endI - maxLen + 1, endI + 1);
}
```

## 零钱兑换

https://leetcode-cn.com/problems/coin-change/

```java
class Solution {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, amount + 1);
        dp[0] = 0;
        for(int i = 1;i <= amount;i++){
            for(int j = 0;j < coins.length;j++){
                if(coins[j] <= i){
                    dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
                }
            }
        }
        return dp[amount] == amount + 1 ? -1 : dp[amount];
    }
}
```

## 零钱兑换2

https://leetcode-cn.com/problems/coin-change-2/

```java
class Solution {
    public int change(int amount, int[] coins) {
        int[] dp = new int[amount + 1];
        dp[0] = 1;
        for(int i = 0;i < coins.length;i++){
            for(int j = 1;j <= amount;j++){
                if(j >= coins[i]){
                    dp[j] += dp[j - coins[i]];
                }
            }
        }
        return dp[amount];
    }
}
```

## 将字符串转化为整数

https://leetcode-cn.com/problems/string-to-integer-atoi/

```java
class Solution {
    public int myAtoi(String s) {
        int len = s.length();
        int i = 0;
        long num = 0;
        boolean flag = true;
        for (i = 0; i < len && s.charAt(i) == ' '; i++) {
        }
        if (i < len) {
            if (s.charAt(i) == '+') {
                i++;
            } else if (s.charAt(i) == '-') {
                flag = false;
                i++;
            }
        }
        while (i < len && s.charAt(i) >= '0' && s.charAt(i) <= '9') {
            num = num * 10 + (s.charAt(i) - '0');
            if (flag && num > Integer.MAX_VALUE) {
                return Integer.MAX_VALUE;
            } else if(!flag && num < Integer.MIN_VALUE){
                return Integer.MIN_VALUE;
            }
            i++;
        }
        if (flag) {
            num = num > Integer.MAX_VALUE ? Integer.MAX_VALUE : num;
        } else {
            num = -num < Integer.MIN_VALUE ? Integer.MIN_VALUE : -num;
        }
        return (int) num;

    }
}
```

## 压缩字符串

https://leetcode-cn.com/problems/string-compression/

```java
public int compress(char[] chars) {
        // read表示当前读取的位置, write表示当前写入的位置, fisrt表示连续的第一个字符
        int read = 0, write = 0, first = 0;
        for(;read < chars.length;read++){
            // 当前位置是最后一个字符或者当前位置的字符是连续的最后一个字符
            if(read == chars.length - 1 || chars[read] != chars[read + 1]){
                // 写入连续的第一个字符
                chars[write] = chars[first];
                write++;
                // read大于first表示需要压缩
                if(read > first){
                    // 写入需要压缩的长度
                    int len = read - first + 1;
                    String s = String.valueOf(len);
                    for(int i = 0;i < s.length();i++){
                        chars[write++] = s.charAt(i);
                    }
                }
                // 更新连续的第一个字符
                first = read + 1;
            }
        }
        return write;
    }
```



## 字符串解码

https://leetcode-cn.com/problems/decode-string/

```java
public String decodeString(String s) {
        Stack<Integer> stack = new Stack<>();
        StringBuilder sb = new StringBuilder(s);
        for(int i = 0;i < sb.length();i++){
            char ch = sb.charAt(i);
            if(ch == '['){
                stack.push(i);
            }else if(ch == ']'){
                // 对应的[的位置
                int start = stack.peek();
                // [和]之间的字符串
                String pattern = sb.substring(start + 1, i);
                stack.pop();
                // [ 前面的数字
                int j = start - 1;
                while(j >= 0 && Character.isDigit(sb.charAt(j))){
                    j--;
                }
                int times = Integer.valueOf(sb.substring(j + 1, start));
                // 解码后的子串
                StringBuilder append = new StringBuilder();
                for(int k = 0;k < times;k++){
                    append.append(pattern);
                }
                sb.replace(j + 1, i + 1, append.toString());
                i = j + append.length();
            }
        }
        return sb.toString();
    }
```

## 把数字翻译成字符串

https://www.nowcoder.com/practice/046a55e6cd274cffb88fc32dba695668?tpId=117&tqId=37840&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int solve (String nums) {
    // 字符串为空或以0开头
    if(nums == null || nums.length() == 0 || nums.charAt(0) == '0') return 0;
    int len = nums.length();
    // dp[i]表示第i个位置之前的子串的翻译种数
    int[] dp = new int[len];
    dp[0] = 1;
    for(int i = 1;i < len;i++){
        // 当前位置不为0,将当前字符单独翻译
        if(nums.charAt(i) != '0'){
            dp[i] = dp[i - 1];
        }
        int num = (nums.charAt(i - 1)  - '0') * 10 + (nums.charAt(i) - '0');
        // 当前位置和前一个位置组成10~26之间的数
        if(num >= 10 && num <= 26){
            // 当前是第二个字符,加一种翻译方式
            if(i == 1){
                dp[i] += 1;
            }
            // 当前不是第二个字符
            else{
                dp[i] += dp[i - 2];
            }
        }
    }
    return dp[len - 1];
}
```

## 验证IP地址

https://www.nowcoder.com/practice/55fb3c68d08d46119f76ae2df7566880?tpId=117&tqId=37837&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public String solve (String IP) {
    if(IP == null || IP.length() == 0) {
        return "Neither";
    }
    if(validIPv4(IP)) return "IPv4";
    if(valirdIPv6(IP)) return "IPv6";
    return "Neither";
}

public boolean validIPv4(String IP){
    String[] strs = IP.split("\\.");
    // IPv4地址包括四部分
    if(strs.length != 4) return false;
    // 每一个部分不能超过255,且长度大于1时不能以0开头
    for (String str : strs) {
        if(Integer.valueOf(str) > 255 || (str.length() > 1 && str.charAt(0) == '0')) return false;
    }
    return true;
}

public boolean valirdIPv6(String IP){
    String[] strs = IP.split(":");
    // IPv6地址包括八部分
    if(strs.length != 8) return false;
    for(String part : strs){
        if(part.length() > 4) return false;
    }
    return true;
}
```

## 还原IP地址

https://leetcode-cn.com/problems/restore-ip-addresses/

```java
class Solution {
    private List<String> res = new ArrayList();

    public List<String> restoreIpAddresses(String s) {
        helper(s,-1,0, new ArrayList<String>());
        return res;
    }

    private void helper(String str, int lastDotIdx, int dotCnt, ArrayList<String> segments){
        if(dotCnt > 4) return;
        if(lastDotIdx == str.length() - 1){
            if(dotCnt == 4){
                res.add(String.join(".", segments));
            }
            return;
        }
        for(int i = lastDotIdx + 1;i < str.length() && i <= lastDotIdx + 3;i++){
            if(judge(str, lastDotIdx + 1, i)){
                String seg = str.substring(lastDotIdx + 1, i + 1);
                segments.add(seg);
                helper(str, i, dotCnt + 1, segments);
                segments.remove(segments.size() - 1);
            }
        }
    }

    private boolean judge(String str, int start, int end){
        // 每一个部分长度为1~3
        if(end - start + 1 >= 4){
            return false;
        }
        // 每一段可以只有一个0但是长度不为1时不能以0开头
        if(end - start > 0 && str.charAt(start) == '0') return false;
        // 计算对应的数
        int num = 0;
        for(int i = start;i <= end;i++){
            num = num * 10 + (str.charAt(i) - '0');
        }
        // 最大是255
        if(num > 255) return false;
        return true;
    }
}
```

## 解码方式

https://leetcode-cn.com/problems/decode-ways/

```java
class Solution {
    private int cnt = 0;
    public int numDecodings(String s) {
        // dp[i]表示从开头开始长度为i的子串的解码方式
        int[] dp = new int[s.length() + 1];
        // 空串有一种解码方式
        dp[0] = 1;
        for(int i = 1;i <= s.length();i++){
            // 当前字符单独解码
            if(s.charAt(i - 1) != '0') dp[i] += dp[i - 1];
            // 当前字符和前一个字符一起解码
            if(i > 1 && s.charAt(i - 2) != '0' && ((s.charAt(i - 2) - '0') * 10 + (s.charAt(i - 1) - '0') <= 26)){
                dp[i] += dp[i - 2];
            }
        }
        return dp[s.length()];
    }
}
```

## **岛屿数量**

https://www.nowcoder.com/practice/0c9664d1554e466aa107d899418e814e?tpId=117&tqId=37833&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
private void dfs(char[][] grid,int i,int j){
    // 边界情况
    if(i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == '0') return;
    // 标记为已访问
    grid[i][j] = '0';
    // 上
    if(i >= 1 && grid[i - 1][j] == '1' ) dfs(grid, i - 1, j);
    // 下
    if(i + 1 < grid.length && grid[i + 1][j] == '1' ) dfs(grid, i + 1, j);
    // 左
    if(j >= 1 && grid[i][j - 1] == '1' ) dfs(grid, i, j - 1);
    // 右
    if(j + 1 < grid[0].length && grid[i][j + 1] == '1' ) dfs(grid, i, j + 1);
}
public int solve (char[][] grid) {
    if(grid == null || grid.length == 0 || grid[0].length == 0) return 0;
    int lands = 0;
    // 每次从一个未访问的节点开始将其周围的岛屿全部标记为已访问
    for(int i = 0;i < grid.length;i++){
        for(int j = 0;j < grid[0].length;j++){
            if(grid[i][j] == '1'){
                lands ++;
                dfs(grid,i,j);
            }
        }
    }
    return lands;
}
```

## **接雨水问题**

https://leetcode-cn.com/problems/trapping-rain-water/

```java
public class NC128 {
    /**
     * 首先求容器边界,
     * 然后使用双指针,
     * 分别从两边往中间扫描,
     * 当左边的高度小于右边的高度时,左指针++,
     * 如果此时当前位置的高度小于容器的边界高度,这个位置上方有水,进行水量累加。
     * 反之，则右指针向左扫描-1。
     * @param arr
     * @return
     */
    public long maxWater (int[] arr) {
        if(arr == null || arr.length < 3) return 0L;
        int left = 0, right = arr.length - 1;
        long waterCnt = 0;
        int min = Math.min(arr[left],arr[right]);
        while(left < right){
            if(arr[left] < arr[right]){
                left ++;
                if(arr[left] < min){
                    waterCnt += (min - arr[left]);
                }else{
                    min = Math.min(arr[left],arr[right]);
                }
            }else{
                right--;
                if(arr[right] < min){
                    waterCnt += (min - arr[right]);
                }else{
                    min = Math.min(arr[left],arr[right]);
                }
            }
        }
        return waterCnt;
    }
}
```

## 盛最多水的容器

https://leetcode-cn.com/problems/container-with-most-water/

```java
class Solution {
    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int area = 0;
        while(left < right){
            int temp = Math.min(height[left], height[right]);
            area = Math.max(area, temp * (right - left));
            if(height[left] < height[right]){
                left++;
            }else{
                right--;
            }
        }
        return area;
    }
}
```



## 分糖果问题

https://leetcode-cn.com/problems/candy/

```java
class Solution {
    public int candy(int[] arr) {
        int sum = 0;
        if(arr == null || arr.length == 0) return sum;
        int len = arr.length;
        // 初始所有人分一颗糖
        int[] candies = new int[len];
        Arrays.fill(candies, 1);
        // 从左向右遍历, 右边的人得分比左边人高，右边人糖果数比左边人多一颗
        for(int i = 1;i < len;i++){
            if(arr[i] > arr[i - 1]){
                candies[i] = candies[i - 1] + 1;
            }
        }
        // 从右向左遍历, 左边的人得分比右边人高，左边人糖果数比右边人多一颗
        for(int i = len - 1;i > 0;i--){
            if(arr[i - 1] > arr[i]){
                candies[i - 1] = Math.max(candies[i - 1], candies[i] + 1);
            }
        }

        for(int i = 0;i < len;i++){
            sum += candies[i];
        }
        return sum;
    }
}
```

## **kmp算法**

https://www.nowcoder.com/practice/bb1615c381cc4237919d1aa448083bcc?tpId=117&tqId=37859&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public int kmp(String S, String T) {
    if (T == null || T.length() == 0) return 0;
    // S是模式串, T是匹配串
    int sLen = S.length();
    int tLen = T.length();
    // 初始化next数组
    int[] next = getNext(S);
    // 匹配个数
    int cnt = 0;
    int sIndex = 0, tIndex = 0;
    while (tIndex < tLen && sIndex < sLen) {
        // 当前字符相等, 继续匹配下一个
        if (sIndex == -1 || S.charAt(sIndex) == T.charAt(tIndex)) {
            tIndex++;
            sIndex++;
            // 找到一个完整的匹配, 模式串跳转到next位置, 匹配串位置不变
            if(sLen == sIndex){
                cnt++;
                sIndex = next[sIndex - 1];
                // 因为前面自增过所以需要还原
                tIndex--;
            }
        }
        // 模式串跳转到next位置
        else {
            sIndex = next[sIndex];
        }
    }
    return cnt;
}

private int[] getNext(String Pattern) {
    // next[i]表示模式串[0 ~ i]的最长公共前后缀的长度
    if(Pattern == null || Pattern.length() == 0) return new int[0];
    int len = Pattern.length();
    int[] next = new int[len];
    next[0] = -1;
    int i = -1;
    for(int j = 1;j < len;j++){
        // 当前串
        while(i != -1 && Pattern.charAt(j) != Pattern.charAt(i + 1)){
            i = next[i];
        }
        if(Pattern.charAt(i + 1) == Pattern.charAt(j)){
            i++;
        }
        next[j] = i;
    }
    return next;
}
```

## 最小生成树

https://www.nowcoder.com/practice/735a34ff4672498b95660f43b7fcd628?tpId=117&tqId=37869&companyId=665&rp=1&ru=%2Fcompany%2Fhome%2Fcode%2F665&qru=%2Fta%2Fjob-code-high%2Fquestion-ranking&tab=answerKey

```java
public class NC159 {

    // 边节点
    class Edge{
        int start;
        int end;
        int cost;
        public Edge(int start, int end, int cost) {
            this.start = start;
            this.end = end;
            this.cost = cost;
        }
    }
    // m表示路的数量
    public int miniSpanningTree (int n, int m, int[][] cost) {
        // 并查集
        int[] unionSet = new int[n + 1];
        // 初始所有节点所在集合的顶点
        Arrays.fill(unionSet, -1);
        // 小根堆
        PriorityQueue<Edge> queue = new PriorityQueue<>(new Comparator<Edge>() {
            @Override
            public int compare(Edge o1, Edge o2) {
                return o1.cost - o2.cost;
            }
        });
        for(int i = 0;i < m;i++){
            queue.add(new Edge(cost[i][0],cost[i][1],cost[i][2]));
        }
        int totalCost = 0;
        // 每次从边集合选出最小的边
        while(!queue.isEmpty()){
            Edge temp = queue.poll();
            // 如果两个顶点不在一个集合，将其并入一个集合
            if(find(unionSet, temp.start) != find(unionSet, temp.end)){
                totalCost += temp.cost;
                union(unionSet, temp.start, temp.end);
            }
        }
        return totalCost;
    }

    // 将两个顶点并入一个集合
    private void union(int[] set, int node1, int node2) {
        int temp1 = find(set, node1);
        int temp2 = find(set, node2);
        if(temp1 != temp2){
            set[temp1] = temp2;
        }
    }

    // 查找顶点所在集合的根顶点
    private int find(int[] set, int node) {
        while(set[node] != -1){
            node = set[node];
        }
        return node;
    }
}
```

## 省份数量

https://leetcode-cn.com/problems/number-of-provinces/

```java
class Solution {
    public int findCircleNum(int[][] isConnected) {
        int citys = isConnected.length;
        int[] set = new int[citys];
        Arrays.fill(set, -1);
        for(int i = 0;i < citys;i++){
            for(int j = i + 1;j < citys;j++){
                if(isConnected[i][j] != 0){
                    int rootX = findRoot(set, i);
                    int rootY = findRoot(set, j);
                    if(rootX != rootY){
                        set[rootX] = rootY;
                    }
                }
            }
        }
        int cnt = 0;
        for(int i = 0;i < citys;i++){
            if(set[i] == -1){
                cnt++;
            }
        }
        return cnt;
    }

    private int findRoot(int[] set, int x){
        int cur = x;
        while(set[cur] != -1){
            cur = set[cur];
        }
        return  cur;
    }
}
```



## 简单计算器

https://leetcode-cn.com/problems/basic-calculator-ii/

```java
class Solution {
    public int calculate(String s) {
        int num = 0;
        char flag = '+';
        Stack<Integer> stack = new Stack();
        for(int i = 0;i < s.length();i++){
            char ch = s.charAt(i);
            if(ch >= '0' && ch <= '9'){
                num = num * 10 + (ch - '0');
            }
            if(ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == s.length() - 1){
                if(flag == '+'){
                    stack.push(num);
                }else if(flag == '-'){
                    stack.push(-num);
                }else if(flag == '*'){
                    int top = stack.peek();
                    stack.pop();
                    stack.push(top * num);
                }else{
                    int top = stack.peek();
                    stack.pop();
                    stack.push(top / num);
                }
                num = 0;
                flag = ch;
            }
        }
        int sum = 0;
        while(!stack.isEmpty()){
            sum += stack.peek();
            stack.pop();
        }
        return sum;
    }
}
```

## 快排

https://leetcode-cn.com/problems/sort-an-array/

递归方式

```java
class Solution {
    public int[] sortArray(int[] nums) {
        quickSort(nums, 0, nums.length - 1);
        return nums;
    }
    public void quickSort(int[] nums, int start, int end){
        if(start < end){
            int pivot = getPivot(nums, start, end);
            quickSort(nums, start, pivot - 1);
            quickSort(nums, pivot + 1, end);
        }
    }
    private int getPivot(int[] nums, int start, int end){
        int left = start, right = end;
        int temp = nums[start];
        while(left < right){
            while(left < right && nums[right] >= temp) right--;
            nums[left] = nums[right];
            while(left < right && nums[left] <= temp) left++;
            nums[right] = nums[left];
        }
        nums[left] = temp;
        return left;
    }
}
```

非递归方式

```java
class Solution {
    Stack<Integer> stack = new Stack();
    public int[] sortArray(int[] nums) {
        stack.push(0);
        stack.push(nums.length - 1);
        quickSort(nums, 0, nums.length - 1);
        return nums;
    }

    public void quickSort(int[] nums, int start, int end){
        while(!stack.isEmpty()){
            int right = stack.pop();
            int left = stack.pop();
            int pivot = getPivot(nums, left, right);
            if(pivot > left){
                stack.push(left);
                stack.push(pivot - 1);
            }
            if(pivot < right){
                stack.push(pivot + 1);
                stack.push(right);
            }
        }
    }

    private int getPivot(int[] nums, int start, int end){
        int left = start, right = end;
        int temp = nums[left];
        while(left < right){
            while(left < right && nums[right] >= temp) right--;
            nums[left] = nums[right];
            while(left < right && nums[left] <= temp) left++;
            nums[right] = nums[left];
        }
        nums[left] = temp;
        return left;
    }
}
```

## 堆排序

https://leetcode-cn.com/problems/sort-an-array/

```java
class Solution {
    public int[] sortArray(int[] nums) {
        heapSort(nums);
        return nums;
    }

    private void heapSort(int[] nums){
        int len = nums.length;
        for(int pos = len / 2 - 1;pos >= 0;pos--){
            adjust(nums, pos, len);
        }
        for(int i = len - 1;i > 0;i--){
            swap(nums, 0, i);
            adjust(nums, 0, i);
        }
    }

    private void adjust(int[] nums, int pos, int len){
        int temp = nums[pos];
        for(int i = pos * 2 + 1; i < len;i = i * 2 + 1){
            if(i + 1 < len && nums[i] < nums[i + 1]){
                i++;
            }
            if(nums[i] > temp){
                nums[pos] = nums[i];
                pos = i;
            }
        }
        nums[pos] = temp;
    }

    private void swap(int[] nums, int i, int j){
        nums[i] = nums[i] ^ nums[j];
        nums[j] = nums[i] ^ nums[j];
        nums[i] = nums[i] ^ nums[j];
    }
}
```

## 2000W学生高考成绩排序

桶排序 + 基数排序

和计数排序类似，不同的是计数排序只能对整数进行排序，每个整数和数组的下标是对应的。而桶排序中的桶是一个范围，将数放到对应范围的桶里，然后对每个桶进行排序。

## 字典树

https://leetcode-cn.com/problems/implement-trie-prefix-tree/

```java
class Trie {

    private Trie[] children;

    private boolean isEnd = false;

    /** Initialize your data structure here. */
    public Trie() {
        children = new Trie[26];
        isEnd = false;
    }
    
    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie cur = this;
        for(int i = 0;i < word.length();i++){
            if(cur.children[word.charAt(i) - 'a'] == null){
                cur.children[word.charAt(i) - 'a'] = new Trie();
            }
            cur = cur.children[word.charAt(i) - 'a'];
        }
        cur.isEnd = true;
    }
    
    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie cur = this;
        for(int i = 0;i < word.length();i++){
            if(cur.children[word.charAt(i) - 'a'] == null){
                return false;
            }
            cur = cur.children[word.charAt(i) - 'a'];
        }
        return cur.isEnd;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        Trie cur = this;
        for(int i = 0;i < prefix.length();i++){
            if(cur.children[prefix.charAt(i) - 'a'] == null){
                return false;
            }
            cur = cur.children[prefix.charAt(i) - 'a'];
        }
        return true;
    }
}
```

## 从无限的字符流中随机选出 10 个字符

https://www.cnblogs.com/snowInPluto/p/5996269.html

采用蓄水池采样算法

假设数据规模是n，需要采样数量为k。

首先构建大小为k的数组，将前k个字符加入数组。从第k+1个字符开始，每个字符以k/n的概率决定是否留在数组中。遍历完所有元素后数组中的元素就是采样的结果。

## 打乱数组

https://leetcode-cn.com/problems/shuffle-an-array/

两种方法:

- 暴力方法，每次从数组中随机选出一个数并将其移除，直到数组为空；时间复杂度为O(n^2)，空间复杂度为O(n)，因为移除数组时需要遍历。
- Fisher-Yates 洗牌算法: 从当前位置到数组末尾之间随机选择一个数与当前位置交换，然后将当前位置加1，直到遍历完数组。时间复杂度和空间复杂度都为O(n)

```java
class Solution {

    private int[] initial;
    private int[] work;

    public Solution(int[] nums) {
        this.work = nums;
        this.initial = nums.clone();
    }

    /**
     * Resets the array to its original configuration and return it.
     */
    public int[] reset() {
        work = initial.clone();
        return work;
    }

    /**
     * Returns a random shuffling of the array.
     */
    public int[] shuffle() {
        for (int i = 0; i < this.work.length; i++) {
            Random rand = new Random();
            int idx = rand.nextInt(this.work.length - i) + i;
            if(idx != i){
                work[i] = work[i] ^ work[idx];
                work[idx] = work[i] ^ work[idx];
                work[i] = work[i] ^ work[idx];
            }
        }
        return this.work;
    }
}
```

## 比目标数大的最小的数

https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/

```java
class Solution {
    public void nextPermutation(int[] nums) {
        int i;
        for(i = nums.length - 2;i >= 0 && nums[i] >= nums[i + 1];i--){}
        int j;
        if(i >= 0) {
            for(j = nums.length - 1;j > i && nums[j] <= nums[i];j--){}
            nums[i] = nums[i] ^ nums[j];
            nums[j] = nums[i] ^ nums[j];
            nums[i] = nums[i] ^ nums[j];
        }
        for(i = i + 1,j = nums.length - 1;i < j;i++,j--){
            nums[i] = nums[i] ^ nums[j];
            nums[j] = nums[i] ^ nums[j];
            nums[i] = nums[i] ^ nums[j];
        }
    }
}
```

## 移除K位数字

https://leetcode-cn.com/problems/remove-k-digits/

```java
class Solution {
    public String removeKdigits(String num, int k) {
        Deque<Character> queue = new LinkedList<>();
        for(int i = 0;i < num.length();i++){
            while(k > 0 && !queue.isEmpty() && queue.peekLast() > num.charAt(i)){
                queue.pollLast();
                k--;
            }
            queue.addLast(num.charAt(i));
        }
        while(k > 0){
            queue.pollLast();
            k--;
        }
        StringBuilder sb = new StringBuilder();
        while(!queue.isEmpty() && queue.peekFirst() == '0'){
            queue.pollFirst();
        }
        while(!queue.isEmpty()){
            sb.append(queue.pollFirst());
        }
        if(sb.length() == 0) return "0";
        return sb.toString();
    }
}
```



## 贪心和动态规划有什么区别

都是将大问题分解成子问题，因此都需要最优子结构。

不同的是贪心算法对某一个子问题做出选择后，只需要对剩下的一个子问题进行求解，因此每一步的最优解一定包含上一步的最优解，上一步之前的最优解不需要保留。

而动态规划中，每一步的最优解不一定包含上一步的最优解，因此需要保留之前所有的最优解。

**贪心算法是动态规划的特例**

动态规划算法通常以自底向上的方式解各子问题，而贪心算法则通常自顶向下的方式进行

## rand7实现rand10

https://leetcode-cn.com/problems/implement-rand10-using-rand7/

```java
/**
 * The rand7() API is already defined in the parent class SolBase.
 * public int rand7();
 * @return a random integer in the range 1 to 7
 */
class Solution extends SolBase {
    public int rand10() {
        while(true){
            int num = (rand7() - 1) * 7 + rand7();
            if(num > 40){
                num = (num - 40 - 1) * 7 + rand7();
                if(num > 60){
                    num = (num - 60 - 1) * 7 + rand7();
                    if(num > 20){
                        continue;
                    }else{
                        return num % 10 + 1;
                    }
                }else{
                    return num % 10 + 1;
                }
            }else{
                return num % 10 + 1;
            }
        }
    }
}
```

## 最小差

https://leetcode-cn.com/problems/smallest-difference-lcci/

```java
class Solution {
    public int smallestDifference(int[] a, int[] b) {
        Arrays.sort(a);
        Arrays.sort(b);
        long minDistance = Long.MAX_VALUE;
        for(int i = 0, j = 0;i < a.length && j < b.length;){
            long distance = a[i] < b[j] ? (long)b[j] - (long)a[i] : (long)a[i] - (long)b[j];
            if(distance < minDistance){
                minDistance = distance;
            }
            if(a[i] < b[j]){
                i++;
            }else if(a[i] == b[j]){
                return 0;
            }else{
                j++;
            }
        }
        return (int)minDistance;
    }
}
```

## 快乐数

https://leetcode-cn.com/problems/happy-number/

### 哈希表法

```java
class Solution {
    public boolean isHappy(int n) {
        HashSet<Integer> set = new HashSet();
        int ret = n;
        while(ret != 1){
            ret = change(ret);
            if(set.contains(ret)){
                return false;
            }
            set.add(ret);
        }
        return true;
    }

    private int change(int n){
        int sum  = 0;
        while(n != 0){
            sum += (n % 10) * (n % 10);
            n = n / 10;
        }
        return sum;
    }
}
```



### 双指针法

```java
class Solution {
     public int getNext(int n) {
        int totalSum = 0;
        while (n > 0) {
            int d = n % 10;
            n = n / 10;
            totalSum += d * d;
        }
        return totalSum;
    }

    public boolean isHappy(int n) {
        int slowRunner = n;
        int fastRunner = getNext(n);
        while (fastRunner != 1 && slowRunner != fastRunner) {
            slowRunner = getNext(slowRunner);
            fastRunner = getNext(getNext(fastRunner));
        }
        return fastRunner == 1;
    }
}
```

## x的平方根

https://leetcode-cn.com/problems/sqrtx/

```java
// 二分法
class Solution {
    public int mySqrt(int x) {
        int left = 0, right = x, ret = 0;
        while(left <= right){
            int mid = left + (right - left) / 2;
            if((long)mid * mid <= x){
                ret = mid;
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        return ret;
    }
}
```

## 对角线遍历

https://leetcode-cn.com/problems/diagonal-traverse/

```java
class Solution {
    public int[] findDiagonalOrder(int[][] mat) {
        int rows = mat.length, cols = mat[0].length;
        int[] ret = new int[rows * cols];
        int pos = 0;
        int sum = 0;
        while(sum < rows + cols - 1){
            int x = sum >= rows ? rows - 1 : sum;
            int y = sum - x;
            while(x >= 0 && y < cols){
                ret[pos++] = mat[x][y];
                x--;
                y++;
            }
            sum++;
            if(sum >= rows + cols - 1) break;
            y = sum >= cols ? cols - 1 : sum;
            x = sum - y;
            while(y >= 0 && x < rows){
                ret[pos++] = mat[x][y];
                y--;
                x++;
            }
            sum++;
        }
        return ret;
    }
}
```

## 圆圈中最后剩下的数字

https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

```java
class Solution {
    public int lastRemaining(int n, int m) {
        int remain = 0;
        for(int i = 2;i <= n;i++){
            remain = (remain + m) % i;
        }
        return remain;
    }
}
```

## 课程表

https://leetcode-cn.com/problems/course-schedule/

```java
public class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        // 保存顶点之间的关系
        ArrayList<ArrayList<Integer>> edges = new ArrayList<>();
        // 每个节点的入度
        int[] indegree = new int[numCourses];
        // 保存所有入度为0的顶点
        Queue<Integer> queue = new LinkedList<>();
        for(int i = 0;i < numCourses;i++){
            edges.add(new ArrayList<>());
        }
        // 更新所有的节点的入度
        for(int[] pre : prerequisites){
            edges.get(pre[1]).add(pre[0]);
            indegree[pre[0]] ++;
        }
        for(int i = 0;i < numCourses;i++){
            if(indegree[i] == 0){
                queue.add(i);
            }
        }
        // 已访问的节点数
        int visited = 0;
        while(!queue.isEmpty()){
            Integer temp = queue.poll();
            visited++;
            // 将邻接点入度减1，如果入度等于0则加入到队列
            for(Integer neighbor : edges.get(temp)){
                indegree[neighbor]--;
                if(indegree[neighbor] == 0){
                    queue.add(neighbor);
                }
            }
        }
        return visited == numCourses;
    }
}
```

## 课程表2

https://leetcode-cn.com/problems/course-schedule-ii/

### 方法一: 拓扑排序

```java
class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        // 保存顶点之间的关系
        ArrayList<ArrayList<Integer>> edges = new ArrayList<>();
        // 每个节点的入度
        int[] indegree = new int[numCourses];
        // 保存所有入度为0的顶点
        Queue<Integer> queue = new LinkedList<>();
        // 返回结果
        int[] ret = new int[numCourses];
        int pos = 0;

        for(int i = 0;i < numCourses;i++){
            edges.add(new ArrayList<>());
        }
        // 更新所有的节点的入度
        for(int[] pre : prerequisites){
            edges.get(pre[1]).add(pre[0]);
            indegree[pre[0]] ++;
        }
        for(int i = 0;i < numCourses;i++){
            if(indegree[i] == 0){
                queue.add(i);
            }
        }
        // 已访问的节点数
        int visited = 0;
        while(!queue.isEmpty()){
            Integer temp = queue.poll();
            ret[pos++] = temp;
            visited++;
            // 将邻接点入度减1，如果入度等于0则加入到队列
            for(Integer neighbor : edges.get(temp)){
                indegree[neighbor]--;
                if(indegree[neighbor] == 0){
                    queue.add(neighbor);
                }
            }
        }
        if(visited == numCourses){
            return ret;
        }
        return new int[0];
    }
}
```

### 方法二: 深度优先搜索

```java
public class Solution {
    // 是否存在环
    private boolean hasCycle = false;
    // 保存结果
    private int[] ret;
    // 栈指针
    private int idx;

    public int[] findOrder(int numCourses, int[][] prerequisites) {
        ret = new int[numCourses];
        idx = numCourses - 1;
        // 标记节点访问状态，0 未访问，1 当前轮访问， 2之前轮访问
        int[] visited = new int[numCourses];
        // 保存顶点之间的关系
        ArrayList<ArrayList<Integer>> edges = new ArrayList<>();
        // 构建邻接表
        for(int i = 0; i < numCourses;i++){
            edges.add(new ArrayList<>());
        }
        for(int[] pre : prerequisites){
            edges.get(pre[1]).add(pre[0]);
        }
        // 每次从一个未访问的节点出发
        for(int i = 0;i < numCourses;i++){
            if(visited[i] == 0){
                dfs(edges, visited, i);
            }
        }
        // 有环
        if(hasCycle){
            return new int[0];
        }else {
            return ret;
        }
    }

    private void dfs(ArrayList<ArrayList<Integer>> edges, int[] visited, int courseNo){
        // 标记为本轮已访问
        visited[courseNo] = 1;
        for(int neighbor : edges.get(courseNo)){
            // 如果邻居节点本轮已经访问，说明有环
            if(visited[neighbor] == 1){
                hasCycle = true;
                return;
            }
            // 如果未访问则可以访问该邻居节点
            else if(visited[neighbor] == 0){
                dfs(edges, visited, neighbor);
                if(hasCycle){
                    return;
                }
            }
        }
        // 所有的邻居节点都访问完毕后将本节点标记为2，因为本节点之后必然不可能有环了，不会对其他节点造成影响
        visited[courseNo] = 2;
        // 将当前节点入栈(从栈顶到栈底)，由于是所有的邻居节点访问完毕后才入栈，所以邻居节点在靠近顶部的位置
        // 因此从栈底到栈顶就是一个排列
        ret[idx --] = courseNo;
    }

    public static void main(String[] args) {
        int[][] arr = {{1,0},{2,0},{3,1},{3,2}};
        Arrays.stream(new Solution().findOrder(4, arr)).forEach(System.out::println);
    }

}
```

## 颜色分类

https://leetcode-cn.com/problems/sort-colors/

```java
class Solution {
    public void sortColors(int[] nums) {
        int left = 0, right = nums.length - 1;
        for(int i = 0;i <= right;i++){
            while(i <= right && nums[i] == 2){
                swap(nums, i, right);
                right--;
            }
            if(nums[i] == 0){
                swap(nums, i, left);
                left++;
            }
        }
    }

    private void swap(int[] nums, int i,int j){
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
```

## 分割链表

https://leetcode-cn.com/problems/partition-list/

```java
class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode smallHead = new ListNode(-1);
        ListNode bigHead = new ListNode(-1);
        ListNode smallPos = smallHead;
        ListNode bigPos = bigHead;
        while(head != null){
            if(head.val < x){
                smallPos.next = head;
                smallPos = smallPos.next;
            }else{
                bigPos.next = head;
                bigPos = bigPos.next;
            }
            head = head.next;
        }
        bigPos.next = null;
        smallPos.next = bigHead.next;
        return smallHead.next;
    }
}
```

## 进制转换

https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/

```java
class Solution {
    public String toHex(int num) {
        if(num == 0) return "0";
        // 数字和十六进制的映射关系
        char[] decimal2HexMap = "0123456789abcdef".toCharArray();
        StringBuilder sb = new StringBuilder();
        while(num != 0){
            // 取余数
            int temp = (num & 15);
            sb.append(decimal2HexMap[temp]);
            // 商继续作为被除数
            num = num >>> 4;
        }
        return sb.reverse().toString();
    }
}
```

## 鸡蛋掉落

https://leetcode-cn.com/problems/super-egg-drop/

```java
class Solution {
    private int[][] dp;
    public int superEggDrop(int k, int n) {

        dp = new int[k + 1][n + 1];
        for(int i = 0;i < k;i++){
            Arrays.fill(dp[i], Integer.MAX_VALUE);
        }
        helper(k, n);
        return dp[k][n];
    }

    private int helper(int k, int n){
        if(k == 1) return n;
        if(n == 0) return 0;
        int minCnt = Integer.MAX_VALUE;
        if(dp[k][n] != Integer.MAX_VALUE) return dp[k][n];
        for(int i = 1;i <= n;i++){
            minCnt = Math.min(Math.max(helper(k, n - i), helper(k - 1, i - 1)) + 1, minCnt);
        }
        dp[k][n] = minCnt;
        return minCnt;
    }
}
```

## 跳跃游戏

https://leetcode-cn.com/problems/jump-game/

```java
class Solution {
    public boolean canJump(int[] nums) {
        // 记录当前能跳跃到的最大位置
        int rightMost = 0;
        for(int i = 0;i < nums.length;i++){
            // 当前位置处于可以跳跃的范围之内
            if(rightMost >= i){
                rightMost = Math.max(rightMost, i + nums[i]);
                if(rightMost >= nums.length - 1) return true;
            }
        }
        return false;
    }
}
```

## 01矩阵

https://leetcode-cn.com/problems/01-matrix/

```java
public class Solution {
    public int[][] updateMatrix(int[][] matrix) {
        int rows = matrix.length;
        int cols = matrix[0].length;

        // 保存最终的结果
        int[][] res = new int[rows][cols];
        Queue<int[]> sources = new LinkedList<>();
        // 初始化距离
        for(int i = 0;i < rows;i ++){
            for(int j = 0;j < cols;j++){
                if(matrix[i][j] == 'O'){
                    sources.add(new int[]{i, j});
                    res[i][j] = 0;
                }else{
                    res[i][j] = Integer.MAX_VALUE;
                }
            }
        }
        // 上右下左,dir[i][0]表示x坐标变化, dir[i][1]表示y坐标变化
        int[][] dir = new int[][]{{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
        while(!sources.isEmpty()){
            int[] point = sources.poll();
            for(int i = 0;i < 4;i++){
                int x = point[0] + dir[i][0], y = point[1] + dir[i][1];
                if(x >= 0 && x < rows && y >= 0 && y < cols && res[x][y] > res[point[0]][point[1]] + 1){
                    res[x][y] = res[point[0]][point[1]] + 1;
                    sources.add(new int[]{x, y});
                }
            }
        }
        return res;
    }
}
```


