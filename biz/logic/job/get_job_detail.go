package job

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/job"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetJobDetailLogic 获取工作详情业务逻辑
func GetJobDetailLogic(req *job.GetJobDetailReq) (*job.GetJobDetailResp, error) {
	// 获取工作详情
	jobModel, err := mysql.GetJobByID(nil, req.JobID)
	if err != nil {
		utils.Errorf("获取工作详情失败: %v", err)
		return &job.GetJobDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if jobModel == nil {
		return &job.GetJobDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "工作不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建工作信息
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

	return &job.GetJobDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取工作详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Job: jobInfo,
	}, nil
}
