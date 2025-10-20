package job

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/job"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetJobListLogic 获取工作列表业务逻辑
func GetJobListLogic(req *job.GetJobListReq) (*job.GetJobListResp, error) {
	// 设置默认分页参数
	page := 1
	limit := 10
	if req.PageReq != nil {
		if req.PageReq.Page > 0 {
			page = int(req.PageReq.Page)
		}
		if req.PageReq.Limit > 0 {
			limit = int(req.PageReq.Limit)
		}
	}

	offset := (page - 1) * limit

	// 根据条件获取工作列表
	var jobs []*models.Job
	var total int64
	var err error

	if req.CategoryID > 0 {
		// 根据分类获取
		jobs, err = mysql.GetJobsByCategory(nil, int64(req.CategoryID), offset, limit)
		if err != nil {
			utils.Errorf("根据分类获取工作列表失败: %v", err)
			return &job.GetJobListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountJobsByCategory(nil, int64(req.CategoryID))
	} else if req.JobType != "" {
		// 根据类型获取
		jobs, err = mysql.GetJobsByType(nil, req.JobType, offset, limit)
		if err != nil {
			utils.Errorf("根据类型获取工作列表失败: %v", err)
			return &job.GetJobListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountJobsByType(nil, req.JobType)
	} else if req.SalaryMin > 0 || req.SalaryMax > 0 {
		// 根据薪资范围获取
		jobs, err = mysql.GetJobsBySalaryRange(nil, req.SalaryMin, req.SalaryMax, offset, limit)
		if err != nil {
			utils.Errorf("根据薪资范围获取工作列表失败: %v", err)
			return &job.GetJobListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountJobs(nil)
	} else {
		// 获取所有工作
		jobs, err = mysql.GetJobs(nil, offset, limit)
		if err != nil {
			utils.Errorf("获取工作列表失败: %v", err)
			return &job.GetJobListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountJobs(nil)
	}

	if err != nil {
		utils.Errorf("获取工作总数失败: %v", err)
		return &job.GetJobListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建工作信息
	var jobInfos []*common.JobInfo
	for _, jobModel := range jobs {
		jobInfo := &common.JobInfo{
			JobID:          jobModel.ID,
			EmployerID:     jobModel.EmployerID,
			BrandID:        jobModel.BrandID,
			CategoryID:     jobModel.CategoryID,
			Title:          jobModel.Title,
			JobType:        jobModel.JobType,
			Description:    jobModel.Description,
			Salary:         jobModel.Salary.InexactFloat64(),
			SalaryUnit:     jobModel.SalaryUnit,
			Location:       jobModel.Location,
			Latitude:       jobModel.Latitude.InexactFloat64(),
			Longitude:      jobModel.Longitude.InexactFloat64(),
			Requirements:   jobModel.Requirements,
			Benefits:       jobModel.Benefits,
			StartTime:      jobModel.StartTime.Format(time.RFC3339),
			EndTime:        jobModel.EndTime.Format(time.RFC3339),
			Status:         jobModel.Status,
			MaxApplicants:  int32(jobModel.MaxApplicants),
			ApplicantCount: int32(jobModel.ApplicantCount),
		}
		jobInfos = append(jobInfos, jobInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &job.GetJobListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取工作列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Jobs:     jobInfos,
	}, nil
}
