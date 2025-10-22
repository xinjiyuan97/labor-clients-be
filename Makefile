.PHONY: build generate_all generate_openapi # 声明为伪目标，强制每次执行命令

build:
	$(eval RUN_NAME=hertz_service)
	$(eval OUTPUT_DIR=output)
	mkdir -p $(OUTPUT_DIR)
	cp script/* $(OUTPUT_DIR) 2>/dev/null
	cp -r conf $(OUTPUT_DIR)/ 2>/dev/null
	chmod +x $(OUTPUT_DIR)/bootstrap.sh
	go build -o $(OUTPUT_DIR)/bin/$(RUN_NAME)


generate_all:
	echo "update idls"
	hz update -idl idls/common.thrift --handler_by_method --handler_dir biz/handler/common  
	hz update -idl idls/auth.thrift --handler_by_method --handler_dir biz/handler/auth  
	hz update -idl idls/schedule.thrift --handler_by_method --handler_dir biz/handler/schedule  
	hz update -idl idls/job.thrift --handler_by_method --handler_dir biz/handler/job  
	hz update -idl idls/job_application.thrift --handler_by_method --handler_dir biz/handler/job_application  
	hz update -idl idls/message.thrift --handler_by_method --handler_dir biz/handler/message  
	hz update -idl idls/user.thrift --handler_by_method --handler_dir biz/handler/user  
	hz update -idl idls/community.thrift --handler_by_method --handler_dir biz/handler/community  
	hz update -idl idls/attendance.thrift --handler_by_method --handler_dir biz/handler/attendance  
	hz update -idl idls/review.thrift --handler_by_method --handler_dir biz/handler/review  
	hz update -idl idls/payment.thrift --handler_by_method --handler_dir biz/handler/payment  
	hz update -idl idls/system.thrift --handler_by_method --handler_dir biz/handler/system  
	hz update -idl idls/upload.thrift --handler_by_method --handler_dir biz/handler/upload  
	hz update -idl idls/admin.thrift --handler_by_method --handler_dir biz/handler/admin  

generate_openapi:
	python3 generate_openapi.py idls