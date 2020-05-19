#include <IOKit/pwr_mgt/IOPMLib.h>

CFStringRef reasonForActivity = CFSTR("Logging time in Go");
IOPMAssertionID assertionID;
IOReturn success;

bool createAssertionSuccess() {
    success = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoIdleSleep, kIOPMAssertionLevelOn,
                                                   reasonForActivity, &assertionID);
    if (success == kIOReturnSuccess) return true;
    return false;
}

void releaseAssertionSuccess() {
    success = IOPMAssertionRelease(assertionID);
}
