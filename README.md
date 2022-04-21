# Github刷小绿点的原理

1. 原理
github小绿点仅以commit的时间为准备，而不是push的时间，所以，当删除项目的时候，以前的某些小绿点会消失，当新增commit的时候，如果commit的时间为过去，则小绿点显示在过去的时间。

2. 思路
在git工具中，commit命令可以指定commit时间（如果不指定，则为系统当前时间，这也是常规的情况），新建若干commit并指定过去的时间，很快就能刷满小绿点。

3. 方法
命令格式：--date
时间格式：[参考这里](https://stackoverflow.com/questions/19742345/what-is-the-format-for-date-parameter-of-git-commit)
举例：git commit -m "Test" --date=format:relative:5.hours.ago 