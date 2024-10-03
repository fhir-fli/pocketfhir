#!/bin/bash

# Navigate to the root directory of the project
cd "$(dirname "$0")/.."

# Clean up previous builds
rm -f pocketfhir_server

# Initialize Go modules in the root directory if go.mod does not exist
if [ ! -f go.mod ]; then
    go mod init pocketfhir
fi

# Tidy up Go modules in the root directory
go mod tidy

# Set the Android target platform
ANDROID_ARCH=arm64
ANDROID_SDK_ROOT=$HOME/Android/Sdk

# Set the Android NDK toolchain path (use the correct NDK version)
NDK_TOOLCHAIN=$ANDROID_SDK_ROOT/ndk/28.0.12433566/toolchains/llvm/prebuilt/linux-x86_64

# Path for Android cross-compilation
export GOARCH=$ANDROID_ARCH
export GOOS=android
export CGO_ENABLED=1

# Set cross-compiler to target Android with the right version
export CC=$NDK_TOOLCHAIN/bin/aarch64-linux-android21-clang
export CXX=$NDK_TOOLCHAIN/bin/aarch64-linux-android21-clang++

# Build the PocketFHIR server for Android
echo "Building PocketFHIR for Android platform..."
if ! go build -o pocketfhir_server_android ./main.go; then
    echo "Android build failed!"
    exit 1
fi
echo "PocketFHIR Android build completed."
