#ifndef TIME_CPP
#define TIME_CPP

#include "Arduino.h"
#include <NTPClient.h>
#include <WiFiUdp.h>

WiFiUDP udp;
NTPClient ntp(udp, "a.st1.ntp.br", -3 * 3600, 60000);

unsigned long getCurrentTime(){
    return ntp.getEpochTime();
}

void startTime(){
    ntp.begin();               
    ntp.forceUpdate();
}

#endif