.PHONY: run
SWIFT_SRC_ROOT := swift/submit-transaction
COURSENAME_SWIFT := $(SWIFT_SRC_ROOT)/Sources/submit-transaction/coursename.out.swift
SWIFT_EXECUTABLE := $(SWIFT_SRC_ROOT)/.build/x86_64-apple-macosx/release/submit-transaction

prepare:
	echo '// Generated code. DO NOT EDIT.' > $(COURSENAME_SWIFT)
	./generate-coursename.sh >> $(COURSENAME_SWIFT)
	(cd $(SWIFT_SRC_ROOT) && swift build -c release)

run:
	go run main.go | $(SWIFT_EXECUTABLE)
