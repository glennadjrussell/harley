package network


import (
    "testing"
)

func TestGetNetworkDevices(t *testing.T) {
    scanner := &Scanner{} // Initialize the scanner

    // Call the function
    devicesPtr, err := scanner.GetNetworkDevices()

    // Check if an error occurred (it shouldn't in this case)
    if err != nil {
        t.Fatalf("Expected no error, but got: %v", err)
    }

    // Ensure that devicesPtr is not nil
    if devicesPtr == nil {
        t.Fatalf("Expected a non-nil pointer to a slice, but got nil")
    }

    // Dereference the pointer to get the slice
    devices := *devicesPtr

    t.Log(devices)
}

