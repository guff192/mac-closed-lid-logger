#include <IOKit/pwr_mgt/IOPMLib.h>

CFStringRef reasonForActivity_Idle = CFSTR("Logging time in Go");
CFStringRef reasonForActivity_Display = CFSTR("Logging time in Go");

IOPMAssertionID assertionID_Idle, assertionID_Display;
IOReturn success_Idle, success_Display;

bool createAssertionSuccess() {
    success_Idle = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoIdleSleep, kIOPMAssertionLevelOn,
                                                   reasonForActivity_Idle, &assertionID_Idle);

    success_Display = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoDisplaySleep, kIOPMAssertionLevelOn,
                                                       reasonForActivity_Display, &assertionID_Display);

    if (success_Display == kIOReturnSuccess && success_Idle == kIOReturnSuccess) return true;
    return false;
}

void releaseAssertionSuccess() {
    success_Display = IOPMAssertionRelease(assertionID_Display);
    success_Idle = IOPMAssertionRelease(assertionID_Idle);
}
