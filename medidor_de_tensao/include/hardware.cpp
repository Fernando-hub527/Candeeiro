#ifndef START_HARDWARE
#define START_HARDWARE

#include "Arduino.h"
#include "LittleFS.h"
#include "ACS712.h"

ACS712 sensor(ACS712_20A, 4);

void startHardware(){
    if(LittleFS.begin(true)){
        
    }
}

float readAcs712(){
    return sensor.getCurrentAC();
}


#endif