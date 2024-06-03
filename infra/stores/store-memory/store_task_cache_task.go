package storememory

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/domain/task"
)

type cacheTask map[task.PrimaryKeyTask]task.TaskInfo

func (cache cacheTask) String() string {
	result := []string{
		fmt.Sprintf("cache Tasks: %d", len(cache)),
	}

	for pk, taskInfo := range cache {
		result = append(result,
			fmt.Sprintf("id: %v, name: %s", pk, taskInfo.Name),
		)
	}

	return strings.Join(result, "\n")
}
