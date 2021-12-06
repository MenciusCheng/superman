package main

import (
	"fmt"
	"github.com/MenciusCheng/superman/util/gendbinfo/tools"
	"github.com/MenciusCheng/superman/util/summoner/leetcode"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	desc := `
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
`

	url := `
https://leetcode-cn.com/problems/two-sum/
`

	cal := `
func twoSum(nums []int, target int) []int {

}
`
	question := "q1"

	month := "m202112"

	subject, err := leetcode.NewSubject(desc, url, cal)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("CheckInSubject").Parse(leetcode.CheckInSubject)
	if err != nil {
		panic(err)
	}

	fileName := "main.go"
	directory := fmt.Sprintf("/Users/chengmengwei/goProject/algorithm-code/leetcode/%s/%s", month, question)
	path := fmt.Sprintf("%s%s%s", directory, string(os.PathSeparator), fileName)

	if err := tools.BuildDir(path); err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, subject)
	if err != nil {
		panic(err)
	}

	cmd, _ := exec.Command("gofmt", "-l", "-w", path).Output()
	fmt.Println(string(cmd))
}
