project_name := "msg"
project_root := justfile_directory()
coverage_data := "coverage.out"
coverage_html := "coverage.html"
goreleaser_dist := "dist"
commit_sha := `git rev-parse HEAD`

# By default print the list of recipes
_default:
    @just --list

# Tidy up dependencies in go.mod and go.sum
tidy:
    go mod tidy

# Run go fmt on all project files
fmt:
    gofumpt -extra -s -w .

# Run all project unit tests
test: fmt
    go test -race ./...

# Lint the project and auto-fix errors if possible
lint: fmt
    golangci-lint run --fix

# Calculate test coverage and render the html
cover:
    go test -race -cover -coverprofile={{ coverage_data }} ./...
    go tool cover -html={{ coverage_data }} -o {{ coverage_html }}
    open {{ coverage_html }}

# Remove build artifacts and other project clutter
clean:
    go clean ./...
    rm -rf {{ coverage_data }} {{ coverage_html }} {{ goreleaser_dist }}

# Run unit tests and linting in one go
check: test lint

# Run all recipes (other than clean) in a sensible order
all: test lint cover
