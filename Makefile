BIN_DIR := bin

CMD := $(shell find cmd -maxdepth 1 -type d ! -path cmd | sed 's|cmd/||')

.PHONY: all clean build

all: build

build: | $(BIN_DIR)
	@echo "Building binaries: $(CMD)"
	@for cmd in $(CMD); do \
		echo " → $$cmd"; \
		go build -o $(BIN_DIR)/$$cmd ./cmd/$$cmd; \
		done

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)
