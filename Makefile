# Gunakan bash untuk perintah shell. Ini akan berfungsi jika dijalankan dari Git Bash atau WSL.
SHELL := /bin/bash

# Set GOPRIVATE jika diperlukan untuk modul pribadi Anda
export GOPRIVATE=https://github.com/tovankamil

# --- Configuration Variables (WAJIB ANDA SESUAIKAN DENGAN LOKASI ASLI DI KOMPUTER ANDA!) ---
# Definisikan jalur modul untuk file Protobuf Anda.
# Contoh: Jika wallet.proto Anda ada di lib/protos/v1/wallet/wallet.proto, MODULE harus 'v1/wallet'.
MODULE ?= v1/wallet

# Path ke direktori include Protobuf compiler (protoc) di Windows.
# Ini harus menunjuk ke tempat file .proto standar protoc berada (misalnya, google/protobuf/timestamp.proto).
# CONTOH UMUM: C:/Program Files/Google/Protobuf/include
# SILAKAN VERIFIKASI JALUR INI DI KOMPUTER ANDA DAN GANTI NILAI DI BAWAH INI!
# >>>>>> HAPUS TANDA KUTIP DARI DEFINISI VARIABEL INI <<<<<<
PROTOC_STANDARD_INCLUDE_PATH := C:/Program Files/protoc/include

# Path ke Google's well-known types (googleapis).
# Ini harus menunjuk ke ROOT dari repositori googleapis yang Anda kloning.
# CONTOH: D:/Project/fastcampus/practice/googleapis
# SILAKAN VERIFIKASI JALUR INI DI KOMPUTER ANDA DAN GANTI NILAI DI BAWAH INI!
# >>>>>> HAPUS TANDA KUTIP DARI DEFINISI VARIABEL INI <<<<<<
GOOGLEAPIS_PATH := D:/Project/fastcampus/practice/googleapis

# Secara dinamis temukan jalur modul grpc-gateway dari cache modul Go.
# Jalur ini berisi protoc-gen-openapiv2/options/openapiv2.proto dan annotations.proto
# VERIFIKASI VERSI (@v2.27.1) INI DENGAN NAMA FOLDER DI GOPATH/pkg/mod ANDA!
# >>>>>> HAPUS TANDA KUTIP DARI DEFINISI VARIABEL INI <<<<<<
GRPC_GATEWAY_MOD_PATH := $(shell go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.27.1

# Temukan file .proto dalam jalur modul yang ditentukan, tidak termasuk direktori vendor.
# Ini menggunakan 'find' Unix, jadi jalankan 'make' di Git Bash atau WSL.
API_PROTO_FILES=$(shell find lib/protos/$(MODULE) -name "*.proto" -not -path '*/vendor/*')

# Temukan file JSON OpenAPI yang dihasilkan.
# Mengasumsikan JSON OpenAPI dihasilkan langsung mencerminkan struktur proto di bawah lib/protos/openapiv2.
API_PROTO_CLIENT=$(shell find lib/protos/openapiv2/$(MODULE) -name "*.json" -not -path '*/vendor/*' -not -path '*/config/*')


# --- Makefile Targets ---

.PHONY: init
init:
	@echo "Installing Protobuf plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@echo "Plugins installed. Pastikan GOPATH/bin ada di PATH sistem Anda dan Protoc terinstal."

.PHONY: print print-client
print:
	@echo "API_PROTO_FILES: $(API_PROTO_FILES)"

print-client:
	@echo "API_PROTO_CLIENT: $(API_PROTO_CLIENT)"

.PHONY: debug-paths
debug-paths:
	@echo "--- DEBUGGING PROTOC INCLUDE PATHS ---"
	@echo "PROTOC_STANDARD_INCLUDE_PATH: $(PROTOC_STANDARD_INCLUDE_PATH)"
	@echo "GOOGLEAPIS_PATH: $(GOOGLEAPIS_PATH)"
	@echo "GRPC_GATEWAY_MOD_PATH: $(GRPC_GATEWAY_MOD_PATH)"
	@echo "--- END DEBUGGING ---"

.PHONY: generate
generate: debug-paths # Menjalankan target debug-paths sebelum generate
	@echo "Generating Protobuf and OpenAPI files for $(MODULE)..."
	# Pastikan direktori output OpenAPI JSONs ada sebelum menjalankan protoc
	mkdir -p ./lib/protos/openapiv2/$(MODULE)
	protoc \
		-I . \
		-I "$(PROTOC_STANDARD_INCLUDE_PATH)" \
		-I "$(GOOGLEAPIS_PATH)" \
		-I "$(GRPC_GATEWAY_MOD_PATH)" \
		--go_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--openapiv2_out ./lib/protos/openapiv2 --openapiv2_opt logtostderr=true \
		$(API_PROTO_FILES)
	@echo "Protobuf dan OpenAPI generation complete."

.PHONY: generate-js-client
generate-js-client:
	@echo "Generating JavaScript client for $(MODULE)..."
	# Pastikan openapi-generator-cli terinstal (misalnya, via npm: npm install @openapitools/openapi-generator-cli -g)
	openapi-generator generate -g javascript --additional-properties=usePromises=true -i $(API_PROTO_CLIENT) -o ./lib/protos/client/js/$(MODULE)
	@echo "JavaScript client generation complete."

.PHONY: generate-go-client
generate-go-client:
	@echo "Generating Go client for $(MODULE)..."
	# Pastikan openapi-generator-cli terinstal
	@openapi-generator generate -g go --additional-properties=packageName=repository -i $(API_PROTO_CLIENT) -o ./lib/protos/client/go/$(MODULE)
	@echo "Go client generation complete."

.PHONY: test
test:
	@echo "=== MENJALANKAN TES ==="
	@go test -v $$(go list ./... | grep -v /vendor/ | grep -v /graph/ | grep -v /cmd/)
	@echo "=== TES SELESAI ==="

.PHONY: lint
lint:
	@echo "Menginstal golangci-lint jika tidak ditemukan: https://golangci-lint.run/usage/install/"
	@golangci-lint run ./... > .golint.txt
	@echo "Linting selesai. Hasil ada di .golint.txt"

.PHONY: clean
clean:
	@echo "Membersihkan file yang dihasilkan..."
	rm -rf lib/protos/**/*.pb.go
	rm -rf lib/protos/**/*.pb.gw.go
	rm -rf lib/protos/**/*.pb.http.go
	rm -rf lib/protos/openapiv2/**/*.json
	rm -rf lib/protos/client/js/$(MODULE)
	rm -rf lib/protos/client/go/$(MODULE)
	@echo "Pembersihan selesai."