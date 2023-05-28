---
title: "计算机考研复试机试题"
date: 2021-03-10T13:35:34+08:00
draft: false
image: "img/skFtz_ycl-I.jpg"
categories: 
tag:
---

## 编程练习

### 约数的个数

输入 n 个整数,依次输出每个数的约数的个数

[网址](https://www.nowcoder.com/practice/04c8a5ea209d41798d23b59f053fa4d6?tpId=40&tqId=21334&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <math.h>
using namespace std;
int main()
{
    int cnt;
    cin >> cnt;
    for (int i = 0; i < cnt; i++)
    {
        int num;
        cin >> num;
        int temp = 0;
        int sq = sqrt(num);
        for (int j = 1; j <= sq; j++)
        {
            if (num % j == 0)
            {
                if (j * j == num)
                {
                    temp += 1;
                }
                else
                {
                    temp += 2;
                }
            }
        }
        cout << temp << endl;
    }
    return 0;
}

```

### 成绩排序

查找和排序

题目：输入任意（用户，成绩）序列，可以获得成绩从高到低或从低到高的排列,相同成绩
都按先录入排列在前的规则处理。

[网址](https://www.nowcoder.com/practice/0383714a1bb749499050d2e0610418b1?tpId=40&tqId=21333&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#define DESC 0
#define ASC 1
using namespace std;
bool comp_desc(pair<string, int> stu1, pair<string, int> stu2)
{
    return stu1.second > stu2.second;
}

bool comp_asc(pair<string, int> stu1, pair<string, int> stu2)
{
    return stu1.second < stu2.second;
}
int main()
{
    vector<pair<string, int> > stu_gra;
    int cnt;
    // 排序方式
    int type;
    cin >> cnt;
    cin >> type;
    for (int i = 0; i < cnt; i++)
    {
        string name;
        int grade;
        cin >> name >> grade;
        stu_gra.push_back(pair<string, int>(name, grade));
    }

    if (type == DESC)
    {
        stable_sort(stu_gra.begin(), stu_gra.end(), comp_desc);
    }
    else
    {
        stable_sort(stu_gra.begin(), stu_gra.end(), comp_asc);
    }
    for (int i = 0; i < cnt; i++)
    {
        cout << stu_gra[i].first << " " << stu_gra[i].second << endl;
    }
    return 0;
}
```

### 反序输出

输入任意 4 个字符(如：abcd)， 并按反序输出(如：dcba)

[网址](https://www.nowcoder.com/practice/171278d170c64d998ab342b3b40171bb?tpId=40&tqId=21336&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;
int main()
{
    string str;
    while (cin >> str)
    {
        reverse(str.begin(), str.end());
        cout << str << endl;
    }
    return 0;
}
```

### 代理服务器

使用代理服务器能够在一定程度上隐藏客户端信息，从而保护用户在互联网上的隐私。我们知道 n 个代理服务器的 IP 地址，现在要用它们去访问 m 个服务器。这 m 个服务器的 IP 地址和访问顺序也已经给出。系统在同一时刻只能使用一个代理服务器，并要求不能用代理服务器去访问和它 IP 地址相同的服务器（不然客户端信息很有可能就会被泄露）。在这样的条件下，找到一种使用代理服务器的方案，使得代理服务器切换的次数尽可能得少。

> 输入

每个测试数据包括 n + m + 2 行。
第 1 行只包含一个整数 n，表示代理服务器的个数。
第 2 行至第 n + 1 行每行是一个字符串，表示代理服务器的 IP 地址。这 n 个 IP 地址两两不相同。
第 n + 2 行只包含一个整数 m，表示要访问的服务器的个数。
第 n + 3 行至第 n + m + 2 行每行是一个字符串，表示要访问的服务器的 IP 地址，按照访问的顺序给出。
每个字符串都是合法的 IP 地址，形式为“xxx.yyy.zzz.www”，其中任何一部分均是 0–255 之间的整数。输入数据的任何一行都不包含空格字符。
其中，1<=n<=1000，1<=m<=5000。

> 输出

可能有多组测试数据，对于每组输入数据， 输出数据只有一行，包含一个整数 s，表示按照要求访问服务器的过程中切换代理服务器的最少次数。第一次使用的代理服务器不计入切换次数中。若没有符合要求的安排方式，则输出-1。

[网址](https://www.nowcoder.com/practice/1284469ee94a4762848816a42281a9e0?tpId=40&tqId=21335&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <vector>
using namespace std;
int main()
{
    int proxyIpCnt;
    int serverIpCnt;
    vector<string> proxyIps;
    vector<string> serverIps;
    // 代理服务器地址
    cin >> proxyIpCnt;
    for (int i = 0; i < proxyIpCnt; i++)
    {
        string temp;
        cin >> temp;
        proxyIps.push_back(temp);
    }
    // 服务器地址
    cin >> serverIpCnt;
    for (int i = 0; i < serverIpCnt; i++)
    {
        string temp;
        cin >> temp;
        serverIps.push_back(temp);
    }

    // start记录每一轮从哪个服务器开始代理，changes记录总的变化次数
    // 当没有符合要求的代理方案时flag为false
    int start = 0, changes = 0, flag = 1, j = 0;
    while (start < serverIpCnt && flag)
    {
        // 每一个代理服务器能够代理的最多的服务器数量
        int cnt = 0;
        for (int i = 0; i < proxyIpCnt; i++)
        {
            // 能够代理的情况
            j = start;
            while (j < serverIpCnt && serverIps[j] != proxyIps[i])
            {
                j++;
            }

            // serverIps[j] == proxyIps[i]，计算该代理服务器能够代理的
            // 服务器的数量
            if (j - start > cnt)
            {
                cnt = j - start;
            }
            if (cnt == 0)
            {
                flag = false;
            }
            changes++;
            // 计算下一个代理服务器代理的起始服务器
            start += cnt;
        }
    }
    if (flag)
    {
        // 计算时将第一次也带进去了，所以需要减去
        cout << changes - 1 << endl;
    }
    else
    {
        // 没有符合的代理方式
        cout << -1 << endl;
    }
    return 0;
}
```

### 手机键盘

按照手机键盘输入字母的方式，计算所花费的时间 如：a,b,c 都在“1”键上，输入 a 只需要按一次，输入 c 需要连续按三次。 如果连续两个字符不在同一个按键上，则可直接按，如：ad 需要按两下，kz 需要按 6 下 如果连续两字符在同一个按键上，则两个按键之间需要等一段时间，如 ac，在按了 a 之后，需要等一会儿才能按 c。 现在假设每按一次需要花费一个时间段，等待时间需要花费两个时间段。 现在给出一串字符，需要计算出它所需要花费的时间。

一个长度不大于 100 的字符串，其中只有手机按键上有的小写字母

一个长度不大于 100 的字符串，其中只有手机按键上有的小写字母

[网址](https://www.nowcoder.com/practice/20082c12f1ec43b29cd27c805cd476cd?tpId=40&tqId=21337&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    // 记录每个字符需要按的次数
    int keys[26] = {1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 4, 1, 2, 3, 1, 2, 3, 4};
    string str;
    cin >> str;
    // 总时间
    int total_time = 0;
    int len = str.length();
    if (len == 0)
    {
        return 0;
    }
    else
    {
        // 第一个字符需要按得次数
        total_time = keys[str[0] - 'a'];
    }
    for (int i = 1; i < len; i++)
    {
        // 按键时间
        total_time += keys[str[i] - 'a'];
        // 等待时间
        if (str[i] - str[i - 1] == keys[str[i] - 'a'] - keys[str[i - 1] - 'a'])
        {
            total_time += 2;
        }
    }
    cout << total_time << endl;
    return 0;
}

```

### 质因素的个数

求正整数 N(N>1)的质因数的个数。 相同的质因数需要重复计算。如 120=2*2*2*3*5，共有 5 个质因数。

可能有多组测试数据，每组测试数据的输入是一个正整数 N，(1 < N < 10 ^ 9)。

对于每组数据，输出 N 的质因数的个数。

[网址](https://www.nowcoder.com/practice/20426b85f7fc4ba8b0844cc04807fbd9?tpId=40&tqId=21338&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <math.h>
using namespace std;
int getPrimeFactor(int num)
{
    int cnt = 0;
    int sq = sqrt(num);
    for (int i = 2; i <= sq; i++)
    {
        if (num % i == 0)
        {
            cnt += (1 + getPrimeFactor(num / i));
            break;
        }
    }
    if (cnt == 0)
    {
        cnt++;
    }
    return cnt;
}
int main()
{
    int num, count = 0;
    while (cin >> num)
    {
        cout << getPrimeFactor(num) << endl;
    }
    return 0;
}
```

### 整数拆分

一个整数总可以拆分为 2 的幂的和，例如： 7=1+2+4 7=1+2+2+2 7=1+1+1+4 7=1+1+1+2+2 7=1+1+1+1+1+2 7=1+1+1+1+1+1+1 总共有六种不同的拆分方式。 再比如：4 可以拆分成：4 = 4，4 = 1 + 1 + 1 + 1，4 = 2 + 2，4=1+1+2。 用 f(n)表示 n 的不同拆分的种数，例如 f(7)=6. 要求编写程序，读入 n(不超过 1000000)，输出 f(n)%1000000000。

每组输入包括一个整数：N(1<=N<=1000000)。

对于每组数据，输出 f(n)%1000000000。

[网址](https://www.nowcoder.com/practice/376537f4609a49d296901db5139639ec?tpId=40&tqId=21339&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

[参考]()https://www.cnblogs.com/multhree/p/10648468.html

```cpp
#include <iostream>
using namespace std;
int main()
{
    int num;
    while (cin >> num)
    {
        int dp[num + 1];
        dp[0] = dp[1] = 1;
        for (int i = 2; i <= num; i++)
        {
            if (i % 2 == 0)
            {
                dp[i] = (dp[i - 1] + dp[i / 2]) % 1000000000;
            }
            else
            {
                dp[i] = dp[i - 1];
            }
        }
        cout << dp[num] << endl;
    }
    return 0;
}
```

### 成绩排序 2

用一维数组存储学号和成绩，然后，按成绩排序输出。

输入第一行包括一个整数 N(1<=N<=100)，代表学生的个数。
接下来的 N 行每行包括两个整数 p 和 q，分别代表每个学生的学号和成绩。

按照学生的成绩从小到大进行排序，并将排序后的学生信息打印出来。
如果学生的成绩相同，则按照学号的大小进行从小到大排序。

[网址](https://www.nowcoder.com/practice/3f27a0a5a59643a8abf0140b9a8cf1f7?tpId=40&tqId=21340&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;
bool comp(pair<int, int> s1, pair<int, int> s2)
{
    return s1.second == s2.second ? s1.first < s2.first :
        s1.second < s2.second;
}
int main()
{
    int cnt = 0;
    cin >> cnt;
    vector<pair<int, int> > vec;
    for (int i = 0; i < cnt; i++)
    {
        int number, grade;
        cin >> number >> grade;
        vec.push_back(pair<int, int>(number, grade));
    }
    stable_sort(vec.begin(), vec.end(), comp);
    for (int i = 0; i < cnt; i++)
    {
        cout << vec[i].first << " " << vec[i].second << endl;
    }
    return 0;
}
```

### 二叉树遍历

编一个程序，读入用户输入的一串先序遍历字符串，根据此字符串建立一个二叉树（以指针方式存储）。 例如如下的先序遍历字符串： ABC##DE#G##F### 其中“#”表示的是空格，空格字符代表空树。建立起此二叉树以后，再对二叉树进行中序遍历，输出遍历结果。

输入包括 1 行字符串，长度不超过 100。

可能有多组测试数据，对于每组数据，
输出将输入字符串建立二叉树后中序遍历的序列，每个字符后面都有一个空格。
每个输出结果占一行。

```cpp
#include <iostream>
#include <string>
#include <vector>
using namespace std;
typedef struct TreeNode
{
    char ch;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(char ch)
    {
        this->ch = ch;
        this->left = NULL;
        this->right = NULL;
    }
} TreeNode;
static int pos = 0;
TreeNode *buildTree(string str)
{
    if (pos >= str.length())
    {
        return NULL;
    }
    if (str[pos] == '#')
    {
        pos++;
        return NULL;
    }
    TreeNode *root = new TreeNode(str[pos]);
    pos++;
    root->left = buildTree(str);
    root->right = buildTree(str);
    return root;
};
void inOrder(TreeNode *root)
{
    if (root == NULL)
    {
        return;
    }
    else
    {
        inOrder(root->left);
        cout << root->ch << " ";
        inOrder(root->right);
    }
}
int main()
{
    string sequence;
    cin >> sequence;
    int len = sequence.length();
    TreeNode *root = buildTree(sequence);
    inOrder(root);
    return 0;
}
```

### 最大最小数

设 a、b、c 均是 0 到 9 之间的数字，abc、bcc 是两个三位数，且有：abc+bcc=532。求满足条件的所有 a、b、c 的值。

[网址](https://www.nowcoder.com/practice/912b15e237ef44148e44018d7b8750b6?tpId=40&tqId=21346&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    for (int a = 1; a <= 5; a++)
    {
        for (int b = 1; b <= 5; b++)
        {
            for (int c = 0; c <= 9; c++)
            {
                if (a * 100 + b * 10 + c + b * 100 + c * 10 + c == 532)
                {
                    cout << a << " " << b << " " << c << endl;
                }
            }
        }
    }
    return 0;
}
```

### 今年的第几天?

输入年、月、日，计算该天是本年的第几天。

[网址](https://www.nowcoder.com/practice/ae7e58fe24b14d1386e13e7d70eaf04d?tpId=40&tqId=21350&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    int dayOfMonth[13] = {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    int year, month, day;
    while (cin >> year >> month >> day)
    {
        int days = 0;
        if (1 <= year && year <= 3000 && month >= 1 && month <= 12 && day >= 1 && day <= 31)
        {
            if ((year % 400 == 0) || (year % 4 == 0 && year % 100 != 0))
            {
                dayOfMonth[2] = 29;
            }
            for (int i = 1; i < month; i++)
            {
                days += dayOfMonth[i];
            }
            days += day;
            cout << days << endl;
            dayOfMonth[2] = 28;
        }
    }
}
```

### 完数和盈数

一个数如果恰好等于它的各因子(该数本身除外)子和，如：6=3+2+1。则称其为“完数”；若因子之和大于该数，则称其为“盈数”。 求出 2 到 60 之间所有“完数”和“盈数”。

[网址](https://www.nowcoder.com/practice/ccc3d1e78014486fb7eed3c50e05c99d?tpId=40&tqId=21351&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int judge(int num)
{
    int sum = 1;
    for (int i = 2; i < num; i++)
    {
        if (num % i == 0)
        {
            sum += i;
        }
    }
    if (sum == num)
    {
        return 0;
    }
    if (sum > num)
    {
        return 1;
    }
    return -1;
}
int main()
{
    vector<int> completion;
    vector<int> surplus;
    for (int i = 2; i <= 60; i++)
    {
        int type = judge(i);
        if (type == 0)
        {
            completion.push_back(i);
        }
        if (type == 1)
        {
            surplus.push_back(i);
        }
    }
    cout << "E: ";
    for (int j = 0; j < completion.size(); j++)
    {
        if (j < completion.size() - 1)
        {
            cout << completion[j] << " ";
        }
        else
        {
            cout << completion[j];
        }
    }
    cout << endl;
    cout << "G: ";
    for (int j = 0; j < surplus.size(); j++)
    {
        if (j < surplus.size() - 1)
        {
            cout << surplus[j] << " ";
        }
        else
        {
            cout << surplus[j];
        }
    }
    cout << endl;
}
```

### 最大序列和

给出一个整数序列 S，其中有 N 个数，定义其中一个非空连续子序列 T 中所有数的和为 T 的“序列和”。 对于 S 的所有非空连续子序列 T，求最大的序列和。 变量条件：N 为正整数，N≤1000000，结果序列和在范围（-2^63,2^63-1）以内。

第一行为一个正整数 N，第二行为 N 个整数，表示序列中的数。
输入可能包括多组数据，对于每一组输入数据，
仅输出一个数，表示最大序列和。

[网址](https://www.nowcoder.com/practice/df219d60a7af4171a981ef56bd597f7b?tpId=40&tqId=21353&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    int cnt;
    while (cin >> cnt)
    {
        vector<int> nums;
        vector<int> dp(cnt);
        for (int i = 0; i < cnt; i++)
        {
            int temp;
            cin >> temp;
            nums.push_back(temp);
        }
        for (int j = 0; j < cnt; j++)
        {
            if (j == 0)
            {
                dp[0] = nums[0];
            }
            else
            {
                dp[j] = nums[j] + dp[j - 1] > nums[j] ? nums[j] + dp[j - 1] : nums[j];
            }
        }
        int max = 0;
        for (int k = 1; k < cnt; k++)
        {
            if (dp[max] < dp[k])
            {
                max = k;
            }
        }
        cout << dp[max];
    }
    return 0;
}
```

### N 的阶乘

输入一个正整数 N，输出 N 的阶乘。

[网址](https://www.nowcoder.com/practice/f54d8e6de61e4efb8cce3eebfd0e0daa?tpId=40&tqId=21355&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#define LEN 1000
using namespace std;
int main()
{
    int factorial[LEN];

    for (int i = 0; i < LEN; i++)
    {
        factorial[i] = 0;
    }

    factorial[0] = 1;

    int num;

    while (cin >> num)
    {
        for (int i = 2; i <= num; i++)
        {
            int carry = 0;
            for (int j = 0; j < LEN; j++)
            {
                int temp = factorial[j] * i + carry;
                factorial[j] = temp % 10;
                carry = temp / 10;
            }
        }

        for (int k = LEN - 1; k >= 0; k--)
        {
            if (factorial[k] != 0)
            {
                while (k >= 0)
                {
                    cout << factorial[k];
                    k--;
                }
            }
        }
        cout << endl;
    }
    return 0;
}

```

### 放置苹果

把 M 个同样的苹果放在 N 个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？（用 K 表示）5，1，1 和 1，5，1 是同一种分法。

[网址](https://www.nowcoder.com/practice/4f0c1e21010e4d849bde5297148e81d9?tpId=40&tqId=21372&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

把 M 个同样的苹果放在 N 个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？（用 K 表示）5，1，1 和 1，5，1 是同一种分法。

```cpp
#include <iostream>
using namespace std;
/**
 *
 * 1、当只有一个盘子或者 含有 0 个 或 1 个苹果的时候只有一种方法
 * 2、当盘子数 n 大于苹果数 m 时，则必有 n - m 个空盘子，所以只需求 m 个盘子
   放 m 个苹果时的方法数即可，
   3、当盘子数 n 小于等于 苹果数 m 时，总方法数 = 当含有一个空盘子时的方法数
  + 不含空盘子时的方法数。
 */
int main()
{
    int appleCnt, dishCnt;
    cin >> appleCnt >> dishCnt;
    // dp[i][j] 表示i个苹果放到j个篮子里面
    int dp[appleCnt + 1][dishCnt + 1];
    for (int i = 0; i <= appleCnt; i++)
    {
        for (int j = 0; j <= dishCnt; j++)
        {
            dp[i][j] = 1;
        }
    }
    for (int i = 2; i <= appleCnt; i++)
    {
        for (int j = 2; j <= dishCnt; j++)
        {
            if (j > i)
            {
                dp[i][j] = dp[i][i];
            }
            else
            {
                dp[i][j] = dp[i][j - 1] + dp[i - j][j];
            }
        }
    }
    cout << dp[appleCnt][dishCnt] << endl;
    return 0;
}
```

### 大整数的因子

已知正整数 k 满足 2<=k<=9，现给出长度最大为 30 位的十进制非负整数 c，求所有能整除 c 的 k.

[网址](https://www.nowcoder.com/practice/3d6cee12fbf54ea99bb165cbaba5823d?tpId=40&tqId=21370&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    string num;
    int temp[30];
    vector<int> vec;
    while (cin >> num)
    {
        bool flag = false;
        for (int k = 2; k <= 9; k++)
        {
            int x = 0;
            for (int i = 0; i < num.size(); i++)
            {
                x = (x * 10 + (num[i] - '0')) % k;
            }
            if (x == 0)
            {
                vec.push_back(k);
                flag = true;
            }
        }
        if (flag == false)
        {
            cout << "none" << endl;
        }
        else
        {
            for (int j = 0; j < vec.size(); j++)
            {
                if (j == vec.size() - 1)
                {
                    cout << vec[j] << endl;
                }
                else
                {
                    cout << vec[j] << " ";
                }
            }
            vec.clear();
        }
    }
    return 0;
}
```

### 小白鼠排队

N 只小白鼠(1 <= N <= 100)，每只鼠头上戴着一顶有颜色的帽子。现在称出每只白鼠的重量，要求按照白鼠重量从大到小的顺序输出它们头上帽子的颜色。帽子的颜色用“red”，“blue”等字符串来表示。不同的小白鼠可以戴相同颜色的帽子。白鼠的重量用整数表示。

多案例输入，每个案例的输入第一行为一个整数 N，表示小白鼠的数目。
下面有 N 行，每行是一只白鼠的信息。第一个为不大于 100 的正整数，表示白鼠的重量，；第二个为字符串，表示白鼠的帽子颜色，字符串长度不超过 10 个字符。

注意：白鼠的重量各不相同。

[网址](https://www.nowcoder.com/practice/27fbaa6c7b2e419bbf4de8ba60cf372b?tpId=40&tqId=21368&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;
bool compDesc(pair<int, string> r1, pair<int, string> r2)
{
    return r1.first > r2.first;
}
int main()
{
    int ratCnt;
    cin >> ratCnt;
    vector<pair<int, string> > rats;
    for (int i = 0; i < ratCnt; i++)
    {
        int weigh;
        string color;
        cin >> weigh >> color;
        rats.push_back(pair<int, string>(weigh, color));
    }
    sort(rats.begin(), rats.end(), compDesc);
    for (int j = 0; j < rats.size(); j++)
    {
        cout << rats[j].second << endl;
    }
    return 0;
}

```

### 最简真分数

给出 n 个正整数，任取两个数分别作为分子和分母组成最简真分数，编程求共有几个这样的组合。

[网址](https://www.nowcoder.com/practice/1f1db273eeb745c6ac83e91ff14d2ec9?tpId=40&tqId=21366&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int gcd(int x, int y)
{
    while (y != 0)
    {

        int temp = x % y;
        x = y;
        y = temp;
    }
    return x;
}
int main()
{
    int cnt;
    while (cin >> cnt)
    {

        int total = 0;
        int num[cnt];
        for (int i = 0; i < cnt; i++)
        {
            int temp;
            cin >> temp;
            num[i] = temp;
        }
        for (int i = 0; i < cnt; i++)
        {
            for (int j = 0; j < cnt; j++)
            {
                if (num[i] == num[j])
                {
                    continue;
                }
                if (num[i] > num[j] && gcd(num[i], num[j]) == 1)
                {
                    total++;
                }
            }
        }
        cout << total << endl;
    }
    return 0;
}
```

### 密码翻译

在情报传递过程中，为了防止情报被截获，往往需要对情报用一定的方式加密，简单的加密算法虽然不足以完全避免情报被破译，但仍然能防止情报被轻易的识别。我们给出一种最简的的加密方法，对给定的一个字符串，把其中从 a-y,A-Y 的字母用其后继字母替代，把 z 和 Z 用 a 和 A 替代，则可得到一个简单的加密字符串。

[网址](https://www.nowcoder.com/practice/136de4a719954361a8e9e41c8c4ad855?tpId=40&tqId=21364&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
using namespace std;
int main()
{
    string str;
    getline(cin, str);
    for (int i = 0; i < str.length(); i++)
    {
        if ((str[i] >= 'A' && str[i] <= 'Y') ||
            (str[i] >= 'a' && str[i] <= 'y'))
        {
            str[i] += 1;
            continue;
        }
        if (str[i] == 'z')
        {
            str[i] = 'a';
            continue;
        }
        if (str[i] == 'Z')
        {
            str[i] = 'A';
            continue;
        }
    }
    cout << str << endl;
    return 0;
}

```

### 日志排序

有一个网络日志，记录了网络中计算任务的执行情况，每个计算任务对应一条如下形式的日志记录： “hs_10000_p”是计算任务的名称， “2007-01-17 19:22:53,315”是计算任务开始执行的时间“年-月-日 时：分：秒，毫秒”， “253.035(s)”是计算任务消耗的时间(以秒计) hs_10000_p 2007-01-17 19:22:53,315 253.035(s) 请你写一个程序，对日志中记录计算任务进行排序。 时间消耗少的计算任务排在前面，时间消耗多的计算任务排在后面。 如果两个计算任务消耗的时间相同，则将开始执行时间早的计算任务排在前面。

日志中每个记录是一个字符串，每个字符串占一行。最后一行为空行，表示日志结束。日志中最多可能有 10000 条记录。
计算任务名称的长度不超过 10，开始执行时间的格式是 YYYY-MM-DD HH:MM:SS,MMM，消耗时间小数点后有三位数字。
计算任务名称与任务开始时间、消耗时间之间以一个或多个空格隔开，行首和行尾可能有多余的空格

排序好的日志记录。每个记录的字符串各占一行。
输入的格式与输入保持一致，输入包括几个空格，你的输出中也应该包含同样多的空格。

[网址](https://www.nowcoder.com/practice/0f64518fea254c0187ccf0ea05019672?tpId=40&tqId=21363&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;
// 分割日志
vector<string> split_log_str(string str)
{
    int len = str.length();
    vector<string> vec;
    string temp = "";
    for (int i = 0; i < len; i++)
    {
        if (str[i] != ' ')
        {
            temp += str[i];
        }
        else
        {
            if (temp != "")
            {
                vec.push_back(temp);
                temp = "";
            }
        }
    }
    if (temp != "")
    {
        vec.push_back(temp);
    }
    return vec;
}
bool compare_log(string a, string b)
{
    vector<string> part_a = split_log_str(a);
    vector<string> part_b = split_log_str(b);
    if (part_a[3].length() < part_b[3].length())
    {
        return true;
    }
    else if (part_a[3].length() > part_b[3].length())
    {
        return false;
    }
    else
    {
        if (part_a[3] == part_b[3])
        {
            return part_a[1] + part_a[2] < part_b[1] + part_b[2];
        }
        else
        {
            return part_a[3] < part_b[3];
        }
    }
}
int main()
{
    string str;
    vector<string> vec;
    while (getline(cin, str) && str.length() != 0)
    {
        vec.push_back(str);
    }

    sort(vec.begin(), vec.end(), compare_log);
    for (int i = 0; i < vec.size(); i++)
    {
        cout << vec[i] << endl;
    }
    return 0;
}
```

### 进制转换

将一个长度最多为 30 位数字的十进制非负整数转换为二进制数输出。

多组数据，每行为一个长度不超过 30 位的十进制非负整数。
（注意是 10 进制数字的个数可能有 30 个，而非 30bits 的整数）

[网址](https://www.nowcoder.com/practice/0337e32b1e5543a19fa380e36d9343d7?tpId=40&tqId=21361&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
// 求余数
vector<int> getRemainer(string &str)
{
    vector<int> ret;
    // 0的情况
    if (str.length() == 1 && str[0] == '0')
    {
        ret.push_back(0);
        return ret;
    }

    // 非0的情况
    else
    {
        while (!str.empty())
        {
            int temp = 0;
            int len = str.length();
            for (int i = 0; i < len; i++)
            {
                // 商
                char ch = (temp * 10 + (str[i] - '0')) / 2 + '0';
                // 余数
                temp = (temp * 10 + (str[i] - '0')) % 2;

                str[i] = ch;
            }

            // 最后一位的余数就是最终的余数
            ret.push_back(temp);
            // 删掉前面的0
            int j = 0;
            while (j < str.length() && str[j] == '0')
            {
                str.erase(str.begin());
            }
        }
    }

    return ret;
}
int main()
{
    string number;
    while (cin >> number)
    {
        // 逆序输出
        vector<int> binary = getRemainer(number);
        for (int i = binary.size() - 1; i >= 0; i--)
        {
            cout << binary[i];
        }
        cout << endl;
    }
    return 0;
}
```

### 10 进制 VS 二进制

对于一个十进制数 A，将 A 转换为二进制数，然后按位逆序排列，再转换为十进制数 B，我们乘 B 为 A 的二进制逆序数。 例如对于十进制数 173，它的二进制形式为 10101101，逆序排列得到 10110101，其十进制数为 181，181 即为 173 的二进制逆序数。

一个 1000 位(即 10^999)以内的十进制数。
输入的十进制数的二进制逆序数。

[网址](https://www.nowcoder.com/practice/fd972d5d5cf04dd4bb4e5f027d4fc11e?tpId=40&tqId=21357&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
using namespace std;
// 十进制转二进制
string decimalToBinary(string number)
{

    // 只有一个0
    if (number.length() == 1 && number[0] == '0')
    {
        return "0";
    }
    else
    {
        vector<char> vec;
        while (!number.empty())
        {
            int temp = 0;
            int len = number.length();
            for (int i = 0; i < len; i++)
            {
                // 每一位的商
                char ch = (temp * 10 + (number[i] - '0')) / 2 + '0';
                // 每一位的余数
                temp = (temp * 10 + (number[i] - '0')) % 2;
                number[i] = ch;
            }

            // 最后一位的余数
            vec.push_back(temp + '0');
            // 删除开头的0对整体的商继续进行求余数
            int j = 0;
            while (j < number.length() && number[j] == '0')
            {
                number.erase(number.begin());
            }
        }
        string str(vec.begin(), vec.end());
        return str;
    }
}

string binaryToDecimal(string binaryString)
{
    vector<char> vec;
    while (!binaryString.empty())
    {
        int temp = 0;
        int len = binaryString.length();
        for (int i = 0; i < len; i++)
        {
            char ch = (temp * 2 + (binaryString[i] - '0')) / 10 + '0';
            temp = (temp * 2 + (binaryString[i] - '0')) % 10;
            binaryString[i] = ch;
        }
        vec.push_back(temp + '0');

        int j = 0;
        while (j < binaryString.length() && binaryString[j] == '0')
        {
            binaryString.erase(binaryString.begin());
        }
    }
    string str(vec.begin(), vec.end());
    reverse(str.begin(), str.end());
    return str;
}
int main()
{
    string number;
    cin >> number;
    string binaryReversed = decimalToBinary(number);
    cout << binaryToDecimal(binaryReversed) << endl;
    return 0;
}
```

### 全排列

给定一个由不同的小写字母组成的字符串，输出这个字符串的所有全排列。 我们假设对于小写字母有'a' < 'b' < ... < 'y' < 'z'，而且给定的字符串中的字母已经按照从小到大的顺序排列。

输入只有一行，是一个由不同的小写字母组成的字符串，已知字符串的长度在 1 到 6 之间。

输出这个字符串的所有排列方式，每行一个排列。要求字母序比较小的排列在前面。字母序如下定义：
已知 S = s1s2...sk , T = t1t2...tk，则 S < T 等价于，存在 p (1 <= p <= k)，使得
s1 = t1, s2 = t2, ..., sp - 1 = tp - 1, sp < tp 成立。

每组样例输出结束后要再输出一个回车。

[网址](https://www.nowcoder.com/practice/5632c23d0d654aecbc9315d1720421c1?tpId=40&tqId=21374&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
using namespace std;
// 结果数组
static vector<string> vec;
// 标记数组
static bool flag[26];

// 三个参数是原字符串、当前位置、临时字符串
void permutation(string originalString, int pos, string sequence)
{
    int len = originalString.length();
    // 一次排列完毕
    if (pos == len)
    {
        vec.push_back(sequence);
        return;
    }
    else
    {
        for (int i = 0; i < len; i++)
        {
            // 没有访问过
            if (flag[originalString[i]] == false)
            {
                // 设置访问标志
                flag[originalString[i]] = true;
                sequence.push_back(originalString[i]);
                // 排列下一个位置
                permutation(originalString, pos + 1, sequence);
                // 还原
                sequence.pop_back();
                flag[originalString[i]] = false;
            }
        }
    }
}
int main()
{
    string str;
    cin >> str;
    sort(str.begin(), str.end());
    int len = str.length();
    for (int i = 0; i < 26; i++)
    {
        flag[i] = false;
    }
    permutation(str, 0, "");
    for (int i = 0; i < vec.size(); i++)
    {
        cout << vec[i] << endl;
    }
    cout << endl;
    return 0;
}
```

### 单词替换

输入一个字符串，以回车结束（字符串长度<=100）。该字符串由若干个单词组成，单词之间用一个空格隔开，所有单词区分大小写。现需要将其中的某个单词替换成另一个单词，并输出替换之后的字符串。

多组数据。每组数据输入包括 3 行，
第 1 行是包含多个单词的字符串 s，
第 2 行是待替换的单词 a，(长度<=100)
第 3 行是 a 将被替换的单词 b。(长度<=100)

每个测试数据输出只有 1 行，
将 s 中所有单词 a 替换成 b 之后的字符串。

```cpp
#include <iostream>
#include <vector>
#include <string>
using namespace std;
int main()
{
    string sentence, toBeReplaced, replaceWord;
    getline(cin, sentence);
    cin >> toBeReplaced;
    cin >> replaceWord;
    int len = sentence.length();
    if (len == 0)
    {
        cout << "" << endl;
    }

    // 分割单词
    vector<string> words;
    string temp = "";
    for (int i = 0; i < len; i++)
    {
        if (sentence[i] == ' ')
        {
            words.push_back(temp);
            temp = "";
        }
        else
        {
            temp += sentence[i];
        }
    }
    words.push_back(temp);
    for (int i = 0; i < words.size(); i++)
    {
        if (words[i] == toBeReplaced)
        {
            words[i] = replaceWord;
        }
    }
    for (int i = 0; i < words.size(); i++)
    {
        if (i < words.size() - 1)
        {
            cout << words[i] << " ";
        }
        else
        {
            cout << words[i] << endl;
        }
    }
}
```

### 二叉树

1
/ \
 2 3
/ \ / \
 4 5 6 7
/\ /\ /\ /\
如上图所示，由正整数 1, 2, 3, ...组成了一棵无限大的二叉树。从某一个结点到根结点（编号是 1 的结点）都有一条唯一的路径，比如从 5 到根结点的路径是（5, 2, 1），从 4 到根结点的路径是（4, 2, 1），从根结点 1 到根结点的路径上只包含一个结点 1，因此路径就是（1）。对于两个结点 x 和 y，假设他们到根结点的路径分别是（x1, x2, ... ,1）和（y1, y2,...,1），那么必然存在两个正整数 i 和 j，使得从 xi 和 yj 开始，有 xi = yj，xi + 1 = yj + 1，xi + 2 = yj + 2，...
现在的问题就是，给定 x 和 y，要求他们的公共父节点，即 xi（也就是 yj）。

输入包含多组数据，每组数据包含两个正整数 x 和 y（1≤x, y≤2^31-1）。

对应每一组数据，输出一个正整数 xi，即它们的首个公共父节点。

[网址](https://www.nowcoder.com/practice/5b80ab166efa4551844657603227caeb?tpId=40&tqId=21378&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    int num1, num2;
    while (cin >> num1 >> num2)
    {
        while (num1 != num2)
        {
            if (num1 > num2)
            {
                num1 /= 2;
            }
            else
            {
                num2 /= 2;
            }
        }
        cout << num1 << endl;
    }
    return 0;
}
```

### 吃糖果

名名的妈妈从外地出差回来，带了一盒好吃又精美的巧克力给名名（盒内共有 N 块巧克力，20 > N >0）。 妈妈告诉名名每天可以吃一块或者两块巧克力。 假设名名每天都吃巧克力，问名名共有多少种不同的吃完巧克力的方案。 例如： 如果 N=1，则名名第 1 天就吃掉它，共有 1 种方案； 如果 N=2，则名名可以第 1 天吃 1 块，第 2 天吃 1 块，也可以第 1 天吃 2 块，共有 2 种方案； 如果 N=3，则名名第 1 天可以吃 1 块，剩 2 块，也可以第 1 天吃 2 块剩 1 块，所以名名共有 2+1=3 种方案； 如果 N=4，则名名可以第 1 天吃 1 块，剩 3 块，也可以第 1 天吃 2 块，剩 2 块，共有 3+2=5 种方案。 现在给定 N，请你写程序求出名名吃巧克力的方案数目。

[网址](https://www.nowcoder.com/practice/72015680c32b449899e81f1470836097?tpId=40&tqId=21379&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    int dp[20];
    dp[0] = 0;
    dp[1] = 1;
    dp[2] = 2;
    for (int i = 3; i < 20; i++)
    {
        dp[i] = dp[i - 1] + dp[i - 2];
    }
    int candyCnt;
    while (cin >> candyCnt)
    {
        cout << dp[candyCnt] << endl;
    }
    return 0;
}
```

### 与 7 无关的数

一个正整数,如果它能被 7 整除,或者它的十进制表示法中某个位数上的数字为 7, 则称其为与 7 相关的数.现求所有小于等于 n(n<100)的与 7 无关的正整数的平方和。

[网址](https://www.nowcoder.com/practice/776d401bf86d446fa783f0bef7d3c096?tpId=40&tqId=21381&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
using namespace std;
int main(){
    int n;
    cin >> n;
    int ret = 0;
    for(int i = 1;i <= n;i++){
        // false表示与7无关
        bool flag = false;
        if(i % 7 == 0){
            flag = true;
        }
        string temp = to_string(i);
        for(int j = temp.length();j >= 0;j --){
            if(temp[j] == '7'){
                flag = true;
            }
        }

        if(!flag){
            ret += i * i;
        }
    }
    cout << ret << endl;
    return 0;
}
```

### 位操作练习

给出两个不大于 65535 的非负整数，判断其中一个的 16 位二进制表示形式，是否能由另一个的 16 位二进制表示形式经过循环左移若干位而得到。 循环左移和普通左移的区别在于：最左边的那一位经过循环左移一位后就会被移到最右边去。比如： 1011 0000 0000 0001 经过循环左移一位后，变成 0110 0000 0000 0011, 若是循环左移 2 位，则变成 1100 0000 0000 0110

[网址](https://www.nowcoder.com/practice/7bdc346ca39841f6a05f73d98477621d?tpId=40&tqId=21383&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <algorithm>
using namespace std;
// 转化为16位二进制
string decimal_to_bits(int number)
{
    string remain;
    while (number != 0)
    {
        int temp = number % 2;
        number /= 2;
        remain.push_back(temp + '0');
    }
    int align0 = 16 - remain.length();
    for (int i = 0; i < align0; i++)
    {
        remain.push_back('0');
    }
    reverse(remain.begin(), remain.end());
    return remain;
}
int main()
{
    int a, b;
    while (cin >> a >> b)
    {
        if (a == b)
        {
            cout << "YES" << endl;
            continue;
        }
        string a_bits = decimal_to_bits(a);
        string b_bits = decimal_to_bits(b);
        int i;
        for (i = 0; i < 16; i++)
        {
            // 左移1位
            a_bits = a_bits.substr(1, 15) + a_bits[0];
            if (a_bits == b_bits)
            {
                cout << "YES" << endl;
                break;
            }
        }
        if (i == 16)
        {
            cout << "NO" << endl;
        }
    }
    return 0;
}
```

### 打印极值点下标

在一个整数数组上，对于下标为 i 的整数，如果它大于所有它相邻的整数， 或者小于所有它相邻的整数，则称为该整数为一个极值点，极值点的下标就是 i。

第一行是此数组的元素个数 k(4<k<80)，第二行是 k 个整数，每两个整数之间用空格分隔。

每个案例输出为 n 行：每行对应于相应数组的所有极值点下标值，下标值之间用空格分隔。

[网址](https://www.nowcoder.com/practice/7fd72f8ac7964ba3b8baa8735246e1f1?tpId=40&tqId=21385&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    int k;
    while (cin >> k)
    {
        // 数
        vector<int> nums;

        // 极值点下标
        vector<int> extremePointIndex;
        for (int i = 0; i < k; i++)
        {
            int temp;
            cin >> temp;
            nums.push_back(temp);
        }
        if (k == 1)
        {
            cout << 0 << endl;
        }
        for (int i = 0; i < k; i++)
        {
            // 第一个元素
            if (i == 0)
            {
                if (nums[0] != nums[1])
                {
                    extremePointIndex.push_back(0);
                }
                continue;
            }

            // 最后一个元素
            else if (i == k - 1)
            {
                if (nums[i] != nums[i - 1])
                {
                    extremePointIndex.push_back(k - 1);
                }
                continue;
            }
            // 1 ~ k - 2之间的值
            if ((nums[i] > nums[i - 1] && nums[i] > nums[i + 1]) ||
                (nums[i] < nums[i - 1] && nums[i] < nums[i + 1]))
            {
                extremePointIndex.push_back(i);
            }
        }

        for (int i = 0; i < extremePointIndex.size(); i++)
        {
            if (i == extremePointIndex.size() - 1)
            {
                cout << extremePointIndex[i] << endl;
            }
            else
            {
                cout << extremePointIndex[i] << " ";
            }
        }
    }
    return 0;
}
```

### 神奇的口袋

有一个神奇的口袋，总的容积是 40，用这个口袋可以变出一些物品，这些物品的总体积必须是 40。John 现在有 n 个想要得到的物品，每个物品的体积分别是 a1，a2……an。John 可以从这些物品中选择一些，如果选出的物体的总体积是 40，那么利用这个神奇的口袋，John 就可以得到这些物品。现在的问题是，John 有多少种不同的选择物品的方式。

输入的第一行是正整数 n(1 <= n <= 20) ，表示不同的物品的数目。接下来的 n 行，每行有一个 1 到 40 之间的正整数，分别给出 a1，a2……an 的值。

[网址](https://www.nowcoder.com/practice/9aaea0b82623466a8b29a9f1a00b5d35?tpId=40&tqId=21390&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;

// 递归方式
int count_v1(int i, vector<int> &vec, int leftVolumn)
{
    if (leftVolumn == 0)
    {
        return 1;
    }
    if (i == vec.size() || leftVolumn < 0)
    {
        return 0;
    }
    // count_v1(i,vec,leftVolumn) 表示从第i个物品开始，剩下的容量为
    // leftVolumn
    //  包括选中当前物品和不选中当前物品两种情况
    return count_v1(i + 1, vec, leftVolumn) +
           count_v1(i + 1, vec, leftVolumn - vec[i]);
}

// 动态规划方式
int count_v2(vector<int> &vec)
{
    // dp[i] 表示口袋容积为i时有多少种方式
    vector<int> dp(41, 0);
    dp[0] = 1;
    for (int i = 0; i < vec.size(); i++)
    {
        for (int j = 40; j >= vec[i]; j--)
        {
            dp[j] += dp[j - vec[i]];
        }
    }
    return dp[40];
}
int main()
{
    int wantedCnt;
    cin >> wantedCnt;
    vector<int> goodVolumn(wantedCnt);
    for (int i = 0; i < wantedCnt; i++)
    {
        cin >> goodVolumn[i];
    }
    //cout << count_v1(0, goodVolumn, 40) << endl;
    cout << count_v2(goodVolumn) << endl;
    return 0;
}
```

### 买房子

某程序员开始工作，年薪 N 万，他希望在中关村公馆买一套 60 平米的房子，现在价格是 200 万，假设房子价格以每年百分之 K 增长，并且该程序员未来年薪不变，且不吃不喝，不用交税，每年所得 N 万全都积攒起来，问第几年能够买下这套房子（第一年房价 200 万，收入 N 万）

有多行，每行两个整数 N（10<=N<=50）, K（1<=K<=20）

针对每组数据，如果在第 21 年或者之前就能买下这套房子，则输出一个整数 M，表示最早需要在第 M 年能买下，否则输出 Impossible，输出需要换行

[网址](https://www.nowcoder.com/practice/a4b46b53773e4a8db60b5f7629ce03e9?tpId=40&tqId=21393&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
#include <math.h>
int main()
{
    int salary, rate, price = 200;
    while (cin >> salary >> rate)
    {
        int total = salary;
        for (int i = 0; i <= 20; i++)
        {
            total += salary;
            price = ((rate + 100) / 100.0) * price;
            if (price <= total)
            {
                cout << i + 2 << endl;
                return 0;
            }
        }
        cout << "Impossible" << endl;
    }
}
```

### 素数判断

给定一个数 n，要求判断其是否为素数（0,1，负数都是非素数）。

[网址](https://www.nowcoder.com/practice/5fd9c28b1ce746dd99287a04d8fa9002?tpId=40&tqId=21494&tPage=8&rp=8&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
using namespace std;
bool isPrime(int num)
{
    if (num <= 1)
    {
        return false;
    }
    for (int i = 2; i * i <= num; i++)
    {
        if (num % i == 0)
        {
            return false;
        }
    }
    return true;
}
int main()
{
    int num;
    cin >> num;
    cout << (isPrime(num) == 1 ? "yes" : "no") << endl;
    return 0;
}
```

### 最大公约数

输入两个正整数，求其最大公约数。

[网址](https://www.nowcoder.com/practice/20216f2c84bc438eb5ef05e382536fd3?tpId=40&tqId=21492&tPage=8&rp=8&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int gcd(int num1, int num2)
{
    int time = 1;
    if (num1 < num2)
    {
        num1 = num1 ^ num2;
        num2 = num2 ^ num1;
        num1 = num1 ^ num2;
    }
    while (num1 != num2)
    {
        if ((num1 & 1) == 0 && (num2 & 1) == 0)
        {
            num1 /= 2;
            num2 /= 2;
            time++;
        }
        else
        {
            int temp = num1 - num2;
            if (temp > num2)
            {
                num1 = temp;
            }
            else
            {
                num1 = num2;
                num2 = temp;
            }
        }
    }
    return time * num2;
}
int main()
{
    int num1, num2;
    cin >> num1 >> num2;
    cout << gcd(num1, num2) << endl;
    return 0;
}
```

### 众数

输入 20 个数，每个数都在 1-10 之间，求 1-10 中的众数（众数就是出现次数最多的数，如果存在一样多次数的众数，则输出权值较小的那一个）。

[网址](https://www.nowcoder.com/practice/1549bbe3d8f546f888f4290250d9e2a6?tpId=40&tqId=21491&tPage=8&rp=8&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    int number[21];
    for (int i = 0; i <= 20; i++)
    {
        number[i] = 0;
    }
    for (int i = 0; i < 20; i++)
    {
        int temp;
        cin >> temp;
        number[temp]++;
    }
    int max = 0;
    for (int i = 1; i < 20; i++)
    {
        if (number[i] > number[max])
        {
            max = i;
        }
    }
    cout << max << endl;
    return 0;
}
```

### problemC

对于给定的字符序列，从左至右将所有的数字字符取出拼接成一个无符号整数（字符序列长度小于 100，拼接出的整数小于 2^31,），计算并输出该整数的最大素因子（如果是素数，则其最大因子为自身）

[网址](https://www.nowcoder.com/practice/2a05dcaa4cde4db989443f206ee3e5c5?tpId=40&tqId=31031&tPage=14&rp=14&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
// https://www.nowcoder.com/practice/2a05dcaa4cde4db989443f206ee3e5c5?tpId=40&tqId=31031&tPage=14&rp=14&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking
#include<iostream>
#include<string>
#include<vector>
using namespace std;
unsigned int transfer(string str){
    int len = str.length();
    unsigned int sum = 0;
    for(int i = 0;i < len;i++){
        if(str[i] >= '0' && str[i] <= '9'){
            sum = sum * 10 + str[i] - '0';
        }
    }
    return sum;
}

unsigned int maxPrime(unsigned num){
    unsigned max = 0;
    for(int i = 2;i * i <= num;i++){
        while(num % i == 0){
            num /= i;
            if(i > max)
            max = i;
        }

    }
    return num > max ? num : max;
}
int main(){
    int n;
    string temp;
    vector<unsigned int> vec;
    cin >> n;
    for(int i = 0;i < n;i++){
        cin >> temp;
        unsigned int un = transfer(temp);
        vec.push_back(maxPrime(un));
    }
    int len = vec.size();
    for(int i = 0;i < len;i++){
        cout << vec[i] << endl;
    }
    return 0;
}
```

### 对称矩阵

输入一个 N 维矩阵，判断是否对称。

可能有多组测试数据，对于每组数据，
输出"Yes!”表示矩阵为对称矩阵。
输出"No!”表示矩阵不是对称矩阵。

[网址](https://www.nowcoder.com/practice/ad11ebc8d44842c78bb0bbfb6d07ad7a?tpId=40&tqId=21552&tPage=11&rp=11&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <stdio.h>
int main()
{
    // 矩阵维数
    int n;
    int matrix[100][100];

    scanf("%d",&n);

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            scanf("%d", &matrix[i][j]);
        }
    }

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (matrix[i][j] != matrix[j][i])
            {
                printf("No!");
                return 0;
            }
        }
    }

    printf("Yes!");
    return 0;
}
```

### 打印日期

给出年分 m 和一年中的第 n 天，算出第 n 天是几月几号。

输入包括两个整数 y(1<=y<=3000)，n(1<=n<=366)。

可能有多组测试数据，对于每组数据，
按 yyyy-mm-dd 的格式将输入中对应的日期打印出来。

[网址](https://www.nowcoder.com/practice/b1f7a77416194fd3abd63737cdfcf82b?tpId=40&tqId=21554&tPage=11&rp=11&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
# include <stdio.h>
int main(){
    int days[13]= {-1,31,28,31,30,31,30,31,31,30,31,30,31};
    // 年份、天数
    int y,n;
    scanf("%d %d",&y,&n);
    // 闰年
    int loop = 0;
    if(y % 400 == 0 || (y % 4 == 0 && y % 100 != 0)){
        days[2]+=1;
    }
    for(int month = 1;month <= 12;month++){
        if(n > days[month]){
            n-=days[month];
        }else{
            printf("%04d-%02d-%02d",y,month,n);
            break;
        }

    }
    return 0;
}
```

### N 阶楼梯上楼问题

N 阶楼梯上楼问题：一次可以走两阶或一阶，问有多少种上楼方式。（要求采用非递归）

[网址](https://www.nowcoder.com/practice/c978e3375b404d598f1808e4f89ac551?tpId=40&tqId=21557&tPage=11&rp=11&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <stdio.h>
#include <stdlib.h>
int main(){
    int N;
    scanf("%d",&N);
    int *arr = (int*)malloc(sizeof(int) * (N + 1));
    arr[1] = 1;
    arr[2] = 2;
    for(int i = 3;i < N + 1;i++){
        arr[i] = arr[i - 1] + arr[i - 2];
    }
    printf("%d",arr[N]);
    return 0;
}
```

### 回文字符串

给出一个长度不超过 1000 的字符串，判断它是不是回文(顺读，逆读均相同)的。

[网址](https://www.nowcoder.com/practice/df00c27320b24278b9c25f6bb1e2f3b8?tpId=40&tqId=21559&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
    string in;
    vector<string> vec;
    while(getline(cin,in)){
        int len = in.length();
        for(int i = 0,j = len - 1;i < j;i++,j--){
            if(in[i] != in[j]){
                vec.push_back("No!");
                break;
            }else
            {
                if(i + 1 == j){
                    vec.push_back("Yes!");
                }
            }

        }
    }
    for(int i = 0;i < vec.size();i++){
        cout << vec[i] << endl;
    }
    return 0;
}
```

### 八进制

输入一个整数，将其转换成八进制数输出。

[网址](https://www.nowcoder.com/practice/eda051c1effc4dffa630bc8507f0c5f7?tpId=40&tqId=21562&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
using namespace std;
int main(){
    int num;
    cin >> num;
    vector<int> vec;
    while(num >= 8){
        vec.push_back(num % 8);
        num /= 8;
    }
    vec.push_back(num);
    int len = vec.size();
    for(int i = len - 1;i >= 0;i--){
        cout << vec.at(i);
    }
    cout << endl;
    return 0;
}
```

### 对称平方数

打印所有不超过 256，其平方具有对称性质的数。如 2，11 就是这样的数，因为 2*2=4，11*11=121。

[网址](https://www.nowcoder.com/practice/a84d46d5321f4e20931cb725e6c74fad?tpId=40&tqId=31022&tPage=14&rp=14&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
using namespace std;
bool symmetrySquare(int num)
{
    int _square = num * num;
    string _square_str = to_string(_square);
    for (int left = 0, right = _square_str.length() - 1; left < right; left++, right--)
    {
        if (_square_str[left] != _square_str[right])
        {
            return false;
        }
    }
    return true;
}
int main()
{
    for (int i = 1; i <= 256; i++)
    {
        if (symmetrySquare(i))
        {
            cout << i << endl;
        }
    }
    return 0;
}
```

### 邮票

某人有 8 角的邮票 5 张，1 元的邮票 4 张，1 元 8 角的邮票 6 张，用这些邮票中的一张或若干张可以得到多少中不同的邮资？

[网址](https://www.nowcoder.com/practice/b6735b1bd4ff488fb9a9032457410f66?tpId=40&tqId=31029&tPage=14&rp=14&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <set>
using namespace std;
int main()
{
    set<int> stampSet;
    // 8角的5张
    for (int i = 0; i <= 5; i++)
    {
        // 一元的4张
        for (int j = 0; j <= 4; j++)
        {
            // 一元八角的6张
            for (int k = 0; k <= 6; k++)
            {
                int total = 8 * i + 10 * j + 18 * k;
                if (stampSet.find(total) == stampSet.end())
                {
                    stampSet.insert(total);
                }
            }
        }
    }
    cout << stampSet.size() - 1 << endl;
    return 0;
}
```

### 一元二次方程

建立一个求一元二次方程解的类（a*x2+b*x+c=0），输入系数 a,b,c 的值后打印出这个方程的解。

[网址](https://www.nowcoder.com/practice/b5a1ceae9f884b6d8f0f798b93404b3d?tpId=40&tqId=31003&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <math.h>
#include <iomanip>
using namespace std;
int main()
{
    int cnt;
    while (cin >> cnt)
    {
        for (int i = 0; i < cnt; i++)
        {
            int a, b, c;
            cin >> a >> b >> c;
            if (a == 0)
            {
                cout << "Not quadratic equation" << endl;
                continue;
            }
            else
            {

                int delta = b * b - 4 * a * c;
                if (delta == 0)
                {
                    cout << setiosflags(ios::fixed) << setprecision(2) << "x=" << endl;
                    cout.unsetf(ios::fixed);
                }
                else if (delta > 0)
                {
                    cout << setiosflags(ios::fixed) << setprecision(2) << "x1=" << (-b - sqrt(delta)) / (2 * a) << ","
                         << "x2=" << (-b + sqrt(delta)) / (2 * a) << endl;
                    cout.unsetf(ios::fixed);
                }
                else
                {
                    cout << -1 << endl;
                }
            }
        }
    }
    return 0;
}
```

### 计算天数

输入年月日，计算该填是本年的第几天。例如 1990 年 9 月 20 日是 1990 年的第 263 天，2000 年 5 月 1 日是 2000 年第 122 天。

输入第一行为样例数 m，接下来 m 行每行 3 个整数分别表示年月日。

[网址](https://www.nowcoder.com/practice/3dc98d482fa84c1ab84384773cce1468?tpId=40&tqId=31001&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main(){
    int days[13] = {0,31,28,31,30,31,30,31,31,30,31,30,31};
    int n;
    cin >> n;
    int *out = new int[n];
    for(int i = 0;i < n;i++){
        int year,month,day,temp = 0;
        cin >> year >> month >> day;

        // 闰年2月天数
        if(year % 400 == 0 || (year % 100 != 0 && year % 4 == 0)){
            days[2]+= 1;
        }

        for(int j = 1;j < month;j++){
            temp += days[j];
        }

        out[i] = temp += day;

        // 还原
        days[2] = 28;
    }

    for(int i = 0;i < n;i++){
        cout << out[i] << endl;
    }

    return 0;
}
```

### A+B

给定两个整数 A 和 B，其表示形式是：从个位开始，每三位数用逗号","隔开。 现在请计算 A+B 的结果，并以正常形式输出。

输入包含多组数据数据，每组数据占一行，由两个整数 A 和 B 组成（-10^9 < A,B < 10^9）。

[网址](https://www.nowcoder.com/practice/b183eac8dfba4de99d47c1ca4ce9571f?tpId=40&tqId=21553&tPage=11&rp=11&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <stdio.h>
#include <string.h>
int main()
{
    char A[20];
    char B[20];
    int lengthA, lengthB;
    // 存放字符串转化成数字后的值
    int numA = 0, numB = 0;
    scanf("%s %s", A, B);
    lengthA = strlen(A);
    lengthB = strlen(B);

    // 字符串转化为数字
    for (int i = 0; i < lengthA; i++)
    {
        if (A[i] >= '0' && A[i] <= '9')
        {
            numA = numA * 10 + (A[i] - '0');
        }
    }
    if (A[0] == '-')
    {
        numA = -numA;
    }

    for (int i = 0; i < lengthB; i++)
    {
        if (B[i] >= '0' && B[i] <= '9')
        {

            numB = numB * 10 + (B[i] - '0');
        }
    }
    if (B[0] == '-')
    {
        numB = -numB;
    }
    printf("%d", numA + numB);

    return 0;
}
```

### 守形数

守形数是这样一种整数，它的平方的低位部分等于它本身。 比如 25 的平方是 625，低位部分是 25，因此 25 是一个守形数。 编一个程序，判断 N 是否为守形数。

可能有多组测试数据，对于每组数据，
输出"Yes!”表示 N 是守形数。
输出"No!”表示 N 不是守形数。

```cpp
#include <iostream>
using namespace std;
int main(){
    int num;
    cin >> num;
    int sq = num * num;
    string num_str = to_string(num);
    string sq_str = to_string(sq);
    for(int i = num_str.length() - 1,j = sq_str.length() - 1;
        i >= 0 && j >= 0;i--,j--){
            if(num_str[i] != sq_str[j]){
                cout << "No!";
                return 0;
            }
    }
    cout << "Yes!";
}
```

### 编排字符串

请输入字符串，最多输入 4 个字符串，要求后输入的字符串排在前面，例如

输入：EricZ

输出：1=EricZ

输入：David

输出：1=David 2=EricZ

输入：Peter

输出：1=Peter 2=David 3=EricZ

输入：Alan

输出：1=Alan 2=Peter 3=David 4=EricZ

输入：Jane

输出：1=Jane 2=Alan 3=Peter 4=David

第一行为字符串个数 m，接下来 m 行每行一个字符床，m 不超过 100，每个字符床长度不超过 20。

[网址](https://www.nowcoder.com/practice/42c0673f04b34f66ae236a1cb7995532?tpId=40&tqId=31014&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <vector>
using namespace std;
int main(){
    int n;
    cin >> n;
    vector<string> vec;
    for(int i = 0;i < n;i++){
        string temp;
        cin >> temp;
        vec.push_back(temp);
    }

    int cnt = vec.size();
    for(int j = 0;j < n;j++){
        for(int k = 0;k <= j;k++){
            if( k < 4){
                cout << k+1 << "=" << vec[j - k] << " ";
            }else{
                continue;
            }
        }
        cout << endl;
    }
}

```

### 多项式

实现一个多项式的类（a+b*x+c*x^2+d\*x^3+...+），要求输入该多项式的系数和 x 的值后打印出这个多项式的值。

[网址](https://www.nowcoder.com/practice/13634a38c0eb4b9db8701953ed453567?tpId=40&tqId=31004&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include<math.h>
using namespace std;
int main()
{
    int cnt;
    cin >> cnt;
    // 指数、系数
    vector<pair<int, int> > vec;
    for (int i = 0; i < cnt; i++)
    {
        int highest;
        cin >> highest;
        for (int j = 0; j <= highest; j++)
        {
            int coef;
            cin >> coef;
            vec.push_back(pair<int, int>(j, coef));
        }
        int x;
        cin >> x;
        int result = 0;
        for (int i = 0; i <= highest; i++)
        {
            result += pow(x,vec[i].first) * vec[i].second;
        }
        cout << result << endl;
    }
    return 0;
}
```

### 日期类

编写一个日期类，要求按 xxxx-xx-xx 的格式输出日期，实现加一天的操作。

[网址](https://www.nowcoder.com/practice/130aa2d7d1f5436b920229dca253893b?tpId=40&tqId=31005&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
using namespace std;
static int daysOfMonth[13] = {0,31,28,31,30,31,30,31,31,30,31,30,31};
int main(){
    int cnt;
    cin >> cnt;
    for(int i = 0;i < cnt;i++){
        int year,month,day;
        cin >> year >> month >> day;
        if((year % 400 == 0) || (year %4 == 0 && year % 100 !=0)){
            daysOfMonth[2] = 29;
        }
        // 当天不是该月的最后一天
        if(day + 1 <= daysOfMonth[month]){
            day += 1;
        }else{
            day = 1;
            // 当月不是12月
            if(month + 1 <= 12){
                month += 1;
            }else{
                month = 1;
                year+= 1;
            }
        }
        daysOfMonth[2] = 28;
        string ret = to_string(year) + "-";
        if(month < 10){
            ret+="0";
        }
        ret+=(to_string(month) + "-");
        if(day < 10){
            ret+="0";
        }
        ret+=(to_string(day));
        cout << ret << endl;
        ret.clear();
    }
    return 0;
}
```

### 单词识别

输入一个英文句子，把句子中的单词(不区分大小写)按出现次数按从多到少把单词和次数在屏幕上输出来，要求能识别英文句号和逗号，即是说单词由空格、句号和逗号隔开。

[网址](https://www.nowcoder.com/practice/16f59b169d904f8898d70d81d4a140a0?tpId=40&tqId=31019&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <map>
using namespace std;
int main(){
    // 输入的句子
    string phrase;

    // 存储单词--次数的映射
    map<string,int> wordMap;
    getline(cin,phrase);
    int len = phrase.length();
    string temp = "";
    for(int i = 0;i < len;i++){
        // 没遇到逗号,句点或空格则表明一个单词还没结束
        if(phrase[i] != ',' && phrase[i] != '.' && phrase[i] != ' '){
            temp += tolower(phrase[i]);
        }
        // 一个单词结束
        else{
            if(temp!=""){
                wordMap[temp]++;
            }
            temp = "";
        }
    }

    for(map<string,int>::iterator p = wordMap.begin();p!= wordMap.end();p++){
        cout << p->first << ":" << p->second << endl;
    }
    return 0;
}
```

### 任务调度

读入任务调度序列，输出 n 个任务适合的一种调度方式。

输入包含多组测试数据。

每组第一行输入一个整数 n（n<100000），表示有 n 个任务。

接下来 n 行，每行第一个表示前序任务，括号中的任务为若干个后序任务，表示只有在前序任务完成的情况下，后序任务才能开始。若后序为 NULL 则表示无后继任务。

输出调度方式，输出如果有多种适合的调度方式，请输出字典序最小的一种。

[网址](https://www.nowcoder.com/practice/88d5fa34fe0748e09062e48c6ae6ffc7?tpId=40&tqId=31018&tPage=13&rp=13&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
#include<vector>
#include<map>
#include<algorithm>
using namespace std;
// 将字符串分割成任务名
vector<string> split_string(string str){
    vector<string> vec;
    int len = str.length();
    string temp = "";
    for(int i = 0;i < len;i++){
        if(str[i]!=')' && str[i]!='(' && str[i]!=','){
            temp.push_back(str[i]);
        }else{
            vec.push_back(temp);
            temp = "";
        }
    }
    if(temp!=""){
        vec.push_back(temp);
    }
    return vec;
}

// 根据字符串填充图
void buildTaskMap(map<string,vector<string> > &taskMap,string str){
    vector<string> vec = split_string(str);

    for(int i = 1;i < vec.size();i++){
        if(vec[i]!="NULL"){
            // 图
            map<string,vector<string> >::iterator it = taskMap.find(vec[0]);
            if(it == taskMap.end()){
                vector<string> temp;
                temp.push_back(vec[i]);
                taskMap.insert(pair<string,vector<string> >(vec[0],temp));
            }else{
                it->second.push_back(vec[i]);
            }
        }else{
            map<string,vector<string> >::iterator it = taskMap.find(vec[0]);
            vector<string> temp;
            taskMap.insert(pair<string,vector<string> >(vec[0],temp));

        }
    }
}

vector<string> findTaskSchedule(map<string,vector<string> > &taskMap){
    // 返回最后的结果
    vector<string> retVec;
    // 入度表
    map<string,int> indegree;
    // 保存入度为0的节点
    vector<string> in0;
    // 填充入度表
    for(map<string,vector<string> >::iterator it = taskMap.begin();it!=taskMap.end();it++){
        map<string,int>::iterator tempIt = indegree.find(it->first);
        if(tempIt == indegree.end()){
            indegree.insert(pair<string,int>(it->first,0));
        }
        for(int i = 0;i < it->second.size();i++){
            tempIt = indegree.find((it->second)[i]);
            if(tempIt!=indegree.end()){
                (tempIt->second)++;
            }else{
                indegree.insert(pair<string,int>((it->second)[i],1));
            }
        }
    }

    map<string,int>::iterator it = indegree.begin();
    while(it->second == 0 && it!=indegree.end()){
        in0.push_back(it->first);
        it++;
    }

    while(!in0.empty()){
        sort(in0.begin(),in0.end());
        retVec.push_back(in0[0]);
        vector<string> vec = taskMap.find(in0[0])->second;
        for(int i = 0;i < vec.size();i++){
            int in = --(indegree.find(vec[i])->second);
            if(in == 0){
                in0.push_back(vec[i]);
            }
        }
        in0.erase(in0.begin());
    }
    return retVec;
}

int main(){
    map<string,vector<string> > taskMap;
    int cnt;
    cin >> cnt;
    for(int i = 0;i < cnt;i++){
        string temp;
        cin >> temp;
        buildTaskMap(taskMap,temp);
    }
    vector<string> vec = findTaskSchedule(taskMap);
    for(int i = 0;i < vec.size();i++){
        cout << vec[i] << " ";
    }
    cout << endl;
    return 0;
}
```

### 进制转换 2

将 M 进制的数 X 转换为 N 进制的数输出。

输入的第一行包括两个整数：M 和 N(2<=M,N<=36)。
下面的一行输入一个数 X，X 是 M 进制的数，现在要求你将 M 进制的数 X 转换成 N 进制的数输出。

[网址](https://www.nowcoder.com/practice/ae4b3c4a968745618d65b866002bbd32?tpId=40&tqId=30990&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
#include<algorithm>
#define MAX_NUMBER_SYSTEM 36
using namespace std;

void init(char *num_map){
    for(int i = 0;i < MAX_NUMBER_SYSTEM;i++){
        if(i <= 9){
            num_map[i] = (i + '0');
        }else{
            num_map[i] = 'A' + (i - 10);
        }
    }
}
// 获取符号对应的十进制数
int getDecimal(char ch){
    if(ch >= '0' && ch <= '9'){
        return ch - '0';
    }else{
        return ch - 'A' + 10;
    }
}
string transferTo(int from,int to,string str){
    // 初始化进制
    char num_map[MAX_NUMBER_SYSTEM];
    init(num_map);

    string result = "";
    while(!str.empty()){
        int remain = 0;
        for(int i = 0;i < str.length();i++){
            int temp = (remain * from + getDecimal(str[i]));
            char ch = temp / to >= 9 ? num_map[temp / to] : temp / to + '0';
            remain =  temp % to;
            str[i] = ch;
        }
        result.push_back(remain > 9 ? num_map[remain] : remain + '0');

        int j = 0;
        while(j < str.length() && str[j] == '0'){
            str.erase(str.begin());
        }
    }
    reverse(result.begin(),result.end());
    return result;
}
int main(){
    int from,to;
    string num;
    cin >> from >> to;
    cin >> num;
    cout << transferTo(from,to,num) << endl;
    return 0;
}
```

### problemB

请写一个程序，对于一个 m 行 m 列的（1<m<10）的方阵，求其每一行，每一列及主对角线元素之和，最后按照从大到小的顺序依次输出。

共一组数据，输入的第一行为一个正整数，表示 m，接下来的 m 行，每行 m 个整数表示方阵元素。

[网址](https://www.nowcoder.com/practice/bcd4ec5971054997a1dc282067aa1d8b?tpId=40&tqId=30991&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
bool com(int a,int b){
    return a  > b;
}
int  main(){
    int rows;
    cin >> rows;
    int matrix[rows][rows];

    // 每行
    vector<int> rowEle(rows);
    // 每列
    vector<int> colEle(rows);
    // 对角线
    int slish = 0;
    for(int i = 0;i < rows;i++){
        for(int j = 0;j < rows;j++){
            int temp;
            cin >> temp;
            matrix[i][j] = temp;
        }
    }
    for(int i = 0;i < rows;i++){
        for(int j = 0;j < rows;j++){
            if(i  == j){
                slish += matrix[i][j];
            }
            rowEle[i] += matrix[i][j];
            colEle[j] += matrix[i][j];
        }
    }

    vector<int> ret(rowEle.begin(),rowEle.end());
    ret.push_back(slish);
    for(int i = 0;i < colEle.size();i++){
        ret.push_back(colEle[i]);
    }

    sort(ret.begin(),ret.end(),com);
    for(int i = 0;i < ret.size();i++){
        cout << ret[i] << " ";
    }
    return 0;
}
```

### ProbleE

请写一个程序，判断给定表达式中的括号是否匹配，表达式中的合法括号为”(“, “)”, “[", "]“, “{“, ”}”,这三个括号可以按照任意的次序嵌套使用。

有多个表达式，输入数据的第一行是表达式的数目，每个表达式占一行。

[网址](https://www.nowcoder.com/practice/3bad4a646b5b47b9b85e3dcb9488a8c3?tpId=40&tqId=30993&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<string>
using namespace std;
int main(){
    int cnt;
    cin >> cnt;
    for(int i = 0;i < cnt;i++){
        vector<char> vec;
        string str;
        cin >> str;
        bool flag = true;
        for(int i =  0;i < str.length();i++){
            if(str[i] == '(' || str[i] == '[' || str[i] == '{'){
                vec.push_back(str[i]);
            }
            else if(str[i] == ')'){
                if(vec.back()!='('|| vec.empty()){
                    flag = false;
                    break;
                }else{
                    vec.pop_back();
                }
            }else if(str[i] == ']'){
                if(vec.back()!='['|| vec.empty()){
                    flag = false;
                    break;
                }else{
                    vec.pop_back();
                }
            }else if(str[i] == '}'){
                if(vec.back()!='{'|| vec.empty()){
                    flag = false;
                    break;
                }else{
                    vec.pop_back();
                }
            }
        }
        if(!flag){
            cout  << "no" << endl;
        }else{
            cout << "yes" << endl;
        }
    }
    return 0;
}

```

### 最大公约数 2

读入 n 个正整数，求出这 n 个数的最小值、最大值以及它们两的最大公约数，并输出。输入中第一行为 n，接下来为 n 个大于零的整数。

```cpp
#include<iostream>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
int maxCommonApproximate(int num1,int num2){
    if(num1 < num2){
        num1 = num1 ^ num2;
        num2 = num2 ^ num1;
        num1 = num1 ^ num2;
    }
    int times = 0;
    while(num1!=num2){
        if((num1 & 1) == 0 && (num2 & 1) == 0){
            num1 /= 2;
            num2 /= 2;
            times++;
        }else{
            int temp = num1 - num2;
            if(temp > num2){
                num1 = temp;
            }else{
                num1 = num2;
                num2 = temp;
            }
        }
    }
    return times == 0 ? num2 : times * 2 * num2;
}
int main(){
    vector<int> nums;
    int cnt;
    cin >> cnt;
    for(int i = 0;i < cnt;i++){
        int temp;
        cin >> temp;
        nums.push_back(temp);
    }
    sort(nums.begin(),nums.end());
    cout << nums[0] << " " << nums[nums.size() - 1]<< " ";
    cout << maxCommonApproximate(nums[0],nums[nums.size() - 1]) <<endl;
    return 0;
}
```

### 围圈报数

N 个人围成一圈顺序编号，从 1 号开始按 1、2、3 顺序报数，报 3 者退出圈外，其余的人再从 1、2、3 开始报数，报 3 的人再退出圈外，依次类推。请按退出顺序输出每个退出人的原序号。要求使用环行链表编程。

[网址](https://www.nowcoder.com/practice/b033062d346c4e42a7191b94164c1cd7?tpId=40&tqId=30997&tPage=12&rp=12&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<list>
using namespace std;
int main(){
    list<int> circle;
    int cnt;
    cin >> cnt;
    for(int i = 1;i  <= cnt;i++){
        circle.push_back(i);
    }
    list<int>::iterator it = circle.begin();
    while(circle.size()!=1){
        int temp = 1;
        while(temp < 3){
            it++;
            if(it == circle.end()){
                it = circle.begin();
            }
            temp++;
        }
        cout << *it << " ";
        it = circle.erase(it);
        if(it == circle.end()){
            it = circle.begin();
        }
    }
    cout << circle.front() << endl;
    return 0;
}

```

### IP 地址

输入一个 ip 地址串，判断是否合法。

[网址](https://www.nowcoder.com/practice/2359e23180194f99828f5cd9c764236a?tpId=40&tqId=21538&tPage=11&rp=11&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
using namespace std;
bool isValidIPAddress(string ip){
    string temp = "";
    for(int i = 0;i < ip.length();i++){
        if(ip[i] == '.'){
            if(atoi(temp.c_str()) < 0 ||
                atoi(temp.c_str()) > 255){
                    return false;
            }
            temp = "";
        }else{
            temp += ip[i];
        }
    }
    if(atoi(temp.c_str()) < 0 ||
                atoi(temp.c_str()) > 255){
                    return false;
    }
    return true;
}
int main(){
    string ip;
    while(getline(cin,ip)){
        isValidIPAddress(ip)?cout << "Yes!" : cout << "No!" << endl;
    }
    return 0;
}
```

### 最大上升子序列和

一个数的序列 bi，当 b1 < b2 < ... < bS 的时候，我们称这个序列是上升的。对于给定的一个序列(a1, a2, ...,aN)，我们可以得到一些上升的子序列(ai1, ai2, ..., aiK)，这里 1 <= i1 < i2 < ... < iK <= N。比如，对于序列(1, 7, 3, 5, 9, 4, 8)，有它的一些上升子序列，如(1, 7), (3, 4, 8)等等。这些子序列中序列和最大为 18，为子序列(1, 3, 5, 9)的和. 你的任务，就是对于给定的序列，求出最大上升子序列和。注意，最长的上升子序列的和不一定是最大的，比如序列(100, 1, 2, 3)的最大上升子序列和为 100，而最长上升子序列为(1, 2, 3)。

[网址](https://www.nowcoder.com/practice/dcb97b18715141599b64dbdb8cdea3bd?tpId=40&tqId=21409&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
using namespace std;
int main(){
    int len;
    while(cin >> len){

        // dp[i] 保存第i个字符之前的子序列的最大值
        vector<int> dp(len);
        vector<int> vec(len);
        for(int i = 0;i < len;i++){
            int temp;
            cin >> temp;
            vec[i] = temp;
            dp[i] = temp;
        }

        for(int i = 1;i < len;i++){
            for(int j = i - 1;j >= 0;j--){
                if(vec[j] < vec[i]){
                    dp[i] = dp[j] + vec[i] > dp[i] ? dp[j] + vec[i] : dp[i];
                }
            }
        }

        // 寻找最大值
        int max = 0;
        for(int i = 1;i < len;i++){
            if(dp[i] > dp[max]){
                max = i;
            }
        }
        cout << dp[max] << endl;
    }
    return 0;

}
```

### 八皇后问题

会下国际象棋的人都很清楚：皇后可以在横、竖、斜线上不限步数地吃掉其他棋子。如何将 8 个皇后放在棋盘上（有 8 \* 8 个方格），使它们谁也不能被吃掉！这就是著名的八皇后问题。 对于某个满足要求的 8 皇后的摆放方法，定义一个皇后串 a 与之对应，即 a=b1b2...b8，其中 bi 为相应摆法中第 i 行皇后所处的列数。已经知道 8 皇后问题一共有 92 组解（即 92 个不同的皇后串）。 给出一个数 b，要求输出第 b 个串。串的比较是这样的：皇后串 x 置于皇后串 y 之前，当且仅当将 x 视为整数时比 y 小。

每组测试数据占 1 行，包括一个正整数 b(1 <= b <= 92)

输出有 n 行，每行输出对应一个输入。输出应是一个正整数，是对应于 b 的皇后串。

[网址](https://www.nowcoder.com/practice/fbf428ecb0574236a2a0295e1fa854cb?tpId=40&tqId=21417&tPage=5&rp=5&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<string>
#include<algorithm>
using namespace std;
class Solution{
private:
    int chessboard[8][8];
    vector<string> vec;
    // 初始化棋盘
    void initChessBoard(){
        for(int i = 0;i < 8;i++){
            for(int j = 0;j < 8;j ++){
                chessboard[i][j] = 0;
            }
        }
    }
    // 检查坐标(x,y)处能否放置皇后
    bool canLay(int x,int y){
        // 同一列
        for(int i = 0;i < x;i++){
            if(chessboard[i][y] == 1){
                return false;
            }
        }
        // 对角线
        for(int i = x - 1,j = y - 1;i >= 0 && j >= 0;i--,j--){
            if(chessboard[i][j] == 1){
                return false;
            }
        }

        // 副对角线
        for(int i = x - 1,j = y + 1;i >= 0 && j <= 7;i--,j++){
            if(chessboard[i][j] == 1){
                return false;
            }
        }

        return true;
    }
    // 获取一个排列方式
    void getLayout(){
        string temp = "";
        for(int i = 0;i < 8;i++){
            for(int j = 0;j < 8;j++){
                if(chessboard[i][j] == 1){
                    temp.push_back((j + 1) + '0');
                    break;
                }
            }
        }
        vec.push_back(temp);
    }
    // 填充棋盘
    void fill(int x){
        if(x > 7){
            getLayout();
            return;
        }
        for(int j = 0;j < 8;j++){
            if(canLay(x,j)){
                chessboard[x][j] = 1;
                fill(x + 1);
                chessboard[x][j] = 0;
            }
        }
    }
public:
    void eightQueen(){
        initChessBoard();
        fill(0);
    }
    // 获取第b个后面的排列
    string getLayoutAfterB(int b){
        sort(vec.begin(),vec.end());
        return vec[b - 1];
    }
};
int main(){
    Solution so;
    int b;
    so.eightQueen();
    while(cin >> b){
        cout << so.getLayoutAfterB(b) << endl;
    }
    return 0;

}
```

### 简单密码

Julius Caesar 曾经使用过一种很简单的密码。 对于明文中的每个字符，将它用它字母表中后 5 位对应的字符来代替，这样就得到了密文。 比如字符 A 用 F 来代替。如下是密文和明文中字符的对应关系。 密文 A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 明文 V W X Y Z A B C D E F G H I J K L M N O P Q R S T U 你的任务是对给定的密文进行解密得到明文。 你需要注意的是，密文中出现的字母都是大写字母。密文中也包括非字母的字符，对这些字符不用进行解码。

[网址](https://www.nowcoder.com/practice/ff99c43dd07f4e95a8f2f5448da3772a?tpId=40&tqId=21421&tPage=5&rp=5&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<string>
using namespace std;
int main(){
    vector<char> encryption(26);
    for(int i = 0;i < 26;i++){
        encryption[i] = (i - 5 + 26) % 26 + 'A';
    }

    string str;
    while(getline(cin,str)){
        if(str == "ENDOFINPUT"){
            return 0;
        }
        if(str == "START" || str == "END"){
            continue;
        }else{
            for(int i = 0;i < str.length();i++){
                if(str[i] >= 'A' && str[i] <= 'Z'){
                    str[i] = encryption[str[i] - 'A'];
                }
            }
            cout << str << endl;
        }

    }
    return 0;
}

```

### 计算表达式

对于一个不存在括号的表达式进行计算

存在多种数据，每组数据一行，表达式不存在空格

[网址](https://www.nowcoder.com/practice/7b18aa6b7cc14f8eaae6b8acdebf890b?tpId=40&tqId=21433&tPage=5&rp=5&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<string>
#include<ctype.h>
using namespace std;
// 用符号栈顶的符号对数字栈顶部的两个数进行运算
void simpleCalculate(vector<double> &num_vec,vector<char> &operator_vec){
    // 运算符栈顶部的运算符
    char operator_top = operator_vec.back();

    double num2 = num_vec.back();
    num_vec.pop_back();
    double num1 = num_vec.back();
    num_vec.pop_back();

    double temp;
    switch(operator_top){
        case '+': temp = num1 + num2;break;
        case '-': temp = num1 - num2;break;
        case '*': temp = num1 * num2;break;
        case '/': temp = num1 / num2;break;
    }
    num_vec.push_back(temp);
    operator_vec.pop_back();
}
int main(){
    // 数字栈
    vector<double> num_vec;
    // 运算符栈
    vector<char> operator_vec;

    string expression;
    cin >> expression;
    for(int i = 0;i < expression.length();i++){

        // 数字
        if(isdigit(expression[i])){
            int temp = expression[i] - '0';
            int j;
            for(j = i+1;j < expression.length() && isdigit(expression[j]);j++){
                temp = temp * 10 + (expression[j] - '0');
            }
            num_vec.push_back(temp);
            i = j - 1;
        }
        // +,- 符号
        else if(expression[i] == '+' || expression[i] == '-'){
            if(operator_vec.empty()){
                operator_vec.push_back(expression[i]);
            }else{
                simpleCalculate(num_vec,operator_vec);
                i--;
            }
        }
        // *,/ 符号
        else if(expression[i] == '*' || expression[i] == '/'){
            if(operator_vec.empty() || operator_vec.back() == '+' || operator_vec.back() == '-'){
                operator_vec.push_back(expression[i]);
            }else{
                simpleCalculate(num_vec,operator_vec);
                i--;
            }
        }
    }

    // 对符号栈剩下的符号运算
    while(num_vec.size()!=1){
        simpleCalculate(num_vec,operator_vec);
    }

    cout << num_vec.back() << endl;
    return 0;
}
```

### 棋盘游戏

有一个 6\*6 的棋盘，每个棋盘上都有一个数值，现在又一个起始位置和终止位置，请找出一个从起始位置到终止位置代价最小的路径： 1、只能沿上下左右四个方向移动 2、总代价是没走一步的代价之和 3、每步（从 a,b 到 c,d）的代价是 c,d 上的值与其在 a,b 上的状态的乘积 4、初始状态为 1 每走一步，状态按如下公式变化：（走这步的代价%4）+1。

[网址](https://www.nowcoder.com/practice/368c98c7bff54a30bba29ae1ba017d55?tpId=40&tqId=21429&tPage=5&rp=5&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <limits.h>
#define CNT 6
using namespace std;
// 方向上右下左(x,y)
static int direction[4][2] = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
class Solution
{
private:
    // 访问状态
    bool visited[CNT][CNT];

    // 棋盘的状态
    int board[CNT][CNT];

    // 最小代价
    unsigned min_cost = INT_MAX;

    // 目标点
    int dest_x, dest_y;

public:
    void setDest(int x, int y)
    {
        dest_x = x;
        dest_y = y;
    }
    // 初始化状态
    void initBoard()
    {
        for (int i = 0; i < CNT; i++)
        {
            for (int j = 0; j < CNT; j++)
            {
                cin >> board[i][j];
                visited[i][j] = false;
            }
        }
    }

    void traverse(int x, int y, unsigned cost, int state)
    {
        if (x >= 0 && x < CNT && y >= 0 && y < CNT && !visited[x][y])
        {
            // 当前代价比最小代价小，必然不符合条件
            if (min_cost < cost)
            {
                return;
            }

            // 目标点
            if (x == dest_x && y == dest_y)
            {
                min_cost = cost;
            }
            visited[x][y] = true;

            for (int i = 0; i < 4; i++)
            {
                // 某个方向的下一步的坐标
                int new_x = x + direction[i][0];
                int new_y = y + direction[i][1];

                // 新的代价
                unsigned newcost = board[new_x][new_y] * state;

                // 新的状态
                int newstate = (newcost % 4) + 1;

                // 四个方向
                traverse(new_x, new_y, cost + newcost, newstate);
            }
            visited[x][y] = false;
        }
    }

    unsigned getCost()
    {
        return min_cost;
    }
};
int main()
{
    Solution so;
    so.initBoard();

    int start_x, start_y, end_x, end_y;
    cin >> start_x >> start_y >> end_x >> end_y;
    so.setDest(end_x, end_y);
    so.traverse(start_x, start_y, 0, 1);
    cout << so.getCost() << endl;
    return 0;
}
```

### 后缀子串排序

对于一个字符串，将其后缀子串进行排序，例如 grain 其子串有： grain rain ain in n 然后对各子串按字典顺序排序，即： ain,grain,in,n,rain

[网址](https://www.nowcoder.com/practice/f89f96ea3145418b8e6c3eb75773f65a?tpId=40&tqId=21446&tPage=6&rp=6&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;
int main()
{
    vector<string> vec;
    string str;
    cin >> str;
    for (int i = 0; i <= str.length() - 1; i++)
    {
        vec.push_back(str.substr(i, str.length() - i));
    }
    sort(vec.begin(), vec.end());
    for (int i = 0; i < vec.size(); i++)
    {
        cout << vec[i] << endl;
    }
    return 0;
}
```

### 字符统计

输入一行字符串，计算其中 A-Z 大写字母出现的次数

[网址](https://www.nowcoder.com/practice/de7bf0945c1c4bd1aa9d49573b831f3c?tpId=40&tqId=21444&tPage=6&rp=6&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <string>
using namespace std;
int main()

{
    vector<int> ch(26);
    for (int i = 0; i < 26; i++)
    {
        ch[i] = 0;
    }
    string str;
    while (cin >> str)
    {
        for (int i = 0; i < str.length(); i++)
        {
            if (str[i] >= 'A' && str[i] <= 'Z')
            {
                ch[str[i] - 'A']++;
            }
        }
        for (int i = 0; i < 26; i++)
        {
            cout << char('A' + i) << ":" << ch[i] << endl;
        }
    }

    return 0;
}
```

### 统计单词

编一个程序，读入用户输入的，以“.”结尾的一行文字，统计一共有多少个单词，并分别输出每个单词含有多少个字符。 （凡是以一个或多个空格隔开的部分就为一个单词）

输入包括 1 行字符串，以“.”结束，字符串中包含多个单词，单词之间以一个或多个空格隔开。

[网址](https://www.nowcoder.com/practice/11c6e7c9987c4ab48f8cdd8834c27064?tpId=40&tqId=21537&tPage=10&rp=10&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
    string phrase;
    vector<int> len_word;
    getline(cin,phrase);
        int len =  phrase.length();
        string temp = "";
        for(int i = 0;i < len;i++){
            if(phrase[i]!=' ' && phrase[i]!='.'){
                temp += phrase[i];
            }else{
                len_word.push_back(temp.length());
                temp = "";
            }
        }

    int vec_size = len_word.size();
    for(int i = 0;i <= vec_size - 1;i++){
        cout << len_word[i] << " ";
    }
    return 0;
}
```

### 字符串反码

一个二进制数，将其每一位取反，称之为这个数的反码。下面我们定义一个字符的反码。如果这是一个小写字符，则它和字符'a’的距离与它的反码和字符'z’的距离相同；如果是一个大写字符，则它和字符'A’的距离与它的反码和字符'Z’的距离相同；如果不是上面两种情况，它的反码就是它自身。 举几个例子，'a’的反码是'z’；'c’的反码是'x’；'W’的反码是'D’；'1’的反码还是'1’；'$'的反码还是'$'。 一个字符串的反码定义为其所有字符的反码。我们的任务就是计算出给定字符串的反码。

[网址](https://www.nowcoder.com/practice/01b7dae14d1b464db5f9259e90d9a35e?tpId=40&tqId=21503&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
using namespace std;
int main()
{
    string str;
    while (getline(cin, str))
    {
        if (str.length() == 1 && str[0] == '!')
        {
            return 0;
        }
        for (int i = 0; i < str.length(); i++)
        {
            if (str[i] >= 'a' && str[i] <= 'z')
            {
                str[i] = 'z' - (str[i] - 'a');
            }
            if (str[i] >= 'A' && str[i] <= 'Z')
            {
                str[i] = 'Z' - (str[i] - 'A');
            }
        }
        cout << str << endl;
    }
    return 0;
}
```

### 连通图

给定一个无向图和其中的所有边，判断这个图是否所有顶点都是连通的。

每组数据的第一行是两个整数 n 和 m（0<=n<=1000）。n 表示图的顶点数目，m 表示图中边的数目。随后有 m 行数据，每行有两个值 x 和 y（0<x, y <=n），表示顶点 x 和 y 相连，顶点的编号从 1 开始计算。输入不保证这些边是否重复。

[网址](https://www.nowcoder.com/practice/569e89823a5141fe8a11ab7d4da21edf?tpId=40&tqId=21506&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <map>
using namespace std;
static int cnt = 0;
void dfs(map<int,vector<int> > graph, vector<int> &flag, int x, int vertexCnt)
{
    if (!flag[x])
    {
        flag[x] = true;
        cnt++;
        vector<int> it = graph.find(x)->second;
        for (int i = 0; i < it.size(); i++)
        {
            if (!flag[it[i]])
            {
                dfs(graph, flag, it[i], vertexCnt);
            }
        }
    }
}
int main()
{
    int vertexs, arcs;
    cin >> vertexs >> arcs;

     map<int,vector<int> > graph;
    for (int i = 0; i < vertexs; i++)
    {
       vector<int> tempVec;
       graph.insert(pair<int,vector<int> >(i,tempVec));
    }

    for(int i = 0;i < arcs;i++){
        int temp1,temp2;
        cin >> temp1  >> temp2;
        map<int,vector<int> >::iterator it1 = graph.find(temp1 - 1);
        map<int,vector<int> >::iterator it2 = graph.find(temp2 - 1);

        for(int i = 0;i < it1->second.size();i++){
                if((it1->second)[i] == temp2 - 1){
                    i--;
                    continue;
                }
            }
            it1->second.push_back(temp2 - 1);
            it2->second.push_back(temp1 - 1);
    }

    // 这里不知道为啥，用bool内存就会占完
    vector<int> visited(vertexs);
    for (int i = 0; i < vertexs; i++)
    {
        visited[i] = 0;
    }

    dfs(graph, visited, 0, vertexs);
    if(cnt == vertexs){
        cout << "YES" << endl;
    }else{
        cout << "NO" << endl;
    }
    return 0;
}
```

### 排列与二进制

在组合数学中，我们学过排列数。从 n 个不同元素中取出 m（m<=n）个元素的所有排列的个数，叫做从 n 中取 m 的排列数，记为 p(n, m)。具体计算方法为 p(n, m)=n(n-1)(n-2)……(n-m+1)= n!/(n-m)! (规定 0!=1).当 n 和 m 不是很小时，这个排列数是比较大的数值，比如 p(10,5)=30240。如果用二进制表示为 p(10,5)=30240=( 111011000100000)b，也就是说，最后面有 5 个零。我们的问题就是，给定一个排列数，算出其二进制表示的后面有多少个连续的零。

[网址](https://www.nowcoder.com/practice/647fc23dc4e147328cc484e3aeb6cc2a?tpId=40&tqId=21507&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
using namespace std;
/*
 * 求二进制末尾有多少个零，等价于求n!/(n-m)!中可以整除2几次，
 * n!/(n-m)!可以写成因子乘积形式 n*(n-1)*...*(n-m+1) ，各因子可以整除2的次数之和即为所求
 */
int main(){
    int n,m;
    while(cin >> n >> m && n != 0){
        int cnt = 0;
        for(int i = n - m + 1;i <= n;i++){
            int x = i;
            while((x & 1) == 0){
                x = x >> 1;
                cnt++;
            }
        }
        cout << cnt << endl;
    }
    return 0;
}
```

### 数字之和

对于给定的正整数 n，计算其十进制形式下所有位置数字之和，并计算其平方的各位数字之和。

[网址](https://www.nowcoder.com/practice/ae759916631f4711a90c4d4d9657f4b0?tpId=40&tqId=21509&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
using namespace std;
int main(){
    int num;
    while(cin >> num){
        string num_str = to_string(num);
        string square_str = to_string(num * num);
        int temp = 0;
        for(int i = 0;i < num_str.length();i++){
            temp += (num_str[i] - '0');
        }
        cout << temp << " ";
        temp = 0;
        for(int i = 0;i < square_str.length();i++){
            temp += (square_str[i] - '0');
        }
        cout << temp << endl;
    }
    return 0;
}
```

### 寻找第 k 小数

查找一个数组的第 K 小的数，注意同样大小算一样大。 如 2 1 3 4 5 2 第三小数为 3。

[网址](https://www.nowcoder.com/practice/204dfa6fcbc8478f993d23f693189ffd?tpId=40&tqId=21522&tPage=10&rp=10&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main(){
    int cnt;
    cin >> cnt;

    int *arr = new int[1001];
    for(int i = 0;i < 1001;i++){
        arr[i] = 0;
    }
    for(int i = 0;i < cnt;i++){
        int temp;
        cin >> temp;
        arr[temp]++;
    }

    int k;
    cin >> k;

    int order = 1;

    for(int j = 1;j < 1001;j++){
        if(arr[j]!=0){
            if(order == k){
                cout << j;
                return 0;
            }else{
                order++;
                continue;
            }
        }
    }
    return 0;
}
```

### 合并字符串

给定两个字符串 S1 和 S2，合并成一个新的字符串 S。 合并规则为，S1 的第一个字符为 S 的第一个字符，将 S2 的最后一个字符作为 S 的第二个字符； 将 S1 的第二个字符作为 S 的第三个字符，将 S2 的倒数第二个字符作为 S 的第四个字符，以此类推。

[网址](https://www.nowcoder.com/practice/7f436c901a0d450ebdec1168e3e57cc2?tpId=40&tqId=21533&tPage=10&rp=10&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<string>
using namespace std;
int main(){
    string str1,str2;
    cin >> str1;
    cin >> str2;
    int len1 = str1.length();
    int len2 = str2.length();
    int index1 = 0;
    int index2 = len2 - 1;
    string temp = "";
    while(index1 < len1 && index2 >= 0){
        temp += str1[index1++];
        temp += str2[index2--];
    }
    cout << temp << endl;
    return 0;
}
```

### 杨辉三角形

输入 n 值，使用递归函数，求杨辉三角形中各个位置上的值。

[网址](https://www.nowcoder.com/practice/ef7f264886a14fdf8a6ed3ac008a23c8?tpId=40&tqId=21535&tPage=10&rp=10&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
using namespace std;
int main(){
    int n;
    cin >> n;
    vector< vector<int> > vec;
    for(int i = 1;i <= n - 1;i++){
        vector<int> temp;
        for(int j = 1;j <= i + 1;j++){
            // 两边
            if(j == 1 || j == i + 1){
                temp.push_back(1);
            }
            // 中间的等于肩上两个元素之和
            else{
                temp.push_back(vec[i - 2][j - 1] + vec[i - 2][j - 2] );
            }
        }
        vec.push_back(temp);
    }

    for(int i = 0;i < n - 1;i++){
        int len = vec[i].size();
        for(int j = 0;j < len;j++){
            cout << vec[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}
```

### 搬水果

在一个果园里，小明已经将所有的水果打了下来，并按水果的不同种类分成了若干堆，小明决定把所有的水果合成一堆。每一次合并，小明可以把两堆水果合并到一起，消耗的体力等于两堆水果的重量之和。当然经过 n‐1 次合并之后，就变成一堆了。小明在合并水果时总共消耗的体力等于每次合并所耗体力之和。 假定每个水果重量都为 1，并且已知水果的种类数和每种水果的数目，你的任务是设计出合并的次序方案，使小明耗费的体力最少，并输出这个最小的体力耗费值。例如有 3 种水果，数目依次为 1，2，9。可以先将 1，2 堆合并，新堆数目为 3，耗费体力为 3。然后将新堆与原先的第三堆合并得到新的堆，耗费体力为 12。所以小明总共耗费体力=3+12=15，可以证明 15 为最小的体力耗费值。

​ 每组数据输入包括两行,第一行是一个整数 n(1<=n<=10000),表示水果的种类数。第二行包含 n 个整数，用空格分隔，第 i 个整数(1<=ai<=1000)是第 i 种水果的数目。

[网址](https://www.nowcoder.com/practice/e4c775b0f3ee42a4bb72c26d2e1eef8a?tpId=40&tqId=21510&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<queue>
#include<vector>
#include<functional>
using namespace std;
/**
* 小顶堆
*/
int main(){
    int fruitType;
    while(cin >> fruitType && fruitType!=0){
        priority_queue<int,vector<int>,greater<int> > qu;
        int energy = 0;
        for(int i = 0;i < fruitType;i++){
            int cnt;
            cin >> cnt;
            qu.push(cnt);
        }
        while(qu.size() > 1){
            int fruitOne = qu.top();
            qu.pop();
            int fruitTwo = qu.top();
            qu.pop();
            energy += (fruitOne + fruitTwo);
            qu.push(fruitOne + fruitTwo);
        }

        cout << energy << endl;
        qu.pop();
    }
    return 0;
}
```

### 堆栈的使用

堆栈是一种基本的数据结构。堆栈具有两种基本操作方式，push 和 pop。Push 一个值会将其压入栈顶，而 pop 则会将栈顶的值弹出。现在我们就来验证一下堆栈的使用。

对于每组测试数据，第一行是一个正整数 n，0<n<=10000(n=0 结束)。而后的 n 行，每行的第一个字符可能是'P’或者'O’或者'A’；如果是'P’，后面还会跟着一个整数，表示把这个数据压入堆栈；如果是'O’，表示将栈顶的值 pop 出来，如果堆栈中没有元素时，忽略本次操作；如果是'A’，表示询问当前栈顶的值，如果当时栈为空，则输出'E'。堆栈开始为空。

[网址](https://www.nowcoder.com/practice/e91982a145944ceab6bb9a4a508e0e26?tpId=40&tqId=21511&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    int cnt;
    while (cin >> cnt && cnt != 0)
    {
        vector<int> vec;

        for (int i = 0; i < cnt; i++)
        {
            char ch;
            cin >> ch;
            switch (ch)
            {
            case 'P':
            {
                int num;
                cin >> num;
                vec.push_back(num);
            }
            break;

            case 'O':
            {
                if (vec.size() > 0)
                {
                    vec.pop_back();
                }
            }
            break;
            case 'A':
            {
                if (vec.size() > 0)
                {
                    cout << vec.back() << endl;
                }
                else
                {
                    cout << (char)'E' << endl;
                }
            }
            break;
            }
        }
        cout << endl;
    }
    return 0;
}
```

### 比较奇偶数的个数

第一行输入一个数，为 n，第二行输入 n 个数，这 n 个数中，如果偶数比奇数多，输出 NO，否则输出 YES。

[网址](https://www.nowcoder.com/practice/6ba928cea3734205b042b9445de9aecb?tpId=40&tqId=21512&tPage=9&rp=9&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
using namespace std;
int main()
{
    int cnt;
    while (cin >> cnt && cnt != 0)
    {
        int oddCnt = 0, evenCnt = 0;
        for (int i = 0; i < cnt; i++)
        {
            int temp;
            cin >> temp;
            if ((temp & 1) == 0)
            {
                evenCnt++;
            }
            else
            {
                oddCnt++;
            }
        }
        if (evenCnt > oddCnt)
        {
            cout << "NO" << endl;
        }
        else
        {
            cout << "YES" << endl;
        }
    }
    return 0;
}
```

### 最少邮票数

有若干张邮票，要求从中选取最少的邮票张数凑成一个给定的总值。 如，有 1 分，3 分，3 分，3 分，4 分五张邮票，要求凑成 10 分，则使用 3 张邮票：3 分、3 分、4 分即可。

有多组数据，对于每组数据，首先是要求凑成的邮票总值 M，M<100。然后是一个数 N，N〈20，表示有 N 张邮票。接下来是 N 个正整数，分别表示这 N 张邮票的面值，且以升序排列

对于每组数据，能够凑成总值 M 的最少邮票张数。若无解，输出 0。

[网址](https://www.nowcoder.com/practice/83800ae3292b4256b7349ded5f178dd1?tpId=40&tqId=21345&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    // 总值,邮票数
    int totalVal, stampCnt;

    cin >> totalVal >> stampCnt;

    // 每张邮票的面值
    vector<int> stampVal(stampCnt);

    for (int i = 0; i < stampCnt; i++)
    {
        cin >> stampVal[i];
    }

    // dp[i]表示能凑成总值为i的邮票的最小数
    vector<int> dp(totalVal + 1, stampCnt + 1);
    dp[0] = 0;
    for (int i = 0; i < stampCnt; i++)
    {

        // 设想假如正序的话，现在有一张邮票的价值为1，DP[1] 会使用到1，
        // DP[2] = min(DP[2], DP[1] + 1), 这样就相当于使用了两次
        // 价值为1的邮票，与题意显然不符合，因此第二层循环一定要倒序。
        for (int j = totalVal; j >= 0; j--)
        {
            // 选中当前邮票后的总张数 j - stampVal[i] + 1
            if (j - stampVal[i] >= 0 && dp[j] > dp[j - stampVal[i]] + 1)
            {
                dp[j] = dp[j - stampVal[i]] + 1;
            }
        }
    }

    if (dp[totalVal] == stampCnt + 1)
    {
        cout << 0 << endl;
    }
    else
    {
        cout << dp[totalVal] << endl;
    }
    return 0;
}
```

### 玛雅人的密码

玛雅人有一种密码，如果字符串中出现连续的 2012 四个数字就能解开密码。给一个长度为 N 的字符串，（2=<N<=13）该字符串中只含有 0,1,2 三种数字，问这个字符串要移位几次才能解开密码，每次只能移动相邻的两个数字。例如 02120 经过一次移位，可以得到 20120,01220,02210,02102，其中 20120 符合要求，因此输出为 1.如果无论移位多少次都解不开密码，输出-1。

输入包含多组测试数据，每组测试数据由两行组成。
第一行为一个整数 N，代表字符串的长度（2<=N<=13）。
第二行为一个仅由 0、1、2 组成的，长度为 N 的字符串。
对于每组测试数据，若可以解出密码，输出最少的移位次数；否则输出-1。

[网址](https://www.nowcoder.com/practice/761fc1e2f03742c2aa929c19ba96dbb0?tpId=40&tqId=21343&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <queue>
#include <map>
#include <string>
using namespace std;

// 是否含有模式串
bool contains(string str, string pattern)
{
    if (str.find(pattern) != string::npos)
    {
        return true;
    }
    return false;
}

// 交换字符串第i个字符及其后面的字符,i不能取最后一个字符位置
string swap_i(string str, int i)
{
    string retStr = str;
    char ch = retStr[i];
    retStr[i] = retStr[i + 1];
    retStr[i + 1] = ch;
    return retStr;
}

// 深度优先搜索
int bfs(string str, string pattern)
{
    // 记录字符串交换的次数
    map<string, int> timesMap;
    // 辅助队列
    queue<string> qu;
    qu.push(str);
    timesMap.insert(pair<string, int>(str, 0));

    while (!qu.empty())
    {
        str = qu.front();
        qu.pop();
        for (int i = 0; i < str.length() - 1; i++)
        {
            string tempStr = swap_i(str, i);
            if (timesMap.find(tempStr) != timesMap.end())
            {
                continue;
            }
            else
            {
                timesMap[tempStr] = timesMap[str] + 1;
                if (contains(tempStr, pattern))
                {
                    return timesMap[tempStr];
                }
                qu.push(tempStr);
            }
        }
    }

    return -1;
}
int main()
{
    string pattern = "2012";
    int len;
    while (cin >> len)
    {
        string str;
        cin >> str;
        if (contains(str, pattern))
        {
            cout << 0 << endl;
        }
        else
        {
            cout << bfs(str, pattern);
        }
    }
    return 0;
}
```

### 剩下的树

有一个长度为整数 L(1<=L<=10000)的马路，可以想象成数轴上长度为 L 的一个线段，起点是坐标原点，在每个整数坐标点有一棵树，即在 0,1,2，...，L 共 L+1 个位置上有 L+1 棵树。 现在要移走一些树，移走的树的区间用一对数字表示，如 100 200 表示移走从 100 到 200 之间（包括端点）所有的树。 可能有 M(1<=M<=100)个区间，区间之间可能有重叠。现在要求移走所有区间的树之后剩下的树的个数。

两个整数 L(1<=L<=10000)和 M(1<=M<=100)。

接下来有 M 组整数，每组有一对数字。

[网址](https://www.nowcoder.com/practice/f5787c69f5cf41499ba4706bc93700a2?tpId=40&tqId=21356&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    int L, M;
    while (cin >> L >> M)
    {
        vector<int> vec(L + 1, 1);
        int start, end, cnt = 0;
        for (int k = 0; k < M; k++)
        {
            cin >> start >> end;
            for (int i = start; i <= end; i++)
            {
                vec[i] = 0;
            }
        }
        for (int j = 0; j < L + 1; j++)
        {
            if (vec[j] == 1)
            {
                cnt++;
            }
        }
        cout << cnt << endl;
    }
    return 0;
}
```

### 谁是你的潜在朋友

“臭味相投”——这是我们描述朋友时喜欢用的词汇。两个人是朋友通常意味着他们存在着许多共同的兴趣。然而作为一个宅男，你发现自己与他人相互了解的机会并不太多。幸运的是，你意外得到了一份北大图书馆的图书借阅记录，于是你挑灯熬夜地编程，想从中发现潜在的朋友。 首先你对借阅记录进行了一番整理，把 N 个读者依次编号为 1,2,…,N，把 M 本书依次编号为 1,2,…,M。同时，按照“臭味相投”的原则，和你喜欢读同一本书的人，就是你的潜在朋友。你现在的任务是从这份借阅记录中计算出每个人有几个潜在朋友。

每个案例第一行两个整数 N,M，2 <= N ，M<= 200。接下来有 N 行，第 i(i = 1,2,…,N)行每一行有一个数，表示读者 i-1 最喜欢的图书的编号 P(1<=P<=M)

每个案例包括 N 行，每行一个数，第 i 行的数表示读者 i 有几个潜在朋友。如果 i 和任何人都没有共同喜欢的书，则输出“BeiJu”（即悲剧，^ ^）

```cpp
#include <iostream>
#include <vector>
using namespace std;
int main()
{
    int N, M;
    cin >> N >> M;
    // 每个人喜欢的书
    vector<int> books(N);
    // 每本书喜欢的人数
    vector<int> commonLikes(M, 0);
    for (int i = 0; i < N; i++)
    {
        int book;
        cin >> book;
        books[i] = book;
        commonLikes[book]++;
    }
    for (int i = 0; i < N; i++)
    {
        if (commonLikes[books[i]] > 1)
        {
            cout << commonLikes[books[i]] - 1 << endl;
        }
        else
        {
            cout << "BeiJu" << endl;
        }
    }
    return 0;
}
```

### 中位数

中位数定义：一组数据按从小到大的顺序依次排列，处在中间位置的一个数（或最中间两个数据的平均数）. 给出一组无序整数，求出中位数，如果求最中间两个数的平均数，向下取整即可（不需要使用浮点数）

该程序包含多组测试数据，每一组测试数据的第一行为 N，代表该组测试数据包含的数据个数，1<=N<=10000.
接着 N 行为 N 个数据的输入，N=0 时结束输入

[网址](https://www.nowcoder.com/practice/2364ff2463984f09904170cf6f67f69a?tpId=40&tqId=21367&tPage=2&rp=2&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    int cnt;
    while (cin >> cnt && cnt != 0)
    {
        vector<int> vec(cnt);
        for (int i = 0; i < cnt; i++)
        {
            cin >> vec[i];
        }
        sort(vec.begin(), vec.end());

        if ((cnt & 1) == 1)
        {
            cout << vec[cnt / 2] << endl;
        }
        else
        {
            cout << (vec[cnt / 2] + vec[cnt / 2 - 1]) / 2 << endl;
        }
    }
    return 0;
}
```

### 首字母大写

对一个字符串中的所有单词，如果单词的首字母不是大写字母，则把单词的首字母变成大写字母。 在字符串中，单词之间通过空白符分隔，空白符包括：空格(' ')、制表符('\t')、回车符('\r')、换行符('\n')。

输入一行：待处理的字符串（长度小于 100）。
可能有多组测试数据，对于每组数据，
输出一行：转换后的字符串。

[网址](https://www.nowcoder.com/practice/91f9c70e7b6f4c0ab23744055632467a?tpId=40&tqId=21388&tPage=3&rp=3&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <ctype.h>
using namespace std;
bool isBlankChar(char ch)
{
    if (ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n')
    {
        return true;
    }
    return false;
}
int main()
{
    string str;
    while (getline(cin, str) && str != " ")
    {
        if (islower(str[0]))
        {
            str[0] = toupper(str[0]);
        }
        for (int i = 1, len = str.length(); i < len; i++)
        {
            if (isBlankChar(str[i - 1]) && islower(str[i]))
            {
                str[i] = toupper(str[i]);
            }
        }
        cout << str << endl;
    }
    return 0;
}
```

### 求平均年龄

班上有学生若干名，给出每名学生的年龄（整数），求班上所有学生的平均年龄，保留到小数点后两位。

第一行有一个整数 n（1<= n <= 100），表示学生的人数。其后 n 行每行有 1 个整数，取值为 15 到 25。

[网址](https://www.nowcoder.com/practice/ca319fdb02714994850cc631d76f5547?tpId=40&tqId=21402&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
using namespace std;
int main(){
    int totalCnt;
    cin >> totalCnt;

    int ageSum = 0;
    for(int i = 0;i < totalCnt;i++){
        int temp;
        cin >> temp;
        ageSum += temp;
    }

    cout.setf(ios::fixed);
    cout.precision(2);
    cout << (double)ageSum / totalCnt << endl;
    return 0;
}
```

### 整数奇偶排序

输入 10 个整数，彼此以空格分隔。重新排序以后输出(也按空格分隔)，要求: 1.先输出其中的奇数,并按从大到小排列； 2.然后输出其中的偶数,并按从小到大排列。

[网址](https://www.nowcoder.com/practice/bbbbf26601b6402c9abfa88de5833163?tpId=40&tqId=21398&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
using namespace std;
int main()
{
    vector<int> nums(10);
    int oddPos = 0, evenPos = 9;
    for (int i = 0; i < 10; i++)
    {
        int temp;
        cin >> temp;
        if ((temp & 1) == 1)
        {
            nums[oddPos++] = temp;
        }
        else
        {
            nums[evenPos--] = temp;
        }
    }
    sort(nums.begin(), nums.begin() + oddPos, greater<int>());
    sort(nums.begin() + oddPos, nums.end(), less<int>());
    for (int i = 0; i < 10; i++)
    {
        cout << nums[i] << " ";
    }
    return 0;
}
```

### 字串计算

给出一个 01 字符串（长度不超过 100），求其每一个子串出现的次数。

输入包含多行，每行一个字符串。
对每个字符串，输出它所有出现次数在 1 次以上的子串和这个子串出现的次数，输出按字典序排序。

[网址](https://www.nowcoder.com/practice/bcad754c91a54994be31a239996e7c11?tpId=40&tqId=21399&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <string>
#include <map>
using namespace std;
int main()
{
    string str;
    while (cin >> str && str != " ")
    {
        map<string, int> substring_map;
        for (int i = 1; i <= str.length(); i++)
        {
            for (int j = 0; j < i; j++)
            {
                substring_map[str.substr(j, i - j)]++;
            }
        }

        for (map<string, int>::iterator it = substring_map.begin(); it != substring_map.end(); it++)
        {
            if (it->second > 1)
            {
                cout << it->first << " " << it->second << endl;
            }
        }
    }
    return 0;
}
```

### 采药问题

辰辰是个很有潜能、天资聪颖的孩子，他的梦想是称为世界上最伟大的医师。 为此，他想拜附近最有威望的医师为师。医师为了判断他的资质，给他出了一个难题。 医师把他带到个到处都是草药的山洞里对他说： “孩子，这个山洞里有一些不同的草药，采每一株都需要一些时间，每一株也有它自身的价值。 我会给你一段时间，在这段时间里，你可以采到一些草药。如果你是一个聪明的孩子，你应该可以让采到的草药的总价值最大。” 如果你是辰辰，你能完成这个任务吗？

输入的第一行有两个整数 T（1 <= T <= 1000）和 M（1 <= M <= 100），T 代表总共能够用来采药的时间，M 代表山洞里的草药的数目。接下来的 M 行每行包括两个在 1 到 100 之间（包括 1 和 100）的的整数，分别表示采摘某株草药的时间和这株草药的价值。

[网址](https://www.nowcoder.com/practice/d7c03b114f0541dd8e32ce9987326c16?tpId=40&tqId=21406&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<math.h>
using namespace std;
int main(){
    // 总时间、药草数量
    int totalTime,herbCnt;
    while(cin >> totalTime >> herbCnt && totalTime!= 0 && herbCnt != 0){

        // dp[i]表示时间为i时能采药的最大价值
        vector<int> dp(totalTime + 1,0);

        // 分别表示每种草药的时间和价值
        int singleTime,singleValue;
        for(int i = 0;i < herbCnt;i++){
            cin >> singleTime >> singleValue;

            // 计算每一时间采药的最大价值
            for(int j = totalTime;j >= 0;j--){
                if(j >= singleTime){
                    dp[j] = max(dp[j],dp[j - singleTime] + singleValue);
                }
            }
        }
        cout << dp[totalTime] << endl;
    }
    return 0;
}
```

### 合唱队形

N 位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的 K 位同学不交换位置就能排成合唱队形。 合唱队形是指这样的一种队形：设 K 位同学从左到右依次编号为 1, 2, …, K，他们的身高分别为 T1, T2, …, TK， 则他们的身高满足 T1 < T2 < … < Ti , Ti > Ti+1 > … > TK (1 <= i <= K)。 你的任务是，已知所有 N 位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。

输入的第一行是一个整数 N（2 <= N <= 100），表示同学的总数。
第一行有 n 个整数，用空格分隔，第 i 个整数 Ti（130 <= Ti <= 230）是第 i 位同学的身高（厘米）。

[网址](https://www.nowcoder.com/practice/cf209ca9ac994015b8caf5bf2cae5c98?tpId=40&tqId=21404&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<math.h>
using namespace std;
int main(){

    // 学生数量
    int studentCnt;
    while(cin >> studentCnt && studentCnt != 0){
        // 学生身高
        vector<int> heights(studentCnt,0);
        for(int i = 0;i < studentCnt;i++){
            cin >> heights[i];
        }

        // leftDp[i]表示i之前的最大递增序列
        // rightDp[i]表示之后的最大递减序列
        vector<int> leftDp(studentCnt,1);
        vector<int> rightDp(studentCnt,1);

        for(int i = 0;i < studentCnt;i++){
            for(int j = 0;j < i;j++){

                // 计算i之前的最长递增序列
                if(heights[j] < heights[i]){
                    leftDp[i] = max(leftDp[i],leftDp[j] + 1);
                }

                // 计算i之后的最长递减序列
                // x位置在y前面
                int x = studentCnt - i - 1;
                int y = studentCnt - j - 1;
                if(heights[x] > heights[y]){
                    rightDp[x] = max(rightDp[x],rightDp[y] + 1);
                }
            }
        }

        // 计算每一个位置前面的递增序列长度和后面的递减序列长度之和
        // 值最大时出列的学生数最少
        int maxSum = leftDp[0] + rightDp[0];
        for(int k = 1;k < studentCnt;k++){
            if(leftDp[k] + rightDp[k] > maxSum){
                maxSum = leftDp[k] + rightDp[k];
            }
        }
        // studentCnt - (maxSum - 1) 减掉1是因为位置i即在递增序列，也在递减序列，加了两次
        cout << studentCnt - maxSum + 1 << endl;
    }
}
```

### 最大子矩阵

已知矩阵的大小定义为矩阵中所有元素的和。给定一个矩阵，你的任务是找到最大的非空(大小至少是 1 _ 1)子矩阵。 比如，如下 4 _ 4 的矩阵 0 -2 -7 0 9 2 -6 2 -4 1 -4 1 -1 8 0 -2 的最大子矩阵是 9 2 -4 1 -1 8 这个子矩阵的大小是 15。

输入是一个 N \* N 的矩阵。输入的第一行给出 N (0 < N <= 100)。
再后面的若干行中，依次（首先从左到右给出第一行的 N 个整数，再从左到右给出第二行的 N 个整数……）给出矩阵中的 N2 个整数，整数之间由空白字符分隔（空格或者空行）。
已知矩阵中整数的范围都在[-127, 127]。

[网址](https://www.nowcoder.com/practice/a5a0b05f0505406ca837a3a76a5419b3?tpId=40&tqId=21394&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<math.h>
#include<limits.h>
using namespace std;

// 获取某一行元素的最大子数组最大和
int getRowMaxSum(vector<int> &vec){
    // dp[i]表示以第i个字符结尾的最大和
    vector<int> dp(vec.size());
    dp[0] = vec[0];
    int maxSum = dp[0];
    for(int i = 1;i < vec.size();i++){
        dp[i] = max(dp[i - 1] + vec[i],vec[i]);
        if(dp[i] > maxSum){
            maxSum = dp[i];
        }
    }
    return maxSum;
}
int getMaxSubMatrix(vector<vector<int> > &vec){
    int maxSum = INT_MIN;

    for(int i = 0;i < vec.size();i++){
        vector<int> temp(vec.size(),0);
        for(int j = i;j < vec.size();j++){
            for(int k = 0;k < vec[i].size();k++){
                temp[k] += vec[j][k];
            }

            // 将上面的行和下面的行对应位置相加，求新行的最大子数组和
            maxSum = max(maxSum,getRowMaxSum(temp));
        }
    }
    return maxSum;
}
int main(){
    int rows;
    while(cin >> rows){
        vector<vector<int> > matrix(rows,vector<int>(rows,0));
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < rows;j++){
                cin >> matrix[i][j];
            }
        }
        cout << getMaxSubMatrix(matrix) << endl;
    }
    return 0;
}
```

### 点菜问题

北大网络实验室经常有活动需要叫外卖，但是每次叫外卖的报销经费的总额最大为 C 元，有 N 种菜可以点，经过长时间的点菜，网络实验室对于每种菜 i 都有一个量化的评价分数（表示这个菜可口程度），为 Vi，每种菜的价格为 Pi, 问如何选择各种菜，使得在报销额度范围内能使点到的菜的总评价分数最大。 注意：由于需要营养多样化，每种菜只能点一次。

输入的第一行有两个整数 C（1 <= C <= 1000）和 N（1 <= N <= 100），C 代表总共能够报销的额度，N>代表能点菜的数目。接下来的 N 行每行包括两个在 1 到 100 之间（包括 1 和 100）的的整数，分别表示菜的>价格和菜的评价分数。

[网址](https://www.nowcoder.com/practice/b44f5be34a9143aa84c478d79401e22a?tpId=40&tqId=21397&tPage=4&rp=4&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include<iostream>
#include<vector>
#include<math.h>
using namespace std;
int main(){
    int totalMoney,dishCnt;
    while(cin >> totalMoney >> dishCnt){
        // dp[i]表示钱数为i时的最大评价
        vector<int> dp(totalMoney + 1,0);

        // 一次加一种菜，每加入一种菜就计算加入后
        // 不同钱数所买菜的评价的最大值
        for(int i = 1;i <= dishCnt;i++){
            // 每种菜的价格和评价
            int price,comment;
            cin >> price >> comment;
            for(int j = totalMoney;j >= 0;j--){
                // 剩余钱数能买该种菜
                if(j >= price){
                    dp[j] = max(dp[j],dp[j - price] + comment);
                }
            }
        }
        cout << dp[totalMoney] << endl;
    }
    return 0;
}
```

### 拦截导弹

某国为了防御敌国的导弹袭击，开发出一种导弹拦截系统。但是这种导弹拦截系统有一个缺陷：虽然它的第一发炮弹能够到达任意的高度，但是以后每一发炮弹都不能高于前一发的高度。某天，雷达捕捉到敌国的导弹来袭，并观测到导弹依次飞来的高度，请计算这套系统最多能拦截多少导弹。拦截来袭导弹时，必须按来袭导弹袭击的时间顺序，不允许先拦截后面的导弹，再拦截前面的导弹。

每组输入有两行，
第一行，输入雷达捕捉到的敌国导弹的数量 k（k<=25），
第二行，输入 k 个正整数，表示 k 枚导弹的高度，按来袭导弹的袭击时间顺序给出，以空格分隔。

每组输出只有一行，包含一个整数，表示最多能拦截多少枚导弹。

[网址](https://www.nowcoder.com/practice/dad3aa23d74b4aaea0749042bba2358a?tpId=40&tqId=21408&tPage=1&rp=1&ru=/ta/kaoyan&qru=/ta/kaoyan/question-ranking)

```cpp
#include <iostream>
#include <vector>
#include<math.h>
using namespace std;
int main(){
    int missileCnt;
    cin >> missileCnt;
    vector<int> heights(missileCnt,0);
    for(int i = 0;i < missileCnt;i++){
        cin >> heights[i];
    }

    // dp[i] 表示前i个导弹所能拦截的最大数量
    vector<int> dp(missileCnt,1);
    int maxIndex = 0;
    for(int i = 0;i < missileCnt;i++){
        for(int j = 0;j < i;j++){
            if(heights[j] >= heights[i]){
                dp[i] = max(dp[i],dp[j] + 1);
            }
            if(dp[i] > dp[maxIndex]){
                maxIndex = i;
            }
        }
    }
    cout << dp[maxIndex] << endl;
}
```


