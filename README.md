# zhaiker

zhaiker.cn API

## Usage

```go
import "github.com/caiguanhao/zhaiker"

ctx := context.Background()
client := zhaiker.NewClient(zhaiker.MustGetKey(ctx, "USERNAME", "PASSWORD"))

var gyms []zhaiker.Gym
client.MustRequest(ctx, zhaiker.API_GET_GYMS, zhaiker.Params("pageSize", 100), &gyms, "DATA.*")

var exams []zhaiker.Exam
client.MustRequest(ctx, zhaiker.API_GET_EXAMS, zhaiker.Params("gymId", gyms[0].Id), &exams, "DATA.*")

var exam zhaiker.ExamWithStandards
client.MustRequest(ctx, zhaiker.API_GET_EXAM, zhaiker.Params("gymId", gyms[0].Id, "id", exams[0].Id), &exam)
```
