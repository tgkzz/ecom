PROTOC = protoc

PROTO_FILES := $(shell find inner -name "*.proto")

.PHONY: proto
proto:
	@echo ">> Generating proto files into innerpb (without inner/)..."
	@for file in $(PROTO_FILES); do \
		echo "Generating $$file"; \
		$(PROTOC) -I=inner \
			--go_out=innerpb --go_opt=paths=source_relative \
			--go-grpc_out=innerpb --go-grpc_opt=paths=source_relative \
			$$file; \
	done
	@echo ">> Done!"
