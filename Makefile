# running all service 
run:
	@echo "Starting Service Product..."
	cd service_product && go run main.go &

	@echo "Starting Service Order..."
	cd service_order && go run main.go &

	@echo "Starting Service Cart..."
	cd service_cart && go run main.go &

	@echo "Starting API Gateway..."
	cd api_gateway && go run main.go &

	# Wait for all processes to continue running
	wait

# Stop all running processes
stop:
	@echo "Stopping all services..."
	@-PID1=$$(lsof -t -i :3000); [ -n "$$PID1" ] && kill -9 $$PID1 || echo "No process on port 3000"
	@-PID1=$$(lsof -t -i :5001); [ -n "$$PID1" ] && kill -9 $$PID1 || echo "No process on port 5001"
	@-PID1=$$(lsof -t -i :5002); [ -n "$$PID1" ] && kill -9 $$PID1 || echo "No process on port 5002"
	@-PID1=$$(lsof -t -i :5003); [ -n "$$PID1" ] && kill -9 $$PID1 || echo "No process on port 5003"
	@echo "All services stopped."



# Restart service
restart: stop run

check:
	@echo "Checking running services..."
	@lsof -i :5001 || echo "Service product is not running"
	@lsof -i :5002 || echo "Service order is not running"
	@lsof -i :5003 || echo "Service cart is not running"
	@lsof -i :3000 || echo "API Gateway is not running"

