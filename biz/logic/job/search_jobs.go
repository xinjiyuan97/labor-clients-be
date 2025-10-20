package job

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/job"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// SearchJobsLogic 搜索工作业务逻辑
func SearchJobsLogic(req *job.SearchJobsReq) (*job.SearchJobsResp, error) {
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

	// 搜索工作
	var jobs []*models.Job
	var total int64
	var err error

	if req.Location != "" || req.SalaryMin > 0 || req.SalaryMax > 0 {
		// 带过滤条件的搜索
		jobs, err = mysql.SearchJobsWithFilters(nil, req.Keyword, req.Location, req.SalaryMin, req.SalaryMax, offset, limit)
		if err != nil {
			utils.Errorf("带过滤条件搜索工作失败: %v", err)
			return &job.SearchJobsResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountSearchJobsWithFilters(nil, req.Keyword, req.Location, req.SalaryMin, req.SalaryMax)
	} else {
		// 简单搜索
		jobs, err = mysql.SearchJobs(nil, req.Keyword, offset, limit)
		if err != nil {
			utils.Errorf("搜索工作失败: %v", err)
			return &job.SearchJobsResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountSearchJobs(nil, req.Keyword)
	}

	if err != nil {
		utils.Errorf("获取搜索结果总数失败: %v", err)
		return &job.SearchJobsResp{
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

	return &job.SearchJobsResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "搜索工作成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Jobs:     jobInfos,
	}, nil
}
