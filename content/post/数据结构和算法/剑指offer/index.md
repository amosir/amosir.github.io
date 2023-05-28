---
title: "剑指offer"
date: 2022-10-28T13:36:31+08:00
draft: false
image: "img/Q6aVkCJL8Uw.jpg"
categories: 
tag:
---


## 数组

### 二维数组的查找

> 题目描述

在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

[网址](https://www.nowcoder.com/practice/abc3fe2ce8e146608e868a70efebf62e?tpId=13&tqId=11154&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
class Solution {
public:
    bool Find(int target, vector<vector<int> > array) {
        // 多少行
        int xLen = array.size();
        if(xLen <= 0){
            return false;
        }
        // 多少列
        int yLen = array[0].size();
        if(yLen <= 0){
            return false;
        }

        bool flag = false;

        for(int i = 0;i < xLen;i++){
            // 在当前行查找
            if(array[i].at(0) <= target && target <= array[i].at(yLen - 1)){
                // 找到
                if(binary_search(array[i].begin(),array[i].end(),target)){
                    flag = true;
                    break;
                }
            }
        }
        return flag;
    }
};
```

### 数组中重复的数字

> 题目描述

在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为 7 的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字 2。

```cpp
class Solution {
public:
    // Parameters:
    //        numbers:     an array of integers
    //        length:      the length of array numbers
    //        duplication: (Output) the duplicated number in the array number
    // Return value:       true if the input is valid, and there are some duplications in the array number
    //                     otherwise false
    bool duplicate(int numbers[], int length, int* duplication) {
        sort(numbers,numbers + length);
        for(int i = 0;i < length - 1;i++){
            if(numbers[i] == numbers[i+1]){
                *duplication = numbers[i];
                return true;
            }
        }
        duplication = NULL;
        return  false;
    }
};
```

### 构建乘积数组

> 题目描述

给定一个数组 A[0,1,...,n-1],请构建一个数组 B[0,1,...,n-1],其中 B 中的元素 B[i]=A[0]_A[1]_...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。

```cpp
class Solution {
public:
    vector<int> multiply(const vector<int>& A) {
        int len = A.size();
        vector<int> ret(len);
        for(int i = 0;i < len;i++){
            int temp = 1;
            for(int j = 0;j < len;j++){
                if(j!=i){
                    temp = temp * A[j];
                }
            }
            ret[i] = temp;
        }
        return ret;
    }
};
```

## 字符串

### 替换空格

> 题目描述

请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为 We Are Happy.则经过替换之后的字符串为 We%20Are%20Happy。

```cpp
class Solution {
public:
	void replaceSpace(char *str,int length) {
       for(int i = 0;i < length;i++){
           if(str[i] == ' '){
               // 空格后面的字符向后移动
               for(int j = length - 1;j > i;j--){
                   str[j+2] = str[j];
               }
               // 修改总长度
               length  +=  2;
               str[i++] = '%';
               str[i++] = '2';
               // 这里不能加 1，因为for循环里面会加1
               str[i] = '0';
           }
       }
       str[length] =  '\0';
	}
};
```

### 正则表达式匹配

> 题目描述

请实现一个函数用来匹配包括'.’和`*`正则表达式。模式中的字符'.'表示任意一个字符，而`*`表示它前面的字符可以出现任意次（包含 0 次）。 在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab\*a"均不匹配

[网址](https://www.nowcoder.com/practice/45327ae22b7b413ea21df13ee7d6429c?tpId=13&tqId=11205&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

**递归回朔法**

```c++
/**
 * 递归回朔法:
 * 1. 没有*的情况: str[i] == pattern[i] || pattern[i] == '.'
 * 2. *在pattern中出现0次:保持str不变，pattern剪去前两个字符，然后继续递归
 * 3. *在pattern中出现1次或多次:比较首字符，首字符相等则剪去str的首字符,pattern不变，然后进行递归
 */
class Solution
{
private:
    bool match(string s, string p)
    {
        // 模式串p为空则根据s判断
        if (p.empty())
        {
            return s.empty();
        }
        bool firstMatch = !s.empty() && (s[0] == p[0] || p[0] == '.');

        // 如果下一个字符是*
        if (p.size() >= 2 && p[1] == '*')
        {
            return match(s, p.substr(2)) || (firstMatch && match(s.substr(1), p));
        }
        // 一般情况
        return firstMatch && match(s.substr(1), p.substr(1));
    }

public:
    bool match(char *str, char *pattern)
    {
        string s(str);
        string p(pattern);
        return match(s, p);
    }
};
```

**动态规划法 **

```cpp
class Solution
{
private:
    bool first_match(string s, int i, string p, int j)
    {
        return (s[i] == p[j] || p[j] == '.');
    }

public:
    bool match(char *str, char *pattern)
    {
        string s(str);
        string p(pattern);

        int len_s = s.length();
        int len_p = p.length();
        // dp[i][j] 表示s的前i个字符是否和p的前j个字符匹配
        bool dp[len_s + 1][len_p + 1];
        for (int i = 0; i <= len_s; i++)
        {
            for (int j = 0; j <= len_p; j++)
            {
                dp[i][j] = false;
            }
        }
        // 匹配串和模式串都是空串的情况
        dp[0][0] = true;

        for (int i = 0; i < len_s; i++)
        {
            for (int j = 0; j < len_p; j++)
            {
                // 模式串当前元素为*,有两种情况:
                // 1, *前面的字符出现0次，则dp[i + 1][j + 1] = dp[i + 1][j - 1]
                // 2. *前面的字符出现1次或多次，比较当前字符是否匹配和dp[i][j + 1](即主串去掉当前
                // 元素后是否与模式串匹配)
                if (p[j] == '*')
                {
                    dp[i + 1][j + 1] = dp[i + 1][j - 1] ||
                                       (first_match(s, i, p, j - 1) && dp[i][j + 1]);
                }
                // 没有*则比较当前元素是否匹配和当前位置之前的元素是否匹配
                else
                {
                    dp[i + 1][j + 1] = first_match(s, i, p, j) && dp[i][j];
                }
            }
        }
        return dp[len_s][len_p];
    }
};
```

### 字符流中第一个不重复的字符

> 题目描述

请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"。如果当前字符流没有存在出现一次的字符，返回#字符。

[网址](https://www.nowcoder.com/practice/00de97733b8e4f97a3fb5c680ee10720?tpId=13&tqId=11207&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution
{
private:
    vector<char> str;
    map<char,int> char_map;
public:
  //Insert one char from stringstream
    void Insert(char ch)
    {
        str.push_back(ch);
        map<char,int>::iterator it = char_map.find(ch);
        if(it == char_map.end()){
          char_map.insert(pair<char,int>(ch,1));
        }else{
          it->second ++;
        }
    }
  //return the first appearence once char in current stringstream
    char FirstAppearingOnce()
    {
      int len = str.size();
      for(int i = 0;i < len;i++){
        map<char,int>::iterator it = char_map.find(str[i]);
        if(it->second == 1){
          return it->first;
        }
      }
      return  '#';
    }

};
```

### 表示数值的字符串

> 题目描述

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100","5e2","-123","3.1416"和"-1E-16"都表示数值。 但是"12e","1a3.14","1.2.3","+-5"和"12e+4.3”都不是。

[网址](https://www.nowcoder.com/practice/6f8c901d091949a5837e24bb82a731f2?tpId=13&tqId=11206&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    bool isNumeric(char* string)
    {
        if(string == NULL) return false;
        // 符号、E、.是否出现过
        bool flag = false,hasE = false,hasDot = false;
        for(int i = 0;string[i]!='\0';i++){
            if(string[i] == 'E' || string[i] == 'e'){
                // e只能出现一次且不在最后一位
                if(hasE || string[i+1] == '\0'){
                    return false;
                }
                hasE = true;
            }else if(string[i] == '.'){
                // .只能出现一次且不能出现在最后一位
                if(hasE || hasDot || string[i + 1] == '\0'){
                    return false;
                }
                hasDot = true;
            }else if(string[i] == '+' || string[i] == '-'){
                // 第一次出现只能在第一位或出现在E后面且不能出现在E后面
                if(!flag && i != 0 && string[i - 1] != 'e' &&
                    string[i - 1] != 'E'){
                        return false;
                    }
                // 第二次出现出现在E后面
                if(flag && string[i - 1]!='e' && string[i-1]!='E'){
                    return false;
                }
                flag = true;
            }else if(string[i] < '0' || string[i] > '9'){
                return false;
            }
        }
    return true;
    }
};
```

## 链表

### 从尾到头打印链表

> 题目描述

输入一个链表，按链表从尾到头的顺序返回一个 ArrayList。

```c++
/**
*  struct ListNode {
*        int val;
*        struct ListNode *next;
*        ListNode(int x) :
*              val(x), next(NULL) {
*        }
*  };
*/
class Solution {
public:
    vector<int> printListFromTailToHead(ListNode* head) {
        vector<int> ret;
        if(head != NULL){
            ret = printListFromTailToHead(head->next);
            ret.push_back(head->val);
        }
        return ret;
    }
};
```

### 链表中环的入口节点

> 题目描述

给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出 null。

```cpp
// 相遇时:
// 快指针路程=a+(b+c)k+b ，k>=1  其中b+c为环的长度，k为绕环的圈数（k>=1,即最少一圈，不能是0圈，不然和慢指针走的一样长，矛盾）。
// 慢指针路程=a+b
// 快指针走的路程是慢指针的两倍，所以：
// (a+b）*2=a+(b+c)k+b
// 化简可得：
// a=(k-1)(b+c)+c
// 这个式子的意思是:链表头到环入口的距离=相遇点到环入口的距离+（k-1）圈环长度。其中k>=1,所以k-1>=0圈。
// 所以两个指针分别从链表头和相遇点出发，最后一定相遇于环入口。

struct ListNode {
        int val;
        struct ListNode *next;
        ListNode(int x) :
              val(x), next(NULL) {
        }
};
class Solution {
public:
    ListNode* EntryNodeOfLoop(ListNode* pHead)
    {
        // 链表为空或只有一个节点
        if(pHead == NULL || pHead->next == NULL){
            return NULL;
        }
        // 快慢指针
        ListNode *fast = pHead,*slow = pHead;
        ListNode *cutIn = NULL;


        while(fast != NULL && fast->next != NULL){
            fast = fast->next->next;
            slow = slow->next;
            if(fast == slow){
                // 交汇点
                cutIn = fast;
                slow = pHead;
                while(slow!=fast){
                    slow = slow->next;
                    fast = fast->next;
                }
                return slow;
            }
        }

        return NULL;
    }
};
```

### 删除链表中重复的结点

> 题目描述

在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。 例如，链表 1->2->3->3->4->4->5 处理后为 1->2->5

```cpp
class Solution {
public:
    ListNode* deleteDuplication(ListNode* pHead){
        // 空链表或只有一个节点
        if(pHead == NULL || pHead->next == NULL){
            return pHead;
        }
        // 防止头结点也是重复节点，需要新建一个头结点
        ListNode *newHead = new ListNode(0);
        newHead->next = pHead;
        // 指向上一个不重复的节点
        ListNode *pre = newHead;
        // 当前工作指针
        ListNode *work = newHead->next;
        while(work != NULL){
            // 当前节点不是最后一个节点且是重复节点
            if(work->next != NULL && work->val == work->next->val){
                //找到重复的最后一个节点
                while(work->next != NULL && work->val == work->next->val){
                    work = work->next;
                }
                pre->next = work->next;
                work = work->next;
            }
            // 不是重复节点则向后继续处理其他节点
            else{
                pre = pre->next;
                work =  work->next;
            }
        }
        return newHead->next;
    }
};
```

## 树

### 重建二叉树

> 题目描述

输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

[网址](https://www.nowcoder.com/practice/8a19cbe657394eeaac2f6ea9b0f6fcf6?tpId=13&tqId=11157&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
class Solution {
public:
    TreeNode* reConstructBinaryTree(vector<int> pre,vector<int> vin) {
        if(pre.size() <= 0 || vin.size() <= 0){
            return NULL;
        }
        return reConstructBinaryTree(pre,0,pre.size() - 1,vin,0,vin.size() - 1);
    }

    // rootIndex 是根节点在前序中的索引
    TreeNode* reConstructBinaryTree(vector<int> pre,int preStart,int preEnd,vector<int> vin,int inStart,int inEnd) {
        if(preStart <= preEnd && inStart <= inEnd){
            TreeNode *root = new TreeNode(pre[preStart]);
            // 根节点在中序序列中的索引
            int rootInVin;
            for(rootInVin = inStart;vin[rootInVin] != pre[preStart];rootInVin++){
            }

            root->left = reConstructBinaryTree(pre,preStart + 1,
                    preStart + rootInVin - inStart,vin,inStart,rootInVin - 1);

            root->right = reConstructBinaryTree(pre,preEnd - (inEnd - rootInVin) + 1,
                    preEnd,vin,rootInVin + 1,inEnd);

            return root;
        }
        return NULL;
    }
};
```

### 二叉树的下一个节点

> 题目描述

给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

[网址](https://www.nowcoder.com/practice/9023a0c988684a53960365b889ceaf5e?tpId=13&tqId=11210&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
/**
 *  1. 如果节点有右子节点，则右子节点的最左节点是该节点的下一个节点
 *  2. 如果节点无右子节点，但该节点是父节点的左子节点，则父节点是该节点的下一个节点。
 *  3. 如果节点无右子节点，且该节点是父节点的右子节点，则沿着父节点的指针向上遍历
 */
class Solution {
public:
    TreeLinkNode* GetNext(TreeLinkNode* pNode)
    {
        // 当前节点为空
        if(pNode == NULL){
            return NULL;
        }
        // 右子节点存在
        if(pNode->right!=NULL){
            TreeLinkNode *work = pNode->right;
            while(work->left){
                work = work->left;
            }
            return work;
        }
        // 右子节点不存在,向上回溯到第一个存在左节点的节点
        TreeLinkNode *work = pNode;
        TreeLinkNode *workParent = pNode->next;
        while(workParent!=NULL && workParent->right == work){
            work = workParent;
            workParent = work->next;
        }
        return workParent;
    }
};
```

### 对称二叉树

> 题目描述

请实现一个函数，用来判断一颗二叉树是不是对称的。注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。

[网址](https://www.nowcoder.com/practice/ff05d44dfdb04e1d83bdbdab320efbcb?tpId=13&tqId=11211&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};
class Solution {
public:
    bool isSymmetrical(TreeNode *node1,TreeNode *node2){
        // 对称位置两个节点为空
        if(node1 == NULL && node2 == NULL){
            return true;
        }
        // 对称位置只有一个为空
        else if(!node1 || !node2){
            return false;
        }
        // 对称位置都不为空，需要值相等且子节点对应位置也对称
        else{
            return node1->val == node2->val && isSymmetrical(node1->left,node2->right)
               && isSymmetrical(node1->right,node2->left);
        }
    }
    bool isSymmetrical(TreeNode* pRoot)
    {
        if(pRoot == NULL)
            return true;
        return isSymmetrical(pRoot->left,pRoot->right);
    }
};
```

### 把二叉树打印成多行

> 题目描述

从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。

[网址](https://www.nowcoder.com/practice/445c44d982d04483b04a54f298796288?tpId=13&tqId=11213&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};
class Solution {
public:
        vector<vector<int> > Print(TreeNode* pRoot) {
            // 存放最终的结果
            vector<vector<int> > ret;

            // 暂存每一行
            vector<int> tempVec;

            TreeNode *work = pRoot;

            // 空树直接返回
            if(pRoot == NULL)
                return ret;

            // 辅助队列
            queue<TreeNode*> qu;
            qu.push(pRoot);
            int count = 1;
            while(qu.size() > 0){
                work = qu.front();
                tempVec.push_back(work->val);
                qu.pop();
                count --;
                if(work->left){
                    qu.push(work->left);
                }
                if(work->right){
                    qu.push(work->right);
                }
                // 某一行全部出队，将该行添加到最终结果中
                if(count == 0){
                    ret.push_back(tempVec);
                    count = qu.size();
                    tempVec.clear();
                }
            }
            return ret;
        }

};
```

### 二叉搜索树的第 k 小节点

> 题目描述

给定一棵二叉搜索树，请找出其中的第 k 小的结点。例如， （5，3，7，2，4，6，8） 中，按结点数值大小顺序第三小结点的值为 4。

[网址](https://www.nowcoder.com/practice/ef068f602dde4d28aab2b210e859150a?tpId=13&tqId=11215&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};
class Solution {
public:
    TreeNode* KthNode(TreeNode* pRoot, int k)
    {
        // 边界检查
        if(!pRoot || k <= 0)
            return NULL;
        // 二叉搜索树中序序列是有序的
        vector<TreeNode*> inOrder;
        // 辅助栈
        vector<TreeNode*> vec;

        TreeNode *work = pRoot;
        while(vec.size() > 0 || work){
            if(work){
                vec.push_back(work);
                work = work->left;

            }else{
                work = vec.back();
                vec.pop_back();
                inOrder.push_back(work);
                work = work->right;
            }
        }

        // 中序序列的第k个元素
        if(inOrder.size() >= k){
            return inOrder[k - 1];
        }
        return NULL;
    }
};
```

### 数据流中的中位数

> 题目描述

如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用 Insert()方法读取数据流，使用 GetMedian()方法获取当前读取数据的中位数。

[网址](https://www.nowcoder.com/practice/9be0172896bd43948f8a32fb954e1be1?tpId=13&tqId=11216&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
#include<functional>
#include <queue>
using namespace std;
/***
    对于一个升序排序的数组，中位数为左半部分的最大值，右半部分的最小值，而左右两部分可以是无需的，只要保证左半部分的数均小于
    右半部分即可。因此，左右两半部分分别可用最大堆、最小堆实现。
    首先定义：如果有奇数个数，则中位数放在左半部分；如果有偶数个数，则取左半部分的最大值、右边部分的最小值之平均值
    分两种情况讨论：
    当目前有偶数个数字时，数字先插入最大堆，然后选择最大堆的最大值插入最小堆（第一个数字插入左半部分的最小堆）
    当目前有奇数个数字时，数字先插入最小堆，然后选择最小堆的最小值插入最大堆
    如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
    ***/
class Solution {
public:
    void Insert(int num)
    {
        cnt++;
        // 当前有偶数个值，先插入到左边大顶堆，再将左边大顶堆的最大值插入右边小顶堆
        if(cnt % 2 == 0){
            leftMaxHeap.push(num);
            int max = leftMaxHeap.top();
            leftMaxHeap.pop();
            rightMinHeap.push(max);
        }
        // 当前有奇数个值，先插入到右边小顶堆，再将右边小顶堆堆的最小值插入左边大顶堆
        else{
            rightMinHeap.push(num);
            int min = rightMinHeap.top();
            rightMinHeap.pop();
            leftMaxHeap.push(min);
        }
    }

    double GetMedian()
    {
        // 偶数个
        if((cnt & 1) == 0){
            return (leftMaxHeap.top() + rightMinHeap.top()) / 2.0;
        }
        // 奇数个
        else{
            return leftMaxHeap.top();
        }

    }

private:
    // 左边的大顶堆
    priority_queue<int,vector<int>,less<int> > leftMaxHeap;
    // 右边的小顶堆
    priority_queue<int,vector<int>,greater<int> > rightMinHeap;
    // 总数
    int cnt = 0;
};
```

### 系列化二叉树

> 题目描述

请实现两个函数，分别用来序列化和反序列化二叉树

二叉树的序列化是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#），以 ！ 表示一个结点值的结束（value!）。

二叉树的反序列化是指：根据某种遍历顺序得到的序列化字符串结果 str，重构二叉树。

[网址](https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84?tpId=13&tqId=11214&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL)
    {
    }
};
class Solution
{
public:
    void SerializeCore(TreeNode *root)
    {
        if (root == NULL)
        {
            stream.push_back('#');
            stream.push_back('!');
            return;
        }
        stream += to_string(root->val);
        stream.push_back('!');
        SerializeCore(root->left);
        SerializeCore(root->right);
    }
    char *Serialize(TreeNode *root)
    {
        if (!root)
        {
            return NULL;
        }
        SerializeCore(root);
        int len = stream.length();
        char *ret = new char(len + 1);
        for (int i = 0; i < len; i++)
        {
            ret[i] = stream[i];
        }
        ret[len] = '\0';
        return ret;
    };

    // 分割字符串
    void split_string(string str, char token)
    {
        string temp = "";
        int len = str.length();
        for (int i = 0; i < len; i++)
        {
            if (str[i] != token)
            {
                temp.push_back(str[i]);
            }
            else
            {
                vec.push_back(temp);
                temp = "";
            }
        }
    };

    TreeNode *Deserialize(char *str)
    {
        if (str == NULL)
        {
            return NULL;
        }
        return DeserializeCore(str);
    };

    TreeNode *DeserializeCore(char *str)
    {
        string s(str);
        split_string(s, '!');
        if (pos < vec.size())
        {
            if (vec[pos] == "#")
            {
                pos++;
                return NULL;
            }
            else
            {
                TreeNode *root = new TreeNode(stoi(vec[pos]));
                pos++;
                root->left = DeserializeCore(str);
                root->right = DeserializeCore(str);
                return root;
            }
        }
        else
        {
            return NULL;
        }
    }

private:
    string stream;
    vector<string> vec;
    int pos = 0;
};
```

### 之字形打印二叉树

> 题目描述

请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。

[网址](https://www.nowcoder.com/practice/91b69814117f4e8097390d107d2efbe0?tpId=13&tqId=11212&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};
class Solution {
public:
    vector<vector<int> > Print(TreeNode* pRoot) {
        // 最终结果
        vector<vector<int> > result;
        // 每一行
        vector<int> row;

        if(pRoot == NULL){
            return result;
        }
        // 奇数层辅助栈
        vector<TreeNode*> oddLayervector;
        // 偶数层辅助栈
        vector<TreeNode*> evenLayervector;

        TreeNode *work = pRoot;
        oddLayervector.push_back(work);
        // 只要有一个不为空
        while(!oddLayervector.empty() || !evenLayervector.empty()){
            // 当前为偶数层
            while(!evenLayervector.empty()){
                work = evenLayervector.back();
                row.push_back(work->val);
                evenLayervector.pop_back();
                if(work->right){
                    oddLayervector.push_back(work->right);
                }
                if(work->left){
                    oddLayervector.push_back(work->left);
                }
            }
            if(row.size()){
                result.push_back(row);
                row.clear();
            }
            // 当前层是奇数层
            while(!oddLayervector.empty()){
                work = oddLayervector.back();
                row.push_back(work->val);
                oddLayervector.pop_back();
                if(work->left){
                    evenLayervector.push_back(work->left);
                }
                if(work->right){
                    evenLayervector.push_back(work->right);
                }
            }
            if(row.size()){
                result.push_back(row);
                row.clear();
            }
        }
        return result;
    }

};
```

## 栈和队列

### 用栈实现队列

> 题目描述

用两个栈来实现一个队列，完成队列的 Push 和 Pop 操作。 队列中的元素为 int 类型。

[网址](https://www.nowcoder.com/practice/54275ddae22f475981afa2244dd448c6?tpId=13&tqId=11158&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
class Solution
{
public:
    void push(int node) {
        stack1.push(node);
    }

    int pop() {
        if(stack2.empty()){
            while(!stack1.empty()){
                stack2.push(stack1.top());
                stack1.pop();
            }
        }
        int to_p = stack2.top();
        stack2.pop();
        return to_p;
    }

private:
    stack<int> stack1;
    stack<int> stack2;
};
```

### 滑动窗口的最大值

> 题目描述

给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小 3，那么一共存在 6 个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下 6 个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

[网址](https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788?tpId=13&tqId=11217&tPage=4&rp=4&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution
{
public:
    vector<int> maxInWindows(const vector<int> &num, unsigned int size)
    {
        // 保存最终结果
        vector<int> result;
        // 窗口，前面的元素一定大于后面的元素，这样前面的元素就是最大值
        deque<int> window;
        int len = num.size();
        if (size == 0 || len == 0 || len < size)
        {
            return result;
        }
        // 第一个窗口
        for (int i = 0; i < size; i++)
        {
            // 如果当前遍历元素比窗口尾部的元素大，那么尾部的元素必然不是
            // 最大值，出队
            while (!window.empty() && num[i] > num[window.back()])
            {
                window.pop_back();
            }
            // 当前元素比尾部元素小，尾部元素出队后依然可能成为最大值，入队
            window.push_back(i);
        }

        result.push_back(num[window.front()]);

        for (int i = size; i < len; i++)
        {
            while (!window.empty() && num[i] > num[window.back()])
            {
                window.pop_back();
            }
            window.push_back(i);
            // 队头元素已滑出窗口，出队
            if (!window.empty() && window.front() <= i - size)
            {
                window.pop_front();
            }
            result.push_back(num[window.front()]);
        }
        return result;
    }
};
```

## 查找和排序

### 旋转数组的最小数字

> 题目描述

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为 1。
NOTE：给出的所有元素都大于 0，若数组大小为 0，请返回 0。

[网址](https://www.nowcoder.com/practice/9f3231a991af4f55b95579b44b7a01ba?tpId=13&tqId=11159&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
class Solution {
public:
    int minNumberInRotateArray(vector<int> rotateArray) {
        if(rotateArray.size() == 0){
            return 0;
        }
        int len = rotateArray.size();
        //最小值所在的索引
        int min = 0;
        for(int i = 0;i < len;i++){
            if(rotateArray[i] < rotateArray[min]){
                min = i;
            }
        }
        rotate(rotateArray,0,len - 1);
        rotate(rotateArray,len - min,len - 1);
        rotate(rotateArray,0,len - 1 - min);
        return rotateArray[0];
    }
    void rotate(vector<int> &vec,int start,int end){
        while(start < end){
            int temp = vec[start];
            vec[start] = vec[end];
            vec[end] = temp;
            start++;
            end--;
        }
    }
};
```

## 递归和循环

### 裴波那契数列

> 题目描述

大家都知道斐波那契数列，现在要求输入一个整数 n，请你输出斐波那契数列的第 n 项（从 0 开始，第 0 项为 0）。n<=39

[网址](https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3?tpId=13&tqId=11160&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int Fibonacci(int n) {
        vector<int> nums;
        if(n < 0){
            return -1;
        }
        if(n == 0){
            return 0;
        }
        if(n == 1 || n == 2){
            return 1;
        }
        nums.push_back(0);
        nums.push_back(1);
        nums.push_back(1);
        for(int i = 3;i <= n;i++){
            nums.push_back(nums[i - 1] + nums[i - 2]);
        }
        return nums[n];
    }
};
```

### 跳台阶

> 题目描述

一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

[网址](https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4?tpId=13&tqId=11161&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int jumpFloor(int number) {
        if(number <= 0){
            return 0;
        }
        if(number <= 2){
            return number;
        }
        return jumpFloor(number - 1) + jumpFloor(number - 2);
    }
};
```

### 变态跳台阶

> 题目描述

一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级……它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

[网址](https://www.nowcoder.com/practice/22243d016f6b47f2a6928b4313c85387?tpId=13&tqId=11162&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
/**
 *  要想跳到第n级台阶，就可以从第n-1级、第n-2级、***、第1级 跳到第n级，再加上直接从地面到第n级的一种情况。
 * F(n-1) = F(n-2) + F(n-3) + ......+ F(1) + F(0);
 * F(n) = F(n-1) + F(n-2) + ......+ F(1) + F(0);
 * F(n) = 2 * F(n-1)
 */
class Solution {
public:
    int jumpFloorII(int number) {
        if(number < 3){
            return number;
        }
        vector<int> vec(number + 1);
        vec[0] = 0;
        vec[1] = 1;
        vec[2] = 2;
        for(int i = 3;i <= number;i++){
            vec[i] = 2 * vec[i - 1];
        }
        return  vec[number];
    }
};
```

### 矩形覆盖

> 题目描述

我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2\*n 的大矩形，总共有多少种方法？

[网址](https://www.nowcoder.com/practice/72a5a919508a4251859fb2cfb987a0e6?tpId=13&tqId=11163&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

[参考](https://www.cnblogs.com/Lune-Qiu/p/9094317.html)

```cpp
class Solution {
public:
    int rectCover(int number) {
        if(number <= 0){
            return 0;
        }
        if(number <= 2){
            return number;
        }
        return rectCover(number - 1) + rectCover(number - 2);
    }
};
```

## 位运算

### 二进制中 1 的个数

> 题目描述

输入一个整数，输出该数二进制表示中 1 的个数。其中负数用补码表示。

[网址](https://www.nowcoder.com/practice/8ee967e43c2c4ec193b040ea7fbb10b8?tpId=13&tqId=11164&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
     int  NumberOf1(int n) {
         unsigned int temp = (unsigned)n;
         int cnt = 0;
         while(temp!=0){
            if((temp & 1) == 1){
                cnt++;
            }
            temp = temp >> 1;
         }
         return cnt;
     }
};
```

## 外码完整性

### 调整数组顺序使奇数位于偶数前面

> 题目描述

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

```cpp
class Solution {
public:
    void reOrderArray(vector<int> &array) {
        int len = array.size();
        if(len == 0 || len == 1){
            return;
        }
        vector<int> vec;
        for(int i = 0;i < len;i++){
            if((array[i] & 1)== 1){
                vec.push_back(array[i]);
            }
        }
        for(int j = 0;j < len;j++){
            if((array[j] & 1)== 0){
                vec.push_back(array[j]);
            }
        }
        for(int k = 0; k < len;k++){
            array[k] = vec[k];
        }
    }
};
```

## 代码的鲁棒性

### 链表中倒数第 k 个结点

> 题目描述

输入一个链表，输出该链表中倒数第 k 个结点。

[网址](https://www.nowcoder.com/practice/529d3ae5a407492994ad2a246518148a?tpId=13&tqId=11167&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {
        if(pListHead == NULL || k == 0){
            return NULL;
        }

        // slow指针指向当前节点，fast指针指向当前节点后的第k个节点
        // 当fast指向链表尾部的下一个时，slow指向目标节点
        ListNode *fast = pListHead;
        ListNode *slow = pListHead;
        for(int count = 0;count < k;count++){
            if(fast == NULL){
                return  NULL;
            }else{
                fast = fast->next;
            }
        }
        while(fast == NULL){
            fast = fast->next;
            slow = slow->next;
        }
        return slow;
    }
};
```

### 反转链表

> 题目描述

输入一个链表，反转链表后，输出新链表的表头。

```
class Solution {
public:
    ListNode* ReverseList(ListNode* pHead) {
        // 空链表或只有一个节点
        if(pHead == NULL || pHead->next == NULL){
            return pHead;
        }

        // 临时头结点
        ListNode *newHead = NULL;

        ListNode  *work = pHead;
        ListNode *next = pHead->next;
        while(work){
            next = work->next;
            work->next = newHead;
            newHead = work;
            work  = next;

        }
        return  newHead;
    }
};
```

### 合并两个有序链表

> 题目描述

输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

```cpp
struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};
class Solution {
public:
    ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
    {
        if(!pHead1 ||!pHead2){
            return pHead1?pHead1:pHead2;
        }
        ListNode *pos1 = pHead1;
        ListNode *pos2 = pHead2;
        // 头结点
        ListNode *head = new ListNode(0);
        // 合并链表的工作指针
        ListNode *pos3 = head;
        while(pos1 && pos2){
            if(pos1->val < pos2->val){
                pos3->next = pos1;
                pos1 = pos1->next;
                pos3 = pos3->next;
                pos3->next = NULL;
            }else{
                pos3->next = pos2;
                pos2 = pos2->next;
                pos3 = pos3->next;
                pos3->next = NULL;
            }
        }

        if(pos1){
            pos3->next = pos1;
        }
        if(pos2){
            pos3->next = pos2;
        }
        return head->next;

    }
};
```

### 树的子结构

> 题目描述

输入两棵二叉树 A，B，判断 B 是不是 A 的子结构。（ps：我们约定空树不是任意一个树的子结构）

[网址](https://www.nowcoder.com/practice/6e196c44c7004d15b1610b9afca8bd88?tpId=13&tqId=11170&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
private:
    bool match(TreeNode *root1,TreeNode *root2){
        // root2遍历完
        if(root2 == NULL){
            return true;
        }
        // root1遍历完，root2没遍历完
        if(root1 == NULL){
            return false;
        }
        // 值不相等
        if(root1->val != root2->val){
            return false;
        }
        // 左右子树
        return match(root1->left,root2->left) &&
            match(root1->right,root2->right);
    }
public:
    bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
    {
        bool flag = false;
        if(pRoot1 && pRoot2){
            flag = match(pRoot1,pRoot2);
            if(!flag){
                //根节点不匹配，则继续匹配左子节点
                flag = match(pRoot1->left,pRoot2);
            }
            if(!flag){
                //根节点不匹配，则继续匹配右子节点
                flag = match(pRoot1->right,pRoot2);
            }
        }
        return flag;
    }
};
```

## 面试思路

### 二叉树的镜像

> 题目描述

操作给定的二叉树，将其变换为源二叉树的镜像

[网址](https://www.nowcoder.com/practice/564f4c26aa584921bc75623e48ca3011?tpId=13&tqId=11171&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    void Mirror(TreeNode *pRoot) {
        if(!pRoot){
            return;
        }
        this->Mirror(pRoot->left);
        this->Mirror(pRoot->right);
        TreeNode *temp = pRoot->left;
        pRoot->left = pRoot->right;
        pRoot->right =  temp;
    }
};
```

## 抽象形象化

### 顺时针打印矩阵

> 题目描述

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，例如，如果输入如下 4 X 4 矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 则依次打印出数字 1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.

```cpp
class Solution {
public:
    vector<int> printMatrix(vector<vector<int> > matrix) {
        vector<int> ret;
        // 矩阵的上边界和左边界
        int rowBegin = 0,colBegin = 0;
        // 矩阵的下边界和右边界
        int rowEnd = matrix.size() - 1,colEnd = matrix[0].size() - 1;

        while(rowBegin <= rowEnd && colBegin <= colEnd){
            // 向右扫描
            for(int j = colBegin;j <= colEnd;j++){
                ret.push_back(matrix[rowBegin][j]);
            }
            // 向下扫描
            for(int i = rowBegin + 1;i <= rowEnd;i++){
                ret.push_back(matrix[i][colEnd]);
            }
            // 向左扫描，如果只有一行则不用执行，故需判断
            if(rowBegin!=rowEnd){
                for(int j = colEnd - 1;j >= colBegin;j--){
                    ret.push_back(matrix[rowEnd][j]);
                }
            }
            // 向上扫描，排除只有一列的情况
            if(colBegin!=colEnd){
                for(int i = rowEnd - 1;i >= rowBegin + 1;i--){
                    ret.push_back(matrix[i][colBegin]);
                }
            }

            // 缩小边界
            rowBegin++;
            rowEnd--;
            colBegin++;
            colEnd--;
        }

        return ret;
    }
};
```

## 抽象具体化

### 包含 min 函数的栈

> 题目描述

定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数（时间复杂度应为 O（1））。

[网址](https://www.nowcoder.com/practice/4c776177d2c04c2494f2555c9fcc1e49?tpId=13&tqId=11173&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    void push(int value) {
        vec1.push_back(value);
        // 栈空
        if(s_top == -1){
            s_top++;
            vec2.push_back(value);
            return;
        }
        // 栈不空
         s_top++;
        int top_min = vec2.back();
        if(top_min < value){
            vec2.push_back(top_min);
        }else{
            vec2.push_back(value);
        }
    }
    void pop() {
        if(s_top!=-1){
            vec1.pop_back();
            vec2.pop_back();
            s_top--;
        }
    }
    int top() {
        if(s_top!=-1){
            return vec1.back();
        }
        return 0;
    }
    int min() {
        if(s_top!=-1){
            return vec2.back();
        }
        return 0;
    }
private:
    // 存放实际的入栈元素
    vector<int> vec1;
    // 存放当前的最小元素
    vector<int> vec2;
    // 栈顶
    int s_top = -1;
};
```

### 栈的压入、弹出序列

> 题目描述

输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列 1,2,3,4,5 是某栈的压入顺序，序列 4,5,3,2,1 是该压栈序列对应的一个弹出序列，但 4,3,5,1,2 就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）

[网址](https://www.nowcoder.com/practice/d77d11405cc7470d82554cb392585106?tpId=13&tqId=11174&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
/**
 * 设计一个辅助栈和一个指针，指针默认指向序列第一个元素，先入栈，每次入栈
 * 后判断栈顶是否与序列的第一个元素相等，不相等继续入栈，相
 * 等时弹出栈顶，指针后移一位，然后继续判断此时栈顶是否与指
 * 针指向的值相等，如果相等继续弹出栈顶，一直比较到栈为空，进
 * 行下一轮入栈或者是循环结束，最后如果辅助栈为空，说明该序列
 * 可能是该压栈序列的弹出序列。
 */
class Solution {
public:
    bool IsPopOrder(vector<int> pushV,vector<int> popV) {
        // 辅助栈
        vector<int> vec;
        int lenPushV = pushV.size();
        int lenPopV = popV.size();
        if(lenPushV == 0 || lenPopV == 0 || lenPopV!=lenPopV){
            return false;
        }

        // 从弹出序列的第一个元素开始
        int popIndex = 0;
        for(int i = 0;i < lenPopV;i++){
            vec.push_back(pushV[i]);
            while(vec.size()!=0 && vec.back() == popV[popIndex]){
                vec.pop_back();
                popIndex++;
            }
        }
        return vec.empty();
    }
};
```

### 从上往下打印二叉树

> 题目描述

从上往下打印出二叉树的每个节点，同层节点从左至右打印。

[网址](https://www.nowcoder.com/practice/7fe2212963db4790b57431d9ed259701?tpId=13&tqId=11175&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};
class Solution {
public:
    vector<int> PrintFromTopToBottom(TreeNode* root) {
        vector<int> vec;
        // 辅助队列
        queue<TreeNode*> qu;
        if(!root){
            return vec;
        }
        TreeNode *work;
        qu.push(root);
        while(qu.size()){
            TreeNode *work = qu.front();
            qu.pop();
            vec.push_back(work->val);
            if(work->left){
                qu.push(work->left);
            }
            if(work->right){
                qu.push(work->right);
            }
        }

        return vec;
    }
};
```

### 二叉搜索树的后序遍历序列

> 题目描述

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出 Yes,否则输出 No。假设输入的数组的任意两个数字都互不相同。

[网址](https://www.nowcoder.com/practice/a861533d45854474ac791d90e447bafd?tpId=13&tqId=11176&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    bool VerifySquenceOfBST(vector<int> sequence) {
        int len =  sequence.size();
        if(len == 0){
            return false;
        }
        if(len == 1){
            return true;
        }
        return judge(sequence,0,len - 1);
    }

    // 序列的起点和终点
    bool judge(vector<int> sequence,int start,int end) {
        if(start >= end){
            return true;
        }
        // 后续序列的根节点在序列最后
        int root = sequence[end];
        // 寻找左子树序列，i最终指向右子树序列的第一个节点
        int i = 0;
        for(;i < end && sequence[i] < root;i++){}
        //  如果右子树中出现了比根节点小的元素则不构成二叉树
        for(int j = i;j < end;j++){
            if(sequence[j] < root){
                return false;
            }
        }
        return  judge(sequence,start,i - 1) &&
            judge(sequence,i,end - 1);
    }
};
```

### 二叉树中和为某一值的路径

> 题目描述

输入一颗二叉树的跟节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。(注意: 在返回值的 list 中，数组长度大的数组靠前)

[网址](https://www.nowcoder.com/practice/b736e784e3e34731af99065031301bca?tpId=13&tqId=11177&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
//非递归后续遍历方式
class Solution {
public:
    vector<vector<int> > FindPath(TreeNode* root,int expectNumber) {
        // 用于返回
        vector<vector<int> > retVec;
        if(!root){
            return retVec;
        }
        // 辅助栈
        vector<TreeNode*> tempVec;
        // 保存一条路径
        vector<int> aPathVec;
        // 临时路径和
        int tempPathSum = 0;
        TreeNode *work = root;
        // 记录上一次访问的节点
        TreeNode *lastVisited = NULL;
        while(work || tempVec.size() > 0){
            if(work){
                // 增加路径和并将当前节点值添加到一条路径中
                tempPathSum += work->val;
                aPathVec.push_back(work->val);
                tempVec.push_back(work);
                // 满足路径和的叶子节点，一条路径完毕，添加到结果中
                if(tempPathSum == expectNumber && !work->left && !work->right){
                    retVec.push_back(aPathVec);
                }
                work = work->left;
            }else{
                work = tempVec.back();
                // 右节点存在且未访问过
                if(work->right && lastVisited!=work->right){
                    work = work->right;
                }
                // 如果右节点不存在或访问过
                else{
                    tempVec.pop_back();
                    aPathVec.pop_back();
                    tempPathSum -= work->val;
                    lastVisited = work;
                    work = NULL;
                }
            }
        }
        return retVec;
    }
};
```

## 分解

### 复杂链表的复制

> 题目描述

输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针指向任意一个节点），返回结果为复制后复杂链表的 head。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）

[网址](https://www.nowcoder.com/practice/f836b2c43afc4b35ad6adc41ec941dba?tpId=13&tqId=11178&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct RandomListNode {
    int label;
    struct RandomListNode *next, *random;
    RandomListNode(int x) :
            label(x), next(NULL), random(NULL) {
    }
};
/**
 * 1.遍历链表，复制链表中的每个结点，并将复制的结点插入到该结点的后面
 * 2.如果原结点的random指针指向的是结点B，那么将复制结点的random指针指向结点B的复制结点B'。
 * 3.将结点和复制结点分为两个链表
 */
class Solution {
public:
    RandomListNode* Clone(RandomListNode* pHead)
    {

        if(!pHead){
            return NULL;
        }
        // 将复制节点接在原始节点的后面
        RandomListNode *work = pHead;
        while(work){
            RandomListNode *copyNode = new RandomListNode(work->label);
            copyNode->next = work->next;
            work->next = copyNode;
            work = copyNode->next;
        }

        // 为 random指针赋值
        work = pHead;
        // 指向赋值节点
        RandomListNode *copyWork = NULL;
        while(work){
            copyWork = work->next;
            if(work->random){
                copyWork->random  = work->random->next;
            }
            work = copyWork->next;
        }

        // 分离原始链表和复制链表
        work = pHead;
        copyWork  = work->next;
        // 复制链表的头部
        RandomListNode *copyHead = pHead->next;
        while(work){
            work->next = copyWork->next;
            work = work->next;
            if(work){
                copyWork->next = work->next;
            }else{
                copyWork->next = NULL;
            }
            copyWork = copyWork->next;
        }
        return copyHead;
    }
};
```

### 字符串的排列

> 题目描述

输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串 abc,则打印出由字符 a,b,c 所能排列出来的所有字符串 abc,acb,bac,bca,cab 和 cba。

> 输入 描述

输入一个字符串,长度不超过 9(可能有字符重复),字符只包括大小写字母。

[网址](https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=13&tqId=11180&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    vector<string> Permutation(string str) {
        vector<string> vec;
        if(str.empty()){
            return vec;
        }
        // 访问标记
        for(int i = 0;i < 9;i++){
            visited[i] = false;
        }
        // 按字典排序
        sort(str.begin(),str.end());
        // 从第一个字符处开始填充
        generate(0,str);
        vec.assign(res_set.begin(),res_set.end());
        return vec;
    }
    void generate(int index,string str){
        // 到达边界则表示完成一次排列
        if(index == str.length()){
            res_set.insert(sequence);
            return;
        }
        // 否则如果当前字符没有访问过则访问当前字符并加入一个排列，然后继续填充
        // 后面的字符
        else{
            for(int j = 0;j < str.length();j++){
                if(visited[j] ==  false){
                    // 设置访问标记
                    visited[j] = true;
                    sequence.push_back(str[j]);
                    // 下一个位置
                    generate(index + 1,str);
                    // 还原
                    visited[j] = false;
                    sequence.pop_back();
                }
            }
        }
    }
private:
    bool visited[9];
    // 为了去重
    set<string> res_set;
    string sequence = "";
};
```

### 二叉搜索树与双向链表

> 题目描述

输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。

[网址](https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5?tpId=13&tqId=11179&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};
class Solution {
public:
    TreeNode* Convert(TreeNode* pRootOfTree)
    {
        if(!pRootOfTree){
            return pRootOfTree;
        }

        // 返回链表的头指针和尾指针
        TreeNode *retHead = NULL;
        TreeNode *retTail = NULL;
        // 辅助栈
        vector<TreeNode*> vec;
        TreeNode *work = pRootOfTree;
        while(work || vec.size()){
            // 到最左端节点
            if(work){
                vec.push_back(work);
                work = work->left;
            }else{
                work = vec.back();
                TreeNode *temp = work;
                vec.pop_back();
                work = work->right;
                if(retHead == NULL){
                    retHead = retTail = temp;
                }else{
                    retTail->right = temp;
                    temp->left = retTail;
                    retTail = retTail->right;
                }
            }
        }
        return retHead;
    }
};
```

## 时间效率

### 数组中出现次数超过一半的数字

> 题目描述

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为 9 的数组{1,2,3,2,2,2,5,4,2}。由于数字 2 在数组中出现了 5 次，超过数组长度的一半，因此输出 2。如果不存在则输出 0。

[网址](https://www.nowcoder.com/practice/e8a1b01a2df14cb2b228b30ee6a92163?tpId=13&tqId=11181&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int MoreThanHalfNum_Solution(vector<int> numbers) {
        int len  = numbers.size();
        if(len == 0){
            return 0;
        }

        // 记录次数最多的数的索引，初始时为第一个元素
        int targetIndex = 0;
        // 出现的次数，出现一次加1，不出现一次则减1
        int cnt = 1;
        for(int i = 1;i < len;i++){
            if(numbers[i] == numbers[targetIndex]){
                cnt++;
            }else{
                // 选定当前数的索引为目标索引
                if(cnt == 0){
                    targetIndex = i;
                    cnt = 1;
                }
                // 抵消一次
                else{
                    cnt--;
                }
            }
        }
        // 验证出现次数最多的数是否超过了一半
        cnt = 0;
        for(int i = 0;i < len;i++){
            if(numbers[targetIndex] == numbers[i]){
                cnt++;
            }
        }
        if(cnt * 2 > len){
            return numbers[targetIndex];
        }else{
            return 0;
        }
    }
};
```

### 连续子数组的最大和

> 题目描述

HZ 偶尔会拿些专业问题来忽悠那些非计算机专业的同学。今天测试组开完会后,他又发话了:在古老的一维模式识别中,常常需要计算连续子向量的最大和,当向量全为正数的时候,问题很好解决。但是,如果向量中包含负数,是否应该包含某个负数,并期望旁边的正数会弥补它呢？例如:{6,-3,-2,7,-15,1,2,2},连续子向量的最大和为 8(从第 0 个开始,到第 3 个为止)。给一个数组，返回它的最大连续子序列的和，你会不会被他忽悠住？(子向量的长度至少是 1)

[网址](https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484?tpId=13&tqId=11183&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int FindGreatestSumOfSubArray(vector<int> array) {
        int len = array.size();
        if(len == 0){
            return 0;
        }
        vector<int> dp;
        for(int i = 0;i < len;i++){
            if(i == 0){
                dp.push_back(array[0]);
            }else{
                int temp = dp[i - 1] + array[i];
                if(temp > array[i]){
                    dp.push_back(temp);
                }else{
                    dp.push_back(array[i]);
                }
            }
        }
        int max = 0;
        for(int i = 1;i < len;i++){
            if(dp[i] > dp[max]){
                max = i;
            }
        }
        return dp[max];
    }
};
```

### 整数中 1 出现的次数（从 1 到 n 整数中 1 出现的次数）

> 题目描述

求出 1-13 的整数中 1 出现的次数,并算出 100-1300 的整数中 1 出现的次数？为此他特别数了一下 1~13 中包含 1 的数字有 1、10、11、12、13 因此共出现 6 次,但是对于后面问题他就没辙了。ACMer 希望你们帮帮他,并把问题更加普遍化,可以很快的求出任意非负整数区间中 1 出现的次数（从 1 到 n 中 1 出现的次数）。

[网址](https://www.nowcoder.com/practice/bd7f978302044eee894445e244c7eee6?tpId=13&tqId=11184&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

[参考](https://www.nowcoder.com/questionTerminal/bd7f978302044eee894445e244c7eee6?f=discussion)

```cpp
class Solution {
public:
    /**
    *
    * 1. 如果第i位（自右至左，从1开始标号）上的数字为0，则第i位可能出现1的次数由更高位决定（若没有高位，视高位为0），等于更高位数字X当前位数的权重10i-1。

    * 2. 如果第i位上的数字为1，则第i位上可能出现1的次数不仅受更高位影响，还受低位影响（若没有低位，视低位为0），等于更高位数字X当前位数的权重10i-1+（低位数字+1）。

    * 3. 如果第i位上的数字大于1，则第i位上可能出现1的次数仅由更高位决定（若没有高位，视高位为0），等于（更高位数字+1）X当前位数的权重10i-1。
    */
    int NumberOf1Between1AndN_Solution(int n)
    {
        // 1的个数
        int count = 0;
        //当前位之前的数，当前位的数，当前位之后的数
        int  before = 0,current = 0,after = 0;
        // 当前位从个位开始
        int i = 1;
        while((n / i)!=0){
            before = n / (i * 10);
            current = (n / i) % 10;
            after = n - (n / i) * i;
            if(current == 0){
                count += before * i;
            }else if(current == 1){
                count += before * i + after + 1;
            }else {
                count += (before + 1) * i;
            }
            i = i * 10;
        }
        return count;
    }
};
```

### 把数组排成最小的数

> 题目描述

输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为 321323。

[网址](https://www.nowcoder.com/practice/8fecd3f8ba334add803bf2a06af1b993?tpId=13&tqId=11185&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
private:
    static bool compare(int x,int y){
        string x_str = to_string(x);
        string y_str = to_string(y);
        return x_str + y_str < y_str + x_str;
    }
public:
    string PrintMinNumber(vector<int> numbers) {
        int len = numbers.size();
        if(len < 1){
            return "";
        }
        string retStr = "";
        sort(numbers.begin(),numbers.end(),Solution::compare);
        for(int i = 0;i < len;i++){
            retStr += to_string(numbers[i]);
        }
        return retStr;
    }
};
```

## 时空效率的平衡

### 丑数

> 题目描述

把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。例如 6、8 都是丑数，但 14 不是，因为它包含质因子 7。 习惯上我们把 1 当做是第一个丑数。求按从小到大的顺序的第 N 个丑数。

```cpp
class Solution {
public:
    int minimum(int a,int b,int c){
        int temp = a < b ? a : b;
        return temp < c ? temp : c;
    }
    int GetUglyNumber_Solution(int index) {
        if(index <= 6){
            return index;
        }
        vector<int> tempVec;
        tempVec.push_back(1);
        int min_num = 1;
        int times2 = 0,times3 = 0,times5 = 0;
        while(tempVec.size() < index){
            min_num = minimum(tempVec[times2] * 2,
                tempVec[times3] * 3,tempVec[times5] * 5);
            if(min_num == tempVec[times2] * 2){
                times2++;
            }
            if(min_num == tempVec[times3] * 3){
                times3++;
            }
            if(min_num == tempVec[times5] * 5){
                times5++;
            }
            tempVec.push_back(min_num);
        }
        return min_num;
    }
};
```

### 第一个只出现一次的字符

> 题目描述

在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.

[网址](https://www.nowcoder.com/practice/1c82e8cf713b4bbeb2a5b31cf5b0417c?tpId=13&tqId=11187&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```c++
class Solution {
public:
    int FirstNotRepeatingChar(string str) {
        map<char,int> char_map;
        int len = str.length();
        for(int i = 0;i < len;i++){
            map<char,int>::iterator it = char_map.find(str[i]);
            if(it == char_map.end()){
                char_map.insert(pair<char,int>(str[i],1));
            }else{
                (it->second)++;
            }
        }
        for(int i = 0;i < len;i++){
            map<char,int>::iterator it  = char_map.find(str[i]);
            if(it->second == 1){
                return i;
            }
        }
        return -1;
    }
};
```

### 两个链表的第一个公共结点

> 题目描述

输入两个链表，找出它们的第一个公共结点。

[网址 ](https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46?tpId=13&tqId=11189&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    ListNode* FindFirstCommonNode( ListNode* pHead1, ListNode* pHead2) {
        if(!pHead1 || !pHead2){
            return NULL;
        }
        int len1 = 0,len2 = 0,distance = 0;
        ListNode *work1 = pHead1;
        ListNode *work2 = pHead2;
        // 两个链表的长度
        while(work1){
            len1++;
            work1 = work1->next;
        }
        while(work2){
            len2++;
            work2 = work2->next;
        }
        // work1始终指向较长链表
        if(len1 > len2){
            work1 = pHead1;
            work2 = pHead2;
            distance = len1 - len2;
        }else{
            work1 = pHead2;
            work2 = pHead1;
            distance = len2 - len1;
        }

        // 较长的链表往后移动若干节点使得两个链表剩余长度相等
        for(int i = 0;i < distance;i++){
            work1 = work1->next;
        }

        while(work1 && work2){
            // 公共节点
            if(work1 == work2){
                return work1;
            }else{
                work1 = work1->next;
                work2 = work2->next;
            }
        }
        return NULL;
    }
};
```

### 数组中的逆序对

> 题目描述

在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数 P。并将 P 对 1000000007 取模的结果输出。 即输出 P%1000000007

[网址]()

[思路参考](https://zhuanlan.zhihu.com/p/92576528)

```cpp
class Solution {
private:
    int cnt = 0;
public:
    int InversePairs(vector<int> data) {
        int len = data.size();
        if(len > 0){
            divide(data,0,len - 1);
        }
        return cnt;
    }
    void merge(vector<int> &vec,int start,int mid,int end){
        // 辅助数组
        vector<int> temp(end - start + 1);
        // i用于遍历第一段,j用于遍历第二段
        int i = start,j = mid + 1,k = 0;
        while(i <= mid && j <= end){
            // 第一段比第二段当前元素小，则较小者入栈
            if(vec[i] <= vec[j]){
                temp[k++] = vec[i++];
            }
            // 第一段当前元素比第二段当前元素大，则构成逆序，由于右侧数组
            // 是有序的，所以逆序数等于当前第二段j之前的元素个数
            else{
                temp[k++] = vec[j++];
                cnt += (mid - i + 1);
                cnt %= 1000000007;
            }
        }

        // 将两端序列剩余的元素添加到临时数组
        while(i <= mid){
            temp[k++] = vec[i++];
        }
        while(j <= end){
            temp[k++] = vec[j++];
        }
        // temp中的元素有序，将其复制到原数组
        for(int k = 0;k < end - start + 1;k++){
            vec[start + k] = temp[k];
        }
    }
    void divide(vector<int> &vec,int start,int end){
        if(start < end){
            int mid = (start + end) >> 1;
            divide(vec,start,mid);
            divide(vec,mid + 1,end);
            merge(vec,start,mid,end);
        }
    }
};
```

## 知识迁移能力

### 数字在排序数组中出现的次数

> 题目描述

统计一个数字在排序数组中出现的次数。

[网址](https://www.nowcoder.com/practice/70610bf967994b22bb1c26f9ae901fa2?tpId=13&tqId=11190&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int GetNumberOfK(vector<int> data ,int k) {
        int len = data.size();
        if(len == 0){
            return 0;
        }
        // 第一个目标元素
        int start = 0,end = len - 1;
        for(int i = 0;i < len;i++){
            if(data[i] == k){
                start = i;
                break;
            }
        }
        // 最后一个目标元素
        for(int j = start;j < len;j++){
            if(data[j]!=k){
                end = j - 1;
                break;
            }
        }
        return end - start + 1;
    }
};
```

### 二叉树的深度

> 题目描述

输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

[网址](https://www.nowcoder.com/practice/435fb86331474282a3499955f0a41e8b?tpId=13&tqId=11191&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};
class Solution {
public:
    int TreeDepth(TreeNode* pRoot)
    {
        if(!pRoot){
            return 0;
        }
        // 深度
        int depth = 1;
        // 用于判断一层是否遍历完
        int count = 0;
        // 辅助队列
        queue<TreeNode*> qu;
        TreeNode *work = pRoot;
        qu.push(work);
        while(qu.size() > 0){
            work = qu.front();
            qu.pop();
            count--;
            if(work->left){
                qu.push(work->left);
            }
            if(work->right){
                qu.push(work->right);
            }
            // 一层遍历完则深度加1
            if(count == 0){
                depth++;
                count = qu.size();
            }
        }
        return depth;
    }
};
```

### 平衡二叉树

> 题目描述

输入一棵二叉树，判断该二叉树是否是平衡二叉树。

[网址](https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222?tpId=13&tqId=11192&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    bool IsBalanced_Solution(TreeNode* pRoot) {
        if(!pRoot){
            return true;
        }
        int leftHeight = getHeight(pRoot->left);
        int rightHeight = getHeight(pRoot->right);
        return abs(leftHeight - rightHeight) <= 1 &&
            IsBalanced_Solution(pRoot->left) && IsBalanced_Solution(pRoot->right);
    }
    int getHeight(TreeNode *root){
        if(!root){
            return 0;
        }
        int leftHeight = getHeight(root->left);
        int rightHeight = getHeight(root->right);
        return max(leftHeight,rightHeight) + 1;
    }
};
```

### 数组中只出现 一次的数字

> 题目描述

一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

[网址](https://www.nowcoder.com/practice/e02fdb54d7524710a7d664d082bb7811?tpId=13&tqId=11193&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
/**
 * 从前向后依次异或数组中的数字，那么得到的结果是两个只
 * 出现一次的数字的异或结果，其他成对出现的数字被抵消了。
 * 由于这两个数字不同，所以异或结果肯定不为0，也就是这个
 * 异或结果一定至少有一位是1，我们在结果中找到第一个为1
 * 的位的位置，记为第n位。接下来，以第n位是不是1为标准，
 * 将数组分为两个子数组，第一个数组中第n位都是1，第二个数
 * 组中第n位都是0。这样，便实现了我们的目标。最后，两个子数
 * 组分别异或则可以找到只出现一次的数字。
 */
class Solution {
public:
    void FindNumsAppearOnce(vector<int> data,int* num1,int *num2) {
        int temp = 0;
        int len = data.size();
        for(int i = 0;i < len;i++){
            temp ^= data[i];
        }

        // 寻找第一个1所在的索引
        int j;
        for(j = 0;j < 32;j++){
            if(((temp >> j) & 1) == 1){
                break;
            }
        }

        int temp_1 = 0;
        int temp_0 = 0;
        // 第j位为1的和为0的所有数字分开
        for(int k = 0;k < len;k++){
            if(((data[k] >> j) & 1) == 1){
                temp_0 ^= data[k];
            }else{
                temp_1 ^= data[k];
            }
        }
        *num1 = temp_0;
        *num2 = temp_1;
        return;
    }
};
```

### 和为 S 的连续正数序列

> 题目描述

小明很喜欢数学,有一天他在做数学作业时,要求计算出 9~16 的和,他马上就写出了正确答案是 100。但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为 100(至少包括两个数)。没多久,他就得到另一组连续正数和为 100 的序列:18,19,20,21,22。现在把问题交给你,你能不能也很快的找出所有和为 S 的连续正数序列? Good Luck!

输出: 输出所有和为 S 的连续正数序列。序列内按照从小至大的顺序，序列间按照开始数字从小到大的顺序

[网址](https://www.nowcoder.com/practice/c451a3fd84b64cb19485dad758a55ebe?tpId=13&tqId=11194&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    vector<vector<int> > FindContinuousSequence(int sum) {
        vector<vector<int> > ret;
        if(sum <= 2){
            return ret;
        }
        int mid = (1 + sum) >> 1;
        int left = 1,right = 2,temp_sum = 3;
        while(left < mid && right < mid + 1){
            while(temp_sum > sum){
                temp_sum -= left;
                left++;
            }
            if(temp_sum == sum){
                vector<int> vec;
                for(int i = left;i <= right;i++){
                    vec.push_back(i);
                }
                ret.push_back(vec);
            }
            ++right;
            temp_sum += right;
        }
        return ret;
    }
};
```

### 和为 S 的两个数字

> 题目描述

输入一个递增排序的数组和一个数字 S，在数组中查找两个数，使得他们的和正好是 S，如果有多对数字的和等于 S，输出两个数的乘积最小的。

输出: 对应每个测试案例，输出两个数，小的先输出。

[网址](https://www.nowcoder.com/practice/390da4f7a00f44bea7c2f3d19491311b?tpId=13&tqId=11195&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    vector<int> FindNumbersWithSum(vector<int> array,int sum) {
        vector<int> ret;
        // 输入数组长度
        int len = array.size();
        if(len < 2){
            return ret;
        }

        // 左指针指向第一个元素，右指针指向最右边元素
        int left = 0,right = len - 1;
        // 用于记录当前最小乘积
        int multi = array[len - 1] * array[len - 1];
        while(left < right){
            int temp_sum = array[left] + array[right];
            // 两数和等于目标数，如果乘积比之前最小乘积小则记录
            if(temp_sum == sum){
                int temp = array[left] * array[right];
                if(temp < multi){
                    ret.clear();
                    ret.push_back(array[left]);
                    ret.push_back(array[right]);
                    multi = temp;
                }
                left++;
                right--;
            }else if(temp_sum > sum){
                right--;
            }else{
                left++;
            }
        }
        return ret;
    }
};
```

### 左旋转字符串

> 题目描述

汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列 S，请你把其循环左移 K 位后的序列输出。例如，字符序列 S=”abcXYZdef”,要求输出循环左移 3 位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它!

[网址](https://www.nowcoder.com/practice/12d959b108cb42b1ab72cef4d36af5ec?tpId=13&tqId=11196&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    void rotate(vector<char> &str,int start,int end){
        for(int i = start,j = end;i < j;){
            char temp = str[i];
            str[i] = str[j];
            str[j] = temp;
            i++;
            j--;
        }
    }
    string LeftRotateString(string str, int n) {
        vector<char> vec(str.begin(),str.end());
        int len = str.length();
        if(len == 0){
            return "";
        }
        n %= len;
        rotate(vec,0,len - 1);
        rotate(vec,0,len - n - 1);
        rotate(vec,len - n,len - 1);
        string retStr(vec.begin(),vec.end());
        return retStr;
    }
};
```

### 翻转单词顺序列

> 题目描述

牛客最近来了一个新员工 Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。同事 Cat 对 Fish 写的内容颇感兴趣，有一天他向 Fish 借来翻看，但却读不懂它的意思。例如，“student. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，正确的句子应该是“I am a student.”。Cat 对一一的翻转这些单词顺序可不在行，你能帮助他么？

[网址](https://www.nowcoder.com/practice/3194a4f4cf814f63919d0790578d51f3?tpId=13&tqId=11197&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
private:
    // 反转字符串
    void reverse(string &str,int left,int right){
        while(left < right){
            char temp = str[left];
            str[left] = str[right];
            str[right] = temp;
            left++;
            right--;
        }
    }

public:
    string ReverseSentence(string str) {
        int len = str.length();
        if(len == 0){
            return "";
        }
        // 反转整个句子
        reverse(str,0,len - 1);
        // 记录每个单词的起始和结束索引
        int left = 0,right = 0;
        for(int i = 0;i < len;i++){
            if(str[i] == ' '){
                // 反转每个单词
                reverse(str,left,right);
                left = i + 1;
            }else{
                right = i;
            }
        }
        // 反转最后一个单词
        reverse(str,left,len - 1);
        return str;
    }
};
```

## 抽象建模能力

### 扑克牌顺子

> 题目描述

LL 今天心情特别好,因为他去买了一副扑克牌,发现里面居然有 2 个大王,2 个小王(一副牌原本是 54 张^\_^)...他随机从中抽出了 5 张牌,想测测自己的手气,看看能不能抽到顺子,如果抽到的话,他决定去买体育彩票,嘿嘿！！“红心 A,黑桃 3,小王,大王,方片 5”,“Oh My God!”不是顺子.....LL 不高兴了,他想了想,决定大\小 王可以看成任何数字,并且 A 看作 1,J 为 11,Q 为 12,K 为 13。上面的 5 张牌就可以变成“1,2,3,4,5”(大小王分别看作 2 和 4),“So Lucky!”。LL 决定去买体育彩票啦。 现在,要求你使用这幅牌模拟上面的过程,然后告诉我们 LL 的运气如何， 如果牌能组成顺子就输出 true，否则就输出 false。为了方便起见,你可以认为大小王是 0。

[网址](https://www.nowcoder.com/practice/762836f4d43d43ca9deb273b3de8e1f4?tpId=13&tqId=11198&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
/**
 * 首先应该对数组进行排序。
 * 统计数组中大小王（0）出现的个数。
 * 统计数组中所有相邻数之间的间隔。
 * 同时还需要排除对子的情况，如果出现了对子，那么肯定不可能是顺子（0除外）。
 * 最后比较0的个数和间隔大小，如果0的个数大于等于间隔数，说明可以组成顺
*/
class Solution {
public:
    bool IsContinuous( vector<int> numbers ) {
        if(numbers.size()!=5){
            return false;
        }
        sort(numbers.begin(),numbers.end());

        // count0记录0的个数，gap记录差距
        int count0 = 0,len = numbers.size(),gap = 0;
        for(int i = 0;i < len - 1;i++){
            if(numbers[i] == 0){
                count0++;
                continue;
            }
            // 处理对子
            if(numbers[i] == numbers[i+1]){
                return false;
            }

            // 相邻两个数的差距
            gap += numbers[i + 1] - numbers[i] - 1;

        }
        if(count0 >= gap){
            return true;
        }else{
            return false;
        }
    }
};
```

### 孩子们的游戏(圆圈中最后剩下的数)

> 题目描述

每年六一儿童节,牛客都会准备一些小礼物去看望孤儿院的小朋友,今年亦是如此。HF 作为牛客的资深元老,自然也准备了一些小游戏。其中,有个游戏是这样的:首先,让小朋友们围成一个大圈。然后,他随机指定一个数 m,让编号为 0 的小朋友开始报数。每次喊到 m-1 的那个小朋友要出列唱首歌,然后可以在礼品箱中任意的挑选礼物,并且不再回到圈中,从他的下一个小朋友开始,继续 0...m-1 报数....这样下去....直到剩下最后一个小朋友,可以不用表演,并且拿到牛客名贵的“名侦探柯南”典藏版(名额有限哦!!^\_^)。请你试着想下,哪个小朋友会得到这份礼品呢？(注：小朋友的编号是从 0 到 n-1)

如果没有小朋友，请返回-1

[网址](https://www.nowcoder.com/practice/f78a359491e64a50bce2d89cff857eb6?tpId=13&tqId=11199&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    // n表示小朋友的数量,m是随机指定的数
    int LastRemaining_Solution(int n, int m)
    {
        if(m < 1 || m < 1){
            return -1;
        }

        // 给孩子编号
        list<int> kids;
        for(int i = 0;i < n;i++){
            kids.push_back(i);
        }

        list<int>::iterator it = kids.begin();
        while(kids.size() > 1){
            int cnt = 0;
            while(cnt < m - 1){
                it++;
                if(it == kids.end()){
                    it = kids.begin();
                }
                cnt++;
            }
            // 去掉第m个编号，erase方法返回删除元素的下一个位置
            it = kids.erase(it);
            if(it == kids.end()){
                it = kids.begin();
            }
        }
        return *it;
    }
};
```

## 发散思维能力

### 求 1+2+3+...+n

> 思维能力

求 1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case 等关键字及条件判断语句（A?B:C）。

[网址](https://www.nowcoder.com/practice/7a0da8fc483247ff8800059e12d7caf1?tpId=13&tqId=11200&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int Sum_Solution(int n) {
        int sum = n;
        bool flag = (n > 0) && (sum += Sum_Solution(n - 1));
        return sum;
    }
};
```

## 综合

### 将字符串转化为数字

> 题目描述

将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。 数值为 0 或者字符串不是一个合法的数值则返回 0

[网址](https://www.nowcoder.com/practice/1277c681251b4372bdef344468e4f26e?tpId=13&tqId=11202&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    int StrToInt(string str) {
        // 数字和符号
        long long num = 0,flag = 1;
        int len = str.length();
        if(len == 0){
            return 0;
        }
        for(int i = 0;i < len;i++){
            // 负数
            if(i == 0 && str[i] == '-'){
                flag = -1;
                continue;
            }
            if(i == 0 && str[i] == '+'){
                flag = 1;
                continue;
            }
            if(str[i] - '0' < 0 || str[i] - '0' > 9){
                return 0;
            }else{
                num = (num << 3) + (num << 1) + (str[i] - '0' + 0);
            }
        }
        long long ret = num * flag;
        if(ret > INT_MAX || ret < INT_MIN){
            return 0;
        }else{
            return ret;
        }
    }
};
```

## 回朔法

### 矩阵中的路径

> 题目描述

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。 例如 a b c e s f c s a d e e 矩阵中包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，因为字符串的第一个字符 b 占据了矩阵中的第一行第二个格子之后，路径不能再次进入该格子。

[网址](https://www.nowcoder.com/practice/c61c6999eecb4b8f88a98f66b273a3cc?tpId=13&tqId=11218&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
public:
    /**
     *  i和j是矩阵中的位置，len是目标串的长度，index是目标串的索引
     */
    bool find(char* matrix, int rows, int cols,int i,int j,
        char* str,int len,int index,vector<bool> &flags){
            int k = cols * i + j;
            if(i < 0 || j < 0 || i >= rows ||
                j >= cols || str[index] != matrix[k]
                || flags[k] == true){
                    return false;
                }
            if(index == len - 1){
                return true;
            }
            flags[k] = true;
            // 前
            if(find(matrix,rows,cols,i,j + 1,str,len,index + 1,flags)){
                return true;
            }
            //后
            if(find(matrix,rows,cols,i,j - 1,str,len,index + 1,flags)){
                return true;
            }
            //上
            if(find(matrix,rows,cols,i - 1,j,str,len,index + 1,flags)){
                return true;
            }
            // 下
            if(find(matrix,rows,cols,i + 1,j,str,len,index + 1,flags)){
                return true;
            }
            flags[k] = false;
            return false;
    }

    bool hasPath(char* matrix, int rows, int cols, char* str)
    {
        // 标记数组
        vector<bool> flags(rows * cols);
        for(int i = 0;i < rows;i++){
            int temp = i * cols;
            for(int j = 0;j < cols;j++){
                flags[temp + j] = false;
            }
        }
        // 目标串的长度
        int len = 0;
        for(;str[len] != '\0';len++){

        }
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < cols;j++){
                if(find(matrix,rows,cols,i,j,str,len,0,flags)){
                    return true;
                }
            }
        }
        return false;
    }


};
```

### 机器人的运动范围

> 题目描述

地上有一个 m 行和 n 列的方格。一个机器人从坐标 0,0 的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于 k 的格子。 例如，当 k 为 18 时，机器人能够进入方格（35,37），因为 3+5+3+7 = 18。但是，它不能进入方格（35,38），因为 3+5+3+8 = 19。请问该机器人能够达到多少个格子？

[网址](https://www.nowcoder.com/practice/6e5207314b5241fb83f2329e89fdecc8?tpId=13&tqId=11219&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

```cpp
class Solution {
private:
    // 坐标各位之和
    int bitSum(int i,int j){
        int sum = 0;
        while(i > 0){
            sum += (i % 10);
            i /= 10;
        }
        while(j > 0){
            sum += (j % 10);
            j /= 10;
        }
        return sum;
    }
    int fillIn(int threshold, int rows, int cols,int i,int j,vector<bool> &flags){
        int cnt = 0;
        int k = i * cols + j;
        if(i >= 0 && i < rows && j >= 0 && j < cols &&
            !flags[k] && bitSum(i,j) <= threshold){

            flags[k] = true;
            // 左
            cnt = 1 + fillIn(threshold,rows,cols,i,j - 1,flags)
            // 右
            + fillIn(threshold,rows,cols,i,j + 1,flags)
            // 上
            + fillIn(threshold,rows,cols,i - 1,j,flags)+
            // 下
            + fillIn(threshold,rows,cols,i + 1,j,flags);
        }

        return cnt;
    }
public:

    int movingCount(int threshold, int rows, int cols)
    {
        if(threshold < 0 || cols <= 0 || rows <= 0){
            return 0;
        }
        // 访问标记
        vector<bool> flags(rows * cols);
        for(int i = 0;i < rows;i++){
            int temp = i * cols;
            for(int j = 0;j < cols;j++){
                flags[temp + j] = false;
            }
        }
        // 从(0,0)开始
        return fillIn(threshold,rows,cols,0,0,flags);
    }
};
```

## 动态规划与贪婪

### 剪绳子

> 题目描述

给你一根长度为 n 的绳子，请把绳子剪成整数长的 m 段（m、n 都是整数，n>1 并且 m>1），每段绳子的长度记为 k[0],k[1],...,k[m]。请问 k[0]xk[1]x...xk[m]可能的最大乘积是多少？例如，当绳子的长度是 8 时，我们把它剪成长度分别为 2、3、3 的三段，此时得到的最大乘积是 18。

[网址](https://www.nowcoder.com/practice/57d85990ba5b440ab888fc72b0751bf8?tpId=13&tqId=33257&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

[参考](https://zhuanlan.zhihu.com/p/55247132)

**动态规划方式**

```cpp
class Solution {
public:
    int cutRope(int number) {
        if(number < 2 || number > 60){
            return 0;
        }else{
            // 由于必须要切一刀所以2和3时为1和2
            if(number == 2){
                return 1;
            }
            if(number == 3){
                return 2;
            }
            // vec[i] 表示长度为i的绳子切割若干段后的最大乘积
            vector<int> vec(number + 1);
            // 这里是子问题，即在此之前已经至少切了一刀，故这里不一样
            vec[0] = 0;
            vec[1] = 1;
            vec[2] = 2;
            vec[3] = 3;
            int max = 0;
            for(int i = 4;i <= number;i++){
                for(int j = 1;j <= i / 2;j ++){
                    int temp = vec[j] * vec[i - j];
                    if(temp > max){
                        max = temp;
                        vec[i] = max;
                    }
                }
            }
            return vec[number];
        }
    }
};
```
